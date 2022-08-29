package datasources

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewAccountsDataSource creates a new account data source
func NewAccountsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `accounts` data source to find accounts",

		ReadContext: accountsRead,

		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(),
			"name":             schemas.DataSourceQuery(),
			"external_ref":     schemas.DataSourceQuery(),
			"accounts": schemas.IntoDataSourceResultsSchema(
				schemas.AccountSchema(),
			),
		},
	}
}

// Operations

// get list of accounts
func accountsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	var managing, name, ref string

	// Filters
	val, ok := res.GetOk("managing_account")
	if ok {
		managing = val.(string)
	}
	val, ok = res.GetOk("name")
	if ok {
		name = val.(string)
	}
	val, ok = res.GetOk("external_ref")
	if ok {
		ref = val.(string)
	}

	// Get all acccounts and filter locally
	accounts, err := api.AccountsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.Account, 0, len(accounts))
	for _, acc := range accounts {
		if managing != "" && acc.ManagingAccount != nil && *acc.ManagingAccount != managing {
			continue
		}
		if name != "" && strings.ToLower(acc.Name) != strings.ToLower(name) {
			continue
		}
		if ref != "" && acc.ExternalRef != nil && *acc.ExternalRef != ref {
			continue
		}
		filtered = append(filtered, acc)
	}

	state := make([]interface{}, len(filtered))
	for i, acc := range filtered {
		state[i] = schemas.FlattenAccount(acc)
	}

	res.Set("accounts", state)
	res.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}

// NewAccountDataSource creates a new account data source for a
// single account referenced by unique name or external id
func NewAccountDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `account` data source to use a single",
		ReadContext: accountRead,
		Schema:      schemas.IntoDataSourceSchema(schemas.AccountSchema()),
	}
}

// Get single account
func accountRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	var name, ref string

	// Filters
	val, ok := res.GetOk("name")
	if ok {
		name = val.(string)
	}
	val, ok = res.GetOk("external_ref")
	if ok {
		ref = val.(string)
	}

	if name == "" && ref == "" {
		return diag.Errorf("at least name or external_ref is required")
	}

	// Get all acccounts and filter locally
	accounts, err := api.AccountsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	filtered := make([]*ixapi.Account, 0, len(accounts))
	for _, acc := range accounts {
		if name != "" && strings.ToLower(acc.Name) != strings.ToLower(name) {
			continue
		}
		if ref != "" && (acc.ExternalRef == nil || (acc.ExternalRef != nil && *acc.ExternalRef != ref)) {
			continue
		}
		filtered = append(filtered, acc)
	}

	if len(filtered) == 0 {
		return diag.Errorf("no account could be found")
	}
	if len(filtered) > 1 {
		return diag.Errorf("account is ambiguous - must have unique name or external_ref: %v", filtered)
	}

	account := filtered[0]

	schemas.AccountSetResourceData(account, res)
	res.SetId(account.ID)

	return nil
}
