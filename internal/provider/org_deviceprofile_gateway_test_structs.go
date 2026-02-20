package provider

type OrgDeviceprofileGatewayModel struct {
	AdditionalConfigCmds    []string                                               `hcl:"additional_config_cmds"`
	BgpConfig               map[string]OrgDeviceprofileGatewayBgpConfigValue       `hcl:"bgp_config"`
	DhcpdConfig             *OrgDeviceprofileGatewayDhcpdConfigValue               `hcl:"dhcpd_config"`
	DnsOverride             *bool                                                  `hcl:"dns_override"`
	DnsServers              []string                                               `hcl:"dns_servers"`
	DnsSuffix               []string                                               `hcl:"dns_suffix"`
	ExtraRoutes             map[string]OrgDeviceprofileGatewayExtraRoutesValue     `hcl:"extra_routes"`
	ExtraRoutes6            map[string]OrgDeviceprofileGatewayExtraRoutes6Value    `hcl:"extra_routes6"`
	IdpProfiles             map[string]OrgDeviceprofileGatewayIdpProfilesValue     `hcl:"idp_profiles"`
	IpConfigs               map[string]OrgDeviceprofileGatewayIpConfigsValue       `hcl:"ip_configs"`
	Name                    string                                                 `hcl:"name"`
	Networks                []OrgDeviceprofileGatewayNetworksValue                 `hcl:"networks"`
	NtpOverride             *bool                                                  `hcl:"ntp_override"`
	NtpServers              []string                                               `hcl:"ntp_servers"`
	OobIpConfig             *OrgDeviceprofileGatewayOobIpConfigValue               `hcl:"oob_ip_config"`
	OrgId                   string                                                 `hcl:"org_id"`
	PathPreferences         map[string]OrgDeviceprofileGatewayPathPreferencesValue `hcl:"path_preferences"`
	PortConfig              map[string]OrgDeviceprofileGatewayPortConfigValue      `hcl:"port_config"`
	RouterId                *string                                                `hcl:"router_id"`
	RoutingPolicies         map[string]OrgDeviceprofileGatewayRoutingPoliciesValue `hcl:"routing_policies"`
	ServicePolicies         []OrgDeviceprofileGatewayServicePoliciesValue          `hcl:"service_policies"`
	SsrAdditionalConfigCmds []string                                               `hcl:"ssr_additional_config_cmds"`
	TunnelConfigs           map[string]OrgDeviceprofileGatewayTunnelConfigsValue   `hcl:"tunnel_configs"`
	TunnelProviderOptions   *OrgDeviceprofileGatewayTunnelProviderOptionsValue     `hcl:"tunnel_provider_options"`
	UrlFilteringDenyMsg     *string                                                `hcl:"url_filtering_deny_msg"`
	VrfConfig               *OrgDeviceprofileGatewayVrfConfigValue                 `hcl:"vrf_config"`
	VrfInstances            map[string]OrgDeviceprofileGatewayVrfInstancesValue    `hcl:"vrf_instances"`
}

type OrgDeviceprofileGatewayBgpConfigValue struct {
	AuthKey                *string                                          `cty:"auth_key" hcl:"auth_key"`
	BfdMinimumInterval     *int64                                           `cty:"bfd_minimum_interval" hcl:"bfd_minimum_interval"`
	BfdMultiplier          *int64                                           `cty:"bfd_multiplier" hcl:"bfd_multiplier"`
	DisableBfd             *bool                                            `cty:"disable_bfd" hcl:"disable_bfd"`
	Export                 *string                                          `cty:"export" hcl:"export"`
	ExportPolicy           *string                                          `cty:"export_policy" hcl:"export_policy"`
	ExtendedV4Nexthop      *bool                                            `cty:"extended_v4_nexthop" hcl:"extended_v4_nexthop"`
	GracefulRestartTime    *int64                                           `cty:"graceful_restart_time" hcl:"graceful_restart_time"`
	HoldTime               *int64                                           `cty:"hold_time" hcl:"hold_time"`
	Import                 *string                                          `cty:"import" hcl:"import"`
	ImportPolicy           *string                                          `cty:"import_policy" hcl:"import_policy"`
	LocalAs                *string                                          `cty:"local_as" hcl:"local_as"`
	NeighborAs             *string                                          `cty:"neighbor_as" hcl:"neighbor_as"`
	Neighbors              map[string]OrgDeviceprofileGatewayNeighborsValue `cty:"neighbors" hcl:"neighbors"`
	Networks               []string                                         `cty:"networks" hcl:"networks"`
	NoPrivateAs            *bool                                            `cty:"no_private_as" hcl:"no_private_as"`
	NoReadvertiseToOverlay *bool                                            `cty:"no_readvertise_to_overlay" hcl:"no_readvertise_to_overlay"`
	TunnelName             *string                                          `cty:"tunnel_name" hcl:"tunnel_name"`
	BgpConfigType          *string                                          `cty:"type" hcl:"type"`
	Via                    string                                           `cty:"via" hcl:"via"`
	VpnName                *string                                          `cty:"vpn_name" hcl:"vpn_name"`
	WanName                *string                                          `cty:"wan_name" hcl:"wan_name"`
}

