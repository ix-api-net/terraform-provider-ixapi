package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewProductOfferingExchangeLan creates a new data source for
// querying exchange lan product offerings
func NewProductOfferingExchangeLan() *schema.Resource {
	return &schema.Resource{
		Description: "Use the data source to get an exchange lan product offering",
		ReadContext: exchangeLanProductOfferingsRead,
		Schema: map[string]*schema.Schema{
			"service_metro_area":         schemas.DataSourceQuery(),
			"service_metro_area_network": schemas.DataSourceQuery(),
			"product_offering": schemas.IntoDataSourceResultsSchema(
				schemas.ExchangeLanNetworkProductOfferingSchema(),
			),
		},
	}
}

// Retrieve the exchange lan product offerings matching the filters
func exchangeLanProductOfferingsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {

	return nil
}
