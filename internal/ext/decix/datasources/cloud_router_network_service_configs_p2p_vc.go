package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewCloudRouterNetworkServiceConfigsP2PVCDataSource returns the schema.Resource for listing Cloud Router network service configs of type p2p_vc.
func NewCloudRouterNetworkServiceConfigsP2PVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_configs_p2p_vc` data source to find Cloud ROUTER NSCs of type p2p_vc",
		ReadContext: cloudRouterNetworkServiceConfigsP2PVCRead,
		Schema: map[string]*schema.Schema{
			"bgp_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Filter by BGP password",
			},
			"bfd_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter by BFD enabled status",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limit the number of results",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Offset for pagination",
			},
			"cloud_router_network_service_configs": schemas.IntoDataSourceResultsSchema(
				decixschemas.CloudRouterNetworkServiceConfigP2PVCSchema(),
			),
		},
	}
}

func cloudRouterNetworkServiceConfigsP2PVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	qry := &ixapi.CloudRouterNetworkServiceConfigsListQuery{
		Type: "p2p_vc",
	}

	if bgpPassword, ok := res.GetOk("bgp_password"); ok {
		qry.BGPPassword = bgpPassword.(string)
	}

	if bfdEnabled, ok := res.GetOk("bfd_enabled"); ok {
		val := bfdEnabled.(bool)
		qry.BFD = &val
	}

	if limit, ok := res.GetOk("limit"); ok {
		qry.Limit = limit.(int)
	}

	if offset, ok := res.GetOk("offset"); ok {
		qry.Offset = offset.(int)
	}

	all, err := api.CloudRouterNetworkServiceConfigsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("cloud_router_network_service_configs", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewCloudRouterNetworkServiceConfigP2PVCDataSource returns the schema.Resource for reading a single Cloud Router network service config of type p2p_vc by ID.
func NewCloudRouterNetworkServiceConfigP2PVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_config_p2p_vc` data source to get a single Cloud ROUTER NSC of type p2p_vc by ID",
		ReadContext: cloudRouterNetworkServiceConfigP2PVCDataRead,
		Schema:      schemas.IntoDataSourceSchema(decixschemas.CloudRouterNetworkServiceConfigP2PVCSchema()),
	}
}

func cloudRouterNetworkServiceConfigP2PVCDataRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	id, hasID := res.GetOk("id")

	if !hasID {
		return diag.Errorf("the cloud_router_network_service_config_p2p_vc `id` is required")
	}

	config, err := api.CloudRouterNetworkServiceConfigsRead(ctx, id.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := schemas.SetResourceData(config, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(config.ID)

	return nil
}
