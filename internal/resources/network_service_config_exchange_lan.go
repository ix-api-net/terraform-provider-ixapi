package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewNetworkServiceConfigExchangeLanResource creates a NSC resource
// for creating an exchange lan network service configuration.
func NewNetworkServiceConfigExchangeLanResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_nsc_exchange_lan` resource to provision an access to an exchange lan for a customer.",

		CreateContext: crud.Create(nscExchangeLanCreate),
		UpdateContext: crud.Update(nscExchangeLanUpdate),
		ReadContext:   crud.Read(nscExchangeLanRead),
		DeleteContext: crud.Delete(nscExchangeLanDelete),

		Schema: schemas.ExchangeLanNetworkServiceConfigSchema(),
	}
}

func nscExchangeLanCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {

	return nil
}

func nscExchangeLanRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

func nscExchangeLanUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}

func nscExchangeLanDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return nil
}
