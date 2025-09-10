package provider

type DeviceGatewayModel struct {
	AdditionalConfigCmds  []string                                  `hcl:"additional_config_cmds"`
	BgpConfig             map[string]BgpConfigValue                 `hcl:"bgp_config"`
	DeviceId              string                                    `hcl:"device_id"`
	DhcpdConfig           *DhcpdConfigValue                         `hcl:"dhcpd_config"`
	DnsServers            []string                                  `hcl:"dns_servers"`
	DnsSuffix             []string                                  `hcl:"dns_suffix"`
	ExtraRoutes           map[string]DeviceGatewayExtraRoutesValue  `hcl:"extra_routes"`
	ExtraRoutes6          map[string]DeviceGatewayExtraRoutes6Value `hcl:"extra_routes6"`
	IdpProfiles           map[string]IdpProfilesValue               `hcl:"idp_profiles"`
	IpConfigs             map[string]IpConfigsValue                 `hcl:"ip_configs"`
	Managed               *bool                                     `hcl:"managed"`
	MapId                 *string                                   `hcl:"map_id"`
	MspId                 *string                                   `hcl:"msp_id"`
	Name                  string                                    `hcl:"name"`
	Networks              []DeviceGatewayNetworksValue              `hcl:"networks"`
	Notes                 *string                                   `hcl:"notes"`
	NtpServers            []string                                  `hcl:"ntp_servers"`
	OobIpConfig           *DeviceGatewayOobIpConfigValue            `hcl:"oob_ip_config"`
	PathPreferences       map[string]PathPreferencesValue           `hcl:"path_preferences"`
	PortConfig            map[string]DeviceGatewayPortConfigValue   `hcl:"port_config"`
	PortMirroring         *DeviceGatewayPortMirroringValue          `hcl:"port_mirroring"`
	RouterId              *string                                   `hcl:"router_id"`
	RoutingPolicies       map[string]RoutingPoliciesValue           `hcl:"routing_policies"`
	ServicePolicies       []ServicePoliciesValue                    `hcl:"service_policies"`
	SiteId                string                                    `hcl:"site_id"`
	TunnelConfigs         map[string]TunnelConfigsValue             `hcl:"tunnel_configs"`
	TunnelProviderOptions *TunnelProviderOptionsValue               `hcl:"tunnel_provider_options"`
	Vars                  map[string]string                         `hcl:"vars"`
	VrfConfig             *DeviceGatewayVrfConfigValue              `hcl:"vrf_config"`
	VrfInstances          map[string]DeviceGatewayVrfInstancesValue `hcl:"vrf_instances"`
	X                     *float64                                  `hcl:"x"`
	Y                     *float64                                  `hcl:"y"`
}

type BgpConfigValue struct {
	AuthKey                *string                   `hcl:"auth_key" cty:"auth_key"`
	BfdMinimumInterval     *int64                    `hcl:"bfd_minimum_interval" cty:"bfd_minimum_interval"`
	BfdMultiplier          *int64                    `hcl:"bfd_multiplier" cty:"bfd_multiplier"`
	DisableBfd             *bool                     `hcl:"disable_bfd" cty:"disable_bfd"`
	Export                 *string                   `hcl:"export" cty:"export"`
	ExportPolicy           *string                   `hcl:"export_policy" cty:"export_policy"`
	ExtendedV4Nexthop      *bool                     `hcl:"extended_v4_nexthop" cty:"extended_v4_nexthop"`
	GracefulRestartTime    *int64                    `hcl:"graceful_restart_time" cty:"graceful_restart_time"`
	HoldTime               *int64                    `hcl:"hold_time" cty:"hold_time"`
	Import                 *string                   `hcl:"import" cty:"import"`
	ImportPolicy           *string                   `hcl:"import_policy" cty:"import_policy"`
	LocalAs                *string                   `hcl:"local_as" cty:"local_as"`
	NeighborAs             *string                   `hcl:"neighbor_as" cty:"neighbor_as"`
	Neighbors              map[string]NeighborsValue `hcl:"neighbors" cty:"neighbors"`
	Networks               []string                  `hcl:"networks" cty:"networks"`
	NoPrivateAs            *bool                     `hcl:"no_private_as" cty:"no_private_as"`
	NoReadvertiseToOverlay *bool                     `hcl:"no_readvertise_to_overlay" cty:"no_readvertise_to_overlay"`
	TunnelName             *string                   `hcl:"tunnel_name" cty:"tunnel_name"`
	BgpConfigType          *string                   `hcl:"type" cty:"type"`
	Via                    string                    `hcl:"via" cty:"via"`
	VpnName                *string                   `hcl:"vpn_name" cty:"vpn_name"`
	WanName                *string                   `hcl:"wan_name" cty:"wan_name"`
}

