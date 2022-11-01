package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkFeatureConfigRouteServerResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkFeatureConfigRouteServerResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_nfc_route_server` resource to configure access to a route server in the exchange lan.",

		CreateContext: crud.Create(nfcRouteServerCreate),
		UpdateContext: crud.Update(nfcRouteServerUpdate),
		ReadContext:   crud.Read(nfcRouteServerRead),
		DeleteContext: crud.Delete(nfcRouteServerDelete),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: schemas.RouteServerNetworkFeatureConfigSchema(),
	}
}

func nfcRouteServerRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.RouteServerNetworkFeatureConfigRequest, error) {
	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.RouteServerNetworkFeatureConfigRequest{
		Type:                 ixapi.RouteServerNetworkFeatureConfigType,
		ManagingAccount:      res.GetString("managing_account"),
		ConsumingAccount:     res.GetString("consuming_account"),
		BillingAccount:       res.GetString("billing_account"),
		ExternalRef:          res.GetStringOpt("external_ref"),
		PurchaseOrder:        res.GetStringOpt("purchase_order"),
		ContractRef:          res.GetStringOpt("contract_ref"),
		RoleAssignments:      res.GetStringList("role_assignments"),
		NetworkFeature:       res.GetString("network_feature"),
		NetworkServiceConfig: res.GetString("network_service_config"),
		// TODO: FeatureFlags
		ASN:            res.GetInt("asn"),
		Password:       res.GetStringOpt("password"),
		AsSetV4:        res.GetStringOpt("as_set_v4"),
		AsSetV6:        res.GetStringOpt("as_set_v6"),
		MaxPrefixV4:    res.GetIntOpt("max_prefix_v4"),
		MaxPrefixV6:    res.GetIntOpt("max_prefix_v6"),
		SessionMode:    res.GetString("session_mode"),
		BGPSessionType: res.GetString("bgp_session_type"),
		IP:             res.GetString("ip"),
	}
	return req, nil
}

func nfcRouteServerPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.RouteServerNetworkFeatureConfigPatch, error) {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.RouteServerNetworkFeatureConfigPatch{
		Type: ixapi.RouteServerNetworkFeatureConfigType,
	}

	if res.HasChange("managing_account") {
		patch.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		patch.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("external_ref") {
		patch.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("asn") {
		patch.ASN = res.GetIntOpt("asn")
	}
	if res.HasChange("password") {
		patch.Password = res.GetStringOpt("password")
	}
	if res.HasChange("as_set_v4") {
		patch.AsSetV4 = res.GetStringOpt("as_set_v4")
	}
	if res.HasChange("as_set_v6") {
		patch.AsSetV6 = res.GetStringOpt("as_set_v6")
	}
	if res.HasChange("session_mode") {
		patch.SessionMode = res.GetStringOpt("session_mode")
	}
	if res.HasChange("bgp_session_type") {
		patch.BGPSessionType = res.GetStringOpt("bgp_session_type")
	}
	if res.HasChange("ip") {
		patch.IP = res.GetStringOpt("ip")
	}
	return patch, nil
}

func nfcRouteServerCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := nfcRouteServerRequestFromResourceData(res)
	if err != nil {
		return err
	}

	nfc, err := api.NetworkFeatureConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	rsnfc, ok := nfc.(*ixapi.RouteServerNetworkFeatureConfig)
	if !ok {
		return fmt.Errorf(
			"API did return a %s instead of a route server network feature config",
			nfc.PolymorphicType())
	}
	res.SetId(rsnfc.ID)
	return nfcRouteServerRead(ctx, res, api)
}

func nfcRouteServerRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()

	nfc, err := api.NetworkFeatureConfigsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("") // Resource not longer available
		return nil
	}
	if err != nil {
		return err
	}
	rsnfc, ok := nfc.(*ixapi.RouteServerNetworkFeatureConfig)
	if !ok {
		return fmt.Errorf(
			"API did return a %s instead of a route server network feature config",
			nfc.PolymorphicType())
	}
	if err := schemas.SetResourceData(rsnfc, res); err != nil {
		return err
	}
	return nil
}

func nfcRouteServerUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := nfcRouteServerPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.NetworkFeatureConfigsPatch(ctx, id, patch)
	if err != nil {
		return err
	}

	return nfcRouteServerRead(ctx, res, api)
}

func nfcRouteServerDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	_, err := api.NetworkFeatureConfigsDestroy(ctx, res.Id())
	if err != nil {
		return nil
	}
	return nfcRouteServerRead(ctx, res, api)
}
