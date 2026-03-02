package datasources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestProductOfferingsCloudVRFRead(t *testing.T) {
	dataSource := NewProductOfferingsCloudVRFDataSource()
	res := dataSource.TestResourceData()

	offerings := []*ixapi.CloudRouterProductOffering{
		{
			ID:                      "po-1",
			Name:                    "cloud-router-1g",
			DisplayName:             "Cloud Router 1G",
			BandwidthMax:            1000,
			BandwidthMin:            100,
			ServiceMetroArea:        "de-fra",
			ServiceMetroAreaName:    "Frankfurt",
			ServiceMetroAreaNetwork: "de-fra-1",
			ContractPeriod:          "P1M",
			Type:                    "cloud_router",
		},
		{
			ID:                      "po-2",
			Name:                    "cloud-router-10g",
			DisplayName:             "Cloud Router 10G",
			BandwidthMax:            10000,
			BandwidthMin:            1000,
			ServiceMetroArea:        "de-fra",
			ServiceMetroAreaName:    "Frankfurt",
			ServiceMetroAreaNetwork: "de-fra-1",
			ContractPeriod:          "P1Y",
			Type:                    "cloud_router",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/product-offerings": offerings,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := productOfferingsCloudVRFRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	pos := res.Get("product_offerings").([]any)
	if len(pos) != 2 {
		t.Errorf("expected 2 product offerings, got %d", len(pos))
	}
}

func TestProductOfferingsCloudVRFReadWithFilter(t *testing.T) {
	dataSource := NewProductOfferingsCloudVRFDataSource()
	res := dataSource.TestResourceData()
	res.Set("contract_period", "P1Y")

	offerings := []*ixapi.CloudRouterProductOffering{
		{
			ID:             "po-2",
			Name:           "cloud-router-10g",
			ContractPeriod: "P1Y",
			Type:           "cloud_router",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/product-offerings": offerings,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := productOfferingsCloudVRFRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	pos := res.Get("product_offerings").([]any)
	if len(pos) != 1 {
		t.Errorf("expected 1 product offering, got %d", len(pos))
	}

	po := pos[0].(map[string]any)
	if po["contract_period"].(string) != "P1Y" {
		t.Errorf("expected contract_period=P1Y, got %v", po["contract_period"])
	}
}

func TestProductOfferingCloudVRFReadByID(t *testing.T) {
	dataSource := NewProductOfferingCloudVRFDataSource()
	res := dataSource.TestResourceData()
	res.Set("id", "po-1")

	offering := &ixapi.CloudRouterProductOffering{
		ID:             "po-1",
		Name:           "cloud-router-1g",
		DisplayName:    "Cloud Router 1G",
		BandwidthMax:   1000,
		BandwidthMin:   100,
		ContractPeriod: "P1M",
		Type:           "cloud_router",
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/product-offerings/po-1": offering,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := productOfferingCloudVRFRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "po-1" {
		t.Errorf("expected ID to be po-1, got %s", res.Id())
	}
	if res.Get("name").(string) != "cloud-router-1g" {
		t.Errorf("unexpected name: %v", res.Get("name"))
	}
}

func TestProductOfferingCloudVRFRead_UniqueFilter(t *testing.T) {
	dataSource := NewProductOfferingCloudVRFDataSource()
	res := dataSource.TestResourceData()
	res.Set("name", "cloud-router-1g")

	offerings := []*ixapi.CloudRouterProductOffering{
		{
			ID:             "po-1",
			Name:           "cloud-router-1g",
			ContractPeriod: "P1M",
			Type:           "cloud_router",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/product-offerings": offerings,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := productOfferingCloudVRFRead(ctx, res, api)
	if err != nil {
		t.Fatal(err)
	}

	if res.Id() != "po-1" {
		t.Errorf("expected ID to be po-1, got %s", res.Id())
	}
}

func TestProductOfferingCloudVRFRead_AmbiguousFilter(t *testing.T) {
	dataSource := NewProductOfferingCloudVRFDataSource()
	res := dataSource.TestResourceData()
	res.Set("service_metro_area", "de-fra")

	offerings := []*ixapi.CloudRouterProductOffering{
		{
			ID:               "po-1",
			Name:             "cloud-router-1g",
			ServiceMetroArea: "de-fra",
			ContractPeriod:   "P1M",
		},
		{
			ID:               "po-2",
			Name:             "cloud-router-10g",
			ServiceMetroArea: "de-fra",
			ContractPeriod:   "P1Y",
		},
	}

	api := ixapi.NewTestClient(map[string]any{
		"/api/v3/decix-vrf-v1/product-offerings": offerings,
	})
	api.CloudRouterEnabled = true

	ctx := context.Background()
	err := productOfferingCloudVRFRead(ctx, res, api)

	if err == nil {
		t.Error("expected error when filter returns multiple product offerings")
	}
}
