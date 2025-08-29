package provider

type DeviceSwitchModel struct {
	AclPolicies           []AclPoliciesValue                  `hcl:"acl_policies"`
	AclTags               map[string]AclTagsValue             `hcl:"acl_tags"`
	AdditionalConfigCmds  []string                            `hcl:"additional_config_cmds"`
	DeviceId              string                              `hcl:"device_id"`
	DhcpSnooping          *DhcpSnoopingValue                  `hcl:"dhcp_snooping"`
	DhcpdConfig           *DhcpdConfigValue                   `hcl:"dhcpd_config"`
	DisableAutoConfig     *bool                               `hcl:"disable_auto_config"`
	DnsServers            []string                            `hcl:"dns_servers"`
	DnsSuffix             []string                            `hcl:"dns_suffix"`
	ExtraRoutes           map[string]ExtraRoutesValue         `hcl:"extra_routes"`
	ExtraRoutes6          map[string]ExtraRoutes6Value        `hcl:"extra_routes6"`
	IpConfig              *DeviceSwitchIpConfigValue          `hcl:"ip_config"`
	LocalPortConfig       map[string]LocalPortConfigValue     `hcl:"local_port_config"`
	Managed               *bool                               `hcl:"managed"`
	MapId                 *string                             `hcl:"map_id"`
	MistNac               *MistNacValue                       `hcl:"mist_nac"`
	Name                  string                              `hcl:"name"`
	Networks              map[string]NetworksValue            `hcl:"networks"`
	Notes                 *string                             `hcl:"notes"`
	NtpServers            []string                            `hcl:"ntp_servers"`
	OobIpConfig           *OobIpConfigValue                   `hcl:"oob_ip_config"`
	OspfAreas             map[string]OspfAreasValue           `hcl:"ospf_areas"`
	OspfConfig            *OspfConfigValue                    `hcl:"ospf_config"`
	OtherIpConfigs        map[string]OtherIpConfigsValue      `hcl:"other_ip_configs"`
	PortConfig            map[string]PortConfigValue          `hcl:"port_config"`
	PortConfigOverwrite   map[string]PortConfigOverwriteValue `hcl:"port_config_overwrite"`
	PortMirroring         map[string]PortMirroringValue       `hcl:"port_mirroring"`
	PortUsages            map[string]PortUsagesValue          `hcl:"port_usages"`
	RadiusConfig          *RadiusConfigValue                  `hcl:"radius_config"`
	RemoteSyslog          *RemoteSyslogValue                  `hcl:"remote_syslog"`
	Role                  *string                             `hcl:"role"`
	RouterId              *string                             `hcl:"router_id"`
	SiteId                string                              `hcl:"site_id"`
	SnmpConfig            *SnmpConfigValue                    `hcl:"snmp_config"`
	StpConfig             *StpConfigValue                     `hcl:"stp_config"`
	SwitchMgmt            *SwitchMgmtValue                    `hcl:"switch_mgmt"`
	UseRouterIdAsSourceIp *bool                               `hcl:"use_router_id_as_source_ip"`
	Vars                  map[string]string                   `hcl:"vars"`
	VirtualChassis        *VirtualChassisValue                `hcl:"virtual_chassis"`
	VrfConfig             *VrfConfigValue                     `hcl:"vrf_config"`
	VrfInstances          map[string]VrfInstancesValue        `hcl:"vrf_instances"`
	VrrpConfig            *VrrpConfigValue                    `hcl:"vrrp_config"`
	X                     *float64                            `hcl:"x"`
	Y                     *float64                            `hcl:"y"`
}

type AclPoliciesValue struct {
	Actions []ActionsValue `cty:"actions" hcl:"actions"`
	Name    *string        `cty:"name" hcl:"name"`
	SrcTags []string       `cty:"src_tags" hcl:"src_tags"`
}

type ActionsValue struct {
	Action *string `cty:"action" hcl:"action"`
	DstTag string  `cty:"dst_tag" hcl:"dst_tag"`
}

type AclTagsValue struct {
	EtherTypes  []string     `cty:"ether_types" hcl:"ether_types"`
	GbpTag      *int64       `cty:"gbp_tag" hcl:"gbp_tag"`
	Macs        []string     `cty:"macs" hcl:"macs"`
	Network     *string      `cty:"network" hcl:"network"`
	PortUsage   *string      `cty:"port_usage" hcl:"port_usage"`
	RadiusGroup *string      `cty:"radius_group" hcl:"radius_group"`
	Specs       []SpecsValue `cty:"specs" hcl:"specs"`
	Subnets     []string     `cty:"subnets" hcl:"subnets"`
	AclTagsType string       `cty:"type" hcl:"type"`
}

