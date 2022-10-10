package resources

import (
	"context"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
)

func TestMaybeSetFQDN(t *testing.T) {
	resource := NewIPAllocationNetworkServiceConfigResource()
	res := resource.Data(nil)
	res.Set("fqdn", "gw1.example.ixp")

	ip := &ixapi.IPAddress{
		ID: "42",
	}
	api := ixapi.NewTestClient(map[string]any{
		"/ips/42": (ixapi.TestResponseFunc)(func(body []byte) (any, error) {
			ixapi.AssertBodyContains(t, body, `gw1.example.ixp`)
			return ip, nil
		}),
	})

	if err := maybeSetFQDN(
		context.Background(),
		res,
		api,
		ip); err != nil {
		t.Fatal(err)
	}

}