type NeighborsValue struct {
	Disabled     *bool   `hcl:"disabled" cty:"disabled"`
	ExportPolicy *string `hcl:"export_policy" cty:"export_policy"`
	HoldTime     *int64  `hcl:"hold_time" cty:"hold_time"`
	ImportPolicy *string `hcl:"import_policy" cty:"import_policy"`
	MultihopTtl  *int64  `hcl:"multihop_ttl" cty:"multihop_ttl"`
	NeighborAs   string  `hcl:"neighbor_as" cty:"neighbor_as"`
}

type DeviceGatewayExtraRoutesValue struct {
	Via string `hcl:"via" cty:"via"`
}

type DeviceGatewayExtraRoutes6Value struct {
	Via string `hcl:"via" cty:"via"`
}

type IdpProfilesValue struct {
	BaseProfile *string           `hcl:"base_profile" cty:"base_profile"`
	Id          *string           `hcl:"id" cty:"id"`
	Name        *string           `hcl:"name" cty:"name"`
	OrgId       *string           `hcl:"org_id" cty:"org_id"`
	Overwrites  []OverwritesValue `hcl:"overwrites" cty:"overwrites"`
}

type OverwritesValue struct {
	Action                      *string                           `hcl:"action" cty:"action"`
	IpdProfileOverwriteMatching *IpdProfileOverwriteMatchingValue `hcl:"matching" cty:"matching"`
	Name                        *string                           `hcl:"name" cty:"name"`
}

type IpdProfileOverwriteMatchingValue struct {
	AttackName []string `hcl:"attack_name" cty:"attack_name"`
	DstSubnet  []string `hcl:"dst_subnet" cty:"dst_subnet"`
	Severity   []string `hcl:"severity" cty:"severity"`
}

type IpConfigsValue struct {
	Ip            *string  `hcl:"ip" cty:"ip"`
	Ip6           *string  `hcl:"ip6" cty:"ip6"`
	Netmask       *string  `hcl:"netmask" cty:"netmask"`
	Netmask6      *string  `hcl:"netmask6" cty:"netmask6"`
	SecondaryIps  []string `hcl:"secondary_ips" cty:"secondary_ips"`
	IpConfigsType *string  `hcl:"type" cty:"type"`
	Type6         *string  `hcl:"type6" cty:"type6"`
}

type DeviceGatewayNetworksValue struct {
	DisallowMistServices *bool                     `hcl:"disallow_mist_services" cty:"disallow_mist_services"`
	Gateway              *string                   `hcl:"gateway" cty:"gateway"`
	Gateway6             *string                   `hcl:"gateway6" cty:"gateway6"`
	InternalAccess       *InternalAccessValue      `hcl:"internal_access" cty:"internal_access"`
	InternetAccess       *InternetAccessValue      `hcl:"internet_access" cty:"internet_access"`
	Isolation            *bool                     `hcl:"isolation" cty:"isolation"`
	Multicast            *MulticastValue           `hcl:"multicast" cty:"multicast"`
	Name                 *string                   `hcl:"name" cty:"name"`
	RoutedForNetworks    []string                  `hcl:"routed_for_networks" cty:"routed_for_networks"`
	Subnet               string                    `hcl:"subnet" cty:"subnet"`
	Subnet6              *string                   `hcl:"subnet6" cty:"subnet6"`
	Tenants              map[string]TenantsValue   `hcl:"tenants" cty:"tenants"`
	VlanId               *string                   `hcl:"vlan_id" cty:"vlan_id"`
	VpnAccess            map[string]VpnAccessValue `hcl:"vpn_access" cty:"vpn_access"`
}

type InternalAccessValue struct {
	Enabled *bool `hcl:"enabled" cty:"enabled"`
}

type InternetAccessValue struct {
	CreateSimpleServicePolicy    *bool                                        `hcl:"create_simple_service_policy" cty:"create_simple_service_policy"`
	Enabled                      *bool                                        `hcl:"enabled" cty:"enabled"`
	InternetAccessDestinationNat map[string]InternetAccessDestinationNatValue `hcl:"destination_nat" cty:"destination_nat"`
	InternetAccessStaticNat      map[string]InternetAccessStaticNatValue      `hcl:"static_nat" cty:"static_nat"`
	Restricted                   *bool                                        `hcl:"restricted" cty:"restricted"`
}

