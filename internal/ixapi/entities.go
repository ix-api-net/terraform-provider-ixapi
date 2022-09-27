package ixapi

import (
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

// CloudNetworkProductOffering Cloud Network Product Offering
type CloudNetworkProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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

	// ServiceProviderRegion The service provider offers the network service for a
	// specific region.
	//
	ServiceProviderRegion string `json:"service_provider_region,omitempty"`

	// ServiceProviderPop The datacenter description of the partner NNI to the service provider.
	//
	ServiceProviderPop string `json:"service_provider_pop,omitempty"`

	// ServiceProviderWorkflow When the workflow is `provider_first` the subscriber creates
	// a circuit with the cloud provider and provides a `cloud_key` for filtering
	// the product-offerings.
	//
	// If the workflow is `exchange_first` the IX will create
	// the cloud circuit on the provider side.
	//
	ServiceProviderWorkflow string `json:"service_provider_workflow,omitempty"`

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
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkProductOffering) PolymorphicType() string {
	return CloudNetworkProductOfferingType
}

// ConnectionProductOffering Connection Product Offering
type ConnectionProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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

// Device Device
type Device struct {
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

	// ID is a id
	ID string `json:"id,omitempty"`
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

	// Device is a device
	Device string `json:"device,omitempty"`

	// ConnectedDevice is a connected_device
	ConnectedDevice string `json:"connected_device,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`
}

// ExchangeLanNetworkProductOffering Exchange Lan Network Product Offering
type ExchangeLanNetworkProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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

	// ExchangeLanNetworkService The id of the exchange lan network service.
	ExchangeLanNetworkService string `json:"exchange_lan_network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkProductOffering) PolymorphicType() string {
	return ExchangeLanNetworkProductOfferingType
}

// Facility Facility
type Facility struct {
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

	// ID is a id
	ID string `json:"id,omitempty"`
}

// MP2MPNetworkProductOffering MP2MP Network Product Offering
type MP2MPNetworkProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkProductOffering) PolymorphicType() string {
	return MP2MPNetworkProductOfferingType
}

// MetroArea MetroArea
type MetroArea struct {
	// ID is a id
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

	// ID is a id
	ID string `json:"id,omitempty"`
}

// P2MPNetworkProductOffering P2MP Network Product Offering
type P2MPNetworkProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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
}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkProductOffering) PolymorphicType() string {
	return P2MPNetworkProductOfferingType
}

// P2PNetworkProductOffering P2P Network Product Offering
type P2PNetworkProductOffering struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
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

	// ContractTerms The contract terms informally describe the contract period and
	// renewals.
	//
	ContractTerms *string `json:"contract_terms,omitempty"`

	// NoticePeriod The notice period informally states constraints
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
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkProductOffering) PolymorphicType() string {
	return P2PNetworkProductOfferingType
}

// PointOfPresence Point Of Presence
type PointOfPresence struct {
	// Name is a name
	Name string `json:"name,omitempty"`

	// Facility The pop is located in this `Facility`.
	Facility string `json:"facility,omitempty"`

	// MetroAreaNetwork is a metro_area_network
	MetroAreaNetwork string `json:"metro_area_network,omitempty"`

	// Devices is a devices
	Devices []string `json:"devices,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`
}

// ProductOffering Polymorphic Product Offering
type ProductOffering interface {
	Polymorphic
}

// PolymorphicProductOffering is a polymorphic base
type PolymorphicProductOffering struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (p PolymorphicProductOffering) PolymorphicType() string {
	return p.Type
}

// ConnectionProductOfferingType is a polymorphic type value for ConnectionProductOffering
const ConnectionProductOfferingType = "connection"

// ExchangeLanNetworkProductOfferingType is a polymorphic type value for ExchangeLanNetworkProductOffering
const ExchangeLanNetworkProductOfferingType = "exchange_lan"

// P2PNetworkProductOfferingType is a polymorphic type value for P2PNetworkProductOffering
const P2PNetworkProductOfferingType = "p2p_vc"

// MP2MPNetworkProductOfferingType is a polymorphic type value for MP2MPNetworkProductOffering
const MP2MPNetworkProductOfferingType = "mp2mp_vc"

// P2MPNetworkProductOfferingType is a polymorphic type value for P2MPNetworkProductOffering
const P2MPNetworkProductOfferingType = "p2mp_vc"

// CloudNetworkProductOfferingType is a polymorphic type value for CloudNetworkProductOffering
const CloudNetworkProductOfferingType = "cloud_vc"

// CloudNetworkServiceConfig Cloud Network Service Config
type CloudNetworkServiceConfig struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService The id of the configured network service.
	NetworkService string `json:"network_service,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

	// Handover The handover enumerates the connection and is
	// required for checking diversity constraints.
	//
	// It must be within `1 <= x <= network_service.diversity`.
	//
	Handover int `json:"handover,omitempty"`

	// CloudVLAN If the `provider_vlans` property of the `ProductOffering` is
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
	CloudVLAN int `json:"cloud_vlan,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfig) PolymorphicType() string {
	return CloudNetworkServiceConfigType
}

