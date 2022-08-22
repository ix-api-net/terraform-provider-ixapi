package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//
// CAUTION:
//   This file is generated from the IX-API
//   openapi specs. DO NOT EDIT.
//

// CancellationPolicySchema is the terraform schema for the model
var CancellationPolicySchema = map[string]*schema.Schema{
	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// CloudNetworkProductOfferingSchema is the terraform schema for the model
var CloudNetworkProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"provider_vlans": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"bandwidth_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"bandwidth_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider_region": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_pop": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_workflow": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"delivery_method": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"diversity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// ConnectionProductOfferingSchema is the terraform schema for the model
var ConnectionProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"cross_connect_initiator": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_pop": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"maximum_port_quantity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// DeviceSchema is the terraform schema for the model
var DeviceSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"pop": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"capabilities": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"facility": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// DeviceCapabilitySchema is the terraform schema for the model
var DeviceCapabilitySchema = map[string]*schema.Schema{
	"media_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"max_lag": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"availability": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// DeviceConnectionSchema is the terraform schema for the model
var DeviceConnectionSchema = map[string]*schema.Schema{
	"capacity_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"device": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"connected_device": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ExchangeLanNetworkProductOfferingSchema is the terraform schema for the model
var ExchangeLanNetworkProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"provider_vlans": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"bandwidth_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"bandwidth_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"exchange_lan_network_service": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// FacilitySchema is the terraform schema for the model
var FacilitySchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address_country": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address_locality": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address_region": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"postal_code": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"street_address": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"peeringdb_facility_id": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"organisation_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"pops": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"latitude": &schema.Schema{
		Type:     schema.TypeFloat,
		Optional: true,
		Computed: true,
	},

	"longitude": &schema.Schema{
		Type:     schema.TypeFloat,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// MP2MPNetworkProductOfferingSchema is the terraform schema for the model
var MP2MPNetworkProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"provider_vlans": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"bandwidth_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"bandwidth_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// MetroAreaSchema is the terraform schema for the model
var MetroAreaSchema = map[string]*schema.Schema{
	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"un_locode": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"iata_code": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"facilities": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"metro_area_networks": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// MetroAreaNetworkSchema is the terraform schema for the model
var MetroAreaNetworkSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"pops": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// P2MPNetworkProductOfferingSchema is the terraform schema for the model
var P2MPNetworkProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"provider_vlans": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"bandwidth_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"bandwidth_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// P2PNetworkProductOfferingSchema is the terraform schema for the model
var P2PNetworkProductOfferingSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_provider_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_logo": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"handover_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"physical_port_speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"service_provider": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"downgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"upgrade_allowed": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"orderable_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"orderable_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_terms": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notice_period": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"provider_vlans": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_metro_area": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"bandwidth_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"bandwidth_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// PointOfPresenceSchema is the terraform schema for the model
var PointOfPresenceSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"facility": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"devices": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// CloudNetworkServiceConfigSchema is the terraform schema for the model
var CloudNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_feature_configs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"vlan_config": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: VLANConfigSchema,
		},
	},

	"handover": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"cloud_vlan": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ConnectionSchema is the terraform schema for the model
var ConnectionSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"mode": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"lacp_timeout": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"port_quantity": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"subscriber_side_demarcs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"connecting_party": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"ports": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"port_reservations": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"pop": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"capacity_allocated": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"capacity_allocation_limit": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"vlan_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"outer_vlan_ethertypes": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// ExchangeLanNetworkServiceConfigSchema is the terraform schema for the model
var ExchangeLanNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_feature_configs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"vlan_config": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: VLANConfigSchema,
		},
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"asns": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"macs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"ips": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"listed": &schema.Schema{
		Type:     schema.TypeBool,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// IXPSpecificFeatureFlagConfigSchema is the terraform schema for the model
var IXPSpecificFeatureFlagConfigSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"enabled": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},
}

// MP2MPNetworkServiceConfigSchema is the terraform schema for the model
var MP2MPNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_feature_configs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"vlan_config": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: VLANConfigSchema,
		},
	},

	"macs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// P2MPNetworkServiceConfigSchema is the terraform schema for the model
var P2MPNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_feature_configs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"vlan_config": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: VLANConfigSchema,
		},
	},

	"role": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// P2PNetworkServiceConfigSchema is the terraform schema for the model
