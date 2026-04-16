data "ixapi_de_cix_cloud_router_bgp_state" "example" {
  nsc_id = "nsc-123"
}

output "bgp_state" {
  value = data.ixapi_de_cix_cloud_router_bgp_state.example.state
}
