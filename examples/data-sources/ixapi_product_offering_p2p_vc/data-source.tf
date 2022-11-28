
# Get a p2p virtual circuit product offering identified by
# handover and service metro area
data "ixapi_metro_area_network" "lon" {
  name = "lon1"  
}

data "ixapi_product_offering_p2p_vc" "lon1" {
  service_metro_area_network = data.ixapi_metro_area_network.lon.id
  handover_metro_area_network = data.ixapi_metro_area_network.lon.id
}
