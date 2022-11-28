
# Get a p2mp virtual circuit product offering identified by
# handover and service metro area
data "ixapi_metro_area_network" "fra1" {
  name = "fra1"  
}

data "ixapi_product_offering_p2mp_vc" "fra1" {
  service_metro_area_network = data.ixapi_metro_area_network.fra1.id
  handover_metro_area_network = data.ixapi_metro_area_network.fra1.id
}
