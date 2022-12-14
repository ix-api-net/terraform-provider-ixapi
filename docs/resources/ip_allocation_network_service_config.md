---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_ip_allocation_network_service_config Resource - terraform-provider-ixapi"
subcategory: ""
description: |-
  Use this resource to reference IP addresses allocated by a referanced network service. The managed IP addresses will be updated with the fqdn provided.
---

# ixapi_ip_allocation_network_service_config (Resource)

Use this resource to reference IP addresses allocated by a referanced network service. The managed IP addresses will be updated with the fqdn provided.

## Example Usage

```terraform
# Get the ip allocation resource and set the fqdn of the IP address(es)
resource "ixapi_ip_allocation_network_service_config" "customernet_fra" {
    network_service_config = resource.ixapi_network_service_config_exchange_lan.customernet_fra.id
    fqdn = "gw1.fra.customer.example.net"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_service_config` (String) Get IP addresses allocated for this network service config

### Optional

- `fqdn` (String) Set the fqdn of the allocated IP addresses
- `ips` (Block List) (see [below for nested schema](#nestedblock--ips))
- `timeout` (Number) Timeout for awaiting the allocation Defaults to `900`.
- `version` (Number) Only get IP addresses with this version (4 or 6)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ips"></a>
### Nested Schema for `ips`

Optional:

- `address` (String) IPv4 or IPv6 Address in the following format: - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) - IPv6: hexadecimal colon separated notation
- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `fqdn` (String)
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.
- `prefix_length` (Number) The CIDR ip prefix length
- `valid_not_after` (String)
- `valid_not_before` (String)
- `version` (Number) The version of the internet protocol.

Read-Only:

- `id` (String) The ID of this resource.


