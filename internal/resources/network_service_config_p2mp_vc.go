package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceConfigP2MPVCResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkServiceConfigP2MPVCResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_network_service_config_p2mp_vc` resource to provision an access to a point to multi point virutal circuit.",

		CreateContext: crud.Create(nscP2MPVCCreate),
		UpdateContext: crud.Update(nscP2MPVCUpdate),
		ReadContext:   crud.Read(nscP2MPVCRead),
		DeleteContext: crud.Delete(nscP2MPVCDelete),

		Schema: schemas.P2MPNetworkServiceConfigSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func nscP2MPVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2MPNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.P2MPNetworkServiceConfigRequest{
		Type:             ixapi.P2MPNetworkServiceConfigType,
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
		ProductOffering:  res.GetStringOpt("product_offering"),
		Capacity:         res.GetIntOpt("capacity"),
		Role:             res.GetStringOptDefault("role", "leaf"),
	}
	return req, nil
}

func nscP2MPVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2MPNetworkServiceConfigPatch, error) {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.P2MPNetworkServiceConfigPatch{
		Type: ixapi.P2MPNetworkServiceConfigType,
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
	if res.HasChange("capacity") {
		patch.Capacity = res.GetIntOpt("capacity")
	}
	if res.HasChange("product_offering") {
		patch.ProductOffering = res.GetStringOpt("product_offering")
	}
	if res.HasChange("role") {
		patch.Role = res.GetStringOpt("role")
	}
	return patch, nil
}

func nscP2MPVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nscP2MPVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request
	nsc, err := api.NetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	p2mpnsc, ok := nsc.(*ixapi.P2MPNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.P2MPNetworkServiceConfigType)
	}
	res.SetId(p2mpnsc.ID)
	return nscP2MPVCRead(ctx, res, api)
}

func nscP2MPVCRead(
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
	p2mpnsc, ok := nsc.(*ixapi.P2MPNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.P2MPNetworkServiceConfigType)
	}
	if err := schemas.SetResourceData(p2mpnsc, res); err != nil {
		return err
	}
	return nil
}

func nscP2MPVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nscP2MPVCPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.NetworkServiceConfigsPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nscP2MPVCRead(ctx, res, api)
}

func nscP2MPVCDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.NetworkServiceConfigsDestroy(ctx, res.Id(), nil)
	if err != nil {
		return err
	}
	return nscP2MPVCRead(ctx, res, api)
}
