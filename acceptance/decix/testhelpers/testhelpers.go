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

// ProviderConfig returns a Terraform provider block string configured from
// the standard TF_VAR_API_* environment variables.
func ProviderConfig() string {
	return fmt.Sprintf(`
provider "ixapi" {
  api                                   = %q
  api_key                               = %q
  api_secret                            = %q
  extension_de_cix_cloud_router_enabled = true
}
`, os.Getenv("TF_VAR_API_URL"), os.Getenv("TF_VAR_API_KEY"), os.Getenv("TF_VAR_API_SECRET"))
}
