package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceCloudVCResource creates a point to point
// virtual circuit network service
func NewNetworkServiceCloudVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_network_service_cloud_vc` resource to create and manage a point to point virtual circuit. Participants need to create an `ixapi_network_service_config_cloud_vc`.",
		CreateContext: crud.Create(nsCloudVCCreate),
		UpdateContext: crud.Update(nsCloudVCUpdate),
		ReadContext:   crud.Read(nsCloudVCRead),
		DeleteContext: crud.Delete(nsCloudVCDelete),

		Schema: schemas.CloudNetworkServiceSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// Make create
func nsCloudVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudNetworkServiceRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.CloudNetworkServiceRequest{
		Type: ixapi.CloudNetworkServiceType,

		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		BillingAccount:   res.GetString("billing_account"),

		ProductOffering: res.GetString("product_offering"),

		ExternalRef:   res.GetStringOpt("external_ref"),
		PurchaseOrder: res.GetStringOpt("purchase_order"),
		ContractRef:   res.GetStringOpt("contract_ref"),

		CloudKey: res.GetString("cloud_key"),
		Capacity: res.GetIntOpt("capacity"),
	}
	return req, nil
}

// Make update
func nsCloudVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudNetworkServicePatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.CloudNetworkServicePatch{
		Type: ixapi.CloudNetworkServiceType,
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
	if res.HasChange("cloud_key") {
		req.CloudKey = res.GetStringOpt("cloud_key")
	}
	if res.HasChange("capacity") {
		req.Capacity = res.GetIntOpt("capacity")
	}
	return req, nil
}

// Create NetworkService
func nsCloudVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nsCloudVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request and create network service
	ns, err := api.NetworkServicesCreate(ctx, req)
	if err != nil {
		return err
	}
	cloudns, ok := ns.(*ixapi.CloudNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.CloudNetworkServiceType)
	}
	res.SetId(cloudns.ID)

	return nsCloudVCRead(ctx, res, api)
}

// Read NetworkService
func nsCloudVCRead(
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
	cloudns, ok := ns.(*ixapi.CloudNetworkService)
	if !ok {
		return ErrUnexpectedPolymorphic(ns, ixapi.CloudNetworkServiceType)
	}
	if err := schemas.SetResourceData(cloudns, res); err != nil {
		return err
	}
	return nil
}

// Update NetworkService
func nsCloudVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nsCloudVCPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.NetworkServicesPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nsCloudVCRead(ctx, res, api)
}

// Delete NetworkService
func nsCloudVCDelete(
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
	return nsCloudVCRead(ctx, res, api)
}
