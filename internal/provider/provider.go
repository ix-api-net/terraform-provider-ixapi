package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/datasources"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/resources"
)

// Configuration
const (
	// EnvAPIKey environment variable for the API key
	// used by legacy auth.
	EnvAPIKey = "IX_API_KEY"
	// EnvAPISecret environment variable for the API secret
	// used by legacy auth.
	EnvAPISecret = "IX_API_SECRET"
	// EnvAPIHost is the environment variable for
	// the host implementing the API including the schema and
	// version.
	EnvAPIHost = "IX_API_HOST"
)

func init() {
	// Support markdown syntax in description
	schema.DescriptionKind = schema.StringMarkdown

	// Add defaults to exported descriptions
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

// New creates a new provider function
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"auth": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Default:     "legacy",
					Description: "Authentication schema used to log in to the API",
				},
				"api": &schema.Schema{
					Type:        schema.TypeString,
					Required:    true,
					Description: "API host, e.g. https://ixapi.myixp.example.com",
					DefaultFunc: schema.EnvDefaultFunc(EnvAPIHost, nil),
				},
				"api_key": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(EnvAPIKey, nil),
					Description: "Legacy auth: api key",
				},
				"api_secret": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc(EnvAPISecret, nil),
					Description: "Legacy auth: api secret",
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"ixapi_metro_areas":         datasources.NewMetroAreasDataSource(),
				"ixapi_metro_area":          datasources.NewMetroAreaDataSource(),
				"ixapi_metro_area_networks": datasources.NewMetroAreaNetworksDataSource(),
				"ixapi_metro_area_network":  datasources.NewMetroAreaNetworkDataSource(),
				"ixapi_roles":               datasources.NewRolesDataSource(),
				"ixapi_role":                datasources.NewRoleDataSource(),
				"ixapi_contacts":            datasources.NewContactsDataSource(),
				"ixapi_contact":             datasources.NewContactDataSource(),
				"ixapi_accounts":            datasources.NewAccountsDataSource(),
				"ixapi_account":             datasources.NewAccountDataSource(),
				"ixapi_pops":                datasources.NewPopsDataSource(),
				"ixapi_pop":                 datasources.NewPopDataSource(),
				"ixapi_facilities":          datasources.NewFacilitiesDataSource(),
				"ixapi_facility":            datasources.NewFacilityDataSource(),
				"ixapi_connections":         datasources.NewConnectionsDataSource(),
				"ixapi_connection":          datasources.NewConnectionDataSource(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"ixapi_contact": resources.NewContactResource(),
				"ixapi_account": resources.NewAccountResource(),
			},

			ConfigureContextFunc: configure,
		}

		return p
	}
}

// Configuration
func configure(
	ctx context.Context,
	res *schema.ResourceData,
) (any, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Get API credentials
	key := res.Get("api_key").(string)
	secret := res.Get("api_secret").(string)
	host := res.Get("api").(string)

	if key == "" || secret == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "IX-API client credentials missing",
			Detail:   "To access the API, a client key and secret are required",
		})
		return nil, diags
	}

	// Create client and authenticate with legacy strategy
	client := ixapi.NewClient(host)
	if err := client.Authenticate(ctx, &ixapi.AuthAPIKeySecret{
		Key:    key,
		Secret: secret,
	}); err != nil {
		return nil, diag.FromErr(err)
	}

	// Make test request to see if we are authenticated

	return client, nil
}
