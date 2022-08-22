package schemas

import "gitlab.com/ix-api/ix-api-terraform-provider/internal/ixapi"

// FlattenMetroArea makes a flat metroarea
func FlattenMetroArea(m *ixapi.MetroArea) map[string]interface{} {
	res := map[string]interface{}{}

	res["id"] = m.ID
	res["un_locode"] = m.UnLocode
	res["iata_code"] = m.IataCode
	res["display_name"] = m.DisplayName

	return res
}

// FlattenMetroAreas flattens the metro areas list
func FlattenMetroAreas(m []*ixapi.MetroArea) []interface{} {
	res := make([]interface{}, len(m))
	for i, met := range m {
		res[i] = FlattenMetroArea(met)
	}
	return res
}
