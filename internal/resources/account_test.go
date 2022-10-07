package resources

import (
	"context"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/testdata"
)

func TestAccountRead(t *testing.T) {
	resource := NewAccountResource()
	res := resource.Data(nil)
	res.SetId("23")

	acc := testdata.NewAccount()

	api := ixapi.NewTestClient(map[string]any{
		"/accounts/23": acc,
	})

	ctx := context.Background()
	diags := accountRead(ctx, res, api)

	if diags != nil {
		t.Fatal(diags)
	}

	if res.Get("name").(string) != "account name" {
		t.Error("unexpected result:", res)
	}
}
