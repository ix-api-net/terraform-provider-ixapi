terraform {
  required_providers {
    ixapi = {
      source = "ix-api-net/ixapi"
    }
  }
}

provider "ixapi" {
  api                                   = "https://ixapi.example.com"
  de_cix_cloud_router_extension_enabled = true
}

# ---------------------------------------------------------------------------
# Account
# ---------------------------------------------------------------------------

data "ixapi_account" "customer" {
  external_ref = "my_account"
}

# ---------------------------------------------------------------------------
# Product offering discovery
# ---------------------------------------------------------------------------

data "ixapi_de_cix_product_offering_cloud_vrf" "vrf" {
  name = "Cloud ROUTER Frankfurt"
}

# ---------------------------------------------------------------------------
# Prefix lists
# ---------------------------------------------------------------------------

resource "ixapi_de_cix_cloud_router_prefix_list" "allowed_from_aws" {
  name              = "allowed-from-aws"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  match_list {
    prefix     = "10.0.0.0/8"
    min_length = 8
    max_length = 24
  }
  match_list {
    prefix     = "172.16.0.0/12"
    min_length = 12
    max_length = 28
  }
}

resource "ixapi_de_cix_cloud_router_prefix_list" "customer_prefixes" {
  name              = "customer-prefixes"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  match_list {
    prefix = "192.0.2.0/24"
  }
}

# ---------------------------------------------------------------------------
# Routing policies
# ---------------------------------------------------------------------------

resource "ixapi_de_cix_cloud_router_policy" "aws_ingress" {
  name              = "aws-ingress-policy"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.allowed_from_aws.name
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

resource "ixapi_de_cix_cloud_router_policy" "aws_egress" {
  name              = "aws-egress-policy"
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id

  entries {
    sequence_number   = 10
    match_prefix_list = ixapi_de_cix_cloud_router_prefix_list.customer_prefixes.name
    action {
      filter = "accept"
      as_path_prepend {
        count = 2
      }
    }
  }
  entries {
    sequence_number = 20
    action {
      filter = "reject"
    }
  }
}

# ---------------------------------------------------------------------------
# Cloud ROUTER (VRF)
# ---------------------------------------------------------------------------

resource "ixapi_de_cix_cloud_router" "main" {
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  product_offering  = data.ixapi_de_cix_product_offering_cloud_vrf.vrf.id
  asn               = 65001
  capacity          = 2000
  external_ref      = "production-cloud-router"
}

# ---------------------------------------------------------------------------
# Network service configs
# ---------------------------------------------------------------------------

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "aws" {
  managing_account  = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  cloud_router      = ixapi_de_cix_cloud_router.main.id
  network_service   = "500"
  address           = "192.0.2.1/30"
  bgp_neighbor      = "192.0.2.2"
  bgp_neighbor_asn  = 64512
  bgp_password      = var.aws_bgp_password
  admin_status      = "enabled"
  bfd_enabled       = true
  policy_ingress    = ixapi_de_cix_cloud_router_policy.aws_ingress.name
  policy_egress     = ixapi_de_cix_cloud_router_policy.aws_egress.name

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 100
  }
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "azure" {
  managing_account  = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  cloud_router      = ixapi_de_cix_cloud_router.main.id
  network_service   = "501"
  address           = "192.0.2.5/30"
  bgp_neighbor      = "192.0.2.6"
  bgp_neighbor_asn  = 64513
  bgp_password      = var.azure_bgp_password
  admin_status      = "enabled"
  bfd_enabled       = true

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 200
  }
}

resource "ixapi_de_cix_cloud_router_network_service_config_p2p_vc" "partner" {
  managing_account  = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  cloud_router      = ixapi_de_cix_cloud_router.main.id
  network_service   = "502"
  nic               = "1"
  address           = "192.0.2.9/30"
  bgp_neighbor      = "192.0.2.10"
  bgp_neighbor_asn  = 65100
  admin_status      = "enabled"
  bfd_enabled       = false

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 300
  }
}

# ---------------------------------------------------------------------------
# Static routes
# ---------------------------------------------------------------------------

resource "ixapi_de_cix_cloud_router_static_route" "aggregate" {
  name   = "default-aggregate"
  prefix = "0.0.0.0/0"
  next_hop = "aggregate"
  network_service_configs = [
    ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id,
    ixapi_de_cix_cloud_router_network_service_config_cloud_vc.azure.id,
  ]
}

# ---------------------------------------------------------------------------
# Operational data sources
# ---------------------------------------------------------------------------

data "ixapi_de_cix_cloud_router_bgp_state" "aws" {
  nsc_id = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id
}

data "ixapi_de_cix_cloud_router_bfd_state" "aws" {
  nsc_id = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id
}

data "ixapi_de_cix_cloud_router_network_service_config_advertised_routes" "aws" {
  network_service_config_id = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id
}

data "ixapi_de_cix_cloud_router_network_service_config_received_routes" "aws" {
  network_service_config_id = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws.id
}

data "ixapi_de_cix_cloud_router_static_routes" "vrf" {
  vrf = ixapi_de_cix_cloud_router.main.id
}

data "ixapi_de_cix_cloud_router_arp_table" "vrf" {
  vrf = ixapi_de_cix_cloud_router.main.id
}

data "ixapi_de_cix_cloud_router_routes" "vrf" {
  vrf = ixapi_de_cix_cloud_router.main.id
}

# ---------------------------------------------------------------------------
# Variables
# ---------------------------------------------------------------------------

variable "aws_bgp_password" {
  description = "BGP password for AWS DirectCloud connection"
  type        = string
  sensitive   = true
}

variable "azure_bgp_password" {
  description = "BGP password for Azure DirectCloud connection"
  type        = string
  sensitive   = true
}

# ---------------------------------------------------------------------------
# Outputs
# ---------------------------------------------------------------------------

output "cloud_router" {
  description = "Cloud ROUTER resource"
  value       = ixapi_de_cix_cloud_router.main
}

output "aws_bgp_state" {
  description = "BGP session state for the AWS connection"
  value       = data.ixapi_de_cix_cloud_router_bgp_state.aws.state
}

output "aws_bfd_state" {
  description = "BFD session state for the AWS connection"
  value       = data.ixapi_de_cix_cloud_router_bfd_state.aws
}

output "aws_advertised_routes" {
  description = "Routes advertised to AWS"
  value       = data.ixapi_de_cix_cloud_router_network_service_config_advertised_routes.aws
}

output "aws_received_routes" {
  description = "Routes received from AWS"
  value       = data.ixapi_de_cix_cloud_router_network_service_config_received_routes.aws
}

output "vrf_static_routes" {
  description = "All static routes in the VRF"
  value       = data.ixapi_de_cix_cloud_router_static_routes.vrf
}

output "vrf_arp_table" {
  description = "ARP table of the VRF"
  value       = data.ixapi_de_cix_cloud_router_arp_table.vrf
}

output "vrf_routing_table" {
  description = "Full routing table of the VRF"
  value       = data.ixapi_de_cix_cloud_router_routes.vrf
}
