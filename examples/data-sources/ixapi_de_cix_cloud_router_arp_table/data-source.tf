data "ixapi_de_cix_cloud_router_arp_table" "example" {
  vrf = "vrf-123"
}

output "arp_table" {
  value = data.ixapi_de_cix_cloud_router_arp_table.example.arp_entries
}