var P2PNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_feature_configs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"vlan_config": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: VLANConfigSchema,
		},
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// PortSchema is the terraform schema for the model
var PortSchema = map[string]*schema.Schema{
	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"speed": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"media_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"operational_state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"device": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"pop": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// PortReservationSchema is the terraform schema for the model
var PortReservationSchema = map[string]*schema.Schema{
	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"subscriber_side_demarc": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"connecting_party": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"cross_connect_id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"connection": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_side_demarc": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"port": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// RouteServerNetworkFeatureConfigSchema is the terraform schema for the model
var RouteServerNetworkFeatureConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"network_feature": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_service_config": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"flags": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"asn": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"password": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"as_set_v4": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"as_set_v6": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"max_prefix_v4": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"max_prefix_v6": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"insert_ixp_asn": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"session_mode": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"bgp_session_type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"ip": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// VLANConfigDot1QSchema is the terraform schema for the model
var VLANConfigDot1QSchema = map[string]*schema.Schema{
	"vlan": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"vlan_ethertype": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// VLANConfigPortSchema is the terraform schema for the model
var VLANConfigPortSchema = map[string]*schema.Schema{}

// VLANConfigQinQSchema is the terraform schema for the model
var VLANConfigQinQSchema = map[string]*schema.Schema{
	"outer_vlan": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"outer_vlan_ethertype": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"inner_vlan": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// AccountSchema is the terraform schema for the model
var AccountSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"legal_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_information": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: BillingInformationSchema,
		},
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"discoverable": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"metro_area_network_presence": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"address": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: AddressSchema,
		},
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// AddressSchema is the terraform schema for the model
var AddressSchema = map[string]*schema.Schema{
	"country": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"locality": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"region": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"postal_code": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"street_address": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"post_office_box_number": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// BillingInformationSchema is the terraform schema for the model
var BillingInformationSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: AddressSchema,
		},
	},

	"vat_number": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ContactSchema is the terraform schema for the model
var ContactSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"telephone": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"email": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// RoleSchema is the terraform schema for the model
var RoleSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"required_fields": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// RoleAssignmentSchema is the terraform schema for the model
var RoleAssignmentSchema = map[string]*schema.Schema{
	"role": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"contact": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// EventSchema is the terraform schema for the model
var EventSchema = map[string]*schema.Schema{
	"serial": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"payload": nil,
	"timestamp": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// StatusSchema is the terraform schema for the model
var StatusSchema = map[string]*schema.Schema{
	"severity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"tag": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"message": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"attrs": nil,
	"timestamp": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// APIExtensionsSchema is the terraform schema for the model
var APIExtensionsSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"publisher": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"documentation_url": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"base_url": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"spec_url": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// APIHealthSchema is the terraform schema for the model
var APIHealthSchema = map[string]*schema.Schema{
	"status": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"version": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"releaseId": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"notes": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"output": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"serviceId": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"description": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"checks": nil,
	"links":  nil,
}

// APIImplementationSchema is the terraform schema for the model
var APIImplementationSchema = map[string]*schema.Schema{
	"schema_version": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"service_version": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"supported_network_service_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"supported_network_service_config_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"supported_network_feature_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"supported_network_feature_config_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"supported_operations": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// IPAddressSchema is the terraform schema for the model
var IPAddressSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"version": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"address": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"prefix_length": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},

	"fqdn": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"valid_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"valid_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// MacAddressSchema is the terraform schema for the model
var MacAddressSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"valid_not_before": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"valid_not_after": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ConflictSchema is the terraform schema for the model
var ConflictSchema = map[string]*schema.Schema{
	"resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"resource_property": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"remote_resource_type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"remote_resource_id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ProblemResponseSchema is the terraform schema for the model
var ProblemResponseSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"title": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"detail": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"instance": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ValidationErrorPropertySchema is the terraform schema for the model
var ValidationErrorPropertySchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"reason": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// AllowMemberJoiningRuleSchema is the terraform schema for the model
var AllowMemberJoiningRuleSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"capacity_min": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"capacity_max": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// CloudNetworkServiceSchema is the terraform schema for the model
var CloudNetworkServiceSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"cloud_key": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"nsc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"diversity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"provider_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// DenyMemberJoiningRuleSchema is the terraform schema for the model
var DenyMemberJoiningRuleSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ExchangeLanNetworkServiceSchema is the terraform schema for the model
var ExchangeLanNetworkServiceSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"nsc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"metro_area_network": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"peeringdb_ixid": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"ixfdb_ixid": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"network_features": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"subnet_v4": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"subnet_v6": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// IXPSpecificFeatureFlagSchema is the terraform schema for the model
var IXPSpecificFeatureFlagSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"description": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"mandatory": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},
}

// MP2MPNetworkServiceSchema is the terraform schema for the model
var MP2MPNetworkServiceSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"public": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"nsc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"nsc_product_offerings": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"member_joining_rules": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"network_features": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// P2MPNetworkServiceSchema is the terraform schema for the model
var P2MPNetworkServiceSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"public": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"nsc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"nsc_product_offerings": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"network_features": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"member_joining_rules": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// P2PNetworkServiceSchema is the terraform schema for the model
var P2PNetworkServiceSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"product_offering": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"consuming_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"purchase_order": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"contract_ref": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"billing_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"joining_member_account": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"state": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"status": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"nsc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"nsc_product_offerings": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"decommission_at": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"charged_until": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"current_billing_start_date": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"capacity": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},
}

// RouteServerNetworkFeatureSchema is the terraform schema for the model
var RouteServerNetworkFeatureSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"required": &schema.Schema{
		Type:     schema.TypeBool,
		Optional: true,
		Computed: true,
	},

	"nfc_required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"flags": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"network_service": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"asn": &schema.Schema{
		Type:     schema.TypeInt,
		Optional: true,
		Computed: true,
	},

	"fqdn": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"looking_glass_url": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"address_families": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"session_mode": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"available_bgp_session_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"ip_v4": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"ip_v6": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}
