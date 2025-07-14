package provider

type DeviceSwitchModel struct {
	AclPolicies           []AclPoliciesValue              `hcl:"acl_policies"`
	AclTags               map[string]AclTagsValue         `hcl:"acl_tags"`
	AdditionalConfigCmds  []string                        `hcl:"additional_config_cmds"`
	DeviceId              string                          `hcl:"device_id"`
	DhcpSnooping          *DhcpSnoopingValue              `hcl:"dhcp_snooping"`
	DhcpdConfig           *DhcpdConfigValue               `hcl:"dhcpd_config"`
	DisableAutoConfig     *bool                           `hcl:"disable_auto_config"`
	DnsServers            []string                        `hcl:"dns_servers"`
	DnsSuffix             []string                        `hcl:"dns_suffix"`
	ExtraRoutes           map[string]ExtraRoutesValue     `hcl:"extra_routes"`
	ExtraRoutes6          map[string]ExtraRoutes6Value    `hcl:"extra_routes6"`
	IpConfig              *SwitchIpConfigValue            `hcl:"ip_config"`
	LocalPortConfig       map[string]LocalPortConfigValue `hcl:"local_port_config"`
	Managed               *bool                           `hcl:"managed"`
	MapId                 *string                         `hcl:"map_id"`
	MistNac               *MistNacValue                   `hcl:"mist_nac"`
	Name                  string                          `hcl:"name"`
	Networks              map[string]NetworksValue        `hcl:"networks"`
	Notes                 *string                         `hcl:"notes"`
	NtpServers            []string                        `hcl:"ntp_servers"`
	OobIpConfig           *OobIpConfigValue               `hcl:"oob_ip_config"`
	OspfAreas             map[string]OspfAreasValue       `hcl:"ospf_areas"`
	OtherIpConfigs        map[string]OtherIpConfigsValue  `hcl:"other_ip_configs"`
	PortConfig            map[string]PortConfigValue      `hcl:"port_config"`
	PortMirroring         map[string]PortMirroringValue   `hcl:"port_mirroring"`
	PortUsages            map[string]PortUsagesValue      `hcl:"port_usages"`
	RadiusConfig          *RadiusConfigValue              `hcl:"radius_config"`
	RemoteSyslog          *RemoteSyslogValue              `hcl:"remote_syslog"`
	Role                  *string                         `hcl:"role"`
	RouterId              *string                         `hcl:"router_id"`
	SiteId                string                          `hcl:"site_id"`
	SnmpConfig            *SnmpConfigValue                `hcl:"snmp_config"`
	StpConfig             *StpConfigValue                 `hcl:"stp_config"`
	SwitchMgmt            *SwitchMgmtValue                `hcl:"switch_mgmt"`
	UseRouterIdAsSourceIp *bool                           `hcl:"use_router_id_as_source_ip"`
	Vars                  map[string]string               `hcl:"vars"`
	VirtualChassis        *VirtualChassisValue            `hcl:"virtual_chassis"`
	VrfConfig             *VrfConfigValue                 `hcl:"vrf_config"`
	VrfInstances          map[string]VrfInstancesValue    `hcl:"vrf_instances"`
	VrrpConfig            *VrrpConfigValue                `hcl:"vrrp_config"`
	X                     *float64                        `hcl:"x"`
	Y                     *float64                        `hcl:"y"`
}

type AclPoliciesValue struct {
	Actions []ActionsValue `cty:"actions"`
	Name    *string        `cty:"name"`
	SrcTags []string       `cty:"src_tags"`
}

type ActionsValue struct {
	Action *string `cty:"action"`
	DstTag string  `cty:"dst_tag"`
}

type AclTagsValue struct {
	GbpTag      *int64       `cty:"gbp_tag"`
	Macs        []string     `cty:"macs"`
	Network     *string      `cty:"network"`
	RadiusGroup *string      `cty:"radius_group"`
	Specs       []SpecsValue `cty:"specs"`
	Subnets     []string     `cty:"subnets"`
	AclTagsType string       `cty:"type"`
}

