package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterNetworkServiceConfigsP2PVCDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterNetworkServiceConfigsP2PVCDataSource()
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
	dataSource := NewDecixCloudRouterNetworkServiceConfigsP2PVCDataSource()
	res := dataSource.TestResourceData()
	_ = res.Set("bfd_enabled", true)

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

func TestCloudRouterNetworkServiceConfigP2PVCDataSourceReadByID(t *testing.T) {
	dataSource := NewDecixCloudRouterNetworkServiceConfigP2PVCDataSource()
	res := dataSource.TestResourceData()
	_ = res.Set("id", "2")

	config := &ixapi.CloudRouterNetworkServiceConfig{
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
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/2": config,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterNetworkServiceConfigP2PVCDataRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "2" {
		t.Errorf("expected ID to be 2, got %s", res.Id())
	}
	if res.Get("bgp_neighbor_asn").(int) != 65100 {
		t.Errorf("unexpected bgp_neighbor_asn: %v", res.Get("bgp_neighbor_asn"))
	}
}

func TestCloudRouterNetworkServiceConfigsP2PVCDataSourceReadWithLimitOffset(t *testing.T) {
	dataSource := NewDecixCloudRouterNetworkServiceConfigsP2PVCDataSource()
	res := dataSource.TestResourceData()
	_ = res.Set("limit", 1)
	_ = res.Set("offset", 1)

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
