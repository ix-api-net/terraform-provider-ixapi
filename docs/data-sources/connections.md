---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_connections Data Source - terraform-provider-ixapi"
subcategory: ""
description: |-
  Use the connections data source find available connections
---

# ixapi_connections (Data Source)

Use the `connections` data source find available connections

## Example Usage

```terraform
# Get metro area by iata code
data "ixapi_metro_area" "fra" {
  iata_code = "FRA"  # Resolve metro area by IATA code
}

# Get metro area networks
data "ixapi_metro_area_networks" "fra" {
  metro_area = data.ixapi_metro_area.fra.id
}

# Get connections available at given metro area network
data "ixapi_connections" "fra" {
  consuming_account = "<my account id>"
  metro_area_network = data.ixapi_metro_area_networks.fra.metro_area_networks[0]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connections` (Block List) (see [below for nested schema](#nestedblock--connections))
- `consuming_account` (String) Filter by account using the connection, e.g. the customer
- `managing_account` (String) Filter by account managing the connection
- `metro_area_network` (String) Filter by metro area network ID, see metro area network data source
- `name` (String) Filter by connection name
- `pop` (String) Filter by PoP ID, see point of presence data source

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--connections"></a>
### Nested Schema for `connections`

Optional:

- `billing_account` (String) An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*
- `capacity_allocated` (Number) Sum of the bandwidth of all network services using the connection in Mbit/s.
- `capacity_allocation_limit` (Number) Maximum allocatable capacity of the connection in Mbit/s. When `null`, the exchange does not impose any limit.
- `charged_until` (String) The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*
- `connecting_party` (String) Name of the service provider who establishes connectivity on your behalf.  This is only relevant, if the cross connect initiator is the `subscriber` and might be `null`.  Please refer to the usage guide of the internet exchange.
- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `contract_ref` (String) A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)*
- `current_billing_start_date` (String) Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*
- `decommission_at` (String) The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.
- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `lacp_timeout` (String) This sets the LACP Timeout mode. Both ends of the connections need to be configured the same.
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.
- `mode` (String) Sets the mode of the connection. The mode can be:  - `lag_lacp`: connection is build as a LAG with LACP enabled - `lag_static`: connection is build as LAG with static configuration - `flex_ethernet`: connect is build as a FlexEthernet channel - `standalone`: only one port is allowed in this connection without any bundling.
- `name` (String)
- `outer_vlan_ethertypes` (List of String)
- `pop` (String) The ID of the point of presence (see `/pops`), where the physical port(s) are present.
- `port_quantity` (Number) The number of ports which should be allocated for this connection.
- `port_reservations` (List of String)
- `ports` (List of String)
- `product_offering` (String) The product offering must match the type `connection`.
- `purchase_order` (String) Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)*
- `role_assignments` (List of String)
- `speed` (Number) Shows the total bandwidth of the connection in Mbit/s.
- `state` (String)
- `status` (Block List) (see [below for nested schema](#nestedblock--connections--status))
- `subscriber_side_demarcs` (List of String)
- `vlan_types` (List of String)

Read-Only:

- `id` (String) The ID of this resource.

<a id="nestedblock--connections--status"></a>
### Nested Schema for `connections.status`

Optional:

- `message` (String) A human readable message, describing the problem and may contain hints for resolution.
- `severity` (Number) We are using syslog severity levels: 0 = Emergency, 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning, 5 = Notice, 6 = Informational, 7 = Debug.
- `tag` (String) A machine readable message identifier.
- `timestamp` (String) The time and date when the event occured.


