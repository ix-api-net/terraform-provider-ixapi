package schemas

import (
	"testing"
)

func TestCloudRouterSchema(t *testing.T) {
	s := CloudRouterSchema()
	required := []string{"managing_account", "billing_account", "consuming_account", "product_offering", "asn", "capacity"}
	for _, key := range required {
		f, ok := s[key]
		if !ok {
			t.Errorf("missing field: %s", key)
			continue
		}
		if f.Required != true {
			t.Errorf("field %s should be Required", key)
		}
	}
}
