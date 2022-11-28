
# Get mp2mp virtual circuit product offerings for the
# metro area network named 'fra1'.
data "ixapi_metro_area_network" "fra" {
  name = "fra1"  
}

data "ixapi_product_offerings_mp2mp_vc" "all" {
  handover_metro_area_network = data.ixapi_metro_area_network.fra.id
}

