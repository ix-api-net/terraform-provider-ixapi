package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VLANConfigSchema is the polymorphic vlan config schema
func VLANConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vlan_type": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The vlan_type determines if a additional configuration is required: For `port` no configuration is required, `qinq` and `dot1q` require the config specific fields to be set.",
		},

		"outer_vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "ID of ther outer VLAN. If not present, the IXP will select a valid ID. Only required for vlan type QinQ.",
		},
		"outer_vlan_ethertype": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Outer vlan ether type, defaults to: 0x8100. Only used with type QinQ.",
		},
		"inner_vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "ID of the inner VLan. Only required with QinQ vlan type.",
		},

		"vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "A VLAN tag. If not present, the IXP will auto-select a valid vlan-id. Only used with VLAN type Dot1Q.",
		},
		"vlan_ethertype": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "VLAN ether type, defaults to: 0x8100. Only used with type Dot1Q.",
		},
	}
}
