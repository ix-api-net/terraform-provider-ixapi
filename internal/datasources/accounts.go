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
			"managing_account": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"accounts": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: schemas.AccountSchema,
				},
			},
		},
	}
}

// Operations

// Read
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
