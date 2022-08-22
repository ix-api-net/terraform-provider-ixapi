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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This field denotes the first possible cancellation date of the service."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will end on this date. Typically `≥ decommission_at`."},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"provider_vlans": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `NetworkService` provides `single` or `multi`ple vlans."},

	"service_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"service_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"bandwidth_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When configuring access to the network service, at least this `capacity` must be provided."},

	"bandwidth_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs."},

	"service_provider_region": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service provider offers the network service for a specific region. "},

	"service_provider_pop": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The datacenter description of the partner NNI to the service provider. "},

	"service_provider_workflow": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "When the workflow is `provider_first` the subscriber creates a circuit with the cloud provider and provides a `cloud_key` for filtering the product-offerings.  If the workflow is `exchange_first` the IX will create the cloud circuit on the provider side. "},

	"delivery_method": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The exchange delivers the service over a `shared` or `dedicated` NNI."},

	"diversity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The service can be delivered over multiple handovers from the exchange to the `service_provider`. The `diversity` denotes the number of handovers between the exchange and the service provider. A value of two signals a redundant service.  Only one network service configuration for each `handover` and `cloud_vlan` can be created."},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"cross_connect_initiator": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A cross connect can be initiated by either the exchange or the subscriber.  This property affects which side has to provide a LOA and demarc information."},

	"handover_pop": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The ID of the point of presence (see `/pops`), where the physical port will be present. "},

	"maximum_port_quantity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The maximum amount of ports which can be aggregated in the connection. `null` means no limit."},

	"required_contact_roles": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},
}

// DeviceSchema is the terraform schema for the model
var DeviceSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the device "},

	"pop": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `PointOfPresence` the device is in."},

	"capabilities": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"facility": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Identifier of the facility where the device is physically based."},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// DeviceCapabilitySchema is the terraform schema for the model
var DeviceCapabilitySchema = map[string]*schema.Schema{
	"media_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The media type of the port (e.g. 1000BASE-LX, 10GBASE-LR, ...) "},

	"speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Speed of port in Mbit/s "},

	"max_lag": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Maximum count of ports which can be bundled to a max_lag"},

	"availability": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Count of available ports on device "},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"provider_vlans": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `NetworkService` provides `single` or `multi`ple vlans."},

	"service_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"service_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"bandwidth_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When configuring access to the network service, at least this `capacity` must be provided."},

	"bandwidth_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs."},

	"exchange_lan_network_service": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The id of the exchange lan network service."},
}

// FacilitySchema is the terraform schema for the model
var FacilitySchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the Datacenter as called by the operator "},

	"metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea` the DC is located in. "},

	"address_country": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "ISO 3166-1 alpha-2 country code, for example DE "},

	"address_locality": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The locality/city. For example, Mountain View."},

	"address_region": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The region. For example, CA"},

	"postal_code": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A postal code. For example, 9404"},

	"street_address": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The street address. For example, 1600 Amphitheatre Pkwy."},

	"peeringdb_facility_id": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "[PeeringDB](https://www.peeringdb.com) facitlity ID, can be extracted from the url https://www.peeringdb.com/fac/$id "},

	"organisation_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of Datacenter operator "},

	"pops": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"latitude": &schema.Schema{
		Type:        schema.TypeFloat,
		Optional:    true,
		Computed:    true,
		Description: "Latitude of the facility's location. "},

	"longitude": &schema.Schema{
		Type:        schema.TypeFloat,
		Optional:    true,
		Computed:    true,
		Description: "Longitude of the facility's location. "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"provider_vlans": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `NetworkService` provides `single` or `multi`ple vlans."},

	"service_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"service_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"bandwidth_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When configuring access to the network service, at least this `capacity` must be provided."},

	"bandwidth_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs."},
}

// MetroAreaSchema is the terraform schema for the model
var MetroAreaSchema = map[string]*schema.Schema{
	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"un_locode": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The UN/LOCODE for identifying the metro area. "},

	"iata_code": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The three letter IATA airport code for identiying the metro area. "},

	"display_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the metro area. Likely the same as the IATA code. "},

	/*
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
	*/
}