type InternetAccessDestinationNatValue struct {
	InternalIp *string `hcl:"internal_ip" cty:"internal_ip"`
	Name       string  `hcl:"name" cty:"name"`
	Port       *string `hcl:"port" cty:"port"`
	WanName    *string `hcl:"wan_name" cty:"wan_name"`
}

type InternetAccessStaticNatValue struct {
	InternalIp string  `hcl:"internal_ip" cty:"internal_ip"`
	Name       string  `hcl:"name" cty:"name"`
	WanName    *string `hcl:"wan_name" cty:"wan_name"`
}

type MulticastValue struct {
	DisableIgmp *bool                               `hcl:"disable_igmp" cty:"disable_igmp"`
	Enabled     *bool                               `hcl:"enabled" cty:"enabled"`
	Groups      map[string]DeviceGatewayGroupsValue `hcl:"groups" cty:"groups"`
}

type DeviceGatewayGroupsValue struct {
	RpIp *string `hcl:"rp_ip" cty:"rp_ip"`
}

type TenantsValue struct {
	Addresses []string `hcl:"addresses" cty:"addresses"`
}

type VpnAccessValue struct {
	AdvertisedSubnet          *string                                 `hcl:"advertised_subnet" cty:"advertised_subnet"`
	AllowPing                 *bool                                   `hcl:"allow_ping" cty:"allow_ping"`
	NatPool                   *string                                 `hcl:"nat_pool" cty:"nat_pool"`
	NoReadvertiseToLanBgp     *bool                                   `hcl:"no_readvertise_to_lan_bgp" cty:"no_readvertise_to_lan_bgp"`
	NoReadvertiseToLanOspf    *bool                                   `hcl:"no_readvertise_to_lan_ospf" cty:"no_readvertise_to_lan_ospf"`
	NoReadvertiseToOverlay    *bool                                   `hcl:"no_readvertise_to_overlay" cty:"no_readvertise_to_overlay"`
	OtherVrfs                 []string                                `hcl:"other_vrfs" cty:"other_vrfs"`
	Routed                    *bool                                   `hcl:"routed" cty:"routed"`
	SourceNat                 *SourceNatValue                         `hcl:"source_nat" cty:"source_nat"`
	SummarizedSubnet          *string                                 `hcl:"summarized_subnet" cty:"summarized_subnet"`
	SummarizedSubnetToLanBgp  *string                                 `hcl:"summarized_subnet_to_lan_bgp" cty:"summarized_subnet_to_lan_bgp"`
	SummarizedSubnetToLanOspf *string                                 `hcl:"summarized_subnet_to_lan_ospf" cty:"summarized_subnet_to_lan_ospf"`
	VpnAccessDestinationNat   map[string]VpnAccessDestinationNatValue `hcl:"destination_nat" cty:"destination_nat"`
	VpnAccessStaticNat        map[string]VpnAccessStaticNatValue      `hcl:"static_nat" cty:"static_nat"`
}

type SourceNatValue struct {
	ExternalIp *string `hcl:"external_ip" cty:"external_ip"`
}

type VpnAccessDestinationNatValue struct {
	InternalIp *string `hcl:"internal_ip" cty:"internal_ip"`
	Name       *string `hcl:"name" cty:"name"`
	Port       *string `hcl:"port" cty:"port"`
}

type VpnAccessStaticNatValue struct {
	InternalIp string `hcl:"internal_ip" cty:"internal_ip"`
	Name       string `hcl:"name" cty:"name"`
}

type DeviceGatewayOobIpConfigValue struct {
	Gateway              *string     `hcl:"gateway" cty:"gateway"`
	Ip                   *string     `hcl:"ip" cty:"ip"`
	Netmask              *string     `hcl:"netmask" cty:"netmask"`
	Node1                *Node1Value `hcl:"node1" cty:"node1"`
	OobIpConfigType      *string     `hcl:"type" cty:"type"`
	UseMgmtVrf           *bool       `hcl:"use_mgmt_vrf" cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool       `hcl:"use_mgmt_vrf_for_host_out" cty:"use_mgmt_vrf_for_host_out"`
	VlanId               *string     `hcl:"vlan_id" cty:"vlan_id"`
}

type Node1Value struct {
	Gateway              *string `hcl:"gateway" cty:"gateway"`
	Ip                   *string `hcl:"ip" cty:"ip"`
	Netmask              *string `hcl:"netmask" cty:"netmask"`
	Node1Type            *string `hcl:"type" cty:"type"`
	UseMgmtVrf           *bool   `hcl:"use_mgmt_vrf" cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `hcl:"use_mgmt_vrf_for_host_out" cty:"use_mgmt_vrf_for_host_out"`
	VlanId               *string `hcl:"vlan_id" cty:"vlan_id"`
}

