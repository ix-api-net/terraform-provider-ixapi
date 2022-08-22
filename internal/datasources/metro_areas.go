package datasources

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewMetroAreasDataSource creates a new metro area datasource.
func NewMetroAreasDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the metro_areas datasource to retrieve all metro areas",

		ReadContext: metroAreasRead,

		Schema: map[string]*schema.Schema{
			"metro_areas": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemas.MetroAreaSchema,
				},
			},
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

	res.Set("metro_areas", schemas.FlattenMetroAreas(metroAreas))

	// Assign pseudoID
	res.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return nil
}
