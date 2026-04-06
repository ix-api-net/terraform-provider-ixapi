package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterRoutesDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterRoutesDataSource()
	res := dataSource.TestResourceData()
	res.Set("vrf", "vrf-1")

	routes := []*ixapi.VrfRoute{
		{
			VRF:                  "vrf-1",
			ReceivedAt:           "2024-10-28T22:22:22Z",
			Prefix:               "192.168.5.0/30",
			DeviceFQDN:           "some.switch.com",
			Metric:               10,
			Protocol:             "bgp",
			Distance:             0,
			NextHop:              "172.16.251.18",
			NetworkServiceConfig: "nsc-abc",
		},
		{
			VRF:        "vrf-1",
			ReceivedAt: "2024-10-28T22:22:22Z",
			Prefix:     "10.0.0.0/8",
			DeviceFQDN: "some.switch.com",
			Metric:     10,
			Protocol:   "local",
			Distance:   0,
			NextHop:    "192.168.5.1",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/routes": routes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterRoutesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	result := res.Get("routes").([]any)
	if len(result) != 2 {
		t.Errorf("expected 2 routes, got %d", len(result))
	}
}

func TestCloudRouterRoutesDataSourceReadEmpty(t *testing.T) {
	dataSource := NewDecixCloudRouterRoutesDataSource()
	res := dataSource.TestResourceData()

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/routes": []*ixapi.VrfRoute{},
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterRoutesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	result := res.Get("routes").([]any)
	if len(result) != 0 {
		t.Errorf("expected 0 routes, got %d", len(result))
	}
}
