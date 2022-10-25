package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/crud"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewProductOfferingsP2MPVCDataSource creates a data source
// for querying product offerings of type p2mp_vc
func NewProductOfferingsP2MPVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can be used to find product offerings for p2mp virtual circuits.",
		ReadContext: crud.Read(productOfferingsP2MPVCRead),
		Schema: productOfferingsVCSchema(
			schemas.IntoDataSourceResultsSchema(
				schemas.P2MPNetworkProductOfferingSchema())),
	}
}

// Load P2MP product offerings from API
func fetchProductOfferingsP2MPVC(
	ctx context.Context,
	api *ixapi.Client,
	qry *ixapi.ProductOfferingsListQuery,
) ([]*ixapi.P2MPNetworkProductOffering, error) {
	results, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}
	offerings := make([]*ixapi.P2MPNetworkProductOffering, 0, len(results))
	for _, off := range results {
		po, ok := off.(*ixapi.P2MPNetworkProductOffering)
		if !ok {
			continue // Skip invalid response from server
		}
		offerings = append(offerings, po)
	}
	return offerings, nil
}

// Read operation for terraform data source
func productOfferingsP2MPVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := productOfferingsVCQuery(ixapi.P2MPNetworkProductOfferingType, res)
	offerings, err := fetchProductOfferingsP2MPVC(ctx, api, qry)
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

// NewProductOfferingP2MPVCDataSource creates a new data source
// for a single p2mp product offering
func NewProductOfferingP2MPVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single p2mp virtual circuit product offering.",
		ReadContext: crud.Read(productOfferingP2MPCVRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.P2MPNetworkProductOfferingSchema()),
	}
}

// Implement read operation for data source
func productOfferingP2MPCVRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")

	var offering *ixapi.P2MPNetworkProductOffering
	if hasID {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		po, ok := result.(*ixapi.P2MPNetworkProductOffering)
		if !ok {
			return fmt.Errorf(
				"API did respond with a %s instead of a p2mp_vc product offering",
				result.PolymorphicType())
		}
		offering = po
	} else {
		qry := productOfferingsVCQuery(ixapi.P2MPNetworkProductOfferingType, res)
		offerings, err := fetchProductOfferingsP2MPVC(ctx, api, qry)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("no such p2mp_vc product offering could be found")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("the p2mp_vc product offering is not uniquely identified")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
