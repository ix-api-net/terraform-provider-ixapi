package ixapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//
// CAUTION:
//   This file is generated from the IX-API
//   openapi specs. DO NOT EDIT.
//

const dateLayout = "2006-01-02"

// Date is a date only time type representing
// a date without time.
type Date time.Time

// String implements the stringer interface for Date
func (d Date) String() string {
    t := time.Time(d)
    return t.Format(dateLayout)
}

// ParseDate decodes a date from a string
func ParseDate(value string) (Date, error) {
    t, err := time.Parse(dateLayout, value)
    if err != nil {
        return Date(t), err
    }
    return Date(t), nil
}

// MustParseDate decodes a Date using ParseDate, but
// will panic in case of an error.
func MustParseDate(value string) Date {
    d, err := ParseDate(value)
    if err != nil {
        panic(err)
    }
    return d
}

// UnmarshalJSON parses the json value of a date
func (d *Date) UnmarshalJSON(b []byte) error {
    s := strings.Trim(string(b), `"`)
    t, err := time.Parse(dateLayout, s)
    if err != nil {
        return err
    }
    *d = Date(t)
    return nil
}

// MarshalJSON returns the time in date format
func (d Date) MarshalJSON() ([]byte, error) {
    val := `"` + d.String() + `"`
    return []byte(val), nil
}

// Polymorphic indicates that the type is polymorphic.
type Polymorphic interface {
    PolymorphicType() string
}

// Response is an IX-API general response
type Response interface{}

// SchemaVersion is the version of the ix-api schema
const SchemaVersion = "2.7.1"

// AuthToken AuthToken
type AuthToken struct {
// AccessToken is a access_token
AccessToken string `json:"access_token,omitempty"`

// RefreshToken is a refresh_token
RefreshToken string `json:"refresh_token,omitempty"`

}

// AuthTokenRequest AuthTokenRequest
type AuthTokenRequest struct {
// APIKey is a api_key
APIKey string `json:"api_key,omitempty"`

// APISecret is a api_secret
APISecret string `json:"api_secret,omitempty"`

}

// RefreshTokenRequest RefreshTokenRequest
type RefreshTokenRequest struct {
// RefreshToken is a refresh_token
RefreshToken string `json:"refresh_token,omitempty"`

}

// CancellationPolicy Cancellation Policy
type CancellationPolicy struct {
// DecommissionAt This field denotes the first possible cancellation
// date of the service.
DecommissionAt Date `json:"decommission_at,omitempty"`

// ChargedUntil Your obligation to pay for the service will end on this date.
// Typically `≥ decommission_at`.
ChargedUntil Date `json:"charged_until,omitempty"`

}

// CancellationRequest Service Cancellation Request
type CancellationRequest struct {
// DecommissionAt An optional date for scheduling the cancellation
// and service decommissioning.
DecommissionAt *Date `json:"decommission_at,omitempty"`

}

// AvailabilityZone AvailabilityZone
type AvailabilityZone struct {
// ID The *primary identifier* of the `AvailabilityZone`.
ID string `json:"id,omitempty"`

// Name The name (description) for the availability zone
// 
Name string `json:"name,omitempty"`

}

// CloudNetworkProductOffering Cloud Network Product Offering
type CloudNetworkProductOffering struct {
// ID The *primary identifier* of the `Cloud Network Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

// ServiceProviderRegion The service provider offers the network service for a
// specific region.
// 
ServiceProviderRegion string `json:"service_provider_region,omitempty"`

// ServiceProviderPop The datacenter id of the partner NNI to the service provider.
// It supposed to be used when identifying a location via
// the cloud provider's APIs.
// 
ServiceProviderPop string `json:"service_provider_pop,omitempty"`

// ServiceProviderPopName The datacenter description of the partner NNI to the service provider.
// 
ServiceProviderPopName *string `json:"service_provider_pop_name,omitempty"`

// ServiceProviderWorkflow When the workflow is `provider_first` the subscriber creates
// a circuit with the cloud provider and provides a `cloud_key` for filtering
// the product-offerings.
// 
// If the workflow is `exchange_first` the IX will create
// the cloud circuit on the provider side.
// 
ServiceProviderWorkflow string `json:"service_provider_workflow,omitempty"`

// ServiceExchangePops A list of object, referencing a `PointOfPresence`
// and providing additional path information, in case the services
// is tethered through another party.
ServiceExchangePops []*ServiceExchangePop `json:"service_exchange_pops,omitempty"`

// DeliveryMethod The exchange delivers the service over a `shared` or `dedicated` NNI.
DeliveryMethod string `json:"delivery_method,omitempty"`

// Diversity The service can be delivered over multiple handovers from
// the exchange to the `service_provider`.
// The `diversity` denotes the number of handovers between the
// exchange and the service provider. A value of two signals a
// redundant service.
// 
// Only one network service configuration for each `handover` and
// `cloud_vlan` can be created.
Diversity int `json:"diversity,omitempty"`

// NscSupportedCloudConfigPeeringTypes The supported peering types for the cloud network service.
// 
// In case selecting a peering type is required, the
// `peering_type` property will be in the
// `nsc_required_cloud_config_fields` list.
// 
NscSupportedCloudConfigPeeringTypes []string `json:"nsc_supported_cloud_config_peering_types,omitempty"`

// NscRequiredCloudConfigFields A list of required attributes in the `cloud_config` of a corresponding
// `NetworkServiceConfig`.
// 
// For example: `"vlan"`, `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredCloudConfigFields []string `json:"nsc_required_cloud_config_fields,omitempty"`

// NscSupportedCloudConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"vlan"`, "`peering_type`",
// `"bgp_password"`, `"bgp_neighbor_address"`, ...
// 
NscSupportedCloudConfigFields []string `json:"nsc_supported_cloud_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkProductOffering) PolymorphicType() string {
return CloudNetworkProductOfferingType
}

// CloudNetworkProductOfferingPatch Cloud Network Product Offering
type CloudNetworkProductOfferingPatch struct {
// ID The *primary identifier* of the `Cloud Network Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs *string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork *string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea *string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

// ServiceProviderRegion The service provider offers the network service for a
// specific region.
// 
ServiceProviderRegion *string `json:"service_provider_region,omitempty"`

// ServiceProviderPop The datacenter id of the partner NNI to the service provider.
// It supposed to be used when identifying a location via
// the cloud provider's APIs.
// 
ServiceProviderPop *string `json:"service_provider_pop,omitempty"`

// ServiceProviderPopName The datacenter description of the partner NNI to the service provider.
// 
ServiceProviderPopName *string `json:"service_provider_pop_name,omitempty"`

// ServiceProviderWorkflow When the workflow is `provider_first` the subscriber creates
// a circuit with the cloud provider and provides a `cloud_key` for filtering
// the product-offerings.
// 
// If the workflow is `exchange_first` the IX will create
// the cloud circuit on the provider side.
// 
ServiceProviderWorkflow *string `json:"service_provider_workflow,omitempty"`

// ServiceExchangePops A list of object, referencing a `PointOfPresence`
// and providing additional path information, in case the services
// is tethered through another party.
ServiceExchangePops []*ServiceExchangePop `json:"service_exchange_pops,omitempty"`

// DeliveryMethod The exchange delivers the service over a `shared` or `dedicated` NNI.
DeliveryMethod *string `json:"delivery_method,omitempty"`

// Diversity The service can be delivered over multiple handovers from
// the exchange to the `service_provider`.
// The `diversity` denotes the number of handovers between the
// exchange and the service provider. A value of two signals a
// redundant service.
// 
// Only one network service configuration for each `handover` and
// `cloud_vlan` can be created.
Diversity *int `json:"diversity,omitempty"`

// NscSupportedCloudConfigPeeringTypes The supported peering types for the cloud network service.
// 
// In case selecting a peering type is required, the
// `peering_type` property will be in the
// `nsc_required_cloud_config_fields` list.
// 
NscSupportedCloudConfigPeeringTypes []string `json:"nsc_supported_cloud_config_peering_types,omitempty"`

// NscRequiredCloudConfigFields A list of required attributes in the `cloud_config` of a corresponding
// `NetworkServiceConfig`.
// 
// For example: `"vlan"`, `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredCloudConfigFields []string `json:"nsc_required_cloud_config_fields,omitempty"`

// NscSupportedCloudConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"vlan"`, "`peering_type`",
// `"bgp_password"`, `"bgp_neighbor_address"`, ...
// 
NscSupportedCloudConfigFields []string `json:"nsc_supported_cloud_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkProductOfferingPatch) PolymorphicType() string {
return CloudNetworkProductOfferingPatchType
}

// ConnectionProductOffering Connection Product Offering
type ConnectionProductOffering struct {
// ID The *primary identifier* of the `Connection Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// CrossConnectInitiator A cross connect can be initiated by either the
// exchange or the subscriber.
// 
// This property affects which side has to provide
// a LOA and demarc information.
CrossConnectInitiator string `json:"cross_connect_initiator,omitempty"`

// HandoverPop The ID of the point of presence (see `/pops`), where
// the physical port will be present.
// 
HandoverPop *string `json:"handover_pop,omitempty"`

// MaximumPortQuantity The maximum amount of ports which can be aggregated
// in the connection. `null` means no limit.
MaximumPortQuantity *int `json:"maximum_port_quantity,omitempty"`

// RequiredContactRoles The connection will require at least one of each of the
// specified roles assigned to contacts.
// 
// The role assignments are associated with the connection
// through the `role_assignments` list property.
RequiredContactRoles []string `json:"required_contact_roles,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c ConnectionProductOffering) PolymorphicType() string {
return ConnectionProductOfferingType
}

// ConnectionProductOfferingPatch Conncetion Product Offering
type ConnectionProductOfferingPatch struct {
// ID The *primary identifier* of the `Conncetion Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// CrossConnectInitiator A cross connect can be initiated by either the
// exchange or the subscriber.
// 
// This property affects which side has to provide
// a LOA and demarc information.
CrossConnectInitiator *string `json:"cross_connect_initiator,omitempty"`

// HandoverPop The ID of the point of presence (see `/pops`), where
// the physical port will be present.
// 
HandoverPop *string `json:"handover_pop,omitempty"`

// MaximumPortQuantity The maximum amount of ports which can be aggregated
// in the connection. `null` means no limit.
MaximumPortQuantity *int `json:"maximum_port_quantity,omitempty"`

// RequiredContactRoles The connection will require at least one of each of the
// specified roles assigned to contacts.
// 
// The role assignments are associated with the connection
// through the `role_assignments` list property.
RequiredContactRoles []string `json:"required_contact_roles,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c ConnectionProductOfferingPatch) PolymorphicType() string {
return ConnectionProductOfferingPatchType
}

// Device Device
type Device struct {
// ID The *primary identifier* of the `Device`.
ID string `json:"id,omitempty"`

// Name Name of the device
// 
Name string `json:"name,omitempty"`

// Pop The `PointOfPresence` the device is in.
Pop string `json:"pop,omitempty"`

// Capabilities is a capabilities
Capabilities []*DeviceCapability `json:"capabilities,omitempty"`

// Facility Identifier of the facility where the device
// is physically based.
Facility string `json:"facility,omitempty"`

}

// DeviceCapability Device Capability
type DeviceCapability struct {
// MediaType The media type of the port (e.g. 1000BASE-LX, 10GBASE-LR, ...)
// 
MediaType string `json:"media_type,omitempty"`

// Speed Speed of port in Mbit/s
// 
Speed int `json:"speed,omitempty"`

// MaxLag Maximum count of ports which can be bundled to a max_lag
MaxLag int `json:"max_lag,omitempty"`

// Availability Count of available ports on device
// 
Availability int `json:"availability,omitempty"`

}

// DeviceConnection Device Connection
type DeviceConnection struct {
// CapacityMax is a capacity_max
CapacityMax int `json:"capacity_max,omitempty"`

// Device The `id` of the related `Device`.
// 
// 
Device string `json:"device,omitempty"`

// ConnectedDevice The `id` of the related `Device`.
// 
// 
ConnectedDevice string `json:"connected_device,omitempty"`

// ID The *primary identifier* of the `Device Connection`.
ID string `json:"id,omitempty"`

}

// ExchangeLanNetworkProductOffering Exchange Lan Network Product Offering
type ExchangeLanNetworkProductOffering struct {
// ID The *primary identifier* of the `Exchange Lan Network Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

// ExchangeLanNetworkService The id of the exchange lan network service.
ExchangeLanNetworkService string `json:"exchange_lan_network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkProductOffering) PolymorphicType() string {
return ExchangeLanNetworkProductOfferingType
}

// ExchangeLanNetworkProductOfferingPatch Exchange Lan Network Product Offering
type ExchangeLanNetworkProductOfferingPatch struct {
// ID The *primary identifier* of the `Exchange Lan Network Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs *string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork *string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea *string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

// ExchangeLanNetworkService The id of the exchange lan network service.
ExchangeLanNetworkService *string `json:"exchange_lan_network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkProductOfferingPatch) PolymorphicType() string {
return ExchangeLanNetworkProductOfferingPatchType
}

// Facility Facility
type Facility struct {
// ID The *primary identifier* of the `Facility`.
ID string `json:"id,omitempty"`

// Name Name of the Datacenter as called by the operator
// 
Name string `json:"name,omitempty"`

// MetroArea Id of the `MetroArea` the DC is located in.
// 
MetroArea string `json:"metro_area,omitempty"`

// AddressCountry ISO 3166-1 alpha-2 country code, for example DE
// 
AddressCountry string `json:"address_country,omitempty"`

// AddressLocality The locality/city. For example, Mountain View.
AddressLocality string `json:"address_locality,omitempty"`

// AddressRegion The region. For example, CA
AddressRegion string `json:"address_region,omitempty"`

// PostalCode A postal code. For example, 9404
PostalCode string `json:"postal_code,omitempty"`

// StreetAddress The street address. For example, 1600 Amphitheatre Pkwy.
StreetAddress string `json:"street_address,omitempty"`

// PeeringdbFacilityID [PeeringDB](https://www.peeringdb.com) facitlity ID,
// can be extracted from the url https://www.peeringdb.com/fac/$id
// 
PeeringdbFacilityID *int `json:"peeringdb_facility_id,omitempty"`

// OrganisationName Name of Datacenter operator
// 
OrganisationName string `json:"organisation_name,omitempty"`

// Pops List of pops reachable from the `Facility`.
Pops []string `json:"pops,omitempty"`

// Latitude Latitude of the facility's location.
// 
Latitude *float64 `json:"latitude,omitempty"`

// Longitude Longitude of the facility's location.
// 
Longitude *float64 `json:"longitude,omitempty"`

}

// MP2MPNetworkProductOffering MP2MP Network Product Offering
type MP2MPNetworkProductOffering struct {
// ID The *primary identifier* of the `MP2MP Network Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkProductOffering) PolymorphicType() string {
return MP2MPNetworkProductOfferingType
}

// MP2MPNetworkProductOfferingPatch MP2MP Network Product Offering
type MP2MPNetworkProductOfferingPatch struct {
// ID The *primary identifier* of the `MP2MP Network Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs *string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork *string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea *string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkProductOfferingPatch) PolymorphicType() string {
return MP2MPNetworkProductOfferingPatchType
}

// MetroArea MetroArea
type MetroArea struct {
// ID The *primary identifier* of the `MetroArea`.
ID string `json:"id,omitempty"`

// UnLocode The UN/LOCODE for identifying the metro area.
// 
UnLocode string `json:"un_locode,omitempty"`

// IataCode The three letter IATA airport code for identiying the
// metro area.
// 
IataCode string `json:"iata_code,omitempty"`

// DisplayName The name of the metro area. Likely the same as the IATA code.
// 
DisplayName string `json:"display_name,omitempty"`

// Facilities List of facilities the metro area network.
Facilities []string `json:"facilities,omitempty"`

// MetroAreaNetworks List of networks in the metro area.
MetroAreaNetworks []string `json:"metro_area_networks,omitempty"`

}

// MetroAreaNetwork MetroAreaNetwork
type MetroAreaNetwork struct {
// ID The *primary identifier* of the `MetroAreaNetwork`.
ID string `json:"id,omitempty"`

// Name The name of the metro area network.
// 
Name string `json:"name,omitempty"`

// MetroArea The id of the metro area.
// 
MetroArea string `json:"metro_area,omitempty"`

// ServiceProvider The service provider is operating the network.
// Usually the exchange.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// Pops List of pops in the metro area network.
Pops []string `json:"pops,omitempty"`

}

