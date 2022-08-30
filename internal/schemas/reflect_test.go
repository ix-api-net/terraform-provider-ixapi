package schemas

import (
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
