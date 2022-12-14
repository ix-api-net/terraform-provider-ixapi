package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/filter"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewProductOfferingsExchangeLanDataSource creates a new data source
// for querying a list of exchange lan product offerings
func NewProductOfferingsExchangeLanDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to get a list of exchange lan product offerings",
		ReadContext: exchangeLanProductOfferingsRead,
		Schema: map[string]*schema.Schema{
			"name": schemas.DataSourceQuery(
				"Filter by name of the product offering"),
			"service_provider": schemas.DataSourceQuery(
				"Filter by service provider name"),
			"service_metro_area": schemas.DataSourceQuery(
				"Filter by metro area id, see releated data source. The service metro area is where the service is delived to the customer."),
			"service_metro_area_network": schemas.DataSourceQuery(
				"Filter by metro area network id, see related data source. The service is directly accessible through this metro area network"),
			"handover_metro_area": schemas.DataSourceQuery(
				"Filter by metro area id, see related data source. The service will be accessed from this metro area"),
			"handover_metro_area_network": schemas.DataSourceQuery(
				"Filter by metro area network id, see related data source. The service will be accessed through the handover metro area network."),
			"product_offerings": schemas.IntoDataSourceResultsSchema(
				schemas.ExchangeLanNetworkProductOfferingSchema(),
			),
		},
	}
}

// Fetch exchange lan product offerings
func fetchExchangeLanProductOfferings(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.ExchangeLanNetworkProductOffering, error) {
	// Filters
	name, hasName := res.GetOk("name")
	serviceProvider, hasServiceProvider := res.GetOk("service_provider")

	serviceMetroArea, hasServiceMetroArea := res.GetOk("service_metro_area")
	serviceMetroAreaNetwork, hasServiceMetroAreaNetwork := res.GetOk("service_metro_area_network")
	handoverMetroArea, hasHandoverMetroArea := res.GetOk("handover_metro_area")
	handoverMetroAreaNetwork, hasHandoverMetroAreaNetwork := res.GetOk("handover_metro_area_network")

	// Query
	qry := &ixapi.ProductOfferingsListQuery{
		Type: ixapi.ExchangeLanNetworkProductOfferingType,
	}
	if hasServiceProvider {
		qry.ServiceProvider = serviceProvider.(string)
	}
	if hasServiceMetroArea {
		qry.ServiceMetroArea = serviceMetroArea.(string)
	}
	if hasServiceMetroAreaNetwork {
		qry.ServiceMetroAreaNetwork = serviceMetroAreaNetwork.(string)
	}
	if hasHandoverMetroArea {
		qry.HandoverMetroArea = handoverMetroArea.(string)
	}
	if hasHandoverMetroAreaNetwork {
		qry.HandoverMetroAreaNetwork = handoverMetroAreaNetwork.(string)
	}

	// Make request
	offerings, err := api.ProductOfferingsList(ctx, qry)
	if err != nil {
		return nil, err
	}

	elpos := make([]*ixapi.ExchangeLanNetworkProductOffering, 0, len(offerings))
	for _, off := range offerings {
		elpo, ok := off.(*ixapi.ExchangeLanNetworkProductOffering)
		if !ok {
			continue // Should not happen with well behaved servers
		}
		if filter.Missing(elpo.Name, name, hasName) {
			continue
		}
		elpos = append(elpos, elpo)
	}
	return elpos, nil
}

// Retrieve a list of exchange lan product offerings
func exchangeLanProductOfferingsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	elpos, err := fetchExchangeLanProductOfferings(ctx, res, api)
	flat, err := schemas.FlattenModels(elpos)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("product_offerings", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewProductOfferingExchangeLanDataSource creates a new data source for
// retrieving an exchange lan product offering.
func NewProductOfferingExchangeLanDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to get an exchange lan product offering",
		ReadContext: exchangeLanProductOfferingRead,
		Schema: schemas.IntoDataSourceSchema(
			schemas.ExchangeLanNetworkProductOfferingSchema()),
	}
}

// Retrieve the exchange lan product offering matching the filters
func exchangeLanProductOfferingRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	id, hasID := res.GetOk("id")

	var offering *ixapi.ExchangeLanNetworkProductOffering
	if hasID && id.(string) != "" {
		result, err := api.ProductOfferingsRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		elpo, ok := result.(*ixapi.ExchangeLanNetworkProductOffering)
		if !ok {
			return diag.Errorf("API did not return an exchange lan network product offering")
		}
		offering = elpo
	} else {
		offerings, err := fetchExchangeLanProductOfferings(ctx, res, api)
		if err != nil {
			return diag.FromErr(err)
		}
		if len(offerings) == 0 {
			return diag.Errorf("No exchange lan network product offering did match")
		}
		if len(offerings) > 1 {
			return diag.Errorf(
				"The query did not return an unique exchange lan product offering")
		}
		offering = offerings[0]
	}
	if err := schemas.SetResourceData(offering, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(offering.ID)
	return nil
}
