package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterNetworkServiceConfigsP2PVCDataSourceRead(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsP2PVCDataSource()
	res := dataSource.TestResourceData()

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "2",
			Type:             "p2p_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "501",
			Address:          "192.0.2.5/30",
			BGPNeighbor:      "192.0.2.6",
			BGPNeighborASN:   65100,
			AdminStatus:      "enabled",
			BFDEnabled:       false,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsP2PVCRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	nscs := res.Get("cloud_router_network_service_configs").([]any)
	if len(nscs) != 1 {
		t.Errorf("expected 1 NSC, got %d", len(nscs))
	}
}

func TestCloudRouterNetworkServiceConfigsP2PVCDataSourceReadWithBFDFilter(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsP2PVCDataSource()
	res := dataSource.TestResourceData()
	res.Set("bfd_enabled", true)

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "2",
			Type:             "p2p_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "501",
			Address:          "192.0.2.5/30",
			BGPNeighbor:      "192.0.2.6",
			BGPNeighborASN:   65100,
			AdminStatus:      "enabled",
			BFDEnabled:       true,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsP2PVCRead(ctx, res, api)
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

func TestCloudRouterNetworkServiceConfigsP2PVCDataSourceReadWithLimitOffset(t *testing.T) {
	dataSource := NewCloudRouterNetworkServiceConfigsP2PVCDataSource()
	res := dataSource.TestResourceData()
	res.Set("limit", 1)
	res.Set("offset", 1)

	configs := []*ixapi.CloudRouterNetworkServiceConfig{
		{
			ID:               "2",
			Type:             "p2p_vc",
			ManagingAccount:  "100",
			BillingAccount:   "100",
			ConsumingAccount: "100",
			CloudRouter:      "274",
			NetworkService:   "501",
			Address:          "192.0.2.5/30",
			BGPNeighbor:      "192.0.2.6",
			BGPNeighborASN:   65100,
			AdminStatus:      "enabled",
			BFDEnabled:       false,
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": configs,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigsP2PVCRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	nscs := res.Get("cloud_router_network_service_configs").([]any)
	if len(nscs) != 1 {
		t.Errorf("expected 1 NSC with limit=1&offset=1, got %d", len(nscs))
	}
}
