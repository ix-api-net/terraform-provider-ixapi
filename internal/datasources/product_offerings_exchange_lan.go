package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewProductOfferingsExchangeLan creates a new data source
// for querying a list of exchange lan product offerings
func NewProductOfferingsExchangeLan() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to get a list of exchange lan product offerings",
		ReadContext: exchangeLanProductOfferingsRead,
		Schema: map[string]*schema.Schema{
			"service_metro_area":          schemas.DataSourceQuery(),
			"service_metro_area_network":  schemas.DataSourceQuery(),
			"handover_metro_area_network": schemas.DataSourceQuery(),
			"handover_metro_area":         schemas.DataSourceQuery(),
			"product_offering": schemas.IntoDataSourceResultsSchema(
				schemas.ExchangeLanNetworkProductOfferingSchema(),
			),
		},
	}
}

// Retrieve a list of exchange lan product offerings
func exchangeLanProductOfferingsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	return nil
}

// NewProductOfferingExchangeLan creates a new data source for
// retrieving an exchange lan product offering.
func NewProductOfferingExchangeLan() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to get an exchange lan product offering",
		ReadContext: exchangeLanProductOfferingsRead,
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

	return nil
}
