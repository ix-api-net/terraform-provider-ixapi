package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func NewCloudRouterResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_de_cix_cloud_router` resource to create and manage a DE-CIX Cloud ROUTER (VRF) instance.",
		CreateContext: crud.Create(cloudRouterCreate),
		ReadContext:   crud.Read(cloudRouterRead),
		DeleteContext: crud.Delete(cloudRouterDelete),

		Schema: schemas.CloudRouterSchema(),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func cloudRouterRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.CloudRouterRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.CloudRouterRequest{
		ManagingAccount:  res.GetString("managing_account"),
		BillingAccount:   res.GetString("billing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		ExternalRef:      res.GetStringOpt("external_ref"),
		ProductOffering:  res.GetString("product_offering"),
		ASN:              res.GetInt("asn"),
		Capacity:         res.GetInt("capacity"),
	}
	return req, nil
}

func cloudRouterCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	req, err := cloudRouterRequestFromResourceData(res)
	if err != nil {
		return err
	}

	cr, err := api.CloudRoutersCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(cr.ID)

	return cloudRouterRead(ctx, res, api)
}

func cloudRouterRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	cr, err := api.CloudRoutersRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	if err := schemas.SetResourceData(cr, res); err != nil {
		return err
	}
	return nil
}

func cloudRouterDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	if err := api.RequireCloudRouterExtension(); err != nil {
		return err
	}

	id := res.Id()
	err := api.CloudRoutersDestroy(ctx, id)
	if err != nil && !ixapi.IsErrNotFound(err) {
		return err
	}
	return cloudRouterRead(ctx, res, api)
}
