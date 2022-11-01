package crud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestCreate(t *testing.T) {
	handler := func(
		ctx context.Context,
		res *schema.ResourceData,
		api *ixapi.Client,
	) error {
		return fmt.Errorf("handler error")
	}
	res := &schema.Resource{
		CreateContext: Create(handler),
	}
	diags := res.CreateContext(
		context.Background(),
		nil,
		&ixapi.Client{})

	if !diags.HasError() {
		t.Fatal("expected error result")
	}
	if diags[0].Summary != "handler error" {
		t.Error("unexpected diags:", diags)
	}
}

func TestRead(t *testing.T) {
	handler := func(
		ctx context.Context,
		res *schema.ResourceData,
		api *ixapi.Client,
	) error {
		return fmt.Errorf("handler error")
	}
	res := &schema.Resource{
		ReadContext: Read(handler),
	}
	diags := res.ReadContext(
		context.Background(),
		nil,
		&ixapi.Client{})

	if !diags.HasError() {
		t.Fatal("expected error result")
	}
	if diags[0].Summary != "handler error" {
		t.Error("unexpected diags:", diags)
	}
}

func TestUpdate(t *testing.T) {
	handler := func(
		ctx context.Context,
		res *schema.ResourceData,
		api *ixapi.Client,
	) error {
		return fmt.Errorf("handler error")
	}
	res := &schema.Resource{
		UpdateContext: Update(handler),
	}
	diags := res.UpdateContext(
		context.Background(),
		nil,
		&ixapi.Client{})

	if !diags.HasError() {
		t.Fatal("expected error result")
	}
	if diags[0].Summary != "handler error" {
		t.Error("unexpected diags:", diags)
	}
}

func TestDelete(t *testing.T) {
	handler := func(
		ctx context.Context,
		res *schema.ResourceData,
		api *ixapi.Client,
	) error {
		return fmt.Errorf("handler error")
	}
	res := &schema.Resource{
		DeleteContext: Delete(handler),
	}
	diags := res.DeleteContext(
		context.Background(),
		nil,
		&ixapi.Client{})

	if !diags.HasError() {
		t.Fatal("expected error result")
	}
	if diags[0].Summary != "handler error" {
		t.Error("unexpected diags:", diags)
	}
}
