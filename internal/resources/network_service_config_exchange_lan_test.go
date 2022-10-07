package resources

import (
	"encoding/json"
	"testing"

	"gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"
	"gitlab.com/ix-api/ix-api-terraform-provider/internal/schemas"
)

// Polymorphics

func TestSetResourceDataPolymorphicVLAN(t *testing.T) {
	resource := NewNetworkServiceConfigExchangeLanResource()
	data := resource.Data(nil)

	vid := 23
	et := "0x8100"

	decom, _ := ixapi.ParseDate("2022-11-02")
	nsc := ixapi.ExchangeLanNetworkServiceConfig{
		Type:             "exchange_lan",
		State:            "requested",
		Status:           []*ixapi.Status{},
		ID:               "42",
		NetworkService:   "23",
		ManagingAccount:  "2342",
		ConsumingAccount: "4242",
		BillingAccount:   "1234",
		RoleAssignments:  []string{"23", "42"},
		Connection:       "12356",
		ASNs:             []int{2084242132},
		Macs:             []string{"42"},
		IPs:              []string{},
		Listed:           true,
		ProductOffering:  "123",
		DecommissionAt:   &decom,
		VLANConfigRaw:    json.RawMessage("foo"),
		VLANConfig: &ixapi.VLANConfigDot1Q{
			VLAN:          &vid,
			VLANType:      "dot1q",
			VLANEthertype: &et,
		},
	}

	err := schemas.SetResourceData(nsc, data)
	if err != nil {
		t.Fatal(err)
	}

	// Check embedded VLAN config
	v := data.Get("vlan_config").([]any)
	vlc := v[0].(map[string]any)

	if vlc["vlan_type"].(string) != "dot1q" {
		t.Log("unexpected vlan config:", vlc)
	}
}