// CloudNetworkServiceConfigPatch Cloud Network Service Config Update
type CloudNetworkServiceConfigPatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection *string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// Handover The handover enumerates the connection and is
	// required for checking diversity constraints.
	//
	// It must be within `1 <= x <= network_service.diversity`.
	//
	Handover *int `json:"handover,omitempty"`

	// CloudVLAN If the `provider_vlans` property of the `ProductOffering` is
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
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigPatch) PolymorphicType() string {
	return CloudNetworkServiceConfigPatchType
}

// CloudNetworkServiceConfigRequest Cloud Network Service Config Request
type CloudNetworkServiceConfigRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

	// Handover The handover enumerates the connection and is
	// required for checking diversity constraints.
	//
	// It must be within `1 <= x <= network_service.diversity`.
	//
	Handover int `json:"handover,omitempty"`

	// CloudVLAN If the `provider_vlans` property of the `ProductOffering` is
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
	CloudVLAN int `json:"cloud_vlan,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigRequest) PolymorphicType() string {
	return CloudNetworkServiceConfigRequestType
}

// CloudNetworkServiceConfigUpdate Cloud Network Service Config Update
type CloudNetworkServiceConfigUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// Handover The handover enumerates the connection and is
	// required for checking diversity constraints.
	//
	// It must be within `1 <= x <= network_service.diversity`.
	//
	Handover int `json:"handover,omitempty"`

	// CloudVLAN If the `provider_vlans` property of the `ProductOffering` is
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
	CloudVLAN int `json:"cloud_vlan,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceConfigUpdate) PolymorphicType() string {
	return CloudNetworkServiceConfigUpdateType
}

// Connection Connection
type Connection struct {
	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

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

	// ID is a id
	ID string `json:"id,omitempty"`

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

	// CapacityAllocated Sum of the bandwidth of all network services using
	// the connection in Mbit/s.
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
}

// ConnectionPatch Connection Update
type ConnectionPatch struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService The id of the configured network service.
	NetworkService string `json:"network_service,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
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
	Capacity *int `json:"capacity,omitempty"`

	// ASNs is a asns
	ASNs []int `json:"asns,omitempty"`

	// Macs A list of mac-address IDs.
	Macs []string `json:"macs,omitempty"`

	// IPs A list of ip-address IDs.
	//
	// Allocation of IP Addresses might be deferred depending on
	// the IXP implementation. No assumption should be made.
	IPs []string `json:"ips,omitempty"`

	// Listed The customer wants to be featured on the member list
	Listed bool `json:"listed,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection *string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ASNs is a asns
	ASNs []int `json:"asns,omitempty"`

	// Macs A list of mac-address IDs.
	Macs []string `json:"macs,omitempty"`

	// IPs A list of ip-address IDs.
	//
	// Allocation of IP Addresses might be deferred depending on
	// the IXP implementation. No assumption should be made.
	IPs []string `json:"ips,omitempty"`

	// Listed The customer wants to be featured on the member list
	Listed *bool `json:"listed,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (e ExchangeLanNetworkServiceConfigPatch) PolymorphicType() string {
	return ExchangeLanNetworkServiceConfigPatchType
}

// ExchangeLanNetworkServiceConfigRequest Exchange Lan Network Service Config Request
type ExchangeLanNetworkServiceConfigRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
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
	Capacity *int `json:"capacity,omitempty"`

	// ASNs is a asns
	ASNs []int `json:"asns,omitempty"`

	// Macs A list of mac-address IDs.
	Macs []string `json:"macs,omitempty"`

	// IPs A list of ip-address IDs.
	//
	// Allocation of IP Addresses might be deferred depending on
	// the IXP implementation. No assumption should be made.
	IPs []string `json:"ips,omitempty"`

	// Listed The customer wants to be featured on the member list
	Listed bool `json:"listed,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ASNs is a asns
	ASNs []int `json:"asns,omitempty"`

	// Macs A list of mac-address IDs.
	Macs []string `json:"macs,omitempty"`

	// IPs A list of ip-address IDs.
	//
	// Allocation of IP Addresses might be deferred depending on
	// the IXP implementation. No assumption should be made.
	IPs []string `json:"ips,omitempty"`

	// Listed The customer wants to be featured on the member list
	Listed bool `json:"listed,omitempty"`
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

// MP2MPNetworkServiceConfig MP2MP Network Service Config
type MP2MPNetworkServiceConfig struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService The id of the configured network service.
	NetworkService string `json:"network_service,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

	// Macs is a macs
	Macs []string `json:"macs,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfig) PolymorphicType() string {
	return MP2MPNetworkServiceConfigType
}

// MP2MPNetworkServiceConfigPatch MP2MP Network Service Config Update
type MP2MPNetworkServiceConfigPatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection *string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// Macs is a macs
	Macs []string `json:"macs,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigPatch) PolymorphicType() string {
	return MP2MPNetworkServiceConfigPatchType
}

// MP2MPNetworkServiceConfigRequest MP2MP Network Service Config Request
type MP2MPNetworkServiceConfigRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// NetworkService The id of the `NetworkService` to configure.
	NetworkService string `json:"network_service,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

	// Macs is a macs
	Macs []string `json:"macs,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigRequest) PolymorphicType() string {
	return MP2MPNetworkServiceConfigRequestType
}

// MP2MPNetworkServiceConfigUpdate MP2MP Network Service Config Update
type MP2MPNetworkServiceConfigUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// Macs is a macs
	Macs []string `json:"macs,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceConfigUpdate) PolymorphicType() string {
	return MP2MPNetworkServiceConfigUpdateType
}

