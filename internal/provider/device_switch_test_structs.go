package provider

type DeviceSwitchModel struct {
	AclPolicies           []DeviceSwitchAclPoliciesValue                  `hcl:"acl_policies"`
	AclTags               map[string]DeviceSwitchAclTagsValue             `hcl:"acl_tags"`
	AdditionalConfigCmds  []string                                        `hcl:"additional_config_cmds"`
	BgpConfig             map[string]DeviceSwitchBgpConfigValue           `hcl:"bgp_config"`
	DeviceId              string                                          `hcl:"device_id"`
	DhcpSnooping          *DeviceSwitchDhcpSnoopingValue                  `hcl:"dhcp_snooping"`
	DhcpdConfig           *DeviceSwitchDhcpdConfigValue                   `hcl:"dhcpd_config"`
	DisableAutoConfig     *bool                                           `hcl:"disable_auto_config"`
	DnsServers            []string                                        `hcl:"dns_servers"`
	DnsSuffix             []string                                        `hcl:"dns_suffix"`
	ExtraRoutes           map[string]DeviceSwitchExtraRoutesValue         `hcl:"extra_routes"`
	ExtraRoutes6          map[string]DeviceSwitchExtraRoutes6Value        `hcl:"extra_routes6"`
	IpConfig              *DeviceSwitchIpConfigValue                      `hcl:"ip_config"`
	LocalPortConfig       map[string]DeviceSwitchLocalPortConfigValue     `hcl:"local_port_config"`
	Managed               *bool                                           `hcl:"managed"`
	MapId                 *string                                         `hcl:"map_id"`
	MistNac               *DeviceSwitchMistNacValue                       `hcl:"mist_nac"`
	Name                  string                                          `hcl:"name"`
	Networks              map[string]DeviceSwitchNetworksValue            `hcl:"networks"`
	Notes                 *string                                         `hcl:"notes"`
	NtpServers            []string                                        `hcl:"ntp_servers"`
	OobIpConfig           *DeviceSwitchOobIpConfigValue                   `hcl:"oob_ip_config"`
	OspfAreas             map[string]DeviceSwitchOspfAreasValue           `hcl:"ospf_areas"`
	OspfConfig            *DeviceSwitchOspfConfigValue                    `hcl:"ospf_config"`
	OtherIpConfigs        map[string]DeviceSwitchOtherIpConfigsValue      `hcl:"other_ip_configs"`
	PortConfig            map[string]DeviceSwitchPortConfigValue          `hcl:"port_config"`
	PortConfigOverwrite   map[string]DeviceSwitchPortConfigOverwriteValue `hcl:"port_config_overwrite"`
	PortMirroring         map[string]DeviceSwitchPortMirroringValue       `hcl:"port_mirroring"`
	PortUsages            map[string]DeviceSwitchPortUsagesValue          `hcl:"port_usages"`
	RadiusConfig          *DeviceSwitchRadiusConfigValue                  `hcl:"radius_config"`
	RemoteSyslog          *DeviceSwitchRemoteSyslogValue                  `hcl:"remote_syslog"`
	Role                  *string                                         `hcl:"role"`
	RouterId              *string                                         `hcl:"router_id"`
	SiteId                string                                          `hcl:"site_id"`
	SnmpConfig            *DeviceSwitchSnmpConfigValue                    `hcl:"snmp_config"`
	StpConfig             *DeviceSwitchStpConfigValue                     `hcl:"stp_config"`
	SwitchMgmt            *DeviceSwitchSwitchMgmtValue                    `hcl:"switch_mgmt"`
	UseRouterIdAsSourceIp *bool                                           `hcl:"use_router_id_as_source_ip"`
	Vars                  map[string]string                               `hcl:"vars"`
	VirtualChassis        *DeviceSwitchVirtualChassisValue                `hcl:"virtual_chassis"`
	VrfConfig             *DeviceSwitchVrfConfigValue                     `hcl:"vrf_config"`
	VrfInstances          map[string]DeviceSwitchVrfInstancesValue        `hcl:"vrf_instances"`
	VrrpConfig            *DeviceSwitchVrrpConfigValue                    `hcl:"vrrp_config"`
	X                     *float64                                        `hcl:"x"`
	Y                     *float64                                        `hcl:"y"`
}

type DeviceSwitchAclPoliciesValue struct {
	Actions []DeviceSwitchActionsValue `cty:"actions" hcl:"actions"`
	Name    *string                    `cty:"name" hcl:"name"`
	SrcTags []string                   `cty:"src_tags" hcl:"src_tags"`
}

type DeviceSwitchActionsValue struct {
	Action *string `cty:"action" hcl:"action"`
	DstTag string  `cty:"dst_tag" hcl:"dst_tag"`
}

type DeviceSwitchAclTagsValue struct {
	EtherTypes  []string                 `cty:"ether_types" hcl:"ether_types"`
	GbpTag      *int64                   `cty:"gbp_tag" hcl:"gbp_tag"`
	Macs        []string                 `cty:"macs" hcl:"macs"`
	Network     *string                  `cty:"network" hcl:"network"`
	PortUsage   *string                  `cty:"port_usage" hcl:"port_usage"`
	RadiusGroup *string                  `cty:"radius_group" hcl:"radius_group"`
	Specs       []DeviceSwitchSpecsValue `cty:"specs" hcl:"specs"`
	Subnets     []string                 `cty:"subnets" hcl:"subnets"`
	AclTagsType string                   `cty:"type" hcl:"type"`
}

