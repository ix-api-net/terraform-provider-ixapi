
# Get customer macs in use in an exchange lan
data "ixapi_macs" "customer1" {
  consuming_customer = "customer1_id"
  network_service_config = "exchange_lan_fra_nsc_id"
}

