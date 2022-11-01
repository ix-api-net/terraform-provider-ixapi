package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
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

func TestAccountUpdate(t *testing.T) {
	resource := NewAccountResource()
	res := resource.Data(nil)
	res.SetId("23")

	// TODO: I do not know how to trigger HasChange()

	// Testclient with data for AccountCreate and Read
	acc := testdata.NewAccount() // ID: 23
	api := ixapi.NewTestClient(map[string]any{
		"/accounts/23": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return acc, nil
		}),
	})

	ctx := context.Background()
	diags := accountUpdate(ctx, res, api)
	if diags != nil {
		t.Fatal(diags)
	}
}

func TestAccountDelete(t *testing.T) {
	resource := NewAccountResource()
	res := resource.Data(nil)
	res.SetId("23")

	// Testclient with data for AccountCreate and Read
	acc := testdata.NewAccount() // ID: 23
	reqs := 0
	api := ixapi.NewTestClient(map[string]any{
		"/accounts/23": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			if reqs == 1 {
				return nil, nil // NotFound
			}
			reqs++
			return acc, nil
		}),
	})

	ctx := context.Background()
	diags := accountDelete(ctx, res, api)
	if diags != nil {
		t.Fatal(diags)
	}
	if res.Id() != "" {
		t.Error("id of resource should be unset but is:", res.Id())
	}
}
