package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterNetworkServiceConfigCloudVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_network_service_config_cloud_vc` resource to configure a Cloud ROUTER connection to a cloud virtual circuit network service.",
		CreateContext: crud.Create(cloudRouterConfigCloudVCCreate),
		UpdateContext: crud.Update(cloudRouterConfigCloudVCUpdate),
		ReadContext:   crud.Read(cloudRouterConfigCloudVCRead),
		DeleteContext: crud.Delete(cloudRouterConfigCloudVCDelete),

		Schema: schemas.CloudRouterNetworkServiceConfigCloudVCSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func cloudRouterConfigCloudVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterNetworkServiceConfigRequest, error) {
	vlanConfig, err := vlanConfigFromResourceData(r)
	if err != nil {
		return nil, err
	}

	res := schemas.ResourceDataFrom(r)
	req := &ixapi.CloudRouterNetworkServiceConfigRequest{
		Type:                  "cloud_vc",
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
		CloudVLAN:             res.GetIntOpt("cloud_vlan"),
		Handover:              res.GetIntOpt("handover"),
		Connection:            nil,
		PurchaseOrder:         res.GetStringOpt("purchase_order"),
		NetworkFeatureConfigs: r.Get("network_feature_configs"),
	}
	return req, nil
}

func cloudRouterConfigCloudVCPatchFromResourceData(
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

func cloudRouterConfigCloudVCCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := cloudRouterConfigCloudVCRequestFromResourceData(res)
	if err != nil {
		return err
	}

	config, err := api.CloudRouterNetworkServiceConfigsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(config.ID)

	return cloudRouterConfigCloudVCRead(ctx, res, api)
}

func cloudRouterConfigCloudVCRead(
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

func cloudRouterConfigCloudVCUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	patch, err := cloudRouterConfigCloudVCPatchFromResourceData(res)
	if err != nil {
		return err
	}

	if patch != nil {
		_, err = api.CloudRouterNetworkServiceConfigsPatch(ctx, id, patch)
		if err != nil {
			return err
		}
	}
	return cloudRouterConfigCloudVCRead(ctx, res, api)
}

func cloudRouterConfigCloudVCDelete(
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
	return cloudRouterConfigCloudVCRead(ctx, res, api)
}
