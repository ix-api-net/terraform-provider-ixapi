
# Get connection by ID
data "ixapi_connection" "conn1" {
  id = "12345"
}

# Get by external ref
data "ixapi_connection" "customer1" {
  external_ref = "customer:1"
}

