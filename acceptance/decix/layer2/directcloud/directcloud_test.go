// Tests the DirectCLOUD use case: creates a cloud_vc network service and a
// cloud_vc network service config on a production connection, verifies
// import, updates the network service external_ref, then tears everything
// down.
package directcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

func directCloudConfig(accountID, nsExternalRef string) string {
	return testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{}) + fmt.Sprintf(`
data "ixapi_metro_area_network" "fra" {
  name = "FRA"
}

data "ixapi_connections" "fra" {
  metro_area_network = data.ixapi_metro_area_network.fra.id
}

data "ixapi_product_offerings_cloud_vc" "all_cloud_vc" {}

locals {
  selected_connection = [
    for c in data.ixapi_connections.fra.connections :
    c if c.state == "production"
  ][0]

  selected_cloud_vc_offering = [
    for o in data.ixapi_product_offerings_cloud_vc.all_cloud_vc.product_offerings :
    o if o.service_metro_area_network == data.ixapi_metro_area_network.fra.id && o.handover_metro_area_network == data.ixapi_metro_area_network.fra.id && o.service_provider_workflow == "exchange_first" && o.service_provider == "AWS"
  ][1]

  cloud_vlan = local.selected_cloud_vc_offering.provider_vlans == "multi" ? 1 : 0
}

%[3]s

resource "ixapi_network_service_cloud_vc" "directcloud" {
  managing_account  = %[1]q
  consuming_account = %[1]q
  billing_account   = %[1]q
  product_offering  = local.selected_cloud_vc_offering.id
  cloud_key         = "999999999999"
  capacity          = local.selected_cloud_vc_offering.bandwidth_min
  external_ref      = %[2]q
}

resource "ixapi_network_service_config_cloud_vc" "directcloud" {
  managing_account   = %[1]q
  consuming_account  = %[1]q
  billing_account    = %[1]q
  network_service    = ixapi_network_service_cloud_vc.directcloud.id
  network_connection = local.selected_connection.id
  role_assignments   = []
  handover           = 1
  cloud_vlan         = local.cloud_vlan
  external_ref       = "l2-tf-acceptance-directcloud-nsc"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = local.free_vlan
  }

  lifecycle {
    ignore_changes = [vlan_config]
  }
}
`, accountID, nsExternalRef,
		testhelpers.FreeDot1QVlanConfig("cloud_vc", "", "local.selected_connection.id", ""))
}

func TestAccDirectCloud(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")
	accountID := testhelpers.RequireTestEnv(t, "ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_network_service_cloud_vc.directcloud"),
			testhelpers.NotExists("ixapi_network_service_config_cloud_vc.directcloud"),
		),
		Steps: []resource.TestStep{
			{
				Config: directCloudConfig(accountID, "l2-tf-acceptance-directcloud-ns"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_network_service_cloud_vc.directcloud", "id"),
					resource.TestCheckResourceAttr("ixapi_network_service_cloud_vc.directcloud", "cloud_key", "999999999999"),
					resource.TestCheckResourceAttr("ixapi_network_service_cloud_vc.directcloud", "external_ref", "l2-tf-acceptance-directcloud-ns"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_cloud_vc.directcloud", "id"),
					resource.TestCheckResourceAttrSet("ixapi_network_service_config_cloud_vc.directcloud", "network_connection"),
					resource.TestCheckResourceAttrPair(
						"ixapi_network_service_config_cloud_vc.directcloud", "network_service",
						"ixapi_network_service_cloud_vc.directcloud", "id",
					),
				),
			},
			{
				ResourceName:      "ixapi_network_service_cloud_vc.directcloud",
				ImportState:       true,
				ImportStateVerify: false,
			},
			{
				Config: directCloudConfig(accountID, "l2-tf-acceptance-directcloud-ns-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ixapi_network_service_cloud_vc.directcloud", "external_ref", "l2-tf-acceptance-directcloud-ns-updated"),
				),
			},
			{Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{})},
		},
	})
}