type DeviceSwitchSpecsValue struct {
	PortRange *string `cty:"port_range" hcl:"port_range"`
	Protocol  *string `cty:"protocol" hcl:"protocol"`
}

type DeviceSwitchBgpConfigValue struct {
	AuthKey            *string                               `cty:"auth_key" hcl:"auth_key"`
	BfdMinimumInterval *int64                                `cty:"bfd_minimum_interval" hcl:"bfd_minimum_interval"`
	ExportPolicy       *string                               `cty:"export_policy" hcl:"export_policy"`
	HoldTime           *int64                                `cty:"hold_time" hcl:"hold_time"`
	ImportPolicy       *string                               `cty:"import_policy" hcl:"import_policy"`
	LocalAs            string                                `cty:"local_as" hcl:"local_as"`
	Neighbors          map[string]DeviceSwitchNeighborsValue `cty:"neighbors" hcl:"neighbors"`
	Networks           []string                              `cty:"networks" hcl:"networks"`
	BgpConfigType      string                                `cty:"type" hcl:"type"`
}

type DeviceSwitchNeighborsValue struct {
	ExportPolicy *string `cty:"export_policy" hcl:"export_policy"`
	HoldTime     *int64  `cty:"hold_time" hcl:"hold_time"`
	ImportPolicy *string `cty:"import_policy" hcl:"import_policy"`
	MultihopTtl  *int64  `cty:"multihop_ttl" hcl:"multihop_ttl"`
	NeighborAs   string  `cty:"neighbor_as" hcl:"neighbor_as"`
}

type DeviceSwitchDhcpSnoopingValue struct {
	AllNetworks         *bool    `cty:"all_networks" hcl:"all_networks"`
	EnableArpSpoofCheck *bool    `cty:"enable_arp_spoof_check" hcl:"enable_arp_spoof_check"`
	EnableIpSourceGuard *bool    `cty:"enable_ip_source_guard" hcl:"enable_ip_source_guard"`
	Enabled             *bool    `cty:"enabled" hcl:"enabled"`
	Networks            []string `cty:"networks" hcl:"networks"`
}

type DeviceSwitchDhcpdConfigValue struct {
	Config  map[string]DeviceSwitchConfigValue `cty:"config" hcl:"config"`
	Enabled *bool                              `cty:"enabled" hcl:"enabled"`
}

type DeviceSwitchConfigValue struct {
	DnsServers         []string                                       `cty:"dns_servers" hcl:"dns_servers"`
	DnsSuffix          []string                                       `cty:"dns_suffix" hcl:"dns_suffix"`
	FixedBindings      map[string]DeviceSwitchFixedBindingsValue      `cty:"fixed_bindings" hcl:"fixed_bindings"`
	Gateway            *string                                        `cty:"gateway" hcl:"gateway"`
	IpEnd              *string                                        `cty:"ip_end" hcl:"ip_end"`
	IpEnd6             *string                                        `cty:"ip_end6" hcl:"ip_end6"`
	IpStart            *string                                        `cty:"ip_start" hcl:"ip_start"`
	IpStart6           *string                                        `cty:"ip_start6" hcl:"ip_start6"`
	LeaseTime          *int64                                         `cty:"lease_time" hcl:"lease_time"`
	Options            map[string]DeviceSwitchOptionsValue            `cty:"options" hcl:"options"`
	ServerIdOverride   *bool                                          `cty:"server_id_override" hcl:"server_id_override"`
	Servers            []string                                       `cty:"servers" hcl:"servers"`
	Servers6           []string                                       `cty:"servers6" hcl:"servers6"`
	ConfigType         *string                                        `cty:"type" hcl:"type"`
	Type6              *string                                        `cty:"type6" hcl:"type6"`
	VendorEncapsulated map[string]DeviceSwitchVendorEncapsulatedValue `cty:"vendor_encapsulated" hcl:"vendor_encapsulated"`
}

type DeviceSwitchFixedBindingsValue struct {
	Ip   *string `cty:"ip" hcl:"ip"`
	Ip6  *string `cty:"ip6" hcl:"ip6"`
	Name *string `cty:"name" hcl:"name"`
}

type DeviceSwitchOptionsValue struct {
	OptionsType *string `cty:"type" hcl:"type"`
	Value       *string `cty:"value" hcl:"value"`
}

type DeviceSwitchVendorEncapsulatedValue struct {
	VendorEncapsulatedType *string `cty:"type" hcl:"type"`
	Value                  *string `cty:"value" hcl:"value"`
}

type DeviceSwitchExtraRoutesValue struct {
	Discard       *bool                                     `cty:"discard" hcl:"discard"`
	Metric        *int64                                    `cty:"metric" hcl:"metric"`
	NextQualified map[string]DeviceSwitchNextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                                     `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                                    `cty:"preference" hcl:"preference"`
	Via           string                                    `cty:"via" hcl:"via"`
}

type DeviceSwitchNextQualifiedValue struct {
	Metric     *int64 `cty:"metric" hcl:"metric"`
	Preference *int64 `cty:"preference" hcl:"preference"`
}

type DeviceSwitchExtraRoutes6Value struct {
	Discard       *bool                                     `cty:"discard" hcl:"discard"`
	Metric        *int64                                    `cty:"metric" hcl:"metric"`
	NextQualified map[string]DeviceSwitchNextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                                     `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                                    `cty:"preference" hcl:"preference"`
	Via           string                                    `cty:"via" hcl:"via"`
}

