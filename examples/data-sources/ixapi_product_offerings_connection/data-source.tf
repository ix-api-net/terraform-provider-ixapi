
# Find product offerings for a connection
data "ixapi_metro_area" "fra" {
  iata_code = "fra"
}

data "ixapi_product_offerings_connection" "conn_10g_fra" {
  handover_metro_area = data.ixapi_metro_area.fra.id
  physical_port_speed = 10000
}

output "product_offerings" {
  value = data.ixapi_product_offerings_connection.conn_10g_fra.product_offerings
}
