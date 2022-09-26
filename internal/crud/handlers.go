package crud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
)

// HandlerFunc implements the IX-API interaction
type HandlerFunc func(context.Context, *schema.ResourceData, *ixapi.Client) error

// Create makes a CreateContext handler func
func Create(handler HandlerFunc) schema.CreateContextFunc {
	return func(
		ctx context.Context,
		res *schema.ResourceData,
		meta any,
	) diag.Diagnostics {
		api := meta.(*ixapi.Client)
		if err := handler(ctx, res, api); err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
}

// Read makes a ReadContext handler func
func Read(handler HandlerFunc) schema.ReadContextFunc {
	return func(
		ctx context.Context,
		res *schema.ResourceData,
		meta any,
	) diag.Diagnostics {
		api := meta.(*ixapi.Client)
		if err := handler(ctx, res, api); err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
}

// Update maks a UpdateContext handler func
func Update(handler HandlerFunc) schema.UpdateContextFunc {
	return func(
		ctx context.Context,
		res *schema.ResourceData,
		meta any,
	) diag.Diagnostics {
		api := meta.(*ixapi.Client)
		if err := handler(ctx, res, api); err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
}

// Delete maks a UpdateContext handler func
func Delete(handler HandlerFunc) schema.DeleteContextFunc {
	return func(
		ctx context.Context,
		res *schema.ResourceData,
		meta any,
	) diag.Diagnostics {
		api := meta.(*ixapi.Client)
		if err := handler(ctx, res, api); err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
}
