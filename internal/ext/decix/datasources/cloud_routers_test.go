package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRoutersDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRoutersDataSource()
	res := dataSource.TestResourceData()

	externalRef := "router-ref-1"
	routers := []*ixapi.CloudRouter{
		{
			ID:               "274",
			State:            "active",
			ManagingAccount:  "100",
			ConsumingAccount: "200",
			BillingAccount:   "300",
			ProductOffering:  "1",
			ASN:              65001,
			Capacity:         1000,
			MetroAreaNetwork: "de-fra",
			ExternalRef:      &externalRef,
		},
		{
			ID:               "275",
			State:            "active",
			ManagingAccount:  "100",
			ConsumingAccount: "200",
			BillingAccount:   "300",
			ProductOffering:  "1",
			ASN:              65002,
			Capacity:         500,
			MetroAreaNetwork: "de-fra",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": routers,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRoutersRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	crs := res.Get("cloud_routers").([]any)
	if len(crs) != 2 {
		t.Errorf("expected 2 cloud routers, got %d", len(crs))
	}
}

func TestCloudRoutersDataSourceReadWithFilter(t *testing.T) {
	dataSource := NewDecixCloudRoutersDataSource()
	res := dataSource.TestResourceData()
	res.Set("managing_account", "100")

	routers := []*ixapi.CloudRouter{
		{
			ID:               "274",
			State:            "active",
			ManagingAccount:  "100",
			ConsumingAccount: "200",
			BillingAccount:   "300",
			ProductOffering:  "1",
			ASN:              65001,
			Capacity:         1000,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": routers,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRoutersRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	crs := res.Get("cloud_routers").([]any)
	if len(crs) != 1 {
		t.Errorf("expected 1 cloud router, got %d", len(crs))
	}

	cr := crs[0].(map[string]any)
	if cr["managing_account"].(string) != "100" {
		t.Errorf("expected managing_account=100, got %v", cr["managing_account"])
	}
}

func TestCloudRouterDataSourceReadByID(t *testing.T) {
	dataSource := NewDecixCloudRouterDataSource()
	res := dataSource.TestResourceData()
	res.Set("id", "274")

	router := &ixapi.CloudRouter{
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
		"/api/v3/decix-vrf-v1/vrfs/274": router,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "274" {
		t.Errorf("expected ID to be 274, got %s", res.Id())
	}
	if res.Get("asn").(int) != 65001 {
		t.Errorf("unexpected asn: %v", res.Get("asn"))
	}
	if res.Get("metro_area_network").(string) != "de-fra" {
		t.Errorf("unexpected metro_area_network: %v", res.Get("metro_area_network"))
	}
}

func TestCloudRouterDataSourceReadByExternalRef(t *testing.T) {
	dataSource := NewDecixCloudRouterDataSource()
	res := dataSource.TestResourceData()
	res.Set("external_ref", "my-router-ref")

	externalRef := "my-router-ref"
	routers := []*ixapi.CloudRouter{
		{
			ID:               "274",
			State:            "active",
			ManagingAccount:  "100",
			ConsumingAccount: "200",
			BillingAccount:   "300",
			ProductOffering:  "1",
			ASN:              65001,
			Capacity:         1000,
			ExternalRef:      &externalRef,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": routers,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "274" {
		t.Errorf("expected ID to be 274, got %s", res.Id())
	}
}

func TestCloudRouterDataSourceRead_MultipleResultsError(t *testing.T) {
	dataSource := NewDecixCloudRouterDataSource()
	res := dataSource.TestResourceData()
	res.Set("external_ref", "shared-ref")

	externalRef := "shared-ref"
	routers := []*ixapi.CloudRouter{
		{
			ID:          "274",
			ExternalRef: &externalRef,
		},
		{
			ID:          "275",
			ExternalRef: &externalRef,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/vrfs": routers,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterRead(ctx, res, api)

	if !diag.HasError() {
		t.Error("expected error when multiple cloud routers match external_ref")
	}
}
