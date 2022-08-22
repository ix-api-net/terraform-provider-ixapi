package ixapi

import (
	"encoding/json"
	"testing"
)

const ValidationErrorJSON = `{"type": "https://errors.ix-api.net/v2/validation-error.html", "title": "Some fields did not validate.", "properties": [{"name": "legal_name", "reason": "This field may not be blank."}, {"name": "billing_information", "reason": {"name": ["This field is required."], "address": {"country": ["This field is required."], "locality": ["This field is required."], "postal_code": ["This field is required."], "street_address": ["This field is required."]}}}, {"name": "external_ref", "reason": "This field may not be blank."}, {"name": "address", "reason": {"country": ["This field is required."], "locality": ["This field is required."], "postal_code": ["This field is required."], "street_address": ["This field is required."]}}]}`

func TestValidationErrorDecoding(t *testing.T) {
	body := []byte(ValidationErrorJSON)
	res := &ValidationError{}
	if err := json.Unmarshal(body, res); err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
