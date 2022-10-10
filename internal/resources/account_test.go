package resources

import (
	"context"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/testdata"
)

func TestAccountRead(t *testing.T) {
	resource := NewAccountResource()
	res := resource.Data(nil)
	res.SetId("23")

	// Testclient with data for AccountsRead
	acc := testdata.NewAccount()
	api := ixapi.NewTestClient(map[string]any{
		"/accounts/23": acc,
	})

	ctx := context.Background()
	diags := accountRead(ctx, res, api)
	if diags != nil {
		t.Fatal(diags)
	}

	// Check resource
	if res.Get("name").(string) != "account name" {
		t.Error("unexpected result:", res)
	}
}

func TestAccountCreate(t *testing.T) {
	resource := NewAccountResource()
	res := resource.Data(nil)

	addr, err := schemas.FlattenModel(testdata.NewAddress())
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Set("address", []any{addr}); err != nil {
		t.Fatal(err)
	}
	billing, err := schemas.FlattenModel(testdata.NewBillingInformation())
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Set("billing_information", []any{billing}); err != nil {
		t.Fatal(err)
	}

	// Testclient with data for AccountCreate and Read
	acc := testdata.NewAccount() // ID: 23
	api := ixapi.NewTestClient(map[string]any{
		"/accounts": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			ixapi.AssertBodyContains(t, body, `"locality":"Berlin"`)
			ixapi.AssertBodyContains(t, body, `billing_information`)
			ixapi.AssertBodyContains(t, body, `NL1235890`)
			return acc, nil
		}),
		"/accounts/23": acc,
	})

	ctx := context.Background()
	diags := accountCreate(ctx, res, api)
	if diags != nil {
		t.Fatal(diags)
	}

	// Check resource
	if res.Get("name").(string) != "account name" {
		t.Error("unexpected result:", res)
	}
}
