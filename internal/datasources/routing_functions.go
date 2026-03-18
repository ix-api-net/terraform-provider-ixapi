package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewRoutingFunctionsDataSource creates a data source for a collection of routing functions
func NewRoutingFunctionsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_routing_functions` data source to find available routing functions",
		ReadContext: routingFunctionsRead,
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by account managing the routing function"),
			"consuming_account": schemas.DataSourceQuery(
				"Filter by account consuming the routing function"),
			"external_ref": schemas.DataSourceQuery(
				"Filter by external reference"),
			"state": schemas.DataSourceQuery(
				"Filter by state"),
			"routing_functions": schemas.IntoDataSourceResultsSchema(
				schemas.RoutingFunctionSchema(),
			),
		},
	}
}

func routingFunctionsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	state, hasState := res.GetOk("state")

	qry := &ixapi.RoutingFunctionsListQuery{}
	if hasManagingAccount {
		qry.ManagingAccount = managingAccount.(string)
	}
	if hasConsumingAccount {
		qry.ConsumingAccount = consumingAccount.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasState {
		qry.State = state.(string)
	}

	all, err := api.RoutingFunctionsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.RoutingFunction, 0, len(all))
	for _, rf := range all {
		if hasManagingAccount && rf.ManagingAccount != managingAccount.(string) {
			continue
		}
		if hasConsumingAccount && rf.ConsumingAccount != consumingAccount.(string) {
			continue
		}
		if hasExternalRef {
			if rf.ExternalRef == nil || *rf.ExternalRef != externalRef.(string) {
				continue
			}
		}
		if hasState && rf.State != state.(string) {
			continue
		}
		filtered = append(filtered, rf)
	}

	routingFunctions, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("routing_functions", routingFunctions); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewRoutingFunctionDataSource creates a data source for a single routing function
func NewRoutingFunctionDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_routing_function` data source to get a single routing function by ID",
		ReadContext: routingFunctionRead,
		Schema:      schemas.IntoDataSourceSchema(schemas.RoutingFunctionSchema()),
	}
}

func routingFunctionRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	id, hasID := res.GetOk("id")
	if !hasID {
		return diag.Errorf("the routing function `id` is required")
	}

	rf, err := api.RoutingFunctionsRead(ctx, id.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := schemas.SetResourceData(rf, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(rf.ID)

	return nil
}
