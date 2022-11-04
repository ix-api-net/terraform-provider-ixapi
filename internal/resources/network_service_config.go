package resources

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// Create polymorphic VLanConfig from resource data
func vlanConfigFromResourceData(r *schema.ResourceData) (ixapi.VLANConfig, error) {
	res := schemas.ResourceData{ResourceData: r}
	c := res.GetResource("vlan_config")
	vType := c["vlan_type"].(string)

	if vType == "port" {
		cfg := &ixapi.VLANConfigPort{
			VLANType: "port",
		}
		return cfg, nil
	}
	if vType == "dot1q" {
		vlan := c.GetIntOpt("vlan")
		ethertype := c.GetStringOptDefault("vlan_ethertype", "0x8100")
		cfg := &ixapi.VLANConfigDot1Q{
			VLANType:      "dot1q",
			VLAN:          vlan,
			VLANEthertype: ethertype,
		}
		return cfg, nil
	}
	if vType == "qinq" {
		outerVlanEthertype := c.GetStringOptDefault("outer_vlan_ethertype", "0x8100")
		outerVlan := c.GetIntOpt("outer_vlan")
		innerVlan := c.GetIntOpt("inner_vlan")
		if innerVlan == nil {
			return nil, fmt.Errorf("The `inner_vlan` property is required for qinq vlan configs")
		}
		cfg := &ixapi.VLANConfigQinQ{
			VLANType:           "qinq",
			OuterVLAN:          outerVlan,
			OuterVLANEthertype: outerVlanEthertype,
			InnerVLAN:          *innerVlan,
		}
		return cfg, nil
	}

	return nil, fmt.Errorf("unknown vlan config type: %s", vType)
}
