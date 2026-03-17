package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestRFRequestFromResourceData(t *testing.T) {
	resource := NewRoutingFunctionResource()
	res := resource.Data(nil)
	res.Set("managing_account", "managing:1")
	res.Set("consuming_account", "consuming:2")
	res.Set("billing_account", "billing:3")
	res.Set("product_offering", "product:4")
	res.Set("asn", 65001)
	res.Set("external_ref", "ext-ref")
	res.Set("capacity", 500)

	req, err := rfRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.ManagingAccount != "managing:1" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if req.ConsumingAccount != "consuming:2" {
		t.Error("unexpected consuming_account:", req.ConsumingAccount)
	}
	if req.BillingAccount != "billing:3" {
		t.Error("unexpected billing_account:", req.BillingAccount)
	}
	if req.ProductOffering != "product:4" {
		t.Error("unexpected product_offering:", req.ProductOffering)
	}
	if req.ASN != 65001 {
		t.Error("unexpected asn:", req.ASN)
	}
	if req.ExternalRef == nil || *req.ExternalRef != "ext-ref" {
		t.Error("unexpected external_ref:", req.ExternalRef)
	}
	if req.Capacity == nil || *req.Capacity != 500 {
		t.Error("unexpected capacity:", req.Capacity)
	}
}

func TestRFRequestFromResourceDataOptionalAbsent(t *testing.T) {
	resource := NewRoutingFunctionResource()
	res := resource.Data(nil)
	res.Set("managing_account", "managing:1")
	res.Set("consuming_account", "consuming:2")
	res.Set("billing_account", "billing:3")
	res.Set("product_offering", "product:4")
	res.Set("asn", 65001)

	req, err := rfRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.ExternalRef != nil {
		t.Error("expected nil external_ref, got:", *req.ExternalRef)
	}
	if req.Capacity != nil {
		t.Error("expected nil capacity, got:", *req.Capacity)
	}
}

func TestRFRead(t *testing.T) {
	resource := NewRoutingFunctionResource()
	res := resource.Data(nil)
	res.SetId("rf:42")

	rf := testdata.NewRoutingFunction()
	api := ixapi.NewTestClient(map[string]any{
		"/routing-functions/rf:42": rf,
	})

	ctx := context.Background()
	err := rfRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("managing_account").(string) != "managing:123" {
		t.Error("unexpected managing_account:", res.Get("managing_account"))
	}
	if res.Get("consuming_account").(string) != "consuming:123" {
		t.Error("unexpected consuming_account:", res.Get("consuming_account"))
	}
	if res.Get("asn").(int) != 65000 {
		t.Error("unexpected asn:", res.Get("asn"))
	}
	if res.Get("state").(string) != "production" {
		t.Error("unexpected state:", res.Get("state"))
	}
}

func TestRFReadNotFound(t *testing.T) {
	resource := NewRoutingFunctionResource()
	res := resource.Data(nil)
	res.SetId("rf:notfound")

	api := ixapi.NewTestClient(map[string]any{})

	ctx := context.Background()
	err := rfRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "" {
		t.Error("expected empty ID after not found, got:", res.Id())
	}
}