// MetroAreaNetworkSchema is the terraform schema for the model
var MetroAreaNetworkSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the metro area network. "},

	"metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The id of the metro area. "},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service provider is operating the network. Usually the exchange. "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"provider_vlans": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `NetworkService` provides `single` or `multi`ple vlans."},

	"service_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"service_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"bandwidth_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When configuring access to the network service, at least this `capacity` must be provided."},

	"bandwidth_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs."},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the product"},

	"display_name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the internet exchange. "},

	"service_provider_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing the logo of the service provider. "},

	"product_logo": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An URI referencing a logo for the product offered. "},

	"resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"handover_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"handover_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. "},

	"physical_port_speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "If the service is dependent on the speed of the physical port this field denotes the speed."},

	"service_provider": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the provider providing the service. "},

	"downgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a lower bandwidth."},

	"upgrade_allowed": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Indicates if the service can be migrated to a higher bandwidth."},

	"orderable_not_before": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering becomes available for ordering after this point in time."},

	"orderable_not_after": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This product offering will become unavailable for ordering after this point in time."},

	"contract_terms": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The contract terms informally describe the contract period and renewals. "},

	"notice_period": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. "},

	"provider_vlans": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `NetworkService` provides `single` or `multi`ple vlans."},

	"service_metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"service_metro_area": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. "},

	"bandwidth_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When configuring access to the network service, at least this `capacity` must be provided."},

	"bandwidth_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs."},
}

// PointOfPresenceSchema is the terraform schema for the model
var PointOfPresenceSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"facility": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The pop is located in this `Facility`."},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"network_service": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the configured network service."},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the connection to use for this `NetworkServiceConfig`."},

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
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The handover enumerates the connection and is required for checking diversity constraints.  It must be within `1 <= x <= network_service.diversity`. "},

	"cloud_vlan": &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: "If the `provider_vlans` property of the `ProductOffering` is `multi`, a numeric value refers to a specific vlan on the service provider side.  Otherwise, if set to `null`, it refers to all unmatched vlan ids on the service provider side. (All vlan ids from the service provider side are presented as tags within any vlans specified in `vlan_config`.)  If the `provider_vlans` property of the `ProductOffering` is `single`, the `cloud_vlan` MUST be `null` or MUST NOT be provided."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// ConnectionSchema is the terraform schema for the model
var ConnectionSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"mode": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Sets the mode of the connection. The mode can be:  - `lag_lacp`: connection is build as a LAG with LACP enabled - `lag_static`: connection is build as LAG with static configuration - `flex_ethernet`: connect is build as a FlexEthernet channel - `standalone`: only one port is allowed in this connection without any bundling. "},

	"lacp_timeout": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This sets the LACP Timeout mode. Both ends of the connections need to be configured the same. "},

	"product_offering": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The product offering must match the type `connection`."},

	"port_quantity": &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The number of ports which should be allocated for this connection."},

	"subscriber_side_demarcs": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"connecting_party": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the service provider who establishes connectivity on your behalf.  This is only relevant, if the cross connect initiator is the `subscriber` and might be `null`.  Please refer to the usage guide of the internet exchange."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The ID of the point of presence (see `/pops`), where the physical port(s) are present. "},

	"speed": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Shows the total bandwidth of the connection in Mbit/s. "},

	"capacity_allocated": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Sum of the bandwidth of all network services using the connection in Mbit/s."},

	"capacity_allocation_limit": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Maximum allocatable capacity of the connection in Mbit/s. When `null`, the exchange does not impose any limit. "},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"network_service": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the configured network service."},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the connection to use for this `NetworkServiceConfig`."},

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
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity."},

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
		Type:        schema.TypeBool,
		Required:    true,
		Description: "The customer wants to be featured on the member list"},

	"product_offering": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The product offering must match the type `exchange_lan` and must refer to the related network service through the `exchange_lan_network_service` property."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// IXPSpecificFeatureFlagConfigSchema is the terraform schema for the model
var IXPSpecificFeatureFlagConfigSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the feature flag. "},

	"enabled": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Enable the feature.  *Mandatory features can not be disabled*."},
}