type PathPreferencesValue struct {
	Paths    []PathsValue `hcl:"paths" cty:"paths"`
	Strategy *string      `hcl:"strategy" cty:"strategy"`
}

type PathsValue struct {
	Cost           *int64   `hcl:"cost" cty:"cost"`
	Disabled       *bool    `hcl:"disabled" cty:"disabled"`
	GatewayIp      *string  `hcl:"gateway_ip" cty:"gateway_ip"`
	InternetAccess *bool    `hcl:"internet_access" cty:"internet_access"`
	Name           *string  `hcl:"name" cty:"name"`
	Networks       []string `hcl:"networks" cty:"networks"`
	TargetIps      []string `hcl:"target_ips" cty:"target_ips"`
	PathsType      string   `hcl:"type" cty:"type"`
	WanName        *string  `hcl:"wan_name" cty:"wan_name"`
}

type DeviceGatewayPortConfigValue struct {
	AeDisableLacp       *bool                           `hcl:"ae_disable_lacp" cty:"ae_disable_lacp"`
	AeIdx               *string                         `hcl:"ae_idx" cty:"ae_idx"`
	AeLacpForceUp       *bool                           `hcl:"ae_lacp_force_up" cty:"ae_lacp_force_up"`
	Aggregated          *bool                           `hcl:"aggregated" cty:"aggregated"`
	Critical            *bool                           `hcl:"critical" cty:"critical"`
	Description         *string                         `hcl:"description" cty:"description"`
	DisableAutoneg      *bool                           `hcl:"disable_autoneg" cty:"disable_autoneg"`
	Disabled            *bool                           `hcl:"disabled" cty:"disabled"`
	DslType             *string                         `hcl:"dsl_type" cty:"dsl_type"`
	DslVci              *int64                          `hcl:"dsl_vci" cty:"dsl_vci"`
	DslVpi              *int64                          `hcl:"dsl_vpi" cty:"dsl_vpi"`
	Duplex              *string                         `hcl:"duplex" cty:"duplex"`
	LteApn              *string                         `hcl:"lte_apn" cty:"lte_apn"`
	LteAuth             *string                         `hcl:"lte_auth" cty:"lte_auth"`
	LteBackup           *bool                           `hcl:"lte_backup" cty:"lte_backup"`
	LtePassword         *string                         `hcl:"lte_password" cty:"lte_password"`
	LteUsername         *string                         `hcl:"lte_username" cty:"lte_username"`
	Mtu                 *int64                          `hcl:"mtu" cty:"mtu"`
	Name                *string                         `hcl:"name" cty:"name"`
	Networks            []string                        `hcl:"networks" cty:"networks"`
	OuterVlanId         *int64                          `hcl:"outer_vlan_id" cty:"outer_vlan_id"`
	PoeDisabled         *bool                           `hcl:"poe_disabled" cty:"poe_disabled"`
	PortIpConfig        *PortIpConfigValue              `hcl:"ip_config" cty:"ip_config"`
	PortNetwork         *string                         `hcl:"port_network" cty:"port_network"`
	PreserveDscp        *bool                           `hcl:"preserve_dscp" cty:"preserve_dscp"`
	Redundant           *bool                           `hcl:"redundant" cty:"redundant"`
	RedundantGroup      *int64                          `hcl:"redundant_group" cty:"redundant_group"`
	RethIdx             *string                         `hcl:"reth_idx" cty:"reth_idx"`
	RethNode            *string                         `hcl:"reth_node" cty:"reth_node"`
	RethNodes           []string                        `hcl:"reth_nodes" cty:"reth_nodes"`
	Speed               *string                         `hcl:"speed" cty:"speed"`
	SsrNoVirtualMac     *bool                           `hcl:"ssr_no_virtual_mac" cty:"ssr_no_virtual_mac"`
	SvrPortRange        *string                         `hcl:"svr_port_range" cty:"svr_port_range"`
	TrafficShaping      *TrafficShapingValue            `hcl:"traffic_shaping" cty:"traffic_shaping"`
	Usage               string                          `hcl:"usage" cty:"usage"`
	VlanId              *string                         `hcl:"vlan_id" cty:"vlan_id"`
	VpnPaths            map[string]VpnPathsValue        `hcl:"vpn_paths" cty:"vpn_paths"`
	WanArpPolicer       *string                         `hcl:"wan_arp_policer" cty:"wan_arp_policer"`
	WanDisableSpeedtest *bool                           `hcl:"wan_disable_speedtest" cty:"wan_disable_speedtest"`
	WanExtIp            *string                         `hcl:"wan_ext_ip" cty:"wan_ext_ip"`
	WanExtraRoutes      map[string]WanExtraRoutesValue  `hcl:"wan_extra_routes" cty:"wan_extra_routes"`
	WanExtraRoutes6     map[string]WanExtraRoutes6Value `hcl:"wan_extra_routes6" cty:"wan_extra_routes6"`
	WanNetworks         []string                        `hcl:"wan_networks" cty:"wan_networks"`
	WanProbeOverride    *WanProbeOverrideValue          `hcl:"wan_probe_override" cty:"wan_probe_override"`
	WanSourceNat        *WanSourceNatValue              `hcl:"wan_source_nat" cty:"wan_source_nat"`
	WanType             *string                         `hcl:"wan_type" cty:"wan_type"`
}