type SpecsValue struct {
	PortRange *string `cty:"port_range"`
	Protocol  *string `cty:"protocol"`
}

type DhcpSnoopingValue struct {
	AllNetworks         *bool    `cty:"all_networks"`
	EnableArpSpoofCheck *bool    `cty:"enable_arp_spoof_check"`
	EnableIpSourceGuard *bool    `cty:"enable_ip_source_guard"`
	Enabled             *bool    `cty:"enabled"`
	Networks            []string `cty:"networks"`
}

type DhcpdConfigValue struct {
	Config  map[string]ConfigValue `cty:"config"`
	Enabled *bool                  `cty:"enabled"`
}

type ConfigValue struct {
	DnsServers         []string                           `cty:"dns_servers"`
	DnsSuffix          []string                           `cty:"dns_suffix"`
	FixedBindings      map[string]FixedBindingsValue      `cty:"fixed_bindings"`
	Gateway            *string                            `cty:"gateway"`
	IpEnd              *string                            `cty:"ip_end"`
	IpEnd6             *string                            `cty:"ip_end6"`
	IpStart            *string                            `cty:"ip_start"`
	IpStart6           *string                            `cty:"ip_start6"`
	LeaseTime          *int64                             `cty:"lease_time"`
	Options            map[string]OptionsValue            `cty:"options"`
	ServerIdOverride   *bool                              `cty:"server_id_override"`
	Servers            []string                           `cty:"servers"`
	Servers6           []string                           `cty:"servers6"`
	ConfigType         *string                            `cty:"type"`
	Type6              *string                            `cty:"type6"`
	VendorEncapsulated map[string]VendorEncapsulatedValue `cty:"vendor_encapsulated"`
}

type FixedBindingsValue struct {
	Ip   string  `cty:"ip"`
	Name *string `cty:"name"`
}

type OptionsValue struct {
	OptionsType *string `cty:"type"`
	Value       *string `cty:"value"`
}

type VendorEncapsulatedValue struct {
	VendorEncapsulatedType *string `cty:"type"`
	Value                  *string `cty:"value"`
}

type ExtraRoutesValue struct {
	Discard       *bool                         `cty:"discard"`
	Metric        *int64                        `cty:"metric"`
	NextQualified map[string]NextQualifiedValue `cty:"next_qualified"`
	NoResolve     *bool                         `cty:"no_resolve"`
	Preference    *int64                        `cty:"preference"`
	Via           string                        `cty:"via"`
}

type NextQualifiedValue struct {
	Metric     *int64 `cty:"metric"`
	Preference *int64 `cty:"preference"`
}

type ExtraRoutes6Value struct {
	Discard       *bool                         `cty:"discard"`
	Metric        *int64                        `cty:"metric"`
	NextQualified map[string]NextQualifiedValue `cty:"next_qualified"`
	NoResolve     *bool                         `cty:"no_resolve"`
	Preference    *int64                        `cty:"preference"`
	Via           string                        `cty:"via"`
}

type SwitchIpConfigValue struct {
	Dns          []string `cty:"dns"`
	DnsSuffix    []string `cty:"dns_suffix"`
	Gateway      *string  `cty:"gateway"`
	Ip           *string  `cty:"ip"`
	Netmask      *string  `cty:"netmask"`
	Network      *string  `cty:"network"`
	IpConfigType *string  `cty:"type"`
}

