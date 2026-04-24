// Tests Cloud Router static routes attached to an existing production P2P NSC:
// create a specific-prefix route and an aggregate route, update the next-hop,
// then delete both.
package static_routes_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

const nscDiscovery = `
data "ixapi_de_cix_cloud_router_network_service_configs_p2p_vc" "all_p2p_vc" {}

locals {
  production_nscs = [
    for nsc in data.ixapi_de_cix_cloud_router_network_service_configs_p2p_vc.all_p2p_vc.cloud_router_network_service_configs :
    nsc if nsc.state == "production"
  ]
  target_nsc = local.production_nscs[0]
}
`

func staticRoutesConfig(nextHop string) string {
	return testhelpers.ProviderConfig() + nscDiscovery + fmt.Sprintf(`
resource "ixapi_de_cix_cloud_router_static_route" "specific_route" {
  name                    = "cr-tf-acceptance-specific-route"
  prefix                  = "198.51.100.0/24"
  next_hop                = %q
  network_service_configs = [local.target_nsc.id]
}

resource "ixapi_de_cix_cloud_router_static_route" "aggregate_route" {
  name                    = "cr-tf-acceptance-aggregate-route"
  prefix                  = "10.0.0.0/8"
  next_hop                = "aggregate"
  network_service_configs = [local.target_nsc.id]
}
`, nextHop)
}

func TestAccStaticRoutes(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		CheckDestroy: resource.ComposeTestCheckFunc(
			testhelpers.NotExists("ixapi_de_cix_cloud_router_static_route.specific_route"),
			testhelpers.NotExists("ixapi_de_cix_cloud_router_static_route.aggregate_route"),
		),
		Steps: []resource.TestStep{
			{
				Config: staticRoutesConfig("10.0.4.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_static_route.specific_route", "id"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.specific_route", "prefix", "198.51.100.0/24"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.specific_route", "next_hop", "10.0.4.2"),
					resource.TestCheckResourceAttrSet("ixapi_de_cix_cloud_router_static_route.aggregate_route", "id"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.aggregate_route", "prefix", "10.0.0.0/8"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.aggregate_route", "next_hop", "aggregate"),
				),
			},
			{
				Config: staticRoutesConfig("10.0.4.3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.specific_route", "next_hop", "10.0.4.3"),
					resource.TestCheckResourceAttr("ixapi_de_cix_cloud_router_static_route.aggregate_route", "next_hop", "aggregate"),
				),
			},
			{Config: testhelpers.ProviderConfig() + nscDiscovery},
		},
	})
}

