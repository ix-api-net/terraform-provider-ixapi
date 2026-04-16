data "ixapi_de_cix_cloud_router_network_service_config_received_routes" "example" {
  network_service_config_id = "nsc-123"
}

output "received_routes" {
  value = data.ixapi_de_cix_cloud_router_network_service_config_received_routes.example.routes
}