type LocalPortConfigValue struct {
	AllNetworks                             *bool              `cty:"all_networks"`
	AllowDhcpd                              *bool              `cty:"allow_dhcpd"`
	AllowMultipleSupplicants                *bool              `cty:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                *bool              `cty:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnkownClient *bool              `cty:"bypass_auth_when_server_down_for_unkown_client"`
	Description                             *string            `cty:"description"`
	DisableAutoneg                          *bool              `cty:"disable_autoneg"`
	Disabled                                *bool              `cty:"disabled"`
	Duplex                                  *string            `cty:"duplex"`
	DynamicVlanNetworks                     []string           `cty:"dynamic_vlan_networks"`
	EnableMacAuth                           *bool              `cty:"enable_mac_auth"`
	EnableQos                               *bool              `cty:"enable_qos"`
	GuestNetwork                            *string            `cty:"guest_network"`
	InterSwitchLink                         *bool              `cty:"inter_switch_link"`
	MacAuthOnly                             *bool              `cty:"mac_auth_only"`
	MacAuthPreferred                        *bool              `cty:"mac_auth_preferred"`
	MacAuthProtocol                         *string            `cty:"mac_auth_protocol"`
	MacLimit                                *int64             `cty:"mac_limit"`
	Mode                                    *string            `cty:"mode"`
	Mtu                                     *int64             `cty:"mtu"`
	Networks                                []string           `cty:"networks"`
	Note                                    *string            `cty:"note"`
	PersistMac                              *bool              `cty:"persist_mac"`
	PoeDisabled                             *bool              `cty:"poe_disabled"`
	PortAuth                                *string            `cty:"port_auth"`
	PortNetwork                             *string            `cty:"port_network"`
	ReauthInterval                          *int64             `cty:"reauth_interval"`
	ServerFailNetwork                       *string            `cty:"server_fail_network"`
	ServerRejectNetwork                     *string            `cty:"server_reject_network"`
	Speed                                   *string            `cty:"speed"`
	StormControl                            *StormControlValue `cty:"storm_control"`
	StpEdge                                 *bool              `cty:"stp_edge"`
	StpNoRootPort                           *bool              `cty:"stp_no_root_port"`
	StpP2p                                  *bool              `cty:"stp_p2p"`
	Usage                                   string             `cty:"usage"`
	UseVstp                                 *bool              `cty:"use_vstp"`
	VoipNetwork                             *string            `cty:"voip_network"`
}

type StormControlValue struct {
	NoBroadcast           *bool  `cty:"no_broadcast"`
	NoMulticast           *bool  `cty:"no_multicast"`
	NoRegisteredMulticast *bool  `cty:"no_registered_multicast"`
	NoUnknownUnicast      *bool  `cty:"no_unknown_unicast"`
	Percentage            *int64 `cty:"percentage"`
}

type MistNacValue struct {
	Enabled *bool   `cty:"enabled"`
	Network *string `cty:"network"`
}

type NetworksValue struct {
	Gateway         *string `cty:"gateway"`
	Gateway6        *string `cty:"gateway6"`
	Isolation       *bool   `cty:"isolation"`
	IsolationVlanId *string `cty:"isolation_vlan_id"`
	Subnet          *string `cty:"subnet"`
	Subnet6         *string `cty:"subnet6"`
	VlanId          string  `cty:"vlan_id"`
}

type OobIpConfigValue struct {
	Gateway              *string `cty:"gateway"`
	Ip                   *string `cty:"ip"`
	Netmask              *string `cty:"netmask"`
	Network              *string `cty:"network"`
	OobIpConfigType      *string `cty:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out"`
}

type OspfAreasValue struct {
	IncludeLoopback *bool                        `cty:"include_loopback"`
	OspfNetworks    map[string]OspfNetworksValue `cty:"networks"`
	OspfAreasType   *string                      `cty:"type"`
}

type OspfNetworksValue struct {
	AuthKeys               map[string]string `cty:"auth_keys"`
	AuthPassword           *string           `cty:"auth_password"`
	AuthType               *string           `cty:"auth_type"`
	BfdMinimumInterval     *int64            `cty:"bfd_minimum_interval"`
	DeadInterval           *int64            `cty:"dead_interval"`
	ExportPolicy           *string           `cty:"export_policy"`
	HelloInterval          *int64            `cty:"hello_interval"`
	ImportPolicy           *string           `cty:"import_policy"`
	InterfaceType          *string           `cty:"interface_type"`
	Metric                 *int64            `cty:"metric"`
	NoReadvertiseToOverlay *bool             `cty:"no_readvertise_to_overlay"`
	Passive                *bool             `cty:"passive"`
}