// P2MPNetworkProductOffering P2MP Network Product Offering
type P2MPNetworkProductOffering struct {
// ID The *primary identifier* of the `P2MP Network Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkProductOffering) PolymorphicType() string {
return P2MPNetworkProductOfferingType
}

// P2MPNetworkProductOfferingPatch P2MP Network Product Offering
type P2MPNetworkProductOfferingPatch struct {
// ID The *primary identifier* of the `P2MP Network Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs *string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork *string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea *string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkProductOfferingPatch) PolymorphicType() string {
return P2MPNetworkProductOfferingPatchType
}

// P2PNetworkProductOffering P2P Network Product Offering
type P2PNetworkProductOffering struct {
// ID The *primary identifier* of the `P2P Network Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkProductOffering) PolymorphicType() string {
return P2PNetworkProductOfferingType
}

// P2PNetworkProductOfferingPatch P2P Network Product Offering
type P2PNetworkProductOfferingPatch struct {
// ID The *primary identifier* of the `P2P Network Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// ProviderVLANs The `NetworkService` provides `single` or `multi`ple vlans.
ProviderVLANs *string `json:"provider_vlans,omitempty"`

// ServiceMetroAreaNetwork Id of the `MetroAreaNetwork`.
// The service is directly provided on the metro area network.
// 
// In case of a `p2p_vc`, the `service_metro_area_network` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroAreaNetwork *string `json:"service_metro_area_network,omitempty"`

// ServiceMetroArea Id of the `MetroArea`. The service is delivered
// in this metro area.
// 
// In case of a `p2p_vc`, the `service_metro_area` refers
// to the B-side of the point-to-point connection.
// The B-side is the accepting party.
// 
ServiceMetroArea *string `json:"service_metro_area,omitempty"`

// BandwidthMin When configuring access to the network service, at least
// this `capacity` must be provided.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax When not `null`, this value enforces a mandatory
// rate limit for all network service configs.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

// NscRequiredL3ConfigFields A list of required attributes in the `l3_config` of a corresponding
// `NetworkServiceConfig` when used with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"bgp_neighbor_asn"`, `"local_address_primary"`, ...
// 
NscRequiredL3ConfigFields []string `json:"nsc_required_l3_config_fields,omitempty"`

// NscSupportedL3ConfigFields The list of fields which are supported in the `l3_config`
// when creating the network service config with a `routing_function`.
// 
// For example:  `"bgp_password"`, `"bgp_neighbor_address"`
// `"local_address_primary"`, ...
// 
NscSupportedL3ConfigFields []string `json:"nsc_supported_l3_config_fields,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkProductOfferingPatch) PolymorphicType() string {
return P2PNetworkProductOfferingPatchType
}

// PointOfPresence Point Of Presence
type PointOfPresence struct {
// Name is a name
Name string `json:"name,omitempty"`

// Facility The pop is located in this `Facility`.
Facility string `json:"facility,omitempty"`

// MetroAreaNetwork The `id` of the related `MetroAreaNetwork`.
// 
// 
MetroAreaNetwork string `json:"metro_area_network,omitempty"`

// Devices A list of `id`s of the related `Device`.
// 
// 
Devices []string `json:"devices,omitempty"`

// ID The *primary identifier* of the `Point Of Presence`.
ID string `json:"id,omitempty"`

// AvailabilityZone Availability zone of the pop.
AvailabilityZone *string `json:"availability_zone,omitempty"`

}

// ProductOffering Polymorphic Product Offering
type ProductOffering struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// ProductOfferingPatch Polymorphic Product Offering
type ProductOfferingPatch struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// RoutingFunctionProductOffering Routing Function Product Offering
type RoutingFunctionProductOffering struct {
// ID The *primary identifier* of the `Routing Function Product Offering`.
ID string `json:"id,omitempty"`

// Name Name of the product
Name string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// BandwidthMin The minimum bandwidth of the routing service in Mbit/s.
BandwidthMin int `json:"bandwidth_min,omitempty"`

// BandwidthMax The maximum bandwidth of the routing service in Mbit/s.
BandwidthMax int `json:"bandwidth_max,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RoutingFunctionProductOffering) PolymorphicType() string {
return RoutingFunctionProductOfferingType
}

// RoutingFunctionProductOfferingPatch Routing Function Product Offering
type RoutingFunctionProductOfferingPatch struct {
// ID The *primary identifier* of the `Routing Function Product Offering`.
ID *string `json:"id,omitempty"`

// Name Name of the product
Name *string `json:"name,omitempty"`

// DisplayName is a display_name
DisplayName *string `json:"display_name,omitempty"`

// ExchangeLogo An URI referencing the logo of the internet exchange.
// 
ExchangeLogo *string `json:"exchange_logo,omitempty"`

// ServiceProviderLogo An URI referencing the logo of the service provider.
// 
ServiceProviderLogo *string `json:"service_provider_logo,omitempty"`

// ProductLogo An URI referencing a logo for the product offered.
// 
ProductLogo *string `json:"product_logo,omitempty"`

// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType *string `json:"resource_type,omitempty"`

// HandoverMetroAreaNetwork Id of the `MetroAreaNetwork`. The service will be accessed
// through the handover metro area network.
// 
// In case of a `p2p_vc`, the `handover_metro_area_network` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroAreaNetwork *string `json:"handover_metro_area_network,omitempty"`

// HandoverMetroArea Id of the `MetroArea`. The network service will be
// accessed from this metro area.
// 
// In case of a `p2p_vc`, the `handover_metro_area` refers
// to the A-side of the point-to-point connection.
// The A-side is the entity which initiates the network service creation.
// 
HandoverMetroArea *string `json:"handover_metro_area,omitempty"`

// PhysicalPortSpeed If the service is dependent on the speed of
// the physical port this field denotes the speed.
PhysicalPortSpeed *int `json:"physical_port_speed,omitempty"`

// ServiceProvider The name of the provider providing the service.
// 
ServiceProvider *string `json:"service_provider,omitempty"`

// DowngradeAllowed Indicates if the service can be migrated to
// a lower bandwidth.
DowngradeAllowed *bool `json:"downgrade_allowed,omitempty"`

// UpgradeAllowed Indicates if the service can be migrated to
// a higher bandwidth.
UpgradeAllowed *bool `json:"upgrade_allowed,omitempty"`

// OrderableNotBefore This product offering becomes available for ordering after
// this point in time.
OrderableNotBefore *time.Time `json:"orderable_not_before,omitempty"`

// OrderableNotAfter This product offering will become unavailable for ordering after
// this point in time.
OrderableNotAfter *time.Time `json:"orderable_not_after,omitempty"`

// ContractTerms This property informally describe the contract's notice- and
// renewal periods as well as additional terms.
// 
// **Note**: This property contains informal information about
// the contract. For a structured representation see:
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// **Example**: A contract with the terms
// _"initially two weeks, renewing for six month afterwards, cancelable with a notice period of one month after and within 5 days during the initial period"_
// can be represented as:
// * `contract_initial_period: "P2W"`
// * `contract_initial_notice_period: "P5D"`
// * `contract_renewal_period: "P6M"`
// * `contract_renewal_notice_period: "P1M"`
// 
ContractTerms *string `json:"contract_terms,omitempty"`

// ContractInitialPeriod _**Format:** ISO8601 Duration_
// 
// The initial duration of the contract. The contract will be
// renewed after this period for the duration of `contract_renewal_period`.
// 
ContractInitialPeriod *string `json:"contract_initial_period,omitempty"`

// ContractInitialNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period for canceling the contract within
// the initial period.
// 
ContractInitialNoticePeriod *string `json:"contract_initial_notice_period,omitempty"`

// ContractRenewalPeriod _**Format:** ISO8601 Duration_
// 
// The duration for which the contract will be renewed after
// the initial period.
// 
// Unless the contract is canceled, it will be
// automatically renewed after the period.
// Cancellation has to be done within the
// `contract_renewal_notice_period`.
// 
ContractRenewalPeriod *string `json:"contract_renewal_period,omitempty"`

// ContractRenewalNoticePeriod _**Format:** ISO8601 Duration_
// 
// The notice period denotes the time before the end of the
// `contract_renewal_period` in which the client has to inform
// the IXP in order to prevent renewal of the contract.
// 
ContractRenewalNoticePeriod *string `json:"contract_renewal_notice_period,omitempty"`

// NoticePeriod **DEPRECATION NOTICE**: This property will be replaced by
// `contract_initial_period`, `contract_initial_notice_period`,
// `contract_renewal_period` and `contract_renewal_notice_period`.
// 
// The notice period informally states constraints
// which define when the client needs to inform the
// IXP in order to prevent renewal of the contract.
// 
NoticePeriod *string `json:"notice_period,omitempty"`

// BandwidthMin The minimum bandwidth of the routing service in Mbit/s.
BandwidthMin *int `json:"bandwidth_min,omitempty"`

// BandwidthMax The maximum bandwidth of the routing service in Mbit/s.
BandwidthMax *int `json:"bandwidth_max,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RoutingFunctionProductOfferingPatch) PolymorphicType() string {
return RoutingFunctionProductOfferingPatchType
}

// ServiceExchangePop Service Exchange PoP
// 
// A list of object, referencing a `PointOfPresence`
// and providing additional path information, in case the services
// is tethered through another party.
type ServiceExchangePop struct {
// Pop The `id` of the `PointOfPresence` the service is provided.
Pop string `json:"pop,omitempty"`

// PathInfo An *optional* text property that describes the path of the service
// where it is tethered through another party.
PathInfo *string `json:"path_info,omitempty"`

}

// CloudConfig The `CloudConfig` provides additional configuration for
// creating the `NetworkServiceConfig` on the cloud provider
// side.
// 
// The `nsc_required_cloud_config_fields` and
// `nsc_supported_cloud_config_fields` attributes of the
// `ProductOffering` specifies which fields are required
// or can be optionally supplied.
// 
// When creating the `NetworkServiceConfig` with a `l3_config`
// and a `routing_functions`, some required fields
// will automatically be derived from the `l3_config`
// and can be ommitted.
// 
// Values in the `cloud_config` will have precedence
// over the `l3_config`.
type CloudConfig struct {
// BGPPassword The password to use for BGP sessions.
BGPPassword *string `json:"bgp_password,omitempty"`

// BGPNeighborAddress The IP address of the BGP neighbor.
BGPNeighborAddress *string `json:"bgp_neighbor_address,omitempty"`

// BGPNeighborAddressPrimary The primary IP address of the BGP neighbor.
BGPNeighborAddressPrimary *string `json:"bgp_neighbor_address_primary,omitempty"`

// BGPNeighborAddressSecondary The secondary IP address of the BGP neighbor.
BGPNeighborAddressSecondary *string `json:"bgp_neighbor_address_secondary,omitempty"`

// BGPNeighborASN The ASN of the BGP neighbor.
BGPNeighborASN *int `json:"bgp_neighbor_asn,omitempty"`

// BGPAddressFamily is a bgp_address_family
BGPAddressFamily *string `json:"bgp_address_family,omitempty"`

// Bfd Enable BFD for the BGP session.
Bfd *bool `json:"bfd,omitempty"`

// LocalASN The local ASN.
LocalASN *int `json:"local_asn,omitempty"`

// LocalAddress The IP address of the router function instance
// in CIDR notation.
LocalAddress *string `json:"local_address,omitempty"`

// LocalAddressPrimary The primary IP address of the router function instance
// in CIDR notation.
LocalAddressPrimary *string `json:"local_address_primary,omitempty"`

// LocalAddressSecondary The secondary IP address of the router function instance
// in CIDR notation.
LocalAddressSecondary *string `json:"local_address_secondary,omitempty"`

// VLAN If the `provider_vlans` property of the `ProductOffering` is
// `multi`, a numeric value refers to a specific vlan on the service
// provider side.
// 
// The `nsc_required_cloud_config_fields` attribute of the
// `ProductOffering` will include `vlan` if `provider_vlans` are
// `multi`.
VLAN *int `json:"vlan,omitempty"`

// PeeringType Some `cloud_vc` network services require selecting
// a peering type.
// 
// See the `nsc_supported_cloud_config_peering_types` attribute of the corresponding
// `ProductOffering` for valid values.
PeeringType *string `json:"peering_type,omitempty"`

}

// CloudNetworkServiceConfig Cloud Network Service Config
type CloudNetworkServiceConfig struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// NetworkService The id of the configured `NetworkService`.
NetworkService string `json:"network_service,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Config`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a polymorphic vlan configuration
VLANConfig VLANConfig `tf:"vlan_config" json:"-"`
// VLANConfigRaw contains the vlan config response data
VLANConfigRaw json.RawMessage `tf:"-" json:"vlan_config,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Handover The handover enumerates the connection and is
// required for checking diversity constraints.
// 
// It must be within `1 <= x <= network_service.diversity`.
// 
Handover int `json:"handover,omitempty"`

// CloudConfig is a cloud_config
CloudConfig *CloudConfig `json:"cloud_config,omitempty"`

// CloudVLAN **Deprecation Notice**: This field is deprecated and will
// be removed in favor of using the `cloud_config.vlan` property.
// The `ProductOffering` will include `vlan` in the
// `nsc_required_cloud_config_fields`, if `provider_vlans` are
// `multi`.
// 
// If the `provider_vlans` property of the `ProductOffering` is
// `multi`, a numeric value refers to a specific vlan on the service
// provider side.
// 
// Otherwise, if set to `null`, it refers to all unmatched
// vlan ids on the service provider side. (All vlan ids from the
// service provider side are presented as tags within any vlans specified
// in `vlan_config`.)
// 
// If the `provider_vlans` property of the `ProductOffering` is `single`,
// the `cloud_vlan` MUST be `null` or MUST NOT be provided.
CloudVLAN *int `json:"cloud_vlan,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfig) PolymorphicType() string {
return CloudNetworkServiceConfigType
}

// CloudNetworkServiceConfigPatch Cloud Network Service Config Update
type CloudNetworkServiceConfigPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Config Update`.
ID *string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Handover The handover enumerates the connection and is
// required for checking diversity constraints.
// 
// It must be within `1 <= x <= network_service.diversity`.
// 
Handover *int `json:"handover,omitempty"`

// CloudConfig is a cloud_config
CloudConfig *CloudConfig `json:"cloud_config,omitempty"`

// CloudVLAN **Deprecation Notice**: This field is deprecated and will
// be removed in favor of using the `cloud_config.vlan` property.
// The `ProductOffering` will include `vlan` in the
// `nsc_required_cloud_config_fields`, if `provider_vlans` are
// `multi`.
// 
// If the `provider_vlans` property of the `ProductOffering` is
// `multi`, a numeric value refers to a specific vlan on the service
// provider side.
// 
// Otherwise, if set to `null`, it refers to all unmatched
// vlan ids on the service provider side. (All vlan ids from the
// service provider side are presented as tags within any vlans specified
// in `vlan_config`.)
// 
// If the `provider_vlans` property of the `ProductOffering` is `single`,
// the `cloud_vlan` MUST be `null` or MUST NOT be provided.
CloudVLAN *int `json:"cloud_vlan,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigPatch) PolymorphicType() string {
return CloudNetworkServiceConfigPatchType
}

// CloudNetworkServiceConfigRequest Cloud Network Service Config Request
type CloudNetworkServiceConfigRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// NetworkService The id of the `NetworkService` to configure.
NetworkService string `json:"network_service,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Config Request`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Handover The handover enumerates the connection and is
// required for checking diversity constraints.
// 
// It must be within `1 <= x <= network_service.diversity`.
// 
Handover int `json:"handover,omitempty"`

// CloudConfig is a cloud_config
CloudConfig *CloudConfig `json:"cloud_config,omitempty"`