type SpecsValue struct {
	PortRange *string `cty:"port_range" hcl:"port_range"`
	Protocol  *string `cty:"protocol" hcl:"protocol"`
}

type DhcpSnoopingValue struct {
	AllNetworks         *bool    `cty:"all_networks" hcl:"all_networks"`
	EnableArpSpoofCheck *bool    `cty:"enable_arp_spoof_check" hcl:"enable_arp_spoof_check"`
	EnableIpSourceGuard *bool    `cty:"enable_ip_source_guard" hcl:"enable_ip_source_guard"`
	Enabled             *bool    `cty:"enabled" hcl:"enabled"`
	Networks            []string `cty:"networks" hcl:"networks"`
}

type DhcpdConfigValue struct {
	Config  map[string]ConfigValue `cty:"config" hcl:"config"`
	Enabled *bool                  `cty:"enabled" hcl:"enabled"`
}

type ConfigValue struct {
	DnsServers         []string                           `cty:"dns_servers" hcl:"dns_servers"`
	DnsSuffix          []string                           `cty:"dns_suffix" hcl:"dns_suffix"`
	FixedBindings      map[string]FixedBindingsValue      `cty:"fixed_bindings" hcl:"fixed_bindings"`
	Gateway            *string                            `cty:"gateway" hcl:"gateway"`
	IpEnd              *string                            `cty:"ip_end" hcl:"ip_end"`
	IpEnd6             *string                            `cty:"ip_end6" hcl:"ip_end6"`
	IpStart            *string                            `cty:"ip_start" hcl:"ip_start"`
	IpStart6           *string                            `cty:"ip_start6" hcl:"ip_start6"`
	LeaseTime          *int64                             `cty:"lease_time" hcl:"lease_time"`
	Options            map[string]OptionsValue            `cty:"options" hcl:"options"`
	ServerIdOverride   *bool                              `cty:"server_id_override" hcl:"server_id_override"`
	Servers            []string                           `cty:"servers" hcl:"servers"`
	Servers6           []string                           `cty:"servers6" hcl:"servers6"`
	ConfigType         *string                            `cty:"type" hcl:"type"`
	Type6              *string                            `cty:"type6" hcl:"type6"`
	VendorEncapsulated map[string]VendorEncapsulatedValue `cty:"vendor_encapsulated" hcl:"vendor_encapsulated"`
}

type FixedBindingsValue struct {
	Ip   string  `cty:"ip" hcl:"ip"`
	Name *string `cty:"name" hcl:"name"`
}

type OptionsValue struct {
	OptionsType *string `cty:"type" hcl:"type"`
	Value       *string `cty:"value" hcl:"value"`
}

type VendorEncapsulatedValue struct {
	VendorEncapsulatedType *string `cty:"type" hcl:"type"`
	Value                  *string `cty:"value" hcl:"value"`
}

