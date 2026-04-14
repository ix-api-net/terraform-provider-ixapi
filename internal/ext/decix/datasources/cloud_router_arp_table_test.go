package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterArpTableDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterArpTableDataSource()
	res := dataSource.TestResourceData()

	entries := []*ixapi.CloudRouterArpEntry{
		{
			VRF:                  "vrf-1",
			NetworkServiceConfig: "nsc-abc",
			DeviceFQDN:           "router1.example.com",
			IPAddress:            "192.168.1.1",
			MACAddress:           "00:11:22:33:44:55",
			ExpirationTime:       3600,
			ReceivedAt:           "2024-03-20T20:45:59Z",
		},
		{
			VRF:                  "vrf-1",
			NetworkServiceConfig: "nsc-abc",
			DeviceFQDN:           "router1.example.com",
			IPAddress:            "192.168.1.2",
			MACAddress:           "00:11:22:33:44:66",
			ExpirationTime:       3600,
			ReceivedAt:           "2024-03-20T20:45:59Z",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/arp-table": entries,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterArpTableRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	result := res.Get("arp_entries").([]any)
	if len(result) != 2 {
		t.Errorf("expected 2 ARP entries, got %d", len(result))
	}
}

func TestCloudRouterArpTableDataSourceReadEmpty(t *testing.T) {
	dataSource := NewDecixCloudRouterArpTableDataSource()
	res := dataSource.TestResourceData()
	res.Set("vrf", "vrf-1")

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/arp-table": []*ixapi.CloudRouterArpEntry{},
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterArpTableRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	result := res.Get("arp_entries").([]any)
	if len(result) != 0 {
		t.Errorf("expected 0 ARP entries, got %d", len(result))
	}
}
