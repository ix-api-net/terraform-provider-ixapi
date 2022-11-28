
# Get a connection product offering in order to participate
# in the metro area network fra1 from the 'Multi-Tiered Data DC41'
# datacenter.
data "ixapi_metro_area_network" "fra" {
  name = "fra1"  
}

data "ixapi_facility" "dc1" {
  name = "Multi-Tiered Data DC41"
}

data "ixapi_pop" "dc1_fra1" {
  facility = data.ixapi_facility.dc1.id
  metro_area_network = data.ixapi_metro_area_network.fra.id
}

data "ixapi_product_offering_connection" "dc1_fra" {
  handover_pop = data.ixapi_pop.dc1_fra1.id
  cross_connect_initiator = "exchange"
}

