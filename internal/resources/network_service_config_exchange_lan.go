package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/crud"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewNetworkServiceConfigExchangeLanResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkServiceConfigExchangeLanResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_nsc_exchange_lan` resource to provision an access to an exchange lan for a customer.",

		CreateContext: crud.Create(nscExchangeLanCreate),
		UpdateContext: crud.Update(nscExchangeLanUpdate),
		ReadContext:   crud.Read(nscExchangeLanRead),
		DeleteContext: crud.Delete(nscExchangeLanDelete),

		Schema: schemas.ExchangeLanNetworkServiceConfigSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

// Create polymorphic VLanConfig from resource data
func vlanConfigFromResourceData(r *schema.ResourceData) (ixapi.VLANConfig, error) {
	res := schemas.ResourceData{ResourceData: r}
	c := res.GetResource("vlan_config")
	vType := c["vlan_type"].(string)

	if vType == "port" {
		cfg := &ixapi.VLANConfigPort{
			VLANType: "port",
		}
		return cfg, nil
	}
	if vType == "dot1q" {
		vlan := c.GetIntOpt("vlan")
		ethertype := c.GetStringOptDefault("vlan_ethertype", "0x8100")
		cfg := &ixapi.VLANConfigDot1Q{
			VLANType:      "dot1q",
			VLAN:          vlan,
			VLANEthertype: ethertype,
		}
		return cfg, nil
	}
	if vType == "qinq" {
		outerVlanEthertype := c.GetStringOptDefault("outer_vlan_ethertype", "0x8100")
		outerVlan := c.GetIntOpt("outer_vlan")
		innerVlan := c.GetIntOpt("inner_vlan")
		if innerVlan == nil {
			return nil, fmt.Errorf("The `inner_vlan` property is required for qinq vlan configs")
		}
		cfg := &ixapi.VLANConfigQinQ{
			VLANType:           "qinq",
			OuterVLAN:          outerVlan,
			OuterVLANEthertype: outerVlanEthertype,
			InnerVLAN:          *innerVlan,
		}
		return cfg, nil
	}

	return nil, fmt.Errorf("unknown vlan config type: %s", vType)
}

func nscExchangeLanRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.ExchangeLanNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.ExchangeLanNetworkServiceConfigRequest{
		Type:             "exchange_lan",
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
		Capacity:         res.GetIntOpt("capacity"),
		ASNs:             res.GetIntList("asns"),
		Macs:             res.GetStringList("macs"),
		Listed:           res.GetBool("listed"),
		ProductOffering:  res.GetString("product_offering"),
	}
	return req, nil
}

func nscExchangeLanPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.ExchangeLanNetworkServiceConfigPatch, error) {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.ExchangeLanNetworkServiceConfigPatch{
		Type: ixapi.ExchangeLanNetworkServiceConfigType,
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
	if res.HasChange("asns") {
		patch.ASNs = res.GetIntList("asns")
	}
	if res.HasChange("macs") {
		patch.Macs = res.GetStringList("macs")
	}
	if res.HasChange("listed") {
		patch.Listed = res.GetBoolOpt("listed")
	}
	return patch, nil
}

func nscExchangeLanCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nscExchangeLanRequestFromResourceData(res)
	if err != nil {
		return err
	}

	// Make API request
	nsc, err := api.NetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	elnsc, ok := nsc.(*ixapi.ExchangeLanNetworkServiceConfig)
	if !ok {
		return fmt.Errorf(
			"API did not return an exchange lan network service config, but: %s",
			nsc.PolymorphicType())
	}
	res.SetId(elnsc.ID)
	return nscExchangeLanRead(ctx, res, api)
}

func nscExchangeLanRead(
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
	elnsc, ok := nsc.(*ixapi.ExchangeLanNetworkServiceConfig)
	if !ok {
		return fmt.Errorf(
			"API did not return an exchange lan network service config, but: %s",
			nsc.PolymorphicType())
	}
	if err := schemas.SetResourceData(elnsc, res); err != nil {
		return err
	}
	return nil
}

func nscExchangeLanUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nscExchangeLanPatchFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.NetworkServiceConfigsPatch(ctx, id, patch)
	if err != nil {
		return err
	}

	return nscExchangeLanRead(ctx, res, api)
}

func nscExchangeLanDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.NetworkServiceConfigsDestroy(ctx, res.Id(), nil)
	if err != nil {
		return err
	}
	return nscExchangeLanRead(ctx, res, api)
}
