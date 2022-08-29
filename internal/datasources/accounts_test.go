package datasources

import (
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

func TestAccountSetResourceData(t *testing.T) {
	accountDataSource := NewAccountDataSource()
	res := accountDataSource.TestResourceData()

	ref := "ref"
	account := &ixapi.Account{
		Name:        "my account",
		ExternalRef: &ref,
	}
	schemas.AccountSetResourceData(account, res)

	val, ok := res.GetOk("name")
	if !ok || val.(string) != "my account" {
		t.Error("unexpected name:", val)
	}
	val, ok = res.GetOk("external_ref")
	if !ok || val.(string) != "ref" {
		t.Error("unexpected external ref:", val)
	}
}
