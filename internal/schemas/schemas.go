package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//
// CAUTION:
//   This file is generated from the IX-API
//   openapi specs. DO NOT EDIT.
//

// SchemaVersion is the version of the IX-API schema
const SchemaVersion = "2.7.1"

// CancellationPolicySchema is the terraform schema for the model
func CancellationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This field denotes the first possible cancellation date of the service.",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will end on this date. Typically `≥ decommission_at`.",
		},
	}
}

// AvailabilityZoneSchema is the terraform schema for the model
func AvailabilityZoneSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `AvailabilityZone`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name (description) for the availability zone ",
		},
	}
}

// CloudNetworkProductOfferingSchema is the terraform schema for the model
func CloudNetworkProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Cloud Network Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"provider_vlans": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `NetworkService` provides `single` or `multi`ple vlans.",
		},

		"service_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"service_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When configuring access to the network service, at least this `capacity` must be provided.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs.",
		},

		"nsc_required_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"service_provider_region": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service provider offers the network service for a specific region. ",
		},

		"service_provider_pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The datacenter id of the partner NNI to the service provider. It supposed to be used when identifying a location via the cloud provider's APIs. ",
		},

		"service_provider_pop_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The datacenter description of the partner NNI to the service provider. ",
		},

		"service_provider_workflow": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "When the workflow is `provider_first` the subscriber creates a circuit with the cloud provider and provides a `cloud_key` for filtering the product-offerings.  If the workflow is `exchange_first` the IX will create the cloud circuit on the provider side. ",
		},

		"service_exchange_pops": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: ServiceExchangePopSchema(),
			},
		},

		"delivery_method": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The exchange delivers the service over a `shared` or `dedicated` NNI.",
		},

		"diversity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The service can be delivered over multiple handovers from the exchange to the `service_provider`. The `diversity` denotes the number of handovers between the exchange and the service provider. A value of two signals a redundant service.  Only one network service configuration for each `handover` and `cloud_vlan` can be created.",
		},

		"nsc_supported_cloud_config_peering_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_required_cloud_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_cloud_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// ConnectionProductOfferingSchema is the terraform schema for the model
func ConnectionProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Connection Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"cross_connect_initiator": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A cross connect can be initiated by either the exchange or the subscriber.  This property affects which side has to provide a LOA and demarc information.",
		},

		"handover_pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ID of the point of presence (see `/pops`), where the physical port will be present. ",
		},

		"maximum_port_quantity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The maximum amount of ports which can be aggregated in the connection. `null` means no limit.",
		},

		"required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// DeviceSchema is the terraform schema for the model
func DeviceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Device`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the device ",
		},

		"pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `PointOfPresence` the device is in.",
		},

		"capabilities": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: DeviceCapabilitySchema(),
			},
		},

		"facility": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Identifier of the facility where the device is physically based.",
		},
	}
}

// DeviceCapabilitySchema is the terraform schema for the model
func DeviceCapabilitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"media_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The media type of the port (e.g. 1000BASE-LX, 10GBASE-LR, ...) ",
		},

		"speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Speed of port in Mbit/s ",
		},

		"max_lag": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Maximum count of ports which can be bundled to a max_lag",
		},

		"availability": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Count of available ports on device ",
		},
	}
}

// DeviceConnectionSchema is the terraform schema for the model
func DeviceConnectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"capacity_max": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},

		"device": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `Device`.  ",
		},

		"connected_device": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `Device`.  ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Device Connection`.",
		},
	}
}

// ExchangeLanNetworkProductOfferingSchema is the terraform schema for the model
func ExchangeLanNetworkProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Exchange Lan Network Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"provider_vlans": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `NetworkService` provides `single` or `multi`ple vlans.",
		},

		"service_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"service_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When configuring access to the network service, at least this `capacity` must be provided.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs.",
		},

		"nsc_required_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"exchange_lan_network_service": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the exchange lan network service.",
		},
	}
}

// FacilitySchema is the terraform schema for the model
func FacilitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Facility`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the Datacenter as called by the operator ",
		},

		"metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea` the DC is located in. ",
		},

		"address_country": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "ISO 3166-1 alpha-2 country code, for example DE ",
		},

		"address_locality": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The locality/city. For example, Mountain View.",
		},

		"address_region": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The region. For example, CA",
		},

		"postal_code": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A postal code. For example, 9404",
		},

		"street_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The street address. For example, 1600 Amphitheatre Pkwy.",
		},

		"peeringdb_facility_id": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "[PeeringDB](https://www.peeringdb.com) facitlity ID, can be extracted from the url https://www.peeringdb.com/fac/$id ",
		},

		"organisation_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of Datacenter operator ",
		},

		"pops": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"latitude": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Latitude of the facility's location. ",
		},

		"longitude": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Longitude of the facility's location. ",
		},
	}
}

// MP2MPNetworkProductOfferingSchema is the terraform schema for the model
func MP2MPNetworkProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `MP2MP Network Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"provider_vlans": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `NetworkService` provides `single` or `multi`ple vlans.",
		},

		"service_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"service_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When configuring access to the network service, at least this `capacity` must be provided.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs.",
		},

		"nsc_required_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// MetroAreaSchema is the terraform schema for the model
func MetroAreaSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `MetroArea`.",
		},

		"un_locode": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The UN/LOCODE for identifying the metro area. ",
		},

		"iata_code": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The three letter IATA airport code for identiying the metro area. ",
		},

		"display_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the metro area. Likely the same as the IATA code. ",
		},

		"facilities": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"metro_area_networks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// MetroAreaNetworkSchema is the terraform schema for the model
func MetroAreaNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `MetroAreaNetwork`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the metro area network. ",
		},

		"metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the metro area. ",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service provider is operating the network. Usually the exchange. ",
		},

		"pops": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// P2MPNetworkProductOfferingSchema is the terraform schema for the model
func P2MPNetworkProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `P2MP Network Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"provider_vlans": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `NetworkService` provides `single` or `multi`ple vlans.",
		},

		"service_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"service_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When configuring access to the network service, at least this `capacity` must be provided.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs.",
		},

		"nsc_required_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// P2PNetworkProductOfferingSchema is the terraform schema for the model