type OtherIpConfigsValue struct {
	EvpnAnycast        *bool   `cty:"evpn_anycast"`
	Ip                 *string `cty:"ip"`
	Ip6                *string `cty:"ip6"`
	Netmask            *string `cty:"netmask"`
	Netmask6           *string `cty:"netmask6"`
	OtherIpConfigsType *string `cty:"type"`
	Type6              *string `cty:"type6"`
}

type PortConfigValue struct {
	AeDisableLacp    *bool   `cty:"ae_disable_lacp"`
	AeIdx            *int64  `cty:"ae_idx"`
	AeLacpSlow       *bool   `cty:"ae_lacp_slow"`
	Aggregated       *bool   `cty:"aggregated"`
	Critical         *bool   `cty:"critical"`
	Description      *string `cty:"description"`
	DisableAutoneg   *bool   `cty:"disable_autoneg"`
	Duplex           *string `cty:"duplex"`
	DynamicUsage     *string `cty:"dynamic_usage"`
	Esilag           *bool   `cty:"esilag"`
	Mtu              *int64  `cty:"mtu"`
	NoLocalOverwrite *bool   `cty:"no_local_overwrite"`
	PoeDisabled      *bool   `cty:"poe_disabled"`
	Speed            *string `cty:"speed"`
	Usage            string  `cty:"usage"`
}

type PortMirroringValue struct {
	InputNetworksIngress []string `cty:"input_networks_ingress"`
	InputPortIdsEgress   []string `cty:"input_port_ids_egress"`
	InputPortIdsIngress  []string `cty:"input_port_ids_ingress"`
	OutputNetwork        *string  `cty:"output_network"`
	OutputPortId         *string  `cty:"output_port_id"`
}

type PortUsagesValue struct {
	AllNetworks                             *bool              `cty:"all_networks"`
	AllowDhcpd                              *bool              `cty:"allow_dhcpd"`
	AllowMultipleSupplicants                *bool              `cty:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                *bool              `cty:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnkownClient *bool              `cty:"bypass_auth_when_server_down_for_unkown_client"`
	Description                             *string            `cty:"description"`
	DisableAutoneg                          *bool              `cty:"disable_autoneg"`
	Disabled                                *bool              `cty:"disabled"`
	Duplex                                  *string            `cty:"duplex"`
	DynamicVlanNetworks                     []string           `cty:"dynamic_vlan_networks"`
	EnableMacAuth                           *bool              `cty:"enable_mac_auth"`
	EnableQos                               *bool              `cty:"enable_qos"`
	GuestNetwork                            *string            `cty:"guest_network"`
	InterSwitchLink                         *bool              `cty:"inter_switch_link"`
	MacAuthOnly                             *bool              `cty:"mac_auth_only"`
	MacAuthPreferred                        *bool              `cty:"mac_auth_preferred"`
	MacAuthProtocol                         *string            `cty:"mac_auth_protocol"`
	MacLimit                                *int64             `cty:"mac_limit"`
	Mode                                    *string            `cty:"mode"`
	Mtu                                     *int64             `cty:"mtu"`
	Networks                                []string           `cty:"networks"`
	PersistMac                              *bool              `cty:"persist_mac"`
	PoeDisabled                             *bool              `cty:"poe_disabled"`
	PortAuth                                *string            `cty:"port_auth"`
	PortNetwork                             *string            `cty:"port_network"`
	ReauthInterval                          *int64             `cty:"reauth_interval"`
	ResetDefaultWhen                        *string            `cty:"reset_default_when"`
	Rules                                   []RulesValue       `cty:"rules"`
	ServerFailNetwork                       *string            `cty:"server_fail_network"`
	ServerRejectNetwork                     *string            `cty:"server_reject_network"`
	Speed                                   *string            `cty:"speed"`
	StormControl                            *StormControlValue `cty:"storm_control"`
	StpEdge                                 *bool              `cty:"stp_edge"`
	StpNoRootPort                           *bool              `cty:"stp_no_root_port"`
	StpP2p                                  *bool              `cty:"stp_p2p"`
	UseVstp                                 *bool              `cty:"use_vstp"`
	VoipNetwork                             *string            `cty:"voip_network"`
}

