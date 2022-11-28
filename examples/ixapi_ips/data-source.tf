
# Get all ipv6 addresses for a customer
data "ixapi_ips" "customer1v6" {
  consuming_customer = "customer_id"
  network_service = "mp2mp_network_service_id"
  version = 6
}

