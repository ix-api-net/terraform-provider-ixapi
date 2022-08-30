package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewPopsDataSource creates a data source for a
// collection of pops
func NewPopsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `pops` data source to find available points of presence",
		ReadContext: popsRead,
		Schema: map[string]*schema.Schema{
			"faciltiy":           schemas.DataSourceQuery(),
			"metro_area_network": schemas.DataSourceQuery(),
			"pops": schemas.IntoDataSourceResultsSchema(
				schemas.PointOfPresenceSchema(),
			),
		},
	}
}

func popsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	pops, err := api.PopsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Filter pops
	facility, hasFacility := res.GetOk("facility")
	metroAreaNetwork, hasMetroAreaNetwork := res.GetOk("metro_area_network")

	filtered := make([]*ixapi.PointOfPresence, 0, len(pops))
	for _, pop := range pops {
		if hasFacility && pop.Facility != facility.(string) {
			continue
		}
		if hasMetroAreaNetwork && pop.MetroAreaNetwork != metroAreaNetwork.(string) {
			continue
		}
		filtered = append(filtered, pop)
	}

	flat, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("pops", flat); err != nil {
		return diag.FromErr(err)
	}

	res.SetId(schemas.Timestamp())

	return nil
}

// NewPopDataSource creates a data source resource for
// referencing a single point of presence
func NewPopDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `pop` data source to reference a single point of presence.",
		ReadContext: popRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(schemas.PointOfPresenceSchema()),
			map[string]*schema.Schema{
				"id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		),
	}
}

// Retrieve a single pop
func popRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Get the pop ID and retrieve from API
	id := res.Get("id").(string)
	pop, err := api.PopsRead(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := schemas.SetResourceData(pop, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(pop.ID)
	return nil
}
