package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewCloudRouterStaticRoutesDataSource returns the schema.Resource for listing static routes attached to a Cloud Router.
func NewCloudRouterStaticRoutesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_static_routes` data source to list static routes, optionally filtered by VRF or network service config",
		ReadContext: cloudRouterStaticRoutesRead,
		Schema: map[string]*schema.Schema{
			"vrf": schemas.DataSourceQuery(
				"Filter by VRF ID"),
			"network_service_config": schemas.DataSourceQuery(
				"Filter by network service config ID"),
			"static_routes": schemas.IntoDataSourceResultsSchema(
				decixschemas.StaticRouteSchema(),
			),
		},
	}
}

func cloudRouterStaticRoutesRead(
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
	nscID := ""
	if v, ok := res.GetOk("network_service_config"); ok {
		nscID = v.(string)
	}

	all, err := api.StaticRoutesList(ctx, vrfID, nscID)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("static_routes", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewCloudRouterStaticRouteDataSource returns the schema.Resource for reading a single static route by ID.
func NewCloudRouterStaticRouteDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_static_route` data source to get a single static route by ID",
		ReadContext: cloudRouterStaticRouteRead,
		Schema:      schemas.IntoDataSourceSchema(decixschemas.StaticRouteSchema()),
	}
}

func cloudRouterStaticRouteRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	id, hasID := res.GetOk("id")
	if !hasID {
		return diag.Errorf("`id` is required")
	}

	route, err := api.StaticRoutesRead(ctx, id.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := schemas.SetResourceData(route, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(route.ID)

	return nil
}
