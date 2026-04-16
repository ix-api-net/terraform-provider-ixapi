package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterConfigCloudVCRequestFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()
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
	res.Set("bfd_enabled", true)
	res.Set("cloud_vlan", 100)

	req, err := cloudRouterConfigCloudVCRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Type != "cloud_vc" {
		t.Error("unexpected type:", req.Type)
	}
	if req.ManagingAccount != "100" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if req.CloudRouter != "274" {
		t.Error("unexpected cloud_router:", req.CloudRouter)
	}
	if req.Address != "192.0.2.1/30" {
		t.Error("unexpected address:", req.Address)
	}
	if req.BGPNeighborASN != 64512 {
		t.Error("unexpected bgp_neighbor_asn:", req.BGPNeighborASN)
	}
	if req.AdminStatus != "enabled" {
		t.Error("unexpected admin_status:", req.AdminStatus)
	}
	if !req.BFDEnabled {
		t.Error("expected bfd_enabled to be true")
	}
	if req.CloudVLAN == nil || *req.CloudVLAN != 100 {
		t.Error("unexpected cloud_vlan:", req.CloudVLAN)
	}
	if req.Connection != nil {
		t.Error("connection should be nil for cloud_vc type")
	}
}

func TestCloudRouterConfigCloudVCRead(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()
	res := resource.Data(nil)
	res.SetId("10")

	cloudVLAN := 100
	config := &ixapi.CloudRouterNetworkServiceConfig{
		ID:               "10",
		Type:             "cloud_vc",
		ManagingAccount:  "100",
		BillingAccount:   "200",
		ConsumingAccount: "300",
		CloudRouter:      "274",
		NetworkService:   "500",
		Address:          "192.0.2.1/30",
		BGPNeighbor:      "192.0.2.2",
		BGPNeighborASN:   64512,
		AdminStatus:      "enabled",
		BFDEnabled:       true,
		CloudVLAN:        &cloudVLAN,
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/10": config,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigCloudVCRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("managing_account").(string) != "100" {
		t.Error("unexpected managing_account:", res.Get("managing_account"))
	}
	if res.Get("admin_status").(string) != "enabled" {
		t.Error("unexpected admin_status:", res.Get("admin_status"))
	}
	if res.Get("bfd_enabled").(bool) != true {
		t.Error("unexpected bfd_enabled:", res.Get("bfd_enabled"))
	}
	if res.Get("cloud_vlan").(int) != 100 {
		t.Error("unexpected cloud_vlan:", res.Get("cloud_vlan"))
	}
}

func TestCloudRouterConfigCloudVCCreate(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()
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

	created := &ixapi.CloudRouterNetworkServiceConfig{
		ID:               "10",
		Type:             "cloud_vc",
		ManagingAccount:  "100",
		BillingAccount:   "200",
		ConsumingAccount: "300",
		CloudRouter:      "274",
		NetworkService:   "500",
		Address:          "192.0.2.1/30",
		BGPNeighbor:      "192.0.2.2",
		BGPNeighborASN:   64512,
		AdminStatus:      "enabled",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return created, nil
		}),
		"/api/v3/decix-vrf-v1/network-service-configs/10": created,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigCloudVCCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "10" {
		t.Error("expected resource ID 10, got:", res.Id())
	}
	if res.Get("address").(string) != "192.0.2.1/30" {
		t.Error("unexpected address after create:", res.Get("address"))
	}
}

func TestCloudRouterConfigCloudVCPatchFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()

	t.Run("set fields are included", func(t *testing.T) {
		res := resource.Data(nil)
		res.Set("admin_status", "enabled")
		res.Set("policy_egress", "my-egress")

		patch := cloudRouterConfigCloudVCPatchFromResourceData(res)

		if patch.AdminStatus == nil || *patch.AdminStatus != "enabled" {
			t.Error("unexpected admin_status:", patch.AdminStatus)
		}
		if patch.PolicyEgress == nil || *patch.PolicyEgress != "my-egress" {
			t.Error("unexpected policy_egress:", patch.PolicyEgress)
		}
		if patch.PolicyIngress != nil {
			t.Error("policy_ingress should be nil when not set:", patch.PolicyIngress)
		}
	})

	t.Run("unset fields are nil", func(t *testing.T) {
		res := resource.Data(nil)

		patch := cloudRouterConfigCloudVCPatchFromResourceData(res)

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

func TestCloudRouterConfigCloudVCUpdate(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()
	res := resource.Data(nil)
	res.SetId("10")
	res.Set("managing_account", "100")
	res.Set("billing_account", "200")
	res.Set("consuming_account", "300")
	res.Set("cloud_router", "274")
	res.Set("network_service", "500")
	res.Set("address", "192.0.2.1/30")
	res.Set("bgp_neighbor", "192.0.2.2")
	res.Set("bgp_neighbor_asn", 64512)
	res.Set("admin_status", "disabled")

	updated := &ixapi.CloudRouterNetworkServiceConfig{
		ID:               "10",
		Type:             "cloud_vc",
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
		"/api/v3/decix-vrf-v1/network-service-configs/10": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			if len(body) > 0 {
				patchCalled = true
				ixapi.AssertBodyContains(t, body, `"admin_status":"disabled"`)
			}
			return updated, nil
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	if err := cloudRouterConfigCloudVCUpdate(ctx, res, api); err != nil {
		t.Fatal(err)
	}

	if !patchCalled {
		t.Error("expected PATCH to be called")
	}
	if res.Get("admin_status").(string) != "disabled" {
		t.Error("unexpected admin_status after update:", res.Get("admin_status"))
	}
}

func TestCloudRouterConfigCloudVCDelete_404Success(t *testing.T) {
	resource := NewDecixCloudRouterNetworkServiceConfigCloudVCResource()
	res := resource.Data(nil)
	res.SetId("10")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := cloudRouterConfigCloudVCDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}
	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}
