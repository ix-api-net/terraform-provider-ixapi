package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterArpTableDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `cloud_router_arp_table` data source to list ARP entries for a cloud router VRF",
		ReadContext: cloudRouterArpTableRead,
		Schema: map[string]*schema.Schema{
			"vrf": schemas.DataSourceQuery(
				"Filter by VRF ID"),
			"network_service_config": schemas.DataSourceQuery(
				"Filter by network service config ID"),
			"arp_entries": schemas.IntoDataSourceResultsSchema(
				schemas.ArpEntrySchema(),
			),
		},
	}
}

func cloudRouterArpTableRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	if err := api.RequireCloudRouterExtension(); err != nil {
		return diag.FromErr(err)
	}

	vrfID := ""
	if v, ok := res.GetOk("vrf"); ok {
		vrfID = v.(string)
	}
	nscID := ""
	if v, ok := res.GetOk("network_service_config"); ok {
		nscID = v.(string)
	}

	all, err := api.ArpTableList(ctx, vrfID, nscID)
	if err != nil {
		return diag.FromErr(err)
	}

	flat, err := schemas.FlattenModels(all)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := res.Set("arp_entries", flat); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(schemas.Timestamp())

	return nil
}
