package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

// NewCloudRouterBGPStateDataSource returns the schema.Resource for reading the BGP session state of a Cloud Router network service config.
func NewCloudRouterBGPStateDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_de_cix_cloud_router_bgp_state` data source to fetch the BGP session state for a cloud router network service config.",
		ReadContext: cloudRouterBGPStateRead,

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
				Description: "BGP connection state (Active, Idle, Established, etc.)",
			},
		},
	}
}

func cloudRouterBGPStateRead(
	ctx context.Context,
	d *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	nscID := d.Get("nsc_id").(string)

	bgpState, err := api.CloudRouterNetworkServiceConfigGetBGPState(ctx, nscID)
	if err != nil {
		return diag.FromErr(err)
	}

	if bgpState == nil {
		d.SetId("")
		return nil
	}

	d.SetId(nscID)
	d.Set("state", bgpState.State)

	return nil
}