func P2PNetworkProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `P2P Network Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"provider_vlans": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `NetworkService` provides `single` or `multi`ple vlans.",
		},

		"service_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service is directly provided on the metro area network.  In case of a `p2p_vc`, the `service_metro_area_network` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"service_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The service is delivered in this metro area.  In case of a `p2p_vc`, the `service_metro_area` refers to the B-side of the point-to-point connection. The B-side is the accepting party. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When configuring access to the network service, at least this `capacity` must be provided.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "When not `null`, this value enforces a mandatory rate limit for all network service configs.",
		},

		"nsc_required_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_supported_l3_config_fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// PointOfPresenceSchema is the terraform schema for the model
func PointOfPresenceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"facility": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The pop is located in this `Facility`.",
		},

		"metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `MetroAreaNetwork`.  ",
		},

		"devices": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Point Of Presence`.",
		},

		"availability_zone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Availability zone of the pop.",
		},
	}
}

// ProductOfferingSchema is the terraform schema for the model
func ProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// RoutingFunctionProductOfferingSchema is the terraform schema for the model
func RoutingFunctionProductOfferingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Routing Function Product Offering`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the product",
		},

		"display_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"exchange_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the internet exchange. ",
		},

		"service_provider_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing the logo of the service provider. ",
		},

		"product_logo": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An URI referencing a logo for the product offered. ",
		},

		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"handover_metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork`. The service will be accessed through the handover metro area network.  In case of a `p2p_vc`, the `handover_metro_area_network` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"handover_metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroArea`. The network service will be accessed from this metro area.  In case of a `p2p_vc`, the `handover_metro_area` refers to the A-side of the point-to-point connection. The A-side is the entity which initiates the network service creation. ",
		},

		"physical_port_speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the service is dependent on the speed of the physical port this field denotes the speed.",
		},

		"service_provider": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the provider providing the service. ",
		},

		"downgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a lower bandwidth.",
		},

		"upgrade_allowed": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the service can be migrated to a higher bandwidth.",
		},

		"orderable_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering becomes available for ordering after this point in time.",
		},

		"orderable_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This product offering will become unavailable for ordering after this point in time.",
		},

		"contract_terms": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This property informally describe the contract's notice- and renewal periods as well as additional terms.  **Note**: This property contains informal information about the contract. For a structured representation see: `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  **Example**: A contract with the terms _\"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period\"_ can be represented as: * `contract_initial_period: \"P2W\"` * `contract_initial_notice_period: \"P5D\"` * `contract_renewal_period: \"P6M\"` * `contract_renewal_notice_period: \"P1M\"` ",
		},

		"contract_initial_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The initial duration of the contract. The contract will be renewed after this period for the duration of `contract_renewal_period`. ",
		},

		"contract_initial_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period for canceling the contract within the initial period. ",
		},

		"contract_renewal_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The duration for which the contract will be renewed after the initial period.  Unless the contract is canceled, it will be automatically renewed after the period. Cancellation has to be done within the `contract_renewal_notice_period`. ",
		},

		"contract_renewal_notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "_**Format:** ISO8601 Duration_  The notice period denotes the time before the end of the `contract_renewal_period` in which the client has to inform the IXP in order to prevent renewal of the contract. ",
		},

		"notice_period": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "**DEPRECATION NOTICE**: This property will be replaced by `contract_initial_period`, `contract_initial_notice_period`, `contract_renewal_period` and `contract_renewal_notice_period`.  The notice period informally states constraints which define when the client needs to inform the IXP in order to prevent renewal of the contract. ",
		},

		"bandwidth_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The minimum bandwidth of the routing service in Mbit/s.",
		},

		"bandwidth_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The maximum bandwidth of the routing service in Mbit/s.",
		},
	}
}

// ServiceExchangePopSchema is the terraform schema for the model
func ServiceExchangePopSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the `PointOfPresence` the service is provided.",
		},

		"path_info": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An *optional* text property that describes the path of the service where it is tethered through another party.",
		},
	}
}

// CloudConfigSchema is the terraform schema for the model
func CloudConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bgp_password": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The password to use for BGP sessions.",
		},

		"bgp_neighbor_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the BGP neighbor.",
		},

		"bgp_neighbor_address_primary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The primary IP address of the BGP neighbor.",
		},

		"bgp_neighbor_address_secondary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The secondary IP address of the BGP neighbor.",
		},

		"bgp_neighbor_asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The ASN of the BGP neighbor.",
		},

		"bgp_address_family": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"bfd": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable BFD for the BGP session.",
		},

		"local_asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The local ASN.",
		},

		"local_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the router function instance in CIDR notation.",
		},

		"local_address_primary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The primary IP address of the router function instance in CIDR notation.",
		},

		"local_address_secondary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The secondary IP address of the router function instance in CIDR notation.",
		},

		"vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "If the `provider_vlans` property of the `ProductOffering` is `multi`, a numeric value refers to a specific vlan on the service provider side.  The `nsc_required_cloud_config_fields` attribute of the `ProductOffering` will include `vlan` if `provider_vlans` are `multi`.",
		},

		"peering_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Some `cloud_vc` network services require selecting a peering type.  See the `nsc_supported_cloud_config_peering_types` attribute of the corresponding `ProductOffering` for valid values.",
		},
	}
}

