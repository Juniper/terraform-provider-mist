package provider

type OrgWlanModel struct {
	AcctImmediateUpdate                  *bool                    `hcl:"acct_immediate_update"`
	AcctInterimInterval                  *int64                   `hcl:"acct_interim_interval"`
	AcctServers                          []AcctServersValue       `hcl:"acct_servers"`
	Airwatch                             *AirwatchValue           `hcl:"airwatch"`
	AllowIpv6Ndp                         *bool                    `hcl:"allow_ipv6_ndp"`
	AllowMdns                            *bool                    `hcl:"allow_mdns"`
	AllowSsdp                            *bool                    `hcl:"allow_ssdp"`
	ApIds                                []string                 `hcl:"ap_ids"`
	AppLimit                             *AppLimitValue           `hcl:"app_limit"`
	AppQos                               *AppQosValue             `hcl:"app_qos"`
	ApplyTo                              *string                  `hcl:"apply_to"`
	ArpFilter                            *bool                    `hcl:"arp_filter"`
	Auth                                 *OrgWlanAuthValue        `hcl:"auth"`
	AuthServerSelection                  *string                  `hcl:"auth_server_selection"`
	AuthServers                          []AuthServersValue       `hcl:"auth_servers"`
	AuthServersNasId                     *string                  `hcl:"auth_servers_nas_id"`
	AuthServersNasIp                     *string                  `hcl:"auth_servers_nas_ip"`
	AuthServersRetries                   *int64                   `hcl:"auth_servers_retries"`
	AuthServersTimeout                   *int64                   `hcl:"auth_servers_timeout"`
	BandSteer                            *bool                    `hcl:"band_steer"`
	BandSteerForceBand5                  *bool                    `hcl:"band_steer_force_band5"`
	Bands                                []string                 `hcl:"bands"`
	BlockBlacklistClients                *bool                    `hcl:"block_blacklist_clients"`
	Bonjour                              *BonjourValue            `hcl:"bonjour"`
	CiscoCwa                             *CiscoCwaValue           `hcl:"cisco_cwa"`
	ClientLimitDown                      *int64                   `hcl:"client_limit_down"`
	ClientLimitDownEnabled               *bool                    `hcl:"client_limit_down_enabled"`
	ClientLimitUp                        *int64                   `hcl:"client_limit_up"`
	ClientLimitUpEnabled                 *bool                    `hcl:"client_limit_up_enabled"`
	CoaServers                           []CoaServersValue        `hcl:"coa_servers"`
	Disable11ax                          *bool                    `hcl:"disable_11ax"`
	DisableHtVhtRates                    *bool                    `hcl:"disable_ht_vht_rates"`
	DisableUapsd                         *bool                    `hcl:"disable_uapsd"`
	DisableV1RoamNotify                  *bool                    `hcl:"disable_v1_roam_notify"`
	DisableV2RoamNotify                  *bool                    `hcl:"disable_v2_roam_notify"`
	DisableWhenGatewayUnreachable        *bool                    `hcl:"disable_when_gateway_unreachable"`
	DisableWhenMxtunnelDown              *bool                    `hcl:"disable_when_mxtunnel_down"`
	DisableWmm                           *bool                    `hcl:"disable_wmm"`
	DnsServerRewrite                     *DnsServerRewriteValue   `hcl:"dns_server_rewrite"`
	Dtim                                 *int64                   `hcl:"dtim"`
	DynamicPsk                           *DynamicPskValue         `hcl:"dynamic_psk"`
	DynamicVlan                          *DynamicVlanValue        `hcl:"dynamic_vlan"`
	EnableLocalKeycaching                *bool                    `hcl:"enable_local_keycaching"`
	EnableWirelessBridging               *bool                    `hcl:"enable_wireless_bridging"`
	EnableWirelessBridgingDhcpTracking   *bool                    `hcl:"enable_wireless_bridging_dhcp_tracking"`
	Enabled                              *bool                    `hcl:"enabled"`
	FastDot1xTimers                      *bool                    `hcl:"fast_dot1x_timers"`
	HideSsid                             *bool                    `hcl:"hide_ssid"`
	HostnameIe                           *bool                    `hcl:"hostname_ie"`
	Hotspot20                            *Hotspot20Value          `hcl:"hotspot20"`
	InjectDhcpOption82                   *InjectDhcpOption82Value `hcl:"inject_dhcp_option_82"`
	Interface                            *string                  `hcl:"interface"`
	Isolation                            *bool                    `hcl:"isolation"`
	L2Isolation                          *bool                    `hcl:"l2_isolation"`
	LegacyOverds                         *bool                    `hcl:"legacy_overds"`
	LimitBcast                           *bool                    `hcl:"limit_bcast"`
	LimitProbeResponse                   *bool                    `hcl:"limit_probe_response"`
	MaxIdletime                          *int64                   `hcl:"max_idletime"`
	MaxNumClients                        *int64                   `hcl:"max_num_clients"`
	MistNac                              *OrgWlanMistNacValue     `hcl:"mist_nac"`
	MxtunnelIds                          []string                 `hcl:"mxtunnel_ids"`
	MxtunnelName                         []string                 `hcl:"mxtunnel_name"`
	NoStaticDns                          *bool                    `hcl:"no_static_dns"`
	NoStaticIp                           *bool                    `hcl:"no_static_ip"`
	OrgId                                string                   `hcl:"org_id"`
	Portal                               *PortalValue             `hcl:"portal"`
	PortalAllowedHostnames               []string                 `hcl:"portal_allowed_hostnames"`
	PortalAllowedSubnets                 []string                 `hcl:"portal_allowed_subnets"`
	PortalDeniedHostnames                []string                 `hcl:"portal_denied_hostnames"`
	Qos                                  *QosValue                `hcl:"qos"`
	Radsec                               *RadsecValue             `hcl:"radsec"`
	Rateset                              map[string]RatesetValue  `hcl:"rateset"`
	ReconnectClientsWhenRoamingMxcluster *bool                    `hcl:"reconnect_clients_when_roaming_mxcluster"`
	RoamMode                             *string                  `hcl:"roam_mode"`
	Schedule                             *ScheduleValue           `hcl:"schedule"`
	SleExcluded                          *bool                    `hcl:"sle_excluded"`
	Ssid                                 string                   `hcl:"ssid"`
	TemplateId                           string                   `hcl:"template_id"`
	UseEapolV1                           *bool                    `hcl:"use_eapol_v1"`
	VlanEnabled                          *bool                    `hcl:"vlan_enabled"`
	VlanId                               *string                  `hcl:"vlan_id"`
	VlanIds                              []string                 `hcl:"vlan_ids"`
	VlanPooling                          *bool                    `hcl:"vlan_pooling"`
	WlanLimitDown                        *int64                   `hcl:"wlan_limit_down"`
	WlanLimitDownEnabled                 *bool                    `hcl:"wlan_limit_down_enabled"`
	WlanLimitUp                          *int64                   `hcl:"wlan_limit_up"`
	WlanLimitUpEnabled                   *bool                    `hcl:"wlan_limit_up_enabled"`
	WxtagIds                             []string                 `hcl:"wxtag_ids"`
	WxtunnelId                           *string                  `hcl:"wxtunnel_id"`
	WxtunnelRemoteId                     *string                  `hcl:"wxtunnel_remote_id"`
}

