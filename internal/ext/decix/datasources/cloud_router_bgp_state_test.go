package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterBGPStateRead(t *testing.T) {
	dataSource := NewDecixCloudRouterBGPStateDataSource()
	res := dataSource.TestResourceData()
	res.Set("nsc_id", "123")

	bgpState := &ixapi.BGPStateResponse{
		State: "Established",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/123/bgp-state": bgpState,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterBGPStateRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "123" {
		t.Errorf("expected ID to be nsc_id 123, got: %s", res.Id())
	}
	if res.Get("state").(string) != "Established" {
		t.Errorf("unexpected state: %v", res.Get("state"))
	}
}

func TestCloudRouterBGPStateRead_NotFound(t *testing.T) {
	dataSource := NewDecixCloudRouterBGPStateDataSource()
	res := dataSource.TestResourceData()
	res.Set("nsc_id", "456")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterBGPStateRead(ctx, res, api)

	if !diag.HasError() {
		t.Error("expected error when BGP state endpoint returns 404")
	}
}