type PortIpConfigValue struct {
	Dns              []string `hcl:"dns" cty:"dns"`
	DnsSuffix        []string `hcl:"dns_suffix" cty:"dns_suffix"`
	Gateway          *string  `hcl:"gateway" cty:"gateway"`
	Gateway6         *string  `hcl:"gateway6" cty:"gateway6"`
	Ip               *string  `hcl:"ip" cty:"ip"`
	Ip6              *string  `hcl:"ip6" cty:"ip6"`
	Netmask          *string  `hcl:"netmask" cty:"netmask"`
	Netmask6         *string  `hcl:"netmask6" cty:"netmask6"`
	Network          *string  `hcl:"network" cty:"network"`
	PoserPassword    *string  `hcl:"poser_password" cty:"poser_password"`
	PppoeAuth        *string  `hcl:"pppoe_auth" cty:"pppoe_auth"`
	PppoeUsername    *string  `hcl:"pppoe_username" cty:"pppoe_username"`
	PortIpConfigType *string  `hcl:"type" cty:"type"`
	Type6            *string  `hcl:"type6" cty:"type6"`
}

type TrafficShapingValue struct {
	ClassPercentages []int64 `hcl:"class_percentages" cty:"class_percentages"`
	Enabled          *bool   `hcl:"enabled" cty:"enabled"`
	MaxTxKbps        *int64  `hcl:"max_tx_kbps" cty:"max_tx_kbps"`
}

type VpnPathsValue struct {
	BfdProfile       *string              `hcl:"bfd_profile" cty:"bfd_profile"`
	BfdUseTunnelMode *bool                `hcl:"bfd_use_tunnel_mode" cty:"bfd_use_tunnel_mode"`
	Preference       *int64               `hcl:"preference" cty:"preference"`
	Role             *string              `hcl:"role" cty:"role"`
	TrafficShaping   *TrafficShapingValue `hcl:"traffic_shaping" cty:"traffic_shaping"`
}

type WanExtraRoutesValue struct {
	Via *string `hcl:"via" cty:"via"`
}

type WanExtraRoutes6Value struct {
	Via *string `hcl:"via" cty:"via"`
}

type WanProbeOverrideValue struct {
	Ip6s         []string `hcl:"ip6s" cty:"ip6s"`
	Ips          []string `hcl:"ips" cty:"ips"`
	ProbeProfile *string  `hcl:"probe_profile" cty:"probe_profile"`
}

type WanSourceNatValue struct {
	Disabled *bool   `hcl:"disabled" cty:"disabled"`
	NatPool  *string `hcl:"nat_pool" cty:"nat_pool"`
}

type DeviceGatewayPortMirroringValue struct {
	PortMirror *PortMirrorValue `hcl:"port_mirror" cty:"port_mirror"`
}

