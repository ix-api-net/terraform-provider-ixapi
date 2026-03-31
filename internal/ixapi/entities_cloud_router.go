package ixapi

import "encoding/json"

// CloudRouter represents a DE-CIX Cloud Router (VRF) instance.
type CloudRouter struct {
	ID               string   `json:"id,omitempty"`
	State            string   `json:"state,omitempty"`
	Status           []Status `json:"status,omitempty"`
	ManagingAccount  string   `json:"managing_account,omitempty"`
	BillingAccount   string   `json:"billing_account,omitempty"`
	ConsumingAccount string   `json:"consuming_account,omitempty"`
	ExternalRef      *string  `json:"external_ref,omitempty"`
	ProductOffering  string   `json:"product_offering,omitempty"`
	ASN              int      `json:"asn,omitempty"`
	Capacity         int      `json:"capacity,omitempty"`
	MetroAreaNetwork string   `json:"metro_area_network,omitempty"`
	DecommissionAt   *string  `json:"decommission_at,omitempty"`
	ChargedUntil     *string  `json:"charged_until,omitempty"`
}

// CloudRouterRequest is the request body for creating a Cloud Router.
type CloudRouterRequest struct {
	ManagingAccount  string  `json:"managing_account"`
	BillingAccount   string  `json:"billing_account"`
	ConsumingAccount string  `json:"consuming_account"`
	ExternalRef      *string `json:"external_ref,omitempty"`
	ProductOffering  string  `json:"product_offering"`
	ASN              int     `json:"asn"`
	Capacity         int     `json:"capacity"`
}

// CloudRouterNetworkServiceConfig represents a network service config attached to a Cloud Router.
type CloudRouterNetworkServiceConfig struct {
	ID                    string      `json:"id,omitempty"`
	State                 string      `json:"state,omitempty"`
	Status                []Status    `json:"status,omitempty"`
	Type                  string      `json:"type,omitempty"`
	ManagingAccount       string      `json:"managing_account,omitempty"`
	BillingAccount        string      `json:"billing_account,omitempty"`
	ConsumingAccount      string      `json:"consuming_account,omitempty"`
	ExternalRef           *string     `json:"external_ref,omitempty"`
	CloudRouter           string      `json:"vrf,omitempty" tf:"cloud_router"`
	NetworkService        string      `json:"network_service,omitempty"`
	Address               string          `json:"address,omitempty"`
	BGPNeighbor           string          `json:"bgp_neighbor,omitempty"`
	BGPNeighborASN        int             `json:"bgp_neighbor_asn,omitempty"`
	BGPPassword           *string         `json:"bgp_password,omitempty"`
	VLANConfig            VLANConfig      `tf:"vlan_config" json:"-"`
	VLANConfigRaw         json.RawMessage `tf:"-" json:"vlan_config,omitempty"`
	PolicyIngress         *string         `json:"policy_ingress,omitempty"`
	PolicyEgress          *string         `json:"policy_egress,omitempty"`
	PolicyIngressID       *string         `json:"policy_ingress_id,omitempty"`
	PolicyEgressID        *string         `json:"policy_egress_id,omitempty"`
	AdminStatus           string      `json:"admin_status,omitempty"`
	BFDEnabled            bool        `json:"bfd,omitempty" tf:"bfd_enabled"`
	CloudVLAN             *int        `json:"cloud_vlan,omitempty"`
	Handover              *int        `json:"handover,omitempty"`
	PurchaseOrder         *string     `json:"purchase_order,omitempty"`
	NetworkFeatureConfigs []string    `json:"network_feature_configs,omitempty"`
	ConnectionID          string      `json:"connection,omitempty" tf:"connection_id"`
}

// CloudRouterNetworkServiceConfigRequest is the request body for creating a Cloud Router network service config.
type CloudRouterNetworkServiceConfigRequest struct{
	Type                  string      `json:"type"`
	ManagingAccount       string      `json:"managing_account"`
	BillingAccount        string      `json:"billing_account"`
	ConsumingAccount      string      `json:"consuming_account"`
	ExternalRef           *string     `json:"external_ref,omitempty"`
	CloudRouter           string      `json:"vrf"`
	NetworkService        string      `json:"network_service"`
	Address               string      `json:"address"`
	BGPNeighbor           string      `json:"bgp_neighbor"`
	BGPNeighborASN        int         `json:"bgp_neighbor_asn"`
	BGPPassword           *string     `json:"bgp_password,omitempty"`
	VLANConfig            interface{} `json:"vlan_config,omitempty"`
	PolicyIngress         *string     `json:"policy_ingress,omitempty"`
	PolicyEgress          *string     `json:"policy_egress,omitempty"`
	AdminStatus           string      `json:"admin_status,omitempty"`
	BFDEnabled            bool        `json:"bfd,omitempty"`
	CloudVLAN             *int        `json:"cloud_vlan,omitempty"`
	Handover              *int        `json:"handover,omitempty"`
	Connection            *string     `json:"connection,omitempty"`
	PurchaseOrder         *string     `json:"purchase_order,omitempty"`
	NetworkFeatureConfigs []string    `json:"network_feature_configs,omitempty"`
}

// CloudRouterNetworkServiceConfigPatch is the request body for partially updating a Cloud Router network service config.
type CloudRouterNetworkServiceConfigPatch struct {
	PolicyIngress *string `json:"policy_ingress,omitempty"`
	PolicyEgress  *string `json:"policy_egress,omitempty"`
	AdminStatus   *string `json:"admin_status,omitempty"`
}

