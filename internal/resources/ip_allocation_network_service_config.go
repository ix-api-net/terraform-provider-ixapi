package resources

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewIPAllocationNetworkServiceConfigResource creates a resource
// for managing allocated IP addresses.
func NewIPAllocationNetworkServiceConfigResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use this resource to reference IP addresses allocated by a referanced network service. The managed IP addresses will be updated with the fqdn provided.",
		CreateContext: crud.Create(ipAllocationCreate),
		UpdateContext: crud.Update(ipAllocationUpdate),
		ReadContext:   crud.Read(ipAllocationRead),
		DeleteContext: crud.Delete(ipAllocationDelete),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fqdn": {
				Description: "Set the fqdn of the allocated IP addresses",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"version": {
				Description: "Only get IP addresses with this version (4 or 6)",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"network_service_config": {
				Description: "Get IP addresses allocated for this network service config",
				Type:        schema.TypeString,
				Required:    true,
			},
			"timeout": {
				Description: "Timeout for awaiting the allocation",
				Default:     600,
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"ips": schemas.IntoDataSourceResultsSchema(
				schemas.IPAddressSchema()),
		},
	}
}

// Retrieve IPs
func fetchIPs(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) ([]*ixapi.IPAddress, error) {
	qry := &ixapi.IPsListQuery{
		NetworkServiceConfig: res.Get("network_service_config").(string),
	}

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

func maybeSetFQDN(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
	ip *ixapi.IPAddress,
) error {
	fqdnOpt, hasFQDN := res.GetOk("fqdn")
	if !hasFQDN {
		return nil // nothing to do here
	}
	fqdn := fqdnOpt.(string)
	patch := &ixapi.IPAddressPatch{
		FQDN: &fqdn,
	}
	_, err := api.IPsPatch(ctx, ip.ID, patch)
	if err != nil {
		return err
	}
	return nil
}

func ipAllocationCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	timeout := time.Duration(res.Get("timeout").(int)) * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	nscID := res.Get("network_service_config").(string)

	// Await IP address allocation
	var ips []*ixapi.IPAddress
	for {
		// Check if the context is still valid befor making the API call
		if err := ctx.Err(); err != nil {
			return err
		}
		result, err := fetchIPs(ctx, res, api)
		if err != nil {
			return err
		}
		if len(result) > 0 {
			ips = result
			break
		}
		time.Sleep(time.Second)
	}

	// Set the FQDN for each IP address
	for _, ip := range ips {
		if err := maybeSetFQDN(ctx, res, api, ip); err != nil {
			return nil
		}
	}
	flat, err := schemas.FlattenModels(ips)
	if err != nil {
		return nil
	}
	if err := res.Set("ips", flat); err != nil {
		return err
	}

	// Use the network service config ID
	res.SetId(nscID)

	// Refresh
	return ipAllocationRead(ctx, res, api)
}

func ipAllocationRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	// Refresh IP address for network service config
	nscID := res.Id()

	// Fetch IPs
	ips, err := api.IPsList(ctx, &ixapi.IPsListQuery{
		NetworkServiceConfig: nscID,
	})

	if len(ips) == 0 {
		// This is gone
		res.SetId("")
		return nil
	}

	flat, err := schemas.FlattenModels(ips)
	if err != nil {
		return nil
	}
	if err := res.Set("ips", flat); err != nil {
		return err
	}

	return nil
}

func ipAllocationUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if res.HasChange("network_service_config") {
		// This is basically a new create
		return ipAllocationCreate(ctx, res, api)
	}

	// Only update if fqdn changed
	updates := []*ixapi.IPAddress{}
	if res.HasChange("fqdn") {
		// Update FQDN for ip addresses
		ips := res.Get("ips").([]any)
		for _, ip := range ips {
			addr := ip.(map[string]any)
			updates = append(updates, &ixapi.IPAddress{
				ID: addr["id"].(string),
			})
		}
	}
	// Apply updates
	for _, ip := range updates {
		if err := maybeSetFQDN(ctx, res, api, ip); err != nil {
			return err
		}
	}

	return ipAllocationRead(ctx, res, api)
}

func ipAllocationDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	// We can just drop this resource
	res.SetId("")
	return nil
}
