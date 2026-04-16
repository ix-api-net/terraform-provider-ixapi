package ixapi

import (
	"encoding/json"
	"fmt"
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
