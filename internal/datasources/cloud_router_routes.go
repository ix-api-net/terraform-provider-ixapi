package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterRoutesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_routes` data source to list routes in the VRF routing table",
		ReadContext: cloudRouterRoutesRead,
		Schema: map[string]*schema.Schema{
			"vrf": schemas.DataSourceQuery(
				"Filter by VRF ID"),
			"routes": schemas.IntoDataSourceResultsSchema(
				schemas.VrfRouteSchema(),
			),
		},
	}
}

func cloudRouterRoutesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	vrfID := ""
	if v, ok := res.GetOk("vrf"); ok {
		vrfID = v.(string)
	}

	all, err := api.VrfRoutesList(ctx, vrfID)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("routes", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}
