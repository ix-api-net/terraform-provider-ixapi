data "ixapi_de_cix_cloud_router_network_service_config_advertised_routes" "example" {
  network_service_config_id = "nsc-123"
}

output "advertised_routes" {
  value = data.ixapi_de_cix_cloud_router_network_service_config_advertised_routes.example.routes
}
