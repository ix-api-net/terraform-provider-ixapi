
# For now only trying to resolve an unique product offering
# by name is supported. If the name is not unique, you
# may resort to the ID or use the `ixapi_product_offerings_connection`
# data-source with additional filter criteria.
data "ixapi_product_offering_connection" "conn" {
  name = "IXP-Global-Connect-10000 Hamburg DC-19"
}

