// Tests a simple exchange_lan peering workflow: creates a MAC and an
// exchange_lan network service config on a production connection
package exchange_lan

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func uniquePrivateASN() int {
	return 64512 + int(time.Now().UnixNano()%(65534-64512))
}

func peeringConfig(accountID, externalRef string, asn int) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{}) + fmt.Sprintf(`
data "ixapi_metro_area_network" "fra" {
  name = "FRA"
}

data "ixapi_connections" "fra" {
  metro_area_network = data.ixapi_metro_area_network.fra.id
}

data "ixapi_product_offerings_exchange_lan" "peering" {
  service_metro_area_network  = data.ixapi_metro_area_network.fra.id
  handover_metro_area_network = data.ixapi_metro_area_network.fra.id
}

locals {
  selected_connection = [
    for c in data.ixapi_connections.fra.connections :
    c if c.state == "production"
  ][0]

  min_offering_bandwidth = min([
    for o in data.ixapi_product_offerings_exchange_lan.peering.product_offerings :
    o.bandwidth_min
  ]...)

  selected_offering = [
    for o in data.ixapi_product_offerings_exchange_lan.peering.product_offerings :
    o if o.bandwidth_min == local.min_offering_bandwidth
  ][0]
}

%[4]s

resource "ixapi_mac" "peering" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  address           = "02:00:00:00:01:01"
}

resource "ixapi_network_service_config_exchange_lan" "peering" {
  managing_account   = %[1]q
  consuming_account  = %[1]q
  billing_account    = %[1]q
  network_service    = local.selected_offering.exchange_lan_network_service
  network_connection = local.selected_connection.id
  product_offering   = local.selected_offering.id
  role_assignments   = []
  listed             = true
  asns               = [%[3]d]
  macs               = [ixapi_mac.peering.id]
  capacity           = local.selected_offering.bandwidth_min
  purchase_order     = "l2-tf-acceptance-peering-po"
  external_ref       = %[2]q

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}
`, accountID, externalRef, asn,
		testhelpers.FreeDot1QVlanConfig("exchange_lan", "", "local.selected_connection.id", ""))
}

func TestAccPeeringSimple(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")
	asn := uniquePrivateASN()

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_network_service_config_exchange_lan.peering"),
			testhelpers.NotExists("ixapi_mac.peering"),
		),
		Steps: []resource.TestStep{
			{
				Config: peeringConfig(accountID, "l2-tf-acceptance-peering-simple", asn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_exchange_lan.peering", "id"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_exchange_lan.peering", "network_service"),
				),
			},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{})},
		},
	})
}
