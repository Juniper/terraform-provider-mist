package provider

type SiteNetworktemplateModel struct {
	AclPolicies                     []SiteNetworktemplateAclPoliciesValue              `hcl:"acl_policies"`
	AclTags                         map[string]SiteNetworktemplateAclTagsValue         `hcl:"acl_tags"`
	AdditionalConfigCmds            []string                                           `hcl:"additional_config_cmds"`
	AutoUpgradeLinecard             *bool                                              `hcl:"auto_upgrade_linecard"`
	DefaultPortUsage                *string                                            `hcl:"default_port_usage"`
	DhcpSnooping                    *SiteNetworktemplateDhcpSnoopingValue              `hcl:"dhcp_snooping"`
	DisabledSystemDefinedPortUsages []string                                           `hcl:"disabled_system_defined_port_usages"`
	DnsServers                      []string                                           `hcl:"dns_servers"`
	DnsSuffix                       []string                                           `hcl:"dns_suffix"`
	ExtraRoutes                     map[string]SiteNetworktemplateExtraRoutesValue     `hcl:"extra_routes"`
	ExtraRoutes6                    map[string]SiteNetworktemplateExtraRoutes6Value    `hcl:"extra_routes6"`
	MistNac                         *SiteNetworktemplateMistNacValue                   `hcl:"mist_nac"`
	Networks                        map[string]SiteNetworktemplateNetworksValue        `hcl:"networks"`
	NtpServers                      []string                                           `hcl:"ntp_servers"`
	OspfAreas                       map[string]SiteNetworktemplateOspfAreasValue       `hcl:"ospf_areas"`
	PortMirroring                   map[string]SiteNetworktemplatePortMirroringValue   `hcl:"port_mirroring"`
	PortUsages                      map[string]SiteNetworktemplatePortUsagesValue      `hcl:"port_usages"`
	RadiusConfig                    *SiteNetworktemplateRadiusConfigValue              `hcl:"radius_config"`
	RemoteSyslog                    *SiteNetworktemplateRemoteSyslogValue              `hcl:"remote_syslog"`
	RoutingPolicies                 map[string]SiteNetworktemplateRoutingPoliciesValue `hcl:"routing_policies"`
	SiteId                          string                                             `hcl:"site_id"`
	SnmpConfig                      *SiteNetworktemplateSnmpConfigValue                `hcl:"snmp_config"`
	SwitchMatching                  *SiteNetworktemplateSwitchMatchingValue            `hcl:"switch_matching"`
	SwitchMgmt                      *SiteNetworktemplateSwitchMgmtValue                `hcl:"switch_mgmt"`
	VrfConfig                       *SiteNetworktemplateVrfConfigValue                 `hcl:"vrf_config"`
	VrfInstances                    map[string]SiteNetworktemplateVrfInstancesValue    `hcl:"vrf_instances"`
}

type SiteNetworktemplateAclPoliciesValue struct {
	Actions []SiteNetworktemplateActionsValue `cty:"actions" hcl:"actions"`
	Name    *string                           `cty:"name" hcl:"name"`
	SrcTags []string                          `cty:"src_tags" hcl:"src_tags"`
}

type SiteNetworktemplateActionsValue struct {
	Action *string `cty:"action" hcl:"action"`
	DstTag string  `cty:"dst_tag" hcl:"dst_tag"`
}

type SiteNetworktemplateAclTagsValue struct {
	EtherTypes  []string                        `cty:"ether_types" hcl:"ether_types"`
	GbpTag      *int64                          `cty:"gbp_tag" hcl:"gbp_tag"`
	Macs        []string                        `cty:"macs" hcl:"macs"`
	Network     *string                         `cty:"network" hcl:"network"`
	PortUsage   *string                         `cty:"port_usage" hcl:"port_usage"`
	RadiusGroup *string                         `cty:"radius_group" hcl:"radius_group"`
	Specs       []SiteNetworktemplateSpecsValue `cty:"specs" hcl:"specs"`
	Subnets     []string                        `cty:"subnets" hcl:"subnets"`
	AclTagsType string                          `cty:"type" hcl:"type"`
}

