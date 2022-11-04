package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceConfigCloudVCResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkServiceConfigCloudVCResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_network_service_config_cloud_vc` resource to provision an access to a point to point virutal circuit.",

		CreateContext: crud.Create(nscCloudVCCreate),
		UpdateContext: crud.Update(nscCloudVCUpdate),
		ReadContext:   crud.Read(nscCloudVCRead),
		DeleteContext: crud.Delete(nscCloudVCDelete),

		Schema: schemas.CloudNetworkServiceConfigSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func nscCloudVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.CloudNetworkServiceConfigRequest{
		Type:             ixapi.CloudNetworkServiceConfigType,
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		BillingAccount:   res.GetString("billing_account"),
		ExternalRef:      res.GetStringOpt("external_ref"),
		NetworkService:   res.GetString("network_service"),
		PurchaseOrder:    res.GetStringOpt("purchase_order"),
		ContractRef:      res.GetStringOpt("contract_ref"),
		RoleAssignments:  res.GetStringList("role_assignments"),
		Connection:       res.GetString("network_connection"),
		VLANConfig:       vlanConfig,
		Handover:         res.GetInt("handover"),
		CloudVLAN:        res.GetInt("cloud_vlan"),
	}
	return req, nil
}

func nscCloudVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudNetworkServiceConfigPatch, error) {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.CloudNetworkServiceConfigPatch{
		Type: ixapi.CloudNetworkServiceConfigType,
	}
	if res.HasChange("vlan_config") {
		vlanConfig, err := vlanConfigFromResourceData(r)
		if err != nil {
			return nil, err
		}
		patch.VLANConfig = vlanConfig
	}
	if res.HasChange("managing_account") {
		patch.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		patch.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("billing_account") {
		patch.BillingAccount = res.GetStringOpt("billing_account")
	}
	if res.HasChange("external_ref") {
		patch.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("purchase_order") {
		patch.PurchaseOrder = res.GetStringOpt("purchase_order")
	}
	if res.HasChange("contract_ref") {
		patch.ContractRef = res.GetStringOpt("contract_ref")
	}
	if res.HasChange("role_assignments") {
		patch.RoleAssignments = res.GetStringList("role_assignments")
	}
	if res.HasChange("network_connection") {
		patch.Connection = res.GetStringOpt("network_connection")
	}
	if res.HasChange("cloud_vlan") {
		patch.CloudVLAN = res.GetIntOpt("cloud_vlan")
	}
	if res.HasChange("handover") {
		patch.Handover = res.GetIntOpt("handover")
	}
	return patch, nil
}

func nscCloudVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nscCloudVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request
	nsc, err := api.NetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	cloudnsc, ok := nsc.(*ixapi.CloudNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.CloudNetworkServiceConfigType)
	}
	res.SetId(cloudnsc.ID)
	return nscCloudVCRead(ctx, res, api)
}

func nscCloudVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()

	nsc, err := api.NetworkServiceConfigsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Resource not longer available
		return nil
	}
	if err != nil {
		return err
	}
	cloudnsc, ok := nsc.(*ixapi.CloudNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.CloudNetworkServiceConfigType)
	}
	if err := schemas.SetResourceData(cloudnsc, res); err != nil {
		return err
	}
	return nil
}

func nscCloudVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nscCloudVCPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.NetworkServiceConfigsPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nscCloudVCRead(ctx, res, api)
}

func nscCloudVCDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.NetworkServiceConfigsDestroy(ctx, res.Id(), nil)
	if err != nil {
		return err
	}
	return nscCloudVCRead(ctx, res, api)
}
