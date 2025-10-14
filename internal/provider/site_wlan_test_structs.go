package provider

type SiteWlanModel struct {
	AcctImmediateUpdate                  *bool                            `hcl:"acct_immediate_update"`
	AcctInterimInterval                  *int64                           `hcl:"acct_interim_interval"`
	AcctServers                          []SiteWlanAcctServersValue       `hcl:"acct_servers"`
	Airwatch                             *SiteWlanAirwatchValue           `hcl:"airwatch"`
	AllowIpv6Ndp                         *bool                            `hcl:"allow_ipv6_ndp"`
	AllowMdns                            *bool                            `hcl:"allow_mdns"`
	AllowSsdp                            *bool                            `hcl:"allow_ssdp"`
	ApIds                                []string                         `hcl:"ap_ids"`
	AppLimit                             *SiteWlanAppLimitValue           `hcl:"app_limit"`
	AppQos                               *SiteWlanAppQosValue             `hcl:"app_qos"`
	ApplyTo                              *string                          `hcl:"apply_to"`
	ArpFilter                            *bool                            `hcl:"arp_filter"`
	Auth                                 *SiteWlanAuthValue               `hcl:"auth"`
	AuthServerSelection                  *string                          `hcl:"auth_server_selection"`
	AuthServers                          []SiteWlanAuthServersValue       `hcl:"auth_servers"`
	AuthServersNasId                     *string                          `hcl:"auth_servers_nas_id"`
	AuthServersNasIp                     *string                          `hcl:"auth_servers_nas_ip"`
	AuthServersRetries                   *int64                           `hcl:"auth_servers_retries"`
	AuthServersTimeout                   *int64                           `hcl:"auth_servers_timeout"`
	BandSteer                            *bool                            `hcl:"band_steer"`
	BandSteerForceBand5                  *bool                            `hcl:"band_steer_force_band5"`
	Bands                                []string                         `hcl:"bands"`
	BlockBlacklistClients                *bool                            `hcl:"block_blacklist_clients"`
	Bonjour                              *SiteWlanBonjourValue            `hcl:"bonjour"`
	CiscoCwa                             *SiteWlanCiscoCwaValue           `hcl:"cisco_cwa"`
	ClientLimitDown                      *string                          `hcl:"client_limit_down"`
	ClientLimitDownEnabled               *bool                            `hcl:"client_limit_down_enabled"`
	ClientLimitUp                        *string                          `hcl:"client_limit_up"`
	ClientLimitUpEnabled                 *bool                            `hcl:"client_limit_up_enabled"`
	CoaServers                           []SiteWlanCoaServersValue        `hcl:"coa_servers"`
	Disable11ax                          *bool                            `hcl:"disable_11ax"`
	Disable11be                          *bool                            `hcl:"disable_11be"`
	DisableHtVhtRates                    *bool                            `hcl:"disable_ht_vht_rates"`
	DisableUapsd                         *bool                            `hcl:"disable_uapsd"`
	DisableV1RoamNotify                  *bool                            `hcl:"disable_v1_roam_notify"`
	DisableV2RoamNotify                  *bool                            `hcl:"disable_v2_roam_notify"`
	DisableWhenGatewayUnreachable        *bool                            `hcl:"disable_when_gateway_unreachable"`
	DisableWhenMxtunnelDown              *bool                            `hcl:"disable_when_mxtunnel_down"`
	DisableWmm                           *bool                            `hcl:"disable_wmm"`
	DnsServerRewrite                     *SiteWlanDnsServerRewriteValue   `hcl:"dns_server_rewrite"`
	Dtim                                 *int64                           `hcl:"dtim"`
	DynamicPsk                           *SiteWlanDynamicPskValue         `hcl:"dynamic_psk"`
	DynamicVlan                          *SiteWlanDynamicVlanValue        `hcl:"dynamic_vlan"`
	EnableLocalKeycaching                *bool                            `hcl:"enable_local_keycaching"`
	EnableWirelessBridging               *bool                            `hcl:"enable_wireless_bridging"`
	EnableWirelessBridgingDhcpTracking   *bool                            `hcl:"enable_wireless_bridging_dhcp_tracking"`
	Enabled                              *bool                            `hcl:"enabled"`
	FastDot1xTimers                      *bool                            `hcl:"fast_dot1x_timers"`
	HideSsid                             *bool                            `hcl:"hide_ssid"`
	HostnameIe                           *bool                            `hcl:"hostname_ie"`
	Hotspot20                            *SiteWlanHotspot20Value          `hcl:"hotspot20"`
	InjectDhcpOption82                   *SiteWlanInjectDhcpOption82Value `hcl:"inject_dhcp_option_82"`
	Interface                            *string                          `hcl:"interface"`
	Isolation                            *bool                            `hcl:"isolation"`
	L2Isolation                          *bool                            `hcl:"l2_isolation"`
	LegacyOverds                         *bool                            `hcl:"legacy_overds"`
	LimitBcast                           *bool                            `hcl:"limit_bcast"`
	LimitProbeResponse                   *bool                            `hcl:"limit_probe_response"`
	MaxIdletime                          *int64                           `hcl:"max_idletime"`
	MaxNumClients                        *int64                           `hcl:"max_num_clients"`
	MistNac                              *SiteWlanMistNacValue            `hcl:"mist_nac"`
	MxtunnelIds                          []string                         `hcl:"mxtunnel_ids"`
	MxtunnelName                         []string                         `hcl:"mxtunnel_name"`
	NoStaticDns                          *bool                            `hcl:"no_static_dns"`
	NoStaticIp                           *bool                            `hcl:"no_static_ip"`
	Portal                               *SiteWlanPortalValue             `hcl:"portal"`
	PortalAllowedHostnames               []string                         `hcl:"portal_allowed_hostnames"`
	PortalAllowedSubnets                 []string                         `hcl:"portal_allowed_subnets"`
	PortalDeniedHostnames                []string                         `hcl:"portal_denied_hostnames"`
	Qos                                  *SiteWlanQosValue                `hcl:"qos"`
	Radsec                               *SiteWlanRadsecValue             `hcl:"radsec"`
	Rateset                              map[string]SiteWlanRatesetValue  `hcl:"rateset"`
	ReconnectClientsWhenRoamingMxcluster *bool                            `hcl:"reconnect_clients_when_roaming_mxcluster"`
	RoamMode                             *string                          `hcl:"roam_mode"`
	Schedule                             *SiteWlanScheduleValue           `hcl:"schedule"`
	SiteId                               string                           `hcl:"site_id"`
	SleExcluded                          *bool                            `hcl:"sle_excluded"`
	Ssid                                 string                           `hcl:"ssid"`
	UseEapolV1                           *bool                            `hcl:"use_eapol_v1"`
	VlanEnabled                          *bool                            `hcl:"vlan_enabled"`
	VlanId                               *string                          `hcl:"vlan_id"`
	VlanIds                              []string                         `hcl:"vlan_ids"`
	VlanPooling                          *bool                            `hcl:"vlan_pooling"`
	WlanLimitDown                        *string                          `hcl:"wlan_limit_down"`
	WlanLimitDownEnabled                 *bool                            `hcl:"wlan_limit_down_enabled"`
	WlanLimitUp                          *string                          `hcl:"wlan_limit_up"`
	WlanLimitUpEnabled                   *bool                            `hcl:"wlan_limit_up_enabled"`
	WxtagIds                             []string                         `hcl:"wxtag_ids"`
	WxtunnelId                           *string                          `hcl:"wxtunnel_id"`
	WxtunnelRemoteId                     *string                          `hcl:"wxtunnel_remote_id"`
}