type OrgDeviceprofileGatewayNeighborsValue struct {
	Disabled     *bool   `cty:"disabled" hcl:"disabled"`
	ExportPolicy *string `cty:"export_policy" hcl:"export_policy"`
	HoldTime     *int64  `cty:"hold_time" hcl:"hold_time"`
	ImportPolicy *string `cty:"import_policy" hcl:"import_policy"`
	MultihopTtl  *int64  `cty:"multihop_ttl" hcl:"multihop_ttl"`
	NeighborAs   string  `cty:"neighbor_as" hcl:"neighbor_as"`
}

type OrgDeviceprofileGatewayDhcpdConfigValue struct {
	Config  map[string]OrgDeviceprofileGatewayConfigValue `cty:"config" hcl:"config"`
	Enabled *bool                                         `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewayConfigValue struct {
	DnsServers         []string                                                  `cty:"dns_servers" hcl:"dns_servers"`
	DnsSuffix          []string                                                  `cty:"dns_suffix" hcl:"dns_suffix"`
	FixedBindings      map[string]OrgDeviceprofileGatewayFixedBindingsValue      `cty:"fixed_bindings" hcl:"fixed_bindings"`
	Gateway            *string                                                   `cty:"gateway" hcl:"gateway"`
	Ip6End             *string                                                   `cty:"ip6_end" hcl:"ip6_end"`
	Ip6Start           *string                                                   `cty:"ip6_start" hcl:"ip6_start"`
	IpEnd4             *string                                                   `cty:"ip_end" hcl:"ip_end"`
	IpStart4           *string                                                   `cty:"ip_start" hcl:"ip_start"`
	LeaseTime          *int64                                                    `cty:"lease_time" hcl:"lease_time"`
	Options            map[string]OrgDeviceprofileGatewayOptionsValue            `cty:"options" hcl:"options"`
	ServerIdOverride   *bool                                                     `cty:"server_id_override" hcl:"server_id_override"`
	Servers4           []string                                                  `cty:"servers" hcl:"servers"`
	Serversv6          []string                                                  `cty:"serversv6" hcl:"serversv6"`
	Type4              *string                                                   `cty:"type" hcl:"type"`
	Type6              *string                                                   `cty:"type6" hcl:"type6"`
	VendorEncapsulated map[string]OrgDeviceprofileGatewayVendorEncapsulatedValue `cty:"vendor_encapsulated" hcl:"vendor_encapsulated"`
}

type OrgDeviceprofileGatewayFixedBindingsValue struct {
	Ip   *string `cty:"ip" hcl:"ip"`
	Ip6  *string `cty:"ip6" hcl:"ip6"`
	Name *string `cty:"name" hcl:"name"`
}

type OrgDeviceprofileGatewayOptionsValue struct {
	OptionsType *string `cty:"type" hcl:"type"`
	Value       *string `cty:"value" hcl:"value"`
}

type OrgDeviceprofileGatewayVendorEncapsulatedValue struct {
	VendorEncapsulatedType *string `cty:"type" hcl:"type"`
	Value                  *string `cty:"value" hcl:"value"`
}

type OrgDeviceprofileGatewayExtraRoutesValue struct {
	Via string `cty:"via" hcl:"via"`
}

type OrgDeviceprofileGatewayExtraRoutes6Value struct {
	Via string `cty:"via" hcl:"via"`
}

type OrgDeviceprofileGatewayIdpProfilesValue struct {
	BaseProfile *string                                  `cty:"base_profile" hcl:"base_profile"`
	Name        *string                                  `cty:"name" hcl:"name"`
	OrgId       *string                                  `cty:"org_id" hcl:"org_id"`
	Overwrites  []OrgDeviceprofileGatewayOverwritesValue `cty:"overwrites" hcl:"overwrites"`
}

type OrgDeviceprofileGatewayOverwritesValue struct {
	Action                      *string                                                  `cty:"action" hcl:"action"`
	IpdProfileOverwriteMatching *OrgDeviceprofileGatewayIpdProfileOverwriteMatchingValue `cty:"matching" hcl:"matching"`
	Name                        *string                                                  `cty:"name" hcl:"name"`
}

type OrgDeviceprofileGatewayIpdProfileOverwriteMatchingValue struct {
	AttackName []string `cty:"attack_name" hcl:"attack_name"`
	DstSubnet  []string `cty:"dst_subnet" hcl:"dst_subnet"`
	Severity   []string `cty:"severity" hcl:"severity"`
}

type OrgDeviceprofileGatewayIpConfigsValue struct {
	Ip            *string  `cty:"ip" hcl:"ip"`
	Ip6           *string  `cty:"ip6" hcl:"ip6"`
	Netmask       *string  `cty:"netmask" hcl:"netmask"`
	Netmask6      *string  `cty:"netmask6" hcl:"netmask6"`
	SecondaryIps  []string `cty:"secondary_ips" hcl:"secondary_ips"`
	IpConfigsType *string  `cty:"type" hcl:"type"`
	Type6         *string  `cty:"type6" hcl:"type6"`
}

type OrgDeviceprofileGatewayNetworksValue struct {
	DisallowMistServices *bool                                            `cty:"disallow_mist_services" hcl:"disallow_mist_services"`
	Gateway              *string                                          `cty:"gateway" hcl:"gateway"`
	Gateway6             *string                                          `cty:"gateway6" hcl:"gateway6"`
	InternalAccess       *OrgDeviceprofileGatewayInternalAccessValue      `cty:"internal_access" hcl:"internal_access"`
	InternetAccess       *OrgDeviceprofileGatewayInternetAccessValue      `cty:"internet_access" hcl:"internet_access"`
	Isolation            *bool                                            `cty:"isolation" hcl:"isolation"`
	Multicast            *OrgDeviceprofileGatewayMulticastValue           `cty:"multicast" hcl:"multicast"`
	Name                 *string                                          `cty:"name" hcl:"name"`
	RoutedForNetworks    []string                                         `cty:"routed_for_networks" hcl:"routed_for_networks"`
	Subnet               string                                           `cty:"subnet" hcl:"subnet"`
	Subnet6              *string                                          `cty:"subnet6" hcl:"subnet6"`
	Tenants              map[string]OrgDeviceprofileGatewayTenantsValue   `cty:"tenants" hcl:"tenants"`
	VlanId               *string                                          `cty:"vlan_id" hcl:"vlan_id"`
	VpnAccess            map[string]OrgDeviceprofileGatewayVpnAccessValue `cty:"vpn_access" hcl:"vpn_access"`
}

type OrgDeviceprofileGatewayInternalAccessValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewayInternetAccessValue struct {
	CreateSimpleServicePolicy    *bool                                                               `cty:"create_simple_service_policy" hcl:"create_simple_service_policy"`
	Enabled                      *bool                                                               `cty:"enabled" hcl:"enabled"`
	InternetAccessDestinationNat map[string]OrgDeviceprofileGatewayInternetAccessDestinationNatValue `cty:"destination_nat" hcl:"destination_nat"`
	InternetAccessStaticNat      map[string]OrgDeviceprofileGatewayInternetAccessStaticNatValue      `cty:"static_nat" hcl:"static_nat"`
	Restricted                   *bool                                                               `cty:"restricted" hcl:"restricted"`
}

type OrgDeviceprofileGatewayInternetAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip" hcl:"internal_ip"`
	Name       string  `cty:"name" hcl:"name"`
	Port       *string `cty:"port" hcl:"port"`
	WanName    *string `cty:"wan_name" hcl:"wan_name"`
}

