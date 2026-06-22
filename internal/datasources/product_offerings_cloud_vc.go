package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewProductOfferingsCloudVCDataSource creates a data source
// for querying product offerings of type cloud_vc
func NewProductOfferingsCloudVCDataSource() *schema.Resource {
	s := productOfferingsVCSchema(
		schemas.IntoDataSourceResultsSchema(
			schemas.CloudNetworkProductOfferingSchema()))
	s["cloud_key"] = schemas.DataSourceQuery(
		"If the service_provider_workflow is provider_first the cloud_key will be used for filtering the relevant offerings.")
	s["delivery_method"] = schemas.DataSourceQuery(
		"Filter cloud network product offerings by delivery method (shared or dedicated).")
	s["service_provider_region"] = schemas.DataSourceQuery(
		"Filter by cloud service provider region.")
	s["service_provider_pop"] = schemas.DataSourceQuery(
		"Filter by id of the service provider pop. See `ixapi_pops` data source.")
	s["handover_metro_area_network_name"] = schemas.DataSourceQuery(
		"Filter by name of the handover metro area network (e.g. FRA).")
	s["service_metro_area_network_name"] = schemas.DataSourceQuery(
		"Filter by name of the service metro area network (e.g. FRA).")

	return &schema.Resource{
		Description: "This data source can be used to find product offerings for cloud virtual circuits.",
		ReadContext: crud.Read(productOfferingsCloudVCRead),
		Schema:      s,
	}
}

// Load Cloud product offerings from API
func fetchProductOfferingsCloudVC(
	ctx context.Context,
	api *ixapi.Client,
	qry *ixapi.ProductOfferingsListQuery,
) ([]*ixapi.CloudNetworkProductOffering, error) {
	results, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}
	offerings := make([]*ixapi.CloudNetworkProductOffering, 0, len(results))
	for _, off := range results {
		po, ok := off.(*ixapi.CloudNetworkProductOffering)
		if !ok {
			continue // Skip invalid response from server
		}
		offerings = append(offerings, po)
	}
	return offerings, nil
}

// Read operation for terraform data source
func productOfferingsCloudVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := productOfferingsVCQuery(ixapi.CloudNetworkProductOfferingType, res)

	// Filters
	cloudKey, hasCloudKey := res.GetOk("cloud_key")
	deliveryMethod, hasDeliveryMethod := res.GetOk("delivery_method")
	serviceProviderRegion, hasServiceProviderRegion := res.GetOk("service_provider_region")
	serviceProviderPop, hasServiceProviderPop := res.GetOk("service_provider_pop")
	handoverMetroAreaNetworkName, hasHandoverMetroAreaNetworkName := res.GetOk("handover_metro_area_network_name")
	serviceMetroAreaNetworkName, hasServiceMetroAreaNetworkName := res.GetOk("service_metro_area_network_name")

	if hasCloudKey {
		qry.CloudKey = cloudKey.(string)
	}
	if hasDeliveryMethod {
		dm := deliveryMethod.(string)
		if dm != "dedicated" && dm != "shared" {
			return fmt.Errorf("delivery_method can only be `dedicated` or `shared`")
		}
		qry.DeliveryMethod = deliveryMethod.(string)
	}
	if hasServiceProviderRegion {
		qry.ServiceProviderRegion = serviceProviderRegion.(string)
	}
	if hasServiceProviderPop {
		qry.ServiceProviderPop = serviceProviderPop.(string)
	}
	if hasHandoverMetroAreaNetworkName {
		qry.HandoverMetroAreaNetworkName = handoverMetroAreaNetworkName.(string)
	}
	if hasServiceMetroAreaNetworkName {
		qry.ServiceMetroAreaNetworkName = serviceMetroAreaNetworkName.(string)
	}

	// API request
	offerings, err := fetchProductOfferingsCloudVC(ctx, api, qry)
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

// NewProductOfferingCloudVCDataSource creates a new data source
// for a single cloud product offering
func NewProductOfferingCloudVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single cloud virtual circuit product offering.",
		ReadContext: crud.Read(productOfferingCloudCVRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.CloudNetworkProductOfferingSchema()),
	}
}

// Implement read operation for data source
func productOfferingCloudCVRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")

	var offering *ixapi.CloudNetworkProductOffering
	if hasID {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		po, ok := result.(*ixapi.CloudNetworkProductOffering)
		if !ok {
			return fmt.Errorf(
				"API did respond with a %s instead of a cloud_vc product offering",
				result.PolymorphicType())
		}
		offering = po
	} else {
		qry := productOfferingsVCQuery(ixapi.CloudNetworkProductOfferingType, res)
		if deliveryMethod, ok := res.GetOk("delivery_method"); ok {
			qry.DeliveryMethod = deliveryMethod.(string)
		}
		if serviceProviderRegion, ok := res.GetOk("service_provider_region"); ok {
			qry.ServiceProviderRegion = serviceProviderRegion.(string)
		}
		if serviceProviderPop, ok := res.GetOk("service_provider_pop"); ok {
			qry.ServiceProviderPop = serviceProviderPop.(string)
		}
		if handoverMetroAreaNetworkName, ok := res.GetOk("handover_metro_area_network_name"); ok {
			qry.HandoverMetroAreaNetworkName = handoverMetroAreaNetworkName.(string)
		}
		if serviceMetroAreaNetworkName, ok := res.GetOk("service_metro_area_network_name"); ok {
			qry.ServiceMetroAreaNetworkName = serviceMetroAreaNetworkName.(string)
		}
		offerings, err := fetchProductOfferingsCloudVC(ctx, api, qry)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("no such cloud_vc product offering could be found")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("the cloud_vc product offering is not uniquely identified")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
