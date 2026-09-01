# Cloud ROUTER Complete Example

This example demonstrates a complete DE-CIX Cloud ROUTER setup with multiple cloud and network connections.

## What This Example Creates

1. **Cloud ROUTER Instance** - A VRF with ASN 65001 and 2000 Mbps capacity
2. **AWS DirectCloud Connection** - BGP peering to AWS with routing policies
3. **Azure DirectCloud Connection** - BGP peering to Azure
4. **Partner Virtual PNI Connection** - BGP peering to a business partner network

## Architecture

```
                    ┌─────────────────────┐
                    │   Cloud ROUTER      │
                    │   ASN: 65001        │
                    │   Capacity: 2000Mbps│
                    └──────────┬──────────┘
                               │
              ┌────────────────┼────────────────┐
              │                │                │
              ▼                ▼                ▼
     ┌────────────────┐ ┌────────────┐ ┌────────────────┐
     │ AWS DirectCloud│ │   Azure    │ │ Partner Network│
     │  ASN: 64512    │ │ASN: 64513  │ │  ASN: 65100    │
     │  VLAN: 100     │ │ VLAN: 200  │ │  VLAN: 300     │
     │  w/ Policies   │ │            │ │                │
     └────────────────┘ └────────────┘ └────────────────┘
```

## Features Demonstrated

- Creating a Cloud ROUTER with specific ASN and capacity
- Configuring multiple DirectCloud connections (AWS, Azure)
- Configuring Virtual PNI to partner networks
- Using BGP authentication (MD5 passwords)
- Applying routing policies (ingress/egress)
- Enabling BFD for fast failure detection
- VLAN configuration for each connection
- Using data sources to query configurations
- Managing sensitive values with variables

## Prerequisites

- IX-API credentials configured
- Access to a metro area network supporting Cloud ROUTER
- Network service IDs for your target clouds/networks
- BGP passwords for secure peering

## Usage

1. Copy this directory to your workspace
2. Create a `terraform.tfvars` file:

```hcl
aws_bgp_password   = "your-aws-password"
azure_bgp_password = "your-azure-password"
```

3. Initialize and apply:

```bash
terraform init
terraform plan
terraform apply
```

## Bandwidth Upgrades

To upgrade the Cloud ROUTER capacity, update the `capacity` value:

```hcl
resource "ixapi_de_cix_cloud_router" "main" {
  # ... other attributes ...
  capacity = 5000  # Upgraded from 2000
}
```

Then apply the change:

```bash
terraform apply
```

## Notes

- The `asn` and `product_offering` fields cannot be changed after creation
- Each configuration requires a unique VLAN ID
- Routing policies must be pre-configured in the IX-API system
- BFD is recommended for production environments for faster failure detection

### AS Override for OCI Multi-Region Connections

Oracle Cloud Infrastructure (OCI) uses the same ASN, `31898`, across its regions. When connecting multiple OCI regions to the same Cloud ROUTER, enable `as_override` on every affected connection so OCI does not reject routes from another region as an AS loop:

```hcl
resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "oci_frankfurt" {
  # ...
  bgp_neighbor_asn = 31898
  as_override      = true
}

resource "ixapi_de_cix_cloud_router_network_service_config_cloud_vc" "oci_amsterdam" {
  # ...
  bgp_neighbor_asn = 31898
  as_override      = true
}
```
