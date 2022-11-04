package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestP2MPVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceP2MPVCResource()
	res := resource.Data(nil)
	res.Set("managing_account", "managing:1")
	res.Set("consuming_account", "consuming:2")
	res.Set("display_name", "vc 1")
	res.Set("public", true)

	req, err := nsP2MPVCRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Type != ixapi.P2MPNetworkServiceType {
		t.Error("unexpected type in req:", req)
	}
	if req.ManagingAccount != "managing:1" {
		t.Error("unexpected request data:", req)
	}
	if req.ConsumingAccount != "consuming:2" {
		t.Error("unexpected request data:", req)
	}
	if *req.Public != true {
		t.Error("unexpected request data:", req)
	}
	if *req.DisplayName != "vc 1" {
		t.Error("unexpected request data:", req)
	}
}

func TestP2MPVCRead(t *testing.T) {
	resource := NewNetworkServiceP2MPVCResource()
	res := resource.Data(nil)
	res.SetId("23")

	ns := testdata.NewP2MPNetworkService()
	api := ixapi.NewTestClient(map[string]any{
		"/network-services/23": ns,
	})

	ctx := context.Background()
	err := nsP2MPVCRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("display_name").(string) != "p2mp network service" {
		t.Error("unexpected display name in resource", res)
	}
	if res.Get("public").(bool) != true {
		t.Error("expected public to be true")
	}
}
