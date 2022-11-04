package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/datasources"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/resources"
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
					Optional:    true,
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
				"ixapi_accounts":             datasources.NewAccountsDataSource(),
				"ixapi_account":              datasources.NewAccountDataSource(),
				"ixapi_connections":          datasources.NewConnectionsDataSource(),
				"ixapi_connection":           datasources.NewConnectionDataSource(),
				"ixapi_contacts":             datasources.NewContactsDataSource(),
				"ixapi_contact":              datasources.NewContactDataSource(),
				"ixapi_devices":              datasources.NewDevicesDataSource(),
				"ixapi_device":               datasources.NewDeviceDataSource(),
				"ixapi_facilities":           datasources.NewFacilitiesDataSource(),
				"ixapi_facility":             datasources.NewFacilityDataSource(),
				"ixapi_ips":                  datasources.NewIPsDataSource(),
				"ixapi_ip":                   datasources.NewIPDataSource(),
				"ixapi_macs":                 datasources.NewMacsDataSource(),
				"ixapi_mac":                  datasources.NewMacDataSource(),
				"ixapi_member_joining_rules": datasources.NewMemberJoiningRulesDataSource(),
				"ixapi_member_joining_rule":  datasources.NewMemberJoiningRuleDataSource(),
				"ixapi_metro_areas":          datasources.NewMetroAreasDataSource(),
				"ixapi_metro_area":           datasources.NewMetroAreaDataSource(),
				"ixapi_metro_area_networks":  datasources.NewMetroAreaNetworksDataSource(),
				"ixapi_metro_area_network":   datasources.NewMetroAreaNetworkDataSource(),
				"ixapi_roles":                datasources.NewRolesDataSource(),
				"ixapi_role":                 datasources.NewRoleDataSource(),
				"ixapi_role_assignments":     datasources.NewRoleAssignmentsDataSource(),
				"ixapi_role_assignment":      datasources.NewRoleAssignmentDataSource(),
				"ixapi_pops":                 datasources.NewPopsDataSource(),
				"ixapi_pop":                  datasources.NewPopDataSource(),

				"ixapi_network_features_route_server": datasources.NewNetworkFeaturesRouteServerDataSource(),
				"ixapi_network_feature_route_server":  datasources.NewNetworkFeatureRouteServerDataSource(),

				"ixapi_network_services_exchange_lan": datasources.NewNetworkServicesExchangeLanDataSource(),
				"ixapi_network_service_exchange_lan":  datasources.NewNetworkServiceExchangeLanDataSource(),
				"ixapi_network_services_p2p_vc":       datasources.NewNetworkServicesP2PDataSource(),
				"ixapi_network_service_p2p_vc":        datasources.NewNetworkServiceP2PDataSource(),
				"ixapi_network_services_p2mp_vc":      datasources.NewNetworkServicesP2MPDataSource(),
				"ixapi_network_service_p2mp_vc":       datasources.NewNetworkServiceP2MPDataSource(),
				"ixapi_network_services_mp2mp_vc":     datasources.NewNetworkServicesMP2MPDataSource(),
				"ixapi_network_service_mp2mp_vc":      datasources.NewNetworkServiceMP2MPDataSource(),
				"ixapi_network_services_cloud_vc":     datasources.NewNetworkServicesCloudDataSource(),
				"ixapi_network_service_cloud_vc":      datasources.NewNetworkServiceCloudDataSource(),

				"ixapi_product_offerings_connection":   datasources.NewProductOfferingsConnectionDataSource(),
				"ixapi_product_offering_connection":    datasources.NewProductOfferingConnectionDataSource(),
				"ixapi_product_offerings_exchange_lan": datasources.NewProductOfferingsExchangeLanDataSource(),
				"ixapi_product_offering_exchange_lan":  datasources.NewProductOfferingExchangeLanDataSource(),
				"ixapi_product_offerings_p2p_vc":       datasources.NewProductOfferingsP2PVCDataSource(),
				"ixapi_product_offering_p2p_vc":        datasources.NewProductOfferingP2PVCDataSource(),
				"ixapi_product_offerings_p2mp_vc":      datasources.NewProductOfferingsP2MPVCDataSource(),
				"ixapi_product_offering_p2mp_vc":       datasources.NewProductOfferingP2MPVCDataSource(),
				"ixapi_product_offerings_mp2mp_vc":     datasources.NewProductOfferingsMP2MPVCDataSource(),
				"ixapi_product_offering_mp2mp_vc":      datasources.NewProductOfferingMP2MPVCDataSource(),
				"ixapi_product_offerings_cloud_vc":     datasources.NewProductOfferingsCloudVCDataSource(),
				"ixapi_product_offering_cloud_vc":      datasources.NewProductOfferingCloudVCDataSource(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"ixapi_contact": resources.NewContactResource(),
				"ixapi_account": resources.NewAccountResource(),
				"ixapi_mac":     resources.NewMACResource(),

				"ixapi_network_service_p2p_vc":   resources.NewNetworkServiceP2PVCResource(),
				"ixapi_network_service_p2mp_vc":  resources.NewNetworkServiceP2MPVCResource(),
				"ixapi_network_service_mp2mp_vc": resources.NewNetworkServiceMP2MPVCResource(),
				"ixapi_network_service_cloud_vc": resources.NewNetworkServiceCloudVCResource(),

				"ixapi_network_service_config_exchange_lan": resources.NewNetworkServiceConfigExchangeLanResource(),

				"ixapi_network_feature_config_route_server": resources.NewNetworkFeatureConfigRouteServerResource(),

				"ixapi_ip_allocation_network_service_config": resources.NewIPAllocationNetworkServiceConfigResource(),
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
	host := os.Getenv(EnvAPIHost)
	hostCfg, hasHostCfg := res.GetOk("api")
	if hasHostCfg {
		host = hostCfg.(string)
	}
	if err := checkEnvConfig("api", host, EnvAPIHost); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	key := os.Getenv(EnvAPIKey)
	keyCfg, hasKeyCfg := res.GetOk("api_key")
	if hasKeyCfg {
		key = keyCfg.(string)
	}
	if err := checkEnvConfig("api_key", key, EnvAPIKey); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	secret := os.Getenv(EnvAPISecret)
	secretCfg, hasSecretCfg := res.GetOk("api_secret")
	if hasSecretCfg {
		secret = secretCfg.(string)
	}
	if err := checkEnvConfig("api_secret", secret, EnvAPISecret); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	if diags.HasError() {
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
	if err := checkAuthenticated(ctx, client); err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}

// Check if the config value is not an empty string.
// If that case make a hint, that the env can be used
// to set a valid value for the provider.
func checkEnvConfig(key, value, env string) error {
	if value == "" {
		return fmt.Errorf("`%s` is not configured, you can set using the `%s` environment variable", key, env)
	}
	return nil
}

// Make an API request to a common endpoint and check
// if the response is permission denied.
func checkAuthenticated(
	ctx context.Context,
	api *ixapi.Client,
) error {
	_, err := api.AccountsList(ctx, &ixapi.AccountsListQuery{
		ExternalRef: "check-authenticated-no-response-required",
	})
	return err
}
