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
			Description: "The vlan_type determines if a additional configuration is required: For `port` no configuration is required, `qinq` and `dot1q` require the `vlan_config_qinq` and `vlan_config_dotiq` to be set",
		},
		"vlan_config_qinq": &schema.Schema{
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "QinQ vlan configuration",
			Elem: &schema.Resource{
				Schema: VLANConfigQinQSchema(),
			},
		},
		"vlan_config_dot1q": &schema.Schema{
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Dot1Q vlan configuration",
			Elem: &schema.Resource{
				Schema: VLANConfigDot1QSchema(),
			},
		},
	}
}
