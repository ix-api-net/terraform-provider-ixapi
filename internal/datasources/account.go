package datasources

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewAccountDataSource creates a new account data source
func NewAccountDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `account` data source to use a single",

		ReadContext: readAccount,

		Schema: map[string]*schema.Schema{
			"managing_account": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The `id` of a managing account. Can be used for creating a customer hierachy. *(Sensitive Property)* "},

			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of the account, how it gets represented in e.g. a \"customers list\". "},

			"legal_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Legal name of the organisation. Only required when it's different from the account name. *(Sensitive Property)* "},

			"billing_information": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemas.BillingInformationSchema,
				},
			},

			"external_ref": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Reference field, free to use for the API user. *(Sensitive Property)*"},

			"discoverable": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "The account will be included for all members of the ix in the list of accounts.  Only `id`, `name` and `present_in_metro_area_networks` are provided to other members."},

			"metro_area_network_presence": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"address": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemas.AddressSchema,
				},
			},

			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			/*
				"status": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
				},
			*/

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

// Get account
func readAccount(
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

	//	res.Set("account", []interface{}{schemas.FlattenAccount(account)})
	schemas.AccountSetResourceData(account, res)
	res.SetId(account.ID)

	return nil
}
