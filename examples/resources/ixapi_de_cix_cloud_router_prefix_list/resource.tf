data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router_prefix_list" "example" {
  name              = "my-prefix-list"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  match_list {
    prefix     = "10.0.0.0/8"
    min_length = 8
    max_length = 24
  }
  match_list {
    prefix     = "192.168.0.0/16"
    min_length = 16
    max_length = 28
  }
}

output "prefix_list_id" {
  value = ixapi_de_cix_cloud_router_prefix_list.example.id
}