// CloudVLAN **Deprecation Notice**: This field is deprecated and will
// be removed in favor of using the `cloud_config.vlan` property.
// The `ProductOffering` will include `vlan` in the
// `nsc_required_cloud_config_fields`, if `provider_vlans` are
// `multi`.
// 
// If the `provider_vlans` property of the `ProductOffering` is
// `multi`, a numeric value refers to a specific vlan on the service
// provider side.
// 
// Otherwise, if set to `null`, it refers to all unmatched
// vlan ids on the service provider side. (All vlan ids from the
// service provider side are presented as tags within any vlans specified
// in `vlan_config`.)
// 
// If the `provider_vlans` property of the `ProductOffering` is `single`,
// the `cloud_vlan` MUST be `null` or MUST NOT be provided.
CloudVLAN *int `json:"cloud_vlan,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigRequest) PolymorphicType() string {
return CloudNetworkServiceConfigRequestType
}

// CloudNetworkServiceConfigUpdate Cloud Network Service Config Update
type CloudNetworkServiceConfigUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Config Update`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Handover The handover enumerates the connection and is
// required for checking diversity constraints.
// 
// It must be within `1 <= x <= network_service.diversity`.
// 
Handover int `json:"handover,omitempty"`

// CloudConfig is a cloud_config
CloudConfig *CloudConfig `json:"cloud_config,omitempty"`

// CloudVLAN **Deprecation Notice**: This field is deprecated and will
// be removed in favor of using the `cloud_config.vlan` property.
// The `ProductOffering` will include `vlan` in the
// `nsc_required_cloud_config_fields`, if `provider_vlans` are
// `multi`.
// 
// If the `provider_vlans` property of the `ProductOffering` is
// `multi`, a numeric value refers to a specific vlan on the service
// provider side.
// 
// Otherwise, if set to `null`, it refers to all unmatched
// vlan ids on the service provider side. (All vlan ids from the
// service provider side are presented as tags within any vlans specified
// in `vlan_config`.)
// 
// If the `provider_vlans` property of the `ProductOffering` is `single`,
// the `cloud_vlan` MUST be `null` or MUST NOT be provided.
CloudVLAN *int `json:"cloud_vlan,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigUpdate) PolymorphicType() string {
return CloudNetworkServiceConfigUpdateType
}

// Connection Connection
type Connection struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Connection`.
ID string `json:"id,omitempty"`

// Mode Sets the mode of the connection. The mode can be:
// 
// - `lag_lacp`: connection is build as a LAG with LACP enabled
// - `lag_static`: connection is build as LAG with static configuration
// - `flex_ethernet`: connect is build as a FlexEthernet channel
// - `standalone`: only one port is allowed in this connection without
// any bundling.
// 
Mode string `json:"mode,omitempty"`

// LacpTimeout This sets the LACP Timeout mode. Both ends of the connections need
// to be configured the same.
// 
LacpTimeout *string `json:"lacp_timeout,omitempty"`

// ProductOffering The product offering must match the type `connection`.
ProductOffering string `json:"product_offering,omitempty"`

// Name is a name
Name string `json:"name,omitempty"`

// Ports References to the port belonging to this connection. Typically
// all ports within one connection are distributed over the same
// device.
// 
Ports []string `json:"ports,omitempty"`

// PortReservations A list of `port-reservations` for this connection.
PortReservations []string `json:"port_reservations,omitempty"`

// Pop The ID of the point of presence (see `/pops`), where
// the physical port(s) are present.
// 
Pop string `json:"pop,omitempty"`

// Speed Shows the total bandwidth of the connection in Mbit/s.
// 
Speed *int `json:"speed,omitempty"`

// CapacityAllocated Sum of the bandwidth of all network service configs
// using the connection in Mbit/s.
CapacityAllocated int `json:"capacity_allocated,omitempty"`

// CapacityAllocationLimit Maximum allocatable capacity of the connection in Mbit/s.
// When `null`, the exchange does not impose any limit.
// 
CapacityAllocationLimit int `json:"capacity_allocation_limit,omitempty"`

// VLANTypes A list of vlan config types you can configure using
// this connection.
VLANTypes []string `json:"vlan_types,omitempty"`

// OuterVLANEthertypes The ethertype of the outer tag in hexadecimal notation.
// 
OuterVLANEthertypes []string `json:"outer_vlan_ethertypes,omitempty"`

// PortQuantity The number of ports which should be allocated
// for this connection.
PortQuantity int `json:"port_quantity,omitempty"`

// SubscriberSideDemarcs The workflow for allocating ports is dependent on the
// `cross_connect_initiator` property of the
// `product_offering`:
// 
// **Cross-Connect initiator: exchange**
// 
// The subscriber needs to provide a
// list of demarc information.
// 
// 
// At least one needs to be provided, but not more than
// `port_quantity`.
// 
// The content is interpreted by the exchange and may
// contain a reference to a pre-existing cross connect order
// or information required for patching in a structured
// format (e.g.
// `<pp-identifier>.<hu-identifier>.<slot-identifier>.<port-identifier>`).
// 
// Please refer to the usage guide of the internet exchange.
// 
// ---
// 
// **Cross-Connect initiator: subscriber**
// 
// This field can be omitted, when the cross connect
// initiator is the `subscriber`.
SubscriberSideDemarcs []string `json:"subscriber_side_demarcs,omitempty"`

// MetroArea Optional ID of the service metro area the connection
// is provided in.
MetroArea *string `json:"metro_area,omitempty"`

// MetroAreaNetwork Optional ID of the service metro area network the
// connection is present on.
MetroAreaNetwork *string `json:"metro_area_network,omitempty"`

}

// ConnectionPatch Connection Update
type ConnectionPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Connection Update`.
ID *string `json:"id,omitempty"`

// Mode Sets the mode of the connection. The mode can be:
// 
// - `lag_lacp`: connection is build as a LAG with LACP enabled
// - `lag_static`: connection is build as LAG with static configuration
// - `flex_ethernet`: connect is build as a FlexEthernet channel
// - `standalone`: only one port is allowed in this connection without
// any bundling.
// 
Mode *string `json:"mode,omitempty"`

// LacpTimeout This sets the LACP Timeout mode. Both ends of the connections need
// to be configured the same.
// 
LacpTimeout *string `json:"lacp_timeout,omitempty"`

// ProductOffering The product offering must match the type `connection`.
ProductOffering *string `json:"product_offering,omitempty"`

}

// ConnectionRequest Request a new connection
type ConnectionRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Request a new connection`.
ID string `json:"id,omitempty"`

// Mode Sets the mode of the connection. The mode can be:
// 
// - `lag_lacp`: connection is build as a LAG with LACP enabled
// - `lag_static`: connection is build as LAG with static configuration
// - `flex_ethernet`: connect is build as a FlexEthernet channel
// - `standalone`: only one port is allowed in this connection without
// any bundling.
// 
Mode string `json:"mode,omitempty"`

// LacpTimeout This sets the LACP Timeout mode. Both ends of the connections need
// to be configured the same.
// 
LacpTimeout *string `json:"lacp_timeout,omitempty"`

// ProductOffering The product offering must match the type `connection`.
ProductOffering string `json:"product_offering,omitempty"`

// PortQuantity The number of `PortReservation`s that will be
// created for this connection.
PortQuantity int `json:"port_quantity,omitempty"`

// SubscriberSideDemarcs The workflow for allocating ports is dependent on the
// `cross_connect_initiator` property of the
// `product_offering`:
// 
// **Cross-Connect initiator: exchange**
// 
// The subscriber needs to provide a
// list of demarc information.
// 
// At least one needs to be provided, but not more than
// `port_quantity`.
// 
// The content is interpreted by the exchange and may
// contain a reference to a pre-existing cross connect order
// or information required for patching in a structured
// format (e.g.
// `<pp-identifier>.<hu-identifier>.<slot-identifier>.<port-identifier>`).
// 
// Please refer to the usage guide of the internet exchange.
// 
// ---
// 
// **Cross-Connect initiator: subscriber**
// 
// This field can be omitted, when the cross connect
// initiator is the `subscriber`.
SubscriberSideDemarcs []string `json:"subscriber_side_demarcs,omitempty"`

// ConnectingParty Name of the service provider who establishes
// connectivity on your behalf.
// 
// This is only relevant, if the cross connect initiator
// is the `subscriber` and might be `null`.
// 
// Please refer to the usage guide of the internet exchange.
ConnectingParty *string `json:"connecting_party,omitempty"`

}

// ConnectionUpdate Connection Update
type ConnectionUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Connection Update`.
ID string `json:"id,omitempty"`

// Mode Sets the mode of the connection. The mode can be:
// 
// - `lag_lacp`: connection is build as a LAG with LACP enabled
// - `lag_static`: connection is build as LAG with static configuration
// - `flex_ethernet`: connect is build as a FlexEthernet channel
// - `standalone`: only one port is allowed in this connection without
// any bundling.
// 
Mode string `json:"mode,omitempty"`

// LacpTimeout This sets the LACP Timeout mode. Both ends of the connections need
// to be configured the same.
// 
LacpTimeout *string `json:"lacp_timeout,omitempty"`

// ProductOffering The product offering must match the type `connection`.
ProductOffering string `json:"product_offering,omitempty"`

}

// ExchangeLanNetworkServiceConfig Exchange Lan Network Service Config
type ExchangeLanNetworkServiceConfig struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// NetworkService The id of the configured `NetworkService`.
NetworkService string `json:"network_service,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Exchange Lan Network Service Config`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a polymorphic vlan configuration
VLANConfig VLANConfig `tf:"vlan_config" json:"-"`
// VLANConfigRaw contains the vlan config response data
VLANConfigRaw json.RawMessage `tf:"-" json:"vlan_config,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// Listed The customer wants to be featured on the member list
Listed bool `json:"listed,omitempty"`

// ConsumerSideReady You can use this optional property to signal to the
// IXP, that your equipment is set up and ready to be
// tested.
// *(Sensitive Property)*
ConsumerSideReady *bool `json:"consumer_side_ready,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// 
// Availability Zones may not be supported for exchange_lan because by
// default they span multiple networks.
// 
// If an availability zone is set then this refers to a circuit that
// is placed on a specific on-ramp to the exchange_lan.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

// SharedStatistics is a shared_statistics
SharedStatistics *SharedStatisticsConfig `json:"shared_statistics,omitempty"`

// ProductOffering The product offering must match the type `exchange_lan`
// and must refer to the related network service through
// the `exchange_lan_network_service` property.
ProductOffering string `json:"product_offering,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkServiceConfig) PolymorphicType() string {
return ExchangeLanNetworkServiceConfigType
}

// ExchangeLanNetworkServiceConfigPatch Exchange Lan Network Service Config Update
type ExchangeLanNetworkServiceConfigPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Exchange Lan Network Service Config Update`.
ID *string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// Listed The customer wants to be featured on the member list
Listed *bool `json:"listed,omitempty"`

// ConsumerSideReady You can use this optional property to signal to the
// IXP, that your equipment is set up and ready to be
// tested.
// *(Sensitive Property)*
ConsumerSideReady *bool `json:"consumer_side_ready,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// 
// Availability Zones may not be supported for exchange_lan because by
// default they span multiple networks.
// 
// If an availability zone is set then this refers to a circuit that
// is placed on a specific on-ramp to the exchange_lan.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

// SharedStatistics is a shared_statistics
SharedStatistics *SharedStatisticsConfig `json:"shared_statistics,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkServiceConfigPatch) PolymorphicType() string {
return ExchangeLanNetworkServiceConfigPatchType
}

// ExchangeLanNetworkServiceConfigRequest Exchange Lan Network Service Config Request
type ExchangeLanNetworkServiceConfigRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// NetworkService The id of the `NetworkService` to configure.
NetworkService string `json:"network_service,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Exchange Lan Network Service Config Request`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// Listed The customer wants to be featured on the member list
Listed bool `json:"listed,omitempty"`

// ConsumerSideReady You can use this optional property to signal to the
// IXP, that your equipment is set up and ready to be
// tested.
// *(Sensitive Property)*
ConsumerSideReady *bool `json:"consumer_side_ready,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// 
// Availability Zones may not be supported for exchange_lan because by
// default they span multiple networks.
// 
// If an availability zone is set then this refers to a circuit that
// is placed on a specific on-ramp to the exchange_lan.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

// SharedStatistics is a shared_statistics
SharedStatistics *SharedStatisticsConfig `json:"shared_statistics,omitempty"`

// ProductOffering The product offering must match the type `exchange_lan`
// and must refer to the related network service through
// the `exchange_lan_network_service` property.
ProductOffering string `json:"product_offering,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkServiceConfigRequest) PolymorphicType() string {
return ExchangeLanNetworkServiceConfigRequestType
}

// ExchangeLanNetworkServiceConfigUpdate Exchange Lan Network Service Config Update
type ExchangeLanNetworkServiceConfigUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Exchange Lan Network Service Config Update`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// Listed The customer wants to be featured on the member list
Listed bool `json:"listed,omitempty"`

// ConsumerSideReady You can use this optional property to signal to the
// IXP, that your equipment is set up and ready to be
// tested.
// *(Sensitive Property)*
ConsumerSideReady *bool `json:"consumer_side_ready,omitempty"`

// AvailabilityZone The availability zone that shall be used on the provider side.
// 
// Availability Zones may not be supported for exchange_lan because by
// default they span multiple networks.
// 
// If an availability zone is set then this refers to a circuit that
// is placed on a specific on-ramp to the exchange_lan.
// *(Sensitive Property)*
AvailabilityZone *string `json:"availability_zone,omitempty"`

// SharedStatistics is a shared_statistics
SharedStatistics *SharedStatisticsConfig `json:"shared_statistics,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkServiceConfigUpdate) PolymorphicType() string {
return ExchangeLanNetworkServiceConfigUpdateType
}

// IXPSpecificFeatureFlagConfig IXP-Specific Feature Flag Configuration
type IXPSpecificFeatureFlagConfig struct {
// Name The name of the feature flag.
// 
Name string `json:"name,omitempty"`

// Enabled Enable the feature.
// 
// *Mandatory features can not be disabled*.
Enabled bool `json:"enabled,omitempty"`

}

// L3Config The layer 3 configuration for the NetworkServiceConfig.
// 
// It is **required** when a `routing_function` is provided.
// It may be required with a `connection`, depending on the
// `ProductOffering`.
// 
// Please check the `nsc_required_l3_config_fields` attribute
// of the corresponding `ProductOffering` to see which fields
// are required.
// 
// For additional optional fields, please check the
// `nsc_supported_l3_config_fields` attribute of the `ProductOffering`.
// 
// *(Sensitive Property)*
type L3Config struct {
// BGPPassword The password to use for BGP sessions.
BGPPassword *string `json:"bgp_password,omitempty"`

// BGPNeighborAddress The IP address of the BGP neighbor.
BGPNeighborAddress *string `json:"bgp_neighbor_address,omitempty"`

// BGPNeighborAddressPrimary The primary IP address of the BGP neighbor.
BGPNeighborAddressPrimary *string `json:"bgp_neighbor_address_primary,omitempty"`

// BGPNeighborAddressSecondary The secondary IP address of the BGP neighbor.
BGPNeighborAddressSecondary *string `json:"bgp_neighbor_address_secondary,omitempty"`

// BGPNeighborASN The ASN of the BGP neighbor.
BGPNeighborASN *int `json:"bgp_neighbor_asn,omitempty"`

// BGPAddressFamily is a bgp_address_family
BGPAddressFamily *string `json:"bgp_address_family,omitempty"`

// Bfd Enable BFD for the BGP session.
Bfd *bool `json:"bfd,omitempty"`

// LocalASN The local ASN.
LocalASN *int `json:"local_asn,omitempty"`

// LocalAddress The IP address of the router function instance
// in CIDR notation.
LocalAddress *string `json:"local_address,omitempty"`

// LocalAddressPrimary The primary IP address of the router function instance
// in CIDR notation.
LocalAddressPrimary *string `json:"local_address_primary,omitempty"`

// LocalAddressSecondary The secondary IP address of the router function instance
// in CIDR notation.
LocalAddressSecondary *string `json:"local_address_secondary,omitempty"`

}

// MP2MPNetworkServiceConfig MP2MP Network Service Config
type MP2MPNetworkServiceConfig struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// NetworkService The id of the configured `NetworkService`.
NetworkService string `json:"network_service,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Config`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a polymorphic vlan configuration
VLANConfig VLANConfig `tf:"vlan_config" json:"-"`
// VLANConfigRaw contains the vlan config response data
VLANConfigRaw json.RawMessage `tf:"-" json:"vlan_config,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfig) PolymorphicType() string {
return MP2MPNetworkServiceConfigType
}

// MP2MPNetworkServiceConfigPatch MP2MP Network Service Config Update
type MP2MPNetworkServiceConfigPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Config Update`.
ID *string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigPatch) PolymorphicType() string {
return MP2MPNetworkServiceConfigPatchType
}

