
# Get p2p virtual circuit product offerings for the
# metro area network named 'lon1'.
data "ixapi_metro_area_network" "lon" {
  name = "lon1"  
}

data "ixapi_product_offerings_p2p_vc" "all" {
 service_metro_area_network = data.ixapi_metro_area_network.lon.id
}

