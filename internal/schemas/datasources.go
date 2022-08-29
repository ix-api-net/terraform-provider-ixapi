package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// IntoDataSourceSchema will mark all properties as optional
// except if it is marked as required - when it is required for
// filtering or lookup.
func IntoDataSourceSchema(
	s map[string]*schema.Schema,
	required ...string,
) map[string]*schema.Schema {
	for name, prop := range s {
		isRequired := false
		for _, requiredProp := range required {
			if name == requiredProp {
				isRequired = true
				break
			}
		}

		if isRequired {
			prop.Optional = false
			prop.Computed = true
			prop.Required = true
		} else {
			prop.Required = false
			prop.Optional = true
			prop.Computed = true
		}

		res, ok := prop.Elem.(*schema.Resource)
		if ok {
			IntoDataSourceSchema(res.Schema, required...)
		}
	}
	return s
}