// NetworkFeatureConfig Polymorphic Network Feature Config
type NetworkFeatureConfig interface {
	Polymorphic
}

// PolymorphicNetworkFeatureConfig is a polymorphic base
type PolymorphicNetworkFeatureConfig struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkFeatureConfig) PolymorphicType() string {
	return n.Type
}

// RouteServerNetworkFeatureConfigType is a polymorphic type value for RouteServerNetworkFeatureConfig
const RouteServerNetworkFeatureConfigType = "route_server"

// NetworkFeatureConfigPatch Polymorphic Network Feauture Config Patch
type NetworkFeatureConfigPatch interface {
	Polymorphic
}

// PolymorphicNetworkFeatureConfigPatch is a polymorphic base
type PolymorphicNetworkFeatureConfigPatch struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkFeatureConfigPatch) PolymorphicType() string {
	return n.Type
}

// RouteServerNetworkFeatureConfigPatchType is a polymorphic type value for RouteServerNetworkFeatureConfigPatch
const RouteServerNetworkFeatureConfigPatchType = "route_server"

// NetworkFeatureConfigRequest Polymorphic Network Feature Config Request
type NetworkFeatureConfigRequest interface {
	Polymorphic
}

// PolymorphicNetworkFeatureConfigRequest is a polymorphic base
type PolymorphicNetworkFeatureConfigRequest struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkFeatureConfigRequest) PolymorphicType() string {
	return n.Type
}

// RouteServerNetworkFeatureConfigRequestType is a polymorphic type value for RouteServerNetworkFeatureConfigRequest
const RouteServerNetworkFeatureConfigRequestType = "route_server"

// NetworkFeatureConfigUpdate Polymorphic Network Feauture Config Update
type NetworkFeatureConfigUpdate interface {
	Polymorphic
}

// PolymorphicNetworkFeatureConfigUpdate is a polymorphic base
type PolymorphicNetworkFeatureConfigUpdate struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkFeatureConfigUpdate) PolymorphicType() string {
	return n.Type
}

// RouteServerNetworkFeatureConfigUpdateType is a polymorphic type value for RouteServerNetworkFeatureConfigUpdate
const RouteServerNetworkFeatureConfigUpdateType = "route_server"

// NetworkServiceConfig Polymorphic Network Service Config
type NetworkServiceConfig interface {
	Polymorphic
}

// PolymorphicNetworkServiceConfig is a polymorphic base
type PolymorphicNetworkServiceConfig struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceConfig) PolymorphicType() string {
	return n.Type
}

// ExchangeLanNetworkServiceConfigType is a polymorphic type value for ExchangeLanNetworkServiceConfig
const ExchangeLanNetworkServiceConfigType = "exchange_lan"

// P2PNetworkServiceConfigType is a polymorphic type value for P2PNetworkServiceConfig
const P2PNetworkServiceConfigType = "p2p_vc"

// P2MPNetworkServiceConfigType is a polymorphic type value for P2MPNetworkServiceConfig
const P2MPNetworkServiceConfigType = "p2mp_vc"

// MP2MPNetworkServiceConfigType is a polymorphic type value for MP2MPNetworkServiceConfig
const MP2MPNetworkServiceConfigType = "mp2mp_vc"

// CloudNetworkServiceConfigType is a polymorphic type value for CloudNetworkServiceConfig
const CloudNetworkServiceConfigType = "cloud_vc"

// NetworkServiceConfigPatch Polymorphic Network Service Config
type NetworkServiceConfigPatch interface {
	Polymorphic
}

// PolymorphicNetworkServiceConfigPatch is a polymorphic base
type PolymorphicNetworkServiceConfigPatch struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceConfigPatch) PolymorphicType() string {
	return n.Type
}

// ExchangeLanNetworkServiceConfigPatchType is a polymorphic type value for ExchangeLanNetworkServiceConfigPatch
const ExchangeLanNetworkServiceConfigPatchType = "exchange_lan"

// P2PNetworkServiceConfigPatchType is a polymorphic type value for P2PNetworkServiceConfigPatch
const P2PNetworkServiceConfigPatchType = "p2p_vc"

// P2MPNetworkServiceConfigPatchType is a polymorphic type value for P2MPNetworkServiceConfigPatch
const P2MPNetworkServiceConfigPatchType = "p2mp_vc"

// MP2MPNetworkServiceConfigPatchType is a polymorphic type value for MP2MPNetworkServiceConfigPatch
const MP2MPNetworkServiceConfigPatchType = "mp2mp_vc"

// CloudNetworkServiceConfigPatchType is a polymorphic type value for CloudNetworkServiceConfigPatch
const CloudNetworkServiceConfigPatchType = "cloud_vc"

// NetworkServiceConfigRequest Polymorhic Network Service Config Request
type NetworkServiceConfigRequest interface {
	Polymorphic
}

