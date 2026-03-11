data "ixapi_de_cix_cloud_router_static_route" "example" {
  id = "route-123"
}

output "static_route" {
  value = data.ixapi_de_cix_cloud_router_static_route.example
}