// MP2MPNetworkServiceConfigRequest MP2MP Network Service Config Request
type MP2MPNetworkServiceConfigRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// NetworkService The id of the `NetworkService` to configure.
NetworkService string `json:"network_service,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Config Request`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigRequest) PolymorphicType() string {
return MP2MPNetworkServiceConfigRequestType
}

// MP2MPNetworkServiceConfigUpdate MP2MP Network Service Config Update
type MP2MPNetworkServiceConfigUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Config Update`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// IPs A list of ip-address IDs.
// 
// Allocation of IP Addresses might be deferred depending on
// the IXP implementation. No assumption should be made.
// *(Sensitive Property)*
IPs []string `json:"ips,omitempty"`

// ASNs A list of AS numbers.
// 
// Depending on the implementation, these can be used for different
// purposes.  For example in the members list on the website, links to
// the looking glass or even generating IPv6 prefixes.
ASNs []int `json:"asns,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigUpdate) PolymorphicType() string {
return MP2MPNetworkServiceConfigUpdateType
}

// NetworkFeatureConfig Polymorphic Network Feature Config
type NetworkFeatureConfig struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkFeatureConfigPatch Polymorphic Network Feauture Config Patch
type NetworkFeatureConfigPatch struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkFeatureConfigRequest Polymorphic Network Feature Config Request
type NetworkFeatureConfigRequest struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkFeatureConfigUpdate Polymorphic Network Feauture Config Update
type NetworkFeatureConfigUpdate struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceConfig Polymorphic Network Service Config
type NetworkServiceConfig struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceConfigPatch Polymorphic Network Service Config
type NetworkServiceConfigPatch struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceConfigRequest Polymorhic Network Service Config Request
type NetworkServiceConfigRequest struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceConfigUpdate Polymorphic Network Service Config
type NetworkServiceConfigUpdate struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// P2MPNetworkServiceConfig P2MP Network Service Config
type P2MPNetworkServiceConfig struct {
// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Config`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a polymorphic vlan configuration
VLANConfig VLANConfig `tf:"vlan_config" json:"-"`
// VLANConfigRaw contains the vlan config response data
VLANConfigRaw json.RawMessage `tf:"-" json:"vlan_config,omitempty"`

// Role A `leaf` can only reach roots and is
// isolated from other leafs. A `root` can
// reach any other point in the virtual circuit
// including other roots.
Role *string `json:"role,omitempty"`

// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// NetworkService The id of the configured `NetworkService`.
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceConfig) PolymorphicType() string {
return P2MPNetworkServiceConfigType
}

// P2MPNetworkServiceConfigPatch P2MP Network Service Config Update
type P2MPNetworkServiceConfigPatch struct {
// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Config Update`.
ID *string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Role A `leaf` can only reach roots and is
// isolated from other leafs. A `root` can
// reach any other point in the virtual circuit
// including other roots.
Role *string `json:"role,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceConfigPatch) PolymorphicType() string {
return P2MPNetworkServiceConfigPatchType
}

// P2MPNetworkServiceConfigRequest P2MP Network Service Config Request
type P2MPNetworkServiceConfigRequest struct {
// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Config Request`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Role A `leaf` can only reach roots and is
// isolated from other leafs. A `root` can
// reach any other point in the virtual circuit
// including other roots.
Role *string `json:"role,omitempty"`

// NetworkService The id of the `NetworkService` to configure.
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceConfigRequest) PolymorphicType() string {
return P2MPNetworkServiceConfigRequestType
}

// P2MPNetworkServiceConfigUpdate P2MP Network Service Config Update
type P2MPNetworkServiceConfigUpdate struct {
// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Config Update`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Role A `leaf` can only reach roots and is
// isolated from other leafs. A `root` can
// reach any other point in the virtual circuit
// including other roots.
Role *string `json:"role,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceConfigUpdate) PolymorphicType() string {
return P2MPNetworkServiceConfigUpdateType
}

// P2PNetworkServiceConfig P2P Network Service Config
type P2PNetworkServiceConfig struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// NetworkService The id of the configured `NetworkService`.
NetworkService string `json:"network_service,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Config`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a polymorphic vlan configuration
VLANConfig VLANConfig `tf:"vlan_config" json:"-"`
// VLANConfigRaw contains the vlan config response data
VLANConfigRaw json.RawMessage `tf:"-" json:"vlan_config,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfig) PolymorphicType() string {
return P2PNetworkServiceConfigType
}

// P2PNetworkServiceConfigPatch P2P Network Service Config Update
type P2PNetworkServiceConfigPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Config Update`.
ID *string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigPatch) PolymorphicType() string {
return P2PNetworkServiceConfigPatchType
}

// P2PNetworkServiceConfigRequest P2P Network Service Config Request
type P2PNetworkServiceConfigRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// NetworkService The id of the `NetworkService` to configure.
NetworkService string `json:"network_service,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Config Request`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigRequest) PolymorphicType() string {
return P2PNetworkServiceConfigRequestType
}

// P2PNetworkServiceConfigUpdate P2P Network Service Config Update
type P2PNetworkServiceConfigUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Config Update`.
ID string `json:"id,omitempty"`

// Connection The id of the connection to use for this `NetworkServiceConfig`.
// 
// If no connection is specified, you have to provide
// a routing function.
// 
// When a connection is provided, you also need to specify
// the `lan_config`. The `routing_function` attribute
// may not be used. Some network services may require the
// use of the `l3_config`, please check the
// `nsc_required_l3_config_fields` attribute of the
// `ProductOffering`.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
Connection *string `json:"connection,omitempty"`

// RoutingFunction The id of the `RoutingFunction` to use for this `NetworkServiceConfig`.
// 
// If no routing function is provided, you need to provide
// the connection to use.
// 
// When a routing function is provided, you also need to
// specify the `l3_config`. The `connection` attribute
// may not be used.
// 
// Connections ans Routing Functions are mutually exclusive.
// *(Sensitive Property)*
RoutingFunction *string `json:"routing_function,omitempty"`

// L3Config is a l3_config
L3Config *L3Config `json:"l3_config,omitempty"`

// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
// 
// *(Sensitive Property)*
NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

// VLANConfig is a vlan_config
VLANConfig VLANConfig `json:"vlan_config,omitempty"`

// Macs A list of MAC address IDs. You may have to register the
// address using the `macs_create` operation.
// *(Sensitive Property)*
Macs []string `json:"macs,omitempty"`

// ProductOffering An optional id of a `ProductOffering`.
// 
// Valid ids of product-offerings can be found in the
// `nsc_product_offerings` property of the `NetworkService`.
ProductOffering *string `json:"product_offering,omitempty"`

// Capacity The capacity of the service in Mbps. If set to Null,
// the maximum capacity will be used, i.e. the virtual circuit is
// not rate-limited.
// 
// An exchange may choose to constrain the available capacity range
// of a `ProductOffering`.
// 
// That means, the service can consume up to the total bandwidth
// of the `Connection`.
// 
// Typically the service is charged based on the capacity.
// *(Sensitive Property)*
Capacity *int `json:"capacity,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigUpdate) PolymorphicType() string {
return P2PNetworkServiceConfigUpdateType
}

// Port Port
type Port struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ID The *primary identifier* of the `Port`.
ID string `json:"id,omitempty"`

// Connection The `id` of the related `Connection`.
// 
// 
Connection *string `json:"connection,omitempty"`

// Speed is a speed
Speed *int `json:"speed,omitempty"`

// Name Name of the port (set by the exchange)
Name *string `json:"name,omitempty"`

// MediaType The media type of the port.
// Query the device's capabilities for available types.
// 
MediaType string `json:"media_type,omitempty"`

// OperationalState The operational state of the port.
OperationalState *string `json:"operational_state,omitempty"`

// Device The device the port.
// 
Device string `json:"device,omitempty"`

// Pop Same as the `pop` of the `device`.
// 
Pop string `json:"pop,omitempty"`

}

// PortReservation A PortReservation
type PortReservation struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `A PortReservation`.
ID string `json:"id,omitempty"`

// SubscriberSideDemarc In an exchange initiated scenario, this field will
// indicated one of the provided `subscriber_side_demarcs`
// from the connection.
SubscriberSideDemarc *string `json:"subscriber_side_demarc,omitempty"`

// ConnectingParty Name of the service provider who establishes
// connectivity on your behalf.
// 
// This is only relevant, if the cross connect initiator
// is the `subscriber`.
// 
// Please refer to the usage guide of the internet exchange.
ConnectingParty *string `json:"connecting_party,omitempty"`

// CrossConnectID An optional identifier of a cross connect.
CrossConnectID *string `json:"cross_connect_id,omitempty"`

// Connection The `Port` will become part of this connection.
Connection string `json:"connection,omitempty"`

// ExchangeSideDemarc Exchange side demarc information. This field will only
// be filled in when the port state is `allocated` or
// in `production`.
// 
// Otherwise this field will be `null`.
ExchangeSideDemarc *string `json:"exchange_side_demarc,omitempty"`

// Port This field will be null, until a port will
// be allocated.
Port *string `json:"port,omitempty"`

}

// PortReservationPatch PortReservation Update
type PortReservationPatch struct {
// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `PortReservation Update`.
ID *string `json:"id,omitempty"`

// SubscriberSideDemarc In an exchange initiated scenario, this field will
// indicated one of the provided `subscriber_side_demarcs`
// from the connection.
SubscriberSideDemarc *string `json:"subscriber_side_demarc,omitempty"`

// ConnectingParty Name of the service provider who establishes
// connectivity on your behalf.
// 
// This is only relevant, if the cross connect initiator
// is the `subscriber`.
// 
// Please refer to the usage guide of the internet exchange.
ConnectingParty *string `json:"connecting_party,omitempty"`

// CrossConnectID An optional identifier of a cross connect.
CrossConnectID *string `json:"cross_connect_id,omitempty"`

}

// PortReservationRequest A PortReservation
type PortReservationRequest struct {
// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `A PortReservation`.
ID string `json:"id,omitempty"`

// SubscriberSideDemarc In an exchange initiated scenario, this field will
// indicated one of the provided `subscriber_side_demarcs`
// from the connection.
SubscriberSideDemarc *string `json:"subscriber_side_demarc,omitempty"`

// ConnectingParty Name of the service provider who establishes
// connectivity on your behalf.
// 
// This is only relevant, if the cross connect initiator
// is the `subscriber`.
// 
// Please refer to the usage guide of the internet exchange.
ConnectingParty *string `json:"connecting_party,omitempty"`

// CrossConnectID An optional identifier of a cross connect.
CrossConnectID *string `json:"cross_connect_id,omitempty"`

// Connection A connection is required for port allocation.
Connection string `json:"connection,omitempty"`

}

// PortReservationUpdate PortReservation Update
type PortReservationUpdate struct {
// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `PortReservation Update`.
ID string `json:"id,omitempty"`

// SubscriberSideDemarc In an exchange initiated scenario, this field will
// indicated one of the provided `subscriber_side_demarcs`
// from the connection.
SubscriberSideDemarc *string `json:"subscriber_side_demarc,omitempty"`

// ConnectingParty Name of the service provider who establishes
// connectivity on your behalf.
// 
// This is only relevant, if the cross connect initiator
// is the `subscriber`.
// 
// Please refer to the usage guide of the internet exchange.
ConnectingParty *string `json:"connecting_party,omitempty"`

// CrossConnectID An optional identifier of a cross connect.
CrossConnectID *string `json:"cross_connect_id,omitempty"`

}

// RouteServerNetworkFeatureConfig Route Server Network Feature Config
type RouteServerNetworkFeatureConfig struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Route Server Network Feature Config`.
ID string `json:"id,omitempty"`

// NetworkFeature The `id` of the related `NetworkFeature`.
// 
// 
NetworkFeature string `json:"network_feature,omitempty"`

// NetworkServiceConfig The `id` of the related `NetworkServiceConfig`.
// 
// 
NetworkServiceConfig string `json:"network_service_config,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ASN The ASN of the peer.
// 
ASN int `json:"asn,omitempty"`

// Password The cleartext BGP session password
Password *string `json:"password,omitempty"`

// AsSetV4 AS-SET of the customer for IPv4 prefix filtering.
// This is used to generate filters on the router servers.
// 
// Only valid referenced prefixes within the AS-SET
// are allowed inbound to the route server. All other routes are
// filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v6` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV4 *string `json:"as_set_v4,omitempty"`

// AsSetV6 AS-SET of the customer for IPv6. This is used to generate filters
// on the router servers. Only valid referenced prefixes within
// the AS-SET are allowed inbound to the route server.
// All other routes are filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet6` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v4` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV6 *string `json:"as_set_v6,omitempty"`

// MaxPrefixV4 Announcing more than `max_prefix` IPv4 prefixes the bgp
// session will be droped.
// 
MaxPrefixV4 *int `json:"max_prefix_v4,omitempty"`

// MaxPrefixV6 Announcing more than `max_prefix` IPv6 prefixes the bgp
// session will be droped.
// 
MaxPrefixV6 *int `json:"max_prefix_v6,omitempty"`

// InsertIxpASN Insert the ASN of the exchange into the AS path. This function is only
// used in special cases. In 99% of all cases, it should be false.
// 
InsertIxpASN *bool `json:"insert_ixp_asn,omitempty"`

// SessionMode Set the session mode with the routeserver.
// 
SessionMode string `json:"session_mode,omitempty"`

// BGPSessionType The session type describes which of the both parties will open the
// connection. If set to passive, the customer router needs to open
// the connection. If its set to active, the route server will open
// the connection. The standard behavior on most exchanges is passive.
// 
BGPSessionType string `json:"bgp_session_type,omitempty"`

// IP The BGP session will be established from this IP address,
// referenced by ID.
// 
// Only IDs of IPs assigned to the corresponding network service
// config can be used.
IP string `json:"ip,omitempty"`

// Flags A list of IXP specific feature flag configs. This can be used
// to enable or disable a specific feature flag.
Flags []*IXPSpecificFeatureFlagConfig `json:"flags,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeatureConfig) PolymorphicType() string {
return RouteServerNetworkFeatureConfigType
}

// RouteServerNetworkFeatureConfigPatch Route Server Network Feature Config Update
type RouteServerNetworkFeatureConfigPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ASN The ASN of the peer.
// 
ASN *int `json:"asn,omitempty"`

// Password The cleartext BGP session password
Password *string `json:"password,omitempty"`

// AsSetV4 AS-SET of the customer for IPv4 prefix filtering.
// This is used to generate filters on the router servers.
// 
// Only valid referenced prefixes within the AS-SET
// are allowed inbound to the route server. All other routes are
// filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v6` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV4 *string `json:"as_set_v4,omitempty"`

// AsSetV6 AS-SET of the customer for IPv6. This is used to generate filters
// on the router servers. Only valid referenced prefixes within
// the AS-SET are allowed inbound to the route server.
// All other routes are filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet6` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v4` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV6 *string `json:"as_set_v6,omitempty"`

// MaxPrefixV4 Announcing more than `max_prefix` IPv4 prefixes the bgp
// session will be droped.
// 
MaxPrefixV4 *int `json:"max_prefix_v4,omitempty"`

// MaxPrefixV6 Announcing more than `max_prefix` IPv6 prefixes the bgp
// session will be droped.
// 
MaxPrefixV6 *int `json:"max_prefix_v6,omitempty"`

// InsertIxpASN Insert the ASN of the exchange into the AS path. This function is only
// used in special cases. In 99% of all cases, it should be false.
// 
InsertIxpASN *bool `json:"insert_ixp_asn,omitempty"`

// SessionMode Set the session mode with the routeserver.
// 
SessionMode *string `json:"session_mode,omitempty"`

// BGPSessionType The session type describes which of the both parties will open the
// connection. If set to passive, the customer router needs to open
// the connection. If its set to active, the route server will open
// the connection. The standard behavior on most exchanges is passive.
// 
BGPSessionType *string `json:"bgp_session_type,omitempty"`