// CloudNetworkServiceConfigSchema is the terraform schema for the model
func CloudNetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the configured `NetworkService`.",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Cloud Network Service Config`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the connection to use for this `NetworkServiceConfig`.  If no connection is specified, you have to provide a routing function.  When a connection is provided, you also need to specify the `lan_config`. The `routing_function` attribute may not be used. Some network services may require the use of the `l3_config`, please check the `nsc_required_l3_config_fields` attribute of the `ProductOffering`.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"routing_function": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.  If no routing function is provided, you need to provide the connection to use.  When a routing function is provided, you also need to specify the `l3_config`. The `connection` attribute may not be used.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"l3_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: L3ConfigSchema(),
			},
		},

		"network_feature_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"vlan_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`.",
		},

		"handover": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The handover enumerates the connection and is required for checking diversity constraints.  It must be within `1 <= x <= network_service.diversity`. ",
		},

		"cloud_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: CloudConfigSchema(),
			},
		},

		"cloud_vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "**Deprecation Notice**: This field is deprecated and will be removed in favor of using the `cloud_config.vlan` property. The `ProductOffering` will include `vlan` in the `nsc_required_cloud_config_fields`, if `provider_vlans` are `multi`.  If the `provider_vlans` property of the `ProductOffering` is `multi`, a numeric value refers to a specific vlan on the service provider side.  Otherwise, if set to `null`, it refers to all unmatched vlan ids on the service provider side. (All vlan ids from the service provider side are presented as tags within any vlans specified in `vlan_config`.)  If the `provider_vlans` property of the `ProductOffering` is `single`, the `cloud_vlan` MUST be `null` or MUST NOT be provided.",
		},

		"availability_zone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The availability zone that shall be used on the provider side. *(Sensitive Property)*",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},
	}
}

// ConnectionSchema is the terraform schema for the model
func ConnectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Connection`.",
		},

		"mode": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Sets the mode of the connection. The mode can be:  - `lag_lacp`: connection is build as a LAG with LACP enabled - `lag_static`: connection is build as LAG with static configuration - `flex_ethernet`: connect is build as a FlexEthernet channel - `standalone`: only one port is allowed in this connection without any bundling. ",
		},

		"lacp_timeout": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This sets the LACP Timeout mode. Both ends of the connections need to be configured the same. ",
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The product offering must match the type `connection`.",
		},

		"port_quantity": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The number of ports which should be allocated for this connection.",
		},

		"subscriber_side_demarcs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"connecting_party": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the service provider who establishes connectivity on your behalf.  This is only relevant, if the cross connect initiator is the `subscriber` and might be `null`.  Please refer to the usage guide of the internet exchange.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
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
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"port_reservations": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ID of the point of presence (see `/pops`), where the physical port(s) are present. ",
		},

		"speed": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Shows the total bandwidth of the connection in Mbit/s. ",
		},

		"capacity_allocated": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Sum of the bandwidth of all network service configs using the connection in Mbit/s.",
		},

		"capacity_allocation_limit": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Maximum allocatable capacity of the connection in Mbit/s. When `null`, the exchange does not impose any limit. ",
		},

		"vlan_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"outer_vlan_ethertypes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"metro_area": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Optional ID of the service metro area the connection is provided in.",
		},

		"metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Optional ID of the service metro area network the connection is present on.",
		},
	}
}

// ExchangeLanNetworkServiceConfigSchema is the terraform schema for the model
func ExchangeLanNetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the configured `NetworkService`.",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Exchange Lan Network Service Config`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the connection to use for this `NetworkServiceConfig`.  If no connection is specified, you have to provide a routing function.  When a connection is provided, you also need to specify the `lan_config`. The `routing_function` attribute may not be used. Some network services may require the use of the `l3_config`, please check the `nsc_required_l3_config_fields` attribute of the `ProductOffering`.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"routing_function": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.  If no routing function is provided, you need to provide the connection to use.  When a routing function is provided, you also need to specify the `l3_config`. The `connection` attribute may not be used.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"l3_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: L3ConfigSchema(),
			},
		},

		"network_feature_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"vlan_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity. *(Sensitive Property)*",
		},

		"asns": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			}},

		"macs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"ips": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"listed": &schema.Schema{
			Type:        schema.TypeBool,
			Required:    true,
			Description: "The customer wants to be featured on the member list",
		},

		"consumer_side_ready": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "You can use this optional property to signal to the IXP, that your equipment is set up and ready to be tested. *(Sensitive Property)*",
		},

		"availability_zone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The availability zone that shall be used on the provider side.  Availability Zones may not be supported for exchange_lan because by default they span multiple networks.  If an availability zone is set then this refers to a circuit that is placed on a specific on-ramp to the exchange_lan. *(Sensitive Property)*",
		},

		"shared_statistics": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: SharedStatisticsConfigSchema(),
			},
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The product offering must match the type `exchange_lan` and must refer to the related network service through the `exchange_lan_network_service` property.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},
	}
}

// IXPSpecificFeatureFlagConfigSchema is the terraform schema for the model
func IXPSpecificFeatureFlagConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the feature flag. ",
		},

		"enabled": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable the feature.  *Mandatory features can not be disabled*.",
		},
	}
}

// L3ConfigSchema is the terraform schema for the model
func L3ConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bgp_password": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The password to use for BGP sessions.",
		},

		"bgp_neighbor_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the BGP neighbor.",
		},

		"bgp_neighbor_address_primary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The primary IP address of the BGP neighbor.",
		},

		"bgp_neighbor_address_secondary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The secondary IP address of the BGP neighbor.",
		},

		"bgp_neighbor_asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The ASN of the BGP neighbor.",
		},

		"bgp_address_family": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},

		"bfd": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable BFD for the BGP session.",
		},

		"local_asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The local ASN.",
		},

		"local_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the router function instance in CIDR notation.",
		},

		"local_address_primary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The primary IP address of the router function instance in CIDR notation.",
		},

		"local_address_secondary": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The secondary IP address of the router function instance in CIDR notation.",
		},
	}
}

