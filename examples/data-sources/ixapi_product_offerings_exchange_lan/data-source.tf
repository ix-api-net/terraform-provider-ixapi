
# Get all exchange lan product offerings for a metro area
data "ixapi_metro_area" "fra" {
  iata_code = "fra"
}

data "ixapi_product_offerings_connection" "peering_fra" {
  handover_metro_area = data.ixapi_metro_area.fra.id
  physical_port_speed = 10000
}

