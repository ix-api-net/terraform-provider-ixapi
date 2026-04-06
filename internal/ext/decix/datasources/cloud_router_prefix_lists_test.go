package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCloudRouterPrefixListsDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterPrefixListsDataSource()
	res := dataSource.TestResourceData()

	prefixLists := []*ixapi.PrefixList{
		{
			ID:               "1",
			Name:             "rfc1918-private",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			MatchList: []ixapi.PrefixMatch{
				{
					Prefix:    "192.168.0.0/16",
					MaxLength: intPtr(24),
				},
				{
					Prefix:    "10.0.0.0/8",
					MinLength: intPtr(16),
					MaxLength: intPtr(24),
				},
			},
		},
		{
			ID:               "2",
			Name:             "customer-networks",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			MatchList: []ixapi.PrefixMatch{
				{
					Prefix:    "203.0.113.0/24",
					MinLength: intPtr(24),
					MaxLength: intPtr(32),
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists": prefixLists,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPrefixListsRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	pls := res.Get("prefix_lists").([]any)
	if len(pls) != 2 {
		t.Errorf("expected 2 prefix lists, got %d", len(pls))
	}
}

func TestCloudRouterPrefixListsDataSourceReadWithFilter(t *testing.T) {
	dataSource := NewDecixCloudRouterPrefixListsDataSource()
	res := dataSource.TestResourceData()
	res.Set("managing_account", "100")

	prefixLists := []*ixapi.PrefixList{
		{
			ID:               "1",
			Name:             "rfc1918-private",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			MatchList: []ixapi.PrefixMatch{
				{
					Prefix:    "192.168.0.0/16",
					MaxLength: intPtr(24),
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists": prefixLists,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPrefixListsRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	pls := res.Get("prefix_lists").([]any)
	if len(pls) != 1 {
		t.Errorf("expected 1 prefix list, got %d", len(pls))
	}

	pl := pls[0].(map[string]any)
	if pl["managing_account"].(string) != "100" {
		t.Errorf("expected managing_account=100, got %v", pl["managing_account"])
	}
}

func intPtr(i int) *int {
	return &i
}

func TestCloudRouterPrefixListDataSourceReadByID(t *testing.T) {
	dataSource := NewDecixCloudRouterPrefixListDataSource()
	res := dataSource.TestResourceData()
	res.Set("id", "1")

	prefixList := &ixapi.PrefixList{
		ID:               "1",
		Name:             "rfc1918-private",
		ManagingAccount:  "100",
		ConsumingAccount: "100",
		MatchList: []ixapi.PrefixMatch{
			{
				Prefix:    "192.168.0.0/16",
				MaxLength: intPtr(24),
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists/1": prefixList,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPrefixListRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "1" {
		t.Errorf("expected ID to be 1, got %s", res.Id())
	}
}

func TestCloudRouterPrefixListDataSourceReadByName(t *testing.T) {
	dataSource := NewDecixCloudRouterPrefixListDataSource()
	res := dataSource.TestResourceData()
	res.Set("name", "rfc1918-private")
	res.Set("managing_account", "100")

	prefixLists := []*ixapi.PrefixList{
		{
			ID:               "1",
			Name:             "rfc1918-private",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			MatchList: []ixapi.PrefixMatch{
				{
					Prefix:    "192.168.0.0/16",
					MaxLength: intPtr(24),
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists": prefixLists,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPrefixListRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "1" {
		t.Errorf("expected ID to be 1, got %s", res.Id())
	}
}