// IP The BGP session will be established from this IP address,
// referenced by ID.
// 
// Only IDs of IPs assigned to the corresponding network service
// config can be used.
IP *string `json:"ip,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeatureConfigPatch) PolymorphicType() string {
return RouteServerNetworkFeatureConfigPatchType
}

// RouteServerNetworkFeatureConfigRequest Route Server Network Feature Config Request
type RouteServerNetworkFeatureConfigRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Route Server Network Feature Config Request`.
ID string `json:"id,omitempty"`

// NetworkFeature The `id` of the related `NetworkFeature`.
// 
// 
NetworkFeature string `json:"network_feature,omitempty"`

// NetworkServiceConfig The `id` of the related `NetworkServiceConfig`.
// 
// 
NetworkServiceConfig string `json:"network_service_config,omitempty"`

// RoleAssignments A set of `RoleAssignment`s. See the documentation
// on the specific `required_contact_roles`,
// `nfc_required_contact_roles` or `nsc_required_contact_roles`
// on what `RoleAssignment`s to provide.
// 
// Please note, that any contact role can additionally be provided.
// The presence of at least one of each required contact roles
// is necessary.
// 
// *(Sensitive Property)*
RoleAssignments []string `json:"role_assignments,omitempty"`

// ASN The ASN of the peer.
// 
ASN int `json:"asn,omitempty"`

// Password The cleartext BGP session password
Password *string `json:"password,omitempty"`

// AsSetV4 AS-SET of the customer for IPv4 prefix filtering.
// This is used to generate filters on the router servers.
// 
// Only valid referenced prefixes within the AS-SET
// are allowed inbound to the route server. All other routes are
// filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v6` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV4 *string `json:"as_set_v4,omitempty"`

// AsSetV6 AS-SET of the customer for IPv6. This is used to generate filters
// on the router servers. Only valid referenced prefixes within
// the AS-SET are allowed inbound to the route server.
// All other routes are filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet6` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v4` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV6 *string `json:"as_set_v6,omitempty"`

// MaxPrefixV4 Announcing more than `max_prefix` IPv4 prefixes the bgp
// session will be droped.
// 
MaxPrefixV4 *int `json:"max_prefix_v4,omitempty"`

// MaxPrefixV6 Announcing more than `max_prefix` IPv6 prefixes the bgp
// session will be droped.
// 
MaxPrefixV6 *int `json:"max_prefix_v6,omitempty"`

// InsertIxpASN Insert the ASN of the exchange into the AS path. This function is only
// used in special cases. In 99% of all cases, it should be false.
// 
InsertIxpASN *bool `json:"insert_ixp_asn,omitempty"`

// SessionMode Set the session mode with the routeserver.
// 
SessionMode string `json:"session_mode,omitempty"`

// BGPSessionType The session type describes which of the both parties will open the
// connection. If set to passive, the customer router needs to open
// the connection. If its set to active, the route server will open
// the connection. The standard behavior on most exchanges is passive.
// 
BGPSessionType string `json:"bgp_session_type,omitempty"`

// IP The BGP session will be established from this IP address,
// referenced by ID.
// 
// Only IDs of IPs assigned to the corresponding network service
// config can be used.
IP string `json:"ip,omitempty"`

// Flags A list of IXP specific feature flag configs. This can be used
// to enable or disable a specific feature flag.
Flags []*IXPSpecificFeatureFlagConfig `json:"flags,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeatureConfigRequest) PolymorphicType() string {
return RouteServerNetworkFeatureConfigRequestType
}

// RouteServerNetworkFeatureConfigUpdate Route Server Network Feature Config Update
type RouteServerNetworkFeatureConfigUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ASN The ASN of the peer.
// 
ASN int `json:"asn,omitempty"`

// Password The cleartext BGP session password
Password *string `json:"password,omitempty"`

// AsSetV4 AS-SET of the customer for IPv4 prefix filtering.
// This is used to generate filters on the router servers.
// 
// Only valid referenced prefixes within the AS-SET
// are allowed inbound to the route server. All other routes are
// filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v6` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV4 *string `json:"as_set_v4,omitempty"`

// AsSetV6 AS-SET of the customer for IPv6. This is used to generate filters
// on the router servers. Only valid referenced prefixes within
// the AS-SET are allowed inbound to the route server.
// All other routes are filtered.
// 
// This field is *required* if the route server network feature only
// supports the `af_inet6` address family.
// If multiple address families are supported, it is optional if the
// `as_set_v4` is provided.
// 
// Important: The format has to be: "AS-SET@IRR". IRR is the database
// where the AS-SET is registred. Typically used IRR's are RADB, RIPE,
// NTTCOM, APNIC, ALTDB, LEVEL3, ARIN, AFRINIC, LACNIC
// 
AsSetV6 *string `json:"as_set_v6,omitempty"`

// MaxPrefixV4 Announcing more than `max_prefix` IPv4 prefixes the bgp
// session will be droped.
// 
MaxPrefixV4 *int `json:"max_prefix_v4,omitempty"`

// MaxPrefixV6 Announcing more than `max_prefix` IPv6 prefixes the bgp
// session will be droped.
// 
MaxPrefixV6 *int `json:"max_prefix_v6,omitempty"`

// InsertIxpASN Insert the ASN of the exchange into the AS path. This function is only
// used in special cases. In 99% of all cases, it should be false.
// 
InsertIxpASN *bool `json:"insert_ixp_asn,omitempty"`

// SessionMode Set the session mode with the routeserver.
// 
SessionMode string `json:"session_mode,omitempty"`

// BGPSessionType The session type describes which of the both parties will open the
// connection. If set to passive, the customer router needs to open
// the connection. If its set to active, the route server will open
// the connection. The standard behavior on most exchanges is passive.
// 
BGPSessionType string `json:"bgp_session_type,omitempty"`

// IP The BGP session will be established from this IP address,
// referenced by ID.
// 
// Only IDs of IPs assigned to the corresponding network service
// config can be used.
IP string `json:"ip,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeatureConfigUpdate) PolymorphicType() string {
return RouteServerNetworkFeatureConfigUpdateType
}

// SharedStatisticsConfig Configure sharing of sensitive statistics with
// other members of the exchange.
// 
// Choosing `allow` will grant access to the shared
// statistics to all accounts on the exchange.
// 
// Accounts can be explicitly prevented from accessing the information,
// by adding them to the list of account ids in the
// `accounts_denied` attribute.
// 
// Choosing `deny` will prevent access to the shared statistics
// by other accounts on the exchange. However, selective
// access can be granted by adding them to the list of
// account ids in the `accounts_allowed` attribute.
// 
// This affects the visibility of `nsc_available_capacity`
// and `nsc_available_capacity_change_perc`
// on the `NetworkServiceConfig` statistics.
// 
// *(Sensitive Property)*
type SharedStatisticsConfig struct {
// Policy is a policy
Policy string `json:"policy,omitempty"`

}

// SharedStatisticsConfigAllow Shared Statistics Policy: Allow
type SharedStatisticsConfigAllow struct {
// NscAvailableCapacity **DEPRECATION NOTICE**: This field will be removed in
// favor of choosing a statistics sharing policy.
// 
// A list of account IDs who can see
// `nsc_available_capacity` and `nsc_available_capacity_change_perc`
// on the `NetworkServiceConfig` statistics.
// 
NscAvailableCapacity []string `json:"nsc_available_capacity,omitempty"`

// AccountsDenied An optional list of account IDs who may not access
// `nsc_available_capacity` and `nsc_available_capacity_change_perc`
// on the `NetworkServiceConfig` statistics.
// 
AccountsDenied []string `json:"accounts_denied,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (s SharedStatisticsConfigAllow) PolymorphicType() string {
return SharedStatisticsConfigAllowType
}

// SharedStatisticsConfigDeny Shared Statistics Policy: Deny
type SharedStatisticsConfigDeny struct {
// NscAvailableCapacity **DEPRECATION NOTICE**: This field will be removed in
// favor of choosing a statistics sharing policy.
// 
// A list of account IDs who can see
// `nsc_available_capacity` and `nsc_available_capacity_change_perc`
// on the `NetworkServiceConfig` statistics.
// 
NscAvailableCapacity []string `json:"nsc_available_capacity,omitempty"`

// AccountsAllowed An optional list of account IDs who is allowed to access
// `nsc_available_capacity` and `nsc_available_capacity_change_perc`
// on the `NetworkServiceConfig` statistics.
// 
// 
AccountsAllowed []string `json:"accounts_allowed,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (s SharedStatisticsConfigDeny) PolymorphicType() string {
return SharedStatisticsConfigDenyType
}

// VLANConfigDot1Q A Dot1Q vlan configuration
type VLANConfigDot1Q struct {
// VLAN A VLAN tag. If `null`, the IXP will auto-select
// a valid vlan-id.
// 
VLAN *int `json:"vlan,omitempty"`

// VLANEthertype The ethertype of the vlan in hexadecimal notation.
VLANEthertype *string `json:"vlan_ethertype,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (v VLANConfigDot1Q) PolymorphicType() string {
return VLANConfigDot1QType
}

// VLANConfigQinQ A QinQ vlan configuration
type VLANConfigQinQ struct {
// OuterVLAN The outer VLAN id.
// If `null`, the IXP will auto-select
// a valid vlan-id.
// 
OuterVLAN *int `json:"outer_vlan,omitempty"`

// OuterVLANEthertype The ethertype of the outer tag in hexadecimal notation.
OuterVLANEthertype *string `json:"outer_vlan_ethertype,omitempty"`

// InnerVLAN The inner VLAN id.
// 
InnerVLAN int `json:"inner_vlan,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (v VLANConfigQinQ) PolymorphicType() string {
return VLANConfigQinQType
}

// VLANConfig The vlan configuration defines how the service
// is made available on the connection.
// 
// Is is **required** when a `connection` is provided.
// 
// *(Sensitive Property)*
type VLANConfig struct {
// VLANType is a vlan_type
VLANType string `json:"vlan_type,omitempty"`

}

// Account Account
type Account struct {
// State is a state
State *string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// ID The *primary identifier* of the `Account`.
ID string `json:"id,omitempty"`

// ManagingAccount The `id` of a managing account. Can be used for creating
// a customer hierachy. *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// Name Name of the account, how it gets represented
// in e.g. a "customers list".
// 
Name string `json:"name,omitempty"`

// LegalName Legal name of the organisation.
// Only required when it's different from the account name.
// *(Sensitive Property)*
// 
LegalName *string `json:"legal_name,omitempty"`

// BillingInformation is a billing_information
BillingInformation *BillingInformation `json:"billing_information,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
ExternalRef *string `json:"external_ref,omitempty"`

// Discoverable The account will be included for all members of the
// ix in the list of accounts.
// 
// Only `id`, `name` and `present_in_metro_area_networks`
// are provided to other members.
Discoverable *bool `json:"discoverable,omitempty"`

// MetroAreaNetworkPresence Informal list of `MetroAreaNetwork` ids, indicating the
// presence to other accounts.
// The list is maintained by the account and can be empty.
// 
MetroAreaNetworkPresence []string `json:"metro_area_network_presence,omitempty"`

// Address is a address
Address *Address `json:"address,omitempty"`

// ASNs List of Autonomous System Numbers (ASNs) associated
// with this account through network service or
// network feature configurations.
// 
ASNs []int `json:"asns,omitempty"`

}

// AccountPatch Account Update
type AccountPatch struct {
// ID The *primary identifier* of the `Account Update`.
ID *string `json:"id,omitempty"`

// ManagingAccount The `id` of a managing account. Can be used for creating
// a customer hierachy. *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// Name Name of the account, how it gets represented
// in e.g. a "customers list".
// 
Name *string `json:"name,omitempty"`

// LegalName Legal name of the organisation.
// Only required when it's different from the account name.
// *(Sensitive Property)*
// 
LegalName *string `json:"legal_name,omitempty"`

// BillingInformation is a billing_information
BillingInformation *BillingInformation `json:"billing_information,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
ExternalRef *string `json:"external_ref,omitempty"`

// Discoverable The account will be included for all members of the
// ix in the list of accounts.
// 
// Only `id`, `name` and `present_in_metro_area_networks`
// are provided to other members.
Discoverable *bool `json:"discoverable,omitempty"`

// MetroAreaNetworkPresence Informal list of `MetroAreaNetwork` ids, indicating the
// presence to other accounts.
// The list is maintained by the account and can be empty.
// 
MetroAreaNetworkPresence []string `json:"metro_area_network_presence,omitempty"`

// Address is a address
Address *Address `json:"address,omitempty"`

}

// AccountRequest Account Request
type AccountRequest struct {
// ID The *primary identifier* of the `Account Request`.
ID string `json:"id,omitempty"`

// ManagingAccount The `id` of a managing account. Can be used for creating
// a customer hierachy. *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// Name Name of the account, how it gets represented
// in e.g. a "customers list".
// 
Name string `json:"name,omitempty"`

// LegalName Legal name of the organisation.
// Only required when it's different from the account name.
// *(Sensitive Property)*
// 
LegalName *string `json:"legal_name,omitempty"`

// BillingInformation is a billing_information
BillingInformation *BillingInformation `json:"billing_information,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
ExternalRef *string `json:"external_ref,omitempty"`

// Discoverable The account will be included for all members of the
// ix in the list of accounts.
// 
// Only `id`, `name` and `present_in_metro_area_networks`
// are provided to other members.
Discoverable *bool `json:"discoverable,omitempty"`

// MetroAreaNetworkPresence Informal list of `MetroAreaNetwork` ids, indicating the
// presence to other accounts.
// The list is maintained by the account and can be empty.
// 
MetroAreaNetworkPresence []string `json:"metro_area_network_presence,omitempty"`

// Address is a address
Address *Address `json:"address,omitempty"`

}

