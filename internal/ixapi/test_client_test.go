package ixapi

import (
	"context"
	"testing"
)

func TestMockResponseClient(t *testing.T) {
	ctx := context.Background()

	c := NewTestClient(map[string]any{
		"/network-services/1": &ExchangeLanNetworkService{
			ID:   "1",
			Type: ExchangeLanNetworkServiceType,
		},
	})

	res, err := c.NetworkServicesRead(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := res.(*ExchangeLanNetworkService); !ok {
		t.Fatalf("expected *ExchangeLanNetworkService, got %T", res)
	}

	_, err = c.NetworkServicesList(ctx)
	if err == nil {
		t.Fatal("endpoint should not be present")
	}
	if _, ok := err.(*NotFoundError); !ok {
		t.Error("response error should be 404 not found")
	}
}

func TestMockResponseClientHandlerFunc(t *testing.T) {
	ctx := context.Background()

	c := NewTestClient(map[string]any{
		"/network-services/1": (TestResponseFunc)(func(body []byte) (any, error) {
			return &ExchangeLanNetworkService{
				ID:   "1",
				Type: ExchangeLanNetworkServiceType,
			}, nil
		}),
	})

	res, err := c.NetworkServicesRead(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := res.(*ExchangeLanNetworkService); !ok {
		t.Fatalf("expected *ExchangeLanNetworkService, got %T", res)
	}

}