// MP2MPNetworkServiceConfigSchema is the terraform schema for the model
var MP2MPNetworkServiceConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"network_service": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the configured network service."},

	"product_offering": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`."},

	"capacity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity."},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the connection to use for this `NetworkServiceConfig`."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"network_service": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the configured network service."},

	"product_offering": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`."},

	"capacity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity."},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the connection to use for this `NetworkServiceConfig`."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A `leaf` can only reach roots and is isolated from other leafs. A `root` can reach any other point in the virtual circuit including other roots."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"network_service": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the configured network service."},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"role_assignments": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
	},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The id of the connection to use for this `NetworkServiceConfig`."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`."},

	"capacity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the port (set by the exchange)"},

	"media_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The media type of the port. Query the device's capabilities for available types. "},

	"operational_state": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The operational state of the port."},

	"device": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The device the port. "},

	"pop": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Same as the `pop` of the `device`. "},
}

// PortReservationSchema is the terraform schema for the model
var PortReservationSchema = map[string]*schema.Schema{
	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"subscriber_side_demarc": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "In an exchange initiated scenario, this field will indicated one of the provided `subscriber_side_demarcs` from the connection."},

	"connecting_party": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the service provider who establishes connectivity on your behalf.  This is only relevant, if the cross connect initiator is the `subscriber`.  Please refer to the usage guide of the internet exchange."},

	"cross_connect_id": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "An optional identifier of a cross connect."},

	"connection": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `Port` will become part of this connection."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},

	"exchange_side_demarc": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Exchange side demarc information. This field will only be filled in when the port state is `allocated` or in `production`.  Otherwise this field will be `null`."},

	"port": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "This field will be null, until a port will be allocated."},
}

// RouteServerNetworkFeatureConfigSchema is the terraform schema for the model
var RouteServerNetworkFeatureConfigSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

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
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The ASN of the peer. "},

	"password": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The cleartext BGP session password"},

	"as_set_v4": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "AS-SET of the customer for IPv4 prefix filtering. This is used to generate filters on the router servers.  Only valid referenced prefixes within the AS-SET are allowed inbound to the route server. All other routes are filtered.  This field is *required* if the route server network feature only supports the `af_inet` address family. If multiple address families are supported, it is optional if the `as_set_v6` is provided.  Important: The format has to be: \"AS-SET@IRR\". IRR is the database where the AS-SET is registred. Typically used IRR's are RADB, RIPE, NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC "},

	"as_set_v6": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "AS-SET of the customer for IPv6. This is used to generate filters on the router servers. Only valid referenced prefixes within the AS-SET are allowed inbound to the route server. All other routes are filtered.  This field is *required* if the route server network feature only supports the `af_inet6` address family. If multiple address families are supported, it is optional if the `as_set_v4` is provided.  Important: The format has to be: \"AS-SET@IRR\". IRR is the database where the AS-SET is registred. Typically used IRR's are RADB, RIPE, NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC "},

	"max_prefix_v4": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Announcing more than `max_prefix` IPv4 prefixes the bgp session will be droped. "},

	"max_prefix_v6": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Announcing more than `max_prefix` IPv6 prefixes the bgp session will be droped. "},

	"insert_ixp_asn": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "Insert the ASN of the exchange into the AS path. This function is only used in special cases. In 99% of all cases, it should be false. "},

	"session_mode": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Set the session mode with the routeserver. "},

	"bgp_session_type": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The session type describes which of the both parties will open the connection. If set to passive, the customer router needs to open the connection. If its set to active, the route server will open the connection. The standard behavior on most exchanges is passive. "},

	"ip": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The BGP session will be established from this IP address, referenced by ID.  Only IDs of IPs assigned to the corresponding network service config can be used."},

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
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "A VLAN tag. If `null`, the IXP will auto-select a valid vlan-id. "},

	"vlan_ethertype": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The ethertype of the vlan in hexadecimal notation."},
}

// VLANConfigPortSchema is the terraform schema for the model
var VLANConfigPortSchema = map[string]*schema.Schema{}

// VLANConfigQinQSchema is the terraform schema for the model
var VLANConfigQinQSchema = map[string]*schema.Schema{
	"outer_vlan": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The outer VLAN id. If `null`, the IXP will auto-select a valid vlan-id. "},

	"outer_vlan_ethertype": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The ethertype of the outer tag in hexadecimal notation."},

	"inner_vlan": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The inner VLAN id. "},
}

