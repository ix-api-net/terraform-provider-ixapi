
**This repository is work in progress and not an official release.**

# IX-API Terraform Provider

This Terraform provider is using the [ix-api](https://ix-api.net)
for configuring and provisisioning IXP services.


## Requirements
 * Terraform >= 1.0
 * Go >= 1.17


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


## DE-CIX Cloud ROUTER Extension

The provider includes a DE-CIX-specific extension for managing Cloud ROUTER (VRF) resources.
This extension is disabled by default and must be explicitly enabled in the provider configuration:

```hcl
provider "ixapi" {
  api        = "https://ixapi.example.com/api/v2"
  api_key    = var.api_key
  api_secret = var.api_secret
  de_cix_cloud_router_extension_enabled = true
}
```

When `de_cix_cloud_router_extension_enabled` is `false` (the default), any attempt to use a
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
