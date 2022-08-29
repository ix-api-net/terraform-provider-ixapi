package schemas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestIntoDataSourceSchema(t *testing.T) {
	s := map[string]*schema.Schema{
		"prop_a": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},

		"prop_b": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"prop_c": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"prop_b": &schema.Schema{
						Type:     schema.TypeInt,
						Required: false,
						Optional: true,
					},
				},
			},
		},
	}

	IntoDataSourceSchema(s, "prop_b")

	if s["prop_a"].Required == true {
		t.Error("prop_a should not be required")
	}
	if s["prop_b"].Required != true {
		t.Error("prop_b should be required")
	}
	if s["prop_c"].Elem.(*schema.Resource).Schema["prop_b"].Required == false {
		t.Error("prop_c.prop_b should be required")
	}

}
