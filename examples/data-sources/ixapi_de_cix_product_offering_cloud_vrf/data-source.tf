data "ixapi_de_cix_product_offering_cloud_vrf" "example" {
  name = "Cloud ROUTER Frankfurt"
}

output "product_offering" {
  value = data.ixapi_de_cix_product_offering_cloud_vrf.example
}
