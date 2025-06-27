package provider

type DeviceGatewayModel struct {
	AdditionalConfigCmds  []string                            `hcl:"additional_config_cmds"`
	BgpConfig             map[string]BgpConfigValue           `hcl:"bgp_config"`
	DeviceId              string                              `hcl:"device_id"`
	DhcpdConfig           *DhcpdConfigValue                   `hcl:"dhcpd_config"`
	DnsServers            []string                            `hcl:"dns_servers"`
	DnsSuffix             []string                            `hcl:"dns_suffix"`
	ExtraRoutes           map[string]GatewayExtraRoutesValue  `hcl:"extra_routes"`
	ExtraRoutes6          map[string]GatewayExtraRoutes6Value `hcl:"extra_routes6"`
	IdpProfiles           map[string]IdpProfilesValue         `hcl:"idp_profiles"`
	IpConfigs             map[string]IpConfigsValue           `hcl:"ip_configs"`
	Managed               *bool                               `hcl:"managed"`
	MapId                 *string                             `hcl:"map_id"`
	MspId                 *string                             `hcl:"msp_id"`
	Name                  string                              `hcl:"name"`
	Networks              []GatewayNetworksValue              `hcl:"networks"`
	Notes                 *string                             `hcl:"notes"`
	NtpServers            []string                            `hcl:"ntp_servers"`
	OobIpConfig           *GatewayOobIpConfigValue            `hcl:"oob_ip_config"`
	PathPreferences       map[string]PathPreferencesValue     `hcl:"path_preferences"`
	PortConfig            map[string]GatewayPortConfigValue   `hcl:"port_config"`
	PortMirroring         *GatewayPortMirroringValue          `hcl:"port_mirroring"`
	RouterId              *string                             `hcl:"router_id"`
	RoutingPolicies       map[string]RoutingPoliciesValue     `hcl:"routing_policies"`
	ServicePolicies       []ServicePoliciesValue              `hcl:"service_policies"`
	SiteId                string                              `hcl:"site_id"`
	TunnelConfigs         map[string]TunnelConfigsValue       `hcl:"tunnel_configs"`
	TunnelProviderOptions *TunnelProviderOptionsValue         `hcl:"tunnel_provider_options"`
	Vars                  map[string]string                   `hcl:"vars"`
	VrfConfig             *VrfConfigValue                     `hcl:"vrf_config"`
	VrfInstances          map[string]GatewayVrfInstancesValue `hcl:"vrf_instances"`
	X                     *float64                            `hcl:"x"`
	Y                     *float64                            `hcl:"y"`
}

type BgpConfigValue struct {
	AuthKey                *string                   `cty:"auth_key"`
	BfdMinimumInterval     *int64                    `cty:"bfd_minimum_interval"`
	BfdMultiplier          *int64                    `cty:"bfd_multiplier"`
	DisableBfd             *bool                     `cty:"disable_bfd"`
	Export                 *string                   `cty:"export"`
	ExportPolicy           *string                   `cty:"export_policy"`
	ExtendedV4Nexthop      *bool                     `cty:"extended_v4_nexthop"`
	GracefulRestartTime    *int64                    `cty:"graceful_restart_time"`
	HoldTime               *int64                    `cty:"hold_time"`
	Import                 *string                   `cty:"import"`
	ImportPolicy           *string                   `cty:"import_policy"`
	LocalAs                *int64                    `cty:"local_as"`
	NeighborAs             *int64                    `cty:"neighbor_as"`
	Neighbors              map[string]NeighborsValue `cty:"neighbors"`
	Networks               []string                  `cty:"networks"`
	NoReadvertiseToOverlay *bool                     `cty:"no_readvertise_to_overlay"`
	TunnelName             *string                   `cty:"tunnel_name"`
	BgpConfigType          *string                   `cty:"type"`
	Via                    *string                   `cty:"via"`
	VpnName                *string                   `cty:"vpn_name"`
	WanName                *string                   `cty:"wan_name"`
}