type OrgDeviceprofileGatewayInternetAccessStaticNatValue struct {
	InternalIp string  `cty:"internal_ip" hcl:"internal_ip"`
	Name       string  `cty:"name" hcl:"name"`
	WanName    *string `cty:"wan_name" hcl:"wan_name"`
}

type OrgDeviceprofileGatewayMulticastValue struct {
	DisableIgmp *bool                                         `cty:"disable_igmp" hcl:"disable_igmp"`
	Enabled     *bool                                         `cty:"enabled" hcl:"enabled"`
	Groups      map[string]OrgDeviceprofileGatewayGroupsValue `cty:"groups" hcl:"groups"`
}

type OrgDeviceprofileGatewayGroupsValue struct {
	RpIp *string `cty:"rp_ip" hcl:"rp_ip"`
}

type OrgDeviceprofileGatewayTenantsValue struct {
	Addresses []string `cty:"addresses" hcl:"addresses"`
}

type OrgDeviceprofileGatewayVpnAccessValue struct {
	AdvertisedSubnet          *string                                                        `cty:"advertised_subnet" hcl:"advertised_subnet"`
	AllowPing                 *bool                                                          `cty:"allow_ping" hcl:"allow_ping"`
	NatPool                   *string                                                        `cty:"nat_pool" hcl:"nat_pool"`
	NoReadvertiseToLanBgp     *bool                                                          `cty:"no_readvertise_to_lan_bgp" hcl:"no_readvertise_to_lan_bgp"`
	NoReadvertiseToLanOspf    *bool                                                          `cty:"no_readvertise_to_lan_ospf" hcl:"no_readvertise_to_lan_ospf"`
	NoReadvertiseToOverlay    *bool                                                          `cty:"no_readvertise_to_overlay" hcl:"no_readvertise_to_overlay"`
	OtherVrfs                 []string                                                       `cty:"other_vrfs" hcl:"other_vrfs"`
	Routed                    *bool                                                          `cty:"routed" hcl:"routed"`
	SourceNat                 *OrgDeviceprofileGatewaySourceNatValue                         `cty:"source_nat" hcl:"source_nat"`
	SummarizedSubnet          *string                                                        `cty:"summarized_subnet" hcl:"summarized_subnet"`
	SummarizedSubnetToLanBgp  *string                                                        `cty:"summarized_subnet_to_lan_bgp" hcl:"summarized_subnet_to_lan_bgp"`
	SummarizedSubnetToLanOspf *string                                                        `cty:"summarized_subnet_to_lan_ospf" hcl:"summarized_subnet_to_lan_ospf"`
	VpnAccessDestinationNat   map[string]OrgDeviceprofileGatewayVpnAccessDestinationNatValue `cty:"destination_nat" hcl:"destination_nat"`
	VpnAccessStaticNat        map[string]OrgDeviceprofileGatewayVpnAccessStaticNatValue      `cty:"static_nat" hcl:"static_nat"`
}

type OrgDeviceprofileGatewaySourceNatValue struct {
	ExternalIp *string `cty:"external_ip" hcl:"external_ip"`
}

