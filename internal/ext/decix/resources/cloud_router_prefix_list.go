package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewDecixCloudRouterPrefixListResource returns the schema.Resource for managing a prefix list for BGP route filtering.
func NewDecixCloudRouterPrefixListResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router_prefix_list` resource to manage a prefix list for BGP route filtering.",
		CreateContext: crud.Create(prefixListCreate),
		UpdateContext: crud.Update(prefixListUpdate),
		ReadContext:   crud.Read(prefixListRead),
		DeleteContext: crud.Delete(prefixListDelete),

		Schema: decixschemas.PrefixListSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func prefixListRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.PrefixListRequest, error) {
	res := schemas.ResourceDataFrom(r)

	matchListRaw := r.Get("match_list").([]interface{})
	matchList := make([]ixapi.PrefixMatch, len(matchListRaw))
	for i, item := range matchListRaw {
		matchMap := item.(map[string]interface{})
		match := ixapi.PrefixMatch{
			Prefix: matchMap["prefix"].(string),
		}
		if minLength, ok := matchMap["min_length"].(int); ok && minLength > 0 {
			match.MinLength = &minLength
		}
		if maxLength, ok := matchMap["max_length"].(int); ok && maxLength > 0 {
			match.MaxLength = &maxLength
		}
		matchList[i] = match
	}

	req := &ixapi.PrefixListRequest{
		Name:             res.GetString("name"),
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		MatchList:        matchList,
	}
	return req, nil
}

func prefixListCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := prefixListRequestFromResourceData(res)
	if err != nil {
		return err
	}

	prefixList, err := api.DecixCloudRouterPrefixListsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(prefixList.ID)

	return prefixListRead(ctx, res, api)
}

func prefixListRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	prefixList, err := api.DecixCloudRouterPrefixListsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	if err := schemas.SetResourceData(prefixList, res); err != nil {
		return err
	}
	return nil
}

func prefixListUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	req, err := prefixListRequestFromResourceData(res)
	if err != nil {
		return err
	}

	_, err = api.DecixCloudRouterPrefixListsUpdate(ctx, id, req)
	if err != nil {
		return err
	}
	return prefixListRead(ctx, res, api)
}

func prefixListDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	_, err := api.DecixCloudRouterPrefixListsDelete(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return prefixListRead(ctx, res, api)
}
