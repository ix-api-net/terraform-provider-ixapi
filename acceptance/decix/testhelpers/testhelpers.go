package testhelpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/provider"
)

// ProviderFactories returns a provider factory map for use in acceptance test cases.
func ProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"ixapi": func() (*schema.Provider, error) {
			return provider.New()(), nil
		},
	}
}

// ForeignProviderFactories serves the provider under two names, "ixapi" and
// "ixapiforeign", so a single test can configure two accounts on separately
// configured instances.
func ForeignProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"ixapi": func() (*schema.Provider, error) {
			return provider.New()(), nil
		},
		"ixapiforeign": func() (*schema.Provider, error) {
			return provider.New()(), nil
		},
	}
}

// RequireTestEnv returns the value of the named environment variable, or
// fatals the test immediately if it is not set.
func RequireTestEnv(t *testing.T, name string) string {
	t.Helper()
	val := os.Getenv(name)
	if val == "" {
		t.Fatalf("required environment variable %s is not set", name)
	}
	return val
}

// NotExists returns a Terraform state check function that fails if the named
// resource is still present in state.
func NotExists(resourceName string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[resourceName]; ok {
			return fmt.Errorf("resource %s should not exist in state", resourceName)
		}
		return nil
	}
}

// ProviderConfigOptions configures the provider block returned by
// ProviderConfig.
type ProviderConfigOptions struct {
	Name             string
	Alias            string
	Foreign          bool
	ExtensionEnabled bool
}

// ProviderConfig returns a Terraform provider block string for the ixapi
// provider, configured per opts.
func ProviderConfig(opts ProviderConfigOptions) string {
	apiKey, apiSecret := os.Getenv("TF_VAR_API_KEY"), os.Getenv("TF_VAR_API_SECRET")
	if opts.Foreign {
		apiKey, apiSecret = os.Getenv("FOREIGN_API_KEY"), os.Getenv("FOREIGN_API_SECRET")
	}

	name := opts.Name
	if name == "" {
		name = "ixapi"
	}
	aliasLine := ""
	if opts.Alias != "" {
		aliasLine = fmt.Sprintf("  alias      = %q\n", opts.Alias)
	}
	extensionLine := ""
	if opts.ExtensionEnabled {
		extensionLine = "  extension_de_cix_cloud_router_enabled = true\n"
	}

	return fmt.Sprintf(`
provider %[1]q {
%[2]s  api        = %[3]q
  api_key    = %[4]q
  api_secret = %[5]q
%[6]s}
`, name, aliasLine, os.Getenv("TF_VAR_API_URL"), apiKey, apiSecret, extensionLine)
}

// FreeDot1QVlanConfig returns HCL computing local.free_vlan<_label>, a
// dot1q VLAN ID not already used by an existing
// ixapi_network_service_configs_<dsType> ("p2p_vc" or "cloud_vc"). label
// disambiguates multiple lookups in one config; providerRef, if set,
// queries a specific provider configuration (e.g. "ixapiforeign.b_side").
func FreeDot1QVlanConfig(dsType, label, connectionExpr, providerRef string) string {
	suffix := ""
	if label != "" {
		suffix = "_" + label
	}
	providerLine := ""
	if providerRef != "" {
		providerLine = fmt.Sprintf("  provider           = %s\n", providerRef)
	}
	return fmt.Sprintf(`
data "ixapi_network_service_configs_%[1]s" "existing%[2]s" {
%[3]s  network_connection = %[4]s
}

locals {
  used_vlans%[2]s = concat([1], [
    for nsc in data.ixapi_network_service_configs_%[1]s.existing%[2]s.network_service_configs :
    nsc.vlan_config[0].vlan
    if length(nsc.vlan_config) > 0 && nsc.vlan_config[0].vlan_type == "dot1q" && nsc.vlan_config[0].vlan != 0
  ])
  free_vlan%[2]s = max(local.used_vlans%[2]s...) + 1
}
`, dsType, suffix, providerLine, connectionExpr)
}
