package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	decixschemas "github.com/ix-api-net/terraform-provider-ixapi/internal/ext/decix/schemas"
)

// NewDecixCloudRoutersDataSource returns the schema.Resource for listing DE-CIX Cloud Routers.
func NewDecixCloudRoutersDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_routers` data source to find available DE-CIX Cloud ROUTER (VRF) instances",
		ReadContext: cloudRoutersRead,
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by managing account ID"),
			"consuming_account": schemas.DataSourceQuery(
				"Filter by consuming account ID"),
			"external_ref": schemas.DataSourceQuery(
				"Filter by external reference"),
			"cloud_routers": schemas.IntoDataSourceResultsSchema(
				decixschemas.CloudRouterSchema(),
			),
		},
	}
}

func cloudRoutersRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	qry := &ixapi.CloudRoutersListQuery{}
	if managingAccount, ok := res.GetOk("managing_account"); ok {
		qry.ManagingAccount = managingAccount.(string)
	}
	if consumingAccount, ok := res.GetOk("consuming_account"); ok {
		qry.ConsumingAccount = consumingAccount.(string)
	}
	if externalRef, ok := res.GetOk("external_ref"); ok {
		qry.ExternalRef = externalRef.(string)
	}

	all, err := api.DecixCloudRoutersList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("cloud_routers", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewDecixCloudRouterDataSource returns the schema.Resource for reading a single DE-CIX Cloud Router by ID or external ref.
func NewDecixCloudRouterDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router` data source to get a single DE-CIX Cloud ROUTER (VRF) by ID or external ref",
		ReadContext: cloudRouterRead,
		Schema:      schemas.IntoDataSourceSchema(decixschemas.CloudRouterSchema()),
	}
}

func cloudRouterRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	id, hasID := res.GetOk("id")
	ref, hasRef := res.GetOk("external_ref")

	if !hasID && !hasRef {
		return diag.Errorf("the cloud_router `id` or `external_ref` is required")
	}

	var cloudRouter *ixapi.CloudRouter
	if hasID {
		cr, err := api.DecixCloudRoutersRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		cloudRouter = cr
	} else {
		qry := &ixapi.CloudRoutersListQuery{
			ExternalRef: ref.(string),
		}
		result, err := api.DecixCloudRoutersList(ctx, qry)
		if err != nil {
			return diag.FromErr(err)
		}

		if len(result) == 0 {
			return diag.Errorf("no cloud router could be found")
		}
		if len(result) > 1 {
			return diag.Errorf("multiple cloud routers returned for this external_ref")
		}
		cloudRouter = result[0]
	}

	if err := schemas.SetResourceData(cloudRouter, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(cloudRouter.ID)

	return nil
}