type AirwatchValue struct {
	ApiKey     *string `cty:"api_key"`
	ConsoleUrl *string `cty:"console_url"`
	Enabled    *bool   `cty:"enabled"`
	Password   *string `cty:"password"`
	Username   *string `cty:"username"`
}

type AppLimitValue struct {
	Apps     map[string]int64 `cty:"apps"`
	Enabled  *bool            `cty:"enabled"`
	WxtagIds map[string]int64 `cty:"wxtag_ids"`
}

type AppQosValue struct {
	Apps    map[string]AppsValue `cty:"apps"`
	Enabled *bool                `cty:"enabled"`
	Others  []OthersValue        `cty:"others"`
}

type AppsValue struct {
	Dscp      *int64  `cty:"dscp"`
	DstSubnet *string `cty:"dst_subnet"`
	SrcSubnet *string `cty:"src_subnet"`
}

type OthersValue struct {
	Dscp       *int64  `cty:"dscp"`
	DstSubnet  *string `cty:"dst_subnet"`
	PortRanges *string `cty:"port_ranges"`
	Protocol   *string `cty:"protocol"`
	SrcSubnet  *string `cty:"src_subnet"`
}

type OrgWlanAuthValue struct {
	AnticlogThreshold  *int64   `cty:"anticlog_threshold"`
	EapReauth          *bool    `cty:"eap_reauth"`
	EnableMacAuth      *bool    `cty:"enable_mac_auth"`
	KeyIdx             *int64   `cty:"key_idx"`
	Keys               []string `cty:"keys"`
	MultiPskOnly       *bool    `cty:"multi_psk_only"`
	Owe                *string  `cty:"owe"`
	Pairwise           []string `cty:"pairwise"`
	PrivateWlan        *bool    `cty:"private_wlan"`
	Psk                *string  `cty:"psk"`
	AuthType           *string  `cty:"type"`
	WepAsSecondaryAuth *bool    `cty:"wep_as_secondary_auth"`
}

