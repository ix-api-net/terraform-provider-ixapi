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
			"managing_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"consuming_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"contacts": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: schemas.ContactSchema(),
				},
			},
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
		c := schemas.FlattenContact(contact)

		// Get roles for contact
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
