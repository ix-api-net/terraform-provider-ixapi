package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewDevicesDataSource creates a data source for querying devices
func NewDevicesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "The Devices data source can be used to query devices available",
		ReadContext: devicesRead,
		Schema: map[string]*schema.Schema{
			"name": schemas.DataSourceQuery(
				"Filter devices by name"),
			"capability_media_type": schemas.DataSourceQuery(
				"Filter devices by media type: e.g. 1000BASE-LX, 10GBASE-LR"),
			"capability_speed": schemas.DataSourceQuery(
				"Filter devices by speed as a string, so you can use modifiers <, <= or >, >=.), for example: >= 10000"),
			"facility": schemas.DataSourceQuery(
				"Filter by facility ID"),
			"devices": schemas.IntoDataSourceResultsSchema(
				schemas.DeviceSchema()),
		},
	}
}

// Fetch devices
func devicesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	return nil
}
