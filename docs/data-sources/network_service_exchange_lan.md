---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_network_service_exchange_lan Data Source - terraform-provider-ixapi"
subcategory: ""
description: |-
  Get an exchange lan network service by ID
---

# ixapi_network_service_exchange_lan (Data Source)

Get an exchange lan network service by ID



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `ixfdb_ixid` (Number) id of ixfdb
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.
- `metro_area_network` (String) Id of the `MetroAreaNetwork` where the exchange lan network service is directly provided.  Same as `service_metro_area_network` on the related `ProductOffering`.
- `name` (String) Exchange-dependent service name, will be shown on the invoice.
- `network_features` (List of String)
- `nsc_required_contact_roles` (List of String)
- `peeringdb_ixid` (Number) PeeringDB ixid
- `product_offering` (String) *deprecation notice*
- `state` (String)
- `status` (Block List) (see [below for nested schema](#nestedblock--status))
- `subnet_v4` (String) IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) CIDR notation.
- `subnet_v6` (String) IPv6 subnet in hexadecimal colon separated CIDR notation.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--status"></a>
### Nested Schema for `status`

Optional:

- `message` (String) A human readable message, describing the problem and may contain hints for resolution.
- `severity` (Number) We are using syslog severity levels: 0 = Emergency, 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning, 5 = Notice, 6 = Informational, 7 = Debug.
- `tag` (String) A machine readable message identifier.
- `timestamp` (String) The time and date when the event occured.