type SiteWlanAcctServersValue struct {
	Host           string  `cty:"host" hcl:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port           *string `cty:"port" hcl:"port"`
	Secret         string  `cty:"secret" hcl:"secret"`
}

type SiteWlanAirwatchValue struct {
	ApiKey     *string `cty:"api_key" hcl:"api_key"`
	ConsoleUrl *string `cty:"console_url" hcl:"console_url"`
	Enabled    *bool   `cty:"enabled" hcl:"enabled"`
	Password   *string `cty:"password" hcl:"password"`
	Username   *string `cty:"username" hcl:"username"`
}

type SiteWlanAppLimitValue struct {
	Apps     map[string]int64 `cty:"apps" hcl:"apps"`
	Enabled  *bool            `cty:"enabled" hcl:"enabled"`
	WxtagIds map[string]int64 `cty:"wxtag_ids" hcl:"wxtag_ids"`
}

type SiteWlanAppQosValue struct {
	Apps    map[string]SiteWlanAppsValue `cty:"apps" hcl:"apps"`
	Enabled *bool                        `cty:"enabled" hcl:"enabled"`
	Others  []SiteWlanOthersValue        `cty:"others" hcl:"others"`
}

type SiteWlanAppsValue struct {
	Dscp      *string `cty:"dscp" hcl:"dscp"`
	DstSubnet *string `cty:"dst_subnet" hcl:"dst_subnet"`
	SrcSubnet *string `cty:"src_subnet" hcl:"src_subnet"`
}

type SiteWlanOthersValue struct {
	Dscp       *string `cty:"dscp" hcl:"dscp"`
	DstSubnet  *string `cty:"dst_subnet" hcl:"dst_subnet"`
	PortRanges *string `cty:"port_ranges" hcl:"port_ranges"`
	Protocol   *string `cty:"protocol" hcl:"protocol"`
	SrcSubnet  *string `cty:"src_subnet" hcl:"src_subnet"`
}

type SiteWlanAuthValue struct {
	AnticlogThreshold  *int64   `cty:"anticlog_threshold" hcl:"anticlog_threshold"`
	EapReauth          *bool    `cty:"eap_reauth" hcl:"eap_reauth"`
	EnableMacAuth      *bool    `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	KeyIdx             *int64   `cty:"key_idx" hcl:"key_idx"`
	Keys               []string `cty:"keys" hcl:"keys"`
	MultiPskOnly       *bool    `cty:"multi_psk_only" hcl:"multi_psk_only"`
	Owe                *string  `cty:"owe" hcl:"owe"`
	Pairwise           []string `cty:"pairwise" hcl:"pairwise"`
	PrivateWlan        *bool    `cty:"private_wlan" hcl:"private_wlan"`
	Psk                *string  `cty:"psk" hcl:"psk"`
	AuthType           *string  `cty:"type" hcl:"type"`
	WepAsSecondaryAuth *bool    `cty:"wep_as_secondary_auth" hcl:"wep_as_secondary_auth"`
}

type SiteWlanAuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type SiteWlanBonjourValue struct {
	AdditionalVlanIds []string                         `cty:"additional_vlan_ids" hcl:"additional_vlan_ids"`
	Enabled           *bool                            `cty:"enabled" hcl:"enabled"`
	Services          map[string]SiteWlanServicesValue `cty:"services" hcl:"services"`
}

type SiteWlanServicesValue struct {
	DisableLocal *bool    `cty:"disable_local" hcl:"disable_local"`
	RadiusGroups []string `cty:"radius_groups" hcl:"radius_groups"`
	Scope        *string  `cty:"scope" hcl:"scope"`
}

type SiteWlanCiscoCwaValue struct {
	AllowedHostnames []string `cty:"allowed_hostnames" hcl:"allowed_hostnames"`
	AllowedSubnets   []string `cty:"allowed_subnets" hcl:"allowed_subnets"`
	BlockedSubnets   []string `cty:"blocked_subnets" hcl:"blocked_subnets"`
	Enabled          *bool    `cty:"enabled" hcl:"enabled"`
}

type SiteWlanCoaServersValue struct {
	DisableEventTimestampCheck *bool   `cty:"disable_event_timestamp_check" hcl:"disable_event_timestamp_check"`
	Enabled                    *bool   `cty:"enabled" hcl:"enabled"`
	Ip                         string  `cty:"ip" hcl:"ip"`
	Port                       *string `cty:"port" hcl:"port"`
	Secret                     string  `cty:"secret" hcl:"secret"`
}

type SiteWlanDnsServerRewriteValue struct {
	Enabled      *bool             `cty:"enabled" hcl:"enabled"`
	RadiusGroups map[string]string `cty:"radius_groups" hcl:"radius_groups"`
}

