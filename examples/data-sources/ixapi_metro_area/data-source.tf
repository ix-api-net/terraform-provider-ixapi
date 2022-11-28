
# Get metro area using IATA code
data "ixapi_metro_area" "fra" {
  iata_code = "FRA"
}

# Get the metro area using an UN LOCODE
data "ixapi_metro_area" "fra" {
  un_locode = "DE FRA"
}