type SiteNetworktemplateSpecsValue struct {
	PortRange *string `cty:"port_range" hcl:"port_range"`
	Protocol  *string `cty:"protocol" hcl:"protocol"`
}

type SiteNetworktemplateDhcpSnoopingValue struct {
	AllNetworks         *bool    `cty:"all_networks" hcl:"all_networks"`
	EnableArpSpoofCheck *bool    `cty:"enable_arp_spoof_check" hcl:"enable_arp_spoof_check"`
	EnableIpSourceGuard *bool    `cty:"enable_ip_source_guard" hcl:"enable_ip_source_guard"`
	Enabled             *bool    `cty:"enabled" hcl:"enabled"`
	Networks            []string `cty:"networks" hcl:"networks"`
}

type SiteNetworktemplateExtraRoutesValue struct {
	Discard       *bool                                            `cty:"discard" hcl:"discard"`
	Metric        *int64                                           `cty:"metric" hcl:"metric"`
	NextQualified map[string]SiteNetworktemplateNextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                                            `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                                           `cty:"preference" hcl:"preference"`
	Via           string                                           `cty:"via" hcl:"via"`
}

type SiteNetworktemplateNextQualifiedValue struct {
	Metric     *int64 `cty:"metric" hcl:"metric"`
	Preference *int64 `cty:"preference" hcl:"preference"`
}

type SiteNetworktemplateExtraRoutes6Value struct {
	Discard       *bool                                            `cty:"discard" hcl:"discard"`
	Metric        *int64                                           `cty:"metric" hcl:"metric"`
	NextQualified map[string]SiteNetworktemplateNextQualifiedValue `cty:"next_qualified" hcl:"next_qualified"`
	NoResolve     *bool                                            `cty:"no_resolve" hcl:"no_resolve"`
	Preference    *int64                                           `cty:"preference" hcl:"preference"`
	Via           string                                           `cty:"via" hcl:"via"`
}

type SiteNetworktemplateMistNacValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Network *string `cty:"network" hcl:"network"`
}

type SiteNetworktemplateNetworksValue struct {
	Gateway         *string `cty:"gateway" hcl:"gateway"`
	Gateway6        *string `cty:"gateway6" hcl:"gateway6"`
	Isolation       *bool   `cty:"isolation" hcl:"isolation"`
	IsolationVlanId *string `cty:"isolation_vlan_id" hcl:"isolation_vlan_id"`
	Subnet          *string `cty:"subnet" hcl:"subnet"`
	Subnet6         *string `cty:"subnet6" hcl:"subnet6"`
	VlanId          string  `cty:"vlan_id" hcl:"vlan_id"`
}

type SiteNetworktemplateOspfAreasValue struct {
	IncludeLoopback *bool                                           `cty:"include_loopback" hcl:"include_loopback"`
	OspfNetworks    map[string]SiteNetworktemplateOspfNetworksValue `cty:"networks" hcl:"networks"`
	OspfAreasType   *string                                         `cty:"type" hcl:"type"`
}

type SiteNetworktemplateOspfNetworksValue struct {
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

type SiteNetworktemplatePortMirroringValue struct {
	InputNetworksIngress []string `cty:"input_networks_ingress" hcl:"input_networks_ingress"`
	InputPortIdsEgress   []string `cty:"input_port_ids_egress" hcl:"input_port_ids_egress"`
	InputPortIdsIngress  []string `cty:"input_port_ids_ingress" hcl:"input_port_ids_ingress"`
	OutputIpAddress      *string  `cty:"output_ip_address" hcl:"output_ip_address"`
	OutputNetwork        *string  `cty:"output_network" hcl:"output_network"`
	OutputPortId         *string  `cty:"output_port_id" hcl:"output_port_id"`
}

type SiteNetworktemplatePortUsagesValue struct {
	AllNetworks                              *bool                                 `cty:"all_networks" hcl:"all_networks"`
	AllowDhcpd                               *bool                                 `cty:"allow_dhcpd" hcl:"allow_dhcpd"`
	AllowMultipleSupplicants                 *bool                                 `cty:"allow_multiple_supplicants" hcl:"allow_multiple_supplicants"`
	BypassAuthWhenServerDown                 *bool                                 `cty:"bypass_auth_when_server_down" hcl:"bypass_auth_when_server_down"`
	BypassAuthWhenServerDownForUnknownClient *bool                                 `cty:"bypass_auth_when_server_down_for_unknown_client" hcl:"bypass_auth_when_server_down_for_unknown_client"`
	BypassAuthWhenServerDownForVoip          *bool                                 `cty:"bypass_auth_when_server_down_for_voip" hcl:"bypass_auth_when_server_down_for_voip"`
	CommunityVlanId                          *int64                                `cty:"community_vlan_id" hcl:"community_vlan_id"`
	Description                              *string                               `cty:"description" hcl:"description"`
	DisableAutoneg                           *bool                                 `cty:"disable_autoneg" hcl:"disable_autoneg"`
	Disabled                                 *bool                                 `cty:"disabled" hcl:"disabled"`
	Duplex                                   *string                               `cty:"duplex" hcl:"duplex"`
	DynamicVlanNetworks                      []string                              `cty:"dynamic_vlan_networks" hcl:"dynamic_vlan_networks"`
	EnableMacAuth                            *bool                                 `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	EnableQos                                *bool                                 `cty:"enable_qos" hcl:"enable_qos"`
	GuestNetwork                             *string                               `cty:"guest_network" hcl:"guest_network"`
	InterIsolationNetworkLink                *bool                                 `cty:"inter_isolation_network_link" hcl:"inter_isolation_network_link"`
	InterSwitchLink                          *bool                                 `cty:"inter_switch_link" hcl:"inter_switch_link"`
	MacAuthOnly                              *bool                                 `cty:"mac_auth_only" hcl:"mac_auth_only"`
	MacAuthPreferred                         *bool                                 `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol                          *string                               `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MacLimit                                 *string                               `cty:"mac_limit" hcl:"mac_limit"`
	Mode                                     *string                               `cty:"mode" hcl:"mode"`
	Mtu                                      *string                               `cty:"mtu" hcl:"mtu"`
	Networks                                 []string                              `cty:"networks" hcl:"networks"`
	PersistMac                               *bool                                 `cty:"persist_mac" hcl:"persist_mac"`
	PoeDisabled                              *bool                                 `cty:"poe_disabled" hcl:"poe_disabled"`
	PoePriority                              *string                               `cty:"poe_priority" hcl:"poe_priority"`
	PortAuth                                 *string                               `cty:"port_auth" hcl:"port_auth"`
	PortNetwork                              *string                               `cty:"port_network" hcl:"port_network"`
	ReauthInterval                           *string                               `cty:"reauth_interval" hcl:"reauth_interval"`
	ResetDefaultWhen                         *string                               `cty:"reset_default_when" hcl:"reset_default_when"`
	Rules                                    []SiteNetworktemplateRulesValue       `cty:"rules" hcl:"rules"`
	ServerFailNetwork                        *string                               `cty:"server_fail_network" hcl:"server_fail_network"`
	ServerRejectNetwork                      *string                               `cty:"server_reject_network" hcl:"server_reject_network"`
	Speed                                    *string                               `cty:"speed" hcl:"speed"`
	StormControl                             *SiteNetworktemplateStormControlValue `cty:"storm_control" hcl:"storm_control"`
	StpDisable                               *bool                                 `cty:"stp_disable" hcl:"stp_disable"`
	StpEdge                                  *bool                                 `cty:"stp_edge" hcl:"stp_edge"`
	StpNoRootPort                            *bool                                 `cty:"stp_no_root_port" hcl:"stp_no_root_port"`
	StpP2p                                   *bool                                 `cty:"stp_p2p" hcl:"stp_p2p"`
	StpRequired                              *bool                                 `cty:"stp_required" hcl:"stp_required"`
	UiEvpntopoId                             *string                               `cty:"ui_evpntopo_id" hcl:"ui_evpntopo_id"`
	UseVstp                                  *bool                                 `cty:"use_vstp" hcl:"use_vstp"`
	VoipNetwork                              *string                               `cty:"voip_network" hcl:"voip_network"`
}

