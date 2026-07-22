terraform {
  required_providers {
    ixapi = {
      source = "ix-api-net/ixapi"
    }
  }
}

# Your side. The DE-CIX Cloud Router extension must be enabled.
provider "ixapi" {
  api                                   = "https://ixapi.example.com"
  api_key                               = "your-api-key"
  extension_de_cix_cloud_router_enabled = true
}

# The foreign partner. In practice each side is usually managed by its own
# Terraform configuration and account; both are shown here for completeness.
provider "ixapi" {
  alias                                 = "partner"
  api                                   = "https://ixapi.example.com"
  api_key                               = "partner-api-key"
  extension_de_cix_cloud_router_enabled = true
}

# ---------------------------------------------------------------------------
# Partner side: publish a connection as discoverable
# ---------------------------------------------------------------------------

data "ixapi_connection" "partner_existing" {
  provider = ixapi.partner
  id       = var.partner_connection_id
}

resource "ixapi_connection" "partner" {
  provider          = ixapi.partner
  managing_account  = data.ixapi_connection.partner_existing.managing_account
  consuming_account = data.ixapi_connection.partner_existing.consuming_account
  billing_account   = data.ixapi_connection.partner_existing.billing_account
  mode              = data.ixapi_connection.partner_existing.mode
  product_offering  = data.ixapi_connection.partner_existing.product_offering
  port_quantity     = data.ixapi_connection.partner_existing.port_quantity
  discoverable      = true
}

# ---------------------------------------------------------------------------
# Your side: find the offering the partner exposes and build the circuit
# ---------------------------------------------------------------------------

data "ixapi_account" "partner" {
  id = var.partner_account_id
}

data "ixapi_metro_area_network" "fra" {
  name = "FRA"
}

data "ixapi_product_offerings_p2p_vc" "from_partner" {
  service_provider            = data.ixapi_account.partner.name
  handover_metro_area_network = data.ixapi_metro_area_network.fra.id

  depends_on = [ixapi_connection.partner]
}

resource "ixapi_network_service_p2p_vc" "pni" {
  managing_account       = var.account_id
  consuming_account      = var.account_id
  billing_account        = var.account_id
  joining_member_account = data.ixapi_account.partner.id
  product_offering       = data.ixapi_product_offerings_p2p_vc.from_partner.product_offerings[0].id
  display_name           = "virtual-pni-foreign"
}

resource "ixapi_network_service_config_p2p_vc" "mine" {
  managing_account   = var.account_id
  consuming_account  = var.account_id
  billing_account    = var.account_id
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = var.my_connection_id
  role_assignments   = []

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 100
  }
}

# ---------------------------------------------------------------------------
# Partner side: attach their port to the same network service
# ---------------------------------------------------------------------------

resource "ixapi_network_service_config_p2p_vc" "partner" {
  provider           = ixapi.partner
  managing_account   = var.partner_account_id
  consuming_account  = var.partner_account_id
  billing_account    = var.partner_account_id
  network_service    = ixapi_network_service_p2p_vc.pni.id
  network_connection = ixapi_connection.partner.id
  role_assignments   = []

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 100
  }
}

variable "account_id" {}
variable "my_connection_id" {}
variable "partner_account_id" {}
variable "partner_connection_id" {}