type RulesValue struct {
	Equals     *string  `cty:"equals"`
	EqualsAny  []string `cty:"equals_any"`
	Expression *string  `cty:"expression"`
	Src        string   `cty:"src"`
	Usage      *string  `cty:"usage"`
}

type RadiusConfigValue struct {
	AcctInterimInterval *int64             `cty:"acct_interim_interval"`
	AcctServers         []AcctServersValue `cty:"acct_servers"`
	AuthServers         []AuthServersValue `cty:"auth_servers"`
	AuthServersRetries  *int64             `cty:"auth_servers_retries"`
	AuthServersTimeout  *int64             `cty:"auth_servers_timeout"`
	Network             *string            `cty:"network"`
	SourceIp            *string            `cty:"source_ip"`
}

type AcctServersValue struct {
	Host           string  `cty:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack"`
	Port           *int64  `cty:"port"`
	Secret         string  `cty:"secret"`
}

type AuthServersValue struct {
	Host                        string  `cty:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack"`
	Port                        *int64  `cty:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator"`
	Secret                      string  `cty:"secret"`
}

type RemoteSyslogValue struct {
	Archive          *ArchiveValue  `cty:"archive"`
	Console          *ConsoleValue  `cty:"console"`
	Enabled          *bool          `cty:"enabled"`
	Files            []FilesValue   `cty:"files"`
	Network          *string        `cty:"network"`
	SendToAllServers *bool          `cty:"send_to_all_servers"`
	Servers          []ServersValue `cty:"servers"`
	TimeFormat       *string        `cty:"time_format"`
	Users            []UsersValue   `cty:"users"`
}

type ArchiveValue struct {
	Files *int64  `cty:"files"`
	Size  *string `cty:"size"`
}

type ConsoleValue struct {
	Contents []ContentsValue `cty:"contents"`
}

type ContentsValue struct {
	Facility *string `cty:"facility"`
	Severity *string `cty:"severity"`
}

type FilesValue struct {
	Archive          *ArchiveValue   `cty:"archive"`
	Contents         []ContentsValue `cty:"contents"`
	ExplicitPriority *bool           `cty:"explicit_priority"`
	File             *string         `cty:"file"`
	Match            *string         `cty:"match"`
	StructuredData   *bool           `cty:"structured_data"`
}

type ServersValue struct {
	Contents         []ContentsValue `cty:"contents"`
	ExplicitPriority *bool           `cty:"explicit_priority"`
	Facility         *string         `cty:"facility"`
	Host             *string         `cty:"host"`
	Match            *string         `cty:"match"`
	Port             *int64          `cty:"port"`
	Protocol         *string         `cty:"protocol"`
	RoutingInstance  *string         `cty:"routing_instance"`
	Severity         *string         `cty:"severity"`
	SourceAddress    *string         `cty:"source_address"`
	StructuredData   *bool           `cty:"structured_data"`
	Tag              *string         `cty:"tag"`
}

type UsersValue struct {
	Contents []ContentsValue `cty:"contents"`
	Match    *string         `cty:"match"`
	User     *string         `cty:"user"`
}

