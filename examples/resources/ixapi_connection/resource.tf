
# Get a connection product offering in order to participate
# in the metro area network fra1 from the 'Multi-Tiered Data DC41'
# datacenter.
data "ixapi_metro_area_network" "fra" {
  name = "fra1"  
}

data "ixapi_facility" "dc1" {
  name = "Multi-Tiered Data DC41"
}

# The connection will be orderered for a reseller account
data "ixapi_account" "reseller" {
  external_ref = "demo_reseller"
}

data "ixapi_pop" "dc1_fra1" {
  facility = data.ixapi_facility.dc1.id
  metro_area_network = data.ixapi_metro_area_network.fra.id
}

data "ixapi_product_offering_connection" "dc1_fra" {
  handover_pop = data.ixapi_pop.dc1_fra1.id
  cross_connect_initiator = "exchange"
}

data "ixapi_role_assignment" "impl" {
  consuming_account = data.ixapi_account.reseller.id
  role = "implementation"
}

locals {
  product = data.ixapi_product_offering_connection.dc1_fra.id
  reseller = data.ixapi_account.reseller.id
  impl = data.ixapi_role_assignment.impl.id
}

# Create connection with 2 ports
resource "ixapi_connection" "fra1" {
  product_offering = local.product
  managing_account = local.reseller
  consuming_account = local.reseller
  billing_account = local.reseller
  role_assignments = [ local.impl ]
  subscriber_side_demarcs = [ "f23.20.49.1", "f23.20.49.2" ]
  mode = "lag_lacp"
  port_quantity = 2
}

output "connection" {
  value = resource.ixapi_connection.fra1
}

