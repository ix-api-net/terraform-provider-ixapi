package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterNetworkServiceConfigPrefixesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_network_service_config_prefixes` data source to get BGP prefixes learned on a network service config",
		ReadContext: cloudRouterNetworkServiceConfigPrefixesRead,
		Schema: map[string]*schema.Schema{
			"network_service_config_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network service config ID",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "gnmi",
				Description: "Query mode (gnmi or netconf)",
			},
			"prefixes": schemas.IntoDataSourceResultsSchema(
				schemas.BGPPrefixSchema(),
			),
		},
	}
}

func cloudRouterNetworkServiceConfigPrefixesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	nscID := res.Get("network_service_config_id").(string)
	mode := res.Get("mode").(string)

	all, err := api.NetworkServiceConfigPrefixesList(ctx, nscID, mode)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("prefixes", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}
