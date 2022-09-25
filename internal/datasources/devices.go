package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
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
			"capability_speed": schemas.DataSourceQueryInt(
				"Filter devices by exact speed"),
			"capability_speed_lt": schemas.DataSourceQueryInt(
				"Filter devices by speed lower than the requested value"),
			"capability_speed_lte": schemas.DataSourceQueryInt(
				"Filter devices by speed lower or equal than the requested value"),
			"capability_speed_gt": schemas.DataSourceQueryInt(
				"Filter devices by speed greater than the requested value"),
			"capability_speed_gte": schemas.DataSourceQueryInt(
				"Filter devices by speed greater or equal than the requested value"),
			"facility": schemas.DataSourceQuery(
				"Filter by facility ID"),
			"devices": schemas.IntoDataSourceResultsSchema(
				schemas.DeviceSchema()),
		},
	}
}

// Create a list query from filters available
func devicesQuery(res *schema.ResourceData) *ixapi.DevicesListQuery {
	// Filters
	name, hasName := res.GetOk("name")
	facility, hasFacility := res.GetOk("facility")
	capMediaType, hasCapMediaType := res.GetOk("capability_media_type")
	capSpeed, hasCapSpeed := res.GetOk("capability_speed")
	capSpeedLt, hasCapSpeedLt := res.GetOk("capability_speed_lt")
	capSpeedLte, hasCapSpeedLte := res.GetOk("capability_speed_lte")
	capSpeedGt, hasCapSpeedGt := res.GetOk("capability_speed_gt")
	capSpeedGte, hasCapSpeedGte := res.GetOk("capability_speed_gte")

	// Query
	qry := &ixapi.DevicesListQuery{}
	if hasName {
		qry.Name = name.(string)
	}
	if hasFacility {
		qry.Facility = facility.(string)
	}
	if hasCapMediaType {
		qry.CapabilityMediaType = capMediaType.(string)
	}
	if hasCapSpeed {
		qry.CapabilitySpeed = capSpeed.(int)
	}
	if hasCapSpeedLt {
		qry.CapabilitySpeedLt = capSpeedLt.(int)
	}
	if hasCapSpeedLte {
		qry.CapabilitySpeedLte = capSpeedLte.(int)
	}
	if hasCapSpeedGt {
		qry.CapabilitySpeedGt = capSpeedGt.(int)
	}
	if hasCapSpeedGte {
		qry.CapabilitySpeedGte = capSpeedGte.(int)
	}

	return qry
}

// Fetch devices
func devicesRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	qry := devicesQuery(res)
	result, err := api.DevicesList(ctx, qry)
	if err != nil {
		return diag.FromErr(err)
	}

	devices, err := schemas.FlattenModels(result)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("devices", devices); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}

// NewDeviceDataSource creates a data source for a single device
func NewDeviceDataSource() *schema.Resource {
	deviceSchema := schemas.IntoDataSourceSchema(schemas.DeviceSchema())
	deviceSchema["id"].Optional = false
	deviceSchema["id"].Required = true
	return &schema.Resource{
		Description: "Use the device data source to reference a single device by ID",
		ReadContext: deviceRead,
		Schema:      deviceSchema,
	}
}

func deviceRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	id := res.Get("id")

	device, err := api.DevicesRead(ctx, id.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := schemas.SetResourceData(device, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(id.(string))

	return nil
}
