package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

func productOfferingsVCSchema(
	results *schema.Schema,
) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": schemas.DataSourceQuery(
			"Filter by name of the product offering"),
		"service_provider": schemas.DataSourceQuery(
			"Filter by service provider name"),
		"service_metro_area": schemas.DataSourceQuery(
			"Filter by metro area id, see releated data source. The service metro area is where the service is delived to the customer."),
		"service_metro_area_network": schemas.DataSourceQuery(
			"Filter by metro area network id, see related data source. The service is directly accessible through this metro area network"),
		"handover_metro_area": schemas.DataSourceQuery(
			"Filter by metro area id, see related data source. The service will be accessed from this metro area"),
		"handover_metro_area_network": schemas.DataSourceQuery(
			"Filter by metro area network id, see related data source. The service will be accessed through the handover metro area network."),
		"bandwidth": schemas.DataSourceQueryInt(
			"Find product offerings where bandwidth is within the range of bandwidth_min and bandwidth_max."),
		"physical_port_speed": schemas.DataSourceQueryInt(
			"Filter by physical port speed"),
		"downgrade_allowed": schemas.DataSourceQueryBool(
			"Find connection product offerings where downgrade is allowed"),
		"upgrade_allowed": schemas.DataSourceQueryBool(
			"Find connection product offerings where upgrade is allowed"),
		"product_offerings": results,
	}
}

func productOfferingsVCQuery(
	t string,
	res *schema.ResourceData,
) *ixapi.ProductOfferingsListQuery {
	qry := &ixapi.ProductOfferingsListQuery{
		Type: t,
	}

	name, hasName := res.GetOk("name")
	serviceProvider, hasServiceProvider := res.GetOk("service_provider")
	serviceMetroArea, hasServiceMetroArea := res.GetOk("service_metro_area")
	serviceMetroAreaNetwork, hasServiceMetroAreaNetwork := res.GetOk("service_metro_area_network")
	handoverMetroArea, hasHandoverMetroArea := res.GetOk("handover_metro_area")
	handoverMetroAreaNetwork, hasHandoverMetroAreaNetwork := res.GetOk("handover_metro_area_network")
	downgradeAllowed, hasDowngradeAllowed := res.GetOk("downgrade_allowed")
	upgradeAllowed, hasUpgradeAllowed := res.GetOk("upgrade_allowed")
	physicalPortSpeed, hasPhysicalPortSpeed := res.GetOk("physical_port_speed")
	bandwidth, hasBandwidth := res.GetOk("bandwidth")

	// Query
	if hasName {
		qry.Name = name.(string)
	}
	if hasServiceProvider {
		qry.ServiceProvider = serviceProvider.(string)
	}
	if hasServiceMetroArea {
		qry.ServiceMetroArea = serviceMetroArea.(string)
	}
	if hasServiceMetroAreaNetwork {
		qry.ServiceMetroAreaNetwork = serviceMetroAreaNetwork.(string)
	}
	if hasHandoverMetroArea {
		qry.HandoverMetroArea = handoverMetroArea.(string)
	}
	if hasHandoverMetroAreaNetwork {
		qry.HandoverMetroAreaNetwork = handoverMetroAreaNetwork.(string)
	}
	if hasDowngradeAllowed {
		if downgradeAllowed.(bool) {
			qry.DowngradeAllowed = "true"
		} else {
			qry.DowngradeAllowed = "false"
		}
	}
	if hasUpgradeAllowed {
		if upgradeAllowed.(bool) {
			qry.UpgradeAllowed = "true"
		} else {
			qry.UpgradeAllowed = "false"
		}
	}
	if hasPhysicalPortSpeed {
		qry.PhysicalPortSpeed = physicalPortSpeed.(int)
	}
	if hasBandwidth {
		qry.Bandwidth = bandwidth.(int)
	}

	return qry
}
