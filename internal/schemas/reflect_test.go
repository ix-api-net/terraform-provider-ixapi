package schemas

import (
	"encoding/json"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/testdata"
)

func TestSetResourceData(t *testing.T) {
	account := testdata.NewAccount()
	res := NewFlatResource()
	if err := SetResourceData(account, res); err != nil {
		t.Fatal(err)
	}
	t.Log(res.Flatten())
}

func TestFlattenModels(t *testing.T) {
	models := []*ixapi.MetroArea{
		&ixapi.MetroArea{
			IataCode: "ASDF",
		},
		&ixapi.MetroArea{
			IataCode: "FOO",
		},
	}

	flat, err := FlattenModels(models)
	if err != nil {
		t.Fatal(err)
	}
	if flat[0].(map[string]interface{})["iata_code"].(string) != "ASDF" {
		t.Error("unexpected item:", flat[0])
	}

	t.Log(flat)
}

const CreateMacResponse = `{"managing_account":"17","consuming_account":"26","external_ref":null,"address":"42:23:bc:8e:b8:b0","valid_not_before":"2022-10-03T08:45:41.389744Z", "valid_not_after":null,"id":"2"}`

func TestSetResourceMac(t *testing.T) {
	mac := &ixapi.MacAddress{}
	json.Unmarshal([]byte(CreateMacResponse), mac)

	res := NewFlatResource()
	err := SetResourceData(mac, res)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

const ExchangeLanProductOfferingsResponse = `[{"id":"520","name":"PlanetPEER","display_name":"Planet Peering (London)","resource_type":"network_service","handover_metro_area_network":"25","handover_metro_area":"LON","physical_port_speed":null,"service_provider":"","downgrade_allowed":true,"upgrade_allowed":true,"orderable_not_before":null,"orderable_not_after":null,"contract_terms":null,"notice_period":null,"provider_vlans":"single","service_metro_area_network":"25","service_metro_area":"LON","bandwidth_min":null,"bandwidth_max":null,"exchange_lan_network_service":"56","type":"exchange_lan"}]`

func TestFlattenExchangeLanProductOfferings(t *testing.T) {

	offerings := []*ixapi.ExchangeLanNetworkProductOffering{}
	err := json.Unmarshal([]byte(ExchangeLanProductOfferingsResponse), &offerings)
	if err != nil {
		t.Fatal(err)
	}

	flat, err := FlattenModels(offerings)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(flat)
}
