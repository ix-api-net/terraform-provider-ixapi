package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/crud"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewMACResource creates an ixapi mac address resource
// for registering mac addresses
func NewMACResource() *schema.Resource {
	return &schema.Resource{
		Description: "Use the `ixapi_mac` resource to register a mac address",

		CreateContext: crud.Create(macCreate),
		UpdateContext: crud.Update(macUpdate),
		ReadContext:   crud.Read(macRead),
		DeleteContext: crud.Delete(macDelete),

		Schema: schemas.MacAddressSchema(),
	}
}

func macRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.MacAddressRequest, error) {
	res := schemas.ResourceData{ResourceData: r}
	vnb, err := res.GetTimeOpt("valid_not_before")
	if err != nil {
		return nil, err
	}
	vna, err := res.GetTimeOpt("valid_not_after")
	if err != nil {
		return nil, err
	}
	req := &ixapi.MacAddressRequest{
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		ExternalRef:      res.GetStringOpt("external_ref"),
		Address:          res.GetString("address"),
		ValidNotBefore:   vnb,
		ValidNotAfter:    vna,
	}
	return req, nil
}

func macCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := macRequestFromResourceData(res)
	if err != nil {
		return err
	}
	mac, err := api.MacsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(mac.ID)
	return macRead(ctx, res, api)
}

func macRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	mac, err := api.MacsRead(ctx, id)
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(mac, res); err != nil {
		return err
	}
	return nil
}

func macDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	_, err := api.MacsDestroy(ctx, id)
	if err != nil {
		return err
	}
	res.SetId("")
	return nil
}

func macUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	return fmt.Errorf("mac addresses can not be updated; you need to create a new resource")
}
