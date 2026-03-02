package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func CloudRouterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},

		"state": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},

		"status": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: StatusSchema()},
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The account responsible for managing the Cloud ROUTER",
		},

		"billing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The account used for billing",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The account consuming the Cloud ROUTER service",
		},

		"external_ref": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "External reference for the Cloud ROUTER",
		},

		"product_offering": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Product offering ID for the Cloud ROUTER",
		},

		"asn": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			ForceNew:    true,
			Description: "Autonomous System Number (ASN) for the Cloud ROUTER",
		},

		"capacity": &schema.Schema{
			Type:        schema.TypeInt,
			Required:    true,
			ForceNew:    true,
			Description: "Bandwidth capacity in Mbps",
		},

		"metro_area_network": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Metro area network where Cloud ROUTER is deployed",
		},

		"decommission_at": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Service will be decommissioned on this date",
		},

		"charged_until": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Service continues incurring charges until this date",
		},
	}
}

func CloudRouterNetworkServiceConfigCommonSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeList,
			Computed: true,
			Elem:     &schema.Resource{Schema: StatusSchema()},
		},
		"managing_account": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Managing account ID",
		},
		"billing_account": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Billing account ID",
		},
		"consuming_account": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Consuming account ID",
		},
		"external_ref": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "External reference",
		},
		"cloud_router": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Cloud ROUTER (VRF) ID",
		},
		"network_service": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Network service ID to connect to the Cloud ROUTER",
		},
		"address": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "IP address for the BGP session (CIDR notation)",
		},
		"bgp_neighbor": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "BGP neighbor IP address",
		},
		"bgp_neighbor_asn": {
			Type:        schema.TypeInt,
			Required:    true,
			ForceNew:    true,
			Description: "BGP neighbor Autonomous System Number",
		},
		"bgp_password": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Sensitive:   true,
			Description: "BGP session password",
		},
		"vlan_config": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "VLAN configuration for the connection",
			Elem: &schema.Resource{
				Schema: VLANConfigSchema(),
			},
		},
		"policy_ingress": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the inbound routing policy",
		},
		"policy_egress": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the outbound routing policy",
		},
		"admin_status": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Administrative status (enabled/disabled)",
		},
		"bfd_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable Bidirectional Forwarding Detection (BFD)",
		},
		"purchase_order": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "Purchase order reference",
		},
		"network_feature_configs": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "Network feature configuration IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func CloudRouterNetworkServiceConfigCloudVCSchema() map[string]*schema.Schema {
	s := CloudRouterNetworkServiceConfigCommonSchema()

	s["cloud_vlan"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		ForceNew:    true,
		Description: "Cloud VLAN ID for cloud connections",
	}
	s["handover"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		ForceNew:    true,
		Description: "Handover ID",
	}
	s["connection_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Connection ID (read-only, if present in API response)",
	}

	return s
}

func CloudRouterNetworkServiceConfigP2PVCSchema() map[string]*schema.Schema {
	s := CloudRouterNetworkServiceConfigCommonSchema()

	s["nic"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Network Interface Card (NIC/Connection) ID to use for this network service config",
	}
	s["connection_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Connection ID (read-only)",
	}

	return s
}

func PrefixListSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the prefix list",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Managing account ID",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Consuming account ID",
		},

		"match_list": &schema.Schema{
			Type:        schema.TypeList,
			Required:    true,
			Description: "List of prefix matches",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"prefix": &schema.Schema{
						Type:        schema.TypeString,
						Required:    true,
						Description: "IP prefix in CIDR notation (e.g., 192.168.0.0/16)",
					},
					"min_length": &schema.Schema{
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Minimum prefix length to match",
					},
					"max_length": &schema.Schema{
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Maximum prefix length to match",
					},
				},
			},
		},
	}
}

func PolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the policy",
		},

		"managing_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Managing account ID",
		},

		"consuming_account": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Consuming account ID",
		},

		"entries": &schema.Schema{
			Type:        schema.TypeList,
			Required:    true,
			Description: "List of policy entries",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"sequence_number": &schema.Schema{
						Type:        schema.TypeInt,
						Required:    true,
						Description: "Sequence number for the policy entry",
					},
					"match_prefix_list": &schema.Schema{
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Name of the prefix list to match",
					},
					"action": &schema.Schema{
						Type:        schema.TypeList,
						Required:    true,
						MaxItems:    1,
						Description: "Action to take when matched",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"filter": &schema.Schema{
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Filter action: accept, reject, or continue",
								},
								"local_preference": &schema.Schema{
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "BGP local preference value",
								},
								"as_path_prepend": &schema.Schema{
									Type:        schema.TypeList,
									Optional:    true,
									MaxItems:    1,
									Description: "AS path prepend configuration",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"count": &schema.Schema{
												Type:        schema.TypeInt,
												Required:    true,
												Description: "Number of times to prepend the ASN",
											},
											"asn": &schema.Schema{
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "AS number to prepend (defaults to local ASN)",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func BGPPrefixSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"as_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "AS path",
		},
		"igp_cost": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "IGP cost",
		},
		"network": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Network prefix in CIDR notation",
		},
		"label": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "MPLS label",
		},
		"local_pref": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Local preference",
		},
		"med": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "Multi-Exit Discriminator (MED)",
		},
		"nexthop": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Next hop IP address",
		},
		"path_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "BGP path identifier",
		},
		"flags": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "BGP prefix flags (e.g., used, valid, best, igp)",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"last_queried_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when prefixes were last queried",
		},
	}
}

func BGPRouteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prefix": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Network prefix in CIDR notation",
		},
		"next_hop": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Next hop IP address",
		},
		"received_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when the route was received",
		},
		"as_path": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "AS path as list of ASNs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}
