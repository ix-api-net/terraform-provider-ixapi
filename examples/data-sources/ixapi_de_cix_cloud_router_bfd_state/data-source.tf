data "ixapi_de_cix_cloud_router_bfd_state" "example" {
  nsc_id = "nsc-123"
}

output "bfd_state" {
  value = data.ixapi_de_cix_cloud_router_bfd_state.example.state
}
