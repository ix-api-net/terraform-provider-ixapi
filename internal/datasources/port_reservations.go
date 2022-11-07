package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewPortReservationsDataSource creates a data source for querying
// port reservations.
func NewPortReservationsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query for port reservations. Use the `ixapi_port_reservation` resource to reference and change port reservations of a connection. You can use the `ixapi_port_reservation` **resource** to import a port reservation, from a connection by passing a `port_index`. In most cases you will most likely use the port reservation resource instead of the data source.",
		ReadContext: crud.Read(portReservationsRead),
		Schema: map[string]*schema.Schema{
			"state": schemas.DataSourceQuery(
				"Filter by states like `production`, `allocated`, `testing`."),
			"connection": schemas.DataSourceQuery(
				"Filter by connection"),
			"port": schemas.DataSourceQuery(
				"Filter by port"),
			"external_ref": schemas.DataSourceQuery(
				"Filter by external ref"),
			"port_reservations": schemas.IntoDataSourceResultsSchema(
				schemas.PortReservationSchema()),
		},
	}
}

// PortReservations list query
func portReservationsQuery(res *schema.ResourceData) *ixapi.PortReservationsListQuery {
	qry := &ixapi.PortReservationsListQuery{}

	state, hasState := res.GetOk("state")
	connection, hasConnection := res.GetOk("connection")
	port, hasPort := res.GetOk("port")
	externalRef, hasExternalRef := res.GetOk("external_ref")

	if hasState {
		qry.State = state.(string)
	}
	if hasConnection {
		qry.Connection = connection.(string)
	}
	if hasPort {
		qry.Port = port.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}

	return qry
}

// Fetch port reservations
func portReservationsRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := portReservationsQuery(res)
	results, err := api.PortReservationsList(ctx, qry)
	if err != nil {
		return err
	}
	// Flatten results
	reservations, err := schemas.FlattenModels(results)
	if err != nil {
		return err
	}
	if err := res.Set("port_reservations", reservations); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewPortReservationDataSource creates a data source for
// retrieving a port reservation identified by ID.
func NewPortReservationDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_port_reservation` data source to retrieve a port reservation by ID. In most usecases consider using the `ixapi_port_reservation` resource.",
		ReadContext: crud.Read(portReservationRead),
		Schema: schemas.IntoDataSourceSchema(
			schemas.PortReservationSchema(),
			"id"),
	}
}

// Fetch port reservation
func portReservationRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Get("id").(string)
	reservation, err := api.PortReservationsRead(ctx, id)
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(reservation, res); err != nil {
		return err
	}
	res.SetId(id)
	return nil
}
