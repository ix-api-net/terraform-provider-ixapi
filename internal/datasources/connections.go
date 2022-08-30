package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewConnectionsDataSource creates a new data source for
// a collection of connections.
func NewConnectionsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `connections` data source find available connections",
		ReadContext: connectionsRead,
		Schema: map[string]*schema.Schema{
			"managing_account":   schemas.DataSourceQuery(),
			"consuming_account":  schemas.DataSourceQuery(),
			"metro_area_network": schemas.DataSourceQuery(),
			"name":               schemas.DataSourceQuery(),
			"pop":                schemas.DataSourceQuery(),
			"connections": schemas.IntoDataSourceResultsSchema(
				schemas.ConnectionSchema(),
			),
		},
	}
}

// Operations

// Get a list of connections
func connectionsRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	all, err := api.ConnectionsList(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	// Filters
	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	name, hasName := res.GetOk("name")
	pop, hasPop := res.GetOk("pop")

	filtered := make([]*ixapi.Connection, 0, len(all))
	for _, conn := range all {
		if hasManagingAccount && conn.ManagingAccount != managingAccount.(string) {
			continue
		}
		if hasConsumingAccount && conn.ConsumingAccount != consumingAccount.(string) {
			continue
		}
		if hasName && conn.Name != name.(string) {
			continue
		}
		if hasPop && conn.Pop != pop.(string) {
			continue
		}
		filtered = append(filtered, conn)
	}

	connections, err := schemas.FlattenModels(filtered)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("connections", connections); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewConnectionDataSource creates a data source schema for
// reading a single connection identified by ID or external ref
func NewConnectionDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `conneciton` data source to get a single connection by ID or external ref",
		ReadContext: connectionRead,
		Schema:      schemas.IntoDataSourceSchema(schemas.ConnectionSchema()),
	}
}

func connectionRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	id, hasID := res.GetOk("id")
	ref, hasRef := res.GetOk("external_ref")
	name, hasName := res.GetOk("name")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")

	if !hasID && !hasRef && !hasName {
		return diag.Errorf("the connection `name`, `id` or `external_ref` is required")
	}

	var conn *ixapi.Connection
	if hasID {
		c, err := api.ConnectionsRead(ctx, id.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		conn = c
	} else {
		qry := &ixapi.ConnectionsListQuery{}
		if hasRef {
			qry.ExternalRef = ref.(string)
		}
		if hasName {
			qry.Name = name.(string)
		}
		result, err := api.ConnectionsList(ctx, qry)
		if err != nil {
			return diag.FromErr(err)
		}
		connections := make([]*ixapi.Connection, 0, len(result))
		for _, conn := range connections {
			if hasConsumingAccount && conn.ConsumingAccount != consumingAccount.(string) {
				continue
			}
			connections = append(connections, conn)
		}

		if len(connections) == 0 {
			return diag.Errorf("no connection could be found")
		}
		if len(connections) > 1 {
			return diag.Errorf("multiple connections returned for this external_ref")
		}
		conn = connections[0]
	}

	if err := schemas.SetResourceData(conn, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(conn.ID)

	return nil
}
