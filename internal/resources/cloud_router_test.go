package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterRequestFromResourceData(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("billing_account", "300")
	res.Set("product_offering", "1")
	res.Set("asn", 65001)
	res.Set("capacity", 1000)
	res.Set("external_ref", "test-ref")

	req, err := cloudRouterRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.ManagingAccount != "100" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if req.ConsumingAccount != "200" {
		t.Error("unexpected consuming_account:", req.ConsumingAccount)
	}
	if req.BillingAccount != "300" {
		t.Error("unexpected billing_account:", req.BillingAccount)
	}
	if req.ProductOffering != "1" {
		t.Error("unexpected product_offering:", req.ProductOffering)
	}
	if req.ASN != 65001 {
		t.Error("unexpected asn:", req.ASN)
	}
	if req.Capacity != 1000 {
		t.Error("unexpected capacity:", req.Capacity)
	}
	if req.ExternalRef == nil || *req.ExternalRef != "test-ref" {
		t.Error("unexpected external_ref:", req.ExternalRef)
	}
}

func TestCloudRouterRead(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.SetId("274")

	cr := &ixapi.CloudRouter{
		ID:               "274",
		State:            "active",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		BillingAccount:   "300",
		ProductOffering:  "1",
		ASN:              65001,
		Capacity:         1000,
		MetroAreaNetwork: "de-fra",
	}
	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs/274": cr,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("asn").(int) != 65001 {
		t.Error("unexpected asn in resource:", res.Get("asn"))
	}
	if res.Get("capacity").(int) != 1000 {
		t.Error("unexpected capacity in resource:", res.Get("capacity"))
	}
	if res.Get("state").(string) != "active" {
		t.Error("unexpected state in resource:", res.Get("state"))
	}
	if res.Get("metro_area_network").(string) != "de-fra" {
		t.Error("unexpected metro_area_network in resource:", res.Get("metro_area_network"))
	}
}

func TestCloudRouterDelete_404Success(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.SetId("274")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}

	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}

func TestCloudRouterDelete_OtherErrorFails(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.SetId("274")

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs/274": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return nil, &ixapi.APIError{
				ProblemResponse: ixapi.ProblemResponse{
					Type:   "internal_error",
					Title:  "Internal Server Error",
					Status: 500,
				},
			}
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterDelete(ctx, res, api)

	if err == nil {
		t.Error("delete should fail on non-404 errors")
	}

	if ixapi.IsErrNotFound(err) {
		t.Error("error should not be NotFoundError")
	}
}

func TestCloudRouterCreate(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("billing_account", "300")
	res.Set("product_offering", "1")
	res.Set("asn", 65001)
	res.Set("capacity", 1000)

	createdRouter := &ixapi.CloudRouter{
		ID:               "274",
		State:            "production",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		BillingAccount:   "300",
		ProductOffering:  "1",
		ASN:              65001,
		Capacity:         1000,
		MetroAreaNetwork: "de-fra",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return createdRouter, nil
		}),
		"/api/v3/decix-vrf-v1/vrfs/274": createdRouter,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "274" {
		t.Error("expected resource ID to be set to 274, got:", res.Id())
	}

	if res.Get("state").(string) != "production" {
		t.Error("unexpected state in resource:", res.Get("state"))
	}

	if res.Get("metro_area_network").(string) != "de-fra" {
		t.Error("unexpected metro_area_network in resource:", res.Get("metro_area_network"))
	}
}

func TestCloudRouterCreate_WithExternalRef(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("billing_account", "300")
	res.Set("product_offering", "1")
	res.Set("asn", 65001)
	res.Set("capacity", 1000)
	res.Set("external_ref", "my-router-ref")

	externalRef := "my-router-ref"
	createdRouter := &ixapi.CloudRouter{
		ID:               "275",
		State:            "production",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		BillingAccount:   "300",
		ProductOffering:  "1",
		ASN:              65001,
		Capacity:         1000,
		ExternalRef:      &externalRef,
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return createdRouter, nil
		}),
		"/api/v3/decix-vrf-v1/vrfs/275": createdRouter,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "275" {
		t.Error("expected resource ID to be set to 275, got:", res.Id())
	}

	if res.Get("external_ref").(string) != "my-router-ref" {
		t.Error("unexpected external_ref in resource:", res.Get("external_ref"))
	}
}

func TestCloudRouterCreate_APIError(t *testing.T) {
	resource := NewCloudRouterResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("billing_account", "300")
	res.Set("product_offering", "1")
	res.Set("asn", 65001)
	res.Set("capacity", 1000)

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return nil, &ixapi.APIError{
				ProblemResponse: ixapi.ProblemResponse{
					Type:   "validation_error",
					Title:  "Invalid Request",
					Status: 400,
				},
			}
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterCreate(ctx, res, api)

	if err == nil {
		t.Error("create should fail when API returns error")
	}

	if res.Id() != "" {
		t.Error("resource ID should not be set on error")
	}
}


