---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_metro_area Data Source - terraform-provider-ixapi"
subcategory: ""
description: |-
  Use the metro_area datasource to retrieve a metro area by unlocode or iatacode
---

# ixapi_metro_area (Data Source)

Use the `metro_area` datasource to retrieve a metro area by un_locode or iata_code

## Example Usage

```terraform
# Get metro area using IATA code
data "ixapi_metro_area" "fra" {
  iata_code = "FRA"
}

# Get the metro area using an UN LOCODE
data "ixapi_metro_area" "fra" {
  un_locode = "DE FRA"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `display_name` (String) The name of the metro area. Likely the same as the IATA code.
- `facilities` (List of String)
- `iata_code` (String) The three letter IATA airport code for identiying the metro area.
- `metro_area_networks` (List of String)
- `un_locode` (String) The UN/LOCODE for identifying the metro area.

### Read-Only

- `id` (String) The ID of this resource.


