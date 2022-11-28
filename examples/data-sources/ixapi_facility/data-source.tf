
# Get a facility by peering db facility ID
data "ixapi_facility" "dc1" {
  peering_db_facility_id = "58"
}

