package schemas

import (
	"testing"

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