type NeighborsValue struct {
	Disabled     *bool   `cty:"disabled"`
	ExportPolicy *string `cty:"export_policy"`
	HoldTime     *int64  `cty:"hold_time"`
	ImportPolicy *string `cty:"import_policy"`
	MultihopTtl  *int64  `cty:"multihop_ttl"`
	NeighborAs   *int64  `cty:"neighbor_as"`
}

type GatewayExtraRoutesValue struct {
	Via string `cty:"via"`
}

type GatewayExtraRoutes6Value struct {
	Via string `cty:"via"`
}

type IdpProfilesValue struct {
	BaseProfile *string           `cty:"base_profile"`
	Id          *string           `cty:"id"`
	Name        *string           `cty:"name"`
	OrgId       *string           `cty:"org_id"`
	Overwrites  []OverwritesValue `cty:"overwrites"`
}

type OverwritesValue struct {
	Action                      *string                           `cty:"action"`
	IpdProfileOverwriteMatching *IpdProfileOverwriteMatchingValue `cty:"matching"`
	Name                        *string                           `cty:"name"`
}

type IpdProfileOverwriteMatchingValue struct {
	AttackName []string `cty:"attack_name"`
	DstSubnet  []string `cty:"dst_subnet"`
	Severity   []string `cty:"severity"`
}

type IpConfigsValue struct {
	Ip            string   `cty:"ip"`
	Netmask       string   `cty:"netmask"`
	SecondaryIps  []string `cty:"secondary_ips"`
	IpConfigsType *string  `cty:"type"`
}

type GatewayNetworksValue struct {
	DisallowMistServices *bool                     `cty:"disallow_mist_services"`
	Gateway              *string                   `cty:"gateway"`
	Gateway6             *string                   `cty:"gateway6"`
	InternalAccess       *InternalAccessValue      `cty:"internal_access"`
	InternetAccess       *InternetAccessValue      `cty:"internet_access"`
	Isolation            *bool                     `cty:"isolation"`
	Multicast            *MulticastValue           `cty:"multicast"`
	Name                 *string                   `cty:"name"`
	RoutedForNetworks    []string                  `cty:"routed_for_networks"`
	Subnet               string                    `cty:"subnet"`
	Subnet6              *string                   `cty:"subnet6"`
	Tenants              map[string]TenantsValue   `cty:"tenants"`
	VlanId               *string                   `cty:"vlan_id"`
	VpnAccess            map[string]VpnAccessValue `cty:"vpn_access"`
}

type InternalAccessValue struct {
	Enabled *bool `cty:"enabled"`
}

type InternetAccessValue struct {
	CreateSimpleServicePolicy    *bool                                        `cty:"create_simple_service_policy"`
	Enabled                      *bool                                        `cty:"enabled"`
	InternetAccessDestinationNat map[string]InternetAccessDestinationNatValue `cty:"destination_nat"`
	InternetAccessStaticNat      map[string]InternetAccessStaticNatValue      `cty:"static_nat"`
	Restricted                   *bool                                        `cty:"restricted"`
}

type InternetAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip"`
	Name       string  `cty:"name"`
	Port       *string `cty:"port"`
	WanName    *string `cty:"wan_name"`
}

type InternetAccessStaticNatValue struct {
	InternalIp string  `cty:"internal_ip"`
	Name       string  `cty:"name"`
	WanName    *string `cty:"wan_name"`
}

type MulticastValue struct {
	DisableIgmp *bool                  `cty:"disable_igmp"`
	Enabled     *bool                  `cty:"enabled"`
	Groups      map[string]GroupsValue `cty:"groups"`
}

type TenantsValue struct {
	Addresses []string `cty:"addresses"`
}