type BonjourValue struct {
	AdditionalVlanIds []string                 `cty:"additional_vlan_ids"`
	Enabled           *bool                    `cty:"enabled"`
	Services          map[string]ServicesValue `cty:"services"`
}

type ServicesValue struct {
	DisableLocal *bool    `cty:"disable_local"`
	RadiusGroups []string `cty:"radius_groups"`
	Scope        *string  `cty:"scope"`
}

type CiscoCwaValue struct {
	AllowedHostnames []string `cty:"allowed_hostnames"`
	AllowedSubnets   []string `cty:"allowed_subnets"`
	BlockedSubnets   []string `cty:"blocked_subnets"`
	Enabled          *bool    `cty:"enabled"`
}

type CoaServersValue struct {
	DisableEventTimestampCheck *bool  `cty:"disable_event_timestamp_check"`
	Enabled                    *bool  `cty:"enabled"`
	Ip                         string `cty:"ip"`
	Port                       *int64 `cty:"port"`
	Secret                     string `cty:"secret"`
}

type DnsServerRewriteValue struct {
	Enabled      *bool             `cty:"enabled"`
	RadiusGroups map[string]string `cty:"radius_groups"`
}

type DynamicPskValue struct {
	DefaultPsk    *string `cty:"default_psk"`
	DefaultVlanId *string `cty:"default_vlan_id"`
	Enabled       *bool   `cty:"enabled"`
	ForceLookup   *bool   `cty:"force_lookup"`
	Source        *string `cty:"source"`
}

type DynamicVlanValue struct {
	DefaultVlanIds  []string          `cty:"default_vlan_ids"`
	Enabled         *bool             `cty:"enabled"`
	LocalVlanIds    []string          `cty:"local_vlan_ids"`
	DynamicVlanType *string           `cty:"type"`
	Vlans           map[string]string `cty:"vlans"`
}

type Hotspot20Value struct {
	DomainName []string `cty:"domain_name"`
	Enabled    *bool    `cty:"enabled"`
	NaiRealms  []string `cty:"nai_realms"`
	Operators  []string `cty:"operators"`
	Rcoi       []string `cty:"rcoi"`
	VenueName  *string  `cty:"venue_name"`
}

type InjectDhcpOption82Value struct {
	CircuitId *string `cty:"circuit_id"`
	Enabled   *bool   `cty:"enabled"`
}

type OrgWlanMistNacValue struct {
	Enabled *bool `cty:"enabled"`
}

