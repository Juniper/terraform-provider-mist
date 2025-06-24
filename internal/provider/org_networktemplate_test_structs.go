package provider

type OrgNetworktemplateModel struct {
	AclPolicies           []AclPoliciesValue                             `hcl:"acl_policies"`
	AclTags               map[string]AclTagsValue                        `hcl:"acl_tags"`
	AdditionalConfigCmds  []string                                       `hcl:"additional_config_cmds"`
	DhcpSnooping          *DhcpSnoopingValue                             `hcl:"dhcp_snooping"`
	DnsServers            []string                                       `hcl:"dns_servers"`
	DnsSuffix             []string                                       `hcl:"dns_suffix"`
	ExtraRoutes           map[string]ExtraRoutesValue                    `hcl:"extra_routes"`
	ExtraRoutes6          map[string]ExtraRoutes6Value                   `hcl:"extra_routes6"`
	MistNac               *MistNacValue                                  `hcl:"mist_nac"`
	Name                  string                                         `hcl:"name"`
	Networks              map[string]NetworksValue                       `hcl:"networks"`
	NtpServers            []string                                       `hcl:"ntp_servers"`
	OrgId                 string                                         `hcl:"org_id"`
	OspfAreas             map[string]OspfAreasValue                      `hcl:"ospf_areas"`
	PortMirroring         map[string]PortMirroringValue                  `hcl:"port_mirroring"`
	PortUsages            map[string]OrgNetworktemplatePortUsagesValue   `hcl:"port_usages"`
	RadiusConfig          *RadiusConfigValue                             `hcl:"radius_config"`
	RemoteSyslog          *RemoteSyslogValue                             `hcl:"remote_syslog"`
	RemoveExistingConfigs *bool                                          `hcl:"remove_existing_configs"`
	SnmpConfig            *SnmpConfigValue                               `hcl:"snmp_config"`
	SwitchMatching        *SwitchMatchingValue                           `hcl:"switch_matching"`
	SwitchMgmt            *SwitchMgmtValue                               `hcl:"switch_mgmt"`
	VrfConfig             *VrfConfigValue                                `hcl:"vrf_config"`
	VrfInstances          map[string]OrgNetworktemplateVrfInstancesValue `hcl:"vrf_instances"`
}

type OrgNetworktemplatePortUsagesValue struct {
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
	MacLimit                                int64              `cty:"mac_limit"`
	Mode                                    *string            `cty:"mode"`
	Mtu                                     int64              `cty:"mtu"`
	Networks                                []string           `cty:"networks"`
	PersistMac                              *bool              `cty:"persist_mac"`
	PoeDisabled                             *bool              `cty:"poe_disabled"`
	PortAuth                                *string            `cty:"port_auth"`
	PortNetwork                             *string            `cty:"port_network"`
	ReauthInterval                          int64              `cty:"reauth_interval"`
	ResetDefaultWhen                        *string            `cty:"reset_default_when"`
	Rules                                   []RulesValue       `cty:"rules"`
	ServerFailNetwork                       *string            `cty:"server_fail_network"`
	ServerRejectNetwork                     *string            `cty:"server_reject_network"`
	Speed                                   *string            `cty:"speed"`
	StormControl                            *StormControlValue `cty:"storm_control"`
	StpEdge                                 *bool              `cty:"stp_edge"`
	StpNoRootPort                           *bool              `cty:"stp_no_root_port"`
	StpP2p                                  *bool              `cty:"stp_p2p"`
	UiEvpntopoId                            *string            `cty:"ui_evpntopo_id"`
	UseVstp                                 *bool              `cty:"use_vstp"`
	VoipNetwork                             *string            `cty:"voip_network"`
}

type SwitchMatchingValue struct {
	Enable        *bool                `cty:"enable"`
	MatchingRules []MatchingRulesValue `cty:"rules"`
}

type MatchingRulesValue struct {
	AdditionalConfigCmds []string                            `cty:"additional_config_cmds"`
	IpConfig             *OrgNetworktemplateIpConfigValue    `cty:"ip_config"`
	MatchModel           *string                             `cty:"match_model"`
	MatchName            *string                             `cty:"match_name"`
	MatchNameOffset      int64                               `cty:"match_name_offset"`
	MatchRole            *string                             `cty:"match_role"`
	MatchType            *string                             `cty:"match_type"`
	MatchValue           *string                             `cty:"match_value"`
	Name                 *string                             `cty:"name"`
	OobIpConfig          *OrgNetworktemplateOobIpConfigValue `cty:"oob_ip_config"`
	PortConfig           map[string]PortConfigValue          `cty:"port_config"`
	PortMirroring        map[string]PortMirroringValue       `cty:"port_mirroring"`
}

type OrgNetworktemplateIpConfigValue struct {
	Network      *string `cty:"network"`
	IpConfigType *string `cty:"type"`
}

type OrgNetworktemplateOobIpConfigValue struct {
	OobIpConfigType      *string `cty:"type"`
	UseMgmtVrf           *bool   `cty:"use_mgmt_vrf"`
	UseMgmtVrfForHostOut *bool   `cty:"use_mgmt_vrf_for_host_out"`
}

type OrgNetworktemplateVrfInstancesValue struct {
	Networks       []string                       `cty:"networks"`
	VrfExtraRoutes map[string]VrfExtraRoutesValue `cty:"extra_routes"`
}