type ExtraRoutesValue struct {
	Discard       *bool                         `cty:"discard" hcl:"discard"`
	Metric        *int64                        `cty:"metric" hcl:"metric"`
	NextQualified map[string]NextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                         `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                        `cty:"preference" hcl:"preference"`
	Via           string                        `cty:"via" hcl:"via"`
}

type NextQualifiedValue struct {
	Metric     *int64 `cty:"metric" hcl:"metric"`
	Preference *int64 `cty:"preference" hcl:"preference"`
}

type ExtraRoutes6Value struct {
	Discard       *bool                         `cty:"discard" hcl:"discard"`
	Metric        *int64                        `cty:"metric" hcl:"metric"`
	NextQualified map[string]NextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                         `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                        `cty:"preference" hcl:"preference"`
	Via           string                        `cty:"via" hcl:"via"`
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

type LocalPortConfigValue struct {
	AllNetworks                              *bool              `cty:"all_networks" hcl:"all_networks"`
	AllowDhcpd                               *bool              `cty:"allow_dhcpd" hcl:"allow_dhcpd"`
	AllowMultipleSupplicants                 *bool              `cty:"allow_multiple_supplicants" hcl:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                 *bool              `cty:"bypass_auth_when_server_down" hcl:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnknownClient *bool              `cty:"bypass_auth_when_server_down_for_unknown_client" hcl:"bypass_auth_when_server_down_for_unknown_client"`
	Description                              *string            `cty:"description" hcl:"description"`
	DisableAutoneg                           *bool              `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled                                 *bool              `cty:"disabled" hcl:"disabled"`
	Duplex                                   *string            `cty:"duplex" hcl:"duplex"`
	DynamicVlanNetworks                      []string           `cty:"dynamic_vlan_networks" hcl:"dynamic_vlan_networks"`
	EnableMacAuth                            *bool              `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	EnableQos                                *bool              `cty:"enable_qos" hcl:"enable_qos"`
	GuestNetwork                             *string            `cty:"guest_network" hcl:"guest_network"`
	InterSwitchLink                          *bool              `cty:"inter_switch_link" hcl:"inter_switch_link"`
	MacAuthOnly                              *bool              `cty:"mac_auth_only" hcl:"mac_auth_only"`
	MacAuthPreferred                         *bool              `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol                          *string            `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MacLimit                                 *int64             `cty:"mac_limit" hcl:"mac_limit"`
	Mode                                     *string            `cty:"mode" hcl:"mode"`
	Mtu                                      *int64             `cty:"mtu" hcl:"mtu"`
	Networks                                 []string           `cty:"networks" hcl:"networks"`
	Note                                     *string            `cty:"note" hcl:"note"`
	PersistMac                               *bool              `cty:"persist_mac" hcl:"persist_mac"`
	PoeDisabled                              *bool              `cty:"poe_disabled" hcl:"poe_disabled"`
	PortAuth                                 *string            `cty:"port_auth" hcl:"port_auth"`
	PortNetwork                              *string            `cty:"port_network" hcl:"port_network"`
	ReauthInterval                           *string            `cty:"reauth_interval" hcl:"reauth_interval"`
	ServerFailNetwork                        *string            `cty:"server_fail_network" hcl:"server_fail_network"`
	ServerRejectNetwork                      *string            `cty:"server_reject_network" hcl:"server_reject_network"`
	Speed                                    *string            `cty:"speed" hcl:"speed"`
	StormControl                             *StormControlValue `cty:"storm_control" hcl:"storm_control"`
	StpEdge                                  *bool              `cty:"stp_edge" hcl:"stp_edge"`
	StpNoRootPort                            *bool              `cty:"stp_no_root_port" hcl:"stp_no_root_port"`
	StpP2p                                   *bool              `cty:"stp_p2p" hcl:"stp_p2p"`
	Usage                                    string             `cty:"usage" hcl:"usage"`
	UseVstp                                  *bool              `cty:"use_vstp" hcl:"use_vstp"`
	VoipNetwork                              *string            `cty:"voip_network" hcl:"voip_network"`
}

type StormControlValue struct {
	DisablePort           *bool  `cty:"disable_port" hcl:"disable_port"`
	NoBroadcast           *bool  `cty:"no_broadcast" hcl:"no_broadcast"`
	NoMulticast           *bool  `cty:"no_multicast" hcl:"no_multicast"`
	NoRegisteredMulticast *bool  `cty:"no_registered_multicast" hcl:"no_registered_multicast"`
	NoUnknownUnicast      *bool  `cty:"no_unknown_unicast" hcl:"no_unknown_unicast"`
	Percentage            *int64 `cty:"percentage" hcl:"percentage"`
}

type MistNacValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Network *string `cty:"network" hcl:"network"`
}

type NetworksValue struct {
	Gateway         *string `cty:"gateway" hcl:"gateway"`
	Gateway6        *string `cty:"gateway6" hcl:"gateway6"`
	Isolation       *bool   `cty:"isolation" hcl:"isolation"`
	IsolationVlanId *string `cty:"isolation_vlan_id" hcl:"isolation_vlan_id"`
	Subnet          *string `cty:"subnet" hcl:"subnet"`
	Subnet6         *string `cty:"subnet6" hcl:"subnet6"`
	VlanId          string  `cty:"vlan_id" hcl:"vlan_id"`
}

type OobIpConfigValue struct {
	Gateway              *string `cty:"gateway" hcl:"gateway"`
	Ip                   *string `cty:"ip" hcl:"ip"`
	Netmask              *string `cty:"netmask" hcl:"netmask"`
	Network              *string `cty:"network" hcl:"network"`
	OobIpConfigType      *string `cty:"type" hcl:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf" hcl:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out" hcl:"use_mgmt_vrf_for_host_out"`
}

type OspfAreasValue struct {
	IncludeLoopback *bool                        `cty:"include_loopback" hcl:"include_loopback"`
	OspfNetworks    map[string]OspfNetworksValue `cty:"networks" hcl:"networks"`
	OspfAreasType   *string                      `cty:"type" hcl:"type"`
}