type PortMirrorValue struct {
	FamilyType     *string  `hcl:"family_type" cty:"family_type"`
	IngressPortIds []string `hcl:"ingress_port_ids" cty:"ingress_port_ids"`
	OutputPortId   *string  `hcl:"output_port_id" cty:"output_port_id"`
	Rate           *int64   `hcl:"rate" cty:"rate"`
	RunLength      *int64   `hcl:"run_length" cty:"run_length"`
}

type RoutingPoliciesValue struct {
	Terms []TermsValue `hcl:"terms" cty:"terms"`
}

type TermsValue struct {
	Actions                   *DeviceGatewayActionsValue      `hcl:"actions" cty:"actions"`
	RoutingPolicyTermMatching *RoutingPolicyTermMatchingValue `hcl:"matching" cty:"matching"`
}

type DeviceGatewayActionsValue struct {
	Accept            *bool    `hcl:"accept" cty:"accept"`
	AddCommunity      []string `hcl:"add_community" cty:"add_community"`
	AddTargetVrfs     []string `hcl:"add_target_vrfs" cty:"add_target_vrfs"`
	Community         []string `hcl:"community" cty:"community"`
	ExcludeAsPath     []string `hcl:"exclude_as_path" cty:"exclude_as_path"`
	ExcludeCommunity  []string `hcl:"exclude_community" cty:"exclude_community"`
	ExportCommunities []string `hcl:"export_communities" cty:"export_communities"`
	LocalPreference   *string  `hcl:"local_preference" cty:"local_preference"`
	PrependAsPath     []string `hcl:"prepend_as_path" cty:"prepend_as_path"`
}

type RoutingPolicyTermMatchingValue struct {
	AsPath         []string          `hcl:"as_path" cty:"as_path"`
	Community      []string          `hcl:"community" cty:"community"`
	Network        []string          `hcl:"network" cty:"network"`
	Prefix         []string          `hcl:"prefix" cty:"prefix"`
	Protocol       []string          `hcl:"protocol" cty:"protocol"`
	RouteExists    *RouteExistsValue `hcl:"route_exists" cty:"route_exists"`
	VpnNeighborMac []string          `hcl:"vpn_neighbor_mac" cty:"vpn_neighbor_mac"`
	VpnPath        []string          `hcl:"vpn_path" cty:"vpn_path"`
	VpnPathSla     *VpnPathSlaValue  `hcl:"vpn_path_sla" cty:"vpn_path_sla"`
}

type RouteExistsValue struct {
	Route   *string `hcl:"route" cty:"route"`
	VrfName *string `hcl:"vrf_name" cty:"vrf_name"`
}

type VpnPathSlaValue struct {
	MaxJitter  *int64 `hcl:"max_jitter" cty:"max_jitter"`
	MaxLatency *int64 `hcl:"max_latency" cty:"max_latency"`
	MaxLoss    *int64 `hcl:"max_loss" cty:"max_loss"`
}

type ServicePoliciesValue struct {
	Action          *string         `hcl:"action" cty:"action"`
	Antivirus       *AntivirusValue `hcl:"antivirus" cty:"antivirus"`
	Appqoe          *AppqoeValue    `hcl:"appqoe" cty:"appqoe"`
	Ewf             []EwfValue      `hcl:"ewf" cty:"ewf"`
	Idp             *IdpValue       `hcl:"idp" cty:"idp"`
	LocalRouting    *bool           `hcl:"local_routing" cty:"local_routing"`
	Name            *string         `hcl:"name" cty:"name"`
	PathPreference  *string         `hcl:"path_preference" cty:"path_preference"`
	ServicepolicyId *string         `hcl:"servicepolicy_id" cty:"servicepolicy_id"`
	Services        []string        `hcl:"services" cty:"services"`
	SslProxy        *SslProxyValue  `hcl:"ssl_proxy" cty:"ssl_proxy"`
	Tenants         []string        `hcl:"tenants" cty:"tenants"`
}

type AntivirusValue struct {
	AvprofileId *string `hcl:"avprofile_id" cty:"avprofile_id"`
	Enabled     *bool   `hcl:"enabled" cty:"enabled"`
	Profile     *string `hcl:"profile" cty:"profile"`
}

type AppqoeValue struct {
	Enabled *bool `hcl:"enabled" cty:"enabled"`
}

