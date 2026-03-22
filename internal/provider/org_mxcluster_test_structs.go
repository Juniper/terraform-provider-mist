package provider

type OrgMxclusterModel struct {
	MistDas                   *OrgMxclusterMistDasValue                      `hcl:"mist_das"`
	MistNac                   *OrgMxclusterMistNacValue                      `hcl:"mist_nac"`
	MxedgeMgmt                *OrgMxclusterMxedgeMgmtValue                   `hcl:"mxedge_mgmt"`
	Name                      string                                         `hcl:"name"`
	OrgId                     string                                         `hcl:"org_id"`
	Proxy                     *OrgMxclusterProxyValue                        `hcl:"proxy"`
	Radsec                    *OrgMxclusterRadsecValue                       `hcl:"radsec"`
	RadsecTls                 *OrgMxclusterRadsecTlsValue                    `hcl:"radsec_tls"`
	SiteId                    *string                                        `hcl:"site_id"`
	TuntermApSubnets          []string                                       `hcl:"tunterm_ap_subnets"`
	TuntermDhcpdConfig        map[string]OrgMxclusterTuntermDhcpdConfigValue `hcl:"tunterm_dhcpd_config"`
	TuntermExtraRoutes        map[string]OrgMxclusterTuntermExtraRoutesValue `hcl:"tunterm_extra_routes"`
	TuntermHosts              []string                                       `hcl:"tunterm_hosts"`
	TuntermHostsOrder         []int64                                        `hcl:"tunterm_hosts_order"`
	TuntermHostsSelection     *string                                        `hcl:"tunterm_hosts_selection"`
	TuntermMonitoring         [][]OrgMxclusterTuntermMonitoringValue         `hcl:"tunterm_monitoring"`
	TuntermMonitoringDisabled *bool                                          `hcl:"tunterm_monitoring_disabled"`
}

type OrgMxclusterMistDasValue struct {
	CoaServers []OrgMxclusterCoaServersValue `cty:"coa_servers" hcl:"coa_servers"`
	Enabled    *bool                         `cty:"enabled" hcl:"enabled"`
}

type OrgMxclusterCoaServersValue struct {
	DisableEventTimestampCheck  *bool   `cty:"disable_event_timestamp_check" hcl:"disable_event_timestamp_check"`
	Enabled                     *bool   `cty:"enabled" hcl:"enabled"`
	Host                        *string `cty:"host" hcl:"host"`
	Port                        *int64  `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      *string `cty:"secret" hcl:"secret"`
}

type OrgMxclusterMistNacValue struct {
	AcctServerPort *int64 `cty:"acct_server_port" hcl:"acct_server_port"`
	AuthServerPort *int64 `cty:"auth_server_port" hcl:"auth_server_port"`
	// ClientIps      map[string]OrgMxclusterClientIpsValue `cty:"client_ips" hcl:"client_ips"` // Commented out: empty struct can't be encoded in HCL
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Secret  *string `cty:"secret" hcl:"secret"`
}

// OrgMxclusterClientIpsValue - Commented out: empty struct without cty tags can't be encoded in HCL
// type OrgMxclusterClientIpsValue struct {
// }

type OrgMxclusterMxedgeMgmtValue struct {
	ConfigAutoRevert *bool   `cty:"config_auto_revert" hcl:"config_auto_revert"`
	FipsEnabled      *bool   `cty:"fips_enabled" hcl:"fips_enabled"`
	MistPassword     *string `cty:"mist_password" hcl:"mist_password"`
	OobIpType        *string `cty:"oob_ip_type" hcl:"oob_ip_type"`
	OobIpType6       *string `cty:"oob_ip_type6" hcl:"oob_ip_type6"`
	RootPassword     *string `cty:"root_password" hcl:"root_password"`
}

type OrgMxclusterProxyValue struct {
	Disabled *bool   `cty:"disabled" hcl:"disabled"`
	Url      *string `cty:"url" hcl:"url"`
}

type OrgMxclusterRadsecValue struct {
	AcctServers     []OrgMxclusterAcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServers     []OrgMxclusterAuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	Enabled         *bool                          `cty:"enabled" hcl:"enabled"`
	MatchSsid       *bool                          `cty:"match_ssid" hcl:"match_ssid"`
	NasIpSource     *string                        `cty:"nas_ip_source" hcl:"nas_ip_source"`
	ProxyHosts      []string                       `cty:"proxy_hosts" hcl:"proxy_hosts"`
	ServerSelection *string                        `cty:"server_selection" hcl:"server_selection"`
	SrcIpSource     *string                        `cty:"src_ip_source" hcl:"src_ip_source"`
}

type OrgMxclusterAcctServersValue struct {
	Host   *string  `cty:"host" hcl:"host"`
	Port   *int64   `cty:"port" hcl:"port"`
	Secret *string  `cty:"secret" hcl:"secret"`
	Ssids  []string `cty:"ssids" hcl:"ssids"`
}

type OrgMxclusterAuthServersValue struct {
	Host                 *string  `cty:"host" hcl:"host"`
	InbandStatusCheck    *bool    `cty:"inband_status_check" hcl:"inband_status_check"`
	InbandStatusInterval *int64   `cty:"inband_status_interval" hcl:"inband_status_interval"`
	KeywrapEnabled       *bool    `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat        *string  `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek           *string  `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack          *string  `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                 *int64   `cty:"port" hcl:"port"`
	Retry                *int64   `cty:"retry" hcl:"retry"`
	Secret               *string  `cty:"secret" hcl:"secret"`
	Ssids                []string `cty:"ssids" hcl:"ssids"`
	Timeout              *int64   `cty:"timeout" hcl:"timeout"`
}

type OrgMxclusterRadsecTlsValue struct {
	Keypair *string `cty:"keypair" hcl:"keypair"`
}

type OrgMxclusterTuntermDhcpdConfigValue struct {
	Enabled                *bool    `cty:"enabled" hcl:"enabled"`
	Servers                []string `cty:"servers" hcl:"servers"`
	TuntermDhcpdConfigType *string  `cty:"type" hcl:"type"`
}

type OrgMxclusterTuntermExtraRoutesValue struct {
	Via *string `cty:"via" hcl:"via"`
}

type OrgMxclusterTuntermMonitoringValue struct {
	Host      *string `cty:"host" hcl:"host"`
	Port      *int64  `cty:"port" hcl:"port"`
	Protocol  *string `cty:"protocol" hcl:"protocol"`
	SrcVlanId *int64  `cty:"src_vlan_id" hcl:"src_vlan_id"`
	Timeout   *int64  `cty:"timeout" hcl:"timeout"`
}