// PolymorphicNetworkServiceConfigRequest is a polymorphic base
type PolymorphicNetworkServiceConfigRequest struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceConfigRequest) PolymorphicType() string {
	return n.Type
}

// ExchangeLanNetworkServiceConfigRequestType is a polymorphic type value for ExchangeLanNetworkServiceConfigRequest
const ExchangeLanNetworkServiceConfigRequestType = "exchange_lan"

// P2PNetworkServiceConfigRequestType is a polymorphic type value for P2PNetworkServiceConfigRequest
const P2PNetworkServiceConfigRequestType = "p2p_vc"

// P2MPNetworkServiceConfigRequestType is a polymorphic type value for P2MPNetworkServiceConfigRequest
const P2MPNetworkServiceConfigRequestType = "p2mp_vc"

// MP2MPNetworkServiceConfigRequestType is a polymorphic type value for MP2MPNetworkServiceConfigRequest
const MP2MPNetworkServiceConfigRequestType = "mp2mp_vc"

// CloudNetworkServiceConfigRequestType is a polymorphic type value for CloudNetworkServiceConfigRequest
const CloudNetworkServiceConfigRequestType = "cloud_vc"

// NetworkServiceConfigUpdate Polymorphic Network Service Config
type NetworkServiceConfigUpdate interface {
	Polymorphic
}

// PolymorphicNetworkServiceConfigUpdate is a polymorphic base
type PolymorphicNetworkServiceConfigUpdate struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceConfigUpdate) PolymorphicType() string {
	return n.Type
}

// ExchangeLanNetworkServiceConfigUpdateType is a polymorphic type value for ExchangeLanNetworkServiceConfigUpdate
const ExchangeLanNetworkServiceConfigUpdateType = "exchange_lan"

// P2PNetworkServiceConfigUpdateType is a polymorphic type value for P2PNetworkServiceConfigUpdate
const P2PNetworkServiceConfigUpdateType = "p2p_vc"

// P2MPNetworkServiceConfigUpdateType is a polymorphic type value for P2MPNetworkServiceConfigUpdate
const P2MPNetworkServiceConfigUpdateType = "p2mp_vc"

// MP2MPNetworkServiceConfigUpdateType is a polymorphic type value for MP2MPNetworkServiceConfigUpdate
const MP2MPNetworkServiceConfigUpdateType = "mp2mp_vc"

// CloudNetworkServiceConfigUpdateType is a polymorphic type value for CloudNetworkServiceConfigUpdate
const CloudNetworkServiceConfigUpdateType = "cloud_vc"

// P2MPNetworkServiceConfig P2MP Network Service Config
type P2MPNetworkServiceConfig struct {
	// Type is a type
	Type string `json:"type,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

	// Role A `leaf` can only reach roots and is
	// isolated from other leafs. A `root` can
	// reach any other point in the virtual circuit
	// including other roots.
	Role *string `json:"role,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService The id of the configured network service.
	NetworkService string `json:"network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkServiceConfig) PolymorphicType() string {
	return P2MPNetworkServiceConfigType
}

// P2MPNetworkServiceConfigPatch P2MP Network Service Config Update
type P2MPNetworkServiceConfigPatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection *string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// NetworkService The id of the `NetworkService` to configure.
	NetworkService string `json:"network_service,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
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
func (p P2MPNetworkServiceConfigRequest) PolymorphicType() string {
	return P2MPNetworkServiceConfigRequestType
}

// P2MPNetworkServiceConfigUpdate P2MP Network Service Config Update
type P2MPNetworkServiceConfigUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService The id of the configured network service.
	NetworkService string `json:"network_service,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfig) PolymorphicType() string {
	return P2PNetworkServiceConfigType
}

// P2PNetworkServiceConfigPatch P2P Network Service Config Update
type P2PNetworkServiceConfigPatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection *string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigPatch) PolymorphicType() string {
	return P2PNetworkServiceConfigPatchType
}

// P2PNetworkServiceConfigRequest P2P Network Service Config Request
type P2PNetworkServiceConfigRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigRequest) PolymorphicType() string {
	return P2PNetworkServiceConfigRequestType
}

// P2PNetworkServiceConfigUpdate P2P Network Service Config Update
type P2PNetworkServiceConfigUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// VLANConfig is a vlan_config
	VLANConfig VLANConfig `json:"vlan_config,omitempty"`

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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection The id of the connection to use for this `NetworkServiceConfig`.
	Connection string `json:"connection,omitempty"`

	// NetworkFeatureConfigs A list of ids of `NetworkFeatureConfig`s.
	//
	NetworkFeatureConfigs []string `json:"network_feature_configs,omitempty"`

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
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceConfigUpdate) PolymorphicType() string {
	return P2PNetworkServiceConfigUpdateType
}

// Port Port
type Port struct {
	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// Connection is a connection
	Connection *string `json:"connection,omitempty"`

	// Speed is a speed
	Speed *int `json:"speed,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

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
	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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

	// ID is a id
	ID string `json:"id,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// NetworkFeature is a network_feature
	NetworkFeature string `json:"network_feature,omitempty"`

	// NetworkServiceConfig is a network_service_config
	NetworkServiceConfig string `json:"network_service_config,omitempty"`

	// Flags A list of IXP specific feature flag configs. This can be used
	// to enable or disable a specific feature flag.
	Flags []*IXPSpecificFeatureFlagConfig `json:"flags,omitempty"`

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

	// ID is a id
	ID string `json:"id,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (r RouteServerNetworkFeatureConfig) PolymorphicType() string {
	return RouteServerNetworkFeatureConfigType
}

