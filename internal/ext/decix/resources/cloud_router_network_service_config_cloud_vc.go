package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
	coreresources "github.com/ix-api-net/terraform-provider-ixapi/internal/resources"
)

// NewDecixCloudRouterNetworkServiceConfigCloudVCResource returns the schema.Resource for managing a Cloud Router connection to a cloud virtual circuit network service.
func NewDecixCloudRouterNetworkServiceConfigCloudVCResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_network_service_config_cloud_vc` resource to configure a Cloud ROUTER connection to a cloud virtual circuit network service.",
		CreateContext: crud.Create(cloudRouterConfigCloudVCCreate),
		UpdateContext: crud.Update(cloudRouterConfigCloudVCUpdate),
		ReadContext:   crud.Read(cloudRouterConfigCloudVCRead),
		DeleteContext: crud.Delete(cloudRouterConfigCloudVCDelete),

		Schema: decixschemas.CloudRouterNetworkServiceConfigCloudVCSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func cloudRouterConfigCloudVCRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterNetworkServiceConfigRequest, error) {
	vlanConfig, err := coreresources.VlanConfigFromResourceData(r)
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
		ASOverride:            res.GetBool("as_override"),
		CloudVLAN:             res.GetIntOpt("cloud_vlan"),
		Handover:              res.GetIntOpt("handover"),
		Connection:            nil,
		PurchaseOrder:         res.GetStringOpt("purchase_order"),
		NetworkFeatureConfigs: res.GetStringList("network_feature_configs"),
	}
	return req, nil
}

func cloudRouterConfigCloudVCPatchFromResourceData(
	r *schema.ResourceData,
) *ixapi.CloudRouterNetworkServiceConfigPatch {
	res := schemas.ResourceDataFrom(r)
	return &ixapi.CloudRouterNetworkServiceConfigPatch{
		PolicyIngress: res.GetStringOpt("policy_ingress"),
		PolicyEgress:  res.GetStringOpt("policy_egress"),
		AdminStatus:   res.GetStringOpt("admin_status"),
	}
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

	config, err := api.DecixCloudRouterNetworkServiceConfigsCreate(ctx, req)
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
	config, err := api.DecixCloudRouterNetworkServiceConfigsRead(ctx, id)
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
	patch := cloudRouterConfigCloudVCPatchFromResourceData(res)
	if _, err := api.DecixCloudRouterNetworkServiceConfigsPatch(ctx, id, patch); err != nil {
		return err
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

	policyIngress := res.Get("policy_ingress").(string)
	policyEgress := res.Get("policy_egress").(string)
	if policyIngress != "" || policyEgress != "" {
		empty := ""
		patch := &ixapi.CloudRouterNetworkServiceConfigPatch{}
		if policyIngress != "" {
			patch.PolicyIngress = &empty
		}
		if policyEgress != "" {
			patch.PolicyEgress = &empty
		}
		if _, err := api.DecixCloudRouterNetworkServiceConfigsPatch(ctx, id, patch); err != nil {
			return err
		}
	}

	err := api.DecixCloudRouterNetworkServiceConfigsDestroy(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return cloudRouterConfigCloudVCRead(ctx, res, api)
}
