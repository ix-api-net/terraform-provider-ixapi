
# Create a new customer
resource "ixapi_account" "customer" {
  managing_account = "<my account id>" 
  name = "customernet"

  legal_name = "Customer Networks Inc."

  billing_information {
    name = "Customer Networks Inc."
    vat_number = "XX1234567890"
    address {
      country = "DE"
      locality = "Berlin"
      postal_code = "10000"
      street_address = "Straßenweg 11"
    }
  }

  address {
    country = "DE"
    locality = "Berlin"
    postal_code = "10000"
    street_address = "Straßenweg 11"
  }

}
