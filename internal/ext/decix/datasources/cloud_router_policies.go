package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewCloudRouterPoliciesDataSource returns the schema.Resource for listing BGP routing policies.
func NewCloudRouterPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_policies` data source to find available policies",
		ReadContext: cloudRouterPoliciesRead,
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by managing account ID"),
			"policies": schemas.IntoDataSourceResultsSchema(
				decixschemas.PolicySchema(),
			),
		},
	}
}

func cloudRouterPoliciesRead(
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

	all, err := api.PoliciesList(ctx, managingAccount)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("policies", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewCloudRouterPolicyDataSource returns the schema.Resource for reading a single BGP routing policy by ID or name.
func NewCloudRouterPolicyDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_policy` data source to get a single policy by ID or name",
		ReadContext: cloudRouterPolicyRead,
		Schema:      schemas.IntoDataSourceSchema(decixschemas.PolicySchema()),
	}
}

func cloudRouterPolicyRead(
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

	var policy *ixapi.Policy
	if hasID {
		p, err := api.PoliciesRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		policy = p
	} else {
		managingAccount := ""
		if ma, ok := res.GetOk("managing_account"); ok {
			managingAccount = ma.(string)
		}

		result, err := api.PoliciesList(ctx, managingAccount)
		if err != nil {
			return diag.FromErr(err)
		}

		var found *ixapi.Policy
		for _, p := range result {
			if p.Name == name.(string) {
				found = p
				break
			}
		}

		if found == nil {
			return diag.Errorf("no policy found with name %s", name.(string))
		}
		policy = found
	}

	if err := schemas.SetResourceData(policy, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(policy.ID)

	return nil
}
