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
	Disable11be                          *bool                    `hcl:"disable_11be"`
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
	ApiKey     *string `hcl:"api_key" cty:"api_key"`
	ConsoleUrl *string `hcl:"console_url" cty:"console_url"`
	Enabled    *bool   `hcl:"enabled" cty:"enabled"`
	Password   *string `hcl:"password" cty:"password"`
	Username   *string `hcl:"username" cty:"username"`
}

type AppLimitValue struct {
	Apps     map[string]int64 `hcl:"apps" cty:"apps"`
	Enabled  *bool            `hcl:"enabled" cty:"enabled"`
	WxtagIds map[string]int64 `hcl:"wxtag_ids" cty:"wxtag_ids"`
}

type AppQosValue struct {
	Apps    map[string]AppsValue `hcl:"apps" cty:"apps"`
	Enabled *bool                `hcl:"enabled" cty:"enabled"`
	Others  []OthersValue        `hcl:"others" cty:"others"`
}

type AppsValue struct {
	Dscp      *int64  `hcl:"dscp" cty:"dscp"`
	DstSubnet *string `hcl:"dst_subnet" cty:"dst_subnet"`
	SrcSubnet *string `hcl:"src_subnet" cty:"src_subnet"`
}

type OthersValue struct {
	Dscp       *int64  `hcl:"dscp" cty:"dscp"`
	DstSubnet  *string `hcl:"dst_subnet" cty:"dst_subnet"`
	PortRanges *string `hcl:"port_ranges" cty:"port_ranges"`
	Protocol   *string `hcl:"protocol" cty:"protocol"`
	SrcSubnet  *string `hcl:"src_subnet" cty:"src_subnet"`
}

type OrgWlanAuthValue struct {
	AnticlogThreshold  *int64   `hcl:"anticlog_threshold" cty:"anticlog_threshold"`
	EapReauth          *bool    `hcl:"eap_reauth" cty:"eap_reauth"`
	EnableMacAuth      *bool    `hcl:"enable_mac_auth" cty:"enable_mac_auth"`
	KeyIdx             *int64   `hcl:"key_idx" cty:"key_idx"`
	Keys               []string `hcl:"keys" cty:"keys"`
	MultiPskOnly       *bool    `hcl:"multi_psk_only" cty:"multi_psk_only"`
	Owe                *string  `hcl:"owe" cty:"owe"`
	Pairwise           []string `hcl:"pairwise" cty:"pairwise"`
	PrivateWlan        *bool    `hcl:"private_wlan" cty:"private_wlan"`
	Psk                *string  `hcl:"psk" cty:"psk"`
	AuthType           *string  `hcl:"type" cty:"type"`
	WepAsSecondaryAuth *bool    `hcl:"wep_as_secondary_auth" cty:"wep_as_secondary_auth"`
}

type BonjourValue struct {
	AdditionalVlanIds []string                 `hcl:"additional_vlan_ids" cty:"additional_vlan_ids"`
	Enabled           *bool                    `hcl:"enabled" cty:"enabled"`
	Services          map[string]ServicesValue `hcl:"services" cty:"services"`
}

type ServicesValue struct {
	DisableLocal *bool    `hcl:"disable_local" cty:"disable_local"`
	RadiusGroups []string `hcl:"radius_groups" cty:"radius_groups"`
	Scope        *string  `hcl:"scope" cty:"scope"`
}

type CiscoCwaValue struct {
	AllowedHostnames []string `hcl:"allowed_hostnames" cty:"allowed_hostnames"`
	AllowedSubnets   []string `hcl:"allowed_subnets" cty:"allowed_subnets"`
	BlockedSubnets   []string `hcl:"blocked_subnets" cty:"blocked_subnets"`
	Enabled          *bool    `hcl:"enabled" cty:"enabled"`
}

type CoaServersValue struct {
	DisableEventTimestampCheck *bool   `hcl:"disable_event_timestamp_check" cty:"disable_event_timestamp_check"`
	Enabled                    *bool   `hcl:"enabled" cty:"enabled"`
	Ip                         string  `hcl:"ip" cty:"ip"`
	Port                       *string `hcl:"port" cty:"port"`
	Secret                     string  `hcl:"secret" cty:"secret"`
}

