package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewNetworkFeatureConfigRouteServerResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkFeatureConfigRouteServerResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_nfc_route_server` resource to configure access to a route server in the exchange lan.",

		CreateContext: crud.Create(nfcRouteServerCreate),
		UpdateContext: crud.Update(nfcRouteServerUpdate),
		ReadContext:   crud.Read(nfcRouteServerRead),
		DeleteContext: crud.Delete(nfcRouteServerDelete),

		Schema: schemas.RouteServerNetworkServiceConfigSchema(),
	}
}

func nfcRouteServerCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {

	return nil
}

func nfcRouteServerRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

func nfcRouteServerUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

func nfcRouteServerDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}
