package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceConfigP2PVCResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkServiceConfigP2PVCResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_network_service_config_p2p_vc` resource to provision an access to a point to point virutal circuit.",

		CreateContext: crud.Create(nscP2PVCCreate),
		UpdateContext: crud.Update(nscP2PVCUpdate),
		ReadContext:   crud.Read(nscP2PVCRead),
		DeleteContext: crud.Delete(nscP2PVCDelete),

		Schema: schemas.P2PNetworkServiceConfigSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func nscP2PVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2PNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.P2PNetworkServiceConfigRequest{
		Type:             ixapi.P2PNetworkServiceConfigType,
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
	}
	return req, nil
}

func nscP2PVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.P2PNetworkServiceConfigPatch, error) {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.P2PNetworkServiceConfigPatch{
		Type: ixapi.P2PNetworkServiceConfigType,
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
	return patch, nil
}

func nscP2PVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nscP2PVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request
	nsc, err := api.NetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	p2pnsc, ok := nsc.(*ixapi.P2PNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.P2PNetworkServiceConfigType)
	}
	res.SetId(p2pnsc.ID)
	return nscP2PVCRead(ctx, res, api)
}

func nscP2PVCRead(
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
	p2pnsc, ok := nsc.(*ixapi.P2PNetworkServiceConfig)
	if !ok {
		return ErrUnexpectedPolymorphic(nsc, ixapi.P2PNetworkServiceConfigType)
	}
	if err := schemas.SetResourceData(p2pnsc, res); err != nil {
		return err
	}
	return nil
}

func nscP2PVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nscP2PVCPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.NetworkServiceConfigsPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return nscP2PVCRead(ctx, res, api)
}

func nscP2PVCDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.NetworkServiceConfigsDestroy(ctx, res.Id(), nil)
	if err != nil {
		return err
	}
	return nscP2PVCRead(ctx, res, api)
}