// AccountUpdate Account Update
type AccountUpdate struct {
// ID The *primary identifier* of the `Account Update`.
ID string `json:"id,omitempty"`

// ManagingAccount The `id` of a managing account. Can be used for creating
// a customer hierachy. *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// Name Name of the account, how it gets represented
// in e.g. a "customers list".
// 
Name string `json:"name,omitempty"`

// LegalName Legal name of the organisation.
// Only required when it's different from the account name.
// *(Sensitive Property)*
// 
LegalName *string `json:"legal_name,omitempty"`

// BillingInformation is a billing_information
BillingInformation *BillingInformation `json:"billing_information,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
ExternalRef *string `json:"external_ref,omitempty"`

// Discoverable The account will be included for all members of the
// ix in the list of accounts.
// 
// Only `id`, `name` and `present_in_metro_area_networks`
// are provided to other members.
Discoverable *bool `json:"discoverable,omitempty"`

// MetroAreaNetworkPresence Informal list of `MetroAreaNetwork` ids, indicating the
// presence to other accounts.
// The list is maintained by the account and can be empty.
// 
MetroAreaNetworkPresence []string `json:"metro_area_network_presence,omitempty"`

// Address is a address
Address *Address `json:"address,omitempty"`

}

// Address A postal address. *(Sensitive Property)*
type Address struct {
// Country ISO 3166-1 alpha-2 country code, for example DE
Country string `json:"country,omitempty"`

// Locality The locality/city. For example, Mountain View.
Locality string `json:"locality,omitempty"`

// Region The region. For example, CA
Region *string `json:"region,omitempty"`

// PostalCode A postal code. For example, 9404
PostalCode string `json:"postal_code,omitempty"`

// StreetAddress The street address. For example, 1600 Amphitheatre Pkwy.
StreetAddress string `json:"street_address,omitempty"`

// PostOfficeBoxNumber The post office box number for PO box addresses.
PostOfficeBoxNumber *string `json:"post_office_box_number,omitempty"`

}

// BillingInformation Optional information required for issuing invoices.
// Only accounts with `billing_information` present can be used
// as a `billing_account`. *(Sensitive Property)*
type BillingInformation struct {
// Name Name of the organisation receiving invoices.
// 
Name string `json:"name,omitempty"`

// Address is a address
Address *Address `json:"address,omitempty"`

// VatNumber Value-added tax number, required for
// european reverse charge system.
// 
VatNumber *string `json:"vat_number,omitempty"`

}

// Contact Contact
type Contact struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `Contact`.
ID string `json:"id,omitempty"`

// Name A name of a person or an organisation
Name *string `json:"name,omitempty"`

// Telephone The telephone number in E.164 Phone Number Formatting
Telephone *string `json:"telephone,omitempty"`

// Email The email of the legal company entity.
// 
Email *string `json:"email,omitempty"`

}

// ContactPatch Contact Update
type ContactPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `Contact Update`.
ID *string `json:"id,omitempty"`

// Name A name of a person or an organisation
Name *string `json:"name,omitempty"`

// Telephone The telephone number in E.164 Phone Number Formatting
Telephone *string `json:"telephone,omitempty"`

// Email The email of the legal company entity.
// 
Email *string `json:"email,omitempty"`

}

// ContactRequest Contact Create Request
type ContactRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `Contact Create Request`.
ID string `json:"id,omitempty"`

// Name A name of a person or an organisation
Name *string `json:"name,omitempty"`

// Telephone The telephone number in E.164 Phone Number Formatting
Telephone *string `json:"telephone,omitempty"`

// Email The email of the legal company entity.
// 
Email *string `json:"email,omitempty"`

}

// ContactUpdate Contact Update
type ContactUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `Contact Update`.
ID string `json:"id,omitempty"`

// Name A name of a person or an organisation
Name *string `json:"name,omitempty"`

// Telephone The telephone number in E.164 Phone Number Formatting
Telephone *string `json:"telephone,omitempty"`

// Email The email of the legal company entity.
// 
Email *string `json:"email,omitempty"`

}

// Role Role for a Contact
type Role struct {
// Name The name of the role.
// 
Name string `json:"name,omitempty"`

// RequiredFields A list of required field names.
// 
RequiredFields []string `json:"required_fields,omitempty"`

// ID The *primary identifier* of the `Role for a Contact`.
ID string `json:"id,omitempty"`

}

// RoleAssignment A role assignment for a contact
type RoleAssignment struct {
// Role The `id` of a role the contact is assigned to.
// 
Role string `json:"role,omitempty"`

// Contact The `id` of a contact the role is assigned to.
// 
Contact string `json:"contact,omitempty"`

// ID The *primary identifier* of the `A role assignment for a contact`.
ID string `json:"id,omitempty"`

}

// RoleAssignmentPatch A role assignment update
type RoleAssignmentPatch struct {
// Role The `id` of a role the contact is assigned to.
// 
Role *string `json:"role,omitempty"`

// Contact The `id` of a contact the role is assigned to.
// 
Contact *string `json:"contact,omitempty"`

}

// RoleAssignmentRequest A role assignment request
type RoleAssignmentRequest struct {
// Role The `id` of a role the contact is assigned to.
// 
Role string `json:"role,omitempty"`

// Contact The `id` of a contact the role is assigned to.
// 
Contact string `json:"contact,omitempty"`

}

// RoleAssignmentUpdate A role assignemnt update
type RoleAssignmentUpdate struct {
// Role The `id` of a role the contact is assigned to.
// 
Role string `json:"role,omitempty"`

// Contact The `id` of a contact the role is assigned to.
// 
Contact string `json:"contact,omitempty"`

}

// RolePatch Role Update
type RolePatch struct {
// Name The name of the role.
// 
Name *string `json:"name,omitempty"`

// RequiredFields A list of required field names.
// 
RequiredFields []string `json:"required_fields,omitempty"`

}

// RoleRequest Create Role
type RoleRequest struct {
// Name The name of the role.
// 
Name string `json:"name,omitempty"`

// RequiredFields A list of required field names.
// 
RequiredFields []string `json:"required_fields,omitempty"`

}

// RoleUpdate Role Update
type RoleUpdate struct {
// Name The name of the role.
// 
Name string `json:"name,omitempty"`

// RequiredFields A list of required field names.
// 
RequiredFields []string `json:"required_fields,omitempty"`

}

// Event Event
type Event struct {
// Serial is a serial
Serial int `json:"serial,omitempty"`

// Account The `id` of the related `Account`.
// 
// 
Account string `json:"account,omitempty"`

// Payload is a payload
Payload map[string]interface{} `json:"payload,omitempty"`

// Timestamp is a timestamp
Timestamp time.Time `json:"timestamp,omitempty"`

}

// Status Status Message
type Status struct {
// Severity We are using syslog severity levels: 0 = Emergency,
// 1 = Alert, 2 = Critical, 3 = Error, 4 = Warning,
// 5 = Notice, 6 = Informational, 7 = Debug.
// 
Severity int `json:"severity,omitempty"`

// Tag A machine readable message identifier.
// 
Tag string `json:"tag,omitempty"`

// Message A human readable message, describing the problem
// and may contain hints for resolution.
// 
Message string `json:"message,omitempty"`

// Attrs Optional machine readable key value pairs
// supplementing the message.
// 
// A custom, detailed or localized error messagen can
// be presented to the user, derived from the `tag` and `attrs`.
// 
Attrs map[string]interface{} `json:"attrs,omitempty"`

// Timestamp The time and date when the event occured.
Timestamp time.Time `json:"timestamp,omitempty"`

}

// APIExtension Implementation specific API extensions
type APIExtension struct {
// Name Name of the extension.
// 
Name string `json:"name,omitempty"`

// Publisher Publisher of the extension.
// 
Publisher string `json:"publisher,omitempty"`

// DocumentationURL URL of the documentation homepage of the extension.
// 
DocumentationURL string `json:"documentation_url,omitempty"`

// BaseURL Extension endpoints are available under this base url.
// 
BaseURL string `json:"base_url,omitempty"`

// SpecURL URL of the extensions schema specifications.
// The schema format schould be OpenAPI v3.
// 
SpecURL string `json:"spec_url,omitempty"`

}

// APIFeatures Optional API Features
type APIFeatures struct {
// Pagination The API implementation supports pagination on `list`
// operations.
Pagination *bool `json:"pagination,omitempty"`

}

// APIHealth Health Status Response
type APIHealth struct {
// Status status indicates whether the service status is
// acceptable or not.
Status string `json:"status,omitempty"`

// Version Public version of the service.
// 
Version string `json:"version,omitempty"`

// Releaseid Release version of the api implementation.
// 
Releaseid string `json:"releaseId,omitempty"`

// Notes Array of notes relevant to current state of health.
Notes []string `json:"notes,omitempty"`

// Output Raw error output, in case of "fail" or "warn" states.
Output string `json:"output,omitempty"`

// Serviceid A unique identifier of the service, in the application scope.
Serviceid string `json:"serviceId,omitempty"`

// Description A human-friendly description of the service.
Description string `json:"description,omitempty"`

// Checks The "checks" object MAY have a number of unique keys,
// one for each logical downstream dependency or sub-component.
// 
// Since each sub-component may be backed by several nodes
// with varying health statuses, these keys point to arrays
// of objects. In case of a single-node sub-component
// (or if presence of nodes is not relevant), a single-element
// array SHOULD be used as the value, for consistency.
// 
// Please see
// https://tools.ietf.org/id/draft-inadarei-api-health-check-04.html#the-checks-object
// for details.
Checks map[string]interface{} `json:"checks,omitempty"`

// Links Is an object containing link relations and URIs [RFC3986]
// for external links that MAY contain more information about
// the health of the endpoint.
Links map[string]interface{} `json:"links,omitempty"`

}

// APIImplementation API Implementation
type APIImplementation struct {
// SchemaVersion Version of the implemented IX-API schema.
// 
SchemaVersion string `json:"schema_version,omitempty"`

// ServiceVersion Version of the API service.
// 
ServiceVersion string `json:"service_version,omitempty"`

// SupportedNetworkServiceTypes Array of network service types, supported by the IX.
// 
SupportedNetworkServiceTypes []string `json:"supported_network_service_types,omitempty"`

// SupportedNetworkServiceConfigTypes Array of supported network service config types.
// 
SupportedNetworkServiceConfigTypes []string `json:"supported_network_service_config_types,omitempty"`

// SupportedNetworkFeatureTypes Array of supported network feature types.
// 
SupportedNetworkFeatureTypes []string `json:"supported_network_feature_types,omitempty"`

// SupportedNetworkFeatureConfigTypes Array of supported network feature config types.
// 
SupportedNetworkFeatureConfigTypes []string `json:"supported_network_feature_config_types,omitempty"`

// SupportedOperations Array of implemented operations of the ix-api schema.
// 
SupportedOperations []string `json:"supported_operations,omitempty"`

// SupportedFeatures is a supported_features
SupportedFeatures *APIFeatures `json:"supported_features,omitempty"`

}

// IPAddress IP-Address
type IPAddress struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `IP-Address`.
ID string `json:"id,omitempty"`

// Version The version of the internet protocol.
// 
Version int `json:"version,omitempty"`

// Address IPv4 or IPv6 Address in the following format:
// - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// - IPv6: hexadecimal colon separated notation
// 
Address string `json:"address,omitempty"`

// PrefixLength The CIDR ip prefix length
// 
PrefixLength int `json:"prefix_length,omitempty"`

// FQDN is a fqdn
FQDN *string `json:"fqdn,omitempty"`

// ValidNotBefore is a valid_not_before
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter is a valid_not_after
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// IPAddressPatch IP-Address Update
type IPAddressPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `IP-Address Update`.
ID *string `json:"id,omitempty"`

// Version The version of the internet protocol.
// 
Version *int `json:"version,omitempty"`

// Address IPv4 or IPv6 Address in the following format:
// - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// - IPv6: hexadecimal colon separated notation
// 
Address *string `json:"address,omitempty"`

// PrefixLength The CIDR ip prefix length
// 
PrefixLength *int `json:"prefix_length,omitempty"`

// FQDN is a fqdn
FQDN *string `json:"fqdn,omitempty"`

// ValidNotBefore is a valid_not_before
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter is a valid_not_after
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// IPAddressRequest IP-Address / Prefix allocation Request
type IPAddressRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `IP-Address / Prefix allocation Request`.
ID string `json:"id,omitempty"`

// Version The version of the internet protocol.
// 
Version int `json:"version,omitempty"`

// Address IPv4 or IPv6 Address in the following format:
// - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// - IPv6: hexadecimal colon separated notation
// 
Address string `json:"address,omitempty"`

// PrefixLength The CIDR ip prefix length
// 
PrefixLength int `json:"prefix_length,omitempty"`

// FQDN is a fqdn
FQDN *string `json:"fqdn,omitempty"`

// ValidNotBefore is a valid_not_before
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter is a valid_not_after
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// IPAddressShort IP-Address
type IPAddressShort struct {
// Version The version of the internet protocol.
// 
Version int `json:"version,omitempty"`

// Address The IP address in the following format:
// - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// - IPv6: hexadecimal colon separated notation
// 
Address string `json:"address,omitempty"`

}

// IPAddressUpdate IP-Address Update
type IPAddressUpdate struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `IP-Address Update`.
ID string `json:"id,omitempty"`

// Version The version of the internet protocol.
// 
Version int `json:"version,omitempty"`

// Address IPv4 or IPv6 Address in the following format:
// - IPv4: [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// - IPv6: hexadecimal colon separated notation
// 
Address string `json:"address,omitempty"`

// PrefixLength The CIDR ip prefix length
// 
PrefixLength int `json:"prefix_length,omitempty"`

// FQDN is a fqdn
FQDN *string `json:"fqdn,omitempty"`

// ValidNotBefore is a valid_not_before
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter is a valid_not_after
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// MacAddress MAC-Address
type MacAddress struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `MAC-Address`.
ID string `json:"id,omitempty"`

// Address Unicast MAC address, formatted hexadecimal values with colons.
// 
Address string `json:"address,omitempty"`

// ValidNotBefore When a mac address is assigned to a NSC, and the current
// datetime is before this value, then the MAC address *cannot*
// be used on the peering platform.
// 
// Afterwards, it is supposed to be available. If the value is
// `null` or the property does not exist, the mac address is
// valid from the creation date.
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter When a mac address is assigned to an NSC, and the current datetime
// is before this value, the MAC address *can* be used on the peering platform.
// 
// Afterwards, it is supposed to be unassigned from the NSC and cannot
// any longer be used on the peering platform.
// 
// If the value is null or the property does not exist, the MAC address
// is valid indefinitely. The value may not be in the past.
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// MacAddressRequest MAC-Address Request
type MacAddressRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `MAC-Address Request`.
ID string `json:"id,omitempty"`

// Address Unicast MAC address, formatted hexadecimal values with colons.
// 
Address string `json:"address,omitempty"`

// ValidNotBefore When a mac address is assigned to a NSC, and the current
// datetime is before this value, then the MAC address *cannot*
// be used on the peering platform.
// 
// Afterwards, it is supposed to be available. If the value is
// `null` or the property does not exist, the mac address is
// valid from the creation date.
ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

// ValidNotAfter When a mac address is assigned to an NSC, and the current datetime
// is before this value, the MAC address *can* be used on the peering platform.
// 
// Afterwards, it is supposed to be unassigned from the NSC and cannot
// any longer be used on the peering platform.
// 
// If the value is null or the property does not exist, the MAC address
// is valid indefinitely. The value may not be in the past.
ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

}

// Peer MAC-, IP-Address and ASN of the peer.
type Peer struct {
// ASN The ASN of the peer.
// 
ASN *int `json:"asn,omitempty"`

// IP is a ip
IP *IPAddressShort `json:"ip,omitempty"`

// MacAddress Unicast MAC address, formatted hexadecimal values with colons.
// 
MacAddress string `json:"mac_address,omitempty"`

}

// Conflict A conflict is preventing success
type Conflict struct {
// ResourceType The resource type refers to an ix-api resource.
// 
ResourceType string `json:"resource_type,omitempty"`

// ResourceID The id of the resource which has a conflict with the
// request operation on the current resource.
// 
ResourceID string `json:"resource_id,omitempty"`

// ResourceProperty Indicates the property where the resource is in use.
// 
ResourceProperty *string `json:"resource_property,omitempty"`

// RemoteResourceType The type of the conflicting resource.
// 
RemoteResourceType string `json:"remote_resource_type,omitempty"`

// RemoteResourceID The id of the conflicting resource. This is in most
// cases the id of the current resource.
// 
RemoteResourceID string `json:"remote_resource_id,omitempty"`

}

// ProblemResponse Encodes a problem into an appropriate response body.
type ProblemResponse struct {
// Title A short, human-readable summary of the problem type.
// 
// It SHOULD NOT change from occurrence to
// occurrence of the problem, except for purposes
// of localization (e.g., using proactive content
// negotiation; see [RFC7231], Section 3.4).
// 
Title string `json:"title,omitempty"`

// Status The HTTP status code ([RFC7231], Section 6)
// generated by the origin server for this occurrence
// of the problem.
Status int `json:"status,omitempty"`

// Detail A human-readable explanation specific to this
// occurrence of the problem.
Detail string `json:"detail,omitempty"`

// Instance A URI reference that identifies the specific
// occurrence of the problem.  It may or may not yield
// further information if dereferenced.
Instance string `json:"instance,omitempty"`

}

// Error implements the error interface for ProblemResponse
func (p ProblemResponse) Error() string {
return fmt.Sprintf("%s (%d), %s",
p.Title, p.Status, p.Detail)
}

// ValidationErrorProperty A failed validation
type ValidationErrorProperty struct {
// Name is a name
Name string `json:"name,omitempty"`

// Reason is a reason
Reason string `json:"reason,omitempty"`

}

// AllowMemberJoiningRule A rule for members joining a private vlan
type AllowMemberJoiningRule struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ID The *primary identifier* of the `A rule for members joining a private vlan`.
ID string `json:"id,omitempty"`

// CapacityMin Require an optional minimum capacity to join
// the network service.
CapacityMin *int `json:"capacity_min,omitempty"`

// CapacityMax An optional rate limit which has precedence over
// the capacity set in the network service config.
CapacityMax *int `json:"capacity_max,omitempty"`

// NetworkService The `id` of the related `NetworkService`.
// 
// 
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRule) PolymorphicType() string {
return AllowMemberJoiningRuleType
}

// AllowMemberJoiningRulePatch A vlan member joining rule update
type AllowMemberJoiningRulePatch struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ID The *primary identifier* of the `A vlan member joining rule update`.
ID *string `json:"id,omitempty"`

// CapacityMin Require an optional minimum capacity to join
// the network service.
CapacityMin *int `json:"capacity_min,omitempty"`

// CapacityMax An optional rate limit which has precedence over
// the capacity set in the network service config.
CapacityMax *int `json:"capacity_max,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRulePatch) PolymorphicType() string {
return AllowMemberJoiningRulePatchType
}

