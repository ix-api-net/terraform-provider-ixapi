// Tests the VirtualPNI use case between a customer-owned port and a foreign
// port: side B marks one of its connections discoverable, side A creates a
// p2p_vc network service and config against the resulting offering, side B
// creates its own config on the discoverable connection, then side B toggles
// the connection to non-discoverable and everything is torn down (NSCs first then NS when possible).
package virtual_pni_foreign_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func bSideConnectionConfig(connectionID string, discoverable bool) string {
	return fmt.Sprintf(`
import {
  to       = ixapi_connection.b_side
  id       = %[1]q
  provider = ixapiforeign.b_side
}

data "ixapi_connection" "b_side_existing" {
  provider = ixapiforeign.b_side
  id       = %[1]q
}

resource "ixapi_connection" "b_side" {
  provider          = ixapiforeign.b_side
  managing_account  = data.ixapi_connection.b_side_existing.managing_account
  consuming_account = data.ixapi_connection.b_side_existing.consuming_account
  billing_account   = data.ixapi_connection.b_side_existing.billing_account
  mode              = data.ixapi_connection.b_side_existing.mode
  product_offering  = data.ixapi_connection.b_side_existing.product_offering
  port_quantity     = data.ixapi_connection.b_side_existing.port_quantity
  discoverable      = %[2]t
  role_assignments  = []

  lifecycle {
    ignore_changes = [
      external_ref,
      purchase_order,
      contract_ref,
      role_assignments,
      subscriber_side_demarcs,
      connecting_party,
    ]
  }
}
`, connectionID, discoverable)
}

func virtualPNIForeignConfig(accountID, foreignAccountID, aSideConnectionID, bSideConnectionID, nsExternalRef string, discoverable bool) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{Alias: "a_side", ExtensionEnabled: true}) + testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{Name: "ixapiforeign", Alias: "b_side", Foreign: true, ExtensionEnabled: true}) + bSideConnectionConfig(bSideConnectionID, discoverable) + fmt.Sprintf(`
data "ixapi_account" "foreign_self" {
  provider = ixapiforeign.b_side
  id       = %[2]q
}

data "ixapi_metro_area_network" "fra" {
  provider = ixapi.a_side
  name     = "FRA"
}

data "ixapi_product_offerings_p2p_vc" "from_foreign" {
  provider                    = ixapi.a_side
  service_provider            = data.ixapi_account.foreign_self.name
  handover_metro_area_network = data.ixapi_metro_area_network.fra.id

  depends_on = [ixapi_connection.b_side]
}

locals {
  selected_p2p_offering_id = try(data.ixapi_product_offerings_p2p_vc.from_foreign.product_offerings[0].id, "")
}

%[4]s
%[5]s

resource "ixapi_network_service_p2p_vc" "pni" {
  provider               = ixapi.a_side
  managing_account       = %[1]q
  consuming_account      = %[1]q
  billing_account        = %[1]q
  joining_member_account = %[2]q
  product_offering       = local.selected_p2p_offering_id
  display_name           = "l2-tf-acceptance-virtualpni-foreign"
  external_ref           = %[3]q

  lifecycle {
    ignore_changes = [joining_member_account, product_offering]
  }
}

resource "ixapi_network_service_config_p2p_vc" "side_a" {
  provider           = ixapi.a_side
  managing_account   = %[1]q
  consuming_account  = %[1]q
  billing_account    = %[1]q
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = %[6]q
  role_assignments   = []
  external_ref       = "l2-tf-acceptance-virtualpni-foreign-a"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan_a
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}

resource "ixapi_network_service_config_p2p_vc" "side_b" {
  provider           = ixapiforeign.b_side
  managing_account   = %[2]q
  consuming_account  = %[2]q
  billing_account    = %[2]q
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = ixapi_connection.b_side.id
  role_assignments   = []
  external_ref       = "l2-tf-acceptance-virtualpni-foreign-b"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan_b
  }

  depends_on = [ixapi_network_service_config_p2p_vc.side_a]

  lifecycle {
    ignore_changes = [vlan_config]
  }
}
`, accountID, foreignAccountID, nsExternalRef,
		testhelpers.FreeDot1QVlanConfig("p2p_vc", "a", fmt.Sprintf("%q", aSideConnectionID), "ixapi.a_side"),
		testhelpers.FreeDot1QVlanConfig("p2p_vc", "b", "ixapi_connection.b_side.id", "ixapiforeign.b_side"),
		aSideConnectionID)
}

func TestAccVirtualPNIForeign(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	foreignAccountID := testhelpers.RequireTestEnv(t, "FOREIGN_ACCOUNT_ID")
	testhelpers.RequireTestEnv(t, "FOREIGN_API_KEY")
	testhelpers.RequireTestEnv(t, "FOREIGN_API_SECRET")
	aSideConnectionID := testhelpers.RequireTestEnv(t, "A_SIDE_CONNECTION_ID")
	bSideConnectionID := testhelpers.RequireTestEnv(t, "B_SIDE_CONNECTION_ID")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ForeignProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_network_service_p2p_vc.pni"),
			testhelpers.NotExists("ixapi_network_service_config_p2p_vc.side_a"),
			testhelpers.NotExists("ixapi_network_service_config_p2p_vc.side_b"),
		),
		Steps: []resource.TestStep{
			{
				Config: virtualPNIForeignConfig(accountID, foreignAccountID, aSideConnectionID, bSideConnectionID, "l2-tf-acceptance-virtualpni-foreign-ns", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ixapi_connection.b_side", "discoverable", "true"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_p2p_vc.pni", "id"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_p2p_vc.side_a", "id"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_p2p_vc.side_b", "id"),
					resource.TestCheckResourceAttrPair(
						"ixapi_network_service_config_p2p_vc.side_a", "network_service",
						"ixapi_network_service_p2p_vc.pni", "id",
					),
					resource.TestCheckResourceAttrPair(
						"ixapi_network_service_config_p2p_vc.side_b", "network_service",
						"ixapi_network_service_p2p_vc.pni", "id",
					),
				),
			},
			{
				Config: virtualPNIForeignConfig(accountID, foreignAccountID, aSideConnectionID, bSideConnectionID, "l2-tf-acceptance-virtualpni-foreign-ns", false),
				Check:  resource.TestCheckResourceAttr("ixapi_connection.b_side", "discoverable", "false"),
			},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{Alias: "a_side", ExtensionEnabled: true}) + testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{Name: "ixapiforeign", Alias: "b_side", Foreign: true, ExtensionEnabled: true})},
		},
	})
}