type DeviceSwitchIpConfigValue struct {
	Dns          []string `cty:"dns" hcl:"dns"`
	DnsSuffix    []string `cty:"dns_suffix" hcl:"dns_suffix"`
	Gateway      *string  `cty:"gateway" hcl:"gateway"`
	Ip           *string  `cty:"ip" hcl:"ip"`
	Netmask      *string  `cty:"netmask" hcl:"netmask"`
	Network      *string  `cty:"network" hcl:"network"`
	IpConfigType *string  `cty:"type" hcl:"type"`
}

type DeviceSwitchLocalPortConfigValue struct {
	AllNetworks                              *bool                          `cty:"all_networks" hcl:"all_networks"`
	AllowDhcpd                               *bool                          `cty:"allow_dhcpd" hcl:"allow_dhcpd"`
	AllowMultipleSupplicants                 *bool                          `cty:"allow_multiple_supplicants" hcl:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                 *bool                          `cty:"bypass_auth_when_server_down" hcl:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnknownClient *bool                          `cty:"bypass_auth_when_server_down_for_unknown_client" hcl:"bypass_auth_when_server_down_for_unknown_client"`
	Description                              *string                        `cty:"description" hcl:"description"`
	DisableAutoneg                           *bool                          `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled                                 *bool                          `cty:"disabled" hcl:"disabled"`
	Duplex                                   *string                        `cty:"duplex" hcl:"duplex"`
	DynamicVlanNetworks                      []string                       `cty:"dynamic_vlan_networks" hcl:"dynamic_vlan_networks"`
	EnableMacAuth                            *bool                          `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	EnableQos                                *bool                          `cty:"enable_qos" hcl:"enable_qos"`
	GuestNetwork                             *string                        `cty:"guest_network" hcl:"guest_network"`
	InterSwitchLink                          *bool                          `cty:"inter_switch_link" hcl:"inter_switch_link"`
	MacAuthOnly                              *bool                          `cty:"mac_auth_only" hcl:"mac_auth_only"`
	MacAuthPreferred                         *bool                          `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol                          *string                        `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MacLimit                                 *int64                         `cty:"mac_limit" hcl:"mac_limit"`
	Mode                                     *string                        `cty:"mode" hcl:"mode"`
	Mtu                                      *int64                         `cty:"mtu" hcl:"mtu"`
	Networks                                 []string                       `cty:"networks" hcl:"networks"`
	Note                                     *string                        `cty:"note" hcl:"note"`
	PersistMac                               *bool                          `cty:"persist_mac" hcl:"persist_mac"`
	PoeDisabled                              *bool                          `cty:"poe_disabled" hcl:"poe_disabled"`
	PortAuth                                 *string                        `cty:"port_auth" hcl:"port_auth"`
	PortNetwork                              *string                        `cty:"port_network" hcl:"port_network"`
	ReauthInterval                           *string                        `cty:"reauth_interval" hcl:"reauth_interval"`
	ServerFailNetwork                        *string                        `cty:"server_fail_network" hcl:"server_fail_network"`
	ServerRejectNetwork                      *string                        `cty:"server_reject_network" hcl:"server_reject_network"`
	Speed                                    *string                        `cty:"speed" hcl:"speed"`
	StormControl                             *DeviceSwitchStormControlValue `cty:"storm_control" hcl:"storm_control"`
	StpEdge                                  *bool                          `cty:"stp_edge" hcl:"stp_edge"`
	StpNoRootPort                            *bool                          `cty:"stp_no_root_port" hcl:"stp_no_root_port"`
	StpP2p                                   *bool                          `cty:"stp_p2p" hcl:"stp_p2p"`
	Usage                                    string                         `cty:"usage" hcl:"usage"`
	UseVstp                                  *bool                          `cty:"use_vstp" hcl:"use_vstp"`
	VoipNetwork                              *string                        `cty:"voip_network" hcl:"voip_network"`
}

type DeviceSwitchStormControlValue struct {
	DisablePort           *bool  `cty:"disable_port" hcl:"disable_port"`
	NoBroadcast           *bool  `cty:"no_broadcast" hcl:"no_broadcast"`
	NoMulticast           *bool  `cty:"no_multicast" hcl:"no_multicast"`
	NoRegisteredMulticast *bool  `cty:"no_registered_multicast" hcl:"no_registered_multicast"`
	NoUnknownUnicast      *bool  `cty:"no_unknown_unicast" hcl:"no_unknown_unicast"`
	Percentage            *int64 `cty:"percentage" hcl:"percentage"`
}

type DeviceSwitchMistNacValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Network *string `cty:"network" hcl:"network"`
}

type DeviceSwitchNetworksValue struct {
	Gateway         *string `cty:"gateway" hcl:"gateway"`
	Gateway6        *string `cty:"gateway6" hcl:"gateway6"`
	Isolation       *bool   `cty:"isolation" hcl:"isolation"`
	IsolationVlanId *string `cty:"isolation_vlan_id" hcl:"isolation_vlan_id"`
	Subnet          *string `cty:"subnet" hcl:"subnet"`
	Subnet6         *string `cty:"subnet6" hcl:"subnet6"`
	VlanId          string  `cty:"vlan_id" hcl:"vlan_id"`
}

