package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestPolicyRequestFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	_ = res.Set("name", "accept-rfc1918")
	_ = res.Set("managing_account", "100")
	_ = res.Set("consuming_account", "200")
	_ = res.Set("entries", []interface{}{
		map[string]interface{}{
			"sequence_number":   10,
			"match_prefix_list": "rfc1918-private",
			"action": []interface{}{
				map[string]interface{}{
					"filter":           "accept",
					"local_preference": 100,
					"as_path_prepend":  []interface{}{},
				},
			},
		},
		map[string]interface{}{
			"sequence_number":   20,
			"match_prefix_list": "",
			"action": []interface{}{
				map[string]interface{}{
					"filter":           "reject",
					"local_preference": 0,
					"as_path_prepend":  []interface{}{},
				},
			},
		},
	})

	req, err := policyRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Name != "accept-rfc1918" {
		t.Error("unexpected name:", req.Name)
	}
	if req.ManagingAccount != "100" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if len(req.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(req.Entries))
	}
	if req.Entries[0].SequenceNumber != 10 {
		t.Error("unexpected sequence_number:", req.Entries[0].SequenceNumber)
	}
	if req.Entries[0].MatchPrefixList == nil || *req.Entries[0].MatchPrefixList != "rfc1918-private" {
		t.Error("unexpected match_prefix_list:", req.Entries[0].MatchPrefixList)
	}
	if req.Entries[0].Action.Filter == nil || *req.Entries[0].Action.Filter != "accept" {
		t.Error("unexpected filter:", req.Entries[0].Action.Filter)
	}
	if req.Entries[0].Action.LocalPreference == nil || *req.Entries[0].Action.LocalPreference != 100 {
		t.Error("unexpected local_preference:", req.Entries[0].Action.LocalPreference)
	}
	if req.Entries[1].MatchPrefixList != nil {
		t.Error("match_prefix_list should be nil when empty, got:", req.Entries[1].MatchPrefixList)
	}
}

func TestPolicyRequestFromResourceData_WithASPathPrepend(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	_ = res.Set("name", "prepend-policy")
	_ = res.Set("managing_account", "100")
	_ = res.Set("consuming_account", "200")
	_ = res.Set("entries", []interface{}{
		map[string]interface{}{
			"sequence_number":   10,
			"match_prefix_list": "",
			"action": []interface{}{
				map[string]interface{}{
					"filter":           "",
					"local_preference": 0,
					"as_path_prepend": []interface{}{
						map[string]interface{}{
							"count": 3,
							"asn":   65001,
						},
					},
				},
			},
		},
	})

	req, err := policyRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(req.Entries))
	}
	if req.Entries[0].Action.ASPathPrepend == nil {
		t.Fatal("expected as_path_prepend to be set")
	}
	if req.Entries[0].Action.ASPathPrepend.Count != 3 {
		t.Error("unexpected count:", req.Entries[0].Action.ASPathPrepend.Count)
	}
	if req.Entries[0].Action.ASPathPrepend.ASN == nil || *req.Entries[0].Action.ASPathPrepend.ASN != 65001 {
		t.Error("unexpected asn:", req.Entries[0].Action.ASPathPrepend.ASN)
	}
}

func TestPolicyRead(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	res.SetId("1")

	filterStr := "accept"
	localPref := 100
	policy := &ixapi.CloudRouterPolicy{
		ID:               "1",
		Name:             "accept-rfc1918",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		Entries: []ixapi.CloudRouterPolicyEntry{
			{
				SequenceNumber: 10,
				Action: ixapi.CloudRouterPolicyAction{
					Filter:          &filterStr,
					LocalPreference: &localPref,
				},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies/1": policy,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := policyRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "accept-rfc1918" {
		t.Error("unexpected name:", res.Get("name"))
	}
	entries := res.Get("entries").([]interface{})
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
}

func TestPolicyCreate(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	_ = res.Set("name", "accept-rfc1918")
	_ = res.Set("managing_account", "100")
	_ = res.Set("consuming_account", "200")
	_ = res.Set("entries", []interface{}{
		map[string]interface{}{
			"sequence_number":   10,
			"match_prefix_list": "",
			"action": []interface{}{
				map[string]interface{}{
					"filter":           "accept",
					"local_preference": 0,
					"as_path_prepend":  []interface{}{},
				},
			},
		},
	})

	filterStr := "accept"
	created := &ixapi.CloudRouterPolicy{
		ID:               "55",
		Name:             "accept-rfc1918",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		Entries: []ixapi.CloudRouterPolicyEntry{
			{
				SequenceNumber: 10,
				Action:         ixapi.CloudRouterPolicyAction{Filter: &filterStr},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return created, nil
		}),
		"/api/v3/decix-vrf-v1/policies/55": created,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := policyCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "55" {
		t.Error("expected resource ID 55, got:", res.Id())
	}
	if res.Get("name").(string) != "accept-rfc1918" {
		t.Error("unexpected name after create:", res.Get("name"))
	}
}

func TestPolicyUpdate(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	res.SetId("55")
	_ = res.Set("name", "updated-policy")
	_ = res.Set("managing_account", "100")
	_ = res.Set("consuming_account", "200")
	_ = res.Set("entries", []interface{}{
		map[string]interface{}{
			"sequence_number":   10,
			"match_prefix_list": "",
			"action": []interface{}{
				map[string]interface{}{
					"filter":           "reject",
					"local_preference": 0,
					"as_path_prepend":  []interface{}{},
				},
			},
		},
	})

	filterStr := "reject"
	updated := &ixapi.CloudRouterPolicy{
		ID:               "55",
		Name:             "updated-policy",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		Entries: []ixapi.CloudRouterPolicyEntry{
			{
				SequenceNumber: 10,
				Action:         ixapi.CloudRouterPolicyAction{Filter: &filterStr},
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies/55": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return updated, nil
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := policyUpdate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "updated-policy" {
		t.Error("unexpected name after update:", res.Get("name"))
	}
}

func TestPolicyDelete_404Success(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	res.SetId("55")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := policyDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}
	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}

func TestPolicyDelete_OtherErrorFails(t *testing.T) {
	resource := NewDecixCloudRouterPolicyResource()
	res := resource.Data(nil)
	res.SetId("55")

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/policies/55": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return nil, &ixapi.APIError{
				ProblemResponse: ixapi.ProblemResponse{
					Type:   "internal_error",
					Title:  "Internal Server Error",
					Status: 500,
				},
			}
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := policyDelete(ctx, res, api)

	if err == nil {
		t.Error("delete should fail on non-404 errors")
	}
}