// MP2MPNetworkServiceConfigSchema is the terraform schema for the model
func MP2MPNetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the configured `NetworkService`.",
		},

		"macs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`.",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity. *(Sensitive Property)*",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `MP2MP Network Service Config`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the connection to use for this `NetworkServiceConfig`.  If no connection is specified, you have to provide a routing function.  When a connection is provided, you also need to specify the `lan_config`. The `routing_function` attribute may not be used. Some network services may require the use of the `l3_config`, please check the `nsc_required_l3_config_fields` attribute of the `ProductOffering`.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"routing_function": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.  If no routing function is provided, you need to provide the connection to use.  When a routing function is provided, you also need to specify the `l3_config`. The `connection` attribute may not be used.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"l3_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: L3ConfigSchema(),
			},
		},

		"network_feature_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"vlan_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},

		"ips": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"asns": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			}},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},
	}
}

// NetworkFeatureConfigSchema is the terraform schema for the model
func NetworkFeatureConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// NetworkServiceConfigSchema is the terraform schema for the model
func NetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// P2MPNetworkServiceConfigSchema is the terraform schema for the model
func P2MPNetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"macs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`.",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity. *(Sensitive Property)*",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `P2MP Network Service Config`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the connection to use for this `NetworkServiceConfig`.  If no connection is specified, you have to provide a routing function.  When a connection is provided, you also need to specify the `lan_config`. The `routing_function` attribute may not be used. Some network services may require the use of the `l3_config`, please check the `nsc_required_l3_config_fields` attribute of the `ProductOffering`.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"routing_function": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.  If no routing function is provided, you need to provide the connection to use.  When a routing function is provided, you also need to specify the `l3_config`. The `connection` attribute may not be used.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"l3_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: L3ConfigSchema(),
			},
		},

		"network_feature_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"vlan_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},

		"role": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A `leaf` can only reach roots and is isolated from other leafs. A `root` can reach any other point in the virtual circuit including other roots.",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the configured `NetworkService`.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},
	}
}

// P2PNetworkServiceConfigSchema is the terraform schema for the model
func P2PNetworkServiceConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the configured `NetworkService`.",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `P2P Network Service Config`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the connection to use for this `NetworkServiceConfig`.  If no connection is specified, you have to provide a routing function.  When a connection is provided, you also need to specify the `lan_config`. The `routing_function` attribute may not be used. Some network services may require the use of the `l3_config`, please check the `nsc_required_l3_config_fields` attribute of the `ProductOffering`.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"routing_function": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.  If no routing function is provided, you need to provide the connection to use.  When a routing function is provided, you also need to specify the `l3_config`. The `connection` attribute may not be used.  Connections ans Routing Functions are mutually exclusive. *(Sensitive Property)*",
		},

		"l3_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: L3ConfigSchema(),
			},
		},

		"network_feature_configs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"vlan_config": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},

		"macs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An optional id of a `ProductOffering`.  Valid ids of product-offerings can be found in the `nsc_product_offerings` property of the `NetworkService`.",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. If set to Null, the maximum capacity will be used, i.e. the virtual circuit is not rate-limited.  An exchange may choose to constrain the available capacity range of a `ProductOffering`.  That means, the service can consume up to the total bandwidth of the `Connection`.  Typically the service is charged based on the capacity. *(Sensitive Property)*",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},
	}
}

// PortSchema is the terraform schema for the model
func PortSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Port`.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `Connection`.  ",
		},

		"speed": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the port (set by the exchange)",
		},

		"media_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The media type of the port. Query the device's capabilities for available types. ",
		},

		"operational_state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The operational state of the port.",
		},

		"device": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The device the port. ",
		},

		"pop": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Same as the `pop` of the `device`. ",
		},
	}
}

// PortReservationSchema is the terraform schema for the model
func PortReservationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `A PortReservation`.",
		},

		"subscriber_side_demarc": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "In an exchange initiated scenario, this field will indicated one of the provided `subscriber_side_demarcs` from the connection.",
		},

		"connecting_party": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the service provider who establishes connectivity on your behalf.  This is only relevant, if the cross connect initiator is the `subscriber`.  Please refer to the usage guide of the internet exchange.",
		},

		"cross_connect_id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "An optional identifier of a cross connect.",
		},

		"network_connection": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `Port` will become part of this connection.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},

		"exchange_side_demarc": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Exchange side demarc information. This field will only be filled in when the port state is `allocated` or in `production`.  Otherwise this field will be `null`.",
		},

		"port": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "This field will be null, until a port will be allocated.",
		},
	}
}

// RouteServerNetworkFeatureConfigSchema is the terraform schema for the model
func RouteServerNetworkFeatureConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Route Server Network Feature Config`.",
		},

		"network_feature": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `NetworkFeature`.  ",
		},

		"network_service_config": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `NetworkServiceConfig`.  ",
		},

		"role_assignments": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The ASN of the peer. ",
		},

		"password": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The cleartext BGP session password",
		},

		"as_set_v4": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "AS-SET of the customer for IPv4 prefix filtering. This is used to generate filters on the router servers.  Only valid referenced prefixes within the AS-SET are allowed inbound to the route server. All other routes are filtered.  This field is *required* if the route server network feature only supports the `af_inet` address family. If multiple address families are supported, it is optional if the `as_set_v6` is provided.  Important: The format has to be: \"AS-SET@IRR\". IRR is the database where the AS-SET is registred. Typically used IRR's are RADB, RIPE, NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC ",
		},

		"as_set_v6": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "AS-SET of the customer for IPv6. This is used to generate filters on the router servers. Only valid referenced prefixes within the AS-SET are allowed inbound to the route server. All other routes are filtered.  This field is *required* if the route server network feature only supports the `af_inet6` address family. If multiple address families are supported, it is optional if the `as_set_v4` is provided.  Important: The format has to be: \"AS-SET@IRR\". IRR is the database where the AS-SET is registred. Typically used IRR's are RADB, RIPE, NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC ",
		},

		"max_prefix_v4": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Announcing more than `max_prefix` IPv4 prefixes the bgp session will be droped. ",
		},

		"max_prefix_v6": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Announcing more than `max_prefix` IPv6 prefixes the bgp session will be droped. ",
		},

		"insert_ixp_asn": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Insert the ASN of the exchange into the AS path. This function is only used in special cases. In 99% of all cases, it should be false. ",
		},

		"session_mode": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Set the session mode with the routeserver. ",
		},

		"bgp_session_type": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The session type describes which of the both parties will open the connection. If set to passive, the customer router needs to open the connection. If its set to active, the route server will open the connection. The standard behavior on most exchanges is passive. ",
		},

		"ip": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The BGP session will be established from this IP address, referenced by ID.  Only IDs of IPs assigned to the corresponding network service config can be used.",
		},

		"flags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: IXPSpecificFeatureFlagConfigSchema(),
			},
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},
	}
}

