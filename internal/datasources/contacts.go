package datasources

import (
	"context"

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
			"managing_account": schemas.DataSourceQuery(
				"Filter by account managing the contact"),
			"consuming_account": schemas.DataSourceQuery(
				"Filter by account the contacts are associated with"),
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
	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")

	// Fetch contacts
	contacts, err := api.ContactsList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.Contact, 0, len(contacts))
	for _, c := range contacts {
		if hasManagingAccount && c.ManagingAccount != managingAccount.(string) {
			continue
		}
		if hasConsumingAccount && c.ConsumingAccount != consumingAccount.(string) {
			continue
		}
		filtered = append(filtered, c)
	}

	// Make state
	state, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("contacts", state); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())
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
			schemas.DataSourceID(),
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
	id := res.Get("id").(string)
	contact, err := api.ContactsRead(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := schemas.SetResourceData(contact, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(id)
	return nil
}