type SnmpConfigValue struct {
	ClientList  []ClientListValue `cty:"client_list"`
	Contact     *string           `cty:"contact"`
	Description *string           `cty:"description"`
	Enabled     *bool             `cty:"enabled"`
	EngineId    *string           `cty:"engine_id"`
	Location    *string           `cty:"location"`
	Name        *string           `cty:"name"`
	Network     *string           `cty:"network"`
	TrapGroups  []TrapGroupsValue `cty:"trap_groups"`
	V2cConfig   []V2cConfigValue  `cty:"v2c_config"`
	V3Config    *V3ConfigValue    `cty:"v3_config"`
	Views       []ViewsValue      `cty:"views"`
}

type ClientListValue struct {
	ClientListName *string  `cty:"client_list_name"`
	Clients        []string `cty:"clients"`
}

type TrapGroupsValue struct {
	Categories []string `cty:"categories"`
	GroupName  *string  `cty:"group_name"`
	Targets    []string `cty:"targets"`
	Version    *string  `cty:"version"`
}

type V2cConfigValue struct {
	Authorization  *string `cty:"authorization"`
	ClientListName *string `cty:"client_list_name"`
	CommunityName  *string `cty:"community_name"`
	View           *string `cty:"view"`
}

type V3ConfigValue struct {
	Notify           []NotifyValue           `cty:"notify"`
	NotifyFilter     []NotifyFilterValue     `cty:"notify_filter"`
	TargetAddress    []TargetAddressValue    `cty:"target_address"`
	TargetParameters []TargetParametersValue `cty:"target_parameters"`
	Usm              []UsmValue              `cty:"usm"`
	Vacm             *VacmValue              `cty:"vacm"`
}

type NotifyValue struct {
	Name       string `cty:"name"`
	Tag        string `cty:"tag"`
	NotifyType string `cty:"type"`
}

type NotifyFilterValue struct {
	ProfileName    *string               `cty:"profile_name"`
	Snmpv3Contents []Snmpv3ContentsValue `cty:"contents"`
}

type Snmpv3ContentsValue struct {
	Include *bool  `cty:"include"`
	Oid     string `cty:"oid"`
}

type TargetAddressValue struct {
	Address           string  `cty:"address"`
	AddressMask       string  `cty:"address_mask"`
	Port              *string `cty:"port"`
	TagList           *string `cty:"tag_list"`
	TargetAddressName string  `cty:"target_address_name"`
	TargetParameters  *string `cty:"target_parameters"`
}

type TargetParametersValue struct {
	MessageProcessingModel string  `cty:"message_processing_model"`
	Name                   string  `cty:"name"`
	NotifyFilter           *string `cty:"notify_filter"`
	SecurityLevel          *string `cty:"security_level"`
	SecurityModel          *string `cty:"security_model"`
	SecurityName           *string `cty:"security_name"`
}

type UsmValue struct {
	EngineType     string             `cty:"engine_type"`
	RemoteEngineId *string            `cty:"remote_engine_id"`
	Snmpv3Users    []Snmpv3UsersValue `cty:"users"`
}

type Snmpv3UsersValue struct {
	AuthenticationPassword *string `cty:"authentication_password"`
	AuthenticationType     *string `cty:"authentication_type"`
	EncryptionPassword     *string `cty:"encryption_password"`
	EncryptionType         *string `cty:"encryption_type"`
	Name                   *string `cty:"name"`
}

type VacmValue struct {
	Access          []AccessValue         `cty:"access"`
	SecurityToGroup *SecurityToGroupValue `cty:"security_to_group"`
}

type AccessValue struct {
	GroupName  *string           `cty:"group_name"`
	PrefixList []PrefixListValue `cty:"prefix_list"`
}

type PrefixListValue struct {
	ContextPrefix  *string `cty:"context_prefix"`
	NotifyView     *string `cty:"notify_view"`
	ReadView       *string `cty:"read_view"`
	SecurityLevel  *string `cty:"security_level"`
	SecurityModel  *string `cty:"security_model"`
	PrefixListType *string `cty:"type"`
	WriteView      *string `cty:"write_view"`
}

