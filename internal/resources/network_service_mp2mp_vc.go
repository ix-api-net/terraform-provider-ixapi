package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceMP2MPVCResource creates a point to point
// virtual circuit network service
func NewNetworkServiceMP2MPVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_network_service_mp2mp_vc` resource to create and manage a multi point to multi point virtual circuit. Participants need to create an `ixapi_network_service_config_mp2mp_vc`. A public mp2mp vc can be joined by everyone on the exchange, unless denied by a member joining rule. You can create and manage joining rules using the `ixapi_member_joining_rule` resource.",
		CreateContext: crud.Create(nsMP2MPVCCreate),
		UpdateContext: crud.Update(nsMP2MPVCUpdate),
		ReadContext:   crud.Read(nsMP2MPVCRead),
		DeleteContext: crud.Delete(nsMP2MPVCDelete),

		Schema: schemas.MP2MPNetworkServiceSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// Make create
func nsMP2MPVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.MP2MPNetworkServiceRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.MP2MPNetworkServiceRequest{
		Type: ixapi.MP2MPNetworkServiceType,

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
func nsMP2MPVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.MP2MPNetworkServicePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.MP2MPNetworkServicePatch{
		Type: ixapi.MP2MPNetworkServiceType,
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
func nsMP2MPVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nsMP2MPVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request and create network service
	ns, err := api.NetworkServicesCreate(ctx, req)
	if err != nil {
		return err
	}
	mp2mpns, ok := ns.(*ixapi.MP2MPNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.MP2MPNetworkServiceType)
	}
	res.SetId(mp2mpns.ID)

	return nsMP2MPVCRead(ctx, res, api)
}

// Read NetworkService
func nsMP2MPVCRead(
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
	mp2mpns, ok := ns.(*ixapi.MP2MPNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.MP2MPNetworkServiceType)
	}
	if err := schemas.SetResourceData(mp2mpns, res); err != nil {
		return err
	}
	return nil
}

// Update NetworkService
func nsMP2MPVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nsMP2MPVCPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.NetworkServicesPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nsMP2MPVCRead(ctx, res, api)
}

// Delete NetworkService
func nsMP2MPVCDelete(
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
	return nsMP2MPVCRead(ctx, res, api)
}
