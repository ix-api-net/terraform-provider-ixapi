data "ixapi_de_cix_cloud_router_network_service_config_p2p_vc" "example" {
  id = "123"
}

output "config_state" {
  value = data.ixapi_de_cix_cloud_router_network_service_config_p2p_vc.example.state
}
