data "ixapi_de_cix_product_offerings_cloud_vrf" "example" {}

output "product_offerings" {
  value = data.ixapi_de_cix_product_offerings_cloud_vrf.example.product_offerings
}
