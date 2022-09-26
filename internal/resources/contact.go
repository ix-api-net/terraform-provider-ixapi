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

		Schema: schemas.Combine(
			schemas.ContactSchema(),
			map[string]*schema.Schema{
				"assigned_roles": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		),
	}
}

// helper
func lookupRoleByName(name string, roles []*ixapi.Role) *ixapi.Role {
	for _, r := range roles {
		if strings.ToLower(r.Name) == strings.ToLower(name) {
			return r
		}
	}
	return nil
}

func lookupRoleByID(id string, roles []*ixapi.Role) *ixapi.Role {
	for _, r := range roles {
		if r.ID == id {
			return r
		}
	}
	return nil
}

// Requests

// contactRequestFromResourceData makes a new structured request
func contactRequestFromResourceData(r *schema.ResourceData) *ixapi.ContactRequest {
	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.ContactRequest{
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		ExternalRef:      res.GetStringOpt("external_ref"),
		Name:             res.GetStringOpt("name"),
		Telephone:        res.GetStringOpt("telephone"),
		Email:            res.GetStringOpt("email"),
	}
	return req
}

// contactPatchFromResourceData creates a contact update
func contactPatchFromResourceData(r *schema.ResourceData) *ixapi.ContactPatch {
	res := schemas.ResourceData{ResourceData: r}
	patch := &ixapi.ContactPatch{}
	if res.HasChange("managing_account") {
		patch.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		patch.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("name") {
		patch.Name = res.GetStringOpt("name")
	}
	if res.HasChange("telephone") {
		patch.Telephone = res.GetStringOpt("telephone")
	}
	if res.HasChange("email") {
		patch.Email = res.GetStringOpt("email")
	}
	return patch
}

// Operations

// Create role assignments by Name
func createRoleAssignments(
	ctx context.Context,
	api *ixapi.Client,
	contact string,
	roleNames []string,
) ([]*ixapi.RoleAssignment, error) {
	// Get roles for assignment
	roles, err := api.RolesList(ctx)
	if err != nil {
		return nil, err
	}
	assignments := []*ixapi.RoleAssignment{}
	for _, name := range assignedRoles {
		r := lookupRoleByName(name, roles)
		if r == nil {
			continue // Skip this role
		}
		// Create role assignment
		assignment, err := api.RoleAssignmentsCreate(ctx, &ixapi.RoleAssignmentRequest{
			Role:    r.ID,
			Contact: contact,
		})
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

// Delete role assignments
func deleteRoleAssignments(
	ctx context.Context,
	api *ixapi.Client,
	contact string,
	roleNames []string,
) error {
	// Get roles for assignment
	roles, err := api.RolesList(ctx)
	if err != nil {
		return err
	}

	assignments, err := api.RoleAssignmentsList(ctx)
	if err != nil {
		return err
	}

	for _, assignment := range assignements {
		if assignment.Contact != contact {
			continue
		}
		for _, roleName := range roleNames {
			r := lookupRoleByName(roleName)
			if r == nil {
				continue // skip unknown role
			}
			if assignment.Role != r.ID {
				continue // skip non match
			}

			// Delete role assignment
			_, err := api.RoleAssignmentDestroy(ctx, assignment.ID)
			if err != nil {
				return err
			}
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

	// Lookup role IDs for assignment
	assignedRoles := res.GetOk("assigned_roles").([]any)

	// Create contact and try to assign role
	req := contactRequestFromResourceData(res)
	contact, err := api.ContactsCreate(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Try to assign roles
	for _, role := range assignedRoleIDs {
		_, err := api.RoleAssignmentsCreate(ctx, &ixapi.RoleAssignmentRequest{
			Role:    role.ID,
			Contact: contact.ID,
		})
		if err != nil {
			diags := diag.FromErr(err)
			// Rollback and return error
			if _, err := api.ContactsDestroy(ctx, contact.ID); err != nil {
				diags = append(diags, diag.FromErr(err)...)
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
	schemas.SetResourceData(contact, res)

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

		role := lookupRoleByID(assignment.Role, roles)
		if role == nil {
			return diag.Errorf("role not found in assignment")
		}
		assignedRoles = append(assignedRoles, role.Name)
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

	req := contactPatchFromResourceData(res)
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
