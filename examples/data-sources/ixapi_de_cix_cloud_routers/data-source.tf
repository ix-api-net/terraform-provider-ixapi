data "ixapi_de_cix_cloud_routers" "example" {}

output "cloud_routers" {
  value = data.ixapi_de_cix_cloud_routers.example.cloud_routers
}
