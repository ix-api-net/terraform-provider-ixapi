package ixapi

import "encoding/json"

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

type CloudRouterRequest struct {
	ManagingAccount  string  `json:"managing_account"`
	BillingAccount   string  `json:"billing_account"`
	ConsumingAccount string  `json:"consuming_account"`
	ExternalRef      *string `json:"external_ref,omitempty"`
	ProductOffering  string  `json:"product_offering"`
	ASN              int     `json:"asn"`
	Capacity         int     `json:"capacity"`
}

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

type CloudRouterNetworkServiceConfigPatch struct {
	PolicyIngress *string `json:"policy_ingress,omitempty"`
	PolicyEgress  *string `json:"policy_egress,omitempty"`
	AdminStatus   *string `json:"admin_status,omitempty"`
}

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

type BGPStateResponse struct {
	State string `json:"state"`
}

type BFDStateResponse struct {
	State string `json:"state"`
}

type PrefixMatch struct {
	Prefix    string `json:"prefix"`
	MinLength *int   `json:"min_length,omitempty"`
	MaxLength *int   `json:"max_length,omitempty"`
}

type PrefixList struct {
	ID               string         `json:"id,omitempty"`
	Name             string         `json:"name,omitempty"`
	ManagingAccount  string         `json:"managing_account,omitempty"`
	ConsumingAccount string         `json:"consuming_account,omitempty"`
	MatchList        []PrefixMatch  `json:"match_list,omitempty"`
}

type PrefixListRequest struct {
	Name             string        `json:"name"`
	ManagingAccount  string        `json:"managing_account"`
	ConsumingAccount string        `json:"consuming_account"`
	MatchList        []PrefixMatch `json:"match_list"`
}

type ASPathPrepend struct {
	Count int     `json:"count"`
	ASN   *int    `json:"asn,omitempty"`
}

type PolicyAction struct {
	Filter          *string        `json:"filter,omitempty"`
	LocalPreference *int           `json:"local_preference,omitempty"`
	ASPathPrepend   *ASPathPrepend `json:"as_path_prepend,omitempty"`
}

type PolicyEntry struct {
	SequenceNumber  int          `json:"sequence_number"`
	MatchPrefixList *string      `json:"match_prefix_list,omitempty"`
	Action          PolicyAction `json:"action"`
}

type Policy struct {
	ID               string        `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	ManagingAccount  string        `json:"managing_account,omitempty"`
	ConsumingAccount string        `json:"consuming_account,omitempty"`
	Entries          []PolicyEntry `json:"entries,omitempty"`
}

type PolicyRequest struct {
	Name             string        `json:"name"`
	ManagingAccount  string        `json:"managing_account"`
	ConsumingAccount string        `json:"consuming_account"`
	Entries          []PolicyEntry `json:"entries"`
}

type BGPRoute struct {
	Prefix     string   `json:"prefix,omitempty"`
	NextHop    string   `json:"next_hop,omitempty"`
	ReceivedAt string   `json:"received_at,omitempty"`
	ASPath     []string `json:"as_path,omitempty"`
}

type ArpEntry struct {
	VRF                  string `json:"vrf,omitempty"`
	NetworkServiceConfig string `json:"network_service_config,omitempty"`
	DeviceFQDN           string `json:"device_fqdn,omitempty"`
	IPAddress            string `json:"ip_address,omitempty"`
	MACAddress           string `json:"mac_address,omitempty"`
	ExpirationTime       int    `json:"expiration_time,omitempty"`
	ReceivedAt           string `json:"received_at,omitempty"`
}

type StaticRoute struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Prefix               string   `json:"prefix,omitempty"`
	NextHop              string   `json:"next_hop,omitempty"`
	NetworkServiceConfigs []string `json:"network_service_configs,omitempty"`
	VRF                  string   `json:"vrf,omitempty"`
}

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

type StaticRouteRequest struct {
	Name                 string   `json:"name"`
	Prefix               string   `json:"prefix"`
	NextHop              string   `json:"next_hop"`
	NetworkServiceConfigs []string `json:"network_service_configs"`
}
