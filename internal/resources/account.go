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

		Schema: schemas.AccountSchema,
	}
}

// Operations

// accountCreate creates a new account
func accountCreate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	// Make request from terraform state
	req, diags := schemas.AccountRequestFromResourceData(res)
	if diags != nil {
		return diags
	}

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
	schemas.AccountSetResourceData(acc, res)

	return nil
}

// accountUpdate updates the account
func accountUpdate(
	ctx context.Context,
	res *schema.ResourceData,
	meta any,
) diag.Diagnostics {
	api := meta.(*ixapi.Client)

	patch, diags := schemas.AccountPatchFromResourceData(res)
	if diags != nil {
		return diags
	}

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