// SharedStatisticsConfigSchema is the terraform schema for the model
func SharedStatisticsConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"policy": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

// SharedStatisticsConfigAllowSchema is the terraform schema for the model
func SharedStatisticsConfigAllowSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"nsc_available_capacity": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"accounts_denied": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// SharedStatisticsConfigDenySchema is the terraform schema for the model
func SharedStatisticsConfigDenySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"nsc_available_capacity": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"accounts_allowed": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// VLANConfigDot1QSchema is the terraform schema for the model
func VLANConfigDot1QSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "A VLAN tag. If `null`, the IXP will auto-select a valid vlan-id. ",
		},

		"vlan_ethertype": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ethertype of the vlan in hexadecimal notation.",
		},
	}
}

// VLANConfigQinQSchema is the terraform schema for the model
func VLANConfigQinQSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"outer_vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The outer VLAN id. If `null`, the IXP will auto-select a valid vlan-id. ",
		},

		"outer_vlan_ethertype": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The ethertype of the outer tag in hexadecimal notation.",
		},

		"inner_vlan": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The inner VLAN id. ",
		},
	}
}

// AccountSchema is the terraform schema for the model
func AccountSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Account`.",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of a managing account. Can be used for creating a customer hierachy. *(Sensitive Property)* ",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the account, how it gets represented in e.g. a \"customers list\". ",
		},

		"legal_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Legal name of the organisation. Only required when it's different from the account name. *(Sensitive Property)* ",
		},

		"billing_information": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: BillingInformationSchema(),
			},
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)*",
		},

		"discoverable": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "The account will be included for all members of the ix in the list of accounts.  Only `id`, `name` and `present_in_metro_area_networks` are provided to other members.",
		},

		"metro_area_network_presence": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"address": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Required: true,
			Elem: &schema.Resource{
				Schema: AddressSchema(),
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
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"asns": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			}},
	}
}

// AddressSchema is the terraform schema for the model
func AddressSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"country": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "ISO 3166-1 alpha-2 country code, for example DE",
		},

		"locality": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The locality/city. For example, Mountain View.",
		},

		"region": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The region. For example, CA",
		},

		"postal_code": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A postal code. For example, 9404",
		},

		"street_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The street address. For example, 1600 Amphitheatre Pkwy.",
		},

		"post_office_box_number": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The post office box number for PO box addresses.",
		},
	}
}

// BillingInformationSchema is the terraform schema for the model
func BillingInformationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the organisation receiving invoices. ",
		},

		"address": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: AddressSchema(),
			},
		},

		"vat_number": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Value-added tax number, required for european reverse charge system. ",
		},
	}
}

// ContactSchema is the terraform schema for the model
func ContactSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Contact`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A name of a person or an organisation",
		},

		"telephone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The telephone number in E.164 Phone Number Formatting",
		},

		"email": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The email of the legal company entity. ",
		},
	}
}