type DeviceSwitchOobIpConfigValue struct {
	Gateway              *string `cty:"gateway" hcl:"gateway"`
	Ip                   *string `cty:"ip" hcl:"ip"`
	Netmask              *string `cty:"netmask" hcl:"netmask"`
	Network              *string `cty:"network" hcl:"network"`
	OobIpConfigType      *string `cty:"type" hcl:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf" hcl:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out" hcl:"use_mgmt_vrf_for_host_out"`
}

type DeviceSwitchOspfAreasValue struct {
	IncludeLoopback *bool                                    `cty:"include_loopback" hcl:"include_loopback"`
	OspfNetworks    map[string]DeviceSwitchOspfNetworksValue `cty:"networks" hcl:"networks"`
	OspfAreasType   *string                                  `cty:"type" hcl:"type"`
}

type DeviceSwitchOspfNetworksValue struct {
	AuthKeys               map[string]string `cty:"auth_keys" hcl:"auth_keys"`
	AuthPassword           *string           `cty:"auth_password" hcl:"auth_password"`
	AuthType               *string           `cty:"auth_type" hcl:"auth_type"`
	BfdMinimumInterval     *int64            `cty:"bfd_minimum_interval" hcl:"bfd_minimum_interval"`
	DeadInterval           *int64            `cty:"dead_interval" hcl:"dead_interval"`
	ExportPolicy           *string           `cty:"export_policy" hcl:"export_policy"`
	HelloInterval          *int64            `cty:"hello_interval" hcl:"hello_interval"`
	ImportPolicy           *string           `cty:"import_policy" hcl:"import_policy"`
	InterfaceType          *string           `cty:"interface_type" hcl:"interface_type"`
	Metric                 *int64            `cty:"metric" hcl:"metric"`
	NoReadvertiseToOverlay *bool             `cty:"no_readvertise_to_overlay" hcl:"no_readvertise_to_overlay"`
	Passive                *bool             `cty:"passive" hcl:"passive"`
}

type DeviceSwitchOspfConfigValue struct {
	Areas              map[string]DeviceSwitchAreasValue `cty:"areas" hcl:"areas"`
	Enabled            *bool                             `cty:"enabled" hcl:"enabled"`
	ExportPolicy       *string                           `cty:"export_policy" hcl:"export_policy"`
	ImportPolicy       *string                           `cty:"import_policy" hcl:"import_policy"`
	ReferenceBandwidth *string                           `cty:"reference_bandwidth" hcl:"reference_bandwidth"`
}

type DeviceSwitchAreasValue struct {
	NoSummary *bool `cty:"no_summary" hcl:"no_summary"`
}

type DeviceSwitchOtherIpConfigsValue struct {
	EvpnAnycast        *bool   `cty:"evpn_anycast" hcl:"evpn_anycast"`
	Ip                 *string `cty:"ip" hcl:"ip"`
	Ip6                *string `cty:"ip6" hcl:"ip6"`
	Netmask            *string `cty:"netmask" hcl:"netmask"`
	Netmask6           *string `cty:"netmask6" hcl:"netmask6"`
	OtherIpConfigsType *string `cty:"type" hcl:"type"`
	Type6              *string `cty:"type6" hcl:"type6"`
}