type VpnAccessValue struct {
	AdvertisedSubnet          *string                                 `cty:"advertised_subnet"`
	AllowPing                 *bool                                   `cty:"allow_ping"`
	NatPool                   *string                                 `cty:"nat_pool"`
	NoReadvertiseToLanBgp     *bool                                   `cty:"no_readvertise_to_lan_bgp"`
	NoReadvertiseToLanOspf    *bool                                   `cty:"no_readvertise_to_lan_ospf"`
	NoReadvertiseToOverlay    *bool                                   `cty:"no_readvertise_to_overlay"`
	OtherVrfs                 []string                                `cty:"other_vrfs"`
	Routed                    *bool                                   `cty:"routed"`
	SourceNat                 *SourceNatValue                         `cty:"source_nat"`
	SummarizedSubnet          *string                                 `cty:"summarized_subnet"`
	SummarizedSubnetToLanBgp  *string                                 `cty:"summarized_subnet_to_lan_bgp"`
	SummarizedSubnetToLanOspf *string                                 `cty:"summarized_subnet_to_lan_ospf"`
	VpnAccessDestinationNat   map[string]VpnAccessDestinationNatValue `cty:"destination_nat"`
	VpnAccessStaticNat        map[string]VpnAccessStaticNatValue      `cty:"static_nat"`
}

type SourceNatValue struct {
	ExternalIp *string `cty:"external_ip"`
}

type VpnAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip"`
	Name       *string `cty:"name"`
	Port       *string `cty:"port"`
}

type VpnAccessStaticNatValue struct {
	InternalIp string `cty:"internal_ip"`
	Name       string `cty:"name"`
}

type GatewayOobIpConfigValue struct {
	Gateway              *string     `cty:"gateway"`
	Ip                   *string     `cty:"ip"`
	Netmask              *string     `cty:"netmask"`
	Node1                *Node1Value `cty:"node1"`
	OobIpConfigType      *string     `cty:"type"`
	UseMgmtVrf           *bool       `cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool       `cty:"use_mgmt_vrf_for_host_out"`
	VlanId               *string     `cty:"vlan_id"`
}

type Node1Value struct {
	Gateway              *string `cty:"gateway"`
	Ip                   *string `cty:"ip"`
	Netmask              *string `cty:"netmask"`
	Node1Type            *string `cty:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out"`
	VlanId               *string `cty:"vlan_id"`
}

type PathPreferencesValue struct {
	Paths    []PathsValue `cty:"paths"`
	Strategy *string      `cty:"strategy"`
}

type PathsValue struct {
	Cost           *int64   `cty:"cost"`
	Disabled       *bool    `cty:"disabled"`
	GatewayIp      *string  `cty:"gateway_ip"`
	InternetAccess *bool    `cty:"internet_access"`
	Name           *string  `cty:"name"`
	Networks       []string `cty:"networks"`
	TargetIps      []string `cty:"target_ips"`
	PathsType      *string  `cty:"type"`
	WanName        *string  `cty:"wan_name"`
}

