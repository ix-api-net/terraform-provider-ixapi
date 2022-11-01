package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServicesCloudDataSource creates a data source for
// querying network services of type cloud virtual circuit.
func NewNetworkServicesCloudDataSource() *schema.Resource {
	s := networkServiceQuerySchema(schemas.CloudNetworkServiceSchema())
	return &schema.Resource{
		Description: "Get network services of type: cloud",
		ReadContext: networkServicesCloudRead,
		Schema:      s,
	}
}

// Retrieve cloud virtual circuit network services and filter by pop or product offering
// or anything else. See `networkServiceQuery`.
func networkServicesCloudRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkServiceQuery(ixapi.CloudNetworkServiceType, res)
	services, err := api.NetworkServicesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Optional metro_are_network filtering
	filtered := make([]*ixapi.CloudNetworkService, 0, len(services))
	for _, svc := range services {
		resns, ok := svc.(*ixapi.CloudNetworkService)
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

// NewNetworkServiceCloudDataSource creates a network service
// data source for cloud virtual circuits, identified by ID
func NewNetworkServiceCloudDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Get a cloud virtual circuit network service by ID",
		ReadContext: networkServiceCloudRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(
				schemas.CloudNetworkServiceSchema()),
			schemas.DataSourceID()),
	}
}

// Retrieve a single cloud virtual circuit network service by ID
func networkServiceCloudRead(
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
	resns, ok := ns.(*ixapi.CloudNetworkService)
	if !ok {
		return diag.Errorf(
			"received %s instead of cloud virtual circuit (cloud_vc) network service from API",
			ns.PolymorphicType(),
		)
	}
	if err := schemas.SetResourceData(resns, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(resns.ID)
	return nil
}
