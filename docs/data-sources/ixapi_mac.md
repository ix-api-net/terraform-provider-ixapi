---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_mac Data Source - ix-api-terraform-provider"
subcategory: ""
description: |-
  Use this data source to reference a single mac by address, external ref or id
---

# ixapi_mac (Data Source)

Use this data source to reference a single mac by address, external ref or id



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address` (String) Unicast MAC address, formatted hexadecimal values with colons.
- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.
- `valid_not_after` (String)
- `valid_not_before` (String)

### Read-Only

- `id` (String) The ID of this resource.

