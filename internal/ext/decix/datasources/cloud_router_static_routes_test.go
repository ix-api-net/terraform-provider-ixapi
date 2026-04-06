package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterStaticRoutesDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterStaticRoutesDataSource()
	res := dataSource.TestResourceData()

	routes := []*ixapi.StaticRoute{
		{
			ID:                    "route-1",
			Name:                  "route-one",
			Prefix:                "10.0.0.0/24",
			NextHop:               "192.168.1.1",
			NetworkServiceConfigs: []string{"nsc-abc"},
			VRF:                   "vrf-xyz",
		},
		{
			ID:                    "route-2",
			Name:                  "route-two",
			Prefix:                "172.16.0.0/16",
			NextHop:               "aggregate",
			NetworkServiceConfigs: []string{"nsc-def"},
			VRF:                   "vrf-xyz",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/static-routes": routes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterStaticRoutesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	result := res.Get("static_routes").([]any)
	if len(result) != 2 {
		t.Errorf("expected 2 static routes, got %d", len(result))
	}
}

func TestCloudRouterStaticRouteDataSourceReadByID(t *testing.T) {
	dataSource := NewDecixCloudRouterStaticRouteDataSource()
	res := dataSource.TestResourceData()
	res.Set("id", "route-1")

	route := &ixapi.StaticRoute{
		ID:                    "route-1",
		Name:                  "route-one",
		Prefix:                "10.0.0.0/24",
		NextHop:               "192.168.1.1",
		NetworkServiceConfigs: []string{"nsc-abc"},
		VRF:                   "vrf-xyz",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/static-routes/route-1": route,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterStaticRouteRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "route-1" {
		t.Errorf("expected ID route-1, got %s", res.Id())
	}
	if res.Get("name").(string) != "route-one" {
		t.Errorf("unexpected name: %s", res.Get("name"))
	}
}
