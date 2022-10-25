package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// Common query schema
func networkServiceQuerySchema(result map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": schemas.DataSourceQuery(
			"Filter by id of account managing the network service"),
		"consuming_account": schemas.DataSourceQuery(
			"Filter by id of the account consuming the network service"),
		"external_ref": schemas.DataSourceQuery(
			"Filter by external reference"),
		"pop": schemas.DataSourceQuery(
			"Filter network services by presence at the pop (id). See point of presence data source."),
		"product_offering": schemas.DataSourceQuery(
			"Filter network services by prouduct offering id"),
		"network_services": schemas.IntoDataSourceResultsSchema(result),
	}
}

// Query
func networkServiceQuery(
	t string, res *schema.ResourceData,
) *ixapi.NetworkServicesListQuery {
	qry := &ixapi.NetworkServicesListQuery{
		Type: t,
	}

	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	pop, hasPop := res.GetOk("pop")
	productOffering, hasProductOffering := res.GetOk("product_offering")

	if hasManagingAccount {
		qry.ManagingAccount = managingAccount.(string)
	}
	if hasConsumingAccount {
		qry.ConsumingAccount = consumingAccount.(string)
	}
	if hasExternalRef {
		qry.ExternalRef = externalRef.(string)
	}
	if hasPop {
		qry.Pop = pop.(string)
	}
	if hasProductOffering {
		qry.ProductOffering = productOffering.(string)
	}
	return qry
}
