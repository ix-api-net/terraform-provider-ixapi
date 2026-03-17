package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestRoutingFunctionSetResourceData(t *testing.T) {
	ds := NewRoutingFunctionDataSource()
	res := ds.TestResourceData()

	rf := testdata.NewRoutingFunction()
	if err := schemas.SetResourceData(rf, res); err != nil {
		t.Fatal(err)
	}

	val, ok := res.GetOk("managing_account")
	if !ok || val.(string) != "managing:123" {
		t.Error("unexpected managing_account:", val)
	}
	val, ok = res.GetOk("asn")
	if !ok || val.(int) != 65000 {
		t.Error("unexpected asn:", val)
	}
	val, ok = res.GetOk("external_ref")
	if !ok || val.(string) != "ext-ref-rf" {
		t.Error("unexpected external_ref:", val)
	}
	val, ok = res.GetOk("state")
	if !ok || val.(string) != "production" {
		t.Error("unexpected state:", val)
	}
}

func TestRoutingFunctionRead(t *testing.T) {
	ds := NewRoutingFunctionDataSource()
	res := ds.TestResourceData()
	res.Set("id", "rf:42")

	rf := testdata.NewRoutingFunction()
	api := ixapi.NewTestClient(map[string]any{
		"/routing-functions/rf:42": rf,
	})

	ctx := context.Background()
	diags := routingFunctionRead(ctx, res, api)
	if diags.HasError() {
		t.Fatal(diags)
	}

	if res.Id() != "rf:42" {
		t.Error("unexpected id:", res.Id())
	}
	if res.Get("asn").(int) != 65000 {
		t.Error("unexpected asn:", res.Get("asn"))
	}
}

func TestRoutingFunctionReadRequiresID(t *testing.T) {
	ds := NewRoutingFunctionDataSource()
	res := ds.TestResourceData()

	api := ixapi.NewTestClient(map[string]any{})

	ctx := context.Background()
	diags := routingFunctionRead(ctx, res, api)
	if !diags.HasError() {
		t.Error("expected error when id is not provided")
	}
}

func TestRoutingFunctionsRead(t *testing.T) {
	ds := NewRoutingFunctionsDataSource()
	res := ds.TestResourceData()

	rf1 := testdata.NewRoutingFunction()
	rf2 := &ixapi.RoutingFunction{
		ID:               "rf:99",
		ManagingAccount:  "managing:999",
		ConsumingAccount: "consuming:999",
		BillingAccount:   "billing:999",
		ProductOffering:  "product:999",
		ASN:              65100,
		State:            "decommission_requested",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/routing-functions": []*ixapi.RoutingFunction{rf1, rf2},
	})

	ctx := context.Background()
	diags := routingFunctionsRead(ctx, res, api)
	if diags.HasError() {
		t.Fatal(diags)
	}

	rfs := res.Get("routing_functions").([]any)
	if len(rfs) != 2 {
		t.Fatalf("expected 2 routing functions, got %d", len(rfs))
	}
}

func TestRoutingFunctionsReadFilterByState(t *testing.T) {
	ds := NewRoutingFunctionsDataSource()
	res := ds.TestResourceData()
	res.Set("state", "production")

	rf1 := testdata.NewRoutingFunction()
	rf2 := &ixapi.RoutingFunction{
		ID:               "rf:99",
		ManagingAccount:  "managing:999",
		ConsumingAccount: "consuming:999",
		BillingAccount:   "billing:999",
		ProductOffering:  "product:999",
		ASN:              65100,
		State:            "decommission_requested",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/routing-functions": []*ixapi.RoutingFunction{rf1, rf2},
	})

	ctx := context.Background()
	diags := routingFunctionsRead(ctx, res, api)
	if diags.HasError() {
		t.Fatal(diags)
	}

	rfs := res.Get("routing_functions").([]any)
	if len(rfs) != 1 {
		t.Fatalf("expected 1 routing function after state filter, got %d", len(rfs))
	}
}
