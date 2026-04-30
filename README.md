

# IX-API Terraform Provider

This Terraform provider is using the [ix-api](https://ix-api.net)
for configuring and provisisioning IXP services.


## Requirements
 * Terraform >= 1.0
 * Go >= 1.26


## Configure the provider

Simple provider configuration example using
the `legacy` authentication strategy.

```hcl
provider "ixapi" {
    api = "http://localhost:8000/api/v2"
    api_key = "my_api_key"
    # api_secret = "" # Use $IX_API_SECRET environment variable
}
```

### Environment Variables

You can also use the environment variables:

 * `$IX_API_AUTH`: Choose the authentication strategy.
   Defaults to `legacy`. Can be set to `oauth2`.
 * `$IX_API_HOST`: The IX-API endpoint in the format: `https://<server>/api/v2`
 * `$IX_API_KEY`: The key provided by the exchange.
 * `$IX_API_SECRET`: Also provided by the exchange.
 * `$IX_API_OAUTH2_TOKEN_URL`: The OAuth2 token endpoint.
 * `$IX_API_OAUTH2_SCOPES`: A comma-separated list of OAuth2 scopes. (Optional)

### OAuth2

In order to use OAuth2 to retrieve an access token, you
need to provide the `oauth2_token_url` in addition to
the `api_key` and `api_secret`. Key and secret will be used
as `client_id` and `client_secret`.

The `auth` strategy must be set to `oauth2`.

```hcl
provider "ixapi" {
    auth = "oauth2"
    api = "http://localhost:8000/api/v2"
    api_key = "my_api_key"
    api_secret = "..."
    oauth2_token_url = "http://localhost:8000/auth/oauth2/token"
    oauth2_scopes = "ix-api"  # Optional
}
```


## Using The Provider

The following examples illustrate basic usage.

```hcl
# Querying: Show all facilities in the metro area FRA
data "ixapi_metro_area" "fra" {
  iata_code = "FRA"  # Resolve metro area by IATA code
}

data "ixapi_facilities" "fra" {
  metro_area = data.ixapi_metro_area.fra.id
}

output "facilities" {
  value = data.ixapi_facilities.fra.facilities
}

data "ixapi_account" "reseller" {
  external_ref = "demo_reseller"
}
```

### Using Resources

Create an account and add a contact.

```hcl
resource "ixapi_account" "hajnet" {
  managing_account = data.ixapi_account.reseller.id
  name = "hajnet"
  legal_name = "Blåhaj Networks Inc."
  address {
    country = "DE"
    locality = "Berlin"
    postal_code = "11111"
    street_address = "Straßenweg 11"
  }
}

locals {
  reseller_id = data.ixapi_account.reseller.id
  customer_id = resource.ixapi_account.hajnet.id
}

resource "ixapi_contact" "hajnet_support" {
  managing_account = local.reseller_id
  consuming_account = local.hajnet_id
  roles = ["noc", "implementation" ]
  email = "mail@example.com"
  telephone = "+0 42 1234567890"
}
```


## Running Acceptance Tests

Acceptance tests run against a real IX-API backend.

```bash
TF_VAR_API_URL=https://ixapi.example.com/api/v2 \
TF_VAR_API_KEY=your-api-key \
TF_VAR_API_SECRET=your-api-secret \
ACCOUNT_ID=your-account-id \
make acceptance
```

To target a specific file:

```bash
TF_ACC=1 go test ./acceptance/decix/vrf/ -v -count=1
```

## Development

For development, you need to add the development build
of the terraform provider. You can do so, by adding
the following snippet to you `~/.terraformrc`:

```hcl
provider_installation {
    dev_overrides {
        "ix-api.net/ix-api/ixapi" = "/<full_path_to>/go/src/github.com/ix-api-net/terraform-provider-ixapi/bin"
    }

    direct {}
}
```

And then in the terraform file use:

```hcl
terraform {
    required_providers {
        ixapi = {
            source = "ix-api.net/ix-api/ixapi"
        }
    }
}
```

Use a `ix-api-sandbox-v2` as local API server.


## IX-API Resources and Data Sources

### Available Resources

| Resource | Description |
|---|---|
| `ixapi_account` | Manage an account |
| `ixapi_contact` | Manage a contact with role assignments |
| `ixapi_connection` | Create a connection and allocate ports |
| `ixapi_port_reservation` | Manage port reservations for a connection |
| `ixapi_mac` | Register a MAC address |
| `ixapi_network_service_p2p_vc` | Point-to-point virtual circuit |
| `ixapi_network_service_p2mp_vc` | Point-to-multipoint virtual circuit |
| `ixapi_network_service_mp2mp_vc` | Multipoint-to-multipoint virtual circuit |
| `ixapi_network_service_cloud_vc` | Cloud virtual circuit |
| `ixapi_network_service_config_exchange_lan` | Exchange LAN access configuration |
| `ixapi_network_service_config_p2p_vc` | P2P virtual circuit access configuration |
| `ixapi_network_service_config_p2mp_vc` | P2MP virtual circuit access configuration |
| `ixapi_network_service_config_mp2mp_vc` | MP2MP virtual circuit access configuration |
| `ixapi_network_service_config_cloud_vc` | Cloud virtual circuit access configuration |
| `ixapi_network_feature_config_route_server` | Route server configuration |
| `ixapi_ip_allocation_network_service_config` | IP address allocation for a network service config |
| `ixapi_member_joining_rule_allow` | Allow a consuming account to join a network service |
| `ixapi_member_joining_rule_deny` | Deny a consuming account from joining a network service |

### Available Data Sources

**Account & Contact:**

| Data Source | Description |
|---|---|
| `ixapi_accounts` / `_account` | Query accounts |
| `ixapi_contacts` / `_contact` | Query contacts |

**Connections & Ports:**

| Data Source | Description |
|---|---|
| `ixapi_connections` | Query connections |
| `ixapi_ports` / `_port` | Query ports |
| `ixapi_port_reservations` / `_port_reservation` | Query port reservations |
| `ixapi_macs` / `_mac` | Query MAC addresses |

**Devices & Facilities:**

| Data Source | Description |
|---|---|
| `ixapi_devices` / `_device` | Query devices |
| `ixapi_facilities` / `_facility` | Query facilities |
| `ixapi_pops` / `_pop` | Query points of presence |
| `ixapi_metro_areas` / `_metro_area` | Query metro areas |
| `ixapi_metro_area_networks` / `_metro_area_network` | Query metro area networks |

**IP Addresses & Roles:**

| Data Source | Description |
|---|---|
| `ixapi_ips` / `_ip` | Query IP addresses |
| `ixapi_roles` / `_role` | Query roles |
| `ixapi_role_assignments` / `_role_assignment` | Query role assignments |
| `ixapi_member_joining_rules` / `_member_joining_rule` | Query member joining rules |

**Network Services:**

| Data Source | Description |
|---|---|
| `ixapi_network_services_exchange_lan` / `_exchange_lan` | Query exchange LAN services |
| `ixapi_network_services_p2p_vc` / `_p2p_vc` | Query P2P virtual circuit services |
| `ixapi_network_services_p2mp_vc` / `_p2mp_vc` | Query P2MP virtual circuit services |
| `ixapi_network_services_mp2mp_vc` / `_mp2mp_vc` | Query MP2MP virtual circuit services |
| `ixapi_network_services_cloud_vc` / `_cloud_vc` | Query cloud virtual circuit services |

**Network Features:**

| Data Source | Description |
|---|---|
| `ixapi_network_features_route_server` / `_route_server` | Query route server features |

**Product Offerings:**

| Data Source | Description |
|---|---|
| `ixapi_product_offerings_connection` / `_offering_connection` | Connection product offerings |
| `ixapi_product_offerings_exchange_lan` / `_offering_exchange_lan` | Exchange LAN product offerings |
| `ixapi_product_offerings_p2p_vc` / `_offering_p2p_vc` | P2P virtual circuit product offerings |
| `ixapi_product_offerings_p2mp_vc` / `_offering_p2mp_vc` | P2MP virtual circuit product offerings |
| `ixapi_product_offerings_mp2mp_vc` / `_offering_mp2mp_vc` | MP2MP virtual circuit product offerings |
| `ixapi_product_offerings_cloud_vc` / `_offering_cloud_vc` | Cloud virtual circuit product offerings |


## DE-CIX Cloud ROUTER Extension

The provider includes a DE-CIX-specific extension for managing Cloud ROUTER (VRF) resources.
This extension is disabled by default and must be explicitly enabled in the provider configuration:

```hcl
provider "ixapi" {
  api        = "https://ixapi.example.com/api/v2"
  api_key    = var.api_key
  api_secret = var.api_secret
  extension_de_cix_cloud_router_enabled = true
}
```

When `extension_de_cix_cloud_router_enabled` is `false` (the default), any attempt to use a
`de_cix_*` resource or data source will return an error. This prevents accidental use against
an IX-API endpoint that does not support the extension.

### Available Resources

| Resource | Description |
|---|---|
| `ixapi_de_cix_cloud_router` | Cloud ROUTER (VRF) instance |
| `ixapi_de_cix_cloud_router_network_service_config_cloud_vc` | BGP peering to a cloud provider (AWS, Azure, …) |
| `ixapi_de_cix_cloud_router_network_service_config_p2p_vc` | BGP peering over a point-to-point virtual circuit |
| `ixapi_de_cix_cloud_router_prefix_list` | IP prefix list for use in routing policies |
| `ixapi_de_cix_cloud_router_policy` | Routing policy (ingress/egress) with prefix-list matching |
| `ixapi_de_cix_cloud_router_static_route` | Static route attached to one or more network service configs |

### Available Data Sources

| Data Source | Description |
|---|---|
| `ixapi_de_cix_cloud_routers` / `_cloud_router` | Query Cloud ROUTER instances |
| `ixapi_de_cix_cloud_router_network_service_configs_cloud_vc` / `_config_cloud_vc` | Query cloud VC configs |
| `ixapi_de_cix_cloud_router_network_service_configs_p2p_vc` / `_config_p2p_vc` | Query P2P VC configs |
| `ixapi_de_cix_cloud_router_prefix_lists` / `_prefix_list` | Query prefix lists |
| `ixapi_de_cix_cloud_router_policies` / `_policy` | Query routing policies |
| `ixapi_de_cix_cloud_router_static_routes` / `_static_route` | Query static routes |
| `ixapi_de_cix_cloud_router_bgp_state` | BGP session state for a network service config |
| `ixapi_de_cix_cloud_router_bfd_state` | BFD session state for a network service config |
| `ixapi_de_cix_cloud_router_network_service_config_advertised_routes` | BGP routes advertised to a peer |
| `ixapi_de_cix_cloud_router_network_service_config_received_routes` | BGP routes received from a peer |
| `ixapi_de_cix_cloud_router_arp_table` | ARP table for a VRF |
| `ixapi_de_cix_cloud_router_routes` | Full routing table for a VRF |
| `ixapi_de_cix_product_offerings_cloud_vrf` / `_offering_cloud_vrf` | Available Cloud ROUTER product offerings |

See [`examples/cloud_router_complete`](examples/cloud_router_complete) for a full example
covering prefix lists, routing policies, static routes, and operational data sources.
