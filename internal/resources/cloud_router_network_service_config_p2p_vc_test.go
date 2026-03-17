package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterConfigP2PVCRequestFromResourceData(t *testing.T) {
	resource := NewCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("billing_account", "200")
	res.Set("consuming_account", "300")
	res.Set("cloud_router", "274")
	res.Set("network_service", "500")
	res.Set("address", "192.0.2.1/30")
	res.Set("bgp_neighbor", "192.0.2.2")
	res.Set("bgp_neighbor_asn", 64512)
	res.Set("admin_status", "enabled")
	res.Set("bfd_enabled", false)
	res.Set("network_connection", "conn-abc123")

	req, err := cloudRouterConfigP2PVCRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Type != "p2p_vc" {
		t.Error("unexpected type:", req.Type)
	}
	if req.ManagingAccount != "100" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if req.CloudRouter != "274" {
		t.Error("unexpected cloud_router:", req.CloudRouter)
	}
	if req.Connection == nil || *req.Connection != "conn-abc123" {
		t.Error("unexpected connection:", req.Connection)
	}
	if req.CloudVLAN != nil {
		t.Error("cloud_vlan should be nil for p2p_vc type")
	}
	if req.Handover != nil {
		t.Error("handover should be nil for p2p_vc type")
	}
}

func TestCloudRouterConfigP2PVCRead(t *testing.T) {
	resource := NewCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	res.SetId("20")

	config := &ixapi.CloudRouterNetworkServiceConfig{
		ID:               "20",
		Type:             "p2p_vc",
		ManagingAccount:  "100",
		BillingAccount:   "200",
		ConsumingAccount: "300",
		CloudRouter:      "274",
		NetworkService:   "500",
		Address:          "192.0.2.1/30",
		BGPNeighbor:      "192.0.2.2",
		BGPNeighborASN:   64512,
		AdminStatus:      "enabled",
		ConnectionID:     "conn-abc123",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/20": config,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigP2PVCRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("managing_account").(string) != "100" {
		t.Error("unexpected managing_account:", res.Get("managing_account"))
	}
	if res.Get("admin_status").(string) != "enabled" {
		t.Error("unexpected admin_status:", res.Get("admin_status"))
	}
	if res.Get("bgp_neighbor_asn").(int) != 64512 {
		t.Error("unexpected bgp_neighbor_asn:", res.Get("bgp_neighbor_asn"))
	}
}

func TestCloudRouterConfigP2PVCCreate(t *testing.T) {
	resource := NewCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	res.Set("managing_account", "100")
	res.Set("billing_account", "200")
	res.Set("consuming_account", "300")
	res.Set("cloud_router", "274")
	res.Set("network_service", "500")
	res.Set("address", "192.0.2.1/30")
	res.Set("bgp_neighbor", "192.0.2.2")
	res.Set("bgp_neighbor_asn", 64512)
	res.Set("admin_status", "enabled")
	res.Set("network_connection", "conn-abc123")

	created := &ixapi.CloudRouterNetworkServiceConfig{
		ID:               "20",
		Type:             "p2p_vc",
		ManagingAccount:  "100",
		BillingAccount:   "200",
		ConsumingAccount: "300",
		CloudRouter:      "274",
		NetworkService:   "500",
		Address:          "192.0.2.1/30",
		BGPNeighbor:      "192.0.2.2",
		BGPNeighborASN:   64512,
		AdminStatus:      "enabled",
		ConnectionID:     "conn-abc123",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return created, nil
		}),
		"/api/v3/decix-vrf-v1/network-service-configs/20": created,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigP2PVCCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "20" {
		t.Error("expected resource ID 20, got:", res.Id())
	}
	if res.Get("address").(string) != "192.0.2.1/30" {
		t.Error("unexpected address after create:", res.Get("address"))
	}
}

func TestCloudRouterConfigP2PVCDelete_404Success(t *testing.T) {
	resource := NewCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	res.SetId("20")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigP2PVCDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}
	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}