type DnsServerRewriteValue struct {
	Enabled      *bool             `hcl:"enabled" cty:"enabled"`
	RadiusGroups map[string]string `hcl:"radius_groups" cty:"radius_groups"`
}

type DynamicPskValue struct {
	DefaultPsk    *string `hcl:"default_psk" cty:"default_psk"`
	DefaultVlanId *string `hcl:"default_vlan_id" cty:"default_vlan_id"`
	Enabled       *bool   `hcl:"enabled" cty:"enabled"`
	ForceLookup   *bool   `hcl:"force_lookup" cty:"force_lookup"`
	Source        *string `hcl:"source" cty:"source"`
}

type DynamicVlanValue struct {
	DefaultVlanIds  []string          `hcl:"default_vlan_ids" cty:"default_vlan_ids"`
	Enabled         *bool             `hcl:"enabled" cty:"enabled"`
	LocalVlanIds    []string          `hcl:"local_vlan_ids" cty:"local_vlan_ids"`
	DynamicVlanType *string           `hcl:"type" cty:"type"`
	Vlans           map[string]string `hcl:"vlans" cty:"vlans"`
}

type Hotspot20Value struct {
	DomainName []string `hcl:"domain_name" cty:"domain_name"`
	Enabled    *bool    `hcl:"enabled" cty:"enabled"`
	NaiRealms  []string `hcl:"nai_realms" cty:"nai_realms"`
	Operators  []string `hcl:"operators" cty:"operators"`
	Rcoi       []string `hcl:"rcoi" cty:"rcoi"`
	VenueName  *string  `hcl:"venue_name" cty:"venue_name"`
}

type InjectDhcpOption82Value struct {
	CircuitId *string `hcl:"circuit_id" cty:"circuit_id"`
	Enabled   *bool   `hcl:"enabled" cty:"enabled"`
}

type OrgWlanMistNacValue struct {
	Enabled *bool `hcl:"enabled" cty:"enabled"`
}