// RoleSchema is the terraform schema for the model
func RoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the role. ",
		},

		"required_fields": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Role for a Contact`.",
		},
	}
}

// RoleAssignmentSchema is the terraform schema for the model
func RoleAssignmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"role": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of a role the contact is assigned to. ",
		},

		"contact": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of a contact the role is assigned to. ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `A role assignment for a contact`.",
		},
	}
}

// EventSchema is the terraform schema for the model
func EventSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"serial": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},

		"account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `Account`.  ",
		},

		"timestamp": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

// StatusSchema is the terraform schema for the model
func StatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"severity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "We are using syslog severity levels: 0 = Emergency, 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning, 5 = Notice, 6 = Informational, 7 = Debug. ",
		},

		"tag": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A machine readable message identifier. ",
		},

		"message": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A human readable message, describing the problem and may contain hints for resolution. ",
		},

		"timestamp": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The time and date when the event occured.",
		},
	}
}

// APIExtensionSchema is the terraform schema for the model
func APIExtensionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the extension. ",
		},

		"publisher": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Publisher of the extension. ",
		},

		"documentation_url": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "URL of the documentation homepage of the extension. ",
		},

		"base_url": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Extension endpoints are available under this base url. ",
		},

		"spec_url": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "URL of the extensions schema specifications. The schema format schould be OpenAPI v3. ",
		},
	}
}

// APIFeaturesSchema is the terraform schema for the model
func APIFeaturesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pagination": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "The API implementation supports pagination on `list` operations.",
		},
	}
}

// APIHealthSchema is the terraform schema for the model
func APIHealthSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"status": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "status indicates whether the service status is acceptable or not.",
		},

		"version": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Public version of the service. ",
		},

		"releaseId": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Release version of the api implementation. ",
		},

		"notes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"output": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Raw error output, in case of \"fail\" or \"warn\" states.",
		},

		"serviceId": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A unique identifier of the service, in the application scope.",
		},

		"description": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A human-friendly description of the service.",
		},
	}
}

// APIImplementationSchema is the terraform schema for the model
func APIImplementationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"schema_version": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Version of the implemented IX-API schema. ",
		},

		"service_version": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Version of the API service. ",
		},

		"supported_network_service_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"supported_network_service_config_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"supported_network_feature_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"supported_network_feature_config_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"supported_operations": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"supported_features": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: APIFeaturesSchema(),
			},
		},
	}
}

// IPAddressSchema is the terraform schema for the model
func IPAddressSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `IP-Address`.",
		},

		"version": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The version of the internet protocol. ",
		},

		"address": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "IPv4 or IPv6 Address in the following format: - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) - IPv6: hexadecimal colon separated notation ",
		},

		"prefix_length": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The CIDR ip prefix length ",
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
	}
}

// IPAddressShortSchema is the terraform schema for the model
func IPAddressShortSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The version of the internet protocol. ",
		},

		"address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address in the following format: - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) - IPv6: hexadecimal colon separated notation ",
		},
	}
}

// MacAddressSchema is the terraform schema for the model
func MacAddressSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `MAC-Address`.",
		},

		"address": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Unicast MAC address, formatted hexadecimal values with colons. ",
		},

		"valid_not_before": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "When a mac address is assigned to a NSC, and the current datetime is before this value, then the MAC address *cannot* be used on the peering platform.  Afterwards, it is supposed to be available. If the value is `null` or the property does not exist, the mac address is valid from the creation date.",
		},

		"valid_not_after": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "When a mac address is assigned to an NSC, and the current datetime is before this value, the MAC address *can* be used on the peering platform.  Afterwards, it is supposed to be unassigned from the NSC and cannot any longer be used on the peering platform.  If the value is null or the property does not exist, the MAC address is valid indefinitely. The value may not be in the past.",
		},
	}
}

// PeerSchema is the terraform schema for the model
func PeerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The ASN of the peer. ",
		},

		"ip": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: IPAddressShortSchema(),
			},
		},

		"mac_address": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Unicast MAC address, formatted hexadecimal values with colons. ",
		},
	}
}

// ConflictSchema is the terraform schema for the model
func ConflictSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The resource type refers to an ix-api resource. ",
		},

		"resource_id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the resource which has a conflict with the request operation on the current resource. ",
		},

		"resource_property": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Indicates the property where the resource is in use. ",
		},

		"remote_resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The type of the conflicting resource. ",
		},

		"remote_resource_id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The id of the conflicting resource. This is in most cases the id of the current resource. ",
		},
	}
}

// ProblemResponseSchema is the terraform schema for the model
func ProblemResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A short, human-readable summary of the problem type.  It SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization (e.g., using proactive content negotiation; see [RFC7231], Section 3.4). ",
		},

		"status": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The HTTP status code ([RFC7231], Section 6) generated by the origin server for this occurrence of the problem.",
		},

		"detail": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A human-readable explanation specific to this occurrence of the problem.",
		},

		"instance": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A URI reference that identifies the specific occurrence of the problem.  It may or may not yield further information if dereferenced.",
		},
	}
}

// ValidationErrorPropertySchema is the terraform schema for the model
func ValidationErrorPropertySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
}

// AllowMemberJoiningRuleSchema is the terraform schema for the model
func AllowMemberJoiningRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account to which access to the network service should be granted or denied. ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `A rule for members joining a private vlan`.",
		},

		"capacity_min": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Require an optional minimum capacity to join the network service.",
		},

		"capacity_max": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "An optional rate limit which has precedence over the capacity set in the network service config.",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `NetworkService`.  ",
		},
	}
}

// CloudNetworkServiceSchema is the terraform schema for the model
func CloudNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `ProductOffering`.  ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Cloud Network Service`.",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. When null, the maximum capacity will be used.",
		},

		"cloud_key": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The cloud key is used to specify to which user or existing circuit of a cloud provider this `network-service` should be provisioned.  For example, for a provider like *AWS*, this would be the *account number* (Example: `123456789876`), or for a provider like Azure, this would be the service key (Example: `acl9edcf-f11c-4681-9c7b-6d16b2973997`)",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"nsc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},

		"diversity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Same value as the corresponding `ProductOffering`.  The service can be delivered over multiple handovers from the exchange to the `service_provider`.  The `diversity` denotes the number of handovers between the exchange and the service provider. A value of two signals a redundant service.  Only one network service configuration for each `handover` and `cloud_vlan` can be created.",
		},

		"provider_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "For a cloud network service with the exchange first workflow, the `provider_ref` will be a reference to a resource of the cloud provider. (E.g. the UUID of a virtual circuit.)  The `provider_ref` is managed by the exchange and its meaning may vary between different cloud services. ",
		},

		"availability_zones": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"network_features": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_product_offerings": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// DenyMemberJoiningRuleSchema is the terraform schema for the model
func DenyMemberJoiningRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account to which access to the network service should be granted or denied. ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `A rule for members joining a private vlan`.",
		},

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `NetworkService`.  ",
		},
	}
}

// ExchangeLanNetworkServiceSchema is the terraform schema for the model
func ExchangeLanNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"nsc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Exchange Lan Network Service`.",
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Exchange-dependent service name, will be shown on the invoice.",
		},

		"metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Id of the `MetroAreaNetwork` where the exchange lan network service is directly provided.  Same as `service_metro_area_network` on the related `ProductOffering`. ",
		},

		"peeringdb_ixid": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "PeeringDB ixid",
		},

		"ixfdb_ixid": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "id of ixfdb",
		},

		"network_features": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"subnet_v4": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) CIDR notation. ",
		},

		"subnet_v6": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv6 subnet in hexadecimal colon separated CIDR notation. ",
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "*deprecation notice*",
		},
	}
}

// IXPSpecificFeatureFlagSchema is the terraform schema for the model
func IXPSpecificFeatureFlagSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the feature flag. ",
		},

		"description": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The description of the feature flag. ",
		},

		"mandatory": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "This feature will always be enabled, even if not provided in the corresponding config's list of `flags`. ",
		},
	}
}

