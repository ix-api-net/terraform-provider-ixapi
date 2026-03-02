package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterNetworkServiceConfigsCloudVCDataSourceRead(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsCloudVCDataSource()
	res := dataSource.TestResourceData()

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "1",
			Type:             "cloud_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "500",
			Address:          "192.0.2.1/30",
			BGPNeighbor:      "192.0.2.2",
			BGPNeighborASN:   64512,
			AdminStatus:      "enabled",
			BFDEnabled:       true,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsCloudVCRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	nscs := res.Get("cloud_router_network_service_configs").([]any)
	if len(nscs) != 1 {
		t.Errorf("expected 1 NSC, got %d", len(nscs))
	}
}

func TestCloudRouterNetworkServiceConfigsCloudVCDataSourceReadWithBFDFilter(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsCloudVCDataSource()
	res := dataSource.TestResourceData()
	res.Set("bfd_enabled", true)

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "1",
			Type:             "cloud_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "500",
			Address:          "192.0.2.1/30",
			BGPNeighbor:      "192.0.2.2",
			BGPNeighborASN:   64512,
			AdminStatus:      "enabled",
			BFDEnabled:       true,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsCloudVCRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	nscs := res.Get("cloud_router_network_service_configs").([]any)
	if len(nscs) != 1 {
		t.Errorf("expected 1 NSC with bfd_enabled=true, got %d", len(nscs))
	}

	nsc := nscs[0].(map[string]any)
	if nsc["bfd_enabled"].(bool) != true {
		t.Errorf("expected bfd_enabled=true, got %v", nsc["bfd_enabled"])
	}
}

func TestCloudRouterNetworkServiceConfigsCloudVCDataSourceReadWithLimitOffset(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsCloudVCDataSource()
	res := dataSource.TestResourceData()
	res.Set("limit", 1)
	res.Set("offset", 0)

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "1",
			Type:             "cloud_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "500",
			Address:          "192.0.2.1/30",
			BGPNeighbor:      "192.0.2.2",
			BGPNeighborASN:   64512,
			AdminStatus:      "enabled",
			BFDEnabled:       false,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsCloudVCRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	nscs := res.Get("cloud_router_network_service_configs").([]any)
	if len(nscs) != 1 {
		t.Errorf("expected 1 NSC with limit=1, got %d", len(nscs))
	}
}
