package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/filter"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewNetworkServicesExchangeLanDataSource creates a data source for
// querying network services of type exchange lan.
func NewNetworkServicesExchangeLanDataSource() *schema.Resource {
	s := networkServiceQuerySchema(schemas.ExchangeLanNetworkServiceSchema())
	s["metro_area_network"] = schemas.DataSourceQuery(
		"Filter by metro area network id, see related data source")
	return &schema.Resource{
		Description: "Get network services of type: exchange lan",
		ReadContext: networkServicesExchangeLanRead,
		Schema:      s,
	}
}

// Retrieve exchange lan network services and filter by pop or product offering
func networkServicesExchangeLanRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Filters
	metroAreaNetwork, hasMetroAreaNetwork := res.GetOk("metro_area_network")

	qry := networkServiceQuery(ixapi.ExchangeLanNetworkServiceType, res)
	services, err := api.NetworkServicesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Optional metro_are_network filtering
	filtered := make([]*ixapi.ExchangeLanNetworkService, 0, len(services))
	for _, svc := range services {
		elns, ok := svc.(*ixapi.ExchangeLanNetworkService)
		if !ok {
			continue // should not happen with well behaving servers
		}
		if filter.Missing(elns.MetroAreaNetwork, metroAreaNetwork, hasMetroAreaNetwork) {
			continue
		}
		filtered = append(filtered, elns)
	}

	flat, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := res.Set("network_services", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewNetworkServiceExchangeLanDataSource creates a network service
// data source for exchange lans, identified by ID
func NewNetworkServiceExchangeLanDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Get an exchange lan network service by ID",
		ReadContext: networkServiceExchangeLanRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(
				schemas.ExchangeLanNetworkServiceSchema()),
			schemas.DataSourceID()),
	}
}

// Retrieve a single exchange lan network service by ID
func networkServiceExchangeLanRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	id := res.Get("id")

	ns, err := api.NetworkServicesRead(ctx, id.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	elns, ok := ns.(*ixapi.ExchangeLanNetworkService)
	if !ok {
		return diag.Errorf(
			"Received non exchange lan network service type from server: %s",
			ns.PolymorphicType(),
		)
	}
	if err := schemas.SetResourceData(elns, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(elns.ID)
	return nil
}
