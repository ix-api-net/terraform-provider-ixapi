package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ptr"
)

func TestCloudRouterPoliciesDataSourceRead(t *testing.T) {
	dataSource := NewDecixCloudRouterPoliciesDataSource()
	res := dataSource.TestResourceData()

	policies := []*ixapi.Policy{
		{
			ID:               "1",
			Name:             "accept-rfc1918",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			Entries: []ixapi.PolicyEntry{
				{
					SequenceNumber:  10,
					MatchPrefixList: ptr.Of("rfc1918-private"),
					Action: ixapi.PolicyAction{
						Filter:          ptr.Of("accept"),
						LocalPreference: ptr.Of(100),
					},
				},
				{
					SequenceNumber: 20,
					Action: ixapi.PolicyAction{
						Filter: ptr.Of("reject"),
					},
				},
			},
		},
		{
			ID:               "2",
			Name:             "reject-customer",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			Entries: []ixapi.PolicyEntry{
				{
					SequenceNumber:  10,
					MatchPrefixList: ptr.Of("customer-networks"),
					Action: ixapi.PolicyAction{
						Filter: ptr.Of("reject"),
						ASPathPrepend: &ixapi.ASPathPrepend{
							Count: 3,
						},
					},
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies": policies,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPoliciesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	pols := res.Get("policies").([]any)
	if len(pols) != 2 {
		t.Errorf("expected 2 policies, got %d", len(pols))
	}
}

func TestCloudRouterPoliciesDataSourceReadWithFilter(t *testing.T) {
	dataSource := NewDecixCloudRouterPoliciesDataSource()
	res := dataSource.TestResourceData()
	res.Set("managing_account", "100")

	policies := []*ixapi.Policy{
		{
			ID:               "1",
			Name:             "accept-rfc1918",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			Entries: []ixapi.PolicyEntry{
				{
					SequenceNumber: 10,
					Action: ixapi.PolicyAction{
						Filter:          ptr.Of("accept"),
						LocalPreference: ptr.Of(100),
					},
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies": policies,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPoliciesRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	pols := res.Get("policies").([]any)
	if len(pols) != 1 {
		t.Errorf("expected 1 policy, got %d", len(pols))
	}

	pol := pols[0].(map[string]any)
	if pol["managing_account"].(string) != "100" {
		t.Errorf("expected managing_account=100, got %v", pol["managing_account"])
	}
}


func TestCloudRouterPolicyDataSourceReadByID(t *testing.T) {
	dataSource := NewDecixCloudRouterPolicyDataSource()
	res := dataSource.TestResourceData()
	res.Set("id", "1")

	policy := &ixapi.Policy{
		ID:               "1",
		Name:             "accept-rfc1918",
		ManagingAccount:  "100",
		ConsumingAccount: "100",
		Entries: []ixapi.PolicyEntry{
			{
				SequenceNumber: 10,
				Action: ixapi.PolicyAction{
					Filter:          ptr.Of("accept"),
					LocalPreference: ptr.Of(100),
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies/1": policy,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPolicyRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "1" {
		t.Errorf("expected ID to be 1, got %s", res.Id())
	}
}

func TestCloudRouterPolicyDataSourceReadByName(t *testing.T) {
	dataSource := NewDecixCloudRouterPolicyDataSource()
	res := dataSource.TestResourceData()
	res.Set("name", "accept-rfc1918")
	res.Set("managing_account", "100")

	policies := []*ixapi.Policy{
		{
			ID:               "1",
			Name:             "accept-rfc1918",
			ManagingAccount:  "100",
			ConsumingAccount: "100",
			Entries: []ixapi.PolicyEntry{
				{
					SequenceNumber: 10,
					Action: ixapi.PolicyAction{
						Filter:          ptr.Of("accept"),
						LocalPreference: ptr.Of(100),
					},
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies": policies,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	diag := cloudRouterPolicyRead(ctx, res, api)
	if diag.HasError() {
		t.Fatal(diag)
	}

	if res.Id() != "1" {
		t.Errorf("expected ID to be 1, got %s", res.Id())
	}
}
