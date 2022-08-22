package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VLANConfigSchema is the polymorphic vlan config schema
var VLANConfigSchema = map[string]*schema.Schema{
	"vlan_type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},
	"vlan_config_qinq": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: VLANConfigQinQSchema,
		},
	},
	"vlan_config_dot1q": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: VLANConfigDot1QSchema,
		},
	},
}