type PortalValue struct {
	AllowWlanIdRoam             *bool             `cty:"allow_wlan_id_roam"`
	AmazonClientId              *string           `cty:"amazon_client_id"`
	AmazonClientSecret          *string           `cty:"amazon_client_secret"`
	AmazonEmailDomains          []string          `cty:"amazon_email_domains"`
	AmazonEnabled               *bool             `cty:"amazon_enabled"`
	AmazonExpire                *int64            `cty:"amazon_expire"`
	Auth                        *string           `cty:"auth"`
	AzureClientId               *string           `cty:"azure_client_id"`
	AzureClientSecret           *string           `cty:"azure_client_secret"`
	AzureEnabled                *bool             `cty:"azure_enabled"`
	AzureExpire                 *int64            `cty:"azure_expire"`
	AzureTenantId               *string           `cty:"azure_tenant_id"`
	BroadnetPassword            *string           `cty:"broadnet_password"`
	BroadnetSid                 *string           `cty:"broadnet_sid"`
	BroadnetUserId              *string           `cty:"broadnet_user_id"`
	BypassWhenCloudDown         *bool             `cty:"bypass_when_cloud_down"`
	ClickatellApiKey            *string           `cty:"clickatell_api_key"`
	CrossSite                   *bool             `cty:"cross_site"`
	EmailEnabled                *bool             `cty:"email_enabled"`
	Enabled                     *bool             `cty:"enabled"`
	Expire                      *int64            `cty:"expire"`
	ExternalPortalUrl           *string           `cty:"external_portal_url"`
	FacebookClientId            *string           `cty:"facebook_client_id"`
	FacebookClientSecret        *string           `cty:"facebook_client_secret"`
	FacebookEmailDomains        []string          `cty:"facebook_email_domains"`
	FacebookEnabled             *bool             `cty:"facebook_enabled"`
	FacebookExpire              *int64            `cty:"facebook_expire"`
	Forward                     *bool             `cty:"forward"`
	ForwardUrl                  *string           `cty:"forward_url"`
	GoogleClientId              *string           `cty:"google_client_id"`
	GoogleClientSecret          *string           `cty:"google_client_secret"`
	GoogleEmailDomains          []string          `cty:"google_email_domains"`
	GoogleEnabled               *bool             `cty:"google_enabled"`
	GoogleExpire                *int64            `cty:"google_expire"`
	GupshupPassword             *string           `cty:"gupshup_password"`
	GupshupUserid               *string           `cty:"gupshup_userid"`
	MicrosoftClientId           *string           `cty:"microsoft_client_id"`
	MicrosoftClientSecret       *string           `cty:"microsoft_client_secret"`
	MicrosoftEmailDomains       []string          `cty:"microsoft_email_domains"`
	MicrosoftEnabled            *bool             `cty:"microsoft_enabled"`
	MicrosoftExpire             *int64            `cty:"microsoft_expire"`
	PassphraseEnabled           *bool             `cty:"passphrase_enabled"`
	PassphraseExpire            *int64            `cty:"passphrase_expire"`
	Password                    *string           `cty:"password"`
	PredefinedSponsorsEnabled   *bool             `cty:"predefined_sponsors_enabled"`
	PredefinedSponsorsHideEmail *bool             `cty:"predefined_sponsors_hide_email"`
	Privacy                     *bool             `cty:"privacy"`
	PuzzelPassword              *string           `cty:"puzzel_password"`
	PuzzelServiceId             *string           `cty:"puzzel_service_id"`
	PuzzelUsername              *string           `cty:"puzzel_username"`
	SmsEnabled                  *bool             `cty:"sms_enabled"`
	SmsExpire                   *int64            `cty:"sms_expire"`
	SmsMessageFormat            *string           `cty:"sms_message_format"`
	SmsProvider                 *string           `cty:"sms_provider"`
	SponsorAutoApprove          *bool             `cty:"sponsor_auto_approve"`
	SponsorEmailDomains         []string          `cty:"sponsor_email_domains"`
	SponsorEnabled              *bool             `cty:"sponsor_enabled"`
	SponsorExpire               *int64            `cty:"sponsor_expire"`
	SponsorLinkValidityDuration *string           `cty:"sponsor_link_validity_duration"`
	SponsorNotifyAll            *bool             `cty:"sponsor_notify_all"`
	SponsorStatusNotify         *bool             `cty:"sponsor_status_notify"`
	Sponsors                    map[string]string `cty:"sponsors"`
	SsoDefaultRole              *string           `cty:"sso_default_role"`
	SsoForcedRole               *string           `cty:"sso_forced_role"`
	SsoIdpCert                  *string           `cty:"sso_idp_cert"`
	SsoIdpSignAlgo              *string           `cty:"sso_idp_sign_algo"`
	SsoIdpSsoUrl                *string           `cty:"sso_idp_sso_url"`
	SsoIssuer                   *string           `cty:"sso_issuer"`
	SsoNameidFormat             *string           `cty:"sso_nameid_format"`
	TelstraClientId             *string           `cty:"telstra_client_id"`
	TelstraClientSecret         *string           `cty:"telstra_client_secret"`
	TwilioAuthToken             *string           `cty:"twilio_auth_token"`
	TwilioPhoneNumber           *string           `cty:"twilio_phone_number"`
	TwilioSid                   *string           `cty:"twilio_sid"`
}

type QosValue struct {
	Class     *string `cty:"class"`
	Overwrite *bool   `cty:"overwrite"`
}

type RadsecValue struct {
	CoaEnabled    *bool                 `cty:"coa_enabled"`
	Enabled       *bool                 `cty:"enabled"`
	IdleTimeout   *int64                `cty:"idle_timeout"`
	MxclusterIds  []string              `cty:"mxcluster_ids"`
	ProxyHosts    []string              `cty:"proxy_hosts"`
	ServerName    *string               `cty:"server_name"`
	Servers       []OrgWlanServersValue `cty:"servers"`
	UseMxedge     *bool                 `cty:"use_mxedge"`
	UseSiteMxedge *bool                 `cty:"use_site_mxedge"`
}

type OrgWlanServersValue struct {
	Host *string `cty:"host"`
	Port *int64  `cty:"port"`
}

type RatesetValue struct {
	Ht       *string  `cty:"ht"`
	Legacy   []string `cty:"legacy"`
	MinRssi  *int64   `cty:"min_rssi"`
	Template *string  `cty:"template"`
	Vht      *string  `cty:"vht"`
}

type ScheduleValue struct {
	Enabled *bool       `cty:"enabled"`
	Hours   *HoursValue `cty:"hours"`
}

type HoursValue struct {
	Fri *string `cty:"fri"`
	Mon *string `cty:"mon"`
	Sat *string `cty:"sat"`
	Sun *string `cty:"sun"`
	Thu *string `cty:"thu"`
	Tue *string `cty:"tue"`
	Wed *string `cty:"wed"`
}
