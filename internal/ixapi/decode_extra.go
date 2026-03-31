package ixapi

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// decode the polymorphic embedded vlan config from response
// entity's VLANConfigRaw attribute
func decodeVLANConfig(data []byte) (VLANConfig, error) {
	if data == nil {
		return nil, nil
	}

	pCfg := &PolymorphicVLANConfig{}
	if err := json.Unmarshal(data, &pCfg); err != nil {
		return nil, err
	}

	switch pCfg.PolymorphicType() {
	case VLANConfigDot1QType:
		cfg := &VLANConfigDot1Q{}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
		return cfg, nil

	case VLANConfigQinQType:
		cfg := &VLANConfigQinQ{}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
		return cfg, nil

	case VLANConfigPortType:
		cfg := &VLANConfigPort{}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
		return cfg, nil
	}

	// Unknown vlan config type
	return nil, fmt.Errorf(
		"unknown vlan config type: %s - do not know how to decode",
		pCfg.PolymorphicType())
}

// dateOnlyPattern matches JSON date strings in YYYY-MM-DD format.
// The DE-CIX extension API returns date-only strings for some timestamp fields
// (e.g. decommission_at), which time.Parse rejects unless normalized to RFC3339.
var dateOnlyPattern = regexp.MustCompile(`"(\d{4}-\d{2}-\d{2})"`)

func preprocessDateFields(data []byte) []byte {
	return dateOnlyPattern.ReplaceAllFunc(data, func(match []byte) []byte {
		dateStr := string(match[1 : len(match)-1])
		if len(dateStr) == 10 {
			return []byte(`"` + dateStr + `T00:00:00Z"`)
		}
		return match
	})
}

// UnmarshalProductOffering unmarshals a DE-CIX Cloud Router product offering response,
// normalizing date-only fields to RFC3339 before decoding.
func UnmarshalProductOffering(data []byte, v interface{}) error {
	return json.Unmarshal(preprocessDateFields(data), v)
}
