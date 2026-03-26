package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func cloudRouterProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"bandwidth_max": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"bandwidth_min": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service_metro_area": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service_metro_area_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service_metro_area_network": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"service_metro_area_network_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"contract_period": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func NewProductOfferingsCloudVRFDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can be used to find product offerings for DE-CIX Cloud ROUTER (VRF).",
		ReadContext: crud.Read(productOfferingsCloudVRFRead),
		Schema: map[string]*schema.Schema{
			"product_offerings": schemas.IntoDataSourceResultsSchema(
				cloudRouterProductOfferingSchema()),
			"limit": schemas.DataSourceQueryInt(
				"Limit the number of results"),
			"offset": schemas.DataSourceQueryInt(
				"Offset for pagination"),
			"bandwidth": schemas.DataSourceQueryInt(
				"Filter by bandwidth in Mbit/s"),
			"name": schemas.DataSourceQuery(
				"Filter by product offering name"),
			"service_metro_area": schemas.DataSourceQuery(
				"Filter by service metro area ID"),
			"service_metro_area_network": schemas.DataSourceQuery(
				"Filter by service metro area network ID"),
			"contract_period": schemas.DataSourceQuery(
				"Filter by contract period (e.g., P1M, P1Y)"),
		},
	}
}

func productOfferingsCloudVRFRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}
	qry := &ixapi.CloudRouterProductOfferingsListQuery{}

	if limit, ok := res.GetOk("limit"); ok {
		qry.Limit = limit.(int)
	}
	if offset, ok := res.GetOk("offset"); ok {
		qry.Offset = offset.(int)
	}
	if bandwidth, ok := res.GetOk("bandwidth"); ok {
		qry.Bandwidth = bandwidth.(int)
	}
	if name, ok := res.GetOk("name"); ok {
		qry.Name = name.(string)
	}
	if serviceMetroArea, ok := res.GetOk("service_metro_area"); ok {
		qry.ServiceMetroArea = serviceMetroArea.(string)
	}
	if serviceMetroAreaNetwork, ok := res.GetOk("service_metro_area_network"); ok {
		qry.ServiceMetroAreaNetwork = serviceMetroAreaNetwork.(string)
	}
	if contractPeriod, ok := res.GetOk("contract_period"); ok {
		qry.ContractPeriod = contractPeriod.(string)
	}

	offerings, err := api.CloudRouterProductOfferingsList(ctx, qry)
	if err != nil {
		return err
	}

	flat, err := schemas.FlattenModels(offerings)
	if err != nil {
		return err
	}
	if err := res.Set("product_offerings", flat); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

func NewProductOfferingCloudVRFDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single Cloud ROUTER (VRF) product offering.",
		ReadContext: crud.Read(productOfferingCloudVRFRead),
		Schema: schemas.IntoDataSourceSchema(
			cloudRouterProductOfferingSchema()),
	}
}

func productOfferingCloudVRFRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}
	id, hasID := res.GetOk("id")

	var offering *ixapi.CloudRouterProductOffering
	if hasID {
		result, err := api.CloudRouterProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		offering = result
	} else {
		qry := &ixapi.CloudRouterProductOfferingsListQuery{}

		if bandwidth, ok := res.GetOk("bandwidth"); ok {
			qry.Bandwidth = bandwidth.(int)
		}
		if name, ok := res.GetOk("name"); ok {
			qry.Name = name.(string)
		}
		if serviceMetroArea, ok := res.GetOk("service_metro_area"); ok {
			qry.ServiceMetroArea = serviceMetroArea.(string)
		}
		if serviceMetroAreaNetwork, ok := res.GetOk("service_metro_area_network"); ok {
			qry.ServiceMetroAreaNetwork = serviceMetroAreaNetwork.(string)
		}
		if contractPeriod, ok := res.GetOk("contract_period"); ok {
			qry.ContractPeriod = contractPeriod.(string)
		}

		offerings, err := api.CloudRouterProductOfferingsList(ctx, qry)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("no such Cloud ROUTER product offering could be found")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("the Cloud ROUTER product offering is not uniquely identified")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
