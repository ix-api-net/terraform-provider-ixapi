
data "ixapi_de_cix_cloud_router" "example" {
  external_ref = "my-cloud-router"
}

data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "aws" {
  managing_account  = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  cloud_router      = data.ixapi_de_cix_cloud_router.example.id
  network_service   = "500"
  address           = "192.0.2.1/30"
  bgp_neighbor      = "192.0.2.2"
  bgp_neighbor_asn  = 64512
  bgp_password      = "my-secret-password"
  admin_status      = "enabled"
  bfd_enabled       = true

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 100
  }
}

output "config_id" {
  value = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id
}

output "config_state" {
  value = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.state
}