type GatewayPortConfigValue struct {
	AeDisableLacp    *bool                          `cty:"ae_disable_lacp"`
	AeIdx            *string                        `cty:"ae_idx"`
	AeLacpForceUp    *bool                          `cty:"ae_lacp_force_up"`
	Aggregated       *bool                          `cty:"aggregated"`
	Critical         *bool                          `cty:"critical"`
	Description      *string                        `cty:"description"`
	DisableAutoneg   *bool                          `cty:"disable_autoneg"`
	Disabled         *bool                          `cty:"disabled"`
	DslType          *string                        `cty:"dsl_type"`
	DslVci           *int64                         `cty:"dsl_vci"`
	DslVpi           *int64                         `cty:"dsl_vpi"`
	Duplex           *string                        `cty:"duplex"`
	LteApn           *string                        `cty:"lte_apn"`
	LteAuth          *string                        `cty:"lte_auth"`
	LteBackup        *bool                          `cty:"lte_backup"`
	LtePassword      *string                        `cty:"lte_password"`
	LteUsername      *string                        `cty:"lte_username"`
	Mtu              *int64                         `cty:"mtu"`
	Name             *string                        `cty:"name"`
	Networks         []string                       `cty:"networks"`
	OuterVlanId      *int64                         `cty:"outer_vlan_id"`
	PoeDisabled      *bool                          `cty:"poe_disabled"`
	PortIpConfig     *PortIpConfigValue             `cty:"ip_config"`
	PortNetwork      *string                        `cty:"port_network"`
	PreserveDscp     *bool                          `cty:"preserve_dscp"`
	Redundant        *bool                          `cty:"redundant"`
	RethIdx          *int64                         `cty:"reth_idx"`
	RethNode         *string                        `cty:"reth_node"`
	RethNodes        []string                       `cty:"reth_nodes"`
	Speed            *string                        `cty:"speed"`
	SsrNoVirtualMac  *bool                          `cty:"ssr_no_virtual_mac"`
	SvrPortRange     *string                        `cty:"svr_port_range"`
	TrafficShaping   *TrafficShapingValue           `cty:"traffic_shaping"`
	Usage            string                         `cty:"usage"`
	VlanId           *string                        `cty:"vlan_id"`
	VpnPaths         map[string]VpnPathsValue       `cty:"vpn_paths"`
	WanArpPolicer    *string                        `cty:"wan_arp_policer"`
	WanExtIp         *string                        `cty:"wan_ext_ip"`
	WanExtraRoutes   map[string]WanExtraRoutesValue `cty:"wan_extra_routes"`
	WanNetworks      []string                       `cty:"wan_networks"`
	WanProbeOverride *WanProbeOverrideValue         `cty:"wan_probe_override"`
	WanSourceNat     *WanSourceNatValue             `cty:"wan_source_nat"`
	WanType          *string                        `cty:"wan_type"`
}

type PortIpConfigValue struct {
	Dns              []string `cty:"dns"`
	DnsSuffix        []string `cty:"dns_suffix"`
	Gateway          *string  `cty:"gateway"`
	Ip               *string  `cty:"ip"`
	Netmask          *string  `cty:"netmask"`
	Network          *string  `cty:"network"`
	PoserPassword    *string  `cty:"poser_password"`
	PppoeAuth        *string  `cty:"pppoe_auth"`
	PppoeUsername    *string  `cty:"pppoe_username"`
	PortIpConfigType *string  `cty:"type"`
}

type TrafficShapingValue struct {
	ClassPercentages []int64 `cty:"class_percentages"`
	Enabled          *bool   `cty:"enabled"`
	MaxTxKbps        *int64  `cty:"max_tx_kbps"`
}

type VpnPathsValue struct {
	BfdProfile       *string              `cty:"bfd_profile"`
	BfdUseTunnelMode *bool                `cty:"bfd_use_tunnel_mode"`
	LinkName         *string              `cty:"link_name"`
	Preference       *int64               `cty:"preference"`
	Role             *string              `cty:"role"`
	TrafficShaping   *TrafficShapingValue `cty:"traffic_shaping"`
}

type WanExtraRoutesValue struct {
	Via *string `cty:"via"`
}

type WanProbeOverrideValue struct {
	Ips          []string `cty:"ips"`
	ProbeProfile *string  `cty:"probe_profile"`
}

type WanSourceNatValue struct {
	Disabled *bool   `cty:"disabled"`
	NatPool  *string `cty:"nat_pool"`
}

type GatewayPortMirroringValue struct {
	PortMirror *PortMirrorValue `cty:"port_mirror"`
}

type PortMirrorValue struct {
	FamilyType     *string  `cty:"family_type"`
	IngressPortIds []string `cty:"ingress_port_ids"`
	OutputPortId   *string  `cty:"output_port_id"`
	Rate           *int64   `cty:"rate"`
	RunLength      *int64   `cty:"run_length"`
}

type RoutingPoliciesValue struct {
	Terms []TermsValue `cty:"terms"`
}

type TermsValue struct {
	Action                    *ActionValue                    `cty:"action"`
	RoutingPolicyTermMatching *RoutingPolicyTermMatchingValue `cty:"matching"`
}

