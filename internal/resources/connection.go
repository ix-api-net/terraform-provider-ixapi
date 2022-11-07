package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewConnectionResource creates the connection resource
func NewConnectionResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_connection` resource to create a new connection and allocate ports. Adding and managing ports can be done using the `ixapi_port_reservation` resource.",
		Schema:      schemas.ConnectionSchema(),

		CreateContext: crud.Create(connectionCreate),
		ReadContext:   crud.Read(connectionRead),
		UpdateContext: crud.Update(connectionUpdate),
		DeleteContext: crud.Delete(connectionDelete),
	}
}

// Connection Request
func connectionRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.ConnectionRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.ConnectionRequest{
		ManagingAccount:       res.GetString("managing_account"),
		ConsumingAccount:      res.GetString("consuming_account"),
		BillingAccount:        res.GetString("billing_account"),
		ExternalRef:           res.GetStringOpt("external_ref"),
		PurchaseOrder:         res.GetStringOpt("purchase_order"),
		ContractRef:           res.GetStringOpt("contract_ref"),
		RoleAssignments:       res.GetStringList("role_assignments"),
		Mode:                  res.GetString("mode"),
		LacpTimeout:           res.GetStringOpt("lacp_timeout"),
		ProductOffering:       res.GetString("product_offering"),
		PortQuantity:          res.GetInt("port_quantity"),
		SubscriberSideDemarcs: res.GetStringList("subscriber_side_demarcs"),
		ConnectingParty:       res.GetStringOpt("connecting_party"),
	}
	return req, nil
}

// Connection Patch
func connectionPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.ConnectionPatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.ConnectionPatch{}

	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		req.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("billing_account") {
		req.BillingAccount = res.GetStringOpt("billing_account")
	}
	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("contract_ref") {
		req.ContractRef = res.GetStringOpt("contract_ref")
	}
	if res.HasChange("purchase_order") {
		req.PurchaseOrder = res.GetStringOpt("purchase_order")
	}
	if res.HasChange("role_assignments") {
		req.RoleAssignments = res.GetStringList("role_assignments")
	}
	if res.HasChange("mode") {
		req.Mode = res.GetStringOpt("mode")
	}
	if res.HasChange("lacp_timeout") {
		req.LacpTimeout = res.GetStringOpt("lacp_timeout")
	}
	if res.HasChange("product_offering") {
		req.ProductOffering = res.GetStringOpt("product_offering")
	}

	return req, nil
}

// Create
func connectionCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := connectionRequestFromResourceData(res)
	if err != nil {
		return err
	}
	// Create Connection
	conn, err := api.ConnectionsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(conn.ID)
	return connectionRead(ctx, res, api)
}

// Read
func connectionRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	// Get the connection and set resource data
	conn, err := api.ConnectionsRead(ctx, res.Id())
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // This connection is gone
		return nil
	}
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(conn, res); err != nil {
		return err
	}
	return nil
}

// Update
func connectionUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	patch, err := connectionPatchFromResourceData(res)
	if err != nil {
		return err
	}
	if _, err := api.ConnectionsPatch(ctx, res.Id(), patch); err != nil {
		return err
	}
	return connectionRead(ctx, res, api)
}

// Delete
func connectionDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return connectionRead(ctx, res, api)
}
