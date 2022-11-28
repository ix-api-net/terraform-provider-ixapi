
# Get joining rules for a network service
data "ixapi_member_joining_rules" "svc1_allow" {
  network_service = "p2mp1_id"
  type = "allow"
}
