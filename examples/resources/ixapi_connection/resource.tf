
# Create a connection for a reseller

data "ixapi_account" "reseller" {
  external_ref = "demo_reseller"
}

data "ixapi_role_assignment" "impl" {
  consuming_account = data.ixapi_account.reseller.id
  role = "implementation"
}

locals {
  product = "502"
  reseller = data.ixapi_account.reseller.id
  impl = data.ixapi_role_assignment.impl.id
}

# Create connection with 4 ports
resource "ixapi_connection" "fra1" {
  product_offering = local.product
  managing_account = local.reseller
  consuming_account = local.reseller
  billing_account = local.reseller
  role_assignments = [ local.impl ]
  mode = "lag_lacp"
  port_quantity = 4
}

