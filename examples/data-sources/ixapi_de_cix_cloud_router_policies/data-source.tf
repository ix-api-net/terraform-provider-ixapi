data "ixapi_de_cix_cloud_router_policies" "example" {}

output "policies" {
  value = data.ixapi_de_cix_cloud_router_policies.example.policies
}
