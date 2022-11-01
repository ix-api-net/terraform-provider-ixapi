
# Create contact for a given account
resource "ixapi_contact" "support" {
  managing_account = resource.ixapi_account.customer.managing_account
  consuming_account = resource.ixapi_account.customer.id
  roles = ["noc", "implementation" ]
  email = "support@customernet.example.net" 
  telephone = "+23 42 52 90"
}

# Use the role assignment data source to get the id of
# the assignemnt for referencing when required.
data "ixapi_role_assignment" "noc" {
  contact = resource.ixapi_contact.support.id
  role = "noc"
}

data "ixapi_role_assignment" "impl" {
  contact = resource.ixapi_contact.support.id
  role = "implementation"
}
