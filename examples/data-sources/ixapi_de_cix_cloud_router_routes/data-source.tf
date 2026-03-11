data "ixapi_de_cix_cloud_router_routes" "example" {
  vrf = "vrf-123"
}

output "routing_table" {
  value = data.ixapi_de_cix_cloud_router_routes.example.routes
}
