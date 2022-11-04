package resources

import "testing"

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
