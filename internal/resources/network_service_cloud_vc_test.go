package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/testdata"
)

func TestCloudVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceCloudVCResource()
	res := resource.Data(nil)
	res.Set("managing_account", "managing:1")
	res.Set("consuming_account", "consuming:2")
	res.Set("cloud_key", "cloudkey")
	res.Set("capacity", 42000)

	req, err := nsCloudVCRequestFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	if req.Type != ixapi.CloudNetworkServiceType {
		t.Error("unexpected type in req:", req)
	}
	if req.ManagingAccount != "managing:1" {
		t.Error("unexpected request data:", req)
	}
	if req.ConsumingAccount != "consuming:2" {
		t.Error("unexpected request data:", req)
	}

	if req.CloudKey != "cloudkey" {
		t.Error("unexpected request data:", req)
	}
	if *req.Capacity != 42000 {
		t.Error("unexpected request data:", req)
	}
}

func TestCloudVCRead(t *testing.T) {
	resource := NewNetworkServiceCloudVCResource()
	res := resource.Data(nil)
	res.SetId("23")

	ns := testdata.NewCloudNetworkService()
	api := ixapi.NewTestClient(map[string]any{
		"/network-services/23": ns,
	})

	ctx := context.Background()
	err := nsCloudVCRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Get("cloud_key").(string) != "cloudkey" {
		t.Error("unexpected cloudkey in resource:", res)
	}
	if res.Get("capacity").(int) != 2300 {
		t.Error("unexpected capacity in resource:", res)
	}
}