type SiteNetworktemplateRulesValue struct {
	Equals     *string  `cty:"equals" hcl:"equals"`
	EqualsAny  []string `cty:"equals_any" hcl:"equals_any"`
	Expression *string  `cty:"expression" hcl:"expression"`
	Src        string   `cty:"src" hcl:"src"`
	Usage      *string  `cty:"usage" hcl:"usage"`
}

type SiteNetworktemplateStormControlValue struct {
	DisablePort           *bool  `cty:"disable_port" hcl:"disable_port"`
	NoBroadcast           *bool  `cty:"no_broadcast" hcl:"no_broadcast"`
	NoMulticast           *bool  `cty:"no_multicast" hcl:"no_multicast"`
	NoRegisteredMulticast *bool  `cty:"no_registered_multicast" hcl:"no_registered_multicast"`
	NoUnknownUnicast      *bool  `cty:"no_unknown_unicast" hcl:"no_unknown_unicast"`
	Percentage            *int64 `cty:"percentage" hcl:"percentage"`
}

type SiteNetworktemplateRadiusConfigValue struct {
	AcctImmediateUpdate *bool                                 `cty:"acct_immediate_update" hcl:"acct_immediate_update"`
	AcctInterimInterval *int64                                `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AcctServers         []SiteNetworktemplateAcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServerSelection *string                               `cty:"auth_server_selection" hcl:"auth_server_selection"`
	AuthServers         []SiteNetworktemplateAuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	AuthServersRetries  *int64                                `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64                                `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool                                 `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *string                               `cty:"coa_port" hcl:"coa_port"`
	FastDot1xTimers     *bool                                 `cty:"fast_dot1x_timers" hcl:"fast_dot1x_timers"`
	Network             *string                               `cty:"network" hcl:"network"`
	SourceIp            *string                               `cty:"source_ip" hcl:"source_ip"`
}