type PortalValue struct {
	AllowWlanIdRoam             *bool             `hcl:"allow_wlan_id_roam" cty:"allow_wlan_id_roam"`
	AmazonClientId              *string           `hcl:"amazon_client_id" cty:"amazon_client_id"`
	AmazonClientSecret          *string           `hcl:"amazon_client_secret" cty:"amazon_client_secret"`
	AmazonEmailDomains          []string          `hcl:"amazon_email_domains" cty:"amazon_email_domains"`
	AmazonEnabled               *bool             `hcl:"amazon_enabled" cty:"amazon_enabled"`
	AmazonExpire                *int64            `hcl:"amazon_expire" cty:"amazon_expire"`
	Auth                        *string           `hcl:"auth" cty:"auth"`
	AzureClientId               *string           `hcl:"azure_client_id" cty:"azure_client_id"`
	AzureClientSecret           *string           `hcl:"azure_client_secret" cty:"azure_client_secret"`
	AzureEnabled                *bool             `hcl:"azure_enabled" cty:"azure_enabled"`
	AzureExpire                 *int64            `hcl:"azure_expire" cty:"azure_expire"`
	AzureTenantId               *string           `hcl:"azure_tenant_id" cty:"azure_tenant_id"`
	BroadnetPassword            *string           `hcl:"broadnet_password" cty:"broadnet_password"`
	BroadnetSid                 *string           `hcl:"broadnet_sid" cty:"broadnet_sid"`
	BroadnetUserId              *string           `hcl:"broadnet_user_id" cty:"broadnet_user_id"`
	BypassWhenCloudDown         *bool             `hcl:"bypass_when_cloud_down" cty:"bypass_when_cloud_down"`
	ClickatellApiKey            *string           `hcl:"clickatell_api_key" cty:"clickatell_api_key"`
	CrossSite                   *bool             `hcl:"cross_site" cty:"cross_site"`
	EmailEnabled                *bool             `hcl:"email_enabled" cty:"email_enabled"`
	Enabled                     *bool             `hcl:"enabled" cty:"enabled"`
	Expire                      *int64            `hcl:"expire" cty:"expire"`
	ExternalPortalUrl           *string           `hcl:"external_portal_url" cty:"external_portal_url"`
	FacebookClientId            *string           `hcl:"facebook_client_id" cty:"facebook_client_id"`
	FacebookClientSecret        *string           `hcl:"facebook_client_secret" cty:"facebook_client_secret"`
	FacebookEmailDomains        []string          `hcl:"facebook_email_domains" cty:"facebook_email_domains"`
	FacebookEnabled             *bool             `hcl:"facebook_enabled" cty:"facebook_enabled"`
	FacebookExpire              *int64            `hcl:"facebook_expire" cty:"facebook_expire"`
	Forward                     *bool             `hcl:"forward" cty:"forward"`
	ForwardUrl                  *string           `hcl:"forward_url" cty:"forward_url"`
	GoogleClientId              *string           `hcl:"google_client_id" cty:"google_client_id"`
	GoogleClientSecret          *string           `hcl:"google_client_secret" cty:"google_client_secret"`
	GoogleEmailDomains          []string          `hcl:"google_email_domains" cty:"google_email_domains"`
	GoogleEnabled               *bool             `hcl:"google_enabled" cty:"google_enabled"`
	GoogleExpire                *int64            `hcl:"google_expire" cty:"google_expire"`
	GupshupPassword             *string           `hcl:"gupshup_password" cty:"gupshup_password"`
	GupshupUserid               *string           `hcl:"gupshup_userid" cty:"gupshup_userid"`
	MicrosoftClientId           *string           `hcl:"microsoft_client_id" cty:"microsoft_client_id"`
	MicrosoftClientSecret       *string           `hcl:"microsoft_client_secret" cty:"microsoft_client_secret"`
	MicrosoftEmailDomains       []string          `hcl:"microsoft_email_domains" cty:"microsoft_email_domains"`
	MicrosoftEnabled            *bool             `hcl:"microsoft_enabled" cty:"microsoft_enabled"`
	MicrosoftExpire             *int64            `hcl:"microsoft_expire" cty:"microsoft_expire"`
	PassphraseEnabled           *bool             `hcl:"passphrase_enabled" cty:"passphrase_enabled"`
	PassphraseExpire            *int64            `hcl:"passphrase_expire" cty:"passphrase_expire"`
	Password                    *string           `hcl:"password" cty:"password"`
	PredefinedSponsorsEnabled   *bool             `hcl:"predefined_sponsors_enabled" cty:"predefined_sponsors_enabled"`
	PredefinedSponsorsHideEmail *bool             `hcl:"predefined_sponsors_hide_email" cty:"predefined_sponsors_hide_email"`
	Privacy                     *bool             `hcl:"privacy" cty:"privacy"`
	PuzzelPassword              *string           `hcl:"puzzel_password" cty:"puzzel_password"`
	PuzzelServiceId             *string           `hcl:"puzzel_service_id" cty:"puzzel_service_id"`
	PuzzelUsername              *string           `hcl:"puzzel_username" cty:"puzzel_username"`
	SmsEnabled                  *bool             `hcl:"sms_enabled" cty:"sms_enabled"`
	SmsExpire                   *int64            `hcl:"sms_expire" cty:"sms_expire"`
	SmsMessageFormat            *string           `hcl:"sms_message_format" cty:"sms_message_format"`
	SmsProvider                 *string           `hcl:"sms_provider" cty:"sms_provider"`
	SponsorAutoApprove          *bool             `hcl:"sponsor_auto_approve" cty:"sponsor_auto_approve"`
	SponsorEmailDomains         []string          `hcl:"sponsor_email_domains" cty:"sponsor_email_domains"`
	SponsorEnabled              *bool             `hcl:"sponsor_enabled" cty:"sponsor_enabled"`
	SponsorExpire               *int64            `hcl:"sponsor_expire" cty:"sponsor_expire"`
	SponsorLinkValidityDuration *string           `hcl:"sponsor_link_validity_duration" cty:"sponsor_link_validity_duration"`
	SponsorNotifyAll            *bool             `hcl:"sponsor_notify_all" cty:"sponsor_notify_all"`
	SponsorStatusNotify         *bool             `hcl:"sponsor_status_notify" cty:"sponsor_status_notify"`
	Sponsors                    map[string]string `hcl:"sponsors" cty:"sponsors"`
	SsoDefaultRole              *string           `hcl:"sso_default_role" cty:"sso_default_role"`
	SsoForcedRole               *string           `hcl:"sso_forced_role" cty:"sso_forced_role"`
	SsoIdpCert                  *string           `hcl:"sso_idp_cert" cty:"sso_idp_cert"`
	SsoIdpSignAlgo              *string           `hcl:"sso_idp_sign_algo" cty:"sso_idp_sign_algo"`
	SsoIdpSsoUrl                *string           `hcl:"sso_idp_sso_url" cty:"sso_idp_sso_url"`
	SsoIssuer                   *string           `hcl:"sso_issuer" cty:"sso_issuer"`
	SsoNameidFormat             *string           `hcl:"sso_nameid_format" cty:"sso_nameid_format"`
	TelstraClientId             *string           `hcl:"telstra_client_id" cty:"telstra_client_id"`
	TelstraClientSecret         *string           `hcl:"telstra_client_secret" cty:"telstra_client_secret"`
	TwilioAuthToken             *string           `hcl:"twilio_auth_token" cty:"twilio_auth_token"`
	TwilioPhoneNumber           *string           `hcl:"twilio_phone_number" cty:"twilio_phone_number"`
	TwilioSid                   *string           `hcl:"twilio_sid" cty:"twilio_sid"`
}