type OrgDeviceprofileGatewayVpnAccessDestinationNatValue struct {
	InternalIp *string `cty:"internal_ip" hcl:"internal_ip"`
	Name       *string `cty:"name" hcl:"name"`
	Port       *string `cty:"port" hcl:"port"`
}

type OrgDeviceprofileGatewayVpnAccessStaticNatValue struct {
	InternalIp string `cty:"internal_ip" hcl:"internal_ip"`
	Name       string `cty:"name" hcl:"name"`
}

type OrgDeviceprofileGatewayOobIpConfigValue struct {
	Gateway              *string                            `cty:"gateway" hcl:"gateway"`
	Ip                   *string                            `cty:"ip" hcl:"ip"`
	Netmask              *string                            `cty:"netmask" hcl:"netmask"`
	Node1                *OrgDeviceprofileGatewayNode1Value `cty:"node1" hcl:"node1"`
	OobIpConfigType      *string                            `cty:"type" hcl:"type"`
	UseMgmtVrf           *bool                              `cty:"use_mgmt_vrf" hcl:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool                              `cty:"use_mgmt_vrf_for_host_out" hcl:"use_mgmt_vrf_for_host_out"`
	VlanId               *string                            `cty:"vlan_id" hcl:"vlan_id"`
}

type OrgDeviceprofileGatewayNode1Value struct {
	Gateway              *string `cty:"gateway" hcl:"gateway"`
	Ip                   *string `cty:"ip" hcl:"ip"`
	Netmask              *string `cty:"netmask" hcl:"netmask"`
	Node1Type            *string `cty:"type" hcl:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf" hcl:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out" hcl:"use_mgmt_vrf_for_host_out"`
	VlanId               *string `cty:"vlan_id" hcl:"vlan_id"`
}

type OrgDeviceprofileGatewayPathPreferencesValue struct {
	Paths    []OrgDeviceprofileGatewayPathsValue `cty:"paths" hcl:"paths"`
	Strategy *string                             `cty:"strategy" hcl:"strategy"`
}

type OrgDeviceprofileGatewayPathsValue struct {
	Cost           *int64   `cty:"cost" hcl:"cost"`
	Disabled       *bool    `cty:"disabled" hcl:"disabled"`
	GatewayIp      *string  `cty:"gateway_ip" hcl:"gateway_ip"`
	InternetAccess *bool    `cty:"internet_access" hcl:"internet_access"`
	Name           *string  `cty:"name" hcl:"name"`
	Networks       []string `cty:"networks" hcl:"networks"`
	TargetIps      []string `cty:"target_ips" hcl:"target_ips"`
	PathsType      string   `cty:"type" hcl:"type"`
	WanName        *string  `cty:"wan_name" hcl:"wan_name"`
}

type OrgDeviceprofileGatewayPortConfigValue struct {
	AeDisableLacp    *bool                                                  `cty:"ae_disable_lacp" hcl:"ae_disable_lacp"`
	AeIdx            *string                                                `cty:"ae_idx" hcl:"ae_idx"`
	AeLacpForceUp    *bool                                                  `cty:"ae_lacp_force_up" hcl:"ae_lacp_force_up"`
	Aggregated       *bool                                                  `cty:"aggregated" hcl:"aggregated"`
	Critical         *bool                                                  `cty:"critical" hcl:"critical"`
	Description      *string                                                `cty:"description" hcl:"description"`
	DisableAutoneg   *bool                                                  `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled         *bool                                                  `cty:"disabled" hcl:"disabled"`
	DslType          *string                                                `cty:"dsl_type" hcl:"dsl_type"`
	DslVci           *int64                                                 `cty:"dsl_vci" hcl:"dsl_vci"`
	DslVpi           *int64                                                 `cty:"dsl_vpi" hcl:"dsl_vpi"`
	Duplex           *string                                                `cty:"duplex" hcl:"duplex"`
	LteApn           *string                                                `cty:"lte_apn" hcl:"lte_apn"`
	LteAuth          *string                                                `cty:"lte_auth" hcl:"lte_auth"`
	LteBackup        *bool                                                  `cty:"lte_backup" hcl:"lte_backup"`
	LtePassword      *string                                                `cty:"lte_password" hcl:"lte_password"`
	LteUsername      *string                                                `cty:"lte_username" hcl:"lte_username"`
	Mtu              *int64                                                 `cty:"mtu" hcl:"mtu"`
	Name             *string                                                `cty:"name" hcl:"name"`
	Networks         []string                                               `cty:"networks" hcl:"networks"`
	OuterVlanId      *int64                                                 `cty:"outer_vlan_id" hcl:"outer_vlan_id"`
	PoeDisabled      *bool                                                  `cty:"poe_disabled" hcl:"poe_disabled"`
	PortIpConfig     *OrgDeviceprofileGatewayPortIpConfigValue              `cty:"ip_config" hcl:"ip_config"`
	PortNetwork      *string                                                `cty:"port_network" hcl:"port_network"`
	PreserveDscp     *bool                                                  `cty:"preserve_dscp" hcl:"preserve_dscp"`
	Redundant        *bool                                                  `cty:"redundant" hcl:"redundant"`
	RedundantGroup   *int64                                                 `cty:"redundant_group" hcl:"redundant_group"`
	RethIdx          *string                                                `cty:"reth_idx" hcl:"reth_idx"`
	RethNode         *string                                                `cty:"reth_node" hcl:"reth_node"`
	RethNodes        []string                                               `cty:"reth_nodes" hcl:"reth_nodes"`
	Speed            *string                                                `cty:"speed" hcl:"speed"`
	SsrNoVirtualMac  *bool                                                  `cty:"ssr_no_virtual_mac" hcl:"ssr_no_virtual_mac"`
	SvrPortRange     *string                                                `cty:"svr_port_range" hcl:"svr_port_range"`
	TrafficShaping   *OrgDeviceprofileGatewayTrafficShapingValue            `cty:"traffic_shaping" hcl:"traffic_shaping"`
	Usage            string                                                 `cty:"usage" hcl:"usage"`
	VlanId           *string                                                `cty:"vlan_id" hcl:"vlan_id"`
	VpnPaths         map[string]OrgDeviceprofileGatewayVpnPathsValue        `cty:"vpn_paths" hcl:"vpn_paths"`
	WanArpPolicer    *string                                                `cty:"wan_arp_policer" hcl:"wan_arp_policer"`
	WanExtIp         *string                                                `cty:"wan_ext_ip" hcl:"wan_ext_ip"`
	WanExtIp6        *string                                                `cty:"wan_ext_ip6" hcl:"wan_ext_ip6"`
	WanExtraRoutes   map[string]OrgDeviceprofileGatewayWanExtraRoutesValue  `cty:"wan_extra_routes" hcl:"wan_extra_routes"`
	WanExtraRoutes6  map[string]OrgDeviceprofileGatewayWanExtraRoutes6Value `cty:"wan_extra_routes6" hcl:"wan_extra_routes6"`
	WanNetworks      []string                                               `cty:"wan_networks" hcl:"wan_networks"`
	WanProbeOverride *OrgDeviceprofileGatewayWanProbeOverrideValue          `cty:"wan_probe_override" hcl:"wan_probe_override"`
	WanSourceNat     *OrgDeviceprofileGatewayWanSourceNatValue              `cty:"wan_source_nat" hcl:"wan_source_nat"`
	WanSpeedtestMode *string                                                `cty:"wan_speedtest_mode" hcl:"wan_speedtest_mode"`
	WanType          *string                                                `cty:"wan_type" hcl:"wan_type"`
}

