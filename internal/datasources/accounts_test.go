package datasources

import (
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/testdata"
)

func TestAccountSetResourceData(t *testing.T) {
	accountDataSource := NewAccountDataSource()
	res := accountDataSource.TestResourceData()

	account := testdata.NewAccount()
	if err := schemas.SetResourceData(account, res); err != nil {
		t.Fatal(err)
	}

	val, ok := res.GetOk("name")
	if !ok || val.(string) != "account name" {
		t.Error("unexpected name:", val)
	}
	val, ok = res.GetOk("external_ref")
	if !ok || val.(string) != "ext ref" {
		t.Error("unexpected external ref:", val)
	}
}
