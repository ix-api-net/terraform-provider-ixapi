package datasources

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewMetroAreasDataSource creates a new metro area datasource.
func NewMetroAreasDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the metro_areas datasource to retrieve all metro areas",

		ReadContext: metroAreasRead,

		Schema: map[string]*schema.Schema{
			"metro_areas": schemas.IntoDataSourceResultsSchema(
				schemas.MetroAreaSchema(),
			),
		},
	}
}

func metroAreasRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	metroAreas, err := api.MetroAreasList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	flatMetroAreas, err := schemas.FlattenModels(metroAreas)
	if err != nil {
		return diag.FromErr(err)
	}
	res.Set("metro_areas", flatMetroAreas)

	// Assign pseudoID
	res.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}

// NewMetroAreaDataSource creates a new metro area datasource
// for a single metro area.
func NewMetroAreaDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `metro_area` datasource to retrieve a metro area by un_locode or iata_code",

		ReadContext: metroAreaRead,

		Schema: schemas.IntoDataSourceSchema(schemas.MetroAreaSchema()),
	}
}

func metroAreaRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Query filters
	var (
		unLoc string
		iata  string
	)
	val, ok := res.GetOk("un_locode")
	if ok {
		unLoc = strings.ToLower(val.(string))
	}
	val, ok = res.GetOk("iata_code")
	if ok {
		iata = strings.ToLower(val.(string))
	}

	if unLoc == "" && iata == "" {
		return diag.Errorf("one of `un_locode` or `iata_code` is required")
	}

	metroAreas, err := api.MetroAreasList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Filter metro areas
	var found *ixapi.MetroArea
	for _, met := range metroAreas {
		if unLoc != "" && strings.ToLower(met.UnLocode) == unLoc {
			found = met
			break
		}
		if iata != "" && strings.ToLower(met.IataCode) == iata {
			found = met
			break
		}
	}
	if found == nil {
		return diag.Errorf("a metro area could not be found")
	}

	res.SetId(found.ID)
	if err := schemas.SetResourceData(found, res); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