// RouteServerNetworkFeatureConfigPatch Route Server Network Feature Config Update
type RouteServerNetworkFeatureConfigPatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	RoleAssignments []string `json:"role_assignments,omitempty"`

	// NetworkFeature is a network_feature
	NetworkFeature string `json:"network_feature,omitempty"`

	// NetworkServiceConfig is a network_service_config
	NetworkServiceConfig string `json:"network_service_config,omitempty"`

	// Flags A list of IXP specific feature flag configs. This can be used
	// to enable or disable a specific feature flag.
	Flags []*IXPSpecificFeatureFlagConfig `json:"flags,omitempty"`

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
func (r RouteServerNetworkFeatureConfigRequest) PolymorphicType() string {
	return RouteServerNetworkFeatureConfigRequestType
}

// RouteServerNetworkFeatureConfigUpdate Route Server Network Feature Config Update
type RouteServerNetworkFeatureConfigUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

// VLANConfigDot1Q A Dot1Q vlan configuration
type VLANConfigDot1Q struct {
	// VLANType is a vlan_type
	VLANType string `json:"vlan_type,omitempty"`

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

// VLANConfigPort A Port vlan configuration
type VLANConfigPort struct {
	// VLANType is a vlan_type
	VLANType string `json:"vlan_type,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (v VLANConfigPort) PolymorphicType() string {
	return VLANConfigPortType
}

// VLANConfigQinQ A QinQ vlan configuration
type VLANConfigQinQ struct {
	// VLANType is a vlan_type
	VLANType string `json:"vlan_type,omitempty"`

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
type VLANConfig interface {
	Polymorphic
}

// PolymorphicVLANConfig is a polymorphic base
type PolymorphicVLANConfig struct {
	VLANType string `json:"vlan_type"`
}

// PolymorphicType implements the polymorphic interface
func (v PolymorphicVLANConfig) PolymorphicType() string {
	return v.VLANType
}

// VLANConfigDot1QType is a polymorphic type value for VLANConfigDot1Q
const VLANConfigDot1QType = "dot1q"

// VLANConfigQinQType is a polymorphic type value for VLANConfigQinQ
const VLANConfigQinQType = "qinq"

// VLANConfigPortType is a polymorphic type value for VLANConfigPort
const VLANConfigPortType = "port"

// Account Account
type Account struct {
	// State is a state
	State *string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

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

	// ID is a id
	ID string `json:"id,omitempty"`

	// Address is a address
	Address *Address `json:"address,omitempty"`
}

// AccountPatch Account Update
type AccountPatch struct {
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// Name A name of a person or an organisation
	Name *string `json:"name,omitempty"`

	// Telephone The telephone number in E.164 Phone Number Formatting
	Telephone *string `json:"telephone,omitempty"`

	// Email The email of the legal company entity.
	//
	Email *string `json:"email,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`
}

// ContactPatch Contact Update
type ContactPatch struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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

	// ID is a id
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

	// ID is a id
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

	// Account is a account
	Account string `json:"account,omitempty"`

	// Type is a type
	Type string `json:"type,omitempty"`

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

// APIExtensions Implementation specific API extensions
type APIExtensions struct {
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
}

// IPAddress IP-Address
type IPAddress struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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

	// ID is a id
	ID string `json:"id,omitempty"`
}

// IPAddressPatch IP-Address Update
type IPAddressPatch struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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

