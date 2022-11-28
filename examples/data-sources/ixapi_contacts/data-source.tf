
# Get all contacts assigned to an account
data "ixapi_contacts" "customer1" {
  consuming_account = "customer1_id"
}

# Get all contacts managed by a reseller
data "ixapi_contacts" "subcustomers" {
  managing_account = "reseller_id"
}