type EwfValue struct {
	AlertOnly    *bool   `hcl:"alert_only" cty:"alert_only"`
	BlockMessage *string `hcl:"block_message" cty:"block_message"`
	Enabled      *bool   `hcl:"enabled" cty:"enabled"`
	Profile      *string `hcl:"profile" cty:"profile"`
}

type IdpValue struct {
	AlertOnly    *bool   `hcl:"alert_only" cty:"alert_only"`
	Enabled      *bool   `hcl:"enabled" cty:"enabled"`
	IdpprofileId *string `hcl:"idpprofile_id" cty:"idpprofile_id"`
	Profile      *string `hcl:"profile" cty:"profile"`
}

type SslProxyValue struct {
	CiphersCategory *string `hcl:"ciphers_category" cty:"ciphers_category"`
	Enabled         *bool   `hcl:"enabled" cty:"enabled"`
}

type TunnelConfigsValue struct {
	AutoProvision  *AutoProvisionValue   `hcl:"auto_provision" cty:"auto_provision"`
	IkeLifetime    *int64                `hcl:"ike_lifetime" cty:"ike_lifetime"`
	IkeMode        *string               `hcl:"ike_mode" cty:"ike_mode"`
	IkeProposals   []IkeProposalsValue   `hcl:"ike_proposals" cty:"ike_proposals"`
	IpsecLifetime  *int64                `hcl:"ipsec_lifetime" cty:"ipsec_lifetime"`
	IpsecProposals []IpsecProposalsValue `hcl:"ipsec_proposals" cty:"ipsec_proposals"`
	LocalId        *string               `hcl:"local_id" cty:"local_id"`
	LocalSubnets   []string              `hcl:"local_subnets" cty:"local_subnets"`
	Mode           *string               `hcl:"mode" cty:"mode"`
	Networks       []string              `hcl:"networks" cty:"networks"`
	Primary        *PrimaryValue         `hcl:"primary" cty:"primary"`
	Probe          *ProbeValue           `hcl:"probe" cty:"probe"`
	Protocol       *string               `hcl:"protocol" cty:"protocol"`
	Provider       string                `hcl:"provider" cty:"provider"`
	Psk            *string               `hcl:"psk" cty:"psk"`
	RemoteSubnets  []string              `hcl:"remote_subnets" cty:"remote_subnets"`
	Secondary      *SecondaryValue       `hcl:"secondary" cty:"secondary"`
	Version        *string               `hcl:"version" cty:"version"`
}

type AutoProvisionValue struct {
	AutoProvisionPrimary   *AutoProvisionPrimaryValue   `hcl:"primary" cty:"primary"`
	AutoProvisionSecondary *AutoProvisionSecondaryValue `hcl:"secondary" cty:"secondary"`
	Enabled                *bool                        `hcl:"enabled" cty:"enabled"`
	Latlng                 *LatlngValue                 `hcl:"latlng" cty:"latlng"`
	Provider               *string                      `hcl:"provider" cty:"provider"`
	Region                 *string                      `hcl:"region" cty:"region"`
	ServiceConnection      *string                      `hcl:"service_connection" cty:"service_connection"`
}

type AutoProvisionPrimaryValue struct {
	ProbeIps []string `hcl:"probe_ips" cty:"probe_ips"`
	WanNames []string `hcl:"wan_names" cty:"wan_names"`
}

type AutoProvisionSecondaryValue struct {
	ProbeIps []string `hcl:"probe_ips" cty:"probe_ips"`
	WanNames []string `hcl:"wan_names" cty:"wan_names"`
}

type LatlngValue struct {
	Lat float64 `hcl:"lat" cty:"lat"`
	Lng float64 `hcl:"lng" cty:"lng"`
}

type IkeProposalsValue struct {
	AuthAlgo *string `hcl:"auth_algo" cty:"auth_algo"`
	DhGroup  *string `hcl:"dh_group" cty:"dh_group"`
	EncAlgo  *string `hcl:"enc_algo" cty:"enc_algo"`
}

type IpsecProposalsValue struct {
	AuthAlgo *string `hcl:"auth_algo" cty:"auth_algo"`
	DhGroup  *string `hcl:"dh_group" cty:"dh_group"`
	EncAlgo  *string `hcl:"enc_algo" cty:"enc_algo"`
}

type PrimaryValue struct {
	Hosts       []string `hcl:"hosts" cty:"hosts"`
	InternalIps []string `hcl:"internal_ips" cty:"internal_ips"`
	ProbeIps    []string `hcl:"probe_ips" cty:"probe_ips"`
	RemoteIds   []string `hcl:"remote_ids" cty:"remote_ids"`
	WanNames    []string `hcl:"wan_names" cty:"wan_names"`
}

