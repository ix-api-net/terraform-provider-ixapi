package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterPrefixListsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_prefix_lists` data source to find available prefix lists",
		ReadContext: cloudRouterPrefixListsRead,
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by managing account ID"),
			"prefix_lists": schemas.IntoDataSourceResultsSchema(
				schemas.PrefixListSchema(),
			),
		},
	}
}

func cloudRouterPrefixListsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	managingAccount := ""
	if ma, ok := res.GetOk("managing_account"); ok {
		managingAccount = ma.(string)
	}

	all, err := api.PrefixListsList(ctx, managingAccount)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("prefix_lists", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

func NewCloudRouterPrefixListDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_prefix_list` data source to get a single prefix list by ID or name",
		ReadContext: cloudRouterPrefixListRead,
		Schema:      schemas.IntoDataSourceSchema(schemas.PrefixListSchema()),
	}
}

func cloudRouterPrefixListRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	id, hasID := res.GetOk("id")
	name, hasName := res.GetOk("name")

	if !hasID && !hasName {
		return diag.Errorf("either `id` or `name` is required")
	}

	var prefixList *ixapi.PrefixList
	if hasID {
		pl, err := api.PrefixListsRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		prefixList = pl
	} else {
		managingAccount := ""
		if ma, ok := res.GetOk("managing_account"); ok {
			managingAccount = ma.(string)
		}

		result, err := api.PrefixListsList(ctx, managingAccount)
		if err != nil {
			return diag.FromErr(err)
		}

		var found *ixapi.PrefixList
		for _, pl := range result {
			if pl.Name == name.(string) {
				found = pl
				break
			}
		}

		if found == nil {
			return diag.Errorf("no prefix list found with name %s", name.(string))
		}
		prefixList = found
	}

	if err := schemas.SetResourceData(prefixList, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(prefixList.ID)

	return nil
}
