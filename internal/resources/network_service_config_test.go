package resources

import (
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

func TestVlanConfigFromResourceDataNil(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	data := resource.Data(nil)

	cfg, err := VlanConfigFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if cfg != nil {
		t.Fatalf("expected nil vlan config, got %v", cfg)
	}
}

func TestVlanConfigFromResourceDataPort(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	data := resource.Data(nil)

	nsc := ixapi.ExchangeLanNetworkServiceConfig{
		Status:          []*ixapi.Status{},
		ASNs:            []int{},
		Macs:            []string{},
		IPs:             []string{},
		RoleAssignments: []string{},
		VLANConfig:      &ixapi.VLANConfigPort{VLANType: "port"},
	}
	if err := schemas.SetResourceData(nsc, data); err != nil {
		t.Fatal(err)
	}

	cfg, err := VlanConfigFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cfg.(*ixapi.VLANConfigPort); !ok {
		t.Fatalf("expected *ixapi.VLANConfigPort, got %T", cfg)
	}
}

func TestNscExchangeLanRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	data := resource.Data(nil)
	data.Set("managing_account", "2342")

	req, err := nscExchangeLanRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.Type != ixapi.ExchangeLanNetworkServiceConfigType {
		t.Errorf("unexpected type: %s", req.Type)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.VLANConfig != nil {
		t.Fatalf("expected nil vlan config, got %v", req.VLANConfig)
	}
}

func TestNscP2PVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceConfigP2PVCResource()
	data := resource.Data(nil)
	data.Set("managing_account", "2342")

	req, err := nscP2PVCRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.Type != ixapi.P2PNetworkServiceConfigType {
		t.Errorf("unexpected type: %s", req.Type)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.VLANConfig != nil {
		t.Fatalf("expected nil vlan config, got %v", req.VLANConfig)
	}
}

func TestNscCloudVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceConfigCloudVCResource()
	data := resource.Data(nil)
	data.Set("managing_account", "2342")

	req, err := nscCloudVCRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.Type != ixapi.CloudNetworkServiceConfigType {
		t.Errorf("unexpected type: %s", req.Type)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.VLANConfig != nil {
		t.Fatalf("expected nil vlan config, got %v", req.VLANConfig)
	}
}

func TestNscMP2MPVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceConfigMP2MPVCResource()
	data := resource.Data(nil)
	data.Set("managing_account", "2342")

	req, err := nscMP2MPVCRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.Type != ixapi.MP2MPNetworkServiceConfigType {
		t.Errorf("unexpected type: %s", req.Type)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.VLANConfig != nil {
		t.Fatalf("expected nil vlan config, got %v", req.VLANConfig)
	}
}

func TestNscP2MPVCRequestFromResourceData(t *testing.T) {
	resource := NewNetworkServiceConfigP2MPVCResource()
	data := resource.Data(nil)
	data.Set("managing_account", "2342")

	req, err := nscP2MPVCRequestFromResourceData(data)
	if err != nil {
		t.Fatal(err)
	}
	if req.Type != ixapi.P2MPNetworkServiceConfigType {
		t.Errorf("unexpected type: %s", req.Type)
	}
	if req.ManagingAccount != "2342" {
		t.Errorf("unexpected managing_account: %s", req.ManagingAccount)
	}
	if req.VLANConfig != nil {
		t.Fatalf("expected nil vlan config, got %v", req.VLANConfig)
	}
}