type SiteWlanDynamicPskValue struct {
	DefaultPsk    *string `cty:"default_psk" hcl:"default_psk"`
	DefaultVlanId *string `cty:"default_vlan_id" hcl:"default_vlan_id"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	ForceLookup   *bool   `cty:"force_lookup" hcl:"force_lookup"`
	Source        *string `cty:"source" hcl:"source"`
}

type SiteWlanDynamicVlanValue struct {
	DefaultVlanIds  []string          `cty:"default_vlan_ids" hcl:"default_vlan_ids"`
	Enabled         *bool             `cty:"enabled" hcl:"enabled"`
	LocalVlanIds    []string          `cty:"local_vlan_ids" hcl:"local_vlan_ids"`
	DynamicVlanType *string           `cty:"type" hcl:"type"`
	Vlans           map[string]string `cty:"vlans" hcl:"vlans"`
}

type SiteWlanHotspot20Value struct {
	DomainName []string `cty:"domain_name" hcl:"domain_name"`
	Enabled    *bool    `cty:"enabled" hcl:"enabled"`
	NaiRealms  []string `cty:"nai_realms" hcl:"nai_realms"`
	Operators  []string `cty:"operators" hcl:"operators"`
	Rcoi       []string `cty:"rcoi" hcl:"rcoi"`
	VenueName  *string  `cty:"venue_name" hcl:"venue_name"`
}

type SiteWlanInjectDhcpOption82Value struct {
	CircuitId *string `cty:"circuit_id" hcl:"circuit_id"`
	Enabled   *bool   `cty:"enabled" hcl:"enabled"`
}

type SiteWlanMistNacValue struct {
	AcctInterimInterval *int64  `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AuthServersRetries  *int64  `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64  `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool   `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *int64  `cty:"coa_port" hcl:"coa_port"`
	Enabled             *bool   `cty:"enabled" hcl:"enabled"`
	FastDot1xTimers     *bool   `cty:"fast_dot1x_timers" hcl:"fast_dot1x_timers"`
	Network             *string `cty:"network" hcl:"network"`
	SourceIp            *string `cty:"source_ip" hcl:"source_ip"`
}

type SiteWlanPortalValue struct {
	AllowWlanIdRoam             *bool             `cty:"allow_wlan_id_roam" hcl:"allow_wlan_id_roam"`
	AmazonClientId              *string           `cty:"amazon_client_id" hcl:"amazon_client_id"`
	AmazonClientSecret          *string           `cty:"amazon_client_secret" hcl:"amazon_client_secret"`
	AmazonEmailDomains          []string          `cty:"amazon_email_domains" hcl:"amazon_email_domains"`
	AmazonEnabled               *bool             `cty:"amazon_enabled" hcl:"amazon_enabled"`
	AmazonExpire                *int64            `cty:"amazon_expire" hcl:"amazon_expire"`
	Auth                        *string           `cty:"auth" hcl:"auth"`
	AzureClientId               *string           `cty:"azure_client_id" hcl:"azure_client_id"`
	AzureClientSecret           *string           `cty:"azure_client_secret" hcl:"azure_client_secret"`
	AzureEnabled                *bool             `cty:"azure_enabled" hcl:"azure_enabled"`
	AzureExpire                 *int64            `cty:"azure_expire" hcl:"azure_expire"`
	AzureTenantId               *string           `cty:"azure_tenant_id" hcl:"azure_tenant_id"`
	BroadnetPassword            *string           `cty:"broadnet_password" hcl:"broadnet_password"`
	BroadnetSid                 *string           `cty:"broadnet_sid" hcl:"broadnet_sid"`
	BroadnetUserId              *string           `cty:"broadnet_user_id" hcl:"broadnet_user_id"`
	BypassWhenCloudDown         *bool             `cty:"bypass_when_cloud_down" hcl:"bypass_when_cloud_down"`
	ClickatellApiKey            *string           `cty:"clickatell_api_key" hcl:"clickatell_api_key"`
	CrossSite                   *bool             `cty:"cross_site" hcl:"cross_site"`
	EmailEnabled                *bool             `cty:"email_enabled" hcl:"email_enabled"`
	Enabled                     *bool             `cty:"enabled" hcl:"enabled"`
	Expire                      *int64            `cty:"expire" hcl:"expire"`
	ExternalPortalUrl           *string           `cty:"external_portal_url" hcl:"external_portal_url"`
	FacebookClientId            *string           `cty:"facebook_client_id" hcl:"facebook_client_id"`
	FacebookClientSecret        *string           `cty:"facebook_client_secret" hcl:"facebook_client_secret"`
	FacebookEmailDomains        []string          `cty:"facebook_email_domains" hcl:"facebook_email_domains"`
	FacebookEnabled             *bool             `cty:"facebook_enabled" hcl:"facebook_enabled"`
	FacebookExpire              *int64            `cty:"facebook_expire" hcl:"facebook_expire"`
	Forward                     *bool             `cty:"forward" hcl:"forward"`
	ForwardUrl                  *string           `cty:"forward_url" hcl:"forward_url"`
	GoogleClientId              *string           `cty:"google_client_id" hcl:"google_client_id"`
	GoogleClientSecret          *string           `cty:"google_client_secret" hcl:"google_client_secret"`
	GoogleEmailDomains          []string          `cty:"google_email_domains" hcl:"google_email_domains"`
	GoogleEnabled               *bool             `cty:"google_enabled" hcl:"google_enabled"`
	GoogleExpire                *int64            `cty:"google_expire" hcl:"google_expire"`
	GupshupPassword             *string           `cty:"gupshup_password" hcl:"gupshup_password"`
	GupshupUserid               *string           `cty:"gupshup_userid" hcl:"gupshup_userid"`
	MicrosoftClientId           *string           `cty:"microsoft_client_id" hcl:"microsoft_client_id"`
	MicrosoftClientSecret       *string           `cty:"microsoft_client_secret" hcl:"microsoft_client_secret"`
	MicrosoftEmailDomains       []string          `cty:"microsoft_email_domains" hcl:"microsoft_email_domains"`
	MicrosoftEnabled            *bool             `cty:"microsoft_enabled" hcl:"microsoft_enabled"`
	MicrosoftExpire             *int64            `cty:"microsoft_expire" hcl:"microsoft_expire"`
	PassphraseEnabled           *bool             `cty:"passphrase_enabled" hcl:"passphrase_enabled"`
	PassphraseExpire            *int64            `cty:"passphrase_expire" hcl:"passphrase_expire"`
	Password                    *string           `cty:"password" hcl:"password"`
	PredefinedSponsorsEnabled   *bool             `cty:"predefined_sponsors_enabled" hcl:"predefined_sponsors_enabled"`
	PredefinedSponsorsHideEmail *bool             `cty:"predefined_sponsors_hide_email" hcl:"predefined_sponsors_hide_email"`
	Privacy                     *bool             `cty:"privacy" hcl:"privacy"`
	PuzzelPassword              *string           `cty:"puzzel_password" hcl:"puzzel_password"`
	PuzzelServiceId             *string           `cty:"puzzel_service_id" hcl:"puzzel_service_id"`
	PuzzelUsername              *string           `cty:"puzzel_username" hcl:"puzzel_username"`
	SmsEnabled                  *bool             `cty:"sms_enabled" hcl:"sms_enabled"`
	SmsExpire                   *int64            `cty:"sms_expire" hcl:"sms_expire"`
	SmsMessageFormat            *string           `cty:"sms_message_format" hcl:"sms_message_format"`
	SmsProvider                 *string           `cty:"sms_provider" hcl:"sms_provider"`
	SmsglobalApiKey             *string           `cty:"smsglobal_api_key" hcl:"smsglobal_api_key"`
	SmsglobalApiSecret          *string           `cty:"smsglobal_api_secret" hcl:"smsglobal_api_secret"`
	SponsorAutoApprove          *bool             `cty:"sponsor_auto_approve" hcl:"sponsor_auto_approve"`
	SponsorEmailDomains         []string          `cty:"sponsor_email_domains" hcl:"sponsor_email_domains"`
	SponsorEnabled              *bool             `cty:"sponsor_enabled" hcl:"sponsor_enabled"`
	SponsorExpire               *int64            `cty:"sponsor_expire" hcl:"sponsor_expire"`
	SponsorLinkValidityDuration *string           `cty:"sponsor_link_validity_duration" hcl:"sponsor_link_validity_duration"`
	SponsorNotifyAll            *bool             `cty:"sponsor_notify_all" hcl:"sponsor_notify_all"`
	SponsorStatusNotify         *bool             `cty:"sponsor_status_notify" hcl:"sponsor_status_notify"`
	Sponsors                    map[string]string `cty:"sponsors" hcl:"sponsors"`
	SsoDefaultRole              *string           `cty:"sso_default_role" hcl:"sso_default_role"`
	SsoForcedRole               *string           `cty:"sso_forced_role" hcl:"sso_forced_role"`
	SsoIdpCert                  *string           `cty:"sso_idp_cert" hcl:"sso_idp_cert"`
	SsoIdpSignAlgo              *string           `cty:"sso_idp_sign_algo" hcl:"sso_idp_sign_algo"`
	SsoIdpSsoUrl                *string           `cty:"sso_idp_sso_url" hcl:"sso_idp_sso_url"`
	SsoIssuer                   *string           `cty:"sso_issuer" hcl:"sso_issuer"`
	SsoNameidFormat             *string           `cty:"sso_nameid_format" hcl:"sso_nameid_format"`
	TelstraClientId             *string           `cty:"telstra_client_id" hcl:"telstra_client_id"`
	TelstraClientSecret         *string           `cty:"telstra_client_secret" hcl:"telstra_client_secret"`
	TwilioAuthToken             *string           `cty:"twilio_auth_token" hcl:"twilio_auth_token"`
	TwilioPhoneNumber           *string           `cty:"twilio_phone_number" hcl:"twilio_phone_number"`
	TwilioSid                   *string           `cty:"twilio_sid" hcl:"twilio_sid"`
}

type SiteWlanQosValue struct {
	Class     *string `cty:"class" hcl:"class"`
	Overwrite *bool   `cty:"overwrite" hcl:"overwrite"`
}

type SiteWlanRadsecValue struct {
	CoaEnabled    *bool                  `cty:"coa_enabled" hcl:"coa_enabled"`
	Enabled       *bool                  `cty:"enabled" hcl:"enabled"`
	IdleTimeout   *string                `cty:"idle_timeout" hcl:"idle_timeout"`
	MxclusterIds  []string               `cty:"mxcluster_ids" hcl:"mxcluster_ids"`
	ProxyHosts    []string               `cty:"proxy_hosts" hcl:"proxy_hosts"`
	ServerName    *string                `cty:"server_name" hcl:"server_name"`
	Servers       []SiteWlanServersValue `cty:"servers" hcl:"servers"`
	UseMxedge     *bool                  `cty:"use_mxedge" hcl:"use_mxedge"`
	UseSiteMxedge *bool                  `cty:"use_site_mxedge" hcl:"use_site_mxedge"`
}

type SiteWlanServersValue struct {
	Host *string `cty:"host" hcl:"host"`
	Port *int64  `cty:"port" hcl:"port"`
}

type SiteWlanRatesetValue struct {
	Eht      *string  `cty:"eht" hcl:"eht"`
	He       *string  `cty:"he" hcl:"he"`
	Ht       *string  `cty:"ht" hcl:"ht"`
	Legacy   []string `cty:"legacy" hcl:"legacy"`
	MinRssi  *int64   `cty:"min_rssi" hcl:"min_rssi"`
	Template *string  `cty:"template" hcl:"template"`
	Vht      *string  `cty:"vht" hcl:"vht"`
}

type SiteWlanScheduleValue struct {
	Enabled *bool               `cty:"enabled" hcl:"enabled"`
	Hours   *SiteWlanHoursValue `cty:"hours" hcl:"hours"`
}

type SiteWlanHoursValue struct {
	Fri *string `cty:"fri" hcl:"fri"`
	Mon *string `cty:"mon" hcl:"mon"`
	Sat *string `cty:"sat" hcl:"sat"`
	Sun *string `cty:"sun" hcl:"sun"`
	Thu *string `cty:"thu" hcl:"thu"`
	Tue *string `cty:"tue" hcl:"tue"`
	Wed *string `cty:"wed" hcl:"wed"`
}