type DeviceSwitchPortConfigValue struct {
	AeDisableLacp    *bool    `cty:"ae_disable_lacp" hcl:"ae_disable_lacp"`
	AeIdx            *int64   `cty:"ae_idx" hcl:"ae_idx"`
	AeLacpSlow       *bool    `cty:"ae_lacp_slow" hcl:"ae_lacp_slow"`
	Aggregated       *bool    `cty:"aggregated" hcl:"aggregated"`
	Critical         *bool    `cty:"critical" hcl:"critical"`
	Description      *string  `cty:"description" hcl:"description"`
	DisableAutoneg   *bool    `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Duplex           *string  `cty:"duplex" hcl:"duplex"`
	DynamicUsage     *string  `cty:"dynamic_usage" hcl:"dynamic_usage"`
	Esilag           *bool    `cty:"esilag" hcl:"esilag"`
	Mtu              *int64   `cty:"mtu" hcl:"mtu"`
	Networks         []string `cty:"networks" hcl:"networks"`
	NoLocalOverwrite *bool    `cty:"no_local_overwrite" hcl:"no_local_overwrite"`
	PoeDisabled      *bool    `cty:"poe_disabled" hcl:"poe_disabled"`
	PortNetwork      *string  `cty:"port_network" hcl:"port_network"`
	Speed            *string  `cty:"speed" hcl:"speed"`
	Usage            string   `cty:"usage" hcl:"usage"`
}

type DeviceSwitchPortConfigOverwriteValue struct {
	Description *string `cty:"description" hcl:"description"`
	Disabled    *bool   `cty:"disabled" hcl:"disabled"`
	Duplex      *string `cty:"duplex" hcl:"duplex"`
	MacLimit    *string `cty:"mac_limit" hcl:"mac_limit"`
	PoeDisabled *bool   `cty:"poe_disabled" hcl:"poe_disabled"`
	PortNetwork *string `cty:"port_network" hcl:"port_network"`
	Speed       *string `cty:"speed" hcl:"speed"`
}

type DeviceSwitchPortMirroringValue struct {
	InputNetworksIngress []string `cty:"input_networks_ingress" hcl:"input_networks_ingress"`
	InputPortIdsEgress   []string `cty:"input_port_ids_egress" hcl:"input_port_ids_egress"`
	InputPortIdsIngress  []string `cty:"input_port_ids_ingress" hcl:"input_port_ids_ingress"`
	OutputIpAddress      *string  `cty:"output_ip_address" hcl:"output_ip_address"`
	OutputNetwork        *string  `cty:"output_network" hcl:"output_network"`
	OutputPortId         *string  `cty:"output_port_id" hcl:"output_port_id"`
}

type DeviceSwitchPortUsagesValue struct {
	AllNetworks                              *bool                          `cty:"all_networks" hcl:"all_networks"`
	AllowDhcpd                               *bool                          `cty:"allow_dhcpd" hcl:"allow_dhcpd"`
	AllowMultipleSupplicants                 *bool                          `cty:"allow_multiple_supplicants" hcl:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                 *bool                          `cty:"bypass_auth_when_server_down" hcl:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnknownClient *bool                          `cty:"bypass_auth_when_server_down_for_unknown_client" hcl:"bypass_auth_when_server_down_for_unknown_client"`
	BypassAuthWhenServerDownForVoip          *bool                          `cty:"bypass_auth_when_server_down_for_voip" hcl:"bypass_auth_when_server_down_for_voip"`
	CommunityVlanId                          *int64                         `cty:"community_vlan_id" hcl:"community_vlan_id"`
	Description                              *string                        `cty:"description" hcl:"description"`
	DisableAutoneg                           *bool                          `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled                                 *bool                          `cty:"disabled" hcl:"disabled"`
	Duplex                                   *string                        `cty:"duplex" hcl:"duplex"`
	DynamicVlanNetworks                      []string                       `cty:"dynamic_vlan_networks" hcl:"dynamic_vlan_networks"`
	EnableMacAuth                            *bool                          `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	EnableQos                                *bool                          `cty:"enable_qos" hcl:"enable_qos"`
	GuestNetwork                             *string                        `cty:"guest_network" hcl:"guest_network"`
	InterIsolationNetworkLink                *bool                          `cty:"inter_isolation_network_link" hcl:"inter_isolation_network_link"`
	InterSwitchLink                          *bool                          `cty:"inter_switch_link" hcl:"inter_switch_link"`
	MacAuthOnly                              *bool                          `cty:"mac_auth_only" hcl:"mac_auth_only"`
	MacAuthPreferred                         *bool                          `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol                          *string                        `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MacLimit                                 *string                        `cty:"mac_limit" hcl:"mac_limit"`
	Mode                                     *string                        `cty:"mode" hcl:"mode"`
	Mtu                                      *string                        `cty:"mtu" hcl:"mtu"`
	Networks                                 []string                       `cty:"networks" hcl:"networks"`
	PersistMac                               *bool                          `cty:"persist_mac" hcl:"persist_mac"`
	PoeDisabled                              *bool                          `cty:"poe_disabled" hcl:"poe_disabled"`
	PoePriority                              *string                        `cty:"poe_priority" hcl:"poe_priority"`
	PortAuth                                 *string                        `cty:"port_auth" hcl:"port_auth"`
	PortNetwork                              *string                        `cty:"port_network" hcl:"port_network"`
	ReauthInterval                           *string                        `cty:"reauth_interval" hcl:"reauth_interval"`
	ResetDefaultWhen                         *string                        `cty:"reset_default_when" hcl:"reset_default_when"`
	Rules                                    []DeviceSwitchRulesValue       `cty:"rules" hcl:"rules"`
	ServerFailNetwork                        *string                        `cty:"server_fail_network" hcl:"server_fail_network"`
	ServerRejectNetwork                      *string                        `cty:"server_reject_network" hcl:"server_reject_network"`
	Speed                                    *string                        `cty:"speed" hcl:"speed"`
	StormControl                             *DeviceSwitchStormControlValue `cty:"storm_control" hcl:"storm_control"`
	StpDisable                               *bool                          `cty:"stp_disable" hcl:"stp_disable"`
	StpEdge                                  *bool                          `cty:"stp_edge" hcl:"stp_edge"`
	StpNoRootPort                            *bool                          `cty:"stp_no_root_port" hcl:"stp_no_root_port"`
	StpP2p                                   *bool                          `cty:"stp_p2p" hcl:"stp_p2p"`
	StpRequired                              *bool                          `cty:"stp_required" hcl:"stp_required"`
	UseVstp                                  *bool                          `cty:"use_vstp" hcl:"use_vstp"`
	VoipNetwork                              *string                        `cty:"voip_network" hcl:"voip_network"`
}

type DeviceSwitchRulesValue struct {
	Equals     *string  `cty:"equals" hcl:"equals"`
	EqualsAny  []string `cty:"equals_any" hcl:"equals_any"`
	Expression *string  `cty:"expression" hcl:"expression"`
	Src        string   `cty:"src" hcl:"src"`
	Usage      *string  `cty:"usage" hcl:"usage"`
}

type DeviceSwitchRadiusConfigValue struct {
	AcctImmediateUpdate *bool                          `cty:"acct_immediate_update" hcl:"acct_immediate_update"`
	AcctInterimInterval *int64                         `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AcctServers         []DeviceSwitchAcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServerSelection *string                        `cty:"auth_server_selection" hcl:"auth_server_selection"`
	AuthServers         []DeviceSwitchAuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	AuthServersRetries  *int64                         `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64                         `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool                          `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *string                        `cty:"coa_port" hcl:"coa_port"`
	FastDot1xTimers     *bool                          `cty:"fast_dot1x_timers" hcl:"fast_dot1x_timers"`
	Network             *string                        `cty:"network" hcl:"network"`
	SourceIp            *string                        `cty:"source_ip" hcl:"source_ip"`
}

type DeviceSwitchAcctServersValue struct {
	Host           string  `cty:"host" hcl:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port           *string `cty:"port" hcl:"port"`
	Secret         string  `cty:"secret" hcl:"secret"`
}

type DeviceSwitchAuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type DeviceSwitchRemoteSyslogValue struct {
	Archive          *DeviceSwitchArchiveValue  `cty:"archive" hcl:"archive"`
	Cacerts          []string                   `cty:"cacerts" hcl:"cacerts"`
	Console          *DeviceSwitchConsoleValue  `cty:"console" hcl:"console"`
	Enabled          *bool                      `cty:"enabled" hcl:"enabled"`
	Files            []DeviceSwitchFilesValue   `cty:"files" hcl:"files"`
	Network          *string                    `cty:"network" hcl:"network"`
	SendToAllServers *bool                      `cty:"send_to_all_servers" hcl:"send_to_all_servers"`
	Servers          []DeviceSwitchServersValue `cty:"servers" hcl:"servers"`
	TimeFormat       *string                    `cty:"time_format" hcl:"time_format"`
	Users            []DeviceSwitchUsersValue   `cty:"users" hcl:"users"`
}

type DeviceSwitchArchiveValue struct {
	Files *string `cty:"files" hcl:"files"`
	Size  *string `cty:"size" hcl:"size"`
}

type DeviceSwitchConsoleValue struct {
	Contents []DeviceSwitchContentsValue `cty:"contents" hcl:"contents"`
}

type DeviceSwitchContentsValue struct {
	Facility *string `cty:"facility" hcl:"facility"`
	Severity *string `cty:"severity" hcl:"severity"`
}

type DeviceSwitchFilesValue struct {
	Archive          *DeviceSwitchArchiveValue   `cty:"archive" hcl:"archive"`
	Contents         []DeviceSwitchContentsValue `cty:"contents" hcl:"contents"`
	EnableTls        *bool                       `cty:"enable_tls" hcl:"enable_tls"`
	ExplicitPriority *bool                       `cty:"explicit_priority" hcl:"explicit_priority"`
	File             *string                     `cty:"file" hcl:"file"`
	Match            *string                     `cty:"match" hcl:"match"`
	StructuredData   *bool                       `cty:"structured_data" hcl:"structured_data"`
}

type DeviceSwitchServersValue struct {
	Contents         []DeviceSwitchContentsValue `cty:"contents" hcl:"contents"`
	ExplicitPriority *bool                       `cty:"explicit_priority" hcl:"explicit_priority"`
	Facility         *string                     `cty:"facility" hcl:"facility"`
	Host             *string                     `cty:"host" hcl:"host"`
	Match            *string                     `cty:"match" hcl:"match"`
	Port             *string                     `cty:"port" hcl:"port"`
	Protocol         *string                     `cty:"protocol" hcl:"protocol"`
	RoutingInstance  *string                     `cty:"routing_instance" hcl:"routing_instance"`
	ServerName       *string                     `cty:"server_name" hcl:"server_name"`
	Severity         *string                     `cty:"severity" hcl:"severity"`
	SourceAddress    *string                     `cty:"source_address" hcl:"source_address"`
	StructuredData   *bool                       `cty:"structured_data" hcl:"structured_data"`
	Tag              *string                     `cty:"tag" hcl:"tag"`
}

type DeviceSwitchUsersValue struct {
	Contents []DeviceSwitchContentsValue `cty:"contents" hcl:"contents"`
	Match    *string                     `cty:"match" hcl:"match"`
	User     *string                     `cty:"user" hcl:"user"`
}

type DeviceSwitchSnmpConfigValue struct {
	ClientList   []DeviceSwitchClientListValue `cty:"client_list" hcl:"client_list"`
	Contact      *string                       `cty:"contact" hcl:"contact"`
	Description  *string                       `cty:"description" hcl:"description"`
	Enabled      *bool                         `cty:"enabled" hcl:"enabled"`
	EngineId     *string                       `cty:"engine_id" hcl:"engine_id"`
	EngineIdType *string                       `cty:"engine_id_type" hcl:"engine_id_type"`
	Location     *string                       `cty:"location" hcl:"location"`
	Name         *string                       `cty:"name" hcl:"name"`
	Network      *string                       `cty:"network" hcl:"network"`
	TrapGroups   []DeviceSwitchTrapGroupsValue `cty:"trap_groups" hcl:"trap_groups"`
	V2cConfig    []DeviceSwitchV2cConfigValue  `cty:"v2c_config" hcl:"v2c_config"`
	V3Config     *DeviceSwitchV3ConfigValue    `cty:"v3_config" hcl:"v3_config"`
	Views        []DeviceSwitchViewsValue      `cty:"views" hcl:"views"`
}

type DeviceSwitchClientListValue struct {
	ClientListName *string  `cty:"client_list_name" hcl:"client_list_name"`
	Clients        []string `cty:"clients" hcl:"clients"`
}

type DeviceSwitchTrapGroupsValue struct {
	Categories []string `cty:"categories" hcl:"categories"`
	GroupName  *string  `cty:"group_name" hcl:"group_name"`
	Targets    []string `cty:"targets" hcl:"targets"`
	Version    *string  `cty:"version" hcl:"version"`
}

