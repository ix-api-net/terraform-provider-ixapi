package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewDecixCloudRouterStaticRouteResource returns the schema.Resource for managing a static route attached to a Cloud Router VRF.
func NewDecixCloudRouterStaticRouteResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_static_route` resource to manage a static route attached to a cloud router VRF.",
		CreateContext: crud.Create(staticRouteCreate),
		UpdateContext: crud.Update(staticRouteUpdate),
		ReadContext:   crud.Read(staticRouteRead),
		DeleteContext: crud.Delete(staticRouteDelete),

		Schema: decixschemas.StaticRouteSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func staticRouteRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterStaticRouteRequest, error) {
	res := schemas.ResourceDataFrom(r)

	nscRaw := r.Get("network_service_configs").([]interface{})
	nscs := make([]string, len(nscRaw))
	for i, v := range nscRaw {
		nscs[i] = v.(string)
	}

	req := &ixapi.CloudRouterStaticRouteRequest{
		Name:                 res.GetString("name"),
		Prefix:               res.GetString("prefix"),
		NextHop:              res.GetString("next_hop"),
		NetworkServiceConfigs: nscs,
	}
	return req, nil
}

func staticRouteCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := staticRouteRequestFromResourceData(res)
	if err != nil {
		return err
	}

	route, err := api.DecixCloudRouterStaticRoutesCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(route.ID)

	return staticRouteRead(ctx, res, api)
}

func staticRouteRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	route, err := api.DecixCloudRouterStaticRoutesRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	if err := schemas.SetResourceData(route, res); err != nil {
		return err
	}
	return nil
}

func staticRouteUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	req, err := staticRouteRequestFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.DecixCloudRouterStaticRoutesUpdate(ctx, id, req)
	if err != nil {
		return err
	}
	return staticRouteRead(ctx, res, api)
}

func staticRouteDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	_, err := api.DecixCloudRouterStaticRoutesDelete(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return staticRouteRead(ctx, res, api)
}
