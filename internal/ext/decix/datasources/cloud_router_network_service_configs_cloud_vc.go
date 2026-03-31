package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

func NewCloudRouterNetworkServiceConfigsCloudVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_configs_cloud_vc` data source to find Cloud ROUTER NSCs of type cloud_vc",
		ReadContext: cloudRouterNetworkServiceConfigsCloudVCRead,
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
				decixschemas.CloudRouterNetworkServiceConfigCloudVCSchema(),
			),
		},
	}
}

func cloudRouterNetworkServiceConfigsCloudVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	qry := &ixapi.CloudRouterNetworkServiceConfigsListQuery{
		Type: "cloud_vc",
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

func NewCloudRouterNetworkServiceConfigCloudVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_config_cloud_vc` data source to get a single Cloud ROUTER NSC of type cloud_vc by ID",
		ReadContext: cloudRouterNetworkServiceConfigCloudVCDataRead,
		Schema:      schemas.IntoDataSourceSchema(decixschemas.CloudRouterNetworkServiceConfigCloudVCSchema()),
	}
}

func cloudRouterNetworkServiceConfigCloudVCDataRead(
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
		return diag.Errorf("the cloud_router_network_service_config_cloud_vc `id` is required")
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
