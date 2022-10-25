package datasources

import (
	"testing"

	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/testdata"
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
