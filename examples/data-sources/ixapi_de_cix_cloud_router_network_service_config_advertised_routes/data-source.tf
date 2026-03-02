data "ixapi_de_cix_cloud_router_network_service_config_advertised_routes" "example" {
  network_service_config_id = "123"
}

output "advertised_routes" {
  value = data.ixapi_de_cix_cloud_router_network_service_config_advertised_routes.example.routes
}

output "first_route_prefix" {
  value = length(data.ixapi_de_cix_cloud_router_network_service_config_advertised_routes.example.routes) > 0 ? data.ixapi_de_cix_cloud_router_network_service_config_advertised_routes.example.routes[0].prefix : null
}
