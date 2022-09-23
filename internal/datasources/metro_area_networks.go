package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/filter"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewMetroAreaNetworksDataSource creates a data source for
// querying metro area networks.
func NewMetroAreaNetworksDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve a list of metro area networks filtered by name, metro area or service provider",
		ReadContext: metroAreaNetworksRead,
		Schema: map[string]*schema.Schema{
			"name":             schemas.DataSourceQuery(),
			"metro_area":       schemas.DataSourceQuery(),
			"service_provider": schemas.DataSourceQuery(),
			"metro_area_networks": schemas.IntoDataSourceResultsSchema(
				schemas.MetroAreaNetworkSchema()),
		},
	}
}

// Fetch filtered metro area networks
func fetchMetroAreaNetworks(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.MetroAreaNetwork, error) {
	all, err := api.MetroAreaNetworksList(ctx)
	if err != nil {
		return nil, err
	}

	// Apply filters
	name, hasName := res.GetOk("name")
	metroArea, hasMetroArea := res.GetOk("metro_area")
	serviceProvider, hasServiceProvider := res.GetOk("service_provider")

	filtered := make([]*ixapi.MetroAreaNetwork, 0, len(all))
	for _, man := range all {
		if filter.String(man.Name, name, hasName) {
			continue
		}
		if filter.String(man.MetroArea, metroArea, hasMetroArea) {
			continue
		}
		if filter.String(man.ServiceProvider, serviceProvider, hasServiceProvider) {
			continue
		}
		filtered = append(filtered, man)
	}
	return filtered, nil
}

// Fetch all metro area networks filtered by params
func metroAreaNetworksRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	networks, err := fetchMetroAreaNetworks(ctx, res, api)
	if err != nil {
		return diag.FromErr(err)
	}
	flat, err := schemas.FlattenModels(networks)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("metro_area_networks", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewMetroAreaNetworkDataSource creates a data source for
// retrieving a single metro area network, either identified
// by id or by a unique combination of metro_area, name and
// service provider
func NewMetroAreaNetworkDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieve a single metro area network identified by ID, or a combination of name, metro are and service provider. Result must be unique.",
		ReadContext: metroAreaNetworkRead,
		Schema:      schemas.IntoDataSourceSchema(schemas.MetroAreaNetworkSchema()),
	}
}

// Retrieve a single metro area network
func metroAreaNetworkRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	id, hasID := res.GetOk("id")

	var man *ixapi.MetroAreaNetwork
	if hasID {
		network, err := api.MetroAreaNetworksRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		man = network
	} else {
		networks, err := fetchMetroAreaNetworks(ctx, res, api)
		if err != nil {
			return diag.FromErr(err)
		}

		if len(networks) == 0 {
			return diag.Errorf("no metro area network could be found matching the criteria")
		}
		if len(networks) > 1 {
			return diag.Errorf("the metro area network could not uniquely identified")
		}
		man = networks[0]
	}

	if err := schemas.SetResourceData(man, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(man.ID)
	return nil
}
