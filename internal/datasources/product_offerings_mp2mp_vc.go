package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewProductOfferingsMP2MPVCDataSource creates a data source
// for querying product offerings of type mp2mp_vc
func NewProductOfferingsMP2MPVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can be used to find product offerings for mp2mp virtual circuits.",
		ReadContext: crud.Read(productOfferingsMP2MPVCRead),
		Schema: productOfferingsVCSchema(
			schemas.IntoDataSourceResultsSchema(
				schemas.MP2MPNetworkProductOfferingSchema())),
	}
}

// Load MP2MP product offerings from API
func fetchProductOfferingsMP2MPVC(
	ctx context.Context,
	api *ixapi.Client,
	qry *ixapi.ProductOfferingsListQuery,
) ([]*ixapi.MP2MPNetworkProductOffering, error) {
	results, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}
	offerings := make([]*ixapi.MP2MPNetworkProductOffering, 0, len(results))
	for _, off := range results {
		po, ok := off.(*ixapi.MP2MPNetworkProductOffering)
		if !ok {
			continue // Skip invalid response from server
		}
		offerings = append(offerings, po)
	}
	return offerings, nil
}

// Read operation for terraform data source
func productOfferingsMP2MPVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := productOfferingsVCQuery(ixapi.MP2MPNetworkProductOfferingType, res)
	offerings, err := fetchProductOfferingsMP2MPVC(ctx, api, qry)
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

// NewProductOfferingMP2MPVCDataSource creates a new data source
// for a single mp2mp product offering
func NewProductOfferingMP2MPVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single mp2mp virtual circuit product offering.",
		ReadContext: crud.Read(productOfferingMP2MPCVRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.MP2MPNetworkProductOfferingSchema()),
	}
}

// Implement read operation for data source
func productOfferingMP2MPCVRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")

	var offering *ixapi.MP2MPNetworkProductOffering
	if hasID {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		po, ok := result.(*ixapi.MP2MPNetworkProductOffering)
		if !ok {
			return fmt.Errorf(
				"API did respond with a %s instead of a mp2mp_vc product offering",
				result.PolymorphicType())
		}
		offering = po
	} else {
		qry := productOfferingsVCQuery(ixapi.MP2MPNetworkProductOfferingType, res)
		offerings, err := fetchProductOfferingsMP2MPVC(ctx, api, qry)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("no such mp2mp_vc product offering could be found")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("the mp2mp_vc product offering is not uniquely identified")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
