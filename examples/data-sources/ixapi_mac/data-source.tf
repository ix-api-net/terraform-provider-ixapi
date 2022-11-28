
# Get a mac address using the external ref
data "ixapi_mac" "customer1" {
  external_ref = "customer:1,new"
}