type ActionValue struct {
	Accept             *bool    `cty:"accept"`
	AddCommunity       []string `cty:"add_community"`
	AddTargetVrfs      []string `cty:"add_target_vrfs"`
	Aggregate          []string `cty:"aggregate"`
	Community          []string `cty:"community"`
	ExcludeAsPath      []string `cty:"exclude_as_path"`
	ExcludeCommunity   []string `cty:"exclude_community"`
	ExportCommunitites []string `cty:"export_communitites"`
	LocalPreference    *string  `cty:"local_preference"`
	PrependAsPath      []string `cty:"prepend_as_path"`
}

type RoutingPolicyTermMatchingValue struct {
	AsPath         []string          `cty:"as_path"`
	Community      []string          `cty:"community"`
	Network        []string          `cty:"network"`
	Prefix         []string          `cty:"prefix"`
	Protocol       []string          `cty:"protocol"`
	RouteExists    *RouteExistsValue `cty:"route_exists"`
	VpnNeighborMac []string          `cty:"vpn_neighbor_mac"`
	VpnPath        []string          `cty:"vpn_path"`
	VpnPathSla     *VpnPathSlaValue  `cty:"vpn_path_sla"`
}

type RouteExistsValue struct {
	Route   *string `cty:"route"`
	VrfName *string `cty:"vrf_name"`
}

type VpnPathSlaValue struct {
	MaxJitter  *int64 `cty:"max_jitter"`
	MaxLatency *int64 `cty:"max_latency"`
	MaxLoss    *int64 `cty:"max_loss"`
}

type ServicePoliciesValue struct {
	Action          *string         `cty:"action"`
	Antivirus       *AntivirusValue `cty:"antivirus"`
	Appqoe          *AppqoeValue    `cty:"appqoe"`
	Ewf             []EwfValue      `cty:"ewf"`
	Idp             *IdpValue       `cty:"idp"`
	LocalRouting    *bool           `cty:"local_routing"`
	Name            *string         `cty:"name"`
	PathPreference  *string         `cty:"path_preference"`
	ServicepolicyId *string         `cty:"servicepolicy_id"`
	Services        []string        `cty:"services"`
	SslProxy        *SslProxyValue  `cty:"ssl_proxy"`
	Tenants         []string        `cty:"tenants"`
}

type AntivirusValue struct {
	AvprofileId *string `cty:"avprofile_id"`
	Enabled     *bool   `cty:"enabled"`
	Profile     *string `cty:"profile"`
}

type AppqoeValue struct {
	Enabled *bool `cty:"enabled"`
}

type EwfValue struct {
	AlertOnly    *bool   `cty:"alert_only"`
	BlockMessage *string `cty:"block_message"`
	Enabled      *bool   `cty:"enabled"`
	Profile      *string `cty:"profile"`
}

type IdpValue struct {
	AlertOnly    *bool   `cty:"alert_only"`
	Enabled      *bool   `cty:"enabled"`
	IdpprofileId *string `cty:"idpprofile_id"`
	Profile      *string `cty:"profile"`
}

type SslProxyValue struct {
	CiphersCategory *string `cty:"ciphers_category"`
	Enabled         *bool   `cty:"enabled"`
}

type TunnelConfigsValue struct {
	AutoProvision  *AutoProvisionValue   `cty:"auto_provision"`
	IkeLifetime    *int64                `cty:"ike_lifetime"`
	IkeMode        *string               `cty:"ike_mode"`
	IkeProposals   []IkeProposalsValue   `cty:"ike_proposals"`
	IpsecLifetime  *int64                `cty:"ipsec_lifetime"`
	IpsecProposals []IpsecProposalsValue `cty:"ipsec_proposals"`
	LocalId        *string               `cty:"local_id"`
	Mode           *string               `cty:"mode"`
	Networks       []string              `cty:"networks"`
	Primary        *PrimaryValue         `cty:"primary"`
	Probe          *ProbeValue           `cty:"probe"`
	Protocol       *string               `cty:"protocol"`
	Provider       string                `cty:"provider"`
	Psk            *string               `cty:"psk"`
	Secondary      *SecondaryValue       `cty:"secondary"`
	Version        *string               `cty:"version"`
}