// AllowMemberJoiningRuleRequest A new vlan member joining rule
type AllowMemberJoiningRuleRequest struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ID The *primary identifier* of the `A new vlan member joining rule`.
ID string `json:"id,omitempty"`

// CapacityMin Require an optional minimum capacity to join
// the network service.
CapacityMin *int `json:"capacity_min,omitempty"`

// CapacityMax An optional rate limit which has precedence over
// the capacity set in the network service config.
CapacityMax *int `json:"capacity_max,omitempty"`

// NetworkService The `id` of the related `NetworkService`.
// 
// 
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRuleRequest) PolymorphicType() string {
return AllowMemberJoiningRuleRequestType
}

// AllowMemberJoiningRuleUpdate A vlan member joining rule update
type AllowMemberJoiningRuleUpdate struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ID The *primary identifier* of the `A vlan member joining rule update`.
ID string `json:"id,omitempty"`

// CapacityMin Require an optional minimum capacity to join
// the network service.
CapacityMin *int `json:"capacity_min,omitempty"`

// CapacityMax An optional rate limit which has precedence over
// the capacity set in the network service config.
CapacityMax *int `json:"capacity_max,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRuleUpdate) PolymorphicType() string {
return AllowMemberJoiningRuleUpdateType
}

// CloudNetworkService Cloud Network Service
type CloudNetworkService struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// NscRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The `RoleAssignment` is associated through the
// `role_assignments` list property of the network service configuration.
NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service`.
ID string `json:"id,omitempty"`

// Capacity The capacity of the service in Mbps. When null,
// the maximum capacity will be used.
Capacity *int `json:"capacity,omitempty"`

// Diversity Same value as the corresponding `ProductOffering`.
// 
// The service can be delivered over multiple handovers from
// the exchange to the `service_provider`.
// 
// The `diversity` denotes the number of handovers between the
// exchange and the service provider. A value of two signals a
// redundant service.
// 
// Only one network service configuration for each `handover` and
// `cloud_vlan` can be created.
Diversity int `json:"diversity,omitempty"`

// ProviderRef For a cloud network service with the exchange first
// workflow, the `provider_ref` will be a reference
// to a resource of the cloud provider. (E.g. the UUID of
// a virtual circuit.)
// 
// The `provider_ref` is managed by the exchange and its
// meaning may vary between different cloud services.
// 
ProviderRef string `json:"provider_ref,omitempty"`

// CloudKey The cloud key is used to specify to which user or
// existing circuit of a cloud provider this `network-service`
// should be provisioned.
// 
// For example, for a provider like *AWS*, this would be the
// *account number* (Example: `123456789876`), or for a provider
// like Azure, this would be the service key
// (Example: `acl9edcf-f11c-4681-9c7b-6d16b2973997`)
CloudKey string `json:"cloud_key,omitempty"`

// AvailabilityZones The availability zones the service can support.
AvailabilityZones []string `json:"availability_zones,omitempty"`

// NetworkFeatures A list of `id`s of the related `NetworkFeature`.
// 
// 
NetworkFeatures []string `json:"network_features,omitempty"`

// NscProductOfferings An optional list of `ProductOffering` which can be used in the
// network service configs for this service.
NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkService) PolymorphicType() string {
return CloudNetworkServiceType
}

// CloudNetworkServicePatch Cloud Network Service Update
type CloudNetworkServicePatch struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering *string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Update`.
ID *string `json:"id,omitempty"`

// Capacity The capacity of the service in Mbps. When null,
// the maximum capacity will be used.
Capacity *int `json:"capacity,omitempty"`

// CloudKey The cloud key is used to specify to which user or
// existing circuit of a cloud provider this `network-service`
// should be provisioned.
// 
// For example, for a provider like *AWS*, this would be the
// *account number* (Example: `123456789876`), or for a provider
// like Azure, this would be the service key
// (Example: `acl9edcf-f11c-4681-9c7b-6d16b2973997`)
// 
// **Please note: *Any update to this field may be rejected if the
// service has successfully been provisioned*.**
CloudKey *string `json:"cloud_key,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServicePatch) PolymorphicType() string {
return CloudNetworkServicePatchType
}

// CloudNetworkServiceRequest Cloud Network Service Request
type CloudNetworkServiceRequest struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Request`.
ID string `json:"id,omitempty"`

// Capacity The capacity of the service in Mbps. When null,
// the maximum capacity will be used.
Capacity *int `json:"capacity,omitempty"`

// CloudKey The cloud key is used to specify to which user or
// existing circuit of a cloud provider this `network-service`
// should be provisioned.
// 
// For example, for a provider like *AWS*, this would be the
// *account number* (Example: `123456789876`), or for a provider
// like Azure, this would be the service key
// (Example: `acl9edcf-f11c-4681-9c7b-6d16b2973997`)
CloudKey string `json:"cloud_key,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceRequest) PolymorphicType() string {
return CloudNetworkServiceRequestType
}

// CloudNetworkServiceUpdate Cloud Network Service Update
type CloudNetworkServiceUpdate struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Cloud Network Service Update`.
ID string `json:"id,omitempty"`

// Capacity The capacity of the service in Mbps. When null,
// the maximum capacity will be used.
Capacity *int `json:"capacity,omitempty"`

// CloudKey The cloud key is used to specify to which user or
// existing circuit of a cloud provider this `network-service`
// should be provisioned.
// 
// For example, for a provider like *AWS*, this would be the
// *account number* (Example: `123456789876`), or for a provider
// like Azure, this would be the service key
// (Example: `acl9edcf-f11c-4681-9c7b-6d16b2973997`)
// 
// **Please note: *Any update to this field may be rejected if the
// service has successfully been provisioned*.**
CloudKey string `json:"cloud_key,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceUpdate) PolymorphicType() string {
return CloudNetworkServiceUpdateType
}

// DenyMemberJoiningRule A rule for members joining a private vlan
type DenyMemberJoiningRule struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// NetworkService The `id` of the related `NetworkService`.
// 
// 
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRule) PolymorphicType() string {
return DenyMemberJoiningRuleType
}

// DenyMemberJoiningRulePatch A vlan member joining rule update
type DenyMemberJoiningRulePatch struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRulePatch) PolymorphicType() string {
return DenyMemberJoiningRulePatchType
}

// DenyMemberJoiningRuleRequest A new vlan member joining rule
type DenyMemberJoiningRuleRequest struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// NetworkService The `id` of the related `NetworkService`.
// 
// 
NetworkService string `json:"network_service,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRuleRequest) PolymorphicType() string {
return DenyMemberJoiningRuleRequestType
}

// DenyMemberJoiningRuleUpdate A vlan member joining rule update
type DenyMemberJoiningRuleUpdate struct {
// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account to which access to the
// network service should be granted or denied.
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRuleUpdate) PolymorphicType() string {
return DenyMemberJoiningRuleUpdateType
}

// ExchangeLanNetworkService Exchange Lan Network Service
type ExchangeLanNetworkService struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// NscRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The `RoleAssignment` is associated through the
// `role_assignments` list property of the network service configuration.
NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// ID The *primary identifier* of the `Exchange Lan Network Service`.
ID string `json:"id,omitempty"`

// Name Exchange-dependent service name, will be shown on the invoice.
Name string `json:"name,omitempty"`

// MetroAreaNetwork Id of the `MetroAreaNetwork` where
// the exchange lan network service is directly provided.
// 
// Same as `service_metro_area_network` on the related
// `ProductOffering`.
// 
MetroAreaNetwork string `json:"metro_area_network,omitempty"`

// PeeringdbIxid PeeringDB ixid
PeeringdbIxid *int `json:"peeringdb_ixid,omitempty"`

// IxfdbIxid id of ixfdb
IxfdbIxid *int `json:"ixfdb_ixid,omitempty"`

// NetworkFeatures A list of `id`s of the related `NetworkFeature`.
// 
// 
NetworkFeatures []string `json:"network_features,omitempty"`

// SubnetV4 IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// CIDR notation.
// 
SubnetV4 string `json:"subnet_v4,omitempty"`

// SubnetV6 IPv6 subnet in hexadecimal colon separated CIDR notation.
// 
SubnetV6 string `json:"subnet_v6,omitempty"`

// ProductOffering *deprecation notice*
ProductOffering *string `json:"product_offering,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkService) PolymorphicType() string {
return ExchangeLanNetworkServiceType
}

// IXPSpecificFeatureFlag IXP-Specific Feature Flag
type IXPSpecificFeatureFlag struct {
// Name The name of the feature flag.
// 
Name string `json:"name,omitempty"`

// Description The description of the feature flag.
// 
Description string `json:"description,omitempty"`

// Mandatory This feature will always be enabled, even if not provided in
// the corresponding config's list of `flags`.
// 
Mandatory bool `json:"mandatory,omitempty"`

}

// MP2MPNetworkService MP2MP Network Service
type MP2MPNetworkService struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// NscRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The `RoleAssignment` is associated through the
// `role_assignments` list property of the network service configuration.
NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// NscProductOfferings An optional list of `ProductOffering` which can be used in the
// network service configs for this service.
NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service`.
ID string `json:"id,omitempty"`

// Public A public mp2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `display_name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

// DisplayName Name of the multi-point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of, for example:
// "Financial Clearance LAN".
// 
DisplayName *string `json:"display_name,omitempty"`

// SubnetV4 IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// CIDR notation.
// 
SubnetV4 *string `json:"subnet_v4,omitempty"`

// SubnetV6 IPv6 subnet in hexadecimal colon separated CIDR notation.
// 
SubnetV6 *string `json:"subnet_v6,omitempty"`

// MemberJoiningRules A list of `id`s of the related `MemberJoiningRule`.
// 
// 
MemberJoiningRules []string `json:"member_joining_rules,omitempty"`

// NetworkFeatures A list of `id`s of the related `NetworkFeature`.
// 
// 
NetworkFeatures []string `json:"network_features,omitempty"`

// MacAclProtection When enabled, only MAC addresses in the referenced in the network
// service config's `macs` property are allowed to send and receive
// traffic on this network service.
MacAclProtection *bool `json:"mac_acl_protection,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkService) PolymorphicType() string {
return MP2MPNetworkServiceType
}

// MP2MPNetworkServicePatch MP2MP Network Service Update
type MP2MPNetworkServicePatch struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering *string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Update`.
ID *string `json:"id,omitempty"`

// Public A public mp2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `display_name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

// DisplayName Name of the multi-point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of, for example:
// "Financial Clearance LAN".
// 
DisplayName *string `json:"display_name,omitempty"`

// SubnetV4 IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// CIDR notation.
// 
SubnetV4 *string `json:"subnet_v4,omitempty"`

// SubnetV6 IPv6 subnet in hexadecimal colon separated CIDR notation.
// 
SubnetV6 *string `json:"subnet_v6,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServicePatch) PolymorphicType() string {
return MP2MPNetworkServicePatchType
}

// MP2MPNetworkServiceRequest MP2MP Network Service Request
type MP2MPNetworkServiceRequest struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Request`.
ID string `json:"id,omitempty"`

// Public A public mp2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `display_name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

// DisplayName Name of the multi-point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of, for example:
// "Financial Clearance LAN".
// 
DisplayName *string `json:"display_name,omitempty"`

// SubnetV4 IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// CIDR notation.
// 
SubnetV4 *string `json:"subnet_v4,omitempty"`

// SubnetV6 IPv6 subnet in hexadecimal colon separated CIDR notation.
// 
SubnetV6 *string `json:"subnet_v6,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceRequest) PolymorphicType() string {
return MP2MPNetworkServiceRequestType
}

// MP2MPNetworkServiceUpdate MP2MP Network Service Update
type MP2MPNetworkServiceUpdate struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `MP2MP Network Service Update`.
ID string `json:"id,omitempty"`

// Public A public mp2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `display_name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

// DisplayName Name of the multi-point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of, for example:
// "Financial Clearance LAN".
// 
DisplayName *string `json:"display_name,omitempty"`

// SubnetV4 IPv4 subnet in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// CIDR notation.
// 
SubnetV4 *string `json:"subnet_v4,omitempty"`

// SubnetV6 IPv6 subnet in hexadecimal colon separated CIDR notation.
// 
SubnetV6 *string `json:"subnet_v6,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceUpdate) PolymorphicType() string {
return MP2MPNetworkServiceUpdateType
}

// MemberJoiningRule Polymorphic Member Joining Rule
type MemberJoiningRule struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// MemberJoiningRulePatch Polymorphic Member Joining Rule Update
type MemberJoiningRulePatch struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// MemberJoiningRuleRequest Polymorphic Member Joining Rule Request
type MemberJoiningRuleRequest struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// MemberJoiningRuleUpdate Polymorphic Member Joining Rule Update
type MemberJoiningRuleUpdate struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkFeature Polymorphic Network Feature
type NetworkFeature struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkService Polymorphic Network Services
type NetworkService struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceChangeRequest NetworkServiceChangeRequest
type NetworkServiceChangeRequest struct {
// ProductOffering Migrate to a diffrent product offering. Please note, that
// the offering only may differ in bandwidth.
ProductOffering string `json:"product_offering,omitempty"`

// Capacity The desired capacity of the service in Mbps.
// 
// Must be within the range of `bandwidth_min` and
// `bandwidth_max` of the `ProductOffering`.
// 
// When `null` the maximum capacity wil be used.
Capacity *int `json:"capacity,omitempty"`

}

// NetworkServiceDeleteResponse Polymorphic Network Service Request
type NetworkServiceDeleteResponse struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServicePatch Polymorphic Network Service Patch
type NetworkServicePatch struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceRequest Polymorphic Network Service Request
type NetworkServiceRequest struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// NetworkServiceUpdate Polymorphic Network Service Update
type NetworkServiceUpdate struct {
// Type is a type
Type string `json:"type,omitempty"`

}

// P2MPNetworkService P2MP Network Service
type P2MPNetworkService struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// NscRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The `RoleAssignment` is associated through the
// `role_assignments` list property of the network service configuration.
NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// NscProductOfferings An optional list of `ProductOffering` which can be used in the
// network service configs for this service.
NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// Public A public p2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made
// available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

// NetworkFeatures A list of `id`s of the related `NetworkFeature`.
// 
// 
NetworkFeatures []string `json:"network_features,omitempty"`

// MemberJoiningRules A list of `id`s of the related `MemberJoiningRule`.
// 
// 
MemberJoiningRules []string `json:"member_joining_rules,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkService) PolymorphicType() string {
return P2MPNetworkServiceType
}

// P2MPNetworkServicePatch P2MP Network Service Update
type P2MPNetworkServicePatch struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering *string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Update`.
ID *string `json:"id,omitempty"`

// DisplayName Name of the point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// Public A public p2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made
// available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServicePatch) PolymorphicType() string {
return P2MPNetworkServicePatchType
}

// P2MPNetworkServiceRequest P2MP Network Service Request
type P2MPNetworkServiceRequest struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Request`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// Public A public p2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made
// available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceRequest) PolymorphicType() string {
return P2MPNetworkServiceRequestType
}

// P2MPNetworkServiceUpdate P2MP Network Service Update
type P2MPNetworkServiceUpdate struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2MP Network Service Update`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to multi-point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// Public A public p2mp network service can be joined
// by everyone on the exchange unless denied by
// a member-joining-rule.
// 
// Public network services are visible to other
// members of the IXP, however only `name`, `type`,
// `product_offering`, `consuming_account` and
// `managing_account` are made
// available.
// 
// Other required fields are redacted.
Public *bool `json:"public,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceUpdate) PolymorphicType() string {
return P2MPNetworkServiceUpdateType
}

// P2PNetworkService P2P Network Service
type P2PNetworkService struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// NscRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The `RoleAssignment` is associated through the
// `role_assignments` list property of the network service configuration.
NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// NscProductOfferings An optional list of `ProductOffering` which can be used in the
// network service configs for this service.
NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

// DecommissionAt The service will be decommissioned on this date.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// *(Sensitive Property)*
DecommissionAt *Date `json:"decommission_at,omitempty"`

