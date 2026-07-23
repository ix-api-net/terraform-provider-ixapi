package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceConfigsExchangeLanDataSource creates a data source for
// querying network service configs of type exchange lan.
func NewNetworkServiceConfigsExchangeLanDataSource() *schema.Resource {
	s := networkServiceConfigQuerySchema(schemas.ExchangeLanNetworkServiceConfigSchema())
	return &schema.Resource{
		Description: "Get network service configs of type: exchange lan",
		ReadContext: networkServiceConfigsExchangeLanRead,
		Schema:      s,
	}
}

// Retrieve exchange lan network service configs and filter by
// managing/consuming account, network service or connection.
func networkServiceConfigsExchangeLanRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkServiceConfigQuery(ixapi.ExchangeLanNetworkServiceConfigType, res)
	configs, err := api.NetworkServiceConfigsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.ExchangeLanNetworkServiceConfig, 0, len(configs))
	for _, nsc := range configs {
		resnsc, ok := nsc.(*ixapi.ExchangeLanNetworkServiceConfig)
		if !ok {
			continue // should not happen with well behaving servers
		}
		filtered = append(filtered, resnsc)
	}

	flat, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("network_service_configs", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())
	return nil
}