// IPAddressUpdate IP-Address Update
type IPAddressUpdate struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// Address Unicast MAC address, formatted hexadecimal values with colons.
	//
	Address string `json:"address,omitempty"`

	// ValidNotBefore is a valid_not_before
	ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

	// ValidNotAfter is a valid_not_after
	ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`
}

// MacAddressRequest MAC-Address Request
type MacAddressRequest struct {
	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// Address Unicast MAC address, formatted hexadecimal values with colons.
	//
	Address string `json:"address,omitempty"`

	// ValidNotBefore is a valid_not_before
	ValidNotBefore *time.Time `json:"valid_not_before,omitempty"`

	// ValidNotAfter is a valid_not_after
	ValidNotAfter *time.Time `json:"valid_not_after,omitempty"`
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
	ResourceProperty string `json:"resource_property,omitempty"`

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
	// Type A URI reference (see RFC3986) that identifies the
	// problem type.
	//
	// This specification encourages that, when
	// dereferenced, it provide human-readable documentation
	// for the problem type (e.g., using HTML
	// [W3C.REC-html5-20141028]).
	//
	// When this member is not present, its value is assumed
	// to be "about:blank".
	//
	Type string `json:"type,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// CapacityMin Require an optional minimum capacity to join
	// the network service.
	CapacityMin *int `json:"capacity_min,omitempty"`

	// CapacityMax An optional rate limit which has precedence over
	// the capacity set in the network service config.
	CapacityMax *int `json:"capacity_max,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRule) PolymorphicType() string {
	return AllowMemberJoiningRuleType
}

// AllowMemberJoiningRulePatch A vlan member joining rule update
type AllowMemberJoiningRulePatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount *string `json:"consuming_account,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// CapacityMin Require an optional minimum capacity to join
	// the network service.
	CapacityMin *int `json:"capacity_min,omitempty"`

	// CapacityMax An optional rate limit which has precedence over
	// the capacity set in the network service config.
	CapacityMax *int `json:"capacity_max,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (a AllowMemberJoiningRuleRequest) PolymorphicType() string {
	return AllowMemberJoiningRuleRequestType
}

// AllowMemberJoiningRuleUpdate A vlan member joining rule update
type AllowMemberJoiningRuleUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NscRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The `RoleAssignment` is associated through the
	// `role_assignments` list property of the network service configuration.
	NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// CloudKey is a cloud_key
	CloudKey string `json:"cloud_key,omitempty"`

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
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkService) PolymorphicType() string {
	return CloudNetworkServiceType
}

// CloudNetworkServicePatch Cloud Network Service Update
type CloudNetworkServicePatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering *string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// CloudKey is a cloud_key
	CloudKey *string `json:"cloud_key,omitempty"`

	// Capacity The capacity of the service in Mbps. When null,
	// the maximum capacity will be used.
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServicePatch) PolymorphicType() string {
	return CloudNetworkServicePatchType
}

// CloudNetworkServiceRequest Cloud Network Service Request
type CloudNetworkServiceRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// CloudKey is a cloud_key
	CloudKey string `json:"cloud_key,omitempty"`

	// Capacity The capacity of the service in Mbps. When null,
	// the maximum capacity will be used.
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceRequest) PolymorphicType() string {
	return CloudNetworkServiceRequestType
}

// CloudNetworkServiceUpdate Cloud Network Service Update
type CloudNetworkServiceUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// CloudKey is a cloud_key
	CloudKey string `json:"cloud_key,omitempty"`

	// Capacity The capacity of the service in Mbps. When null,
	// the maximum capacity will be used.
	Capacity *int `json:"capacity,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (c CloudNetworkServiceUpdate) PolymorphicType() string {
	return CloudNetworkServiceUpdateType
}

// DenyMemberJoiningRule A rule for members joining a private vlan
type DenyMemberJoiningRule struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRule) PolymorphicType() string {
	return DenyMemberJoiningRuleType
}

// DenyMemberJoiningRulePatch A vlan member joining rule update
type DenyMemberJoiningRulePatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account to which access to the
	// network service should be granted or denied.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (d DenyMemberJoiningRuleRequest) PolymorphicType() string {
	return DenyMemberJoiningRuleRequestType
}

