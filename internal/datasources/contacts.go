package datasources

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewContactsDataSource creates a contacts data source schema
func NewContactsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "The contacts data source can be used to load contacts for an account",

		ReadContext: contactsRead,

		Schema: map[string]*schema.Schema{
			"managing_account":  schemas.DataSourceQuery(),
			"consuming_account": schemas.DataSourceQuery(),
			"contacts": schemas.IntoDataSourceResultsSchema(
				schemas.ContactSchema(),
			),
		},
	}
}

// Read and query the contacts resource
func contactsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	var (
		managing  string
		consuming string
	)

	val, ok := res.GetOk("managing_account")
	if ok {
		managing = val.(string)
	}
	val, ok = res.GetOk("consuming_account")
	if ok {
		consuming = val.(string)
	}

	// Fetch contacts
	contacts, err := api.ContactsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Load role assignments
	assignments, err := api.RoleAssignmentsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Load roles
	roles, err := api.RolesList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Filter locally
	filtered := make([]*ixapi.Contact, 0, len(contacts))
	for _, c := range contacts {
		if managing != "" && c.ManagingAccount != managing {
			continue
		}
		if consuming != "" && c.ConsumingAccount != consuming {
			continue
		}
		filtered = append(filtered, c)
	}

	// Make state
	state := make([]interface{}, len(filtered))
	for i, contact := range filtered {
		c, err := schemas.FlattenModel(contact)
		if err != nil {
			return diag.FromErr(err)
		}

		// Get roles for contact and flatten assigned roles
		assigned := []interface{}{}
		for _, a := range assignments {
			if a.Contact != contact.ID {
				continue
			}
			var role *ixapi.Role
			for _, r := range roles {
				if r.ID == a.Role {
					role = r
					break
				}
			}
			if role == nil {
				return diag.Errorf("could not resolve role in assignment")
			}
			assigned = append(assigned, map[string]interface{}{
				"name":       role.Name,
				"id":         role.ID,
				"assignment": a.ID,
			})

		}
		c["assigned_roles"] = assigned
		state[i] = c
	}

	res.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	res.Set("contacts", state)

	return nil
}

// NewContactDataSource creates a data source for a contact
// referenced by ID or external ref
func NewContactDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `contact` data source to fetch a single contact.",
		ReadContext: contactRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(schemas.ContactSchema()),
			map[string]*schema.Schema{
				"role_id": &schema.Schema{
					Type:        schema.TypeString,
					Description: "query contact by role id, only with consuming_account",
					Optional:    true,
				},
				"role": &schema.Schema{
					Type:        schema.TypeString,
					Description: "query contact by role name, only with consuming_account",
					Optional:    true,
				},
			},
		),
	}
}

// Retrieve a single contact.
func contactRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	id, hasID := res.GetOk("id")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	roleID, hasRoleID := res.GetOk("role_id")
	roleName, hasRoleName := res.GetOk("role")

	// Query
	qry := &ixapi.ContactsListQuery{}
	if hasID {
		qry.ID = []string{id.(string)}
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasConsumingAccount {
		qry.ConsumingAccount = consumingAccount.(string)
	}

	var queryRole *ixapi.Role
	if hasRoleID {
		result, err := api.RolesRead(ctx, roleID.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		queryRole = result
	} else if hasRoleName {
		results, err := api.RolesList(ctx, &ixapi.RolesListQuery{
			Name: roleName.(string),
		})
		if err != nil {
			return diag.FromErr(err)
		}
		filtered := make([]*ixapi.Role, 0, len(results))
		for _, role := range results {
			if role.Name != roleName.(string) {
				continue
			}
			filtered = append(filtered, role)
		}
		if len(filtered) == 0 {
			return diag.Errorf("a role with the name %s could not be found",
				roleName.(string))
		}
		if len(filtered) > 1 {
			return diag.Errorf("multiple roles were returned for %s",
				roleName.(string))
		}
		queryRole = filtered[0]
	}

	results, err := api.ContactsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Further filter roles
	filtered := make([]*ixapi.Contact, 0, len(results))
	for _, contact := range results {
		if hasID && contact.ID != id.(string) {
			continue
		}

		if queryRole != nil {
			// Check role assignments
			assignments, err := api.RoleAssignmentsList(ctx, &ixapi.RoleAssignmentsListQuery{
				Contact: contact.ID,
				Role:    queryRole.ID,
			})
			if err != nil {
				return diag.FromErr(err)
			}
			if len(assignments) == 0 {
				continue // no match
			}
		}

		filtered = append(filtered, contact)
	}

	if len(filtered) == 0 {
		return diag.Errorf("no matching contact could be found")
	}
	if len(filtered) > 1 {
		return diag.Errorf("multiple contacts were returned")
	}

	if err := schemas.SetResourceData(filtered[0], res); err != nil {
		return diag.FromErr(err)
	}

	res.SetId(filtered[0].ID)

	return nil
}
