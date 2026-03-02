package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterNetworkServiceConfigPrefixesDataSourceRead(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigPrefixesDataSource()
	res := dataSource.TestResourceData()
	res.Set("network_service_config_id", "123")
	res.Set("mode", "gnmi")

	prefixes := []*ixapi.BGPPrefix{
		{
			ASPath:        "3320",
			IGPCost:       "0",
			Network:       "192.168.1.1/32",
			Label:         "-",
			Nexthop:       "10.0.1.1",
			Flags:         []string{"used", "valid", "best", "igp"},
			LastQueriedAt: "2022-04-14T21:22:20.837229+00:00",
		},
		{
			ASPath:        "174 3356",
			IGPCost:       "0",
			Network:       "203.0.113.0/24",
			Label:         "-",
			Nexthop:       "10.0.1.1",
			LocalPref:     intPtrPrefix(100),
			Flags:         []string{"used", "valid", "best"},
			LastQueriedAt: "2022-04-14T21:22:20.837229+00:00",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/123/prefixes": prefixes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigPrefixesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	prefs := res.Get("prefixes").([]any)
	if len(prefs) != 2 {
		t.Errorf("expected 2 prefixes, got %d", len(prefs))
	}
}

func TestCloudRouterNetworkServiceConfigPrefixesDataSourceReadEmpty(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigPrefixesDataSource()
	res := dataSource.TestResourceData()
	res.Set("network_service_config_id", "456")

	prefixes := []*ixapi.BGPPrefix{}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/456/prefixes": prefixes,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigPrefixesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	prefs := res.Get("prefixes").([]any)
	if len(prefs) != 0 {
		t.Errorf("expected 0 prefixes, got %d", len(prefs))
	}
}

func intPtrPrefix(i int) *int {
	return &i
}
