
data "ixapi_facility" "dc1" {
  peeringdb_facility_id = "58"    
}

# Find devices supporting the media type 1000BASE-LX
# in the facility
data "ixapi_devices" "devices" {
  capability_media_type = "1000BASE-LX"
  facility = data.ixapi_facility.dc1.id
}

