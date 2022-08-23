package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewMetroAreaDataSource creates a new metro area datasource.
func NewMetroAreaDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `metro_area` datasource to retrieve a metro area by un_locode or iata_code",

		ReadContext: metroAreaRead,

		Schema: map[string]*schema.Schema{
			"un_locode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iata_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"metro_area": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: schemas.MetroAreaSchema,
				},
			},
		},
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
	res.Set("id", found.ID)
	res.Set("metro_area", []interface{}{
		schemas.FlattenMetroArea(found),
	})

	return nil
}
