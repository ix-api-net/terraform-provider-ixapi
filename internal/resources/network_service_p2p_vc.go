package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceP2PVCResource creates a point to point
// virtual circuit network service
func NewNetworkServiceP2PVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_network_service_p2p_vc` resource to create and manage a point to point virtual circuit. Participants need to create an `ixapi_network_service_config_p2p_vc`.",
		CreateContext: crud.Create(nsP2PVCCreate),
		UpdateContext: crud.Update(nsP2PVCUpdate),
		ReadContext:   crud.Read(nsP2PVCRead),
		DeleteContext: crud.Delete(nsP2PVCDelete),

		Schema: schemas.P2PNetworkServiceSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// Make create
func nsP2PVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2PNetworkServiceRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.P2PNetworkServiceRequest{
		Type: ixapi.P2PNetworkServiceType,

		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		BillingAccount:   res.GetString("billing_account"),

		ProductOffering: res.GetString("product_offering"),

		ExternalRef:   res.GetStringOpt("external_ref"),
		PurchaseOrder: res.GetStringOpt("purchase_order"),
		ContractRef:   res.GetStringOpt("contract_ref"),

		DisplayName: res.GetStringOpt("display_name"),

		JoiningMemberAccount: res.GetString("joining_member_account"),
	}
	return req, nil
}

// Make update
func nsP2PVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2PNetworkServicePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.P2PNetworkServicePatch{
		Type: ixapi.P2PNetworkServiceType,
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
	if res.HasChange("joining_member_account") {
		req.JoiningMemberAccount = res.GetStringOpt("joining_member_account")
	}
	return req, nil
}

// Create NetworkService
func nsP2PVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nsP2PVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request and create network service
	ns, err := api.NetworkServicesCreate(ctx, req)
	if err != nil {
		return err
	}
	p2pns, ok := ns.(*ixapi.P2PNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.P2PNetworkServiceType)
	}
	res.SetId(p2pns.ID)

	return nsP2PVCRead(ctx, res, api)
}

// Read NetworkService
func nsP2PVCRead(
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
	p2pns, ok := ns.(*ixapi.P2PNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.P2PNetworkServiceType)
	}
	if err := schemas.SetResourceData(p2pns, res); err != nil {
		return err
	}
	return nil
}

// Update NetworkService
func nsP2PVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nsP2PVCPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.NetworkServicesPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nsP2PVCRead(ctx, res, api)
}

// Delete NetworkService
func nsP2PVCDelete(
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

	return nsP2PVCRead(ctx, res, api)
}
