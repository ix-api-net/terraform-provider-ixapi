
# Role assignment required for configuring a network service
data "ixapi_role_assignment" "impl" {
  consuming_account = "customer_id"
  role = "implementation"
}

