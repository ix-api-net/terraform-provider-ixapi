package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestStaticRouteRequestFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterStaticRouteResource()
	res := resource.Data(nil)
	_ = res.Set("name", "test-route")
	_ = res.Set("prefix", "10.0.0.0/24")
	_ = res.Set("next_hop", "192.168.1.1")
	_ = res.Set("network_service_configs", []interface{}{"nsc-1", "nsc-2"})

	req, err := staticRouteRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Name != "test-route" {
		t.Error("unexpected name:", req.Name)
	}
	if req.Prefix != "10.0.0.0/24" {
		t.Error("unexpected prefix:", req.Prefix)
	}
	if req.NextHop != "192.168.1.1" {
		t.Error("unexpected next_hop:", req.NextHop)
	}
	if len(req.NetworkServiceConfigs) != 2 {
		t.Fatalf("expected 2 NSCs, got %d", len(req.NetworkServiceConfigs))
	}
	if req.NetworkServiceConfigs[0] != "nsc-1" {
		t.Error("unexpected first NSC:", req.NetworkServiceConfigs[0])
	}
}

func TestStaticRouteRead(t *testing.T) {
	resource := NewDecixCloudRouterStaticRouteResource()
	res := resource.Data(nil)
	res.SetId("route-1")

	route := &ixapi.CloudRouterStaticRoute{
		ID:                    "route-1",
		Name:                  "my-route",
		Prefix:                "10.0.0.0/24",
		NextHop:               "192.168.1.1",
		NetworkServiceConfigs: []string{"nsc-abc"},
		VRF:                   "vrf-xyz",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/static-routes/route-1": route,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := staticRouteRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "my-route" {
		t.Error("unexpected name:", res.Get("name"))
	}
	if res.Get("prefix").(string) != "10.0.0.0/24" {
		t.Error("unexpected prefix:", res.Get("prefix"))
	}
	if res.Get("next_hop").(string) != "192.168.1.1" {
		t.Error("unexpected next_hop:", res.Get("next_hop"))
	}
	if res.Get("vrf").(string) != "vrf-xyz" {
		t.Error("unexpected vrf:", res.Get("vrf"))
	}
	nscs := res.Get("network_service_configs").([]interface{})
	if len(nscs) != 1 || nscs[0].(string) != "nsc-abc" {
		t.Error("unexpected network_service_configs:", nscs)
	}
}

func TestStaticRouteCreate(t *testing.T) {
	resource := NewDecixCloudRouterStaticRouteResource()
	res := resource.Data(nil)
	_ = res.Set("name", "my-route")
	_ = res.Set("prefix", "10.0.0.0/24")
	_ = res.Set("next_hop", "aggregate")
	_ = res.Set("network_service_configs", []interface{}{"nsc-abc"})

	created := &ixapi.CloudRouterStaticRoute{
		ID:                    "route-42",
		Name:                  "my-route",
		Prefix:                "10.0.0.0/24",
		NextHop:               "aggregate",
		NetworkServiceConfigs: []string{"nsc-abc"},
		VRF:                   "vrf-xyz",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/static-routes": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return created, nil
		}),
		"/api/v3/decix-vrf-v1/static-routes/route-42": created,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := staticRouteCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "route-42" {
		t.Error("expected resource ID route-42, got:", res.Id())
	}
	if res.Get("name").(string) != "my-route" {
		t.Error("unexpected name after create:", res.Get("name"))
	}
}

func TestStaticRouteUpdate(t *testing.T) {
	resource := NewDecixCloudRouterStaticRouteResource()
	res := resource.Data(nil)
	res.SetId("route-42")
	_ = res.Set("name", "updated-route")
	_ = res.Set("prefix", "172.16.0.0/16")
	_ = res.Set("next_hop", "10.0.0.1")
	_ = res.Set("network_service_configs", []interface{}{"nsc-abc"})

	updated := &ixapi.CloudRouterStaticRoute{
		ID:                    "route-42",
		Name:                  "updated-route",
		Prefix:                "172.16.0.0/16",
		NextHop:               "10.0.0.1",
		NetworkServiceConfigs: []string{"nsc-abc"},
		VRF:                   "vrf-xyz",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/static-routes/route-42": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return updated, nil
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := staticRouteUpdate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "updated-route" {
		t.Error("unexpected name after update:", res.Get("name"))
	}
	if res.Get("prefix").(string) != "172.16.0.0/16" {
		t.Error("unexpected prefix after update:", res.Get("prefix"))
	}
}

func TestStaticRouteDelete_404Success(t *testing.T) {
	resource := NewDecixCloudRouterStaticRouteResource()
	res := resource.Data(nil)
	res.SetId("route-42")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := staticRouteDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}
	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}
