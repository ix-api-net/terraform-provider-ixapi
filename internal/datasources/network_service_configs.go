package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// Common query schema
func networkServiceConfigQuerySchema(result map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": schemas.DataSourceQuery(
			"Filter by id of account managing the network service config"),
		"consuming_account": schemas.DataSourceQuery(
			"Filter by id of the account consuming the network service config"),
		"external_ref": schemas.DataSourceQuery(
			"Filter by external reference"),
		"network_service": schemas.DataSourceQuery(
			"Filter by id of the network service"),
		"network_connection": schemas.DataSourceQuery(
			"Filter by id of the connection"),
		"network_service_configs": schemas.IntoDataSourceResultsSchema(result),
	}
}

// Query
func networkServiceConfigQuery(
	t string, res *schema.ResourceData,
) *ixapi.NetworkServiceConfigsListQuery {
	qry := &ixapi.NetworkServiceConfigsListQuery{
		Type: t,
	}

	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	networkService, hasNetworkService := res.GetOk("network_service")
	networkConnection, hasNetworkConnection := res.GetOk("network_connection")

	if hasManagingAccount {
		qry.ManagingAccount = managingAccount.(string)
	}
	if hasConsumingAccount {
		qry.ConsumingAccount = consumingAccount.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasNetworkService {
		qry.NetworkService = networkService.(string)
	}
	if hasNetworkConnection {
		qry.Connection = networkConnection.(string)
	}
	return qry
}
