data "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "example" {
  id = "123"
}

output "config_state" {
  value = data.ixapi_de_cix_cloud_router_network_service_config_cloud_vc.example.state
}
