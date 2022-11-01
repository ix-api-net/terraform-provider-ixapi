
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
