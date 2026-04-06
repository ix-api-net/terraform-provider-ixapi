package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

// NewDecixCloudRouterBFDStateDataSource returns the schema.Resource for reading the BFD session state of a Cloud Router network service config.
func NewDecixCloudRouterBFDStateDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_de_cix_cloud_router_bfd_state` data source to fetch the BFD session state for a cloud router network service config.",
		ReadContext: cloudRouterBFDStateRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Same as nsc_id",
			},
			"nsc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network service config ID",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "BFD session state (Up, Down, AdminDown, Init)",
			},
		},
	}
}

func cloudRouterBFDStateRead(
	ctx context.Context,
	d *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	nscID := d.Get("nsc_id").(string)

	bfdState, err := api.DecixCloudRouterNetworkServiceConfigGetBFDState(ctx, nscID)
	if err != nil {
		return diag.FromErr(err)
	}

	if bfdState == nil {
		d.SetId("")
		return nil
	}

	d.SetId(nscID)
	d.Set("state", bfdState.State)

	return nil
}
