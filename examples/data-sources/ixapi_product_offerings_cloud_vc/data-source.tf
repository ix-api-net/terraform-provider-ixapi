
# Get cloud virtual circuit product offerings
data "ixapi_product_offerings_cloud_vc" "all" {
  cloud_key = "fc4201-dn311-fj03992"
  bandwidth = 1000
  delivery_method = "dedicated"
}