// AccountSchema is the terraform schema for the model
var AccountSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `id` of a managing account. Can be used for creating a customer hierachy. *(Sensitive Property)* "},

	"name": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the account, how it gets represented in e.g. a \"customers list\". "},

	"legal_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Legal name of the organisation. Only required when it's different from the account name. *(Sensitive Property)* "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)*"},

	"discoverable": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "The account will be included for all members of the ix in the list of accounts.  Only `id`, `name` and `present_in_metro_area_networks` are provided to other members."},

	"metro_area_network_presence": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "ISO 3166-1 alpha-2 country code, for example DE"},

	"locality": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The locality/city. For example, Mountain View."},

	"region": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The region. For example, CA"},

	"postal_code": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A postal code. For example, 9404"},

	"street_address": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The street address. For example, 1600 Amphitheatre Pkwy."},

	"post_office_box_number": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The post office box number for PO box addresses."},
}

// BillingInformationSchema is the terraform schema for the model
var BillingInformationSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the organisation receiving invoices. "},

	"address": &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: AddressSchema,
		},
	},

	"vat_number": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Value-added tax number, required for european reverse charge system. "},
}

// ContactSchema is the terraform schema for the model
var ContactSchema = map[string]*schema.Schema{
	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A name of a person or an organisation"},

	"telephone": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The telephone number in E.164 Phone Number Formatting"},

	"email": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The email of the legal company entity. "},

	"id": &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
}

// RoleSchema is the terraform schema for the model
var RoleSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the role. "},

	"required_fields": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of a role the contact is assigned to. "},

	"contact": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of a contact the role is assigned to. "},

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
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "We are using syslog severity levels: 0 = Emergency, 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning, 5 = Notice, 6 = Informational, 7 = Debug. "},

	"tag": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A machine readable message identifier. "},

	"message": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A human readable message, describing the problem and may contain hints for resolution. "},

	"attrs": nil,
	"timestamp": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The time and date when the event occured."},
}

// APIExtensionsSchema is the terraform schema for the model
var APIExtensionsSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the extension. "},

	"publisher": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Publisher of the extension. "},

	"documentation_url": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "URL of the documentation homepage of the extension. "},

	"base_url": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Extension endpoints are available under this base url. "},

	"spec_url": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "URL of the extensions schema specifications. The schema format schould be OpenAPI v3. "},
}

// APIHealthSchema is the terraform schema for the model
var APIHealthSchema = map[string]*schema.Schema{
	"status": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "status indicates whether the service status is acceptable or not."},

	"version": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Public version of the service. "},

	"releaseId": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Release version of the api implementation. "},

	"notes": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"output": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Raw error output, in case of \"fail\" or \"warn\" states."},

	"serviceId": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A unique identifier of the service, in the application scope."},

	"description": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A human-friendly description of the service."},

	"checks": nil,
	"links":  nil,
}

// APIImplementationSchema is the terraform schema for the model
var APIImplementationSchema = map[string]*schema.Schema{
	"schema_version": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Version of the implemented IX-API schema. "},

	"service_version": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Version of the API service. "},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"version": &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The version of the internet protocol. "},

	"address": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "IPv4 or IPv6 Address in the following format: - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) - IPv6: hexadecimal colon separated notation "},

	"prefix_length": &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The CIDR ip prefix length "},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"address": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "Unicast MAC address, formatted hexadecimal values with colons. "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The resource type refers to an ix-api resource. "},

	"resource_id": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The id of the resource which has a conflict with the request operation on the current resource. "},

	"resource_property": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Indicates the property where the resource is in use. "},

	"remote_resource_type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The type of the conflicting resource. "},

	"remote_resource_id": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The id of the conflicting resource. This is in most cases the id of the current resource. "},
}

