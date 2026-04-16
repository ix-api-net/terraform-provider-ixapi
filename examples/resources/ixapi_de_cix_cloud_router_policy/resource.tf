data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router_prefix_list" "allowed" {
  name              = "allowed-prefixes"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  match_list {
    prefix     = "10.0.0.0/8"
    min_length = 8
    max_length = 24
  }
}

resource "ixapi_de_cix_cloud_router_policy" "example" {
  name              = "my-ingress-policy"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.allowed.name
    action {
      filter           = "accept"
      local_preference = 200
    }
  }
  entries {
    sequence_number = 20
    action {
      filter = "reject"
    }
  }
}

output "policy_id" {
  value = ixapi_de_cix_cloud_router_policy.example.id
}
