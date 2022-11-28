package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewProductOfferingsConnectionDataSource creates a new data source
// for querying connection product offerings
func NewProductOfferingsConnectionDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the product offerings data source to get a list of connection product offerings",
		ReadContext: crud.Read(productOfferingsConnectionRead),
		Schema: map[string]*schema.Schema{
			"name": schemas.DataSourceQuery(
				"Filter product offerings by name"),
			"handover_metro_area": schemas.DataSourceQuery(
				"Filter by handover metro area ID, see related metro area data source."),
			"handover_metro_area_network": schemas.DataSourceQuery(
				"Filter by handover metro area network ID, see related data source."),
			"handover_pop": schemas.DataSourceQuery(
				"Filter by ID of the point of presence, where the physical port will be present."),
			"physical_port_speed": schemas.DataSourceQueryInt(
				"Filter by physical port speed"),
			"service_provider": schemas.DataSourceQuery(
				"Filter by service provider name"),
			"downgrade_allowed": schemas.DataSourceQueryBool(
				"Find connection product offerings where downgrade is allowed"),
			"upgrade_allowed": schemas.DataSourceQueryBool(
				"Find connection product offerings where upgrade is allowed"),
			"cross_connect_initiator": schemas.DataSourceQuery(
				"Filter by cross connect initiator"),
			"product_offerings": schemas.IntoDataSourceResultsSchema(
				schemas.ConnectionProductOfferingSchema()),
		},
	}
}

func productOfferingsConnectionQuery(
	res *schema.ResourceData,
) *ixapi.ProductOfferingsListQuery {
	// Filters
	name, hasName := res.GetOk("name")
	handoverMetroArea, hasHandoverMetroArea := res.GetOk("handover_metro_area")
	handoverMetroAreaNetwork, hasHandoverMetroAreaNetwork := res.GetOk("handover_metro_area_network")
	handoverPop, hasHandoverPop := res.GetOk("handover_pop")
	serviceProvider, hasServiceProvider := res.GetOk("service_provider")
	downgradeAllowed, hasDowngradeAllowed := res.GetOk("downgrade_allowed")
	upgradeAllowed, hasUpgradeAllowed := res.GetOk("upgrade_allowed")
	physicalPortSpeed, hasPhysicalPortSpeed := res.GetOk("physical_port_speed")

	// Query
	qry := &ixapi.ProductOfferingsListQuery{
		Type: ixapi.ConnectionProductOfferingType,
	}

	if hasName {
		qry.Name = name.(string)
	}
	if hasHandoverMetroArea {
		qry.HandoverMetroArea = handoverMetroArea.(string)
	}
	if hasHandoverMetroAreaNetwork {
		qry.HandoverMetroAreaNetwork = handoverMetroAreaNetwork.(string)
	}
	if hasHandoverPop {
		qry.HandoverPop = handoverPop.(string)
	}
	if hasServiceProvider {
		qry.ServiceProvider = serviceProvider.(string)
	}
	if hasDowngradeAllowed {
		if downgradeAllowed.(bool) {
			qry.DowngradeAllowed = "true"
		} else {
			qry.DowngradeAllowed = "false"
		}
	}
	if hasUpgradeAllowed {
		if upgradeAllowed.(bool) {
			qry.UpgradeAllowed = "true"
		} else {
			qry.UpgradeAllowed = "false"
		}
	}
	if hasPhysicalPortSpeed {
		qry.PhysicalPortSpeed = physicalPortSpeed.(int)
	}

	return qry
}

func fetchProductOfferingsConnection(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.ConnectionProductOffering, error) {
	// Query
	qry := productOfferingsConnectionQuery(res)
	offerings, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}

	// Apply additional filters
	cci, hasCci := res.GetOk("cross_connect_initiator")
	hop, hasHop := res.GetOk("handover_pop")

	cpos := make([]*ixapi.ConnectionProductOffering, 0, len(offerings))
	for _, off := range offerings {
		c, ok := off.(*ixapi.ConnectionProductOffering)
		if !ok {
			continue // should not happen on well behaved servers
		}
		if hasCci && c.CrossConnectInitiator != cci.(string) {
			continue
		}
		if hasHop && c.HandoverPop != nil && *c.HandoverPop != hop.(string) {
			continue
		}
		cpos = append(cpos, c)
	}
	return cpos, nil
}

// Fetch product offerings of type: connection
func productOfferingsConnectionRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	offerings, err := fetchProductOfferingsConnection(ctx, res, api)
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

// NewProductOfferingConnectionDataSource retrievs a product
// offering by ID or by name.
func NewProductOfferingConnectionDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_product_offering_connection` data source to select a specific product offering by ID or a combination of attributes that match an unique product offering.",

		ReadContext: crud.Read(productOfferingConnectionRead),

		Schema: schemas.IntoDataSourceSchema(
			schemas.ConnectionProductOfferingSchema()),
	}
}

// Read single product offering of type connection
func productOfferingConnectionRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")
	var offering *ixapi.ConnectionProductOffering
	if hasID && id.(string) != "" {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		po, ok := result.(*ixapi.ConnectionProductOffering)
		if !ok {
			return fmt.Errorf(
				"API did return an %s instead of a connection product offering",
				result.PolymorphicType())
		}
		offering = po
	} else {
		offerings, err := fetchProductOfferingsConnection(ctx, res, api)
		if err != nil {
			return err
		}
		if len(offerings) == 0 {
			return fmt.Errorf("No connection product offering did match")
		}
		if len(offerings) > 1 {
			return fmt.Errorf("The query did not return an unique connection product offering")
		}
		offering = offerings[0]
	}

	if err := schemas.SetResourceData(offering, res); err != nil {
		return err
	}
	res.SetId(offering.ID)
	return nil
}