type OrgDeviceprofileGatewayPortIpConfigValue struct {
	Dns              []string `cty:"dns" hcl:"dns"`
	DnsSuffix        []string `cty:"dns_suffix" hcl:"dns_suffix"`
	Gateway          *string  `cty:"gateway" hcl:"gateway"`
	Gateway6         *string  `cty:"gateway6" hcl:"gateway6"`
	Ip               *string  `cty:"ip" hcl:"ip"`
	Ip6              *string  `cty:"ip6" hcl:"ip6"`
	Netmask          *string  `cty:"netmask" hcl:"netmask"`
	Netmask6         *string  `cty:"netmask6" hcl:"netmask6"`
	Network          *string  `cty:"network" hcl:"network"`
	PoserPassword    *string  `cty:"poser_password" hcl:"poser_password"`
	PppoeAuth        *string  `cty:"pppoe_auth" hcl:"pppoe_auth"`
	PppoeUsername    *string  `cty:"pppoe_username" hcl:"pppoe_username"`
	PortIpConfigType *string  `cty:"type" hcl:"type"`
	Type6            *string  `cty:"type6" hcl:"type6"`
}

type OrgDeviceprofileGatewayTrafficShapingValue struct {
	ClassPercentages []int64 `cty:"class_percentages" hcl:"class_percentages"`
	Enabled          *bool   `cty:"enabled" hcl:"enabled"`
	MaxTxKbps        *int64  `cty:"max_tx_kbps" hcl:"max_tx_kbps"`
}

type OrgDeviceprofileGatewayVpnPathsValue struct {
	BfdProfile       *string                                     `cty:"bfd_profile" hcl:"bfd_profile"`
	BfdUseTunnelMode *bool                                       `cty:"bfd_use_tunnel_mode" hcl:"bfd_use_tunnel_mode"`
	Preference       *int64                                      `cty:"preference" hcl:"preference"`
	Role             *string                                     `cty:"role" hcl:"role"`
	TrafficShaping   *OrgDeviceprofileGatewayTrafficShapingValue `cty:"traffic_shaping" hcl:"traffic_shaping"`
}

type OrgDeviceprofileGatewayWanExtraRoutesValue struct {
	Via *string `cty:"via" hcl:"via"`
}

type OrgDeviceprofileGatewayWanExtraRoutes6Value struct {
	Via *string `cty:"via" hcl:"via"`
}

type OrgDeviceprofileGatewayWanProbeOverrideValue struct {
	Ip6s         []string `cty:"ip6s" hcl:"ip6s"`
	Ips          []string `cty:"ips" hcl:"ips"`
	ProbeProfile *string  `cty:"probe_profile" hcl:"probe_profile"`
}

type OrgDeviceprofileGatewayWanSourceNatValue struct {
	Disabled *bool   `cty:"disabled" hcl:"disabled"`
	Nat6Pool *string `cty:"nat6_pool" hcl:"nat6_pool"`
	NatPool  *string `cty:"nat_pool" hcl:"nat_pool"`
}

type OrgDeviceprofileGatewayRoutingPoliciesValue struct {
	Terms []OrgDeviceprofileGatewayTermsValue `cty:"terms" hcl:"terms"`
}

