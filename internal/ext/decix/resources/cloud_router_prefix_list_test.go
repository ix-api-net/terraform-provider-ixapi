package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestPrefixListRequestFromResourceData(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.Set("name", "test-list")
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("match_list", []interface{}{
		map[string]interface{}{
			"prefix":     "192.168.0.0/16",
			"min_length": 0,
			"max_length": 24,
		},
		map[string]interface{}{
			"prefix":     "10.0.0.0/8",
			"min_length": 16,
			"max_length": 24,
		},
	})

	req, err := prefixListRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Name != "test-list" {
		t.Error("unexpected name:", req.Name)
	}
	if req.ManagingAccount != "100" {
		t.Error("unexpected managing_account:", req.ManagingAccount)
	}
	if req.ConsumingAccount != "200" {
		t.Error("unexpected consuming_account:", req.ConsumingAccount)
	}
	if len(req.MatchList) != 2 {
		t.Fatalf("expected 2 match entries, got %d", len(req.MatchList))
	}
	if req.MatchList[0].Prefix != "192.168.0.0/16" {
		t.Error("unexpected first prefix:", req.MatchList[0].Prefix)
	}
	if req.MatchList[0].MaxLength == nil || *req.MatchList[0].MaxLength != 24 {
		t.Error("unexpected first max_length:", req.MatchList[0].MaxLength)
	}
	if req.MatchList[0].MinLength != nil {
		t.Error("min_length should be nil when 0, got:", req.MatchList[0].MinLength)
	}
	if req.MatchList[1].MinLength == nil || *req.MatchList[1].MinLength != 16 {
		t.Error("unexpected second min_length:", req.MatchList[1].MinLength)
	}
}

func TestPrefixListRead(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.SetId("1")

	pl := &ixapi.CloudRouterPrefixList{
		ID:               "1",
		Name:             "rfc1918-private",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		MatchList: []ixapi.CloudRouterPrefixMatch{
			{
				Prefix:    "192.168.0.0/16",
				MaxLength: new(24),
			},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists/1": pl,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := prefixListRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "rfc1918-private" {
		t.Error("unexpected name:", res.Get("name"))
	}
	if res.Get("managing_account").(string) != "100" {
		t.Error("unexpected managing_account:", res.Get("managing_account"))
	}
	matchList := res.Get("match_list").([]interface{})
	if len(matchList) != 1 {
		t.Fatalf("expected 1 match entry, got %d", len(matchList))
	}
}

func TestPrefixListCreate(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.Set("name", "rfc1918-private")
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("match_list", []interface{}{
		map[string]interface{}{
			"prefix":     "192.168.0.0/16",
			"min_length": 0,
			"max_length": 0,
		},
	})

	created := &ixapi.CloudRouterPrefixList{
		ID:               "42",
		Name:             "rfc1918-private",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		MatchList: []ixapi.CloudRouterPrefixMatch{
			{Prefix: "192.168.0.0/16"},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return created, nil
		}),
		"/api/v3/decix-vrf-v1/prefix-lists/42": created,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := prefixListCreate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "42" {
		t.Error("expected resource ID 42, got:", res.Id())
	}
	if res.Get("name").(string) != "rfc1918-private" {
		t.Error("unexpected name after create:", res.Get("name"))
	}
}

func TestPrefixListUpdate(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.SetId("42")
	res.Set("name", "updated-list")
	res.Set("managing_account", "100")
	res.Set("consuming_account", "200")
	res.Set("match_list", []interface{}{
		map[string]interface{}{
			"prefix":     "10.0.0.0/8",
			"min_length": 0,
			"max_length": 0,
		},
	})

	updated := &ixapi.CloudRouterPrefixList{
		ID:               "42",
		Name:             "updated-list",
		ManagingAccount:  "100",
		ConsumingAccount: "200",
		MatchList: []ixapi.CloudRouterPrefixMatch{
			{Prefix: "10.0.0.0/8"},
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists/42": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			return updated, nil
		}),
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := prefixListUpdate(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("name").(string) != "updated-list" {
		t.Error("unexpected name after update:", res.Get("name"))
	}
}

func TestPrefixListDelete_404Success(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.SetId("42")

	api := ixapi.NewTestClient(map[string]any{})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := prefixListDelete(ctx, res, api)

	if err != nil {
		t.Error("delete should succeed on 404, got error:", err)
	}
	if res.Id() != "" {
		t.Error("resource ID should be cleared after 404")
	}
}

func TestPrefixListDelete_OtherErrorFails(t *testing.T) {
	resource := NewDecixCloudRouterPrefixListResource()
	res := resource.Data(nil)
	res.SetId("42")

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/prefix-lists/42": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
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
	err := prefixListDelete(ctx, res, api)

	if err == nil {
		t.Error("delete should fail on non-404 errors")
	}
}
