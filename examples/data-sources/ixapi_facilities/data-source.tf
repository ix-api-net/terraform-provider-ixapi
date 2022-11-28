
# Get all facilities in Frankfurt (Hessen)
data "ixapi_facilities" "fra" {
  postal_code = "60314"
}

# Get facilities matching an organisation
data "ixapi_facilities" "interxion" {
  organisation_name = "interxion"   
  metro_area = data.ixapi_metro_area.fra.id
}

# Metro area frankfurt
data "ixapi_metro_area" "fra" {
  iata_code = "fra"
}
