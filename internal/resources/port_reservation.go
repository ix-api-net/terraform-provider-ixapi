package resources

import (
	"context"

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
			"port_index": schemas.DataSourceQueryInt(
				"When the port_index is provided, the resource will reference an already allocated port reservation."),
		})
	return &schema.Resource{
		Description: "Use the `ixapi_port_reservation` resource to manage port reservations for a connection. You can create a new port reservation by only passing the `connection` to the resource, or you can reference an already allocated port reservation, which was created when the connection was allocated, by specifying a `port_index`.",
		Schema:      s,

		CreateContext: crud.Create(portReservationCreate),
		ReadContext:   crud.Read(portReservationRead),
		UpdateContext: crud.Update(portReservationUpdate),
		DeleteContext: crud.Delete(portReservationDelete),
	}
}

// Create a port reservation resource
func portReservationCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

// Read a port reservation
func portReservationRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil

}

// Update a port reservation: Only the external_ref field
// is updateable. Updates to other fields will result in an error.
func portReservationUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

// Destroy a port reservation
func portReservationDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}