type DeviceSwitchV2cConfigValue struct {
	Authorization  *string `cty:"authorization" hcl:"authorization"`
	ClientListName *string `cty:"client_list_name" hcl:"client_list_name"`
	CommunityName  *string `cty:"community_name" hcl:"community_name"`
	View           *string `cty:"view" hcl:"view"`
}

type DeviceSwitchV3ConfigValue struct {
	Notify           []DeviceSwitchNotifyValue           `cty:"notify" hcl:"notify"`
	NotifyFilter     []DeviceSwitchNotifyFilterValue     `cty:"notify_filter" hcl:"notify_filter"`
	TargetAddress    []DeviceSwitchTargetAddressValue    `cty:"target_address" hcl:"target_address"`
	TargetParameters []DeviceSwitchTargetParametersValue `cty:"target_parameters" hcl:"target_parameters"`
	Usm              []DeviceSwitchUsmValue              `cty:"usm" hcl:"usm"`
	Vacm             *DeviceSwitchVacmValue              `cty:"vacm" hcl:"vacm"`
}

type DeviceSwitchNotifyValue struct {
	Name       string `cty:"name" hcl:"name"`
	Tag        string `cty:"tag" hcl:"tag"`
	NotifyType string `cty:"type" hcl:"type"`
}

type DeviceSwitchNotifyFilterValue struct {
	ProfileName    *string                           `cty:"profile_name" hcl:"profile_name"`
	Snmpv3Contents []DeviceSwitchSnmpv3ContentsValue `cty:"contents" hcl:"contents"`
}

type DeviceSwitchSnmpv3ContentsValue struct {
	Include *bool  `cty:"include" hcl:"include"`
	Oid     string `cty:"oid" hcl:"oid"`
}

type DeviceSwitchTargetAddressValue struct {
	Address           string  `cty:"address" hcl:"address"`
	AddressMask       string  `cty:"address_mask" hcl:"address_mask"`
	Port              *string `cty:"port" hcl:"port"`
	TagList           *string `cty:"tag_list" hcl:"tag_list"`
	TargetAddressName string  `cty:"target_address_name" hcl:"target_address_name"`
	TargetParameters  *string `cty:"target_parameters" hcl:"target_parameters"`
}

type DeviceSwitchTargetParametersValue struct {
	MessageProcessingModel string  `cty:"message_processing_model" hcl:"message_processing_model"`
	Name                   string  `cty:"name" hcl:"name"`
	NotifyFilter           *string `cty:"notify_filter" hcl:"notify_filter"`
	SecurityLevel          *string `cty:"security_level" hcl:"security_level"`
	SecurityModel          *string `cty:"security_model" hcl:"security_model"`
	SecurityName           *string `cty:"security_name" hcl:"security_name"`
}

type DeviceSwitchUsmValue struct {
	EngineType     string                         `cty:"engine_type" hcl:"engine_type"`
	RemoteEngineId *string                        `cty:"remote_engine_id" hcl:"remote_engine_id"`
	Snmpv3Users    []DeviceSwitchSnmpv3UsersValue `cty:"users" hcl:"users"`
}

type DeviceSwitchSnmpv3UsersValue struct {
	AuthenticationPassword *string `cty:"authentication_password" hcl:"authentication_password"`
	AuthenticationType     *string `cty:"authentication_type" hcl:"authentication_type"`
	EncryptionPassword     *string `cty:"encryption_password" hcl:"encryption_password"`
	EncryptionType         *string `cty:"encryption_type" hcl:"encryption_type"`
	Name                   *string `cty:"name" hcl:"name"`
}

type DeviceSwitchVacmValue struct {
	Access          []DeviceSwitchAccessValue         `cty:"access" hcl:"access"`
	SecurityToGroup *DeviceSwitchSecurityToGroupValue `cty:"security_to_group" hcl:"security_to_group"`
}

type DeviceSwitchAccessValue struct {
	GroupName  *string                       `cty:"group_name" hcl:"group_name"`
	PrefixList []DeviceSwitchPrefixListValue `cty:"prefix_list" hcl:"prefix_list"`
}

type DeviceSwitchPrefixListValue struct {
	ContextPrefix  *string `cty:"context_prefix" hcl:"context_prefix"`
	NotifyView     *string `cty:"notify_view" hcl:"notify_view"`
	ReadView       *string `cty:"read_view" hcl:"read_view"`
	SecurityLevel  *string `cty:"security_level" hcl:"security_level"`
	SecurityModel  *string `cty:"security_model" hcl:"security_model"`
	PrefixListType *string `cty:"type" hcl:"type"`
	WriteView      *string `cty:"write_view" hcl:"write_view"`
}

type DeviceSwitchSecurityToGroupValue struct {
	SecurityModel     *string                              `cty:"security_model" hcl:"security_model"`
	Snmpv3VacmContent []DeviceSwitchSnmpv3VacmContentValue `cty:"content" hcl:"content"`
}

type DeviceSwitchSnmpv3VacmContentValue struct {
	Group        *string `cty:"group" hcl:"group"`
	SecurityName *string `cty:"security_name" hcl:"security_name"`
}

type DeviceSwitchViewsValue struct {
	Include  *bool   `cty:"include" hcl:"include"`
	Oid      *string `cty:"oid" hcl:"oid"`
	ViewName *string `cty:"view_name" hcl:"view_name"`
}

