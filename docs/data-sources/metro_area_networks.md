---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ixapi_metro_area_networks Data Source - terraform-provider-ixapi"
subcategory: ""
description: |-
  Retrieve a list of metro area networks filtered by name, metro area or service provider
---

# ixapi_metro_area_networks (Data Source)

Retrieve a list of metro area networks filtered by name, metro area or service provider

## Example Usage

```terraform
# Get all metro area networks in a metro area
data "ixapi_metro_area" "fra" {
  iata_code = "fra"
}

data "ixapi_metro_area_networks" "fra" {
  metro_area = data.ixapi_metro_area.fra.id
}

# Get all metro area networks by name
data "ixapi_metro_area_networks" "dus" {
  name = "peering_dus1"
}

# Or using a service provider
data "ixapi_metro_area_networks" "ixp" {
  service_provider = "de-cix"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `metro_area` (String) Filter by metro area id, see metro area data source
- `metro_area_networks` (Block List) (see [below for nested schema](#nestedblock--metro_area_networks))
- `name` (String) Filter metro area network by name, e.g. FRA
- `service_provider` (String) Filter by service provider operating the network

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--metro_area_networks"></a>
### Nested Schema for `metro_area_networks`

Optional:

- `metro_area` (String) The id of the metro area.
- `name` (String) The name of the metro area network.
- `pops` (List of String)
- `service_provider` (String) The service provider is operating the network. Usually the exchange.

Read-Only:

- `id` (String) The ID of this resource.


