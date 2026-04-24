// Tests attaching and detaching BGP routing policies on a Cloud Router
// P2P network service config: create an NSC without a policy, attach a
// prefix-list-backed policy, remove the NSC while the policy remains, then
// clean up.
package nsc_policies_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

const discoveryLocals = `
data "ixapi_connections" "all" {}
data "ixapi_product_offerings_p2p_vc" "p2p" {
  service_provider            = "DE-CIX Management GmbH"
  handover_metro_area_network = "202"
  service_metro_area_network  = "202"
  bandwidth                   = 50
}
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

  selected_p2p_offering = data.ixapi_product_offerings_p2p_vc.p2p.product_offerings[0]

  used_dot1q_vlans = concat([1], [
    for nsc in data.ixapi_de_cix_cloud_router_network_service_configs_p2p_vc.pre_existing.cloud_router_network_service_configs :
    nsc.vlan_config[0].vlan
    if length(nsc.vlan_config) > 0 && nsc.vlan_config[0].vlan_type == "dot1q" && nsc.vlan_config[0].vlan != 0
  ])
  free_p2p_vc_vlan = max(local.used_dot1q_vlans...) + 1
}
`

func vrfAndP2PResources(accountID string) string {
	return fmt.Sprintf(`
resource "ixapi_de_cix_cloud_router" "test_vrf" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  billing_account   = %[1]q
  product_offering  = local.selected_cr_offering.id
  asn               = 65893
  capacity          = local.selected_cr_offering.bandwidth_max
  external_ref      = "cr-tf-acceptance-test-cloud-router"
}

resource "ixapi_network_service_p2p_vc" "p2p_service" {
  managing_account       = %[1]q
  consuming_account      = %[1]q
  billing_account        = %[1]q
  joining_member_account = %[1]q
  product_offering       = local.selected_p2p_offering.id
  external_ref           = "cr-tf-acceptance-test-p2p-network-service"
}
`, accountID)
}

func partnerConnectionResource(accountID, policyEgress string) string {
	return fmt.Sprintf(`
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
  admin_status       = "enabled"
  bfd_enabled        = true
  external_ref       = "cr-tf-acceptance-test-p2p-vc"
  policy_egress      = %[2]s

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_p2p_vc_vlan
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}
`, accountID, policyEgress)
}

func nscPolicyResources(accountID string) string {
	return fmt.Sprintf(`
resource "ixapi_de_cix_cloud_router_prefix_list" "nsc_prefixes" {
  name              = "cr-tf-acceptance-nsc-policy-prefixes"
  managing_account  = %[1]q
  consuming_account = %[1]q

  match_list {
    prefix     = "10.0.0.0/8"
    max_length = 24
  }
}

resource "ixapi_de_cix_cloud_router_policy" "nsc_policy" {
  name              = "cr-tf-acceptance-nsc-policy"
  managing_account  = %[1]q
  consuming_account = %[1]q

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.nsc_prefixes.name
    action {
      filter           = "accept"
      local_preference = 100
    }
  }
  entries {
    sequence_number = 20
    action {
      filter = "reject"
    }
  }
}
`, accountID)
}

func vrfOnlyConfig(accountID string) string {
	return testhelpers.ProviderConfig() + fmt.Sprintf(`
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

func TestAccNSCPolicies(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	base := testhelpers.ProviderConfig() + discoveryLocals + vrfAndP2PResources(accountID)
	withNSC := base + partnerConnectionResource(accountID, "null")
	withPolicy := base + nscPolicyResources(accountID) + partnerConnectionResource(accountID, "ixapi_de_cix_cloud_router_policy.nsc_policy.name")
	vrfWithPolicyNoNSC := base + nscPolicyResources(accountID)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_de_cix_cloud_router.test_vrf"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_policy.nsc_policy"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_prefix_list.nsc_prefixes"),
		),
		Steps: []resource.TestStep{
			{
				Config: withNSC,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router.test_vrf", "id"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router.test_vrf", "asn", "65893"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection", "id"),
				),
			},
			{
				Config: withPolicy,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_policy.nsc_policy", "id"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_prefix_list.nsc_prefixes", "id"),
					resource.TestCheckResourceAttrSet(
						"ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_connection", "policy_egress",
					),
				),
			},
			{
				Config: vrfWithPolicyNoNSC,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_policy.nsc_policy", "id"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_prefix_list.nsc_prefixes", "id"),
				),
			},
			{Config: vrfOnlyConfig(accountID)},
			{Config: testhelpers.ProviderConfig()},
		},
	})
}