// ChargedUntil The service continues incurring charges until this date.
// Typically `≥ decommission_at`.
// 
// This field is only used when
// the state is `DECOMMISSION_REQUESTED` or
// `DECOMMISSIONED`.
// 
// *(Sensitive Property)*
ChargedUntil *Date `json:"charged_until,omitempty"`

// CurrentBillingStartDate Your obligation to pay for the service will start on this date.
// 
// However, this date may change after an upgrade and not reflect
// the inital start date of the service.
// 
// *(Sensitive Property)*
CurrentBillingStartDate *Date `json:"current_billing_start_date,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2P Network Service`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// JoiningMemberAccount The account of the B-side member joining the virtual circuit.
// 
JoiningMemberAccount string `json:"joining_member_account,omitempty"`

// AvailabilityZones The availability zones for the service.
AvailabilityZones []string `json:"availability_zones,omitempty"`

// Capacity The capacity of the service in Mbps. When null,
// the maximum capacity will be used.
Capacity *int `json:"capacity,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkService) PolymorphicType() string {
return P2PNetworkServiceType
}

// P2PNetworkServicePatch P2P Network Service Update
type P2PNetworkServicePatch struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering *string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Update`.
ID *string `json:"id,omitempty"`

// DisplayName Name of the point to point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// JoiningMemberAccount The account of the B-side member joining the virtual circuit.
// 
JoiningMemberAccount *string `json:"joining_member_account,omitempty"`

// AvailabilityZones The availability zones for the service.
AvailabilityZones []string `json:"availability_zones,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServicePatch) PolymorphicType() string {
return P2PNetworkServicePatchType
}

// P2PNetworkServiceRequest P2P Network Service Request
type P2PNetworkServiceRequest struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Request`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// JoiningMemberAccount The account of the B-side member joining the virtual circuit.
// 
JoiningMemberAccount string `json:"joining_member_account,omitempty"`

// AvailabilityZones The availability zones for the service.
AvailabilityZones []string `json:"availability_zones,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceRequest) PolymorphicType() string {
return P2PNetworkServiceRequestType
}

// P2PNetworkServiceUpdate P2P Network Service Update
type P2PNetworkServiceUpdate struct {
// ProductOffering The `id` of the related `ProductOffering`.
// 
// 
ProductOffering string `json:"product_offering,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `P2P Network Service Update`.
ID string `json:"id,omitempty"`

// DisplayName Name of the point to point virtual circuit.
// 
// It is visible to all parties allowed to connect
// to this virtual circuit.
// 
// It is intended for humans to make sense of.
// 
DisplayName *string `json:"display_name,omitempty"`

// JoiningMemberAccount The account of the B-side member joining the virtual circuit.
// 
JoiningMemberAccount string `json:"joining_member_account,omitempty"`

// AvailabilityZones The availability zones for the service.
AvailabilityZones []string `json:"availability_zones,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceUpdate) PolymorphicType() string {
return P2PNetworkServiceUpdateType
}

// RouteServerNetworkFeature Route Server Network Feature
type RouteServerNetworkFeature struct {
// ID The *primary identifier* of the `Route Server Network Feature`.
ID string `json:"id,omitempty"`

// Name is a name
Name string `json:"name,omitempty"`

// Required is a required
Required bool `json:"required,omitempty"`

// NetworkService The `id` of the related `NetworkService`.
// 
// 
NetworkService string `json:"network_service,omitempty"`

// NfcRequiredContactRoles The configuration will require at least one of each of the
// specified roles assigned to contacts.
// 
// The role assignments is associated with the network feature
// config through the `role_assignments` list property.
NfcRequiredContactRoles []string `json:"nfc_required_contact_roles,omitempty"`

// Flags A list of IXP specific feature flags. This can be used
// to see if e.g. RPKI hard filtering is available.
Flags []*IXPSpecificFeatureFlag `json:"flags,omitempty"`

// ASN is a asn
ASN int `json:"asn,omitempty"`

// FQDN The FQDN of the route server.
// 
FQDN string `json:"fqdn,omitempty"`

// LookingGlassURL The url of the looking glass.
// 
LookingGlassURL *string `json:"looking_glass_url,omitempty"`

// AddressFamilies When creating a route server feature config, remember
// to specify which address family or families to use:
// 
// If the route server network feature only supports `af_inet`,
// then the `as_set_v4` in the network feature config is required.
// 
// If only `af_inet6` is supported, then the `as_set_v6` is required.
// 
// If both `af_inet` and `af_inet6` are supported, either
// `as_set_v4` or `as_set_v6` is required, but both can be provided
// in the network service config.
// 
AddressFamilies []string `json:"address_families,omitempty"`

// SessionMode When creating a route server feature config, remember
// to specify the same session_mode as the route server.
// 
SessionMode string `json:"session_mode,omitempty"`

// AvailableBGPSessionTypes The route server provides the following session modes.
// 
AvailableBGPSessionTypes []string `json:"available_bgp_session_types,omitempty"`

// IPV4 IPv4 address in [dot-decimal notation](https://en.wikipedia.org/wiki/Dot-decimal_notation)
// notation.
// 
// This field is only set if the `address_families` include `af_inet`.
// 
IPV4 *string `json:"ip_v4,omitempty"`

// IPV6 IPv6 address in hexadecimal colon separated notation.
// 
// This field is only set if the `address_families` include `af_inet6`.
// 
IPV6 *string `json:"ip_v6,omitempty"`

}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeature) PolymorphicType() string {
return RouteServerNetworkFeatureType
}

// RoutingFunction Routing Function
type RoutingFunction struct {
// State The state of the object.
// *(Sensitive Property)*
State string `json:"state,omitempty"`

// Status Status information about the object.
// *(Sensitive Property)*
Status []*Status `json:"status,omitempty"`

// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Routing Function`.
ID string `json:"id,omitempty"`

// ProductOffering The product offering to be used for the
// routing function.
ProductOffering string `json:"product_offering,omitempty"`

// ASN Any routing function instance needs to be
// assigned a 2-byte or 4-byte ASN of the
// customer's choice. There is no restriction on
// private or public ASNs.
ASN int `json:"asn,omitempty"`

// Capacity The desired upper bound of the capacity for
// the routing function.
Capacity *int `json:"capacity,omitempty"`

}

// RoutingFunctionPatch Routing Function Patch
type RoutingFunctionPatch struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount *string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount *string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount *string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Routing Function Patch`.
ID *string `json:"id,omitempty"`

// ProductOffering The product offering to be used for the
// routing function.
ProductOffering *string `json:"product_offering,omitempty"`

// ASN Any routing function instance needs to be
// assigned a 2-byte or 4-byte ASN of the
// customer's choice. There is no restriction on
// private or public ASNs.
ASN *int `json:"asn,omitempty"`

// Capacity The desired upper bound of the capacity for
// the routing function.
Capacity *int `json:"capacity,omitempty"`

}

// RoutingFunctionRequest Routing Function Request
type RoutingFunctionRequest struct {
// ManagingAccount The `id` of the account responsible for managing the service via
// the API. A manager can read and update the state of entities.
// *(Sensitive Property)*
// 
ManagingAccount string `json:"managing_account,omitempty"`

// ConsumingAccount The `id` of the account consuming a service.
// 
// Used to be `owning_customer`.
// *(Sensitive Property)*
// 
ConsumingAccount string `json:"consuming_account,omitempty"`

// ExternalRef Reference field, free to use for the API user.
// *(Sensitive Property)*
// 
ExternalRef *string `json:"external_ref,omitempty"`

// PurchaseOrder Purchase Order ID which will be displayed on the invoice.
// *(Sensitive Property)*
// 
PurchaseOrder *string `json:"purchase_order,omitempty"`

// ContractRef A reference to a contract. If no specific contract is used,
// a default MAY be chosen by the implementer.
// *(Sensitive Property)*
// 
ContractRef *string `json:"contract_ref,omitempty"`

// BillingAccount An account requires billing_information to be used as a `billing_account`.
// *(Sensitive Property)*
BillingAccount string `json:"billing_account,omitempty"`

// ID The *primary identifier* of the `Routing Function Request`.
ID string `json:"id,omitempty"`

// ProductOffering The product offering to be used for the
// routing function.
ProductOffering string `json:"product_offering,omitempty"`

// ASN Any routing function instance needs to be
// assigned a 2-byte or 4-byte ASN of the
// customer's choice. There is no restriction on
// private or public ASNs.
ASN int `json:"asn,omitempty"`

// Capacity The desired upper bound of the capacity for
// the routing function.
Capacity *int `json:"capacity,omitempty"`

}

// NetworkServiceConfigAggregate Statistics for NetworkServiceConfig
type NetworkServiceConfigAggregate struct {
// Aggregates Aggregated statistics for a connection or service configuration.
// 
// For the **property name** the string representation of the
// aggregate interval in ISO8601 period notation is recommended.
// 
// For example: `PT5M`, `P1D`, `P30D`,`P1Y`.
// 
// If a window is defined via the `gtart` and `end` query parameter,
// the **property name** will be `custom`.
// 
// The available intervals can differ by implementation.
// 
Aggregates map[string]interface{} `json:"aggregates,omitempty"`

}

// NetworkServiceConfigAggregateStatistics AggregateStatistics for NetworkServiceConfig
type NetworkServiceConfigAggregateStatistics struct {
// Title Title of the aggregated statistics.
// 
Title string `json:"title,omitempty"`

// Start Start of the traffic aggregation.
Start time.Time `json:"start,omitempty"`

// End End of the traffic aggregation.
End time.Time `json:"end,omitempty"`

// Accuracy The accuracy is the ratio of *total aggregated samples* to
// *expected samples*.
// 
// The expected number of samples is the size of the window
// of the aggregate, divided by the aggregation resolution.
// 
// For example: A window of `24 h` with an aggregation resolution
// of `5 m` would yield `288` samples.
// 
// If only `275` samples are available for aggregation, the
// accuracy would be `0.95`.
// 
Accuracy float64 `json:"accuracy,omitempty"`

// CreatedAt Timestamp when the statistics were created.
CreatedAt time.Time `json:"created_at,omitempty"`

// NextUpdateAt Next update of the statistical data.
// This may not correspond to the aggregate interval.
NextUpdateAt time.Time `json:"next_update_at,omitempty"`

// AveragePpsIn Average number of inbound **packets per second**.
// 
AveragePpsIn int `json:"average_pps_in,omitempty"`

// AveragePpsOut Average number outbound **packets per second**.
// 
AveragePpsOut int `json:"average_pps_out,omitempty"`

// AverageOpsIn Average inbound **octets per second**.
// 
AverageOpsIn int `json:"average_ops_in,omitempty"`

// AverageOpsOut Average outbound **octets per second**.
// 
AverageOpsOut int `json:"average_ops_out,omitempty"`

// AverageEpsIn Average **errors per second** inbound.
// 
AverageEpsIn *float64 `json:"average_eps_in,omitempty"`

// AverageEpsOut Averages **errors per second** outbound.
// 
AverageEpsOut *float64 `json:"average_eps_out,omitempty"`

// AverageDpsIn Average **discards per second** inbound.
// 
AverageDpsIn *float64 `json:"average_dps_in,omitempty"`

// AverageDpsOut Averages **discards per second** outbound.
// 
AverageDpsOut *float64 `json:"average_dps_out,omitempty"`

// Percentile95PpsIn 95th percentile of the inbound **octets per second**.
// 
Percentile95PpsIn *int `json:"percentile95_pps_in,omitempty"`

// Percentile95PpsOut 95th percentile of outbound **packets per second**.
// 
Percentile95PpsOut *int `json:"percentile95_pps_out,omitempty"`

// Percentile95OpsOut 95th percentile of outbound **octets per second**.
// 
Percentile95OpsOut *int `json:"percentile95_ops_out,omitempty"`

// MaximumPpsIn Peak inbound **packets per second** during the interval.
// 
MaximumPpsIn *int `json:"maximum_pps_in,omitempty"`

// MaximumPpsOut Peak outbound **packets per second** during the interval.
// 
MaximumPpsOut *int `json:"maximum_pps_out,omitempty"`

// MaximumOpsIn Peak inbound **octets per second** during the interval.
// 
MaximumOpsIn *int `json:"maximum_ops_in,omitempty"`

// MaximumOpsOut Peak outbound **octets per second** during the interval.
// 
MaximumOpsOut *int `json:"maximum_ops_out,omitempty"`

// MaximumInAt Timestamp when the inbound peak occured.
MaximumInAt *time.Time `json:"maximum_in_at,omitempty"`

// MaximumOutAt Timestamp when the outbound peak occured.
MaximumOutAt *time.Time `json:"maximum_out_at,omitempty"`

// NscAvailableCapacity The capacity left on the `NetworkServiceConfig` in
// **megabits per second** (Mbps).
// 
NscAvailableCapacity *int `json:"nsc_available_capacity,omitempty"`

// NscAvailableCapacityChangePerc The percentage change of the available capacity since
// the last update.
// 
NscAvailableCapacityChangePerc *float64 `json:"nsc_available_capacity_change_perc,omitempty"`

}

// PeerAggregate PeerStatistics
type PeerAggregate struct {
// Aggregates Aggregated statistics for a connection or service configuration.
// 
// For the **property name** the string representation of the
// aggregate interval in ISO8601 period notation is recommended.
// 
// For example: `PT5M`, `P1D`, `P30D`,`P1Y`.
// 
// If a window is defined via the `gtart` and `end` query parameter,
// the **property name** will be `custom`.
// 
// The available intervals can differ by implementation.
// 
Aggregates map[string]interface{} `json:"aggregates,omitempty"`

// Peer is a peer
Peer *Peer `json:"peer,omitempty"`

}

// PeerRTT Peer RTT Statistics
type PeerRTT struct {
// TimeMs The total duration of the measurement in
// milliseconds.
// 
TimeMs int `json:"time_ms,omitempty"`

// Tx The number of probe packets *transmitted*
// within the duration of the measurement.
// 
Tx int `json:"tx,omitempty"`

// Rx The number of probe packets *received*
// within the duration of the measurement.
// 
Rx int `json:"rx,omitempty"`

// Loss Ratio of *transmitted packets* to *received packets*:
// `loss = 1.0 - (rx / tx)`.
// 
Loss float64 `json:"loss,omitempty"`

// RttMinMs The minimum RTT in milliseconds.
// 
RttMinMs float64 `json:"rtt_min_ms,omitempty"`

// RttAvgMs The average RTT in milliseconds.
// 
RttAvgMs float64 `json:"rtt_avg_ms,omitempty"`

// RttMaxMs The maximum RTT in milliseconds.
// 
RttMaxMs float64 `json:"rtt_max_ms,omitempty"`

// RttMdevMs The median RTT in milliseconds.
// Standard deviation in milliseconds.
// 
RttMdevMs float64 `json:"rtt_mdev_ms,omitempty"`

// Neighbor The name of the peer.
// 
Neighbor string `json:"neighbor,omitempty"`

// ASN The Autonomous System Number (ASN) of the peer.
// 
ASN *int `json:"asn,omitempty"`

// IP The IP address of the peer.
// For IPv6 addresses the canonical form is used.
// 
IP string `json:"ip,omitempty"`

// Timestamp The date and time when the RTT statistic was measured.
Timestamp time.Time `json:"timestamp,omitempty"`

// Serial The `serial` is an incrementing counter. You can use it
// to poll for changes.
// 
Serial int `json:"serial,omitempty"`

}

// PortStatistics Port Statistics
type PortStatistics struct {
// Aggregates Aggregated statistics for a connection or service configuration.
// 
// For the **property name** the string representation of the
// aggregate interval in ISO8601 period notation is recommended.
// 
// For example: `PT5M`, `P1D`, `P30D`,`P1Y`.
// 
// If a window is defined via the `start` and `end` query parameter,
// the **property name** will be `custom`.
// 
// The available intervals can differ by implementation.
// 
Aggregates map[string]interface{} `json:"aggregates,omitempty"`

// LightLevelsTx A list of light levels in **dBm** for each channel.
// 
LightLevelsTx []float64 `json:"light_levels_tx,omitempty"`

// LightLevelsRx A list of light levels in **dBm** for each channel.
// 
LightLevelsRx []float64 `json:"light_levels_rx,omitempty"`

}

