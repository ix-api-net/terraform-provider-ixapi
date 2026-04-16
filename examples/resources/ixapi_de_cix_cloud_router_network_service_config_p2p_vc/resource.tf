
data "ixapi_de_cix_cloud_router" "example" {
  external_ref = "my-cloud-router"
}

data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router_network_service_config_p2p_vc" "partner" {
  managing_account   = data.ixapi_account.customer.id
  billing_account    = data.ixapi_account.customer.id
  consuming_account  = data.ixapi_account.customer.id
  cloud_router       = data.ixapi_de_cix_cloud_router.example.id
  network_service    = "502"
  network_connection = "1"
  address            = "192.0.2.9/30"
  bgp_neighbor       = "192.0.2.10"
  bgp_neighbor_asn   = 65100
  admin_status       = "enabled"
  bfd_enabled        = false

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 300
  }
}

output "config_id" {
  value = ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner.id
}

output "config_state" {
  value = ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner.state
}
