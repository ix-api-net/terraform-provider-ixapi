package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/crud"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// NewRoutingFunctionResource creates a routing function resource
func NewRoutingFunctionResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Use the `ixapi_routing_function` resource to create and manage a routing function.",
		CreateContext: crud.Create(rfCreate),
		UpdateContext: crud.Update(rfUpdate),
		ReadContext:   crud.Read(rfRead),
		DeleteContext: crud.Delete(rfDelete),

		Schema: schemas.RoutingFunctionSchema(),

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func rfRequestFromResourceData(
	r *schema.ResourceData,
) (*ixapi.RoutingFunctionRequest, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.RoutingFunctionRequest{
		ManagingAccount:  res.GetString("managing_account"),
		ConsumingAccount: res.GetString("consuming_account"),
		BillingAccount:   res.GetString("billing_account"),
		ProductOffering:  res.GetString("product_offering"),
		ASN:              res.GetInt("asn"),

		ExternalRef:   res.GetStringOpt("external_ref"),
		PurchaseOrder: res.GetStringOpt("purchase_order"),
		ContractRef:   res.GetStringOpt("contract_ref"),
		Capacity:      res.GetIntOpt("capacity"),
	}
	return req, nil
}

func rfPatchFromResourceData(
	r *schema.ResourceData,
) (*ixapi.RoutingFunctionPatch, error) {
	res := schemas.ResourceDataFrom(r)
	req := &ixapi.RoutingFunctionPatch{}

	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("consuming_account") {
		req.ConsumingAccount = res.GetStringOpt("consuming_account")
	}
	if res.HasChange("billing_account") {
		req.BillingAccount = res.GetStringOpt("billing_account")
	}
	if res.HasChange("product_offering") {
		req.ProductOffering = res.GetStringOpt("product_offering")
	}
	if res.HasChange("asn") {
		req.ASN = res.GetIntOpt("asn")
	}
	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("purchase_order") {
		req.PurchaseOrder = res.GetStringOpt("purchase_order")
	}
	if res.HasChange("contract_ref") {
		req.ContractRef = res.GetStringOpt("contract_ref")
	}
	if res.HasChange("capacity") {
		req.Capacity = res.GetIntOpt("capacity")
	}
	return req, nil
}

func rfCreate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	req, err := rfRequestFromResourceData(res)
	if err != nil {
		return err
	}
	rf, err := api.RoutingFunctionsCreate(ctx, req)
	if err != nil {
		return err
	}
	res.SetId(rf.ID)
	return rfRead(ctx, res, api)
}

func rfRead(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	rf, err := api.RoutingFunctionsRead(ctx, id)
	if err != nil && ixapi.IsErrNotFound(err) {
		res.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	if err := schemas.SetResourceData(rf, res); err != nil {
		return err
	}
	return nil
}

func rfUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	patch, err := rfPatchFromResourceData(res)
	if err != nil {
		return err
	}
	_, err = api.RoutingFunctionsPatch(ctx, id, patch)
	if err != nil {
		return err
	}
	return rfRead(ctx, res, api)
}

func rfDelete(
	ctx context.Context,
	res *schema.ResourceData,
	api *ixapi.Client,
) error {
	id := res.Id()
	req := &ixapi.CancellationRequest{}
	_, err := api.RoutingFunctionsDestroy(ctx, id, req)
	if err != nil {
		return err
	}
	return rfRead(ctx, res, api)
}
