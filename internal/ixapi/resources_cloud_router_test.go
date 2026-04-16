package ixapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

)

func TestCloudRoutersCreate(t *testing.T) {
	externalRef := "test-vrf"
	expectedReq := &CloudRouterRequest{
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		BillingAccount:   "acc-1",
		ProductOffering:  "po-123",
		ASN:              65893,
		Capacity:         1000,
		ExternalRef:      &externalRef,
	}

	expectedResp := &CloudRouter{
		ID:               "vrf-1",
		State:            "active",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		BillingAccount:   "acc-1",
		ExternalRef:      &externalRef,
		ProductOffering:  "po-123",
		ASN:              65893,
		Capacity:         1000,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/vrfs" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/vrfs path, got %s", r.URL.Path)
		}

		var req CloudRouterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.ASN != expectedReq.ASN {
			t.Errorf("Expected ASN %d, got %d", expectedReq.ASN, req.ASN)
		}
		if req.Capacity != expectedReq.Capacity {
			t.Errorf("Expected capacity %d, got %d", expectedReq.Capacity, req.Capacity)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRoutersCreate(ctx, expectedReq)
	if err != nil {
		t.Fatalf("CloudRoutersCreate failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
	if result.ASN != expectedResp.ASN {
		t.Errorf("Expected ASN %d, got %d", expectedResp.ASN, result.ASN)
	}
}

func TestCloudRoutersRead(t *testing.T) {
	externalRef := "test-vrf"
	expectedResp := &CloudRouter{
		ID:               "vrf-1",
		State:            "active",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		BillingAccount:   "acc-1",
		ExternalRef:      &externalRef,
		ProductOffering:  "po-123",
		ASN:              65893,
		Capacity:         1000,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/vrfs/vrf-1" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/vrfs/vrf-1 path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRoutersRead(ctx, "vrf-1")
	if err != nil {
		t.Fatalf("CloudRoutersRead failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
	if result.State != expectedResp.State {
		t.Errorf("Expected state %s, got %s", expectedResp.State, result.State)
	}
}

func TestCloudRoutersDestroy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/vrfs/vrf-1" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/vrfs/vrf-1 path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	err := client.DecixCloudRoutersDestroy(ctx, "vrf-1")
	if err != nil {
		t.Fatalf("CloudRoutersDestroy failed: %v", err)
	}
}

func TestCloudRoutersList(t *testing.T) {
	externalRef := "test-vrf"
	expectedResp := []*CloudRouter{
		{
			ID:               "vrf-1",
			State:            "active",
			ManagingAccount:  "acc-1",
			ConsumingAccount: "acc-1",
			ExternalRef:      &externalRef,
			ASN:              65893,
		},
		{
			ID:               "vrf-2",
			State:            "active",
			ManagingAccount:  "acc-1",
			ConsumingAccount: "acc-1",
			ASN:              65894,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/vrfs" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/vrfs path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRoutersList(ctx)
	if err != nil {
		t.Fatalf("CloudRoutersList failed: %v", err)
	}

	if len(result) != len(expectedResp) {
		t.Errorf("Expected %d routers, got %d", len(expectedResp), len(result))
	}
	if result[0].ID != expectedResp[0].ID {
		t.Errorf("Expected first router ID %s, got %s", expectedResp[0].ID, result[0].ID)
	}
}

func TestCloudRoutersListWithQuery(t *testing.T) {
	expectedResp := []*CloudRouter{
		{
			ID:               "vrf-1",
			State:            "active",
			ManagingAccount:  "acc-1",
			ConsumingAccount: "acc-1",
			ASN:              65893,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/vrfs" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/vrfs path, got %s", r.URL.Path)
		}

		query := r.URL.Query()
		if query.Get("managing_account") != "acc-1" {
			t.Errorf("Expected managing_account=acc-1, got %s", query.Get("managing_account"))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	qry := &CloudRoutersListQuery{
		ManagingAccount: "acc-1",
	}
	result, err := client.DecixCloudRoutersList(ctx, qry)
	if err != nil {
		t.Fatalf("CloudRoutersList failed: %v", err)
	}

	if len(result) != len(expectedResp) {
		t.Errorf("Expected %d routers, got %d", len(expectedResp), len(result))
	}
}

func TestPrefixListsCreate(t *testing.T) {
	expectedReq := &CloudRouterPrefixListRequest{
		Name:             "test-prefix-list",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		MatchList: []CloudRouterPrefixMatch{
			{Prefix: "192.168.0.0/16", MaxLength: new(24)},
			{Prefix: "10.0.0.0/8", MinLength: new(16), MaxLength: new(24)},
		},
	}

	expectedResp := &CloudRouterPrefixList{
		ID:               "pl-1",
		Name:             "test-prefix-list",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		MatchList: []CloudRouterPrefixMatch{
			{Prefix: "192.168.0.0/16", MaxLength: new(24)},
			{Prefix: "10.0.0.0/8", MinLength: new(16), MaxLength: new(24)},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/prefix-lists" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/prefix-lists path, got %s", r.URL.Path)
		}

		var req CloudRouterPrefixListRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.Name != expectedReq.Name {
			t.Errorf("Expected name %s, got %s", expectedReq.Name, req.Name)
		}
		if len(req.MatchList) != len(expectedReq.MatchList) {
			t.Errorf("Expected %d matches, got %d", len(expectedReq.MatchList), len(req.MatchList))
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRouterPrefixListsCreate(ctx, expectedReq)
	if err != nil {
		t.Fatalf("PrefixListsCreate failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
	if result.Name != expectedResp.Name {
		t.Errorf("Expected name %s, got %s", expectedResp.Name, result.Name)
	}
}

func TestPrefixListsRead(t *testing.T) {
	expectedResp := &CloudRouterPrefixList{
		ID:               "pl-1",
		Name:             "test-prefix-list",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		MatchList: []CloudRouterPrefixMatch{
			{Prefix: "192.168.0.0/16", MaxLength: new(24)},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/prefix-lists/pl-1" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/prefix-lists/pl-1 path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRouterPrefixListsRead(ctx, "pl-1")
	if err != nil {
		t.Fatalf("PrefixListsRead failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
}

func TestPrefixListsDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/prefix-lists/pl-1" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/prefix-lists/pl-1 path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(&CloudRouterPrefixList{ID: "pl-1"})
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	_, err := client.DecixCloudRouterPrefixListsDelete(ctx, "pl-1")
	if err != nil {
		t.Fatalf("PrefixListsDelete failed: %v", err)
	}
}

func TestPoliciesCreate(t *testing.T) {
	localPref := 100
	count := 3
	expectedReq := &CloudRouterPolicyRequest{
		Name:             "test-policy",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		Entries: []CloudRouterPolicyEntry{
			{
				SequenceNumber:  10,
				MatchPrefixList: new("test-list"),
				Action: CloudRouterPolicyAction{
					Filter:          new("accept"),
					LocalPreference: &localPref,
					ASPathPrepend: &CloudRouterASPathPrepend{
						Count: count,
					},
				},
			},
		},
	}

	expectedResp := &CloudRouterPolicy{
		ID:               "pol-1",
		Name:             "test-policy",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		Entries: []CloudRouterPolicyEntry{
			{
				SequenceNumber:  10,
				MatchPrefixList: new("test-list"),
				Action: CloudRouterPolicyAction{
					Filter:          new("accept"),
					LocalPreference: &localPref,
					ASPathPrepend: &CloudRouterASPathPrepend{
						Count: count,
					},
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/policies" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/policies path, got %s", r.URL.Path)
		}

		var req CloudRouterPolicyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request: %v", err)
		}

		if req.Name != expectedReq.Name {
			t.Errorf("Expected name %s, got %s", expectedReq.Name, req.Name)
		}
		if len(req.Entries) != len(expectedReq.Entries) {
			t.Errorf("Expected %d entries, got %d", len(expectedReq.Entries), len(req.Entries))
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRouterPoliciesCreate(ctx, expectedReq)
	if err != nil {
		t.Fatalf("PoliciesCreate failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
	if result.Name != expectedResp.Name {
		t.Errorf("Expected name %s, got %s", expectedResp.Name, result.Name)
	}
}

func TestPoliciesRead(t *testing.T) {
	localPref := 100
	expectedResp := &CloudRouterPolicy{
		ID:               "pol-1",
		Name:             "test-policy",
		ManagingAccount:  "acc-1",
		ConsumingAccount: "acc-1",
		Entries: []CloudRouterPolicyEntry{
			{
				SequenceNumber:  10,
				MatchPrefixList: new("test-list"),
				Action: CloudRouterPolicyAction{
					Filter:          new("accept"),
					LocalPreference: &localPref,
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/api/v3/decix-vrf-v1/policies/pol-1" {
			t.Errorf("Expected /api/v3/decix-vrf-v1/policies/pol-1 path, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client := NewClient(server.URL)
	ctx := context.Background()

	result, err := client.DecixCloudRouterPoliciesRead(ctx, "pol-1")
	if err != nil {
		t.Fatalf("PoliciesRead failed: %v", err)
	}

	if result.ID != expectedResp.ID {
		t.Errorf("Expected ID %s, got %s", expectedResp.ID, result.ID)
	}
}
