
# Get all implementation role assignments for an account
data "ixapi_role_assignments" "customer1_billing" {
  contact = "customer1_contact_id"
  role_name = "billing"
}

# Or using the role data source
data "ixapi_role" "noc" {
  name = "noc"
}

data "ixapi_role_assignments" "noc" {
  role = data.ixapi_role.noc.id
}

