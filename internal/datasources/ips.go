package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/crud"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ix-api/internal/schemas"
)

// NewIPsDataSource creates a new data source for retrieving
// IP addresses associated with a service.
func NewIPsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Get IP addresses associated with a network service (config) or network feature (config) with this data source",
		ReadContext: crud.Read(ipsRead),
		Schema: map[string]*schema.Schema{
			"managing_account": schemas.DataSourceQuery(
				"Filter by account managing the service or config"),
			"consuming_account": schemas.DataSourceQuery(
				"Filter by account using the ip addresses"),
			"external_ref": schemas.DataSourceQuery(
				"Filter by external reference"),
			"network_service": schemas.DataSourceQuery(
				"Filter by ID of the network service, see related data source(s)"),
			"network_service_config": schemas.DataSourceQuery(
				"Filter by ID of the network service config, e.g. in case of an exchange lan. See related data source(s)."),
			"network_feature": schemas.DataSourceQuery(
				"Filter by ID of the network feature"),
			"network_feature_config": schemas.DataSourceQuery(
				"Filter by ID of the network feature config"),
			"version": schemas.DataSourceQueryInt(
				"Filter by IP address version (4 or 6)"),

			"ips": schemas.IntoDataSourceResultsSchema(
				schemas.IPAddressSchema()),
		},
	}
}

// Create query from provided resource data
func ipsQuery(res *schema.ResourceData) *ixapi.IPsListQuery {
	qry := &ixapi.IPsListQuery{}

	managingAccount, hasManagingAccount := res.GetOk("managing_account")
	consumingAccount, hasConsumingAccount := res.GetOk("consuming_account")
	externalRef, hasExternalRef := res.GetOk("external_ref")
	networkService, hasNetworkService := res.GetOk("network_service")
	networkServiceConfig, hasNetworkServiceConfig := res.GetOk("network_service_config")
	networkFeature, hasNetworkFeature := res.GetOk("network_feature")
	networkFeatureConfig, hasNetworkFeatureConfig := res.GetOk("network_feature_config")

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
	if hasNetworkServiceConfig {
		qry.NetworkServiceConfig = networkServiceConfig.(string)
	}
	if hasNetworkFeature {
		qry.NetworkFeature = networkFeature.(string)
	}
	if hasNetworkFeatureConfig {
		qry.NetworkFeatureConfig = networkFeatureConfig.(string)
	}

	return qry
}

// Retrieve IPs
func fetchIPs(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.IPAddress, error) {
	qry := ipsQuery(res)
	result, err := api.IPsList(ctx, qry)
	if err != nil {
		return nil, err
	}

	// Custom filter
	version, hasVersion := res.GetOk("version")
	filtered := make([]*ixapi.IPAddress, 0, len(result))
	for _, ip := range result {
		if hasVersion && ip.Version != version.(int) {
			continue
		}
		filtered = append(filtered, ip)
	}

	return filtered, nil
}

// Fetch ip addresses
func ipsRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	results, err := fetchIPs(ctx, res, api)
	if err != nil {
		return err
	}

	ips, err := schemas.FlattenModels(results)
	if err != nil {
		return err
	}

	if err := res.Set("ips", ips); err != nil {
		return err
	}
	res.SetId(schemas.Timestamp())
	return nil
}

// NewIPDataSource creates a data source to retrieve a single IP
// address identified by ID.
func NewIPDataSource() *schema.Resource {
	ipSchema := schemas.Combine(
		schemas.IntoDataSourceSchema(schemas.IPAddressSchema()),
	)

	return &schema.Resource{
		Description: "Use the ixapi_ip data source to retrieve a single ip address, identified by ID",
		ReadContext: ipRead,
		Schema:      ipSchema,
	}
}

func ipRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)
	id := res.Get("id").(string)

	ip, err := api.IPsRead(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := schemas.SetResourceData(ip, res); err != nil {
		return diag.FromErr(err)
	}
	res.SetId(id)
	return nil
}
