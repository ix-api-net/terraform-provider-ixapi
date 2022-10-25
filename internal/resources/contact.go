package resources

import (
	"context"
	"errors"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewContactResource creates a new resource schema
func NewContactResource() *schema.Resource {
	return &schema.Resource{
		Description: "A contact for an account",

		CreateContext: contactCreate,
		ReadContext:   contactRead,
		UpdateContext: contactUpdate,
		DeleteContext: contactDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: schemas.Combine(
			schemas.ContactSchema(),
			map[string]*schema.Schema{
				"roles": {
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
	for _, name := range roleNames {
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

	for _, assignment := range assignments {
		if assignment.Contact != contact {
			continue
		}
		for _, roleName := range roleNames {
			r := lookupRoleByName(roleName, roles)
			if r == nil {
				continue // skip unknown role
			}
			if assignment.Role != r.ID {
				continue // skip non match
			}

			// Delete role assignment
			_, err := api.RoleAssignmentsDestroy(ctx, assignment.ID)
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

	assignedRoles := schemas.MustStringListFromAny(res.Get("roles"))

	// Create contact and try to assign role
	req := contactRequestFromResourceData(res)
	contact, err := api.ContactsCreate(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Try to assign roles
	_, err = createRoleAssignments(ctx, api, contact.ID, assignedRoles)
	if err != nil {
		// Rollback
		diags := diag.FromErr(err)
		err := deleteRoleAssignments(ctx, api, contact.ID, assignedRoles)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
		_, err = api.ContactsDestroy(ctx, contact.ID)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
		return diags
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
	if err := schemas.SetResourceData(contact, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(res.Id())
	return nil
}

func diffRoleAssignments(
	roles []*ixapi.Role,
	prev []string,
	next []string,
) ([]string, []string) {
	deletes := []string{}
	creates := []string{}
	// Get deletions
	for _, p := range prev {
		found := false
		for _, n := range next {
			if p == n {
				found = true
				break
			}
		}
		if !found {
			deletes = append(deletes, p)
		}
	}
	// Get Creates
	for _, n := range next {
		found := false
		for _, p := range prev {
			if p == n {
				found = true
				break
			}
		}
		if found {
			continue
		}
		creates = append(creates, n)
	}
	return deletes, creates
}

// Update
func contactUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	contactID := res.Get("id").(string)

	req := contactPatchFromResourceData(res)
	_, err := api.ContactsPatch(ctx, res.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update role assignments
	if res.HasChange("roles") {

		roles, err := api.RolesList(ctx)
		if err != nil {
			return diag.FromErr(err)
		}
		prev, next := res.GetChange("roles")
		deletes, creates := diffRoleAssignments(
			roles,
			schemas.MustStringListFromAny(prev),
			schemas.MustStringListFromAny(next))
		err = deleteRoleAssignments(ctx, api, contactID, deletes)
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = createRoleAssignments(ctx, api, contactID, creates)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	// Refresh contact
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
