package provider

type OrgMxedgeModel struct {
	ForSite                   *bool                                          `hcl:"for_site"`
	Magic                     *string                                        `hcl:"magic"`
	Model                     string                                         `hcl:"model"`
	MxagentRegistered         *bool                                          `hcl:"mxagent_registered"`
	MxclusterId               *string                                        `hcl:"mxcluster_id"`
	MxedgeMgmt                *OrgMxedgeMxedgeMgmtValue                      `hcl:"mxedge_mgmt"`
	Name                      string                                         `hcl:"name"`
	Note                      *string                                        `hcl:"note"`
	NtpServers                []string                                       `hcl:"ntp_servers"`
	OobIpConfig               *OrgMxedgeOobIpConfigValue                     `hcl:"oob_ip_config"`
	OrgId                     string                                         `hcl:"org_id"`
	Proxy                     *OrgMxedgeProxyValue                           `hcl:"proxy"`
	Services                  []string                                       `hcl:"services"`
	SiteId                    *string                                        `hcl:"site_id"`
	TuntermDhcpdConfig        map[string]OrgMxedgeTuntermDhcpdConfigValue    `hcl:"tunterm_dhcpd_config"`
	TuntermExtraRoutes        map[string]OrgMxedgeTuntermExtraRoutesValue    `hcl:"tunterm_extra_routes"`
	TuntermIgmpSnoopingConfig *OrgMxedgeTuntermIgmpSnoopingConfigValue       `hcl:"tunterm_igmp_snooping_config"`
	TuntermIpConfig           *OrgMxedgeTuntermIpConfigValue                 `hcl:"tunterm_ip_config"`
	TuntermMonitoring         []string                                       `hcl:"tunterm_monitoring"`
	TuntermMulticastConfig    *OrgMxedgeTuntermMulticastConfigValue          `hcl:"tunterm_multicast_config"`
	TuntermOtherIpConfigs     map[string]OrgMxedgeTuntermOtherIpConfigsValue `hcl:"tunterm_other_ip_configs"`
	TuntermPortConfig         *OrgMxedgeTuntermPortConfigValue               `hcl:"tunterm_port_config"`
	TuntermRegistered         *bool                                          `hcl:"tunterm_registered"`
	TuntermSwitchConfig       map[string]OrgMxedgeTuntermSwitchConfigValue   `hcl:"tunterm_switch_config"`
	Versions                  *OrgMxedgeVersionsValue                        `hcl:"versions"`
}

type OrgMxedgeMxedgeMgmtValue struct {
	ConfigAutoRevert *bool   `cty:"config_auto_revert" hcl:"config_auto_revert"`
	FipsEnabled      *bool   `cty:"fips_enabled" hcl:"fips_enabled"`
	MistPassword     *string `cty:"mist_password" hcl:"mist_password"`
	OobIpType        *string `cty:"oob_ip_type" hcl:"oob_ip_type"`
	OobIpType6       *string `cty:"oob_ip_type6" hcl:"oob_ip_type6"`
	RootPassword     *string `cty:"root_password" hcl:"root_password"`
}

type OrgMxedgeOobIpConfigValue struct {
	Autoconf6       *bool    `cty:"autoconf6" hcl:"autoconf6"`
	Dhcp6           *bool    `cty:"dhcp6" hcl:"dhcp6"`
	Dns             []string `cty:"dns" hcl:"dns"`
	Gateway         *string  `cty:"gateway" hcl:"gateway"`
	Gateway6        *string  `cty:"gateway6" hcl:"gateway6"`
	Ip              *string  `cty:"ip" hcl:"ip"`
	Ip6             *string  `cty:"ip6" hcl:"ip6"`
	Netmask         *string  `cty:"netmask" hcl:"netmask"`
	Netmask6        *string  `cty:"netmask6" hcl:"netmask6"`
	OobIpConfigType *string  `cty:"type" hcl:"type"`
	Type6           *string  `cty:"type6" hcl:"type6"`
}

type OrgMxedgeProxyValue struct {
	Url *string `cty:"url" hcl:"url"`
}

type OrgMxedgeTuntermDhcpdConfigValue struct {
	Enabled                *bool    `cty:"enabled" hcl:"enabled"`
	Servers                []string `cty:"servers" hcl:"servers"`
	TuntermDhcpdConfigType *string  `cty:"type" hcl:"type"`
}

type OrgMxedgeTuntermExtraRoutesValue struct {
	Via *string `cty:"via" hcl:"via"`
}

type OrgMxedgeTuntermIgmpSnoopingConfigValue struct {
	Enabled *bool                  `cty:"enabled" hcl:"enabled"`
	Querier *OrgMxedgeQuerierValue `cty:"querier" hcl:"querier"`
	VlanIds []int64                `cty:"vlan_ids" hcl:"vlan_ids"`
}

type OrgMxedgeQuerierValue struct {
	MaxResponseTime *int64 `cty:"max_response_time" hcl:"max_response_time"`
	Mtu             *int64 `cty:"mtu" hcl:"mtu"`
	QueryInterval   *int64 `cty:"query_interval" hcl:"query_interval"`
	Robustness      *int64 `cty:"robustness" hcl:"robustness"`
	Version         *int64 `cty:"version" hcl:"version"`
}

type OrgMxedgeTuntermIpConfigValue struct {
	Gateway  string  `cty:"gateway" hcl:"gateway"`
	Gateway6 *string `cty:"gateway6" hcl:"gateway6"`
	Ip       string  `cty:"ip" hcl:"ip"`
	Ip6      *string `cty:"ip6" hcl:"ip6"`
	Netmask  string  `cty:"netmask" hcl:"netmask"`
	Netmask6 *string `cty:"netmask6" hcl:"netmask6"`
}

type OrgMxedgeTuntermMulticastConfigValue struct {
	Mdns *OrgMxedgeMdnsValue `cty:"mdns" hcl:"mdns"`
	Ssdp *OrgMxedgeSsdpValue `cty:"ssdp" hcl:"ssdp"`
}

type OrgMxedgeMdnsValue struct {
	Enabled *bool    `cty:"enabled" hcl:"enabled"`
	VlanIds []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type OrgMxedgeSsdpValue struct {
	Enabled *bool    `cty:"enabled" hcl:"enabled"`
	VlanIds []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type OrgMxedgeTuntermOtherIpConfigsValue struct {
	Ip      string `cty:"ip" hcl:"ip"`
	Netmask string `cty:"netmask" hcl:"netmask"`
}

type OrgMxedgeTuntermPortConfigValue struct {
	DownstreamPorts            []string `cty:"downstream_ports" hcl:"downstream_ports"`
	SeparateUpstreamDownstream *bool    `cty:"separate_upstream_downstream" hcl:"separate_upstream_downstream"`
	UpstreamPortVlanId         *int64   `cty:"upstream_port_vlan_id" hcl:"upstream_port_vlan_id"`
	UpstreamPorts              []string `cty:"upstream_ports" hcl:"upstream_ports"`
}

type OrgMxedgeTuntermSwitchConfigValue struct {
	PortVlanId *int64   `cty:"port_vlan_id" hcl:"port_vlan_id"`
	VlanIds    []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type OrgMxedgeVersionsValue struct {
	Mxagent *string `cty:"mxagent" hcl:"mxagent"`
	Tunterm *string `cty:"tunterm" hcl:"tunterm"`
}