type SiteNetworktemplateAcctServersValue struct {
	Host           string  `cty:"host" hcl:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port           *string `cty:"port" hcl:"port"`
	Secret         string  `cty:"secret" hcl:"secret"`
}

type SiteNetworktemplateAuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type SiteNetworktemplateRemoteSyslogValue struct {
	Archive          *SiteNetworktemplateArchiveValue  `cty:"archive" hcl:"archive"`
	Cacerts          []string                          `cty:"cacerts" hcl:"cacerts"`
	Console          *SiteNetworktemplateConsoleValue  `cty:"console" hcl:"console"`
	Enabled          *bool                             `cty:"enabled" hcl:"enabled"`
	Files            []SiteNetworktemplateFilesValue   `cty:"files" hcl:"files"`
	Network          *string                           `cty:"network" hcl:"network"`
	SendToAllServers *bool                             `cty:"send_to_all_servers" hcl:"send_to_all_servers"`
	Servers          []SiteNetworktemplateServersValue `cty:"servers" hcl:"servers"`
	TimeFormat       *string                           `cty:"time_format" hcl:"time_format"`
	Users            []SiteNetworktemplateUsersValue   `cty:"users" hcl:"users"`
}

type SiteNetworktemplateArchiveValue struct {
	Files *string `cty:"files" hcl:"files"`
	Size  *string `cty:"size" hcl:"size"`
}

type SiteNetworktemplateConsoleValue struct {
	Contents []SiteNetworktemplateContentsValue `cty:"contents" hcl:"contents"`
}

type SiteNetworktemplateContentsValue struct {
	Facility *string `cty:"facility" hcl:"facility"`
	Severity *string `cty:"severity" hcl:"severity"`
}

type SiteNetworktemplateFilesValue struct {
	Archive          *SiteNetworktemplateArchiveValue   `cty:"archive" hcl:"archive"`
	Contents         []SiteNetworktemplateContentsValue `cty:"contents" hcl:"contents"`
	EnableTls        *bool                              `cty:"enable_tls" hcl:"enable_tls"`
	ExplicitPriority *bool                              `cty:"explicit_priority" hcl:"explicit_priority"`
	File             *string                            `cty:"file" hcl:"file"`
	Match            *string                            `cty:"match" hcl:"match"`
	StructuredData   *bool                              `cty:"structured_data" hcl:"structured_data"`
}

type SiteNetworktemplateServersValue struct {
	Contents         []SiteNetworktemplateContentsValue `cty:"contents" hcl:"contents"`
	ExplicitPriority *bool                              `cty:"explicit_priority" hcl:"explicit_priority"`
	Facility         *string                            `cty:"facility" hcl:"facility"`
	Host             *string                            `cty:"host" hcl:"host"`
	Match            *string                            `cty:"match" hcl:"match"`
	Port             *string                            `cty:"port" hcl:"port"`
	Protocol         *string                            `cty:"protocol" hcl:"protocol"`
	RoutingInstance  *string                            `cty:"routing_instance" hcl:"routing_instance"`
	ServerName       *string                            `cty:"server_name" hcl:"server_name"`
	Severity         *string                            `cty:"severity" hcl:"severity"`
	SourceAddress    *string                            `cty:"source_address" hcl:"source_address"`
	StructuredData   *bool                              `cty:"structured_data" hcl:"structured_data"`
	Tag              *string                            `cty:"tag" hcl:"tag"`
}

type SiteNetworktemplateUsersValue struct {
	Contents []SiteNetworktemplateContentsValue `cty:"contents" hcl:"contents"`
	Match    *string                            `cty:"match" hcl:"match"`
	User     *string                            `cty:"user" hcl:"user"`
}

type SiteNetworktemplateRoutingPoliciesValue struct {
	Terms []SiteNetworktemplateTermsValue `cty:"terms" hcl:"terms"`
}

type SiteNetworktemplateTermsValue struct {
	Matching                 *SiteNetworktemplateMatchingValue                 `cty:"matching" hcl:"matching"`
	Name                     string                                            `cty:"name" hcl:"name"`
	RoutingPolicyTermActions *SiteNetworktemplateRoutingPolicyTermActionsValue `cty:"actions" hcl:"actions"`
}

type SiteNetworktemplateMatchingValue struct {
	AsPath    []string `cty:"as_path" hcl:"as_path"`
	Community []string `cty:"community" hcl:"community"`
	Prefix    []string `cty:"prefix" hcl:"prefix"`
	Protocol  []string `cty:"protocol" hcl:"protocol"`
}

type SiteNetworktemplateRoutingPolicyTermActionsValue struct {
	Accept          *bool    `cty:"accept" hcl:"accept"`
	Community       []string `cty:"community" hcl:"community"`
	LocalPreference *string  `cty:"local_preference" hcl:"local_preference"`
	PrependAsPath   []string `cty:"prepend_as_path" hcl:"prepend_as_path"`
}

type SiteNetworktemplateSnmpConfigValue struct {
	ClientList   []SiteNetworktemplateClientListValue `cty:"client_list" hcl:"client_list"`
	Contact      *string                              `cty:"contact" hcl:"contact"`
	Description  *string                              `cty:"description" hcl:"description"`
	Enabled      *bool                                `cty:"enabled" hcl:"enabled"`
	EngineId     *string                              `cty:"engine_id" hcl:"engine_id"`
	EngineIdType *string                              `cty:"engine_id_type" hcl:"engine_id_type"`
	Location     *string                              `cty:"location" hcl:"location"`
	Name         *string                              `cty:"name" hcl:"name"`
	Network      *string                              `cty:"network" hcl:"network"`
	TrapGroups   []SiteNetworktemplateTrapGroupsValue `cty:"trap_groups" hcl:"trap_groups"`
	V2cConfig    []SiteNetworktemplateV2cConfigValue  `cty:"v2c_config" hcl:"v2c_config"`
	V3Config     *SiteNetworktemplateV3ConfigValue    `cty:"v3_config" hcl:"v3_config"`
	Views        []SiteNetworktemplateViewsValue      `cty:"views" hcl:"views"`
}

type SiteNetworktemplateClientListValue struct {
	ClientListName *string  `cty:"client_list_name" hcl:"client_list_name"`
	Clients        []string `cty:"clients" hcl:"clients"`
}

type SiteNetworktemplateTrapGroupsValue struct {
	Categories []string `cty:"categories" hcl:"categories"`
	GroupName  *string  `cty:"group_name" hcl:"group_name"`
	Targets    []string `cty:"targets" hcl:"targets"`
	Version    *string  `cty:"version" hcl:"version"`
}

type SiteNetworktemplateV2cConfigValue struct {
	Authorization  *string `cty:"authorization" hcl:"authorization"`
	ClientListName *string `cty:"client_list_name" hcl:"client_list_name"`
	CommunityName  *string `cty:"community_name" hcl:"community_name"`
	View           *string `cty:"view" hcl:"view"`
}

type SiteNetworktemplateV3ConfigValue struct {
	Notify           []SiteNetworktemplateNotifyValue           `cty:"notify" hcl:"notify"`
	NotifyFilter     []SiteNetworktemplateNotifyFilterValue     `cty:"notify_filter" hcl:"notify_filter"`
	TargetAddress    []SiteNetworktemplateTargetAddressValue    `cty:"target_address" hcl:"target_address"`
	TargetParameters []SiteNetworktemplateTargetParametersValue `cty:"target_parameters" hcl:"target_parameters"`
	Usm              []SiteNetworktemplateUsmValue              `cty:"usm" hcl:"usm"`
	Vacm             *SiteNetworktemplateVacmValue              `cty:"vacm" hcl:"vacm"`
}

type SiteNetworktemplateNotifyValue struct {
	Name       string `cty:"name" hcl:"name"`
	Tag        string `cty:"tag" hcl:"tag"`
	NotifyType string `cty:"type" hcl:"type"`
}

type SiteNetworktemplateNotifyFilterValue struct {
	ProfileName    *string                                  `cty:"profile_name" hcl:"profile_name"`
	Snmpv3Contents []SiteNetworktemplateSnmpv3ContentsValue `cty:"contents" hcl:"contents"`
}

type SiteNetworktemplateSnmpv3ContentsValue struct {
	Include *bool  `cty:"include" hcl:"include"`
	Oid     string `cty:"oid" hcl:"oid"`
}

type SiteNetworktemplateTargetAddressValue struct {
	Address           string  `cty:"address" hcl:"address"`
	AddressMask       string  `cty:"address_mask" hcl:"address_mask"`
	Port              *string `cty:"port" hcl:"port"`
	TagList           *string `cty:"tag_list" hcl:"tag_list"`
	TargetAddressName string  `cty:"target_address_name" hcl:"target_address_name"`
	TargetParameters  *string `cty:"target_parameters" hcl:"target_parameters"`
}

type SiteNetworktemplateTargetParametersValue struct {
	MessageProcessingModel string  `cty:"message_processing_model" hcl:"message_processing_model"`
	Name                   string  `cty:"name" hcl:"name"`
	NotifyFilter           *string `cty:"notify_filter" hcl:"notify_filter"`
	SecurityLevel          *string `cty:"security_level" hcl:"security_level"`
	SecurityModel          *string `cty:"security_model" hcl:"security_model"`
	SecurityName           *string `cty:"security_name" hcl:"security_name"`
}

type SiteNetworktemplateUsmValue struct {
	EngineType     string                                `cty:"engine_type" hcl:"engine_type"`
	RemoteEngineId *string                               `cty:"remote_engine_id" hcl:"remote_engine_id"`
	Snmpv3Users    []SiteNetworktemplateSnmpv3UsersValue `cty:"users" hcl:"users"`
}

type SiteNetworktemplateSnmpv3UsersValue struct {
	AuthenticationPassword *string `cty:"authentication_password" hcl:"authentication_password"`
	AuthenticationType     *string `cty:"authentication_type" hcl:"authentication_type"`
	EncryptionPassword     *string `cty:"encryption_password" hcl:"encryption_password"`
	EncryptionType         *string `cty:"encryption_type" hcl:"encryption_type"`
	Name                   *string `cty:"name" hcl:"name"`
}

type SiteNetworktemplateVacmValue struct {
	Access          []SiteNetworktemplateAccessValue         `cty:"access" hcl:"access"`
	SecurityToGroup *SiteNetworktemplateSecurityToGroupValue `cty:"security_to_group" hcl:"security_to_group"`
}

type SiteNetworktemplateAccessValue struct {
	GroupName  *string                              `cty:"group_name" hcl:"group_name"`
	PrefixList []SiteNetworktemplatePrefixListValue `cty:"prefix_list" hcl:"prefix_list"`
}

type SiteNetworktemplatePrefixListValue struct {
	ContextPrefix  *string `cty:"context_prefix" hcl:"context_prefix"`
	NotifyView     *string `cty:"notify_view" hcl:"notify_view"`
	ReadView       *string `cty:"read_view" hcl:"read_view"`
	SecurityLevel  *string `cty:"security_level" hcl:"security_level"`
	SecurityModel  *string `cty:"security_model" hcl:"security_model"`
	PrefixListType *string `cty:"type" hcl:"type"`
	WriteView      *string `cty:"write_view" hcl:"write_view"`
}

type SiteNetworktemplateSecurityToGroupValue struct {
	SecurityModel     *string                                     `cty:"security_model" hcl:"security_model"`
	Snmpv3VacmContent []SiteNetworktemplateSnmpv3VacmContentValue `cty:"content" hcl:"content"`
}

type SiteNetworktemplateSnmpv3VacmContentValue struct {
	Group        *string `cty:"group" hcl:"group"`
	SecurityName *string `cty:"security_name" hcl:"security_name"`
}

type SiteNetworktemplateViewsValue struct {
	Include  *bool   `cty:"include" hcl:"include"`
	Oid      *string `cty:"oid" hcl:"oid"`
	ViewName *string `cty:"view_name" hcl:"view_name"`
}

type SiteNetworktemplateSwitchMatchingValue struct {
	Enable        *bool                                   `cty:"enable" hcl:"enable"`
	MatchingRules []SiteNetworktemplateMatchingRulesValue `cty:"rules" hcl:"rules"`
}

type SiteNetworktemplateMatchingRulesValue struct {
	AdditionalConfigCmds []string                                         `cty:"additional_config_cmds" hcl:"additional_config_cmds"`
	IpConfig             *SiteNetworktemplateIpConfigValue                `cty:"ip_config" hcl:"ip_config"`
	MatchModel           *string                                          `cty:"match_model" hcl:"match_model"`
	MatchName            *string                                          `cty:"match_name" hcl:"match_name"`
	MatchNameOffset      *int64                                           `cty:"match_name_offset" hcl:"match_name_offset"`
	MatchRole            *string                                          `cty:"match_role" hcl:"match_role"`
	Name                 *string                                          `cty:"name" hcl:"name"`
	OobIpConfig          *SiteNetworktemplateOobIpConfigValue             `cty:"oob_ip_config" hcl:"oob_ip_config"`
	PortConfig           map[string]SiteNetworktemplatePortConfigValue    `cty:"port_config" hcl:"port_config"`
	PortMirroring        map[string]SiteNetworktemplatePortMirroringValue `cty:"port_mirroring" hcl:"port_mirroring"`
	StpConfig            *SiteNetworktemplateStpConfigValue               `cty:"stp_config" hcl:"stp_config"`
}

type SiteNetworktemplateIpConfigValue struct {
	Network      *string `cty:"network" hcl:"network"`
	IpConfigType *string `cty:"type" hcl:"type"`
}

type SiteNetworktemplateOobIpConfigValue struct {
	OobIpConfigType      *string `cty:"type" hcl:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf" hcl:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out" hcl:"use_mgmt_vrf_for_host_out"`
}

type SiteNetworktemplatePortConfigValue struct {
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

type SiteNetworktemplateStpConfigValue struct {
	BridgePriority *string `cty:"bridge_priority" hcl:"bridge_priority"`
}

type SiteNetworktemplateSwitchMgmtValue struct {
	ApAffinityThreshold   *int64                                           `cty:"ap_affinity_threshold" hcl:"ap_affinity_threshold"`
	CliBanner             *string                                          `cty:"cli_banner" hcl:"cli_banner"`
	CliIdleTimeout        *int64                                           `cty:"cli_idle_timeout" hcl:"cli_idle_timeout"`
	ConfigRevertTimer     *int64                                           `cty:"config_revert_timer" hcl:"config_revert_timer"`
	DhcpOptionFqdn        *bool                                            `cty:"dhcp_option_fqdn" hcl:"dhcp_option_fqdn"`
	DisableOobDownAlarm   *bool                                            `cty:"disable_oob_down_alarm" hcl:"disable_oob_down_alarm"`
	FipsEnabled           *bool                                            `cty:"fips_enabled" hcl:"fips_enabled"`
	LocalAccounts         map[string]SiteNetworktemplateLocalAccountsValue `cty:"local_accounts" hcl:"local_accounts"`
	MxedgeProxyHost       *string                                          `cty:"mxedge_proxy_host" hcl:"mxedge_proxy_host"`
	MxedgeProxyPort       *string                                          `cty:"mxedge_proxy_port" hcl:"mxedge_proxy_port"`
	ProtectRe             *SiteNetworktemplateProtectReValue               `cty:"protect_re" hcl:"protect_re"`
	RemoveExistingConfigs *bool                                            `cty:"remove_existing_configs" hcl:"remove_existing_configs"`
	RootPassword          *string                                          `cty:"root_password" hcl:"root_password"`
	Tacacs                *SiteNetworktemplateTacacsValue                  `cty:"tacacs" hcl:"tacacs"`
	UseMxedgeProxy        *bool                                            `cty:"use_mxedge_proxy" hcl:"use_mxedge_proxy"`
}

type SiteNetworktemplateLocalAccountsValue struct {
	Password *string `cty:"password" hcl:"password"`
	Role     *string `cty:"role" hcl:"role"`
}

type SiteNetworktemplateProtectReValue struct {
	AllowedServices []string                         `cty:"allowed_services" hcl:"allowed_services"`
	Custom          []SiteNetworktemplateCustomValue `cty:"custom" hcl:"custom"`
	Enabled         *bool                            `cty:"enabled" hcl:"enabled"`
	HitCount        *bool                            `cty:"hit_count" hcl:"hit_count"`
	TrustedHosts    []string                         `cty:"trusted_hosts" hcl:"trusted_hosts"`
}

type SiteNetworktemplateCustomValue struct {
	PortRange *string  `cty:"port_range" hcl:"port_range"`
	Protocol  *string  `cty:"protocol" hcl:"protocol"`
	Subnets   []string `cty:"subnets" hcl:"subnets"`
}

type SiteNetworktemplateTacacsValue struct {
	DefaultRole    *string                                  `cty:"default_role" hcl:"default_role"`
	Enabled        *bool                                    `cty:"enabled" hcl:"enabled"`
	Network        *string                                  `cty:"network" hcl:"network"`
	TacacctServers []SiteNetworktemplateTacacctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	TacplusServers []SiteNetworktemplateTacplusServersValue `cty:"tacplus_servers" hcl:"tacplus_servers"`
}

type SiteNetworktemplateTacacctServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type SiteNetworktemplateTacplusServersValue struct {
	Host    *string `cty:"host" hcl:"host"`
	Port    *string `cty:"port" hcl:"port"`
	Secret  *string `cty:"secret" hcl:"secret"`
	Timeout *int64  `cty:"timeout" hcl:"timeout"`
}

type SiteNetworktemplateVrfConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteNetworktemplateVrfInstancesValue struct {
	EvpnAutoLoopbackSubnet  *string                                            `cty:"evpn_auto_loopback_subnet" hcl:"evpn_auto_loopback_subnet"`
	EvpnAutoLoopbackSubnet6 *string                                            `cty:"evpn_auto_loopback_subnet6" hcl:"evpn_auto_loopback_subnet6"`
	Networks                []string                                           `cty:"networks" hcl:"networks"`
	VrfExtraRoutes          map[string]SiteNetworktemplateVrfExtraRoutesValue  `cty:"extra_routes" hcl:"extra_routes"`
	VrfExtraRoutes6         map[string]SiteNetworktemplateVrfExtraRoutes6Value `cty:"extra_routes6" hcl:"extra_routes6"`
}

type SiteNetworktemplateVrfExtraRoutesValue struct {
	Via string `cty:"via" hcl:"via"`
}

type SiteNetworktemplateVrfExtraRoutes6Value struct {
	Via *string `cty:"via" hcl:"via"`
}