// ProblemResponseSchema is the terraform schema for the model
var ProblemResponseSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A URI reference (see RFC3986) that identifies the problem type.  This specification encourages that, when dereferenced, it provide human-readable documentation for the problem type (e.g., using HTML [W3C.REC-html5-20141028]).  When this member is not present, its value is assumed to be \"about:blank\". "},

	"title": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A short, human-readable summary of the problem type.  It SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization (e.g., using proactive content negotiation; see [RFC7231], Section 3.4). "},

	"status": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The HTTP status code ([RFC7231], Section 6) generated by the origin server for this occurrence of the problem."},

	"detail": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A human-readable explanation specific to this occurrence of the problem."},

	"instance": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A URI reference that identifies the specific occurrence of the problem.  It may or may not yield further information if dereferenced."},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account to which access to the network service should be granted or denied. "},

	"capacity_min": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Require an optional minimum capacity to join the network service."},

	"capacity_max": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "An optional rate limit which has precedence over the capacity set in the network service config."},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"cloud_key": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"capacity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. When null, the maximum capacity will be used."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

	"diversity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "Same value as the corresponding `ProductOffering`.  The service can be delivered over multiple handovers from the exchange to the `service_provider`.  The `diversity` denotes the number of handovers between the exchange and the service provider. A value of two signals a redundant service.  Only one network service configuration for each `handover` and `cloud_vlan` can be created."},

	"provider_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "For a cloud network service with the exchange first workflow, the `provider_ref` will be a reference to a resource of the cloud provider. (E.g. the UUID of a virtual circuit.)  The `provider_ref` is managed by the exchange and its meaning may vary between different cloud services. "},
}

// DenyMemberJoiningRuleSchema is the terraform schema for the model
var DenyMemberJoiningRuleSchema = map[string]*schema.Schema{
	"type": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"managing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account to which access to the network service should be granted or denied. "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Exchange-dependent service name, will be shown on the invoice."},

	"metro_area_network": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Id of the `MetroAreaNetwork` where the exchange lan network service is directly provided.  Same as `service_metro_area_network` on the related `ProductOffering`. "},

	"peeringdb_ixid": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "PeeringDB ixid"},

	"ixfdb_ixid": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "id of ixfdb"},

	"network_features": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"subnet_v4": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) CIDR notation. "},

	"subnet_v6": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "IPv6 subnet in hexadecimal colon separated CIDR notation. "},

	"product_offering": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "*deprecation notice*"},
}

// IXPSpecificFeatureFlagSchema is the terraform schema for the model
var IXPSpecificFeatureFlagSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The name of the feature flag. "},

	"description": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The description of the feature flag. "},

	"mandatory": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "This feature will always be enabled, even if not provided in the corresponding config's list of `flags`. "},
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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"public": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "A public mp2mp network service can be joined by everyone on the exchange unless denied by a member-joining-rule.  Public network services are visible to other members of the IXP, however only `display_name`, `type`, `product_offering`, `consuming_account` and `managing_account` are made available.  Other required fields are redacted."},

	"display_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the multi-point to multi-point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of, for example: \"Financial Clearance LAN\". "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"display_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the point to multi-point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of. "},

	"public": &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Computed:    true,
		Description: "A public p2mp network service can be joined by everyone on the exchange unless denied by a member-joining-rule.  Public network services are visible to other members of the IXP, however only `name`, `type`, `product_offering`, `consuming_account` and `managing_account` are made available.  Other required fields are redacted."},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

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
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. "},

	"consuming_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. "},

	"external_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Reference field, free to use for the API user. *(Sensitive Property)* "},

	"purchase_order": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* "},

	"contract_ref": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* "},

	"billing_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*"},

	"display_name": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Name of the point to point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of. "},

	"joining_member_account": &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The account of the B-side member joining the virtual circuit. "},

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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`."},

	"charged_until": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*"},

	"current_billing_start_date": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*"},

	"capacity": &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Computed:    true,
		Description: "The capacity of the service in Mbps. When null, the maximum capacity will be used."},
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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The FQDN of the route server. "},

	"looking_glass_url": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "The url of the looking glass. "},

	"address_families": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"session_mode": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "When creating a route server feature config, remember to specify the same session_mode as the route server. "},

	"available_bgp_session_types": &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
	},

	"ip_v4": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "IPv4 address in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) notation.  This field is only set if the `address_families` include `af_inet`. "},

	"ip_v6": &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: "IPv6 address in hexadecimal colon separated notation.  This field is only set if the `address_families` include `af_inet6`. "},
}