// CloudRouterProductOffering represents a product offering for a DE-CIX Cloud Router (VRF) service.
type CloudRouterProductOffering struct {
	ID                          string `json:"id"`
	DisplayName                 string `json:"display_name"`
	BandwidthMax                int    `json:"bandwidth_max"`
	BandwidthMin                int    `json:"bandwidth_min"`
	Name                        string `json:"name"`
	ServiceMetroArea            string `json:"service_metro_area"`
	ServiceMetroAreaName        string `json:"service_metro_area_name"`
	ServiceMetroAreaNetwork     string `json:"service_metro_area_network"`
	ServiceMetroAreaNetworkName string `json:"service_metro_area_network_name"`
	ContractPeriod              string `json:"contract_period"`
	Type                        string `json:"type"`
}

// BGPStateResponse holds the BGP session state for a network service config.
type BGPStateResponse struct {
	State string `json:"state"`
}

// BFDStateResponse holds the BFD session state for a network service config.
type BFDStateResponse struct {
	State string `json:"state"`
}

// PrefixMatch represents a single prefix match entry in a prefix list.
type PrefixMatch struct {
	Prefix    string `json:"prefix"`
	MinLength *int   `json:"min_length,omitempty"`
	MaxLength *int   `json:"max_length,omitempty"`
}

// PrefixList represents a named list of IP prefixes used for BGP route filtering.
type PrefixList struct {
	ID               string         `json:"id,omitempty"`
	Name             string         `json:"name,omitempty"`
	ManagingAccount  string         `json:"managing_account,omitempty"`
	ConsumingAccount string         `json:"consuming_account,omitempty"`
	MatchList        []PrefixMatch  `json:"match_list,omitempty"`
}

// PrefixListRequest is the request body for creating or updating a prefix list.
type PrefixListRequest struct {
	Name             string        `json:"name"`
	ManagingAccount  string        `json:"managing_account"`
	ConsumingAccount string        `json:"consuming_account"`
	MatchList        []PrefixMatch `json:"match_list"`
}

// ASPathPrepend configures BGP AS path prepending in a policy action.
type ASPathPrepend struct {
	Count int     `json:"count"`
	ASN   *int    `json:"asn,omitempty"`
}

// PolicyAction defines the action to take when a policy entry matches.
type PolicyAction struct {
	Filter          *string        `json:"filter,omitempty"`
	LocalPreference *int           `json:"local_preference,omitempty"`
	ASPathPrepend   *ASPathPrepend `json:"as_path_prepend,omitempty"`
}

// PolicyEntry represents a single rule in a BGP routing policy.
type PolicyEntry struct {
	SequenceNumber  int          `json:"sequence_number"`
	MatchPrefixList *string      `json:"match_prefix_list,omitempty"`
	Action          PolicyAction `json:"action"`
}

// Policy represents a BGP routing policy consisting of ordered entries.
type Policy struct {
	ID               string        `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	ManagingAccount  string        `json:"managing_account,omitempty"`
	ConsumingAccount string        `json:"consuming_account,omitempty"`
	Entries          []PolicyEntry `json:"entries,omitempty"`
}

// PolicyRequest is the request body for creating or updating a routing policy.
type PolicyRequest struct {
	Name             string        `json:"name"`
	ManagingAccount  string        `json:"managing_account"`
	ConsumingAccount string        `json:"consuming_account"`
	Entries          []PolicyEntry `json:"entries"`
}

// BGPRoute represents a BGP route entry with prefix, next hop, and AS path information.
type BGPRoute struct {
	Prefix     string   `json:"prefix,omitempty"`
	NextHop    string   `json:"next_hop,omitempty"`
	ReceivedAt string   `json:"received_at,omitempty"`
	ASPath     []string `json:"as_path,omitempty"`
}

// ArpEntry represents a single ARP table entry for a Cloud Router VRF.
type ArpEntry struct {
	VRF                  string `json:"vrf,omitempty"`
	NetworkServiceConfig string `json:"network_service_config,omitempty"`
	DeviceFQDN           string `json:"device_fqdn,omitempty"`
	IPAddress            string `json:"ip_address,omitempty"`
	MACAddress           string `json:"mac_address,omitempty"`
	ExpirationTime       int    `json:"expiration_time,omitempty"`
	ReceivedAt           string `json:"received_at,omitempty"`
}

// StaticRoute represents a static route attached to a Cloud Router VRF.
type StaticRoute struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Prefix               string   `json:"prefix,omitempty"`
	NextHop              string   `json:"next_hop,omitempty"`
	NetworkServiceConfigs []string `json:"network_service_configs,omitempty"`
	VRF                  string   `json:"vrf,omitempty"`
}

// VrfRoute represents a route entry in a Cloud Router VRF routing table.
type VrfRoute struct {
	VRF                  string `json:"vrf,omitempty"`
	ReceivedAt           string `json:"received_at,omitempty"`
	Prefix               string `json:"prefix,omitempty"`
	DeviceFQDN           string `json:"device_fqdn,omitempty"`
	Metric               int    `json:"metric,omitempty"`
	Protocol             string `json:"protocol,omitempty"`
	Distance             int    `json:"distance,omitempty"`
	NextHop              string `json:"next_hop,omitempty"`
	NetworkServiceConfig string `json:"network_service_config,omitempty"`
}

// StaticRouteRequest is the request body for creating or updating a static route.
type StaticRouteRequest struct {
	Name                 string   `json:"name"`
	Prefix               string   `json:"prefix"`
	NextHop              string   `json:"next_hop"`
	NetworkServiceConfigs []string `json:"network_service_configs"`
}
