package ixapi

import (
	"encoding/json"
	"testing"
)

func TestDecodeVLANConfigUnknown(t *testing.T) {
	_, err := decodeVLANConfig([]byte(`{"vlan_type": "unknown"}`))
	if err == nil {
		t.Fatal("expected error for unknown vlan_type")
	}
}

func TestUnmarshalProductOffering(t *testing.T) {
	type testPO struct {
		AvailableUntil *APITimestamp `json:"available_until"`
	}
	data := []byte(`{"available_until": "2024-01-15"}`)
	var po testPO
	if err := UnmarshalProductOffering(data, &po); err != nil {
		t.Fatal(err)
	}
	if po.AvailableUntil == nil {
		t.Fatal("expected non-nil AvailableUntil")
	}
	ts := po.AvailableUntil.Time
	if ts.Year() != 2024 || ts.Month() != 1 || ts.Day() != 15 {
		t.Errorf("unexpected date: %v", ts)
	}
}

func TestDecodeVLANConfigPort(t *testing.T) {
	data := json.RawMessage(`{"vlan_type": "port"}`)
	cfg, err := decodeVLANConfig(data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cfg.(*VLANConfigPort); !ok {
		t.Error("unexpected vlan type:", cfg)
	}
	t.Log(cfg.PolymorphicType())
}

func TestDecodeVLANConfigDot1Q(t *testing.T) {
	data := json.RawMessage(`{
		"vlan_type": "dot1q",
		"vlan": 42,
		"vlan_ethertype": "0x8100"
	}`)
	cfg, err := decodeVLANConfig(data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cfg.(*VLANConfigDot1Q); !ok {
		t.Error("unexpected vlan type:", cfg)
	}
	if *cfg.(*VLANConfigDot1Q).VLANEthertype != "0x8100" {
		t.Error("unexpected ethertype in:", cfg)
	}
	t.Log(cfg.PolymorphicType())
}

func TestDecodeVLANConfigQinQ(t *testing.T) {
	data := json.RawMessage(`{
		"vlan_type": "qinq",
		"outer_vlan": 42,
		"inner_vlan": 23,
		"outer_vlan_ethertype": "0x8100"
	}`)
	cfg, err := decodeVLANConfig(data)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := cfg.(*VLANConfigQinQ); !ok {
		t.Error("unexpected vlan type:", cfg)
	}
	if *cfg.(*VLANConfigQinQ).OuterVLANEthertype != "0x8100" {
		t.Error("unexpected ethertype in:", cfg)
	}
	t.Log(cfg.PolymorphicType())
}
