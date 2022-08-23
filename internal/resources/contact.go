package resources

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewContactResource creates a new resource schema
func NewContactResource() *schema.Resource {
	return &schema.Resource{
		Description: "A contact for an account",

		CreateContext: contactCreate,
		ReadContext:   contactRead,
		UpdateContext: contactUpdate,
		DeleteContext: contactDelete,

		Schema: schemas.ContactSchema,
	}
}

// Operations

func lookupRole(roles []*ixapi.Role, assignment map[string]interface{}) *ixapi.Role {
	var (
		id   string
		name string
	)
	val, ok := assignment["id"]
	if ok {
		id = val.(string)
	}
	val, ok = assignment["name"]
	if ok {
		name = strings.ToLower(val.(string))
	}

	for _, role := range roles {
		if id != "" && role.ID == id {
			return role
		}
		if name != "" && strings.ToLower(role.Name) == name {
			return role
		}
	}
	return nil
}

func getRoleByID(roles []*ixapi.Role, id string) *ixapi.Role {
	for _, r := range roles {
		if r.ID == id {
			return r
		}
	}
	return nil
}

// Create
func contactCreate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Get roles for assignment
	roles, err := api.RolesList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Lookup roles for assignment
	val, ok := res.GetOk("assigned_roles")
	assignRoles := []*ixapi.Role{}
	if ok {
		for _, assign := range val.([]interface{}) {
			role := lookupRole(roles, assign.(map[string]interface{}))
			if role == nil {
				return diag.Errorf("can not find role for assignment: %v", assign)
			}
			assignRoles = append(assignRoles, role)
		}
	}

	// Create contact and try to assign role
	req := schemas.ContactRequestFromResourceData(res)
	contact, err := api.ContactsCreate(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Try to assign roles
	for _, role := range assignRoles {
		_, err := api.RoleAssignmentsCreate(ctx, &ixapi.RoleAssignmentRequest{
			Role:    role.ID,
			Contact: contact.ID,
		})
		if err != nil {
			diags := diag.FromErr(err)
			// Rollback and return error
			if _, eerr := api.ContactsDestroy(ctx, contact.ID); eerr != nil {
				diags = append(diags, diag.FromErr(eerr)...)
			}
			return diags
		}
	}

	res.SetId(contact.ID)
	return contactRead(ctx, res, meta)
}

// Read
func contactRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Fetch contact
	var notFoundErr *ixapi.NotFoundError
	contact, err := api.ContactsRead(ctx, res.Id())
	if err != nil && errors.As(err, &notFoundErr) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	// Set resource data
	schemas.ContactSetResourceData(contact, res)

	// Get assigned roles
	roles, err := api.RolesList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	assignments, err := api.RoleAssignmentsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	assignedRoles := []interface{}{}
	for _, assignment := range assignments {
		if assignment.Contact != contact.ID {
			continue
		}

		role := getRoleByID(roles, assignment.Role)
		if role == nil {
			return diag.Errorf("role not found in assignment")
		}
		assignedRoles = append(assignedRoles, map[string]interface{}{
			"name":       role.Name,
			"id":         role.ID,
			"assignment": assignment.ID,
		})
	}

	res.Set("assigned_roles", assignedRoles)
	res.SetId(res.Id())
	return nil
}

func diffRoleAssignments(
	roles []*ixapi.Role,
	prev []interface{},
	next []interface{},
) ([]string, []string, error) {
	deletes := []string{}
	creates := []string{}

	// Get deletions
	for _, p := range prev {
		aCur := p.(map[string]interface{})
		found := false
		for _, n := range next {
			aNext := n.(map[string]interface{})
			if aNext["assignment"] == aCur["assignment"] {
				found = true
				break
			}
		}
		if !found {
			deletes = append(deletes, aCur["assignment"].(string))
		}
	}

	// Get Creates
	for _, n := range next {
		a := n.(map[string]interface{})

		var aID string
		val, ok := a["assignment"]
		if ok {
			aID = val.(string)
		}
		if aID != "" {
			continue
		}
		role := lookupRole(roles, a)
		if role == nil {
			return nil, nil, fmt.Errorf(
				"role not found for assignment: %v", a)
		}
		creates = append(creates, role.ID)
	}
	return deletes, creates, nil
}

// Update
func contactUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	req := schemas.ContactPatchFromResourceData(res)
	_, err := api.ContactsPatch(ctx, res.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update role assignments
	if res.HasChange("assigned_roles") {

		roles, err := api.RolesList(ctx)
		if err != nil {
			return diag.FromErr(err)
		}
		prev, next := res.GetChange("assigned_roles")
		if _, ok := prev.([]interface{}); !ok {
			prev = []interface{}{}
		}
		if _, ok := next.([]interface{}); !ok {
			next = []interface{}{}
		}
		deletes, creates, err := diffRoleAssignments(
			roles,
			prev.([]interface{}),
			next.([]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}

		for _, del := range deletes {
			if _, err := api.RoleAssignmentsDestroy(ctx, del); err != nil {
				return diag.FromErr(err)
			}
		}

		for _, roleID := range creates {
			if _, err := api.RoleAssignmentsCreate(ctx, &ixapi.RoleAssignmentRequest{
				Contact: res.Id(),
				Role:    roleID,
			}); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return contactRead(ctx, res, meta)
}

// Delete
func contactDelete(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Delete role assignments
	assignments, err := api.RoleAssignmentsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, assignment := range assignments {
		if assignment.Contact != res.Id() {
			continue
		}
		if _, err := api.RoleAssignmentsDestroy(ctx, assignment.ID); err != nil {
			return diag.FromErr(err)
		}
	}

	if _, err := api.ContactsDestroy(ctx, res.Id()); err != nil {
		return diag.FromErr(err)
	}

	res.SetId("") // contact is gone
	return nil
}
