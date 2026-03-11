resource "ixapi_de_cix_cloud_router_static_route" "example" {
  name     = "my-aggregate"
  prefix   = "10.0.0.0/8"
  next_hop = "aggregate"
  network_service_configs = ["nsc-123"]
}

output "static_route_id" {
  value = ixapi_de_cix_cloud_router_static_route.example.id
}

output "static_route_vrf" {
  value = ixapi_de_cix_cloud_router_static_route.example.vrf
}
