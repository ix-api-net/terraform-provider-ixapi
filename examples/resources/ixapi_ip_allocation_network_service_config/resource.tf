

# Get the ip allocation resource and set the fqdn of the IP address(es)
resource "ixapi_ip_allocation_network_service_config" "customernet_fra" {
    network_service_config = resource.ixapi_network_service_config_exchange_lan.customernet_fra.id
    fqdn = "gw1.fra.customer.example.net"
}

