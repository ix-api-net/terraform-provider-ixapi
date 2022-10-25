package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/crud"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewProductOfferingsP2PVCDataSource creates a data source
// for querying product offerings of type p2p_vc
func NewProductOfferingsP2PVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "This data source can be used to find product offerings for p2p virtual circuits.",
		ReadContext: crud.Read(productOfferingsP2PVCRead),
		Schema: productOfferingsVCSchema(
			schemas.IntoDataSourceResultsSchema(
				schemas.P2PNetworkProductOfferingSchema())),
	}
}

// Load P2P product offerings from API
func fetchProductOfferingsP2PVC(
	ctx context.Context,
	api *ixapi.Client,
	qry *ixapi.ProductOfferingsListQuery,
) ([]*ixapi.P2PNetworkProductOffering, error) {
	results, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}
	offerings := make([]*ixapi.P2PNetworkProductOffering, 0, len(results))
	for _, off := range results {
		po, ok := off.(*ixapi.P2PNetworkProductOffering)
		if !ok {
			continue // Skip invalid response from server
		}
		offerings = append(offerings, po)
	}
	return offerings, nil
}

// Read operation for terraform data source
func productOfferingsP2PVCRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := productOfferingsVCQuery(ixapi.P2PNetworkProductOfferingType, res)
	offerings, err := fetchProductOfferingsP2PVC(ctx, api, qry)
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

// NewProductOfferingP2PVCDataSource creates a new data source
// for a single p2p product offering
func NewProductOfferingP2PVCDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single p2p virtual circuit product offering.",
		ReadContext: crud.Read(productOfferingP2PCVRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.P2PNetworkProductOfferingSchema()),
	}
}

// Implement read operation for data source
func productOfferingP2PCVRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")

	var offering *ixapi.P2PNetworkProductOffering
	if hasID {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		po, ok := result.(*ixapi.P2PNetworkProductOffering)
		if !ok {
			return fmt.Errorf(
				"API did respond with a %s instead of a p2p_vc product offering",
				result.PolymorphicType())
		}
		offering = po
	} else {
		qry := productOfferingsVCQuery(ixapi.P2PNetworkProductOfferingType, res)
		offerings, err := fetchProductOfferingsP2PVC(ctx, api, qry)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("no such p2p_vc product offering could be found")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("the p2p_vc product offering is not uniquely identified")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
