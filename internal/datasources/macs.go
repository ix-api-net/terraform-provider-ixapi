package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewMacsDataSource creates a data source for querying mac addresses
func NewMacsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the ixapi_macs data source to query mac addresses",
		ReadContext: crud.Read(macsRead),
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Find mac addresses managed by this account id"),
			"consuming_account": schemas.DataSourceQuery(
				"Find mac addresses used by this account id"),
			"external_ref": schemas.DataSourceQuery(
				"Find mac addresses with this external_ref"),
			"network_service_config": schemas.DataSourceQuery(
				"Find mac addresses in use by a network service config"),
			"address": schemas.DataSourceQuery(
				"Find mac addresses with this value"),
			"macs": schemas.IntoDataSourceResultsSchema(
				schemas.MacAddressSchema()),
		},
	}
}

// make mac addresses query
func macAddressesQuery(
	res *schema.ResourceData,
) *ixapi.MacsListQuery {
	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	nsc, hasNsc := res.GetOk("network_service_config")
	address, hasAddress := res.GetOk("address")

	qry := &ixapi.MacsListQuery{}
	if hasManagingAccount {
		qry.ManagingAccount = managingAccount.(string)
	}
	if hasConsumingAccount {
		qry.ConsumingAccount = consumingAccount.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasNsc {
		qry.NetworkServiceConfig = nsc.(string)
	}
	if hasAddress {
		qry.Address = address.(string)
	}

	return qry
}

// Read macs resource
func macsRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	qry := macAddressesQuery(res)
	macs, err := api.MacsList(ctx, qry)
	if err != nil {
		return err
	}

	flat, err := schemas.FlattenModels(macs)
	if err != nil {
		return err
	}
	if err := res.Set("macs", flat); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewMacDataSource creates a data source for a single mac address
func NewMacDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to reference a single mac by address, external ref or id",
		Schema: schemas.IntoDataSourceSchema(
			schemas.MacAddressSchema()),
		ReadContext: crud.Read(macRead),
	}
}

// Fetch a single mac address
func macRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id, hasID := res.GetOk("id")

	var mac *ixapi.MacAddress
	if hasID {
		result, err := api.MacsRead(ctx, id.(string))
		if err != nil {
			return err
		}
		mac = result
	} else {
		qry := macAddressesQuery(res)
		results, err := api.MacsList(ctx, qry)
		if err != nil {
			return err
		}
		if len(results) == 0 {
			return fmt.Errorf("no such mac address could be found")
		}
		if len(results) > 1 {
			return fmt.Errorf("the mac address could not be uniquely identified")
		}
		mac = results[0]
	}

	if err := schemas.SetResourceData(mac, res); err != nil {
		return err
	}
	res.SetId(mac.ID)
	return nil
}
