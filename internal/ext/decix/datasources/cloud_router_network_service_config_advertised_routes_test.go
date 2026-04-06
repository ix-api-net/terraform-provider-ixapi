package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterNetworkServiceConfigAdvertisedRoutesDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterNetworkServiceConfigAdvertisedRoutesDataSource()
	res := dataSource.TestResourceData()
	res.Set("network_service_config_id", "123")

	routes := []*ixapi.BGPRoute{
		{
			Prefix:     "2001:db8:10::/48",
			NextHop:    "fe80::42:acff:fe11:2",
			ReceivedAt: "2024-03-20 20:45:59+01:00",
			ASPath:     []string{"65100", "65102"},
		},
		{
			Prefix:     "2001:db8:11::/48",
			NextHop:    "fe80::42:acff:fe11:2",
			ReceivedAt: "2024-03-20 20:45:59+01:00",
			ASPath:     []string{"65100", "65102"},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/123/advertised-routes": routes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigAdvertisedRoutesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	routesList := res.Get("routes").([]any)
	if len(routesList) != 2 {
		t.Errorf("expected 2 routes, got %d", len(routesList))
	}
}

func TestCloudRouterNetworkServiceConfigAdvertisedRoutesDataSourceReadEmpty(t *testing.T) {
	dataSource := NewDecixCloudRouterNetworkServiceConfigAdvertisedRoutesDataSource()
	res := dataSource.TestResourceData()
	res.Set("network_service_config_id", "456")

	routes := []*ixapi.BGPRoute{}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/456/advertised-routes": routes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigAdvertisedRoutesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	routesList := res.Get("routes").([]any)
	if len(routesList) != 0 {
		t.Errorf("expected 0 routes, got %d", len(routesList))
	}
}
