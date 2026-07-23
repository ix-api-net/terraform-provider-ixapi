// Tests the VirtualPNI use case between two ports owned by the same customer:
// creates a p2p_vc network service and two p2p_vc network service configs, one
// on each of two production connections with distinct VLANs,
// updates the network service external_ref, then tears everything down.
package virtual_pni_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func virtualPNIConfig(accountID, nsExternalRef string) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{}) + fmt.Sprintf(`
data "ixapi_metro_area_network" "fra" {
  name = "FRA"
}

data "ixapi_connections" "fra" {
  metro_area_network = data.ixapi_metro_area_network.fra.id
}

data "ixapi_product_offerings_p2p_vc" "p2p_vc" {
  service_provider            = "DE-CIX Management GmbH"
  handover_metro_area_network = data.ixapi_metro_area_network.fra.id
  service_metro_area_network  = data.ixapi_metro_area_network.fra.id
  bandwidth                   = 10
}

locals {
  production_connections = [
    for c in data.ixapi_connections.fra.connections :
    c if c.state == "production"
  ]
  connection_a = local.production_connections[0]
  connection_b = local.production_connections[1]

  selected_p2p_offering = data.ixapi_product_offerings_p2p_vc.p2p_vc.product_offerings[0]
}

%[3]s
%[4]s

resource "ixapi_network_service_p2p_vc" "pni" {
  managing_account       = %[1]q
  consuming_account      = %[1]q
  billing_account        = %[1]q
  joining_member_account = %[1]q
  product_offering       = local.selected_p2p_offering.id
  display_name           = "l2-tf-acceptance-virtualpni"
  external_ref           = %[2]q

  lifecycle {
    ignore_changes = [joining_member_account]
  }
}

resource "ixapi_network_service_config_p2p_vc" "side_a" {
  managing_account   = %[1]q
  consuming_account  = %[1]q
  billing_account    = %[1]q
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = local.connection_a.id
  role_assignments   = []
  external_ref       = "l2-tf-acceptance-virtualpni-a"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan_a
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}

resource "ixapi_network_service_config_p2p_vc" "side_b" {
  managing_account   = %[1]q
  consuming_account  = %[1]q
  billing_account    = %[1]q
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = local.connection_b.id
  role_assignments   = []
  external_ref       = "l2-tf-acceptance-virtualpni-b"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan_b
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}
`, accountID, nsExternalRef,
		testhelpers.FreeDot1QVlanConfig("p2p_vc", "a", "local.connection_a.id", ""),
		testhelpers.FreeDot1QVlanConfig("p2p_vc", "b", "local.connection_b.id", ""))
}

func TestAccVirtualPNI(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_network_service_p2p_vc.pni"),
			testhelpers.NotExists("ixapi_network_service_config_p2p_vc.side_a"),
			testhelpers.NotExists("ixapi_network_service_config_p2p_vc.side_b"),
		),
		Steps: []resource.TestStep{
			{
				Config: virtualPNIConfig(accountID, "l2-tf-acceptance-virtualpni-ns"),
				Check: resource.ComposeTestCheckFunc(
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
				ResourceName:      "ixapi_network_service_p2p_vc.pni",
				ImportState:       true,
				ImportStateVerify: false,
			},
			{
				Config: virtualPNIConfig(accountID, "l2-tf-acceptance-virtualpni-ns-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ixapi_network_service_p2p_vc.pni", "external_ref", "l2-tf-acceptance-virtualpni-ns-updated"),
				),
			},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{})},
		},
	})
}
