package resources

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// NewAccountResource creates a new account resource
func NewAccountResource() *schema.Resource {
	return &schema.Resource{
		Description: "An IX-API account",

		CreateContext: accountCreate,
		ReadContext:   accountRead,
		UpdateContext: accountUpdate,
		DeleteContext: accountDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: schemas.AccountSchema(),
	}
}

// Operations

// addressFromResource decodes embedded address infos
func addressFromResource(res schemas.Resource) *ixapi.Address {
	if res == nil {
		return nil
	}
	return &ixapi.Address{
		Country:             res.GetString("country"),
		Locality:            res.GetString("locality"),
		Region:              res.GetStringOpt("region"),
		PostalCode:          res.GetString("postal_code"),
		StreetAddress:       res.GetString("street_address"),
		PostOfficeBoxNumber: res.GetStringOpt("post_office_box_number"),
	}
}

// billingInformationFromResource decodes billing information
// from an embedded resource
func billingInformationFromResource(res schemas.Resource) *ixapi.BillingInformation {
	if res == nil {
		return nil
	}
	address := addressFromResource(res.GetResource("address"))
	return &ixapi.BillingInformation{
		Name:      res.GetString("name"),
		VatNumber: res.GetStringOpt("vat_number"),
		Address:   address,
	}
}

// accountRequestFromResourceData builds an account create request
func accountRequestFromResourceData(
	r *schema.ResourceData,
) *ixapi.AccountRequest {
	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.AccountRequest{
		ManagingAccount: res.GetStringOpt("managing_account"),
		Name:            res.GetString("name"),
		LegalName:       res.GetStringOpt("legal_name"),
		BillingInformation: billingInformationFromResource(
			res.GetResource("billing_information")),
		ExternalRef:  res.GetStringOpt("external_ref"),
		Discoverable: res.GetBoolOpt("discoverable"),
		Address: addressFromResource(
			res.GetResource("address")),
	}
	return req
}

// accountPatchFromResourceData makes a new Patch payload
func accountPatchFromResourceData(
	r *schema.ResourceData,
) *ixapi.AccountPatch {
	res := schemas.ResourceData{ResourceData: r}
	req := &ixapi.AccountPatch{}
	if res.HasChange("managing_account") {
		req.ManagingAccount = res.GetStringOpt("managing_account")
	}
	if res.HasChange("name") {
		req.Name = res.GetStringOpt("name")
	}
	if res.HasChange("legal_name") {
		req.LegalName = res.GetStringOpt("legal_name")
	}
	if res.HasChange("billing_information") {
		req.BillingInformation = billingInformationFromResource(
			res.GetResource("billing_information"))
	}
	if res.HasChange("external_ref") {
		req.ExternalRef = res.GetStringOpt("external_ref")
	}
	if res.HasChange("discoverable") {
		req.Discoverable = res.GetBoolOpt("discoverable")
	}
	if res.HasChange("address") {
		req.Address = addressFromResource(res.GetResource("address"))
	}
	return req
}

// accountCreate creates a new account
func accountCreate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Make request from terraform state
	req := accountRequestFromResourceData(res)

	// Call api, set ID and fetch account
	acc, err := api.AccountsCreate(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	res.SetId(acc.ID)

	return accountRead(ctx, res, meta)
}

// accountCreate fetches the the account from the API
func accountRead(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	acc, err := api.AccountsRead(ctx, res.Id())
	var notFoundErr *ixapi.NotFoundError
	if err != nil && errors.As(err, &notFoundErr) {
		// The ID is not longer available, so we remove it
		res.SetId("")
		return nil
	} else if err != nil {
		return diag.FromErr(err)
	}

	// Update resource
	schemas.SetResourceData(acc, res)

	return nil
}

// accountUpdate updates the account
func accountUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	patch := accountPatchFromResourceData(res)
	_, err := api.AccountsPatch(ctx, res.Id(), patch)
	if err != nil {
		return diag.FromErr(err)
	}

	return accountRead(ctx, res, meta)
}

// accountDelete requests the deletion of the account
func accountDelete(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	_, err := api.AccountsDestroy(ctx, res.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return accountRead(ctx, res, meta)
}
