
data "ixapi_port_reservations" "fra1" {
  network_connection = resource.ixapi_connection.fra1.id
}

