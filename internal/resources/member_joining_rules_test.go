package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestMJRAllowRequestFromResourceData(t *testing.T) {
	resource := NewMemberJoiningRuleAllowResource()
	res := resource.Data(nil)
	res.Set("network_service", "ns:42")
	res.Set("managing_account", "acc:23")
	res.Set("consuming_account", "acc:42")
	res.Set("capacity_max", 42000)

	req, err := mjrAllowRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if *req.CapacityMax != 42000 {
		t.Error("unexpected request data:", req)
	}
	if req.ConsumingAccount != "acc:42" {
		t.Error("unexpected request data:", req)
	}
}

func TestMJRAllowRead(t *testing.T) {
	resource := NewMemberJoiningRuleAllowResource()
	res := resource.Data(nil)
	res.SetId("23")

	rule := testdata.NewMemberJoiningRuleAllow()
	api := ixapi.NewTestClient(map[string]any{
		"/member-joining-rules/23": rule,
	})

	ctx := context.Background()
	if err := mjrAllowRead(ctx, res, api); err != nil {
		t.Fatal(err)
	}

	// Check resource data
	if res.Get("managing_account").(string) != "managing:123" {
		t.Error("unexpected resource data:", res)
	}
	if res.Get("network_service").(string) != "ns:23" {
		t.Error("unexpected resource data:", res)
	}
}

func TestMJRDenyRequestFromResourceData(t *testing.T) {
	resource := NewMemberJoiningRuleDenyResource()
	res := resource.Data(nil)
	res.Set("network_service", "ns:42")
	res.Set("managing_account", "acc:23")
	res.Set("consuming_account", "acc:42")

	req, err := mjrDenyRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.NetworkService != "ns:42" {
		t.Error("unexpected request data:", req)
	}
	if req.ConsumingAccount != "acc:42" {
		t.Error("unexpected request data:", req)
	}
}

func TestMJRDenyRead(t *testing.T) {
	resource := NewMemberJoiningRuleDenyResource()
	res := resource.Data(nil)
	res.SetId("23")

	rule := testdata.NewMemberJoiningRuleDeny()
	api := ixapi.NewTestClient(map[string]any{
		"/member-joining-rules/23": rule,
	})

	ctx := context.Background()
	if err := mjrDenyRead(ctx, res, api); err != nil {
		t.Fatal(err)
	}

	// Check resource data
	if res.Get("managing_account").(string) != "managing:123" {
		t.Error("unexpected resource data:", res)
	}
	if res.Get("network_service").(string) != "ns:23" {
		t.Error("unexpected resource data:", res)
	}
}
