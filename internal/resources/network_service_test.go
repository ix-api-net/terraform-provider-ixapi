package resources

import (
	"testing"

	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
)

func TestVLanConfigFromResourceDataPort(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	res := resource.Data(nil)

	vlanConfig := map[string]any{
		"vlan_type": "port",
	}
	res.Set("vlan_config", []any{vlanConfig})

	config, err := vlanConfigFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	c := config.(*ixapi.VLANConfigPort)
	if c.VLANType != "port" {
		t.Fatal("expected port vlan config type")
	}
}

func TestVLanConfigFromResourceDataDot1Q(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	res := resource.Data(nil)

	vlanConfig := map[string]any{
		"vlan_type": "dot1q",
		"vlan":      42,
	}
	res.Set("vlan_config", []any{vlanConfig})

	config, err := vlanConfigFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	c := config.(*ixapi.VLANConfigDot1Q)
	if c.VLANType != "dot1q" {
		t.Error("expected dot1q vlan config type")
	}
	if *c.VLAN != 42 {
		t.Error("unexpected vlan:", c)
	}
	if *c.VLANEthertype != "0x8100" {
		t.Error("unexpected ethertype:", c)
	}
}

func TestVLanConfigFromResourceDataQinQ(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	res := resource.Data(nil)

	vlanConfig := map[string]any{
		"vlan_type":            "qinq",
		"inner_vlan":           42,
		"outer_vlan":           23,
		"outer_vlan_ethertype": "0x2342",
	}
	res.Set("vlan_config", []any{vlanConfig})

	config, err := vlanConfigFromResourceData(res)
	if err != nil {
		t.Fatal(err)
	}

	c := config.(*ixapi.VLANConfigQinQ)
	if c.VLANType != "qinq" {
		t.Error("unexpected type:", c)
	}
	if c.InnerVLAN != 42 {
		t.Error("unexpected vlan:", c)
	}
	if *c.OuterVLAN != 23 {
		t.Error("unexpected vlan:", c)
	}
	if *c.OuterVLANEthertype != "0x2342" {
		t.Error("unexpected ethertype:", c)
	}
}