// DenyMemberJoiningRuleUpdate A vlan member joining rule update
type DenyMemberJoiningRuleUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NscRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The `RoleAssignment` is associated through the
	// `role_assignments` list property of the network service configuration.
	NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
	//
	ConsumingAccount string `json:"consuming_account,omitempty"`

	// ExternalRef Reference field, free to use for the API user.
	// *(Sensitive Property)*
	//
	ExternalRef *string `json:"external_ref,omitempty"`

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

	// NetworkFeatures is a network_features
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NscRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The `RoleAssignment` is associated through the
	// `role_assignments` list property of the network service configuration.
	NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// NscProductOfferings An optional list of `ProductOffering`s which can be used in the
	// network service configs for this service.
	NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// MemberJoiningRules is a member_joining_rules
	MemberJoiningRules []string `json:"member_joining_rules,omitempty"`

	// NetworkFeatures is a network_features
	NetworkFeatures []string `json:"network_features,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkService) PolymorphicType() string {
	return MP2MPNetworkServiceType
}

// MP2MPNetworkServicePatch MP2MP Network Service Update
type MP2MPNetworkServicePatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering *string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServicePatch) PolymorphicType() string {
	return MP2MPNetworkServicePatchType
}

// MP2MPNetworkServiceRequest MP2MP Network Service Request
type MP2MPNetworkServiceRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceRequest) PolymorphicType() string {
	return MP2MPNetworkServiceRequestType
}

// MP2MPNetworkServiceUpdate MP2MP Network Service Update
type MP2MPNetworkServiceUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (m MP2MPNetworkServiceUpdate) PolymorphicType() string {
	return MP2MPNetworkServiceUpdateType
}

// MemberJoiningRule Polymorphic Member Joining Rule
type MemberJoiningRule interface {
	Polymorphic
}

// PolymorphicMemberJoiningRule is a polymorphic base
type PolymorphicMemberJoiningRule struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (m PolymorphicMemberJoiningRule) PolymorphicType() string {
	return m.Type
}

// AllowMemberJoiningRuleType is a polymorphic type value for AllowMemberJoiningRule
const AllowMemberJoiningRuleType = "allow"

// DenyMemberJoiningRuleType is a polymorphic type value for DenyMemberJoiningRule
const DenyMemberJoiningRuleType = "deny"

// MemberJoiningRulePatch Polymorphic Member Joining Rule Update
type MemberJoiningRulePatch interface {
	Polymorphic
}

// PolymorphicMemberJoiningRulePatch is a polymorphic base
type PolymorphicMemberJoiningRulePatch struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (m PolymorphicMemberJoiningRulePatch) PolymorphicType() string {
	return m.Type
}

// AllowMemberJoiningRulePatchType is a polymorphic type value for AllowMemberJoiningRulePatch
const AllowMemberJoiningRulePatchType = "allow"

// DenyMemberJoiningRulePatchType is a polymorphic type value for DenyMemberJoiningRulePatch
const DenyMemberJoiningRulePatchType = "deny"

// MemberJoiningRuleRequest Polymorphic Member Joining Rule Request
type MemberJoiningRuleRequest interface {
	Polymorphic
}

// PolymorphicMemberJoiningRuleRequest is a polymorphic base
type PolymorphicMemberJoiningRuleRequest struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (m PolymorphicMemberJoiningRuleRequest) PolymorphicType() string {
	return m.Type
}

// AllowMemberJoiningRuleRequestType is a polymorphic type value for AllowMemberJoiningRuleRequest
const AllowMemberJoiningRuleRequestType = "allow"

// DenyMemberJoiningRuleRequestType is a polymorphic type value for DenyMemberJoiningRuleRequest
const DenyMemberJoiningRuleRequestType = "deny"

// MemberJoiningRuleUpdate Polymorphic Member Joining Rule Update
type MemberJoiningRuleUpdate interface {
	Polymorphic
}

// PolymorphicMemberJoiningRuleUpdate is a polymorphic base
type PolymorphicMemberJoiningRuleUpdate struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (m PolymorphicMemberJoiningRuleUpdate) PolymorphicType() string {
	return m.Type
}

// AllowMemberJoiningRuleUpdateType is a polymorphic type value for AllowMemberJoiningRuleUpdate
const AllowMemberJoiningRuleUpdateType = "allow"

// DenyMemberJoiningRuleUpdateType is a polymorphic type value for DenyMemberJoiningRuleUpdate
const DenyMemberJoiningRuleUpdateType = "deny"

// NetworkFeature Polymorphic Network Feature
type NetworkFeature interface {
	Polymorphic
}

// PolymorphicNetworkFeature is a polymorphic base
type PolymorphicNetworkFeature struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkFeature) PolymorphicType() string {
	return n.Type
}

// RouteServerNetworkFeatureType is a polymorphic type value for RouteServerNetworkFeature
const RouteServerNetworkFeatureType = "route_server"

// NetworkService Polymorphic Network Services
type NetworkService interface {
	Polymorphic
}

// PolymorphicNetworkService is a polymorphic base
type PolymorphicNetworkService struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkService) PolymorphicType() string {
	return n.Type
}

// ExchangeLanNetworkServiceType is a polymorphic type value for ExchangeLanNetworkService
const ExchangeLanNetworkServiceType = "exchange_lan"

// P2PNetworkServiceType is a polymorphic type value for P2PNetworkService
const P2PNetworkServiceType = "p2p_vc"

// P2MPNetworkServiceType is a polymorphic type value for P2MPNetworkService
const P2MPNetworkServiceType = "p2mp_vc"

// MP2MPNetworkServiceType is a polymorphic type value for MP2MPNetworkService
const MP2MPNetworkServiceType = "mp2mp_vc"

// CloudNetworkServiceType is a polymorphic type value for CloudNetworkService
const CloudNetworkServiceType = "cloud_vc"

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
type NetworkServiceDeleteResponse interface {
	Polymorphic
}

// PolymorphicNetworkServiceDeleteResponse is a polymorphic base
type PolymorphicNetworkServiceDeleteResponse struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceDeleteResponse) PolymorphicType() string {
	return n.Type
}

// NetworkServicePatch Polymorphic Network Service Patch
type NetworkServicePatch interface {
	Polymorphic
}

// PolymorphicNetworkServicePatch is a polymorphic base
type PolymorphicNetworkServicePatch struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServicePatch) PolymorphicType() string {
	return n.Type
}

// P2PNetworkServicePatchType is a polymorphic type value for P2PNetworkServicePatch
const P2PNetworkServicePatchType = "p2p_vc"

// P2MPNetworkServicePatchType is a polymorphic type value for P2MPNetworkServicePatch
const P2MPNetworkServicePatchType = "p2mp_vc"

// MP2MPNetworkServicePatchType is a polymorphic type value for MP2MPNetworkServicePatch
const MP2MPNetworkServicePatchType = "mp2mp_vc"

// CloudNetworkServicePatchType is a polymorphic type value for CloudNetworkServicePatch
const CloudNetworkServicePatchType = "cloud_vc"

// NetworkServiceRequest Polymorphic Network Service Request
type NetworkServiceRequest interface {
	Polymorphic
}

// PolymorphicNetworkServiceRequest is a polymorphic base
type PolymorphicNetworkServiceRequest struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceRequest) PolymorphicType() string {
	return n.Type
}

// P2PNetworkServiceRequestType is a polymorphic type value for P2PNetworkServiceRequest
const P2PNetworkServiceRequestType = "p2p_vc"

// P2MPNetworkServiceRequestType is a polymorphic type value for P2MPNetworkServiceRequest
const P2MPNetworkServiceRequestType = "p2mp_vc"

// MP2MPNetworkServiceRequestType is a polymorphic type value for MP2MPNetworkServiceRequest
const MP2MPNetworkServiceRequestType = "mp2mp_vc"

// CloudNetworkServiceRequestType is a polymorphic type value for CloudNetworkServiceRequest
const CloudNetworkServiceRequestType = "cloud_vc"

// NetworkServiceUpdate Polymorphic Network Service Update
type NetworkServiceUpdate interface {
	Polymorphic
}

// PolymorphicNetworkServiceUpdate is a polymorphic base
type PolymorphicNetworkServiceUpdate struct {
	Type string `json:"type"`
}

// PolymorphicType implements the polymorphic interface
func (n PolymorphicNetworkServiceUpdate) PolymorphicType() string {
	return n.Type
}

// P2PNetworkServiceUpdateType is a polymorphic type value for P2PNetworkServiceUpdate
const P2PNetworkServiceUpdateType = "p2p_vc"

// P2MPNetworkServiceUpdateType is a polymorphic type value for P2MPNetworkServiceUpdate
const P2MPNetworkServiceUpdateType = "p2mp_vc"

// MP2MPNetworkServiceUpdateType is a polymorphic type value for MP2MPNetworkServiceUpdate
const MP2MPNetworkServiceUpdateType = "mp2mp_vc"

// CloudNetworkServiceUpdateType is a polymorphic type value for CloudNetworkServiceUpdate
const CloudNetworkServiceUpdateType = "cloud_vc"

// P2MPNetworkService P2MP Network Service
type P2MPNetworkService struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NscRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The `RoleAssignment` is associated through the
	// `role_assignments` list property of the network service configuration.
	NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// NscProductOfferings An optional list of `ProductOffering`s which can be used in the
	// network service configs for this service.
	NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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

	// NetworkFeatures is a network_features
	NetworkFeatures []string `json:"network_features,omitempty"`

	// MemberJoiningRules is a member_joining_rules
	MemberJoiningRules []string `json:"member_joining_rules,omitempty"`
}

// PolymorphicType implements the polymorphic interface
func (p P2MPNetworkService) PolymorphicType() string {
	return P2MPNetworkServiceType
}

// P2MPNetworkServicePatch P2MP Network Service Update
type P2MPNetworkServicePatch struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering *string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// State is a state
	State string `json:"state,omitempty"`

	// Status is a status
	Status []*Status `json:"status,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// NscRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The `RoleAssignment` is associated through the
	// `role_assignments` list property of the network service configuration.
	NscRequiredContactRoles []string `json:"nsc_required_contact_roles,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// NscProductOfferings An optional list of `ProductOffering`s which can be used in the
	// network service configs for this service.
	NscProductOfferings []string `json:"nsc_product_offerings,omitempty"`

	// DecommissionAt The service will be decommissioned on this date.
	//
	// This field is only used when
	// the state is `DECOMMISSION_REQUESTED` or
	// `DECOMMISSIONED`.
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
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering *string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount *string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServicePatch) PolymorphicType() string {
	return P2PNetworkServicePatchType
}

// P2PNetworkServiceRequest P2P Network Service Request
type P2PNetworkServiceRequest struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceRequest) PolymorphicType() string {
	return P2PNetworkServiceRequestType
}

// P2PNetworkServiceUpdate P2P Network Service Update
type P2PNetworkServiceUpdate struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ProductOffering is a product_offering
	ProductOffering string `json:"product_offering,omitempty"`

	// ManagingAccount The `id` of the account responsible for managing the service via
	// the API. A manager can read and update the state of entities.
	//
	ManagingAccount string `json:"managing_account,omitempty"`

	// ConsumingAccount The `id` of the account consuming a service.
	//
	// Used to be `owning_customer`.
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
}

// PolymorphicType implements the polymorphic interface
func (p P2PNetworkServiceUpdate) PolymorphicType() string {
	return P2PNetworkServiceUpdateType
}

// RouteServerNetworkFeature Route Server Network Feature
type RouteServerNetworkFeature struct {
	// Type is a type
	Type string `json:"type,omitempty"`

	// ID is a id
	ID string `json:"id,omitempty"`

	// Name is a name
	Name string `json:"name,omitempty"`

	// Required is a required
	Required bool `json:"required,omitempty"`

	// NfcRequiredContactRoles The configuration will require at least one of each of the
	// specified roles assigned to contacts.
	//
	// The role assignments is associated with the network feature
	// config through the `role_assignments` list property.
	NfcRequiredContactRoles []string `json:"nfc_required_contact_roles,omitempty"`

	// Flags A list of IXP specific feature flags. This can be used
	// to see if e.g. RPKI hard filtering is available.
	Flags []*IXPSpecificFeatureFlag `json:"flags,omitempty"`

	// NetworkService is a network_service
	NetworkService string `json:"network_service,omitempty"`

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

// PortStatistics Port Statistics
type PortStatistics struct {
	// Aggregates Aggregated statistics for a connection or service configuration
	//
	// For the **property name** the aggregate interval as a
	// string representation is used. For example: `5m`, `1d`, `30d`,
	// `1y`.
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
