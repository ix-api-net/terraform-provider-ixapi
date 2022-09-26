package datasources

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewRoleAssignmentDataSource creates a role assignment datasource
func NewRoleAssignmentDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_role_assignment` data source to retrieve an assignment between a contact and a role. Assignments can be used in configs.",
		ReadContext: roleAssignmentRead,
		Schema: map[string]*schema.Schema{
			"role_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name (e.g. `noc`) of a role the contact is assigned to. ",
			},

			"contact": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The `id` of a contact the role is assigned to. ",
			},

			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The `id` of the assignment. Can be used in configs.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func fetchRoleByName(
	ctx context.Context,
	api *ixapi.Client,
	name string,
) (*ixapi.Role, error) {
	roles, err := api.RolesList(ctx)
	if err != nil {
		return nil, err
	}

	for _, role := range roles {
		// Match role by name (should be unique)
		if strings.ToLower(role.Name) == strings.ToLower(name) {
			return role, nil
		}
	}

	err = fmt.Errorf("a role with the name %s could not be found", name)
	return nil, err
}

func fetchRoleAssignment(
	ctx context.Context,
	api *ixapi.Client,
	roleID string,
	contactID string,
) (*ixapi.RoleAssignment, error) {
	qry := &ixapi.RoleAssignmentsListQuery{
		Contact: contactID,
		Role:    roleID,
	}
	assignments, err := api.RoleAssignmentsList(ctx, qry)
	if err != nil {
		return nil, err
	}

	for _, a := range assignments {
		if a.Role == roleID && a.Contact == contactID {
			return a, nil
		}
	}
	err = fmt.Errorf("a role assignment for the role (%s) and contact (%s) could not be found", roleID, contactID)
	return nil, err
}

// Fetch role assignment
func roleAssignmentRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Data
	roleName := res.Get("role").(string)   // Name
	contact := res.Get("contact").(string) // ID

	role, err := fetchRoleByName(ctx, api, roleName)
	if err != nil {
		return diag.FromErr(err)
	}

	// Fetch role assignments
	assignment, err := fetchRoleAssignment(ctx, api, role.ID, contact)
	if err != nil {
		return diag.FromErr(err)
	}

	if schemas.SetResourceData(assignment, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(assignment.ID)

	return nil
}

// NewRoleAssignmentsDataSource creates a data source for querying
// role assignments.
func NewRoleAssignmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_role_assignments` data source to query a list of role assigments. You can filter by contact and role.",
		ReadContext: roleAssignmentsRead,
		Schema: map[string]*schema.Schema{
			"contact":   schemas.DataSourceQuery("Filter by contact ID"),
			"role":      schemas.DataSourceQuery("Filter by role ID"),
			"role_name": schemas.DataSourceQuery("Filter by role name"),
			"role_assignments": schemas.IntoDataSourceResultsSchema(
				schemas.RoleAssignmentSchema()),
		},
	}
}

func roleAssignmentsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	contactID, hasContact := res.GetOk("contact")
	role, hasRole := res.GetOk("role")
	roleName, hasRoleName := res.GetOk("role_name")

	var roleID string
	if hasRole {
		roleID = role.(string)
	}

	if hasRoleName {
		r, err := fetchRoleByName(ctx, api, roleName.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		roleID = r.ID
		hasRole = true
	}

	results, err := api.RoleAssignmentsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.RoleAssignment, 0, len(results))
	for _, assignment := range results {
		if hasRole && assignment.Role != roleID {
			continue
		}
		if hasContact && assignment.Contact != contactID.(string) {
			continue
		}
		filtered = append(filtered, assignment)
	}

	flat, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("role_assignments", flat); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
