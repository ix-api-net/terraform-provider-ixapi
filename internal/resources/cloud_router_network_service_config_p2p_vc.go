package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterNetworkServiceConfigP2PVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_network_service_config_p2p_vc` resource to configure a Cloud ROUTER connection to a point-to-point virtual circuit network service.",
		CreateContext: crud.Create(cloudRouterConfigP2PVCCreate),
		UpdateContext: crud.Update(cloudRouterConfigP2PVCUpdate),
		ReadContext:   crud.Read(cloudRouterConfigP2PVCRead),
		DeleteContext: crud.Delete(cloudRouterConfigP2PVCDelete),

		Schema: schemas.CloudRouterNetworkServiceConfigP2PVCSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func cloudRouterConfigP2PVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceDataFrom(r)
	req := &ixapi.CloudRouterNetworkServiceConfigRequest{
		Type:                  "p2p_vc",
		ManagingAccount:       res.GetString("managing_account"),
		BillingAccount:        res.GetString("billing_account"),
		ConsumingAccount:      res.GetString("consuming_account"),
		ExternalRef:           res.GetStringOpt("external_ref"),
		CloudRouter:           res.GetString("cloud_router"),
		NetworkService:        res.GetString("network_service"),
		Address:               res.GetString("address"),
		BGPNeighbor:           res.GetString("bgp_neighbor"),
		BGPNeighborASN:        res.GetInt("bgp_neighbor_asn"),
		BGPPassword:           res.GetStringOpt("bgp_password"),
		VLANConfig:            vlanConfig,
		PolicyIngress:         res.GetStringOpt("policy_ingress"),
		PolicyEgress:          res.GetStringOpt("policy_egress"),
		AdminStatus:           res.GetString("admin_status"),
		BFDEnabled:            res.GetBool("bfd_enabled"),
		CloudVLAN:             nil,
		Handover:              nil,
		Connection:            res.GetStringOpt("nic"),
		PurchaseOrder:         res.GetStringOpt("purchase_order"),
		NetworkFeatureConfigs: r.Get("network_feature_configs"),
	}
	return req, nil
}

func cloudRouterConfigP2PVCPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterNetworkServiceConfigPatch, error) {
	res := schemas.ResourceDataFrom(r)
	patch := &ixapi.CloudRouterNetworkServiceConfigPatch{}
	hasChanges := false

	if res.HasChange("policy_ingress") {
		patch.PolicyIngress = res.GetStringOpt("policy_ingress")
		hasChanges = true
	}
	if res.HasChange("policy_egress") {
		patch.PolicyEgress = res.GetStringOpt("policy_egress")
		hasChanges = true
	}
	if res.HasChange("admin_status") {
		patch.AdminStatus = res.GetStringOpt("admin_status")
		hasChanges = true
	}

	if !hasChanges {
		return nil, nil
	}
	return patch, nil
}

func cloudRouterConfigP2PVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := cloudRouterConfigP2PVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	config, err := api.CloudRouterNetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(config.ID)

	return cloudRouterConfigP2PVCRead(ctx, res, api)
}

func cloudRouterConfigP2PVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	config, err := api.CloudRouterNetworkServiceConfigsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	if err := schemas.SetResourceData(config, res); err != nil {
		return err
	}
	return nil
}

func cloudRouterConfigP2PVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	patch, err := cloudRouterConfigP2PVCPatchFromResourceData(res)
	if err != nil {
		return err
	}

	if patch != nil {
		_, err = api.CloudRouterNetworkServiceConfigsPatch(ctx, id, patch)
		if err != nil {
			return err
		}
	}
	return cloudRouterConfigP2PVCRead(ctx, res, api)
}

func cloudRouterConfigP2PVCDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	err := api.CloudRouterNetworkServiceConfigsDestroy(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return cloudRouterConfigP2PVCRead(ctx, res, api)
}
