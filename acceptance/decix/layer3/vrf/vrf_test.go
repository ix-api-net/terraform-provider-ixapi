// Tests the full Cloud Router (VRF) lifecycle: create a VRF with a cloud-VC
// and a P2P partner connection, update the P2P admin status, then tear down
// connections before removing the VRF.
package vrf_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

const discoveryLocals = `
data "ixapi_connections" "all" {}
data "ixapi_product_offerings_p2p_vc" "p2p_vc" {
  service_provider            = "DE-CIX Management GmbH"
  handover_metro_area_network = "202"
  service_metro_area_network  = "202"
  bandwidth                   = 50
}
data "ixapi_product_offerings_cloud_vc" "all_cloud_vc" {}
data "ixapi_de_cix_product_offerings_cloud_vrf" "all" {}
data "ixapi_de_cix_cloud_router_network_service_configs_p2p_vc" "pre_existing" {}

locals {
  selected_cr_offering = [
    for o in data.ixapi_de_cix_product_offerings_cloud_vrf.all.product_offerings :
    o if o.service_metro_area == "201" && o.contract_period == "P3Y"
  ][0]

  selected_connection = [
    for c in data.ixapi_connections.all.connections :
    c if c.state == "production"
  ][0]

  selected_p2p_offering = data.ixapi_product_offerings_p2p_vc.p2p_vc.product_offerings[0]

  selected_cloud_vc_offering = [
    for o in data.ixapi_product_offerings_cloud_vc.all_cloud_vc.product_offerings :
    o if o.service_metro_area_network == "202" && o.handover_metro_area_network == "202" && o.service_provider_workflow == "exchange_first" && o.service_provider == "AWS"
  ][0]

  used_dot1q_vlans = concat([1], [
    for nsc in data.ixapi_de_cix_cloud_router_network_service_configs_p2p_vc.pre_existing.cloud_router_network_service_configs :
    nsc.vlan_config[0].vlan
    if length(nsc.vlan_config) > 0 && nsc.vlan_config[0].vlan_type == "dot1q" && nsc.vlan_config[0].vlan != 0
  ])
  free_p2p_vc_vlan = max(local.used_dot1q_vlans...) + 1
}
`

func fullVRFConfig(accountID, p2pAdminStatus string) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{ExtensionEnabled: true}) + discoveryLocals + fmt.Sprintf(`
resource "ixapi_de_cix_cloud_router" "test_vrf" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  billing_account   = %[1]q
  product_offering  = local.selected_cr_offering.id
  asn               = 65893
  capacity          = local.selected_cr_offering.bandwidth_max
  external_ref      = "cr-tf-acceptance-test-cloud-router"
}

resource "ixapi_network_service_cloud_vc" "cloud_service" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  billing_account   = %[1]q
  product_offering  = local.selected_cloud_vc_offering.id
  cloud_key         = "999999999999"
  capacity          = local.selected_cloud_vc_offering.bandwidth_min
  external_ref      = "cr-tf-acceptance-test-cloud-network-service"
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "cloud_connection" {
  managing_account  = %[1]q
  billing_account   = %[1]q
  consuming_account = %[1]q
  cloud_router      = ixapi_de_cix_cloud_router.test_vrf.id
  network_service   = ixapi_network_service_cloud_vc.cloud_service.id
  address           = "10.0.2.1/30"
  bgp_neighbor      = "10.0.2.2"
  bgp_neighbor_asn  = 65000
  bgp_password      = "test-cloud-bgp-password"
  admin_status      = "enabled"
  bfd_enabled       = true
  handover          = 1
  external_ref      = "cr-tf-acceptance-test-cloud-vc"
}

resource "ixapi_network_service_p2p_vc" "p2p_service" {
  managing_account       = %[1]q
  consuming_account      = %[1]q
  billing_account        = %[1]q
  joining_member_account = %[1]q
  product_offering       = local.selected_p2p_offering.id
  external_ref           = "cr-tf-acceptance-test-p2p-network-service"
}

resource "ixapi_de_cix_cloud_router_network_service_config_p2p_vc" "partner_connection" {
  managing_account   = %[1]q
  billing_account    = %[1]q
  consuming_account  = %[1]q
  cloud_router       = ixapi_de_cix_cloud_router.test_vrf.id
  network_service    = ixapi_network_service_p2p_vc.p2p_service.id
  network_connection = local.selected_connection.id
  address            = "10.0.1.1/30"
  bgp_neighbor       = "10.0.1.2"
  bgp_neighbor_asn   = 64512
  bgp_password       = "test-bgp-password"
  admin_status       = %[2]q
  bfd_enabled        = true
  external_ref       = "cr-tf-acceptance-test-p2p-vc"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_p2p_vc_vlan
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }

  depends_on = [ixapi_de_cix_cloud_router_network_service_config_cloud_vc.cloud_connection]
}
`, accountID, p2pAdminStatus)
}

func vrfOnlyConfig(accountID string) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{ExtensionEnabled: true}) + fmt.Sprintf(`
data "ixapi_de_cix_product_offerings_cloud_vrf" "all" {}

locals {
  selected_cr_offering = [
    for o in data.ixapi_de_cix_product_offerings_cloud_vrf.all.product_offerings :
    o if o.service_metro_area == "201" && o.contract_period == "P3Y"
  ][0]
}

resource "ixapi_de_cix_cloud_router" "test_vrf" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  billing_account   = %[1]q
  product_offering  = local.selected_cr_offering.id
  asn               = 65893
  capacity          = local.selected_cr_offering.bandwidth_max
  external_ref      = "cr-tf-acceptance-test-cloud-router"
}
`, accountID)
}

func TestAccVRF(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_de_cix_cloud_router.test_vrf"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_network_service_config_cloud_vc.cloud_connection"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection"),
		),
		Steps: []resource.TestStep{
			{
				Config: fullVRFConfig(accountID, "enabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router.test_vrf", "asn", "65893"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router.test_vrf", "external_ref", "cr-tf-acceptance-test-cloud-router"),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_network_service_config_cloud_vc.cloud_connection",
						"external_ref", "cr-tf-acceptance-test-cloud-vc",
					),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_network_service_config_cloud_vc.cloud_connection",
						"bfd_enabled", "true",
					),
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection",
						"external_ref", "cr-tf-acceptance-test-p2p-vc",
					),
				),
			},
			{
				Config: fullVRFConfig(accountID, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection",
						"admin_status", "DISABLED",
					),
				),
			},
			{Config: vrfOnlyConfig(accountID)},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{ExtensionEnabled: true})},
		},
	})
}
