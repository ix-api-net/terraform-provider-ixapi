data "ixapi_de_cix_cloud_router_static_routes" "example" {
  vrf = "vrf-123"
}

output "static_routes" {
  value = data.ixapi_de_cix_cloud_router_static_routes.example.static_routes
}
