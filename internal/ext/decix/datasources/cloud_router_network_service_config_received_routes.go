package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewCloudRouterNetworkServiceConfigReceivedRoutesDataSource returns the schema.Resource for listing BGP routes received by a Cloud Router network service config.
func NewCloudRouterNetworkServiceConfigReceivedRoutesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_config_received_routes` data source to get BGP routes received from peers on a network service config",
		ReadContext: cloudRouterNetworkServiceConfigReceivedRoutesRead,
		Schema: map[string]*schema.Schema{
			"network_service_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network service config ID",
			},
			"routes": schemas.IntoDataSourceResultsSchema(
				decixschemas.BGPRouteSchema(),
			),
		},
	}
}

func cloudRouterNetworkServiceConfigReceivedRoutesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	nscID := res.Get("network_service_config_id").(string)

	all, err := api.NetworkServiceConfigReceivedRoutesList(ctx, nscID)
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
