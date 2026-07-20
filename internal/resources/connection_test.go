package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/gocty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		ID:               "42",
		ManagingAccount:  "2342",
		ConsumingAccount: "4242",
		BillingAccount:   "1234",
		Mode:             "standalone",
		ProductOffering:  "123",
		Status:           []*ixapi.Status{},
		RoleAssignments:  []string{},
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

func TestConnectionImport(t *testing.T) {
	t.Run("extension enabled", func(t *testing.T) {
		resource := NewConnectionResource()
		data := resource.Data(nil)
		data.SetId("42")

		api := ixapi.NewTestClient(nil)
		api.CloudRouterEnabled = true

		results, err := connectionImport(context.Background(), data, api)
		if err != nil {
			t.Fatal(err)
		}
		if len(results) != 1 {
			t.Fatalf("expected 1 result, got %d", len(results))
		}
		if results[0].Id() != "42" {
			t.Errorf("unexpected id: %s", results[0].Id())
		}
	})

	t.Run("extension disabled", func(t *testing.T) {
		resource := NewConnectionResource()
		data := resource.Data(nil)
		data.SetId("42")

		api := ixapi.NewTestClient(nil)

		_, err := connectionImport(context.Background(), data, api)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if !strings.Contains(err.Error(), "CloudRouter extension is not enabled") {
			t.Errorf("unexpected error: %s", err)
		}
	})
}

// rawConfigInstanceState builds an InstanceState carrying a RawConfig cty.Value,
// which is what CustomizeDiff's GetRawConfig() reads from.
func rawConfigInstanceState(
	t *testing.T,
	res *schema.Resource,
	values map[string]interface{},
) *terraform.InstanceState {
	implType := res.CoreConfigSchema().ImpliedType()
	attrs := map[string]cty.Value{}
	for name, attrType := range implType.AttributeTypes() {
		v, ok := values[name]
		if !ok {
			attrs[name] = cty.NullVal(attrType)
			continue
		}
		val, err := gocty.ToCtyValue(v, attrType)
		if err != nil {
			t.Fatalf("building raw config for %q: %v", name, err)
		}
		attrs[name] = val
	}
	return &terraform.InstanceState{RawConfig: cty.ObjectVal(attrs)}
}

func TestConnectionCustomizeDiff(t *testing.T) {
	baseValues := map[string]interface{}{
		"managing_account":  "2342",
		"consuming_account": "4242",
		"billing_account":   "1234",
		"mode":              "standalone",
		"product_offering":  "123",
		"port_quantity":     2,
		"role_assignments":  []interface{}{},
	}

	tests := []struct {
		name             string
		extensionEnabled bool
		setDiscoverable  bool
		expectErr        bool
	}{
		{
			name:             "disabled with discoverable set",
			extensionEnabled: false,
			setDiscoverable:  true,
			expectErr:        true,
		},
		{
			name:             "disabled without discoverable",
			extensionEnabled: false,
			setDiscoverable:  false,
			expectErr:        false,
		},
		{
			name:             "enabled with discoverable set",
			extensionEnabled: true,
			setDiscoverable:  true,
			expectErr:        false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resource := NewConnectionResource()

			values := map[string]interface{}{}
			for k, v := range baseValues {
				values[k] = v
			}
			if tc.setDiscoverable {
				values["discoverable"] = true
			}

			conf := terraform.NewResourceConfigRaw(values)
			state := rawConfigInstanceState(t, resource, values)

			api := ixapi.NewTestClient(nil)
			api.CloudRouterEnabled = tc.extensionEnabled

			_, err := resource.Diff(context.Background(), state, conf, api)
			if tc.expectErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				if !strings.Contains(err.Error(), "CloudRouter extension is not enabled") {
					t.Errorf("unexpected error: %s", err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
		})
	}
}
