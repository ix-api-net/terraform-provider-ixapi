package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterBFDStateRead(t *testing.T) {
	dataSource := NewDecixCloudRouterBFDStateDataSource()
	res := dataSource.TestResourceData()
	res.Set("nsc_id", "123")

	bfdState := &ixapi.BFDStateResponse{
		State: "Up",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/network-service-configs/123/bfd-state": bfdState,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterBFDStateRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "123" {
		t.Errorf("expected ID to be nsc_id 123, got: %s", res.Id())
	}
	if res.Get("state").(string) != "Up" {
		t.Errorf("unexpected state: %v", res.Get("state"))
	}
}

func TestCloudRouterBFDStateRead_NotFound(t *testing.T) {
	dataSource := NewDecixCloudRouterBFDStateDataSource()
	res := dataSource.TestResourceData()
	res.Set("nsc_id", "456")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterBFDStateRead(ctx, res, api)

	if !diag.HasError() {
		t.Error("expected error when BFD state endpoint returns 404")
	}
}