type OspfNetworksValue struct {
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

type OspfConfigValue struct {
	Areas              map[string]AreasValue `cty:"areas" hcl:"areas"`
	Enabled            *bool                 `cty:"enabled" hcl:"enabled"`
	ExportPolicy       *string               `cty:"export_policy" hcl:"export_policy"`
	ImportPolicy       *string               `cty:"import_policy" hcl:"import_policy"`
	ReferenceBandwidth *string               `cty:"reference_bandwidth" hcl:"reference_bandwidth"`
}

type AreasValue struct {
	NoSummary *bool `cty:"no_summary" hcl:"no_summary"`
}

type OtherIpConfigsValue struct {
	EvpnAnycast        *bool   `cty:"evpn_anycast" hcl:"evpn_anycast"`
	Ip                 *string `cty:"ip" hcl:"ip"`
	Ip6                *string `cty:"ip6" hcl:"ip6"`
	Netmask            *string `cty:"netmask" hcl:"netmask"`
	Netmask6           *string `cty:"netmask6" hcl:"netmask6"`
	OtherIpConfigsType *string `cty:"type" hcl:"type"`
	Type6              *string `cty:"type6" hcl:"type6"`
}

type PortConfigValue struct {
	AeDisableLacp    *bool   `cty:"ae_disable_lacp" hcl:"ae_disable_lacp"`
	AeIdx            *int64  `cty:"ae_idx" hcl:"ae_idx"`
	AeLacpSlow       *bool   `cty:"ae_lacp_slow" hcl:"ae_lacp_slow"`
	Aggregated       *bool   `cty:"aggregated" hcl:"aggregated"`
	Critical         *bool   `cty:"critical" hcl:"critical"`
	Description      *string `cty:"description" hcl:"description"`
	DisableAutoneg   *bool   `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Duplex           *string `cty:"duplex" hcl:"duplex"`
	DynamicUsage     *string `cty:"dynamic_usage" hcl:"dynamic_usage"`
	Esilag           *bool   `cty:"esilag" hcl:"esilag"`
	Mtu              *int64  `cty:"mtu" hcl:"mtu"`
	NoLocalOverwrite *bool   `cty:"no_local_overwrite" hcl:"no_local_overwrite"`
	PoeDisabled      *bool   `cty:"poe_disabled" hcl:"poe_disabled"`
	PortNetwork      *string `cty:"port_network" hcl:"port_network"`
	Speed            *string `cty:"speed" hcl:"speed"`
	Usage            string  `cty:"usage" hcl:"usage"`
}

type PortConfigOverwriteValue struct {
	Description *string `cty:"description" hcl:"description"`
	Disabled    *bool   `cty:"disabled" hcl:"disabled"`
	Duplex      *string `cty:"duplex" hcl:"duplex"`
	MacLimit    *string `cty:"mac_limit" hcl:"mac_limit"`
	PoeDisabled *bool   `cty:"poe_disabled" hcl:"poe_disabled"`
	PortNetwork *string `cty:"port_network" hcl:"port_network"`
	Speed       *string `cty:"speed" hcl:"speed"`
}

type PortMirroringValue struct {
	InputNetworksIngress []string `cty:"input_networks_ingress" hcl:"input_networks_ingress"`
	InputPortIdsEgress   []string `cty:"input_port_ids_egress" hcl:"input_port_ids_egress"`
	InputPortIdsIngress  []string `cty:"input_port_ids_ingress" hcl:"input_port_ids_ingress"`
	OutputIpAddress      *string  `cty:"output_ip_address" hcl:"output_ip_address"`
	OutputNetwork        *string  `cty:"output_network" hcl:"output_network"`
	OutputPortId         *string  `cty:"output_port_id" hcl:"output_port_id"`
}

type PortUsagesValue struct {
	AllNetworks                              *bool              `cty:"all_networks" hcl:"all_networks"`
	AllowDhcpd                               *bool              `cty:"allow_dhcpd" hcl:"allow_dhcpd"`
	AllowMultipleSupplicants                 *bool              `cty:"allow_multiple_supplicants" hcl:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                 *bool              `cty:"bypass_auth_when_server_down" hcl:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnknownClient *bool              `cty:"bypass_auth_when_server_down_for_unknown_client" hcl:"bypass_auth_when_server_down_for_unknown_client"`
	CommunityVlanId                          *int64             `cty:"community_vlan_id" hcl:"community_vlan_id"`
	Description                              *string            `cty:"description" hcl:"description"`
	DisableAutoneg                           *bool              `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled                                 *bool              `cty:"disabled" hcl:"disabled"`
	Duplex                                   *string            `cty:"duplex" hcl:"duplex"`
	DynamicVlanNetworks                      []string           `cty:"dynamic_vlan_networks" hcl:"dynamic_vlan_networks"`
	EnableMacAuth                            *bool              `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	EnableQos                                *bool              `cty:"enable_qos" hcl:"enable_qos"`
	GuestNetwork                             *string            `cty:"guest_network" hcl:"guest_network"`
	InterIsolationNetworkLink                *bool              `cty:"inter_isolation_network_link" hcl:"inter_isolation_network_link"`
	InterSwitchLink                          *bool              `cty:"inter_switch_link" hcl:"inter_switch_link"`
	MacAuthOnly                              *bool              `cty:"mac_auth_only" hcl:"mac_auth_only"`
	MacAuthPreferred                         *bool              `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol                          *string            `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MacLimit                                 *string            `cty:"mac_limit" hcl:"mac_limit"`
	Mode                                     *string            `cty:"mode" hcl:"mode"`
	Mtu                                      *string            `cty:"mtu" hcl:"mtu"`
	Networks                                 []string           `cty:"networks" hcl:"networks"`
	PersistMac                               *bool              `cty:"persist_mac" hcl:"persist_mac"`
	PoeDisabled                              *bool              `cty:"poe_disabled" hcl:"poe_disabled"`
	PortAuth                                 *string            `cty:"port_auth" hcl:"port_auth"`
	PortNetwork                              *string            `cty:"port_network" hcl:"port_network"`
	ReauthInterval                           *string            `cty:"reauth_interval" hcl:"reauth_interval"`
	ResetDefaultWhen                         *string            `cty:"reset_default_when" hcl:"reset_default_when"`
	Rules                                    []RulesValue       `cty:"rules" hcl:"rules"`
	ServerFailNetwork                        *string            `cty:"server_fail_network" hcl:"server_fail_network"`
	ServerRejectNetwork                      *string            `cty:"server_reject_network" hcl:"server_reject_network"`
	Speed                                    *string            `cty:"speed" hcl:"speed"`
	StormControl                             *StormControlValue `cty:"storm_control" hcl:"storm_control"`
	StpEdge                                  *bool              `cty:"stp_edge" hcl:"stp_edge"`
	StpNoRootPort                            *bool              `cty:"stp_no_root_port" hcl:"stp_no_root_port"`
	StpP2p                                   *bool              `cty:"stp_p2p" hcl:"stp_p2p"`
	UseVstp                                  *bool              `cty:"use_vstp" hcl:"use_vstp"`
	VoipNetwork                              *string            `cty:"voip_network" hcl:"voip_network"`
}

type RulesValue struct {
	Equals     *string  `cty:"equals" hcl:"equals"`
	EqualsAny  []string `cty:"equals_any" hcl:"equals_any"`
	Expression *string  `cty:"expression" hcl:"expression"`
	Src        string   `cty:"src" hcl:"src"`
	Usage      *string  `cty:"usage" hcl:"usage"`
}

type RadiusConfigValue struct {
	AcctImmediateUpdate *bool              `cty:"acct_immediate_update" hcl:"acct_immediate_update"`
	AcctInterimInterval *int64             `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AcctServers         []AcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServerSelection *string            `cty:"auth_server_selection" hcl:"auth_server_selection"`
	AuthServers         []AuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	AuthServersRetries  *int64             `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64             `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool              `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *string            `cty:"coa_port" hcl:"coa_port"`
	FastDot1xTimers     *bool              `cty:"fast_dot1x_timers" hcl:"fast_dot1x_timers"`
	Network             *string            `cty:"network" hcl:"network"`
	SourceIp            *string            `cty:"source_ip" hcl:"source_ip"`
}

type AcctServersValue struct {
	Host                        string  `hcl:"host" cty:"host"`
	KeywrapEnabled              *bool   `hcl:"keywrap_enabled" cty:"keywrap_enabled"`
	KeywrapFormat               *string `hcl:"keywrap_format" cty:"keywrap_format"`
	KeywrapKek                  *string `hcl:"keywrap_kek" cty:"keywrap_kek"`
	KeywrapMack                 *string `hcl:"keywrap_mack" cty:"keywrap_mack"`
	Port                        *int64  `hcl:"port" cty:"port"`
	RequireMessageAuthenticator *bool   `hcl:"require_message_authenticator" cty:"require_message_authenticator"`
	Secret                      string  `hcl:"secret" cty:"secret"`
}

type AuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type RemoteSyslogValue struct {
	Archive          *ArchiveValue  `cty:"archive" hcl:"archive"`
	Cacerts          []string       `cty:"cacerts" hcl:"cacerts"`
	Console          *ConsoleValue  `cty:"console" hcl:"console"`
	Enabled          *bool          `cty:"enabled" hcl:"enabled"`
	Files            []FilesValue   `cty:"files" hcl:"files"`
	Network          *string        `cty:"network" hcl:"network"`
	SendToAllServers *bool          `cty:"send_to_all_servers" hcl:"send_to_all_servers"`
	Servers          []ServersValue `cty:"servers" hcl:"servers"`
	TimeFormat       *string        `cty:"time_format" hcl:"time_format"`
	Users            []UsersValue   `cty:"users" hcl:"users"`
}

type ArchiveValue struct {
	Files *string `cty:"files" hcl:"files"`
	Size  *string `cty:"size" hcl:"size"`
}

type ConsoleValue struct {
	Contents []ContentsValue `cty:"contents" hcl:"contents"`
}

type ContentsValue struct {
	Facility *string `cty:"facility" hcl:"facility"`
	Severity *string `cty:"severity" hcl:"severity"`
}

type FilesValue struct {
	Archive          *ArchiveValue   `cty:"archive" hcl:"archive"`
	Contents         []ContentsValue `cty:"contents" hcl:"contents"`
	EnableTls        *bool           `cty:"enable_tls" hcl:"enable_tls"`
	ExplicitPriority *bool           `cty:"explicit_priority" hcl:"explicit_priority"`
	File             *string         `cty:"file" hcl:"file"`
	Match            *string         `cty:"match" hcl:"match"`
	StructuredData   *bool           `cty:"structured_data" hcl:"structured_data"`
}

type ServersValue struct {
	Contents         []ContentsValue `cty:"contents" hcl:"contents"`
	ExplicitPriority *bool           `cty:"explicit_priority" hcl:"explicit_priority"`
	Facility         *string         `cty:"facility" hcl:"facility"`
	Host             *string         `cty:"host" hcl:"host"`
	Match            *string         `cty:"match" hcl:"match"`
	Port             *string         `cty:"port" hcl:"port"`
	Protocol         *string         `cty:"protocol" hcl:"protocol"`
	RoutingInstance  *string         `cty:"routing_instance" hcl:"routing_instance"`
	ServerName       *string         `cty:"server_name" hcl:"server_name"`
	Severity         *string         `cty:"severity" hcl:"severity"`
	SourceAddress    *string         `cty:"source_address" hcl:"source_address"`
	StructuredData   *bool           `cty:"structured_data" hcl:"structured_data"`
	Tag              *string         `cty:"tag" hcl:"tag"`
}

type UsersValue struct {
	Contents []ContentsValue `cty:"contents" hcl:"contents"`
	Match    *string         `cty:"match" hcl:"match"`
	User     *string         `cty:"user" hcl:"user"`
}

type SnmpConfigValue struct {
	ClientList   []ClientListValue `cty:"client_list" hcl:"client_list"`
	Contact      *string           `cty:"contact" hcl:"contact"`
	Description  *string           `cty:"description" hcl:"description"`
	Enabled      *bool             `cty:"enabled" hcl:"enabled"`
	EngineId     *string           `cty:"engine_id" hcl:"engine_id"`
	EngineIdType *string           `cty:"engine_id_type" hcl:"engine_id_type"`
	Location     *string           `cty:"location" hcl:"location"`
	Name         *string           `cty:"name" hcl:"name"`
	Network      *string           `cty:"network" hcl:"network"`
	TrapGroups   []TrapGroupsValue `cty:"trap_groups" hcl:"trap_groups"`
	V2cConfig    []V2cConfigValue  `cty:"v2c_config" hcl:"v2c_config"`
	V3Config     *V3ConfigValue    `cty:"v3_config" hcl:"v3_config"`
	Views        []ViewsValue      `cty:"views" hcl:"views"`
}

type ClientListValue struct {
	ClientListName *string  `cty:"client_list_name" hcl:"client_list_name"`
	Clients        []string `cty:"clients" hcl:"clients"`
}

type TrapGroupsValue struct {
	Categories []string `cty:"categories" hcl:"categories"`
	GroupName  *string  `cty:"group_name" hcl:"group_name"`
	Targets    []string `cty:"targets" hcl:"targets"`
	Version    *string  `cty:"version" hcl:"version"`
}

type V2cConfigValue struct {
	Authorization  *string `cty:"authorization" hcl:"authorization"`
	ClientListName *string `cty:"client_list_name" hcl:"client_list_name"`
	CommunityName  *string `cty:"community_name" hcl:"community_name"`
	View           *string `cty:"view" hcl:"view"`
}

type V3ConfigValue struct {
	Notify           []NotifyValue           `cty:"notify" hcl:"notify"`
	NotifyFilter     []NotifyFilterValue     `cty:"notify_filter" hcl:"notify_filter"`
	TargetAddress    []TargetAddressValue    `cty:"target_address" hcl:"target_address"`
	TargetParameters []TargetParametersValue `cty:"target_parameters" hcl:"target_parameters"`
	Usm              []UsmValue              `cty:"usm" hcl:"usm"`
	Vacm             *VacmValue              `cty:"vacm" hcl:"vacm"`
}

type NotifyValue struct {
	Name       string `cty:"name" hcl:"name"`
	Tag        string `cty:"tag" hcl:"tag"`
	NotifyType string `cty:"type" hcl:"type"`
}

type NotifyFilterValue struct {
	ProfileName    *string               `cty:"profile_name" hcl:"profile_name"`
	Snmpv3Contents []Snmpv3ContentsValue `cty:"contents" hcl:"contents"`
}

type Snmpv3ContentsValue struct {
	Include *bool  `cty:"include" hcl:"include"`
	Oid     string `cty:"oid" hcl:"oid"`
}

type TargetAddressValue struct {
	Address           string  `cty:"address" hcl:"address"`
	AddressMask       string  `cty:"address_mask" hcl:"address_mask"`
	Port              *string `cty:"port" hcl:"port"`
	TagList           *string `cty:"tag_list" hcl:"tag_list"`
	TargetAddressName string  `cty:"target_address_name" hcl:"target_address_name"`
	TargetParameters  *string `cty:"target_parameters" hcl:"target_parameters"`
}

type TargetParametersValue struct {
	MessageProcessingModel string  `cty:"message_processing_model" hcl:"message_processing_model"`
	Name                   string  `cty:"name" hcl:"name"`
	NotifyFilter           *string `cty:"notify_filter" hcl:"notify_filter"`
	SecurityLevel          *string `cty:"security_level" hcl:"security_level"`
	SecurityModel          *string `cty:"security_model" hcl:"security_model"`
	SecurityName           *string `cty:"security_name" hcl:"security_name"`
}

type UsmValue struct {
	EngineType     string             `cty:"engine_type" hcl:"engine_type"`
	RemoteEngineId *string            `cty:"remote_engine_id" hcl:"remote_engine_id"`
	Snmpv3Users    []Snmpv3UsersValue `cty:"users" hcl:"users"`
}

type Snmpv3UsersValue struct {
	AuthenticationPassword *string `cty:"authentication_password" hcl:"authentication_password"`
	AuthenticationType     *string `cty:"authentication_type" hcl:"authentication_type"`
	EncryptionPassword     *string `cty:"encryption_password" hcl:"encryption_password"`
	EncryptionType         *string `cty:"encryption_type" hcl:"encryption_type"`
	Name                   *string `cty:"name" hcl:"name"`
}

type VacmValue struct {
	Access          []AccessValue         `cty:"access" hcl:"access"`
	SecurityToGroup *SecurityToGroupValue `cty:"security_to_group" hcl:"security_to_group"`
}

type AccessValue struct {
	GroupName  *string           `cty:"group_name" hcl:"group_name"`
	PrefixList []PrefixListValue `cty:"prefix_list" hcl:"prefix_list"`
}

type PrefixListValue struct {
	ContextPrefix  *string `cty:"context_prefix" hcl:"context_prefix"`
	NotifyView     *string `cty:"notify_view" hcl:"notify_view"`
	ReadView       *string `cty:"read_view" hcl:"read_view"`
	SecurityLevel  *string `cty:"security_level" hcl:"security_level"`
	SecurityModel  *string `cty:"security_model" hcl:"security_model"`
	PrefixListType *string `cty:"type" hcl:"type"`
	WriteView      *string `cty:"write_view" hcl:"write_view"`
}

type SecurityToGroupValue struct {
	SecurityModel     *string                  `cty:"security_model" hcl:"security_model"`
	Snmpv3VacmContent []Snmpv3VacmContentValue `cty:"content" hcl:"content"`
}

type Snmpv3VacmContentValue struct {
	Group        *string `cty:"group" hcl:"group"`
	SecurityName *string `cty:"security_name" hcl:"security_name"`
}

type ViewsValue struct {
	Include  *bool   `cty:"include" hcl:"include"`
	Oid      *string `cty:"oid" hcl:"oid"`
	ViewName *string `cty:"view_name" hcl:"view_name"`
}

type StpConfigValue struct {
	BridgePriority *string `cty:"bridge_priority" hcl:"bridge_priority"`
}

type SwitchMgmtValue struct {
	ApAffinityThreshold   *int64                        `cty:"ap_affinity_threshold" hcl:"ap_affinity_threshold"`
	CliBanner             *string                       `cty:"cli_banner" hcl:"cli_banner"`
	CliIdleTimeout        *int64                        `cty:"cli_idle_timeout" hcl:"cli_idle_timeout"`
	ConfigRevertTimer     *int64                        `cty:"config_revert_timer" hcl:"config_revert_timer"`
	DhcpOptionFqdn        *bool                         `cty:"dhcp_option_fqdn" hcl:"dhcp_option_fqdn"`
	DisableOobDownAlarm   *bool                         `cty:"disable_oob_down_alarm" hcl:"disable_oob_down_alarm"`
	FipsEnabled           *bool                         `cty:"fips_enabled" hcl:"fips_enabled"`
	LocalAccounts         map[string]LocalAccountsValue `cty:"local_accounts" hcl:"local_accounts"`
	MxedgeProxyHost       *string                       `cty:"mxedge_proxy_host" hcl:"mxedge_proxy_host"`
	MxedgeProxyPort       *string                       `cty:"mxedge_proxy_port" hcl:"mxedge_proxy_port"`
	ProtectRe             *ProtectReValue               `cty:"protect_re" hcl:"protect_re"`
	RemoveExistingConfigs *bool                         `cty:"remove_existing_configs" hcl:"remove_existing_configs"`
	RootPassword          *string                       `cty:"root_password" hcl:"root_password"`
	Tacacs                *TacacsValue                  `cty:"tacacs" hcl:"tacacs"`
	UseMxedgeProxy        *bool                         `cty:"use_mxedge_proxy" hcl:"use_mxedge_proxy"`
}

type LocalAccountsValue struct {
	Password *string `cty:"password" hcl:"password"`
	Role     *string `cty:"role" hcl:"role"`
}

type ProtectReValue struct {
	AllowedServices []string      `cty:"allowed_services" hcl:"allowed_services"`
	Custom          []CustomValue `cty:"custom" hcl:"custom"`
	Enabled         *bool         `cty:"enabled" hcl:"enabled"`
	HitCount        *bool         `cty:"hit_count" hcl:"hit_count"`
	TrustedHosts    []string      `cty:"trusted_hosts" hcl:"trusted_hosts"`
}

type CustomValue struct {
	PortRange *string  `cty:"port_range" hcl:"port_range"`
	Protocol  *string  `cty:"protocol" hcl:"protocol"`
	Subnets   []string `cty:"subnets" hcl:"subnets"`
}

type TacacsValue struct {
	DefaultRole    *string               `cty:"default_role" hcl:"default_role"`
	Enabled        *bool                 `cty:"enabled" hcl:"enabled"`
	Network        *string               `cty:"network" hcl:"network"`
	TacacctServers []TacacctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	TacplusServers []TacplusServersValue `cty:"tacplus_servers" hcl:"tacplus_servers"`
}

type TacacctServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type TacplusServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type VirtualChassisValue struct {
	Members        []MembersValue `cty:"members" hcl:"members"`
	Preprovisioned *bool          `cty:"preprovisioned" hcl:"preprovisioned"`
}

type MembersValue struct {
	Mac      string `cty:"mac" hcl:"mac"`
	MemberId int64  `cty:"member_id" hcl:"member_id"`
	VcRole   string `cty:"vc_role" hcl:"vc_role"`
}

type VrfConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type VrfInstancesValue struct {
	EvpnAutoLoopbackSubnet  *string                         `cty:"evpn_auto_loopback_subnet" hcl:"evpn_auto_loopback_subnet"`
	EvpnAutoLoopbackSubnet6 *string                         `cty:"evpn_auto_loopback_subnet6" hcl:"evpn_auto_loopback_subnet6"`
	Networks                []string                        `cty:"networks" hcl:"networks"`
	VrfExtraRoutes          map[string]VrfExtraRoutesValue  `cty:"extra_routes" hcl:"extra_routes"`
	VrfExtraRoutes6         map[string]VrfExtraRoutes6Value `cty:"extra_routes6" hcl:"extra_routes6"`
}

type VrfExtraRoutesValue struct {
	Via string `cty:"via" hcl:"via"`
}

type VrfExtraRoutes6Value struct {
	Via *string `cty:"via" hcl:"via"`
}

type VrrpConfigValue struct {
	Enabled *bool                  `cty:"enabled" hcl:"enabled"`
	Groups  map[string]GroupsValue `cty:"groups" hcl:"groups"`
}

type GroupsValue struct {
	Preempt  *bool  `cty:"preempt" hcl:"preempt"`
	Priority *int64 `cty:"priority" hcl:"priority"`
}
