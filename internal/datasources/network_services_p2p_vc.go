package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewNetworkServicesP2PDataSource creates a data source for
// querying network services of type p2p virtual circuit.
func NewNetworkServicesP2PDataSource() *schema.Resource {
	s := networkServiceQuerySchema(schemas.P2PNetworkServiceSchema())
	return &schema.Resource{
		Description: "Get network services of type: p2p virtual circuit",
		ReadContext: networkServicesP2PRead,
		Schema:      s,
	}
}

// Retrieve p2p virtual circuit network services and filter by pop or product offering
func networkServicesP2PRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := networkServiceQuery(ixapi.P2PNetworkServiceType, res)
	services, err := api.NetworkServicesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	// Optional metro_are_network filtering
	filtered := make([]*ixapi.P2PNetworkService, 0, len(services))
	for _, svc := range services {
		resns, ok := svc.(*ixapi.P2PNetworkService)
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

// NewNetworkServiceP2PDataSource creates a network service
// data source for p2p virtual circuits, identified by ID
func NewNetworkServiceP2PDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Get an p2p virtual circuit network service by ID",
		ReadContext: networkServiceP2PRead,
		Schema: schemas.Combine(
			schemas.IntoDataSourceSchema(
				schemas.P2PNetworkServiceSchema()),
			schemas.DataSourceID()),
	}
}

// Retrieve a single p2p virtual circuit network service by ID
func networkServiceP2PRead(
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
	resns, ok := ns.(*ixapi.P2PNetworkService)
	if !ok {
		return diag.Errorf(
			"received %s instead of p2p virtual circuit (p2p_vc) network service from API",
			ns.PolymorphicType(),
		)
	}
	if err := schemas.SetResourceData(resns, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(resns.ID)
	return nil
}
