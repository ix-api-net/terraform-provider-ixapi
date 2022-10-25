package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewRolesDataSource creates a data source for querying roles
func NewRolesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to query IX-API roles",
		Schema: map[string]*schema.Schema{
			"contact": schemas.DataSourceQuery(
				"Filter by id of the contact to retrieve assigned roles"),
			"roles": schemas.IntoDataSourceResultsSchema(
				schemas.RoleSchema()),
		},
		ReadContext: rolesRead,
	}
}

// fetch all roles matching the query
func rolesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	contact, hasContact := res.GetOk("contact")

	qry := &ixapi.RolesListQuery{}
	if hasContact {
		qry.Contact = contact.(string)
	}
	results, err := api.RolesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	roles, err := schemas.FlattenModels(results)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("roles", roles); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewRoleDataSource creates a new role data source schema
func NewRoleDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `role` data source to get a specifc role",

		ReadContext: roleRead,

		Schema: schemas.IntoDataSourceSchema(schemas.RoleSchema()),
	}
}

// Retrieve a single role
func roleRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Get all roles and filter by name
	name := strings.ToLower(res.Get("name").(string))

	roles, err := api.RolesList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	var found *ixapi.Role
	for _, role := range roles {
		if strings.ToLower(role.Name) == name {
			found = role
			break
		}
	}
	if found == nil {
		return diag.Errorf("a role matching the name could not be found")
	}

	if err := schemas.SetResourceData(found, res); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