type OrgDeviceprofileGatewayTermsValue struct {
	Actions                   *OrgDeviceprofileGatewayActionsValue                   `cty:"actions" hcl:"actions"`
	RoutingPolicyTermMatching *OrgDeviceprofileGatewayRoutingPolicyTermMatchingValue `cty:"matching" hcl:"matching"`
}

type OrgDeviceprofileGatewayActionsValue struct {
	Accept            *bool    `cty:"accept" hcl:"accept"`
	AddCommunity      []string `cty:"add_community" hcl:"add_community"`
	AddTargetVrfs     []string `cty:"add_target_vrfs" hcl:"add_target_vrfs"`
	Community         []string `cty:"community" hcl:"community"`
	ExcludeAsPath     []string `cty:"exclude_as_path" hcl:"exclude_as_path"`
	ExcludeCommunity  []string `cty:"exclude_community" hcl:"exclude_community"`
	ExportCommunities []string `cty:"export_communities" hcl:"export_communities"`
	LocalPreference   *string  `cty:"local_preference" hcl:"local_preference"`
	PrependAsPath     []string `cty:"prepend_as_path" hcl:"prepend_as_path"`
}

type OrgDeviceprofileGatewayRoutingPolicyTermMatchingValue struct {
	AsPath         []string                                 `cty:"as_path" hcl:"as_path"`
	Community      []string                                 `cty:"community" hcl:"community"`
	Network        []string                                 `cty:"network" hcl:"network"`
	Prefix         []string                                 `cty:"prefix" hcl:"prefix"`
	Protocol       []string                                 `cty:"protocol" hcl:"protocol"`
	RouteExists    *OrgDeviceprofileGatewayRouteExistsValue `cty:"route_exists" hcl:"route_exists"`
	VpnNeighborMac []string                                 `cty:"vpn_neighbor_mac" hcl:"vpn_neighbor_mac"`
	VpnPath        []string                                 `cty:"vpn_path" hcl:"vpn_path"`
	VpnPathSla     *OrgDeviceprofileGatewayVpnPathSlaValue  `cty:"vpn_path_sla" hcl:"vpn_path_sla"`
}

type OrgDeviceprofileGatewayRouteExistsValue struct {
	Route   *string `cty:"route" hcl:"route"`
	VrfName *string `cty:"vrf_name" hcl:"vrf_name"`
}

type OrgDeviceprofileGatewayVpnPathSlaValue struct {
	MaxJitter  *int64 `cty:"max_jitter" hcl:"max_jitter"`
	MaxLatency *int64 `cty:"max_latency" hcl:"max_latency"`
	MaxLoss    *int64 `cty:"max_loss" hcl:"max_loss"`
}

type OrgDeviceprofileGatewayServicePoliciesValue struct {
	Action          *string                                `cty:"action" hcl:"action"`
	Antivirus       *OrgDeviceprofileGatewayAntivirusValue `cty:"antivirus" hcl:"antivirus"`
	Appqoe          *OrgDeviceprofileGatewayAppqoeValue    `cty:"appqoe" hcl:"appqoe"`
	Ewf             []OrgDeviceprofileGatewayEwfValue      `cty:"ewf" hcl:"ewf"`
	Idp             *OrgDeviceprofileGatewayIdpValue       `cty:"idp" hcl:"idp"`
	LocalRouting    *bool                                  `cty:"local_routing" hcl:"local_routing"`
	Name            *string                                `cty:"name" hcl:"name"`
	PathPreference  *string                                `cty:"path_preference" hcl:"path_preference"`
	ServicepolicyId *string                                `cty:"servicepolicy_id" hcl:"servicepolicy_id"`
	Services        []string                               `cty:"services" hcl:"services"`
	Skyatp          *OrgDeviceprofileGatewaySkyatpValue    `cty:"skyatp" hcl:"skyatp"`
	SslProxy        *OrgDeviceprofileGatewaySslProxyValue  `cty:"ssl_proxy" hcl:"ssl_proxy"`
	Syslog          *OrgDeviceprofileGatewaySyslogValue    `cty:"syslog" hcl:"syslog"`
	Tenants         []string                               `cty:"tenants" hcl:"tenants"`
}

