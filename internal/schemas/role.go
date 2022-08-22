package schemas

import "gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"

// FlattenRole makes a flat role
func FlattenRole(role *ixapi.Role) map[string]interface{} {
	res := map[string]interface{}{}
	res["name"] = role.Name
	res["id"] = role.ID
	res["required_fields"] = role.RequiredFields
	return res
}
