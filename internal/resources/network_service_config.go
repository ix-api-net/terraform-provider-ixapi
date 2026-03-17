package resources

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/ixapi"
	"github.com/ix-api-net/terraform-provider-ixapi/internal/schemas"
)

// addConnectionConstraints adds ConflictsWith and AtLeastOneOf validation
// to the network_connection and routing_function fields in a NSC schema.
// The IX-API requires exactly one of these fields to be provided.
func addConnectionConstraints(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["network_connection"].ConflictsWith = []string{"routing_function"}
	s["network_connection"].AtLeastOneOf = []string{"network_connection", "routing_function"}
	s["routing_function"].ConflictsWith = []string{"network_connection"}
	s["routing_function"].AtLeastOneOf = []string{"network_connection", "routing_function"}
	return s
}

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
