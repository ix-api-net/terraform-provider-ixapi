package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewPortsDataSource creates a data source for querying pops
func NewPortsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query for a set of ports. You can filter by `pop`, `media_type`, `speed` and much more. Ports are allocated using port-reservations and connections. You can request a new connection and ports will be allocated. You can add ports to a connection by creating port reservations.",
		ReadContext: crud.Read(portsRead),
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by id of account managing the port"),
			"consuming_account": schemas.DataSourceQuery(
				"Filter by id of the account using the port"),
			"state": schemas.DataSourceQuery(
				"Filter by states like `production`, `allocated`, `testing`."),
			"media_type": schemas.DataSourceQuery(
				"Filter ports by media type"),
			"pop": schemas.DataSourceQuery(
				"Filter by Point Of Presence"),
			"name": schemas.DataSourceQuery(
				"Filter by name of the port"),
			"external_ref": schemas.DataSourceQuery(
				"Filter by external reference"),
			"device": schemas.DataSourceQuery(
				"Filter by device"),
			"speed": schemas.DataSourceQueryInt(
				"Filter by speed"),
			"network_connection": schemas.DataSourceQuery(
				"Filter by connection"),
			"ports": schemas.IntoDataSourceResultsSchema(
				schemas.PortSchema()),
		},
	}
}

// Make ports query from resource data
func portsQuery(res *schema.ResourceData) *ixapi.PortsListQuery {
	qry := &ixapi.PortsListQuery{}

	state, hasState := res.GetOk("state")
	mediaType, hasMediaType := res.GetOk("media_type")
	pop, hasPop := res.GetOk("pop")
	name, hasName := res.GetOk("name")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	device, hasDevice := res.GetOk("device")
	speed, hasSpeed := res.GetOk("speed")
	connection, hasConnection := res.GetOk("network_connection")

	if hasState {
		qry.State = state.(string)
	}
	if hasMediaType {
		qry.MediaType = mediaType.(string)
	}
	if hasPop {
		qry.Pop = pop.(string)
	}
	if hasName {
		qry.Name = name.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasDevice {
		qry.Device = device.(string)
	}
	if hasSpeed {
		qry.Speed = fmt.Sprintf("%d", speed.(int))
	}
	if hasConnection {
		qry.Connection = connection.(string)
	}

	return qry
}

// Fetch Ports
func fetchPorts(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.Port, error) {
	qry := portsQuery(res)
	results, err := api.PortsList(ctx, qry)
	if err != nil {
		return nil, err
	}

	// Additional filters: managing / consuming account
	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")

	filtered := make([]*ixapi.Port, 0, len(results))
	for _, port := range results {
		if hasManagingAccount && port.ManagingAccount != managingAccount.(string) {
			continue
		}
		if hasConsumingAccount && port.ConsumingAccount != consumingAccount.(string) {
			continue
		}
		filtered = append(filtered, port)
	}

	return filtered, nil
}

// Query ports
func portsRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	results, err := fetchPorts(ctx, res, api)
	if err != nil {
		return err
	}
	ports, err := schemas.FlattenModels(results)
	if err != nil {
		return err
	}
	if err := res.Set("ports", ports); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewPortDataSource is for querying and referencing a specific port
func NewPortDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get information about a specific port identified by ID",
		ReadContext: crud.Read(portRead),
		Schema:      schemas.IntoDataSourceSchema(schemas.PortSchema(), "id"),
	}
}

// Read a single port
func portRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Get("id").(string)
	port, err := api.PortsRead(ctx, id)
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(port, res); err != nil {
		return err
	}
	res.SetId(id)
	return nil
}
