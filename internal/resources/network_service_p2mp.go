package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceP2MPVCResource creates a point to point
// virtual circuit network service
func NewNetworkServiceP2MPVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_network_service_p2mp_vc` resource to create and manage a point to multi point virtual circuit. Participants need to create an `ixapi_network_service_config_p2p_vc`. A public p2p vc can be joined by everyone on the exchange, unless denied by a member joining rule. You can create and manage joining rules using the `ixapi_member_joining_rule` resource.",
		CreateContext: crud.Create(nsP2MPVCCreate),
		UpdateContext: crud.Update(nsP2MPVCUpdate),
		ReadContext:   crud.Read(nsP2MPVCRead),
		DeleteContext: crud.Delete(nsP2MPVCDelete),

		Schema: schemas.P2MPNetworkServiceSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// Make create
func nsP2MPVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2MPNetworkServiceRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.P2MPNetworkServiceRequest{
		Type: ixapi.P2MPNetworkServiceType,

		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		BillingAccount:   res.GetString("billing_account"),

		ProductOffering: res.GetString("product_offering"),

		ExternalRef:   res.GetStringOpt("external_ref"),
		PurchaseOrder: res.GetStringOpt("purchase_order"),
		ContractRef:   res.GetStringOpt("contract_ref"),

		DisplayName: res.GetStringOpt("display_name"),

		Public: res.GetBoolOpt("public"),
	}
	return req, nil
}

// Make update
func nsP2MPVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2MPNetworkServicePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.P2MPNetworkServicePatch{
		Type: ixapi.P2MPNetworkServiceType,
	}

	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		req.ManagingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("billing_account") {
		req.ManagingAccount = res.GetStringOpt("billing_account")
	}

	if res.HasChange("product_offering") {
		req.ProductOffering = res.GetStringOpt("product_offering")
	}

	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("purchase_order") {
		req.PurchaseOrder = res.GetStringOpt("purchase_order")
	}
	if res.HasChange("contract_ref") {
		req.ContractRef = res.GetStringOpt("contract_ref")
	}
	if res.HasChange("display_name") {
		req.DisplayName = res.GetStringOpt("display_name")
	}

	if res.HasChange("public") {
		req.Public = res.GetBoolOpt("public")
	}
	return req, nil
}

// Create NetworkService
func nsP2MPVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nsP2MPVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request and create network service
	ns, err := api.NetworkServicesCreate(ctx, req)
	if err != nil {
		return err
	}
	p2mpns, ok := ns.(*ixapi.P2MPNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.P2MPNetworkServiceType)
	}
	res.SetId(p2mpns.ID)

	return nsP2MPVCRead(ctx, res, api)
}

// Read NetworkService
func nsP2MPVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	ns, err := api.NetworkServicesRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Network service is gone
		return nil
	}
	if err != nil {
		return err
	}
	p2mpns, ok := ns.(*ixapi.P2MPNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.P2MPNetworkServiceType)
	}
	if err := schemas.SetResourceData(p2mpns, res); err != nil {
		return err
	}
	return nil
}

// Update NetworkService
func nsP2MPVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nsP2MPVCPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.NetworkServicesPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nsP2MPVCRead(ctx, res, api)
}

// Delete NetworkService
func nsP2MPVCDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	req := &ixapi.CancellationRequest{}
	_, err := api.NetworkServicesDestroy(ctx, id, req)
	if err != nil {
		return err
	}
	return nsP2MPVCRead(ctx, res, api)
}
