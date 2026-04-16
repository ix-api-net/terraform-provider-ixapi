data "ixapi_de_cix_cloud_router_prefix_lists" "example" {}

output "prefix_lists" {
  value = data.ixapi_de_cix_cloud_router_prefix_lists.example.prefix_lists
}
