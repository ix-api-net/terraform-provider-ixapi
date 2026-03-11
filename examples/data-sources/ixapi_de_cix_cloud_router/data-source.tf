data "ixapi_de_cix_cloud_router" "example" {
  id = "vrf-123"
}

output "cloud_router" {
  value = data.ixapi_de_cix_cloud_router.example
}