type AutoProvisionValue struct {
	AutoProvisionPrimary   *AutoProvisionPrimaryValue   `cty:"primary"`
	AutoProvisionSecondary *AutoProvisionSecondaryValue `cty:"secondary"`
	Enable                 *bool                        `cty:"enable"`
	Latlng                 *LatlngValue                 `cty:"latlng"`
	Provider               *string                      `cty:"provider"`
	Region                 *string                      `cty:"region"`
}

type AutoProvisionPrimaryValue struct {
	ProbeIps []string `cty:"probe_ips"`
	WanNames []string `cty:"wan_names"`
}

type AutoProvisionSecondaryValue struct {
	ProbeIps []string `cty:"probe_ips"`
	WanNames []string `cty:"wan_names"`
}

type LatlngValue struct {
	Lat float64 `cty:"lat"`
	Lng float64 `cty:"lng"`
}

type IkeProposalsValue struct {
	AuthAlgo *string `cty:"auth_algo"`
	DhGroup  *string `cty:"dh_group"`
	EncAlgo  *string `cty:"enc_algo"`
}

type IpsecProposalsValue struct {
	AuthAlgo *string `cty:"auth_algo"`
	DhGroup  *string `cty:"dh_group"`
	EncAlgo  *string `cty:"enc_algo"`
}

type PrimaryValue struct {
	Hosts       []string `cty:"hosts"`
	InternalIps []string `cty:"internal_ips"`
	ProbeIps    []string `cty:"probe_ips"`
	RemoteIds   []string `cty:"remote_ids"`
	WanNames    []string `cty:"wan_names"`
}

type ProbeValue struct {
	Interval  *int64  `cty:"interval"`
	Threshold *int64  `cty:"threshold"`
	Timeout   *int64  `cty:"timeout"`
	ProbeType *string `cty:"type"`
}

type SecondaryValue struct {
	Hosts       []string `cty:"hosts"`
	InternalIps []string `cty:"internal_ips"`
	ProbeIps    []string `cty:"probe_ips"`
	RemoteIds   []string `cty:"remote_ids"`
	WanNames    []string `cty:"wan_names"`
}

type TunnelProviderOptionsValue struct {
	Jse     *JseValue     `cty:"jse"`
	Zscaler *ZscalerValue `cty:"zscaler"`
}

type JseValue struct {
	NumUsers *int64  `cty:"num_users"`
	OrgName  *string `cty:"org_name"`
}

type ZscalerValue struct {
	AupBlockInternetUntilAccepted       *bool               `cty:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool               `cty:"aup_enabled"`
	AupForceSslInspection               *bool               `cty:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64              `cty:"aup_timeout_in_days"`
	AuthRequired                        *bool               `cty:"auth_required"`
	CautionEnabled                      *bool               `cty:"caution_enabled"`
	DnBandwidth                         *float64            `cty:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64              `cty:"idle_time_in_minutes"`
	OfwEnabled                          *bool               `cty:"ofw_enabled"`
	SubLocations                        []SubLocationsValue `cty:"sub_locations"`
	SurrogateIp                         *bool               `cty:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool               `cty:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64              `cty:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64            `cty:"up_bandwidth"`
	XffForwardEnabled                   *bool               `cty:"xff_forward_enabled"`
}

type SubLocationsValue struct {
	AupBlockInternetUntilAccepted       *bool    `cty:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool    `cty:"aup_enabled"`
	AupForceSslInspection               *bool    `cty:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64   `cty:"aup_timeout_in_days"`
	AuthRequired                        *bool    `cty:"auth_required"`
	CautionEnabled                      *bool    `cty:"caution_enabled"`
	DnBandwidth                         *float64 `cty:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64   `cty:"idle_time_in_minutes"`
	Name                                *string  `cty:"name"`
	OfwEnabled                          *bool    `cty:"ofw_enabled"`
	SurrogateIp                         *bool    `cty:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool    `cty:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64   `cty:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64 `cty:"up_bandwidth"`
}

type GatewayVrfInstancesValue struct {
	Networks []string `cty:"networks"`
}