type SecurityToGroupValue struct {
	SecurityModel     *string                  `cty:"security_model"`
	Snmpv3VacmContent []Snmpv3VacmContentValue `cty:"content"`
}

type Snmpv3VacmContentValue struct {
	Group        *string `cty:"group"`
	SecurityName *string `cty:"security_name"`
}

type ViewsValue struct {
	Include  *bool   `cty:"include"`
	Oid      *string `cty:"oid"`
	ViewName *string `cty:"view_name"`
}

type StpConfigValue struct {
	BridgePriority *string `cty:"bridge_priority"`
}

type SwitchMgmtValue struct {
	ApAffinityThreshold *int64                        `cty:"ap_affinity_threshold"`
	CliBanner           *string                       `cty:"cli_banner"`
	CliIdleTimeout      *int64                        `cty:"cli_idle_timeout"`
	ConfigRevertTimer   *int64                        `cty:"config_revert_timer"`
	DhcpOptionFqdn      *bool                         `cty:"dhcp_option_fqdn"`
	DisableOobDownAlarm *bool                         `cty:"disable_oob_down_alarm"`
	LocalAccounts       map[string]LocalAccountsValue `cty:"local_accounts"`
	MxedgeProxyHost     *string                       `cty:"mxedge_proxy_host"`
	MxedgeProxyPort     *int64                        `cty:"mxedge_proxy_port"`
	ProtectRe           *ProtectReValue               `cty:"protect_re"`
	RootPassword        *string                       `cty:"root_password"`
	Tacacs              *TacacsValue                  `cty:"tacacs"`
	UseMxedgeProxy      *bool                         `cty:"use_mxedge_proxy"`
}

type LocalAccountsValue struct {
	Password *string `cty:"password"`
	Role     *string `cty:"role"`
}

type ProtectReValue struct {
	AllowedServices []string      `cty:"allowed_services"`
	Custom          []CustomValue `cty:"custom"`
	Enabled         *bool         `cty:"enabled"`
	TrustedHosts    []string      `cty:"trusted_hosts"`
}

type CustomValue struct {
	PortRange *string  `cty:"port_range"`
	Protocol  *string  `cty:"protocol"`
	Subnets   []string `cty:"subnets"`
}

type TacacsValue struct {
	DefaultRole    *string               `cty:"default_role"`
	Enabled        *bool                 `cty:"enabled"`
	Network        *string               `cty:"network"`
	TacacctServers []TacacctServersValue `cty:"acct_servers"`
	TacplusServers []TacplusServersValue `cty:"tacplus_servers"`
}

type TacacctServersValue struct {
	Host    *string `cty:"host"`
	Port    *string `cty:"port"`
	Secret  *string `cty:"secret"`
	Timeout *int64  `cty:"timeout"`
}

type TacplusServersValue struct {
	Host    *string `cty:"host"`
	Port    *string `cty:"port"`
	Secret  *string `cty:"secret"`
	Timeout *int64  `cty:"timeout"`
}

type VirtualChassisValue struct {
	Members        []MembersValue `cty:"members"`
	Preprovisioned *bool          `cty:"preprovisioned"`
}

type MembersValue struct {
	Mac      string `cty:"mac"`
	MemberId int64  `cty:"member_id"`
	VcRole   string `cty:"vc_role"`
}

type VrfConfigValue struct {
	Enabled *bool `cty:"enabled"`
}

type VrfInstancesValue struct {
	Networks       []string                       `cty:"networks"`
	VrfExtraRoutes map[string]VrfExtraRoutesValue `cty:"vrf_extra_routes"`
}

type VrfExtraRoutesValue struct {
	Via string `cty:"via"`
}

type VrrpConfigValue struct {
	Enabled *bool                  `cty:"enabled"`
	Groups  map[string]GroupsValue `cty:"groups"`
}

type GroupsValue struct {
	Priority *int64 `cty:"priority"`
}
