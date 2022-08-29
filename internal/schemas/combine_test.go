package schemas

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCombine(t *testing.T) {
	s := Combine(
		map[string]*schema.Schema{
			"prop_a": &schema.Schema{
				Optional: true,
			},
			"prop_b": &schema.Schema{
				Optional: true,
			},
		},
		map[string]*schema.Schema{
			"prop_b": &schema.Schema{
				Required: true,
			},
		},
	)

	if s["prop_a"].Optional != true {
		t.Error("expected prop_a to be optional")
	}
	if s["prop_b"].Optional == true {
		t.Error("expected prop_b to be not optional")
	}
	if s["prop_b"].Required != true {
		t.Error("expected prop_b to be required")
	}
}
