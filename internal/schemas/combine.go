package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Combine merges two or more schemas. The last override wins.
// Can be used for making data source query props required.
func Combine(all ...map[string]*schema.Schema) map[string]*schema.Schema {
	combined := map[string]*schema.Schema{}
	for _, s := range all {
		for k, v := range s {
			combined[k] = v
		}
	}
	return combined
}
