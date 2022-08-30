package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

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