type DeviceSwitchStpConfigValue struct {
	BridgePriority *string `cty:"bridge_priority" hcl:"bridge_priority"`
}

type DeviceSwitchSwitchMgmtValue struct {
	ApAffinityThreshold   *int64                                    `cty:"ap_affinity_threshold" hcl:"ap_affinity_threshold"`
	CliBanner             *string                                   `cty:"cli_banner" hcl:"cli_banner"`
	CliIdleTimeout        *int64                                    `cty:"cli_idle_timeout" hcl:"cli_idle_timeout"`
	ConfigRevertTimer     *int64                                    `cty:"config_revert_timer" hcl:"config_revert_timer"`
	DhcpOptionFqdn        *bool                                     `cty:"dhcp_option_fqdn" hcl:"dhcp_option_fqdn"`
	DisableOobDownAlarm   *bool                                     `cty:"disable_oob_down_alarm" hcl:"disable_oob_down_alarm"`
	FipsEnabled           *bool                                     `cty:"fips_enabled" hcl:"fips_enabled"`
	LocalAccounts         map[string]DeviceSwitchLocalAccountsValue `cty:"local_accounts" hcl:"local_accounts"`
	MxedgeProxyHost       *string                                   `cty:"mxedge_proxy_host" hcl:"mxedge_proxy_host"`
	MxedgeProxyPort       *string                                   `cty:"mxedge_proxy_port" hcl:"mxedge_proxy_port"`
	ProtectRe             *DeviceSwitchProtectReValue               `cty:"protect_re" hcl:"protect_re"`
	RemoveExistingConfigs *bool                                     `cty:"remove_existing_configs" hcl:"remove_existing_configs"`
	RootPassword          *string                                   `cty:"root_password" hcl:"root_password"`
	Tacacs                *DeviceSwitchTacacsValue                  `cty:"tacacs" hcl:"tacacs"`
	UseMxedgeProxy        *bool                                     `cty:"use_mxedge_proxy" hcl:"use_mxedge_proxy"`
}

type DeviceSwitchLocalAccountsValue struct {
	Password *string `cty:"password" hcl:"password"`
	Role     *string `cty:"role" hcl:"role"`
}

type DeviceSwitchProtectReValue struct {
	AllowedServices []string                  `cty:"allowed_services" hcl:"allowed_services"`
	Custom          []DeviceSwitchCustomValue `cty:"custom" hcl:"custom"`
	Enabled         *bool                     `cty:"enabled" hcl:"enabled"`
	HitCount        *bool                     `cty:"hit_count" hcl:"hit_count"`
	TrustedHosts    []string                  `cty:"trusted_hosts" hcl:"trusted_hosts"`
}

type DeviceSwitchCustomValue struct {
	PortRange *string  `cty:"port_range" hcl:"port_range"`
	Protocol  *string  `cty:"protocol" hcl:"protocol"`
	Subnets   []string `cty:"subnets" hcl:"subnets"`
}

type DeviceSwitchTacacsValue struct {
	DefaultRole    *string                           `cty:"default_role" hcl:"default_role"`
	Enabled        *bool                             `cty:"enabled" hcl:"enabled"`
	Network        *string                           `cty:"network" hcl:"network"`
	TacacctServers []DeviceSwitchTacacctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	TacplusServers []DeviceSwitchTacplusServersValue `cty:"tacplus_servers" hcl:"tacplus_servers"`
}

type DeviceSwitchTacacctServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type DeviceSwitchTacplusServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type DeviceSwitchVirtualChassisValue struct {
	Members        []DeviceSwitchMembersValue `cty:"members" hcl:"members"`
	Preprovisioned *bool                      `cty:"preprovisioned" hcl:"preprovisioned"`
}

type DeviceSwitchMembersValue struct {
	Mac      string `cty:"mac" hcl:"mac"`
	MemberId int64  `cty:"member_id" hcl:"member_id"`
	VcRole   string `cty:"vc_role" hcl:"vc_role"`
}

type DeviceSwitchVrfConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type DeviceSwitchVrfInstancesValue struct {
	EvpnAutoLoopbackSubnet  *string                                     `cty:"evpn_auto_loopback_subnet" hcl:"evpn_auto_loopback_subnet"`
	EvpnAutoLoopbackSubnet6 *string                                     `cty:"evpn_auto_loopback_subnet6" hcl:"evpn_auto_loopback_subnet6"`
	Networks                []string                                    `cty:"networks" hcl:"networks"`
	VrfExtraRoutes          map[string]DeviceSwitchVrfExtraRoutesValue  `cty:"extra_routes" hcl:"extra_routes"`
	VrfExtraRoutes6         map[string]DeviceSwitchVrfExtraRoutes6Value `cty:"extra_routes6" hcl:"extra_routes6"`
}

type DeviceSwitchVrfExtraRoutesValue struct {
	Via string `cty:"via" hcl:"via"`
}

type DeviceSwitchVrfExtraRoutes6Value struct {
	Via *string `cty:"via" hcl:"via"`
}

type DeviceSwitchVrrpConfigValue struct {
	Enabled *bool                              `cty:"enabled" hcl:"enabled"`
	Groups  map[string]DeviceSwitchGroupsValue `cty:"groups" hcl:"groups"`
}

type DeviceSwitchGroupsValue struct {
	Preempt  *bool  `cty:"preempt" hcl:"preempt"`
	Priority *int64 `cty:"priority" hcl:"priority"`
}
