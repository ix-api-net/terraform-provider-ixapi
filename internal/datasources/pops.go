package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/filter"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewPopsDataSource creates a data source for a
// collection of pops
func NewPopsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `pops` data source to find available points of presence",
		ReadContext: crud.Read(popsRead),
		Schema: map[string]*schema.Schema{
			"facility": schemas.DataSourceQuery(
				"Filter by facility id, see facilities data source"),
			"metro_area_network": schemas.DataSourceQuery(
				"Filter by metro area network id, see related data source"),
			"pops": schemas.IntoDataSourceResultsSchema(
				schemas.PointOfPresenceSchema(),
			),
		},
	}
}

// Fetch pops from API
func fetchPops(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.PointOfPresence, error) {
	// Filter pops
	facility, hasFacility := res.GetOk("facility")
	metroAreaNetwork, hasMetroAreaNetwork := res.GetOk("metro_area_network")
	name, hasName := res.GetOk("name")

	// Query
	pops, err := api.PopsList(ctx)
	if err != nil {
		return nil, err
	}

	// Apply local filters
	filtered := make([]*ixapi.PointOfPresence, 0, len(pops))
	for _, pop := range pops {
		if filter.Missing(pop.Facility, facility, hasFacility) {
			continue
		}
		if filter.Missing(pop.MetroAreaNetwork, metroAreaNetwork, hasMetroAreaNetwork) {
			continue
		}
		if filter.Missing(pop.Name, name, hasName) {
			continue
		}
		filtered = append(filtered, pop)
	}

	return filtered, nil
}

// Fetch all pops
func popsRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {

	pops, err := fetchPops(ctx, res, api)
	if err != nil {
		return err
	}

	flat, err := schemas.FlattenModels(pops)
	if err != nil {
		return err
	}

	if err := res.Set("pops", flat); err != nil {
		return err
	}

	res.SetId(schemas.Timestamp())

	return nil
}

// NewPopDataSource creates a data source resource for
// referencing a single point of presence
func NewPopDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `pop` data source to reference a single point of presence.",
		ReadContext: crud.Read(popRead),
		Schema:      schemas.IntoDataSourceSchema(schemas.PointOfPresenceSchema()),
	}
}

// Retrieve a single pop
func popRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	var pop *ixapi.PointOfPresence

	// Get the pop ID and retrieve from API
	id, hasID := res.GetOk("id")
	if hasID && id != "" {
		result, err := api.PopsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		pop = result
	} else {
		results, err := fetchPops(ctx, res, api)
		if err != nil {
			return err
		}
		if len(results) == 0 {
			return fmt.Errorf("No point of presence did match the query")
		}
		if len(results) > 1 {
			return fmt.Errorf("The API returned more than one point of presence")
		}
		pop = results[0]
	}

	if err := schemas.SetResourceData(pop, res); err != nil {
		return err
	}
	res.SetId(pop.ID)
	return nil
}
