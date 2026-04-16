data "ixapi_de_cix_cloud_router_prefix_list" "example" {
  name = "my-prefix-list"
}

output "prefix_list" {
  value = data.ixapi_de_cix_cloud_router_prefix_list.example
}
