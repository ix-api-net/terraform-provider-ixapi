package resources

import (
	"context"
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestConnectionRequestFromResourceData(t *testing.T) {
	resource := NewConnectionResource()
	data := resource.Data(nil)
	_ = data.Set("managing_account", "2342")
	_ = data.Set("consuming_account", "4242")
	_ = data.Set("mode", "standalone")
	_ = data.Set("port_quantity", 2)

	req, err := connectionRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.ConsumingAccount != "4242" {
		t.Errorf("unexpected consuming_account: %s", req.ConsumingAccount)
	}
	if req.Mode != "standalone" {
		t.Errorf("unexpected mode: %s", req.Mode)
	}
	if req.PortQuantity != 2 {
		t.Errorf("unexpected port_quantity: %d", req.PortQuantity)
	}
}

func TestConnectionPatchFromResourceDataNoChange(t *testing.T) {
	resource := NewConnectionResource()
	data := resource.Data(nil)

	patch, err := connectionPatchFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if patch != nil {
		t.Fatal("expected nil patch when nothing changed")
	}
}

func TestConnectionRead(t *testing.T) {
	resource := NewConnectionResource()
	data := resource.Data(nil)
	data.SetId("42")

	conn := &ixapi.Connection{
		ID:              "42",
		ManagingAccount: "2342",
		ConsumingAccount: "4242",
		BillingAccount:  "1234",
		Mode:            "standalone",
		ProductOffering: "123",
		Status:          []*ixapi.Status{},
		RoleAssignments: []string{},
	}
	api := ixapi.NewTestClient(map[string]any{
		"/connections/42": conn,
	})

	if err := connectionRead(context.Background(), data, api); err != nil {
		t.Fatal(err)
	}
	if data.Get("managing_account").(string) != "2342" {
		t.Errorf("unexpected managing_account: %s", data.Get("managing_account"))
	}
	if data.Get("mode").(string) != "standalone" {
		t.Errorf("unexpected mode: %s", data.Get("mode"))
	}
}
