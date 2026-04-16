data "ixapi_de_cix_cloud_router_policy" "example" {
  name = "my-routing-policy"
}

output "policy" {
  value = data.ixapi_de_cix_cloud_router_policy.example
}
