# Get allocated ports for a connection
data "ixapi_ports" "fra1" {
  network_connection = resource.ixapi_connection.fra1.id
}