type QosValue struct {
	Class     *string `hcl:"class" cty:"class"`
	Overwrite *bool   `hcl:"overwrite" cty:"overwrite"`
}

type RadsecValue struct {
	CoaEnabled    *bool                 `hcl:"coa_enabled" cty:"coa_enabled"`
	Enabled       *bool                 `hcl:"enabled" cty:"enabled"`
	IdleTimeout   *string               `hcl:"idle_timeout" cty:"idle_timeout"`
	MxclusterIds  []string              `hcl:"mxcluster_ids" cty:"mxcluster_ids"`
	ProxyHosts    []string              `hcl:"proxy_hosts" cty:"proxy_hosts"`
	ServerName    *string               `hcl:"server_name" cty:"server_name"`
	Servers       []OrgWlanServersValue `hcl:"servers" cty:"servers"`
	UseMxedge     *bool                 `hcl:"use_mxedge" cty:"use_mxedge"`
	UseSiteMxedge *bool                 `hcl:"use_site_mxedge" cty:"use_site_mxedge"`
}

type OrgWlanServersValue struct {
	Host *string `hcl:"host" cty:"host"`
	Port *int64  `hcl:"port" cty:"port"`
}

type RatesetValue struct {
	Eht      *string  `hcl:"eht" cty:"eht"`
	He       *string  `hcl:"he" cty:"he"`
	Ht       *string  `hcl:"ht" cty:"ht"`
	Legacy   []string `hcl:"legacy" cty:"legacy"`
	MinRssi  *int64   `hcl:"min_rssi" cty:"min_rssi"`
	Template *string  `hcl:"template" cty:"template"`
	Vht      *string  `hcl:"vht" cty:"vht"`
}

type ScheduleValue struct {
	Enabled *bool       `hcl:"enabled" cty:"enabled"`
	Hours   *HoursValue `hcl:"hours" cty:"hours"`
}

type HoursValue struct {
	Fri *string `hcl:"fri" cty:"fri"`
	Mon *string `hcl:"mon" cty:"mon"`
	Sat *string `hcl:"sat" cty:"sat"`
	Sun *string `hcl:"sun" cty:"sun"`
	Thu *string `hcl:"thu" cty:"thu"`
	Tue *string `hcl:"tue" cty:"tue"`
	Wed *string `hcl:"wed" cty:"wed"`
}
