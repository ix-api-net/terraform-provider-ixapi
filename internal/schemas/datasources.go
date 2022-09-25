package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func boolFromOption(b []bool) bool {
	if len(b) > 0 {
		return b[0]
	}
	return false
}

// DataSourceQuery creates a string data source schema
// used for querying.
// We assume string as the type, as this is the case for
// almost every query property.
func DataSourceQuery(description string) *schema.Schema {
	s := &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: description,
	}
	return s
}

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

// IntoDataSourceResultsSchema makes a new result schema for a
// to be used in a data source returning a list of results.
func IntoDataSourceResultsSchema(
	s map[string]*schema.Schema,
) *schema.Schema {
	IntoDataSourceSchema(s)
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: s,
		},
	}
}

// DataSourceID creates a schema mapping for combination
// with a required ID parameter
func DataSourceID() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}
