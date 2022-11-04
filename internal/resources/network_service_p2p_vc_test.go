package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestP2PVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceP2PVCResource()
	res := resource.Data(nil)
	res.Set("managing_account", "managing:1")
	res.Set("consuming_account", "consuming:2")
	res.Set("joining_member_account", "joining:3")
	res.Set("display_name", "p2p vc 1")

	req, err := nsP2PVCRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Type != ixapi.P2PNetworkServiceType {
		t.Error("unexpected type in req:", req)
	}
	if req.ManagingAccount != "managing:1" {
		t.Error("unexpected request data:", req)
	}
	if req.ConsumingAccount != "consuming:2" {
		t.Error("unexpected request data:", req)
	}
	if req.JoiningMemberAccount != "joining:3" {
		t.Error("unexpected request data:", req)
	}
	if *req.DisplayName != "p2p vc 1" {
		t.Error("unexpected request data:", req)
	}
}

func TestP2PVCRead(t *testing.T) {
	resource := NewNetworkServiceP2PVCResource()
	res := resource.Data(nil)
	res.SetId("23")

	ns := testdata.NewP2PNetworkService()
	api := ixapi.NewTestClient(map[string]any{
		"/network-services/23": ns,
	})

	ctx := context.Background()
	err := nsP2PVCRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("display_name").(string) != "p2p network service" {
		t.Error("unexpected display name in resource", res)
	}
	if res.Get("joining_member_account").(string) != "joining:123" {
		t.Error("unexpected joining member account in resource", res)
	}
}