type ProbeValue struct {
	Interval  *int64  `hcl:"interval" cty:"interval"`
	Threshold *int64  `hcl:"threshold" cty:"threshold"`
	Timeout   *int64  `hcl:"timeout" cty:"timeout"`
	ProbeType *string `hcl:"type" cty:"type"`
}

type SecondaryValue struct {
	Hosts       []string `hcl:"hosts" cty:"hosts"`
	InternalIps []string `hcl:"internal_ips" cty:"internal_ips"`
	ProbeIps    []string `hcl:"probe_ips" cty:"probe_ips"`
	RemoteIds   []string `hcl:"remote_ids" cty:"remote_ids"`
	WanNames    []string `hcl:"wan_names" cty:"wan_names"`
}

type TunnelProviderOptionsValue struct {
	Jse     *JseValue     `hcl:"jse" cty:"jse"`
	Prisma  *PrismaValue  `hcl:"prisma" cty:"prisma"`
	Zscaler *ZscalerValue `hcl:"zscaler" cty:"zscaler"`
}

type JseValue struct {
	NumUsers *int64  `hcl:"num_users" cty:"num_users"`
	OrgName  *string `hcl:"org_name" cty:"org_name"`
}

type PrismaValue struct {
	ServiceAccountName *string `hcl:"service_account_name" cty:"service_account_name"`
}

type ZscalerValue struct {
	AupBlockInternetUntilAccepted       *bool               `hcl:"aup_block_internet_until_accepted" cty:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool               `hcl:"aup_enabled" cty:"aup_enabled"`
	AupForceSslInspection               *bool               `hcl:"aup_force_ssl_inspection" cty:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64              `hcl:"aup_timeout_in_days" cty:"aup_timeout_in_days"`
	AuthRequired                        *bool               `hcl:"auth_required" cty:"auth_required"`
	CautionEnabled                      *bool               `hcl:"caution_enabled" cty:"caution_enabled"`
	DnBandwidth                         *float64            `hcl:"dn_bandwidth" cty:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64              `hcl:"idle_time_in_minutes" cty:"idle_time_in_minutes"`
	OfwEnabled                          *bool               `hcl:"ofw_enabled" cty:"ofw_enabled"`
	SubLocations                        []SubLocationsValue `hcl:"sub_locations" cty:"sub_locations"`
	SurrogateIp                         *bool               `hcl:"surrogate_ip" cty:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool               `hcl:"surrogate_ip_enforced_for_known_browsers" cty:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64              `hcl:"surrogate_refresh_time_in_minutes" cty:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64            `hcl:"up_bandwidth" cty:"up_bandwidth"`
	XffForwardEnabled                   *bool               `hcl:"xff_forward_enabled" cty:"xff_forward_enabled"`
}

type SubLocationsValue struct {
	AupBlockInternetUntilAccepted       *bool    `hcl:"aup_block_internet_until_accepted" cty:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool    `hcl:"aup_enabled" cty:"aup_enabled"`
	AupForceSslInspection               *bool    `hcl:"aup_force_ssl_inspection" cty:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64   `hcl:"aup_timeout_in_days" cty:"aup_timeout_in_days"`
	AuthRequired                        *bool    `hcl:"auth_required" cty:"auth_required"`
	CautionEnabled                      *bool    `hcl:"caution_enabled" cty:"caution_enabled"`
	DnBandwidth                         *float64 `hcl:"dn_bandwidth" cty:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64   `hcl:"idle_time_in_minutes" cty:"idle_time_in_minutes"`
	Name                                *string  `hcl:"name" cty:"name"`
	OfwEnabled                          *bool    `hcl:"ofw_enabled" cty:"ofw_enabled"`
	SurrogateIp                         *bool    `hcl:"surrogate_ip" cty:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool    `hcl:"surrogate_ip_enforced_for_known_browsers" cty:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64   `hcl:"surrogate_refresh_time_in_minutes" cty:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64 `hcl:"up_bandwidth" cty:"up_bandwidth"`
}

type DeviceGatewayVrfConfigValue struct {
	Enabled *bool `hcl:"enabled" cty:"enabled"`
}

type DeviceGatewayVrfInstancesValue struct {
	Networks []string `hcl:"networks" cty:"networks"`
}
