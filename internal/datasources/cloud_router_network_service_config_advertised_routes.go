package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterNetworkServiceConfigAdvertisedRoutesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_config_advertised_routes` data source to get BGP routes advertised to peers on a network service config",
		ReadContext: cloudRouterNetworkServiceConfigAdvertisedRoutesRead,
		Schema: map[string]*schema.Schema{
			"network_service_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network service config ID",
			},
			"routes": schemas.IntoDataSourceResultsSchema(
				schemas.BGPRouteSchema(),
			),
		},
	}
}

func cloudRouterNetworkServiceConfigAdvertisedRoutesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	nscID := res.Get("network_service_config_id").(string)

	all, err := api.NetworkServiceConfigAdvertisedRoutesList(ctx, nscID)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("routes", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}
