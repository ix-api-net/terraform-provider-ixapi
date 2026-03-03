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

data "ixapi_account" "customer" {
  external_ref = "my_account"
}

resource "ixapi_de_cix_cloud_router" "main" {
  managing_account  = data.ixapi_account.customer.id
  consuming_account = data.ixapi_account.customer.id
  billing_account   = data.ixapi_account.customer.id
  product_offering  = "1"
  asn               = 65001
  capacity          = 2000
  external_ref      = "production-cloud-router"
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "aws_directcloud" {
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
  policy_ingress    = "aws-ingress-policy"
  policy_egress     = "aws-egress-policy"

  vlan_config {
    vlan_type = "dot1q"
    vlan      = 100
  }
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "azure_directcloud" {
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

resource "ixapi_de_cix_cloud_router_network_service_config_p2p_vc" "partner_vpni" {
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

data "ixapi_de_cix_cloud_router_network_service_configs_cloud_vc" "cloud_configs" {
  depends_on = [
    ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws_directcloud,
    ixapi_de_cix_cloud_router_network_service_config_cloud_vc.azure_directcloud,
  ]
}

data "ixapi_de_cix_cloud_router_network_service_configs_p2p_vc" "p2p_configs" {
  depends_on = [
    ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_vpni
  ]
}

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

output "cloud_router_id" {
  description = "The ID of the Cloud ROUTER"
  value       = ixapi_de_cix_cloud_router.main.id
}

output "cloud_router_asn" {
  description = "The ASN of the Cloud ROUTER"
  value       = ixapi_de_cix_cloud_router.main.asn
}

output "cloud_router_state" {
  description = "The state of the Cloud ROUTER"
  value       = ixapi_de_cix_cloud_router.main.state
}

output "configuration_count" {
  description = "Total number of Cloud ROUTER configurations"
  value       = length(data.ixapi_de_cix_cloud_router_network_service_configs_cloud_vc.cloud_configs.cloud_router_network_service_configs) + length(data.ixapi_de_cix_cloud_router_network_service_configs_p2p_vc.p2p_configs.cloud_router_network_service_configs)
}

output "aws_config_id" {
  description = "AWS DirectCloud configuration ID"
  value       = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.aws_directcloud.id
}

output "azure_config_id" {
  description = "Azure DirectCloud configuration ID"
  value       = ixapi_de_cix_cloud_router_network_service_config_cloud_vc.azure_directcloud.id
}

output "partner_config_id" {
  description = "Partner Virtual PNI configuration ID"
  value       = ixapi_de_cix_cloud_router_network_service_config_p2p_vc.partner_vpni.id
}
