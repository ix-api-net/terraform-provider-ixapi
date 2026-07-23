package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServiceConfigsP2PVCDataSource creates a data source for
// querying network service configs of type p2p virtual circuit.
func NewNetworkServiceConfigsP2PVCDataSource() *schema.Resource {
	s := networkServiceConfigQuerySchema(schemas.P2PNetworkServiceConfigSchema())
	return &schema.Resource{
		Description: "Get network service configs of type: p2p virtual circuit",
		ReadContext: networkServiceConfigsP2PVCRead,
		Schema:      s,
	}
}

// Retrieve p2p virtual circuit network service configs and filter by
// managing/consuming account, network service or connection.
func networkServiceConfigsP2PVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkServiceConfigQuery(ixapi.P2PNetworkServiceConfigType, res)
	configs, err := api.NetworkServiceConfigsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.P2PNetworkServiceConfig, 0, len(configs))
	for _, nsc := range configs {
		resnsc, ok := nsc.(*ixapi.P2PNetworkServiceConfig)
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
