package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServicesMP2MPDataSource creates a data source for
// querying network services of type mp2mp virtual circuit.
func NewNetworkServicesMP2MPDataSource() *schema.Resource {
	s := networkServiceQuerySchema(schemas.MP2MPNetworkServiceSchema())
	return &schema.Resource{
		Description: "Get network services of type: mp2mp virtual circuit",
		ReadContext: networkServicesMP2MPRead,
		Schema:      s,
	}
}

// Retrieve mp2mp virtual circuit network services and filter by pop or product offering
func networkServicesMP2MPRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkServiceQuery(ixapi.MP2MPNetworkServiceType, res)
	services, err := api.NetworkServicesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Optional metro_are_network filtering
	filtered := make([]*ixapi.MP2MPNetworkService, 0, len(services))
	for _, svc := range services {
		resns, ok := svc.(*ixapi.MP2MPNetworkService)
		if !ok {
			continue // should not happen with well behaving servers
		}
		filtered = append(filtered, resns)
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

// NewNetworkServiceMP2MPDataSource creates a network service
// data source for mp2mp virtual circuits, identified by ID
func NewNetworkServiceMP2MPDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Get a mp2mp virtual circuit network service by ID",
		ReadContext: networkServiceMP2MPRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(
				schemas.MP2MPNetworkServiceSchema()),
			schemas.DataSourceID()),
	}
}

// Retrieve a single mp2mp virtual circuit network service by ID
func networkServiceMP2MPRead(
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
	resns, ok := ns.(*ixapi.MP2MPNetworkService)
	if !ok {
		return diag.Errorf(
			"received %s instead of mp2mp virtual circuit (mp2mp_vc) network service from API",
			ns.PolymorphicType(),
		)
	}
	if err := schemas.SetResourceData(resns, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(resns.ID)
	return nil
}
