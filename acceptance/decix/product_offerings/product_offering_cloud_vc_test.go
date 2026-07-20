package product_offering_cloud_vc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ix-api-net/terraform-provider-ixapi/acceptance/decix/testhelpers"
)

const cloudVCOfferingsConfig = `
data "ixapi_product_offerings_cloud_vc" "filtered" {
  service_provider                 = "AWS"
  bandwidth                        = 50
  contract_period                  = "P1M"
  delivery_method                  = "shared"
  handover_metro_area_network_name = "FRA"
  service_metro_area_network_name  = "FRA"
  service_provider_region          = "eu-central-1"
  service_provider_pop             = "EqFA5"
}
`

const cloudVCOfferingConfig = `
data "ixapi_product_offering_cloud_vc" "aws_euc1_50_2" {
  service_provider                 = "AWS"
  bandwidth                        = 50
  contract_period                  = "P1M"
  delivery_method                  = "shared"
  handover_metro_area_network_name = "FRA"
  service_metro_area_network_name  = "FRA"
  service_provider_region          = "eu-central-1"
  service_provider_pop             = "EqFA5"
}
`

func TestAccProductOfferingCloudVC(t *testing.T) {
	testhelpers.RequireTestEnv(t, "TF_VAR_API_URL")

	resource.Test(t, resource.TestCase{
		ProviderFactories: testhelpers.ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{}) + cloudVCOfferingsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ixapi_product_offerings_cloud_vc.filtered", "product_offerings.#", "1"),
					resource.TestCheckResourceAttrSet("data.ixapi_product_offerings_cloud_vc.filtered", "product_offerings.0.id"),
					resource.TestCheckResourceAttr("data.ixapi_product_offerings_cloud_vc.filtered", "product_offerings.0.contract_period", "P1M"),
					resource.TestCheckResourceAttr("data.ixapi_product_offerings_cloud_vc.filtered", "product_offerings.0.service_provider", "AWS"),
					resource.TestCheckResourceAttr("data.ixapi_product_offerings_cloud_vc.filtered", "product_offerings.0.delivery_method", "shared"),
				),
			},
			{
				Config: testhelpers.ProviderConfig(testhelpers.ProviderConfigOptions{}) + cloudVCOfferingConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ixapi_product_offering_cloud_vc.aws_euc1_50_2", "id"),
					resource.TestCheckResourceAttr("data.ixapi_product_offering_cloud_vc.aws_euc1_50_2", "contract_period", "P1M"),
					resource.TestCheckResourceAttr("data.ixapi_product_offering_cloud_vc.aws_euc1_50_2", "delivery_method", "shared"),
					resource.TestCheckResourceAttr("data.ixapi_product_offering_cloud_vc.aws_euc1_50_2", "service_provider_region", "eu-central-1"),
					resource.TestCheckResourceAttr("data.ixapi_product_offering_cloud_vc.aws_euc1_50_2", "service_provider_pop", "EqFA5"),
				),
			},
		},
	})
}