// MP2MPNetworkServiceSchema is the terraform schema for the model
func MP2MPNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `ProductOffering`.  ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `MP2MP Network Service`.",
		},

		"public": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "A public mp2mp network service can be joined by everyone on the exchange unless denied by a member-joining-rule.  Public network services are visible to other members of the IXP, however only `display_name`, `type`, `product_offering`, `consuming_account` and `managing_account` are made available.  Other required fields are redacted.",
		},

		"display_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the multi-point to multi-point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of, for example: \"Financial Clearance LAN\". ",
		},

		"subnet_v4": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) CIDR notation. ",
		},

		"subnet_v6": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv6 subnet in hexadecimal colon separated CIDR notation. ",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"nsc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_product_offerings": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},

		"member_joining_rules": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"network_features": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"mac_acl_protection": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "When enabled, only MAC addresses in the referenced in the network service config's `macs` property are allowed to send and receive traffic on this network service.",
		},
	}
}

// MemberJoiningRuleSchema is the terraform schema for the model
func MemberJoiningRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// NetworkFeatureSchema is the terraform schema for the model
func NetworkFeatureSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// NetworkServiceSchema is the terraform schema for the model
func NetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// NetworkServiceDeleteResponseSchema is the terraform schema for the model
func NetworkServiceDeleteResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// P2MPNetworkServiceSchema is the terraform schema for the model
func P2MPNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `ProductOffering`.  ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `P2MP Network Service`.",
		},

		"display_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the point to multi-point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of. ",
		},

		"public": &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "A public p2mp network service can be joined by everyone on the exchange unless denied by a member-joining-rule.  Public network services are visible to other members of the IXP, however only `name`, `type`, `product_offering`, `consuming_account` and `managing_account` are made available.  Other required fields are redacted.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"nsc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_product_offerings": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},

		"network_features": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"member_joining_rules": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},
	}
}

// P2PNetworkServiceSchema is the terraform schema for the model
func P2PNetworkServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the related `ProductOffering`.  ",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `P2P Network Service`.",
		},

		"display_name": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Name of the point to point virtual circuit.  It is visible to all parties allowed to connect to this virtual circuit.  It is intended for humans to make sense of. ",
		},

		"joining_member_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The account of the B-side member joining the virtual circuit. ",
		},

		"availability_zones": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},

		"nsc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"nsc_product_offerings": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service will be decommissioned on this date.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`. *(Sensitive Property)*",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The service continues incurring charges until this date. Typically `≥ decommission_at`.  This field is only used when the state is `DECOMMISSION_REQUESTED` or `DECOMMISSIONED`.  *(Sensitive Property)*",
		},

		"current_billing_start_date": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Your obligation to pay for the service will start on this date.  However, this date may change after an upgrade and not reflect the inital start date of the service.  *(Sensitive Property)*",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity of the service in Mbps. When null, the maximum capacity will be used.",
		},
	}
}

// RouteServerNetworkFeatureSchema is the terraform schema for the model
func RouteServerNetworkFeatureSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The *primary identifier* of the `Route Server Network Feature`.",
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

		"network_service": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The `id` of the related `NetworkService`.  ",
		},

		"nfc_required_contact_roles": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"flags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: IXPSpecificFeatureFlagSchema(),
			},
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
			Description: "The FQDN of the route server. ",
		},

		"looking_glass_url": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The url of the looking glass. ",
		},

		"address_families": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"session_mode": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "When creating a route server feature config, remember to specify the same session_mode as the route server. ",
		},

		"available_bgp_session_types": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"ip_v4": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv4 address in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation) notation.  This field is only set if the `address_families` include `af_inet`. ",
		},

		"ip_v6": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "IPv6 address in hexadecimal colon separated notation.  This field is only set if the `address_families` include `af_inet6`. ",
		},
	}
}

// RoutingFunctionSchema is the terraform schema for the model
func RoutingFunctionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account responsible for managing the service via the API. A manager can read and update the state of entities. *(Sensitive Property)* ",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The `id` of the account consuming a service.  Used to be `owning_customer`. *(Sensitive Property)* ",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Reference field, free to use for the API user. *(Sensitive Property)* ",
		},

		"purchase_order": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Purchase Order ID which will be displayed on the invoice. *(Sensitive Property)* ",
		},

		"contract_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "A reference to a contract. If no specific contract is used, a default MAY be chosen by the implementer. *(Sensitive Property)* ",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "An account requires billing_information to be used as a `billing_account`. *(Sensitive Property)*",
		},

		"id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The *primary identifier* of the `Routing Function`.",
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The product offering to be used for the routing function.",
		},

		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Any routing function instance needs to be assigned a 2-byte or 4-byte ASN of the customer's choice. There is no restriction on private or public ASNs.",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The desired upper bound of the capacity for the routing function.",
		},

		"state": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The state of the object. *(Sensitive Property)*",
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: StatusSchema(),
			},
		},
	}
}

// AggregateSchema is the terraform schema for the model
func AggregateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// AggregateStatisticsSchema is the terraform schema for the model
func AggregateStatisticsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Title of the aggregated statistics. ",
		},

		"start": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Start of the traffic aggregation.",
		},

		"end": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "End of the traffic aggregation.",
		},

		"accuracy": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The accuracy is the ratio of *total aggregated samples* to *expected samples*.  The expected number of samples is the size of the window of the aggregate, divided by the aggregation resolution.  For example: A window of `24 h` with an aggregation resolution of `5 m` would yield `288` samples.  If only `275` samples are available for aggregation, the accuracy would be `0.95`. ",
		},

		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the statistics were created.",
		},

		"next_update_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Next update of the statistical data. This may not correspond to the aggregate interval.",
		},

		"average_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average number of inbound **packets per second**. ",
		},

		"average_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average number outbound **packets per second**. ",
		},

		"average_ops_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average inbound **octets per second**. ",
		},

		"average_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average outbound **octets per second**. ",
		},

		"average_eps_in": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Average **errors per second** inbound. ",
		},

		"average_eps_out": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Averages **errors per second** outbound. ",
		},

		"average_dps_in": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Average **discards per second** inbound. ",
		},

		"average_dps_out": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Averages **discards per second** outbound. ",
		},

		"percentile95_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of the inbound **octets per second**. ",
		},

		"percentile95_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of outbound **packets per second**. ",
		},

		"percentile95_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of outbound **octets per second**. ",
		},

		"maximum_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak inbound **packets per second** during the interval. ",
		},

		"maximum_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak outbound **packets per second** during the interval. ",
		},

		"maximum_ops_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak inbound **octets per second** during the interval. ",
		},

		"maximum_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak outbound **octets per second** during the interval. ",
		},

		"maximum_in_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the inbound peak occured.",
		},

		"maximum_out_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the outbound peak occured.",
		},
	}
}

