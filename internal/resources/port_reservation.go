package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewPortReservationResource creates the resource for
// managing port reservations.
func NewPortReservationResource() *schema.Resource {
	s := schemas.Combine(
		schemas.PortReservationSchema(),
		map[string]*schema.Schema{
			"port_num": schemas.DataSourceQueryInt(
				"When the `port_num` is provided, the resource will reference an already allocated port reservation. Starting at 1."),
		})
	return &schema.Resource{
		Description: "Use the `ixapi_port_reservation` resource to manage port reservations for a connection. You can create a new port reservation by only passing the `connection` to the resource, or you can reference an already allocated port reservation, which was created when the connection was allocated, by specifying a `port_num` (starting at 1).",
		Schema:      s,

		CreateContext: crud.Create(portReservationCreate),
		ReadContext:   crud.Read(portReservationRead),
		UpdateContext: crud.Update(portReservationUpdate),
		DeleteContext: crud.Delete(portReservationDelete),
	}
}

// Fetch connection port reservations. This will poll the conenction
// until the number of port_reservations will match `port_quantity`.
func fetchConnectionPortReservations(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]string, error) {
	connection := res.Get("connection").(string)
	for {
		// Check if the context is still valid
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		conn, err := api.ConnectionsRead(ctx, connection)
		if err != nil {
			return nil, err
		}
		if len(conn.PortReservations) >= conn.PortQuantity {
			return conn.PortReservations, nil
		}
		time.Sleep(30 * time.Second)
	}
}

// Make a port reservation create request from resource data
func portReservationRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.PortReservationRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.PortReservationRequest{
		PurchaseOrder:        res.GetStringOpt("purchase_order"),
		ContractRef:          res.GetStringOpt("contract_ref"),
		ExternalRef:          res.GetStringOpt("external_ref"),
		SubscriberSideDemarc: res.GetStringOpt("subscriber_side_demarc"),
		ConnectingParty:      res.GetStringOpt("connecting_party"),
		CrossConnectID:       res.GetStringOpt("cross_connect_id"),
		Connection:           res.GetString("connection"),
	}
	return req, nil
}

// Make a port reservation patch from resource data
func portReservationPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.PortReservationPatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.PortReservationPatch{}

	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}

	// It feels a bit off, that the following fields are
	// updateable, but the API specs allow this.
	// I guess the API will respond with an error in most of the times here.
	if res.HasChange("purchase_order") {
		req.PurchaseOrder = res.GetStringOpt("purchase_order")
	}
	if res.HasChange("contract_ref") {
		req.ContractRef = res.GetStringOpt("contract_ref")
	}
	if res.HasChange("subscriber_side_demarc") {
		req.SubscriberSideDemarc = res.GetStringOpt("subscriber_side_demarc")
	}
	if res.HasChange("connecting_party") {
		req.ConnectingParty = res.GetStringOpt("connecting_party")
	}
	if res.HasChange("cross_connect_id") {
		req.CrossConnectID = res.GetStringOpt("cross_connect_id")
	}

	return req, nil
}

// Create a port reservation resource
func portReservationCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	// A connection is required, if the port_index is given
	// use the connection.port_reservation[i].
	portNum, hasPortNum := res.GetOk("port_num")
	if hasPortNum {
		reservationIds, err := fetchConnectionPortReservations(ctx, res, api)
		if err != nil {
			return err
		}
		idx := portNum.(int) - 1
		if idx < 0 {
			return fmt.Errorf("`port_num` must be >= 1")
		}
		if idx >= len(reservationIds) {
			return fmt.Errorf("`port_num` may not be greater then the number of allocated port reservations on the connection")
		}
		id := reservationIds[idx]
		res.SetId(id)
		return portReservationRead(ctx, res, api)
	}

	// Otherwise: Create a new port reservation for a connection
	req, err := portReservationRequestFromResourceData(res)
	if err != nil {
		return err
	}
	reservation, err := api.PortReservationsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(reservation.ID)
	return portReservationRead(ctx, res, api)
}

// Read a port reservation
func portReservationRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	reservation, err := api.PortReservationsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Port reservation is gone
		return nil
	}
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(reservation, res); err != nil {
		return err
	}
	return nil
}

// Update a port reservation: Only the external_ref field
// is updateable. Updates to other fields will result in an error.
func portReservationUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if res.HasChange("port_num") {
		return fmt.Errorf(
			"The `port_num` property can not be updated, you need to create a new resource")
	}
	if res.HasChange("connection") {
		return fmt.Errorf(
			"The `connection` can not be changed, you need to create a new resource")
	}

	// Make patch from resource data
	patch, err := portReservationPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.PortReservationsPatch(ctx, res.Id(), patch)
	if err != nil {
		return err
	}
	return portReservationRead(ctx, res, api)
}

// Destroy a port reservation
func portReservationDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.PortReservationsDestroy(ctx, res.Id(), nil)
	if err != nil {
		return err
	}
	return portReservationRead(ctx, res, api)
}
