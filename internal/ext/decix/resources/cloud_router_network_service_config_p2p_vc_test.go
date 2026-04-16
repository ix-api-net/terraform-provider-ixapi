package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterConfigP2PVCRequestFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	_ = res.Set("managing_account", "100")
	_ = res.Set("billing_account", "200")
	_ = res.Set("consuming_account", "300")
	_ = res.Set("cloud_router", "274")
	_ = res.Set("network_service", "500")
	_ = res.Set("address", "192.0.2.1/30")
	_ = res.Set("bgp_neighbor", "192.0.2.2")
	_ = res.Set("bgp_neighbor_asn", 64512)
	_ = res.Set("admin_status", "enabled")
	_ = res.Set("bfd_enabled", false)
	_ = res.Set("network_connection", "conn-abc123")

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
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()
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
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	_ = res.Set("managing_account", "100")
	_ = res.Set("billing_account", "200")
	_ = res.Set("consuming_account", "300")
	_ = res.Set("cloud_router", "274")
	_ = res.Set("network_service", "500")
	_ = res.Set("address", "192.0.2.1/30")
	_ = res.Set("bgp_neighbor", "192.0.2.2")
	_ = res.Set("bgp_neighbor_asn", 64512)
	_ = res.Set("admin_status", "enabled")
	_ = res.Set("network_connection", "conn-abc123")

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

func TestCloudRouterConfigP2PVCPatchFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()

	t.Run("set fields are included", func(t *testing.T) {
		res := resource.Data(nil)
		_ = res.Set("admin_status", "enabled")
		_ = res.Set("policy_ingress", "my-ingress")

		patch := cloudRouterConfigP2PVCPatchFromResourceData(res)

		if patch.AdminStatus == nil || *patch.AdminStatus != "enabled" {
			t.Error("unexpected admin_status:", patch.AdminStatus)
		}
		if patch.PolicyIngress == nil || *patch.PolicyIngress != "my-ingress" {
			t.Error("unexpected policy_ingress:", patch.PolicyIngress)
		}
		if patch.PolicyEgress != nil {
			t.Error("policy_egress should be nil when not set:", patch.PolicyEgress)
		}
	})

	t.Run("unset fields are nil", func(t *testing.T) {
		res := resource.Data(nil)

		patch := cloudRouterConfigP2PVCPatchFromResourceData(res)

		if patch.AdminStatus != nil {
			t.Error("admin_status should be nil when not set:", patch.AdminStatus)
		}
		if patch.PolicyIngress != nil {
			t.Error("policy_ingress should be nil when not set:", patch.PolicyIngress)
		}
		if patch.PolicyEgress != nil {
			t.Error("policy_egress should be nil when not set:", patch.PolicyEgress)
		}
	})
}

func TestCloudRouterConfigP2PVCUpdate(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()
	res := resource.Data(nil)
	res.SetId("20")
	_ = res.Set("managing_account", "100")
	_ = res.Set("billing_account", "200")
	_ = res.Set("consuming_account", "300")
	_ = res.Set("cloud_router", "274")
	_ = res.Set("network_service", "500")
	_ = res.Set("address", "192.0.2.1/30")
	_ = res.Set("bgp_neighbor", "192.0.2.2")
	_ = res.Set("bgp_neighbor_asn", 64512)
	_ = res.Set("admin_status", "disabled")

	updated := &ixapi.CloudRouterNetworkServiceConfig{
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
		AdminStatus:      "disabled",
	}

	patchCalled := false
	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/20": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			if len(body) > 0 {
				patchCalled = true
				ixapi.AssertBodyContains(t, body, `"admin_status":"disabled"`)
			}
			return updated, nil
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	if err := cloudRouterConfigP2PVCUpdate(ctx, res, api); err != nil {
		t.Fatal(err)
	}

	if !patchCalled {
		t.Error("expected PATCH to be called")
	}
	if res.Get("admin_status").(string) != "disabled" {
		t.Error("unexpected admin_status after update:", res.Get("admin_status"))
	}
}

func TestCloudRouterConfigP2PVCDelete_404Success(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigP2PVCResource()
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
