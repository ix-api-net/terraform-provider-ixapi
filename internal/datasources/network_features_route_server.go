package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewNetworkFeaturesRouteServerDataSource creates a data source for
// retrieving route server network features, related to an exchange
// lan network service config.
func NewNetworkFeaturesRouteServerDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the ixapi_network_feature_route_server data source to retrieve a list of route server network features, required for an exchange lan network service configuration",
		ReadContext: networkFeaturesRouteServerRead,
		Schema: map[string]*schema.Schema{
			"network_service": &schema.Schema{
				Description: "Filter by related (exchange lan) network service",
				Required:    true,
				Type:        schema.TypeString,
			},
			"network_features": schemas.IntoDataSourceResultsSchema(
				schemas.RouteServerNetworkFeatureSchema()),
		},
	}
}

func networkFeatureRouteServerQuery(
	res *schema.ResourceData,
) *ixapi.NetworkFeaturesListQuery {
	// Filters
	networkService := res.Get("network_service").(string)
	return &ixapi.NetworkFeaturesListQuery{
		Type:           "route_server",
		NetworkService: networkService,
	}
}

// Fetch route server network features from server
func networkFeaturesRouteServerRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkFeatureRouteServerQuery(res)
	results, err := api.NetworkFeaturesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}
	features := make([]*ixapi.RouteServerNetworkFeature, 0, len(results))
	for _, f := range results {
		rsnf, ok := f.(*ixapi.RouteServerNetworkFeature)
		if !ok {
			continue // Should not happen on well behaved servers
		}
		features = append(features, rsnf)
	}

	flat, err := schemas.FlattenModels(features)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("network_features", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewNetworkFeatureRouteServerDataSource creates a data source to
// fetch a single route server network feature identified by ID
func NewNetworkFeatureRouteServerDataSource() *schema.Resource {
	rsnfSchema := schemas.IntoDataSourceSchema(
		schemas.RouteServerNetworkFeatureSchema())
	rsnfSchema["id"].Computed = false
	rsnfSchema["id"].Optional = false
	rsnfSchema["id"].Required = true
	return &schema.Resource{
		Description: "Use the ixapi_network_feature_route_server data source to retrieve a single route server network feature identified by ID",
		ReadContext: networkFeatureRouteServerRead,
		Schema:      rsnfSchema,
	}
}

// Fetch a route server network feature from the server
func networkFeatureRouteServerRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	id := res.Get("id").(string)

	result, err := api.NetworkFeaturesRead(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if this is a route server network feature
	rsnf, ok := result.(*ixapi.RouteServerNetworkFeature)
	if !ok {
		return diag.Errorf("Received unexpected network feature type: %s",
			result.PolymorphicType())
	}
	if err := schemas.SetResourceData(rsnf, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(id)
	return nil
}
