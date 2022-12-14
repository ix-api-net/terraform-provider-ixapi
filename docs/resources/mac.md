---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_mac Resource - terraform-provider-ixapi"
subcategory: ""
description: |-
  Use the ixapi_mac resource to register a mac address. Attention: MAC addresses can only be created and destroyed. To change a MAC-Address, you have to create a new resource.
---

# ixapi_mac (Resource)

Use the `ixapi_mac` resource to register a mac address. *Attention:* MAC addresses can only be created and destroyed. To change a MAC-Address, you have to create a new resource.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Unicast MAC address, formatted hexadecimal values with colons.
- `consuming_account` (String) The `id` of the account consuming a service.  Used to be `owning_customer`.
- `managing_account` (String) The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities.

### Optional

- `external_ref` (String) Reference field, free to use for the API user. *(Sensitive Property)*
- `valid_not_after` (String) When a mac address is assigned to an NSC, and the current datetime is before this value, the MAC address *can* be used on the peering platform.  Afterwards, it is supposed to be unassigned from the NSC and cannot any longer be used on the peering platform.  If the value is null or the property does not exist, the MAC address is valid indefinitely. The value may not be in the past.
- `valid_not_before` (String) When a mac address is assigned to a NSC, and the current datetime is before this value, then the MAC address *cannot* be used on the peering platform.  Afterwards, it is supposed to be available. If the value is `null` or the property does not exist, the mac address is valid from the creation date.

### Read-Only

- `id` (String) The ID of this resource.