type OrgDeviceprofileGatewayAntivirusValue struct {
	AvprofileId *string `cty:"avprofile_id" hcl:"avprofile_id"`
	Enabled     *bool   `cty:"enabled" hcl:"enabled"`
	Profile     *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewayAppqoeValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewayEwfValue struct {
	AlertOnly    *bool   `cty:"alert_only" hcl:"alert_only"`
	BlockMessage *string `cty:"block_message" hcl:"block_message"`
	Enabled      *bool   `cty:"enabled" hcl:"enabled"`
	Profile      *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewayIdpValue struct {
	AlertOnly    *bool   `cty:"alert_only" hcl:"alert_only"`
	Enabled      *bool   `cty:"enabled" hcl:"enabled"`
	IdpprofileId *string `cty:"idpprofile_id" hcl:"idpprofile_id"`
	Profile      *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewaySkyatpValue struct {
	DnsDgaDetection    *OrgDeviceprofileGatewayDnsDgaDetectionValue    `cty:"dns_dga_detection" hcl:"dns_dga_detection"`
	DnsTunnelDetection *OrgDeviceprofileGatewayDnsTunnelDetectionValue `cty:"dns_tunnel_detection" hcl:"dns_tunnel_detection"`
	HttpInspection     *OrgDeviceprofileGatewayHttpInspectionValue     `cty:"http_inspection" hcl:"http_inspection"`
	IotDevicePolicy    *OrgDeviceprofileGatewayIotDevicePolicyValue    `cty:"iot_device_policy" hcl:"iot_device_policy"`
}

type OrgDeviceprofileGatewayDnsDgaDetectionValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Profile *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewayDnsTunnelDetectionValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Profile *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewayHttpInspectionValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Profile *string `cty:"profile" hcl:"profile"`
}

type OrgDeviceprofileGatewayIotDevicePolicyValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewaySslProxyValue struct {
	CiphersCategory *string `cty:"ciphers_category" hcl:"ciphers_category"`
	Enabled         *bool   `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewaySyslogValue struct {
	Enabled     *bool    `cty:"enabled" hcl:"enabled"`
	ServerNames []string `cty:"server_names" hcl:"server_names"`
}

type OrgDeviceprofileGatewayTunnelConfigsValue struct {
	AutoProvision  *OrgDeviceprofileGatewayAutoProvisionValue   `cty:"auto_provision" hcl:"auto_provision"`
	IkeLifetime    *int64                                       `cty:"ike_lifetime" hcl:"ike_lifetime"`
	IkeMode        *string                                      `cty:"ike_mode" hcl:"ike_mode"`
	IkeProposals   []OrgDeviceprofileGatewayIkeProposalsValue   `cty:"ike_proposals" hcl:"ike_proposals"`
	IpsecLifetime  *int64                                       `cty:"ipsec_lifetime" hcl:"ipsec_lifetime"`
	IpsecProposals []OrgDeviceprofileGatewayIpsecProposalsValue `cty:"ipsec_proposals" hcl:"ipsec_proposals"`
	LocalId        *string                                      `cty:"local_id" hcl:"local_id"`
	LocalSubnets   []string                                     `cty:"local_subnets" hcl:"local_subnets"`
	Mode           *string                                      `cty:"mode" hcl:"mode"`
	Networks       []string                                     `cty:"networks" hcl:"networks"`
	Primary        *OrgDeviceprofileGatewayPrimaryValue         `cty:"primary" hcl:"primary"`
	Probe          *OrgDeviceprofileGatewayProbeValue           `cty:"probe" hcl:"probe"`
	Protocol       *string                                      `cty:"protocol" hcl:"protocol"`
	Provider       string                                       `cty:"provider" hcl:"provider"`
	Psk            *string                                      `cty:"psk" hcl:"psk"`
	RemoteSubnets  []string                                     `cty:"remote_subnets" hcl:"remote_subnets"`
	Secondary      *OrgDeviceprofileGatewaySecondaryValue       `cty:"secondary" hcl:"secondary"`
	Version        *string                                      `cty:"version" hcl:"version"`
}

type OrgDeviceprofileGatewayAutoProvisionValue struct {
	AutoProvisionPrimary   *OrgDeviceprofileGatewayAutoProvisionPrimaryValue   `cty:"primary" hcl:"primary"`
	AutoProvisionSecondary *OrgDeviceprofileGatewayAutoProvisionSecondaryValue `cty:"secondary" hcl:"secondary"`
	Enabled                *bool                                               `cty:"enabled" hcl:"enabled"`
	Latlng                 *OrgDeviceprofileGatewayLatlngValue                 `cty:"latlng" hcl:"latlng"`
	Provider               *string                                             `cty:"provider" hcl:"provider"`
	Region                 *string                                             `cty:"region" hcl:"region"`
	ServiceConnection      *string                                             `cty:"service_connection" hcl:"service_connection"`
}

type OrgDeviceprofileGatewayAutoProvisionPrimaryValue struct {
	ProbeIps []string `cty:"probe_ips" hcl:"probe_ips"`
	WanNames []string `cty:"wan_names" hcl:"wan_names"`
}

type OrgDeviceprofileGatewayAutoProvisionSecondaryValue struct {
	ProbeIps []string `cty:"probe_ips" hcl:"probe_ips"`
	WanNames []string `cty:"wan_names" hcl:"wan_names"`
}

type OrgDeviceprofileGatewayLatlngValue struct {
	Lat float64 `cty:"lat" hcl:"lat"`
	Lng float64 `cty:"lng" hcl:"lng"`
}

type OrgDeviceprofileGatewayIkeProposalsValue struct {
	AuthAlgo *string `cty:"auth_algo" hcl:"auth_algo"`
	DhGroup  *string `cty:"dh_group" hcl:"dh_group"`
	EncAlgo  *string `cty:"enc_algo" hcl:"enc_algo"`
}

type OrgDeviceprofileGatewayIpsecProposalsValue struct {
	AuthAlgo *string `cty:"auth_algo" hcl:"auth_algo"`
	DhGroup  *string `cty:"dh_group" hcl:"dh_group"`
	EncAlgo  *string `cty:"enc_algo" hcl:"enc_algo"`
}

type OrgDeviceprofileGatewayPrimaryValue struct {
	Hosts       []string `cty:"hosts" hcl:"hosts"`
	InternalIps []string `cty:"internal_ips" hcl:"internal_ips"`
	ProbeIps    []string `cty:"probe_ips" hcl:"probe_ips"`
	RemoteIds   []string `cty:"remote_ids" hcl:"remote_ids"`
	WanNames    []string `cty:"wan_names" hcl:"wan_names"`
}

type OrgDeviceprofileGatewayProbeValue struct {
	Interval  *int64  `cty:"interval" hcl:"interval"`
	Threshold *int64  `cty:"threshold" hcl:"threshold"`
	Timeout   *int64  `cty:"timeout" hcl:"timeout"`
	ProbeType *string `cty:"type" hcl:"type"`
}

type OrgDeviceprofileGatewaySecondaryValue struct {
	Hosts       []string `cty:"hosts" hcl:"hosts"`
	InternalIps []string `cty:"internal_ips" hcl:"internal_ips"`
	ProbeIps    []string `cty:"probe_ips" hcl:"probe_ips"`
	RemoteIds   []string `cty:"remote_ids" hcl:"remote_ids"`
	WanNames    []string `cty:"wan_names" hcl:"wan_names"`
}

type OrgDeviceprofileGatewayTunnelProviderOptionsValue struct {
	Jse     *OrgDeviceprofileGatewayJseValue     `cty:"jse" hcl:"jse"`
	Prisma  *OrgDeviceprofileGatewayPrismaValue  `cty:"prisma" hcl:"prisma"`
	Zscaler *OrgDeviceprofileGatewayZscalerValue `cty:"zscaler" hcl:"zscaler"`
}

type OrgDeviceprofileGatewayJseValue struct {
	NumUsers *int64  `cty:"num_users" hcl:"num_users"`
	OrgName  *string `cty:"org_name" hcl:"org_name"`
}

type OrgDeviceprofileGatewayPrismaValue struct {
	ServiceAccountName *string `cty:"service_account_name" hcl:"service_account_name"`
}

type OrgDeviceprofileGatewayZscalerValue struct {
	AupBlockInternetUntilAccepted       *bool                                      `cty:"aup_block_internet_until_accepted" hcl:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool                                      `cty:"aup_enabled" hcl:"aup_enabled"`
	AupForceSslInspection               *bool                                      `cty:"aup_force_ssl_inspection" hcl:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64                                     `cty:"aup_timeout_in_days" hcl:"aup_timeout_in_days"`
	AuthRequired                        *bool                                      `cty:"auth_required" hcl:"auth_required"`
	CautionEnabled                      *bool                                      `cty:"caution_enabled" hcl:"caution_enabled"`
	DnBandwidth                         *float64                                   `cty:"dn_bandwidth" hcl:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64                                     `cty:"idle_time_in_minutes" hcl:"idle_time_in_minutes"`
	OfwEnabled                          *bool                                      `cty:"ofw_enabled" hcl:"ofw_enabled"`
	SubLocations                        []OrgDeviceprofileGatewaySubLocationsValue `cty:"sub_locations" hcl:"sub_locations"`
	SurrogateIp                         *bool                                      `cty:"surrogate_ip" hcl:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool                                      `cty:"surrogate_ip_enforced_for_known_browsers" hcl:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64                                     `cty:"surrogate_refresh_time_in_minutes" hcl:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64                                   `cty:"up_bandwidth" hcl:"up_bandwidth"`
	XffForwardEnabled                   *bool                                      `cty:"xff_forward_enabled" hcl:"xff_forward_enabled"`
}

type OrgDeviceprofileGatewaySubLocationsValue struct {
	AupBlockInternetUntilAccepted       *bool    `cty:"aup_block_internet_until_accepted" hcl:"aup_block_internet_until_accepted"`
	AupEnabled                          *bool    `cty:"aup_enabled" hcl:"aup_enabled"`
	AupForceSslInspection               *bool    `cty:"aup_force_ssl_inspection" hcl:"aup_force_ssl_inspection"`
	AupTimeoutInDays                    *int64   `cty:"aup_timeout_in_days" hcl:"aup_timeout_in_days"`
	AuthRequired                        *bool    `cty:"auth_required" hcl:"auth_required"`
	CautionEnabled                      *bool    `cty:"caution_enabled" hcl:"caution_enabled"`
	DnBandwidth                         *float64 `cty:"dn_bandwidth" hcl:"dn_bandwidth"`
	IdleTimeInMinutes                   *int64   `cty:"idle_time_in_minutes" hcl:"idle_time_in_minutes"`
	Name                                *string  `cty:"name" hcl:"name"`
	OfwEnabled                          *bool    `cty:"ofw_enabled" hcl:"ofw_enabled"`
	SurrogateIp                         *bool    `cty:"surrogate_ip" hcl:"surrogate_ip"`
	SurrogateIpEnforcedForKnownBrowsers *bool    `cty:"surrogate_ip_enforced_for_known_browsers" hcl:"surrogate_ip_enforced_for_known_browsers"`
	SurrogateRefreshTimeInMinutes       *int64   `cty:"surrogate_refresh_time_in_minutes" hcl:"surrogate_refresh_time_in_minutes"`
	UpBandwidth                         *float64 `cty:"up_bandwidth" hcl:"up_bandwidth"`
}

type OrgDeviceprofileGatewayVrfConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileGatewayVrfInstancesValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
}