// AggregateTimeseriesSchema is the terraform schema for the model
func AggregateTimeseriesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Title of the timeseries. ",
		},

		"precision": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Precision indicates the sampling rate of the aggregated traffic data in seconds. For example if the data is aggregated over 5 minutes, the precision would be 300. ",
		},

		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the statistics were created.",
		},

		"next_update_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Next update of the statistical data. This may not correspond to the aggregate interval.",
		},

		"origin_timezone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The timezone where the data was collected in tz database format. ",
		},

		"fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"samples": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeList,
			}},
	}
}

// NetworkServiceConfigAggregateSchema is the terraform schema for the model
func NetworkServiceConfigAggregateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// NetworkServiceConfigAggregateStatisticsSchema is the terraform schema for the model
func NetworkServiceConfigAggregateStatisticsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Title of the aggregated statistics. ",
		},

		"start": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Start of the traffic aggregation.",
		},

		"end": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "End of the traffic aggregation.",
		},

		"accuracy": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The accuracy is the ratio of *total aggregated samples* to *expected samples*.  The expected number of samples is the size of the window of the aggregate, divided by the aggregation resolution.  For example: A window of `24 h` with an aggregation resolution of `5 m` would yield `288` samples.  If only `275` samples are available for aggregation, the accuracy would be `0.95`. ",
		},

		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the statistics were created.",
		},

		"next_update_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Next update of the statistical data. This may not correspond to the aggregate interval.",
		},

		"average_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average number of inbound **packets per second**. ",
		},

		"average_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average number outbound **packets per second**. ",
		},

		"average_ops_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average inbound **octets per second**. ",
		},

		"average_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Average outbound **octets per second**. ",
		},

		"average_eps_in": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Average **errors per second** inbound. ",
		},

		"average_eps_out": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Averages **errors per second** outbound. ",
		},

		"average_dps_in": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Average **discards per second** inbound. ",
		},

		"average_dps_out": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Averages **discards per second** outbound. ",
		},

		"percentile95_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of the inbound **octets per second**. ",
		},

		"percentile95_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of outbound **packets per second**. ",
		},

		"percentile95_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "95th percentile of outbound **octets per second**. ",
		},

		"maximum_pps_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak inbound **packets per second** during the interval. ",
		},

		"maximum_pps_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak outbound **packets per second** during the interval. ",
		},

		"maximum_ops_in": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak inbound **octets per second** during the interval. ",
		},

		"maximum_ops_out": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Peak outbound **octets per second** during the interval. ",
		},

		"maximum_in_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the inbound peak occured.",
		},

		"maximum_out_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the outbound peak occured.",
		},

		"nsc_available_capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The capacity left on the `NetworkServiceConfig` in **megabits per second** (Mbps). ",
		},

		"nsc_available_capacity_change_perc": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The percentage change of the available capacity since the last update. ",
		},
	}
}

// PeerAggregateSchema is the terraform schema for the model
func PeerAggregateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"peer": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: PeerSchema(),
			},
		},
	}
}

// PeerRTTSchema is the terraform schema for the model
func PeerRTTSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"time_ms": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The total duration of the measurement in milliseconds. ",
		},

		"tx": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The number of probe packets *transmitted* within the duration of the measurement. ",
		},

		"rx": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The number of probe packets *received* within the duration of the measurement. ",
		},

		"loss": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "Ratio of *transmitted packets* to *received packets*: `loss = 1.0 - (rx / tx)`. ",
		},

		"rtt_min_ms": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The minimum RTT in milliseconds. ",
		},

		"rtt_avg_ms": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The average RTT in milliseconds. ",
		},

		"rtt_max_ms": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The maximum RTT in milliseconds. ",
		},

		"rtt_mdev_ms": &schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Computed:    true,
			Description: "The median RTT in milliseconds. Standard deviation in milliseconds. ",
		},

		"neighbor": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The name of the peer. ",
		},

		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The Autonomous System Number (ASN) of the peer. ",
		},

		"ip": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The IP address of the peer. For IPv6 addresses the canonical form is used. ",
		},

		"timestamp": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The date and time when the RTT statistic was measured.",
		},

		"serial": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The `serial` is an incrementing counter. You can use it to poll for changes. ",
		},
	}
}

// PeerTimeseriesSchema is the terraform schema for the model
func PeerTimeseriesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"title": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Title of the timeseries. ",
		},

		"precision": &schema.Schema{
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "Precision indicates the sampling rate of the aggregated traffic data in seconds. For example if the data is aggregated over 5 minutes, the precision would be 300. ",
		},

		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Timestamp when the statistics were created.",
		},

		"next_update_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Next update of the statistical data. This may not correspond to the aggregate interval.",
		},

		"origin_timezone": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The timezone where the data was collected in tz database format. ",
		},

		"fields": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			}},

		"samples": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeList,
			}},

		"peer": &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: PeerSchema(),
			},
		},
	}
}

// PortStatisticsSchema is the terraform schema for the model
func PortStatisticsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"light_levels_tx": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			}},

		"light_levels_rx": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeFloat,
			}},
	}
}
