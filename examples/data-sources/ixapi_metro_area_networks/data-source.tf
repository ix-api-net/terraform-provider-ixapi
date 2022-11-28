

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
