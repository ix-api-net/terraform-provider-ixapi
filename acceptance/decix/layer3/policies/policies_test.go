// Tests standalone Cloud Router prefix lists and routing policies: create
// two prefix lists and two policies, update local-preference and AS-path
// prepend counts, then delete everything.
package policies_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func policiesConfig(accountID string, localPref, asPrependCount int) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{ExtensionEnabled: true}) + fmt.Sprintf(`
resource "ixapi_de_cix_cloud_router_prefix_list" "rfc1918" {
  name              = "cr-tf-acceptance-test-rfc1918"
  managing_account  = %[1]q
  consuming_account = %[1]q

  match_list {
    prefix     = "192.168.0.0/16"
    max_length = 24
  }
  match_list {
    prefix     = "10.0.0.0/8"
    min_length = 16
    max_length = 24
  }
  match_list {
    prefix     = "172.16.0.0/12"
    max_length = 24
  }
}

resource "ixapi_de_cix_cloud_router_prefix_list" "customer_networks" {
  name              = "cr-tf-acceptance-test-customer"
  managing_account  = %[1]q
  consuming_account = %[1]q

  match_list {
    prefix     = "203.0.113.0/24"
    min_length = 24
    max_length = 32
  }
  match_list {
    prefix     = "198.51.100.0/24"
    max_length = 28
  }
}

resource "ixapi_de_cix_cloud_router_policy" "accept_private" {
  name              = "cr-tf-acceptance-test-accept"
  managing_account  = %[1]q
  consuming_account = %[1]q

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.rfc1918.name
    action {
      filter           = "accept"
      local_preference = %[2]d
    }
  }
  entries {
    sequence_number = 20
    action {
      filter = "reject"
    }
  }
}

resource "ixapi_de_cix_cloud_router_policy" "reject_customer" {
  name              = "cr-tf-acceptance-test-reject"
  managing_account  = %[1]q
  consuming_account = %[1]q

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.customer_networks.name
    action {
      filter = "reject"
      as_path_prepend {
        count = %[3]d
      }
    }
  }
  entries {
    sequence_number = 20
    action {
      filter           = "continue"
      local_preference = 50
    }
  }
}
`, accountID, localPref, asPrependCount)
}

func TestAccPolicies(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_de_cix_cloud_router_prefix_list.rfc1918"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_prefix_list.customer_networks"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_policy.accept_private"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_policy.reject_customer"),
		),
		Steps: []resource.TestStep{
			{
				Config: policiesConfig(accountID, 100, 3),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_prefix_list.rfc1918", "id"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_prefix_list.rfc1918", "match_list.#", "3"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_prefix_list.customer_networks", "id"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_prefix_list.customer_networks", "match_list.#", "2"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_policy.accept_private", "id"),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_policy.accept_private",
						"entries.0.action.0.local_preference", "100",
					),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_policy.reject_customer", "id"),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_policy.reject_customer",
						"entries.0.action.0.as_path_prepend.0.count", "3",
					),
				),
			},
			{
				Config: policiesConfig(accountID, 150, 5),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_policy.accept_private",
						"entries.0.action.0.local_preference", "150",
					),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_policy.reject_customer",
						"entries.0.action.0.as_path_prepend.0.count", "5",
					),
				),
			},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{ExtensionEnabled: true})},
		},
	})
}
