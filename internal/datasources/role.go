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

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"required_fields": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: schemas.RoleSchema(),
				},
			},
		},
	}
}

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

	res.SetId(found.ID)
	res.Set("id", found.ID)
	res.Set("role", []interface{}{
		schemas.FlattenRole(found),
	})
	res.Set("required_fields", found.RequiredFields)

	return nil
}
