data "ixapi_de_cix_cloud_router_network_service_configs_p2p_vc" "example" {
  bfd_enabled = true
}

output "configs" {
  value = data.ixapi_de_cix_cloud_router_network_service_configs_p2p_vc.example.cloud_router_network_service_configs
}
