
# Get role by name
data "ixapi_role" "impl" {
  name = "implementation"
}

data "ixapi_role" "billing" {
  name = "billing"
}

data "ixapi_role" "peering" {
  name = "peering"
}

data "ixapi_role" "noc" {
  name = "noc"
}

data "ixapi_role" "legal" {
  name = "legal"
}
