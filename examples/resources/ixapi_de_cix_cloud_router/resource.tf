
data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router" "example" {
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  product_offering  = "1"
  asn               = 65001
  capacity          = 1000
  external_ref      = "my-cloud-router"
}

output "cloud_router_id" {
  value = ixapi_de_cix_cloud_router.example.id
}

output "cloud_router_state" {
  value = ixapi_de_cix_cloud_router.example.state
}
