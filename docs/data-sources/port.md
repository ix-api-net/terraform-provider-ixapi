---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_port Data Source - terraform-provider-ixapi"
subcategory: ""
description: |-
  Use this data source to get information about a specific port identified by ID
---

# ixapi_port (Data Source)

Use this data source to get information about a specific port identified by ID



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `billing_account` (String) An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*
- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `contract_ref` (String) A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)*
- `device` (String) The device the port.
- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.
- `media_type` (String) The media type of the port. Query the device's capabilities for available types.
- `name` (String) Name of the port (set by the exchange)
- `network_connection` (String)
- `operational_state` (String) The operational state of the port.
- `pop` (String) Same as the `pop` of the `device`.
- `purchase_order` (String) Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)*
- `role_assignments` (List of String)
- `speed` (Number)
- `state` (String)
- `status` (Block List) (see [below for nested schema](#nestedblock--status))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--status"></a>
### Nested Schema for `status`

Optional:

- `message` (String) A human readable message, describing the problem and may contain hints for resolution.
- `severity` (Number) We are using syslog severity levels: 0 = Emergency, 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning, 5 = Notice, 6 = Informational, 7 = Debug.
- `tag` (String) A machine readable message identifier.
- `timestamp` (String) The time and date when the event occured.

