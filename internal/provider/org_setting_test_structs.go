package provider

type OrgSettingModel struct {
	ApUpdownThreshold      *int64                          `hcl:"ap_updown_threshold"`
	ApiPolicy              *ApiPolicyValue                 `hcl:"api_policy"`
	Cacerts                []string                        `hcl:"cacerts"`
	Celona                 *CelonaValue                    `hcl:"celona"`
	Cloudshark             *CloudsharkValue                `hcl:"cloudshark"`
	DeviceCert             *DeviceCertValue                `hcl:"device_cert"`
	DeviceUpdownThreshold  *int64                          `hcl:"device_updown_threshold"`
	DisablePcap            *bool                           `hcl:"disable_pcap"`
	DisableRemoteShell     *bool                           `hcl:"disable_remote_shell"`
	GatewayUpdownThreshold *int64                          `hcl:"gateway_updown_threshold"`
	Installer              *InstallerValue                 `hcl:"installer"`
	Jcloud                 *JcloudValue                    `hcl:"jcloud"`
	JcloudRa               *JcloudRaValue                  `hcl:"jcloud_ra"`
	Mgmt                   *MgmtValue                      `hcl:"mgmt"`
	MistNac                *OrgSettingMistNacValue         `hcl:"mist_nac"`
	MxedgeFipsEnabled      *bool                           `hcl:"mxedge_fips_enabled"`
	MxedgeMgmt             *MxedgeMgmtValue                `hcl:"mxedge_mgmt"`
	OpticPortConfig        map[string]OpticPortConfigValue `hcl:"optic_port_config"`
	OrgId                  string                          `hcl:"org_id"`
	PasswordPolicy         *PasswordPolicyValue            `hcl:"password_policy"`
	Pcap                   *PcapValue                      `hcl:"pcap"`
	Security               *SecurityValue                  `hcl:"security"`
	SwitchMgmt             *OrgSettingSwitchMgmtValue      `hcl:"switch_mgmt"`
	SwitchUpdownThreshold  *int64                          `hcl:"switch_updown_threshold"`
	SyntheticTest          *SyntheticTestValue             `hcl:"synthetic_test"`
	UiIdleTimeout          *int64                          `hcl:"ui_idle_timeout"`
	VpnOptions             *VpnOptionsValue                `hcl:"vpn_options"`
	WanPma                 *WanPmaValue                    `hcl:"wan_pma"`
	WiredPma               *WiredPmaValue                  `hcl:"wired_pma"`
	WirelessPma            *WirelessPmaValue               `hcl:"wireless_pma"`
}

type ApiPolicyValue struct {
	NoReveal *bool `cty:"no_reveal" hcl:"no_reveal"`
}

type CelonaValue struct {
	ApiKey    string `cty:"api_key" hcl:"api_key"`
	ApiPrefix string `cty:"api_prefix" hcl:"api_prefix"`
}

type CloudsharkValue struct {
	Apitoken string  `cty:"apitoken" hcl:"apitoken"`
	Url      *string `cty:"url" hcl:"url"`
}

type CradlepointValue struct {
}

type DeviceCertValue struct {
	Cert string `cty:"cert" hcl:"cert"`
	Key  string `cty:"key" hcl:"key"`
}

type InstallerValue struct {
	AllowAllDevices *bool    `cty:"allow_all_devices" hcl:"allow_all_devices"`
	AllowAllSites   *bool    `cty:"allow_all_sites" hcl:"allow_all_sites"`
	ExtraSiteIds    []string `cty:"extra_site_ids" hcl:"extra_site_ids"`
	GracePeriod     *int64   `cty:"grace_period" hcl:"grace_period"`
}

type JcloudValue struct {
	OrgApitoken     string `cty:"org_apitoken" hcl:"org_apitoken"`
	OrgApitokenName string `cty:"org_apitoken_name" hcl:"org_apitoken_name"`
	OrgId           string `cty:"org_id" hcl:"org_id"`
}

type JcloudRaValue struct {
	OrgApitoken     *string `cty:"org_apitoken" hcl:"org_apitoken"`
	OrgApitokenName *string `cty:"org_apitoken_name" hcl:"org_apitoken_name"`
	OrgId           *string `cty:"org_id" hcl:"org_id"`
}

type JuniperValue struct {
}

type AccountsValue struct {
}

type MgmtValue struct {
	MxtunnelIds []string `cty:"mxtunnel_ids" hcl:"mxtunnel_ids"`
	UseMxtunnel *bool    `cty:"use_mxtunnel" hcl:"use_mxtunnel"`
	UseWxtunnel *bool    `cty:"use_wxtunnel" hcl:"use_wxtunnel"`
}

type OrgSettingMistNacValue struct {
	Cacerts                   []string         `cty:"cacerts" hcl:"cacerts"`
	DefaultIdpId              *string          `cty:"default_idp_id" hcl:"default_idp_id"`
	DisableRsaeAlgorithms     *bool            `cty:"disable_rsae_algorithms" hcl:"disable_rsae_algorithms"`
	EapSslSecurityLevel       *int64           `cty:"eap_ssl_security_level" hcl:"eap_ssl_security_level"`
	EuOnly                    *bool            `cty:"eu_only" hcl:"eu_only"`
	IdpMachineCertLookupField *string          `cty:"idp_machine_cert_lookup_field" hcl:"idp_machine_cert_lookup_field"`
	IdpUserCertLookupField    *string          `cty:"idp_user_cert_lookup_field" hcl:"idp_user_cert_lookup_field"`
	Idps                      []IdpsValue      `cty:"idps" hcl:"idps"`
	ServerCert                *ServerCertValue `cty:"server_cert" hcl:"server_cert"`
	UseIpVersion              *string          `cty:"use_ip_version" hcl:"use_ip_version"`
	UseSslPort                *bool            `cty:"use_ssl_port" hcl:"use_ssl_port"`
}

type IdpsValue struct {
	ExcludeRealms []string `cty:"exclude_realms" hcl:"exclude_realms"`
	Id            string   `cty:"id" hcl:"id"`
	UserRealms    []string `cty:"user_realms" hcl:"user_realms"`
}

type ServerCertValue struct {
	Cert     *string `cty:"cert" hcl:"cert"`
	Key      *string `cty:"key" hcl:"key"`
	Password *string `cty:"password" hcl:"password"`
}

type MxedgeMgmtValue struct {
	FipsEnabled  *bool   `cty:"fips_enabled" hcl:"fips_enabled"`
	MistPassword *string `cty:"mist_password" hcl:"mist_password"`
	OobIpType    *string `cty:"oob_ip_type" hcl:"oob_ip_type"`
	OobIpType6   *string `cty:"oob_ip_type6" hcl:"oob_ip_type6"`
	RootPassword *string `cty:"root_password" hcl:"root_password"`
}

type OpticPortConfigValue struct {
	Channelized *bool   `cty:"channelized" hcl:"channelized"`
	Speed       *string `cty:"speed" hcl:"speed"`
}

type PasswordPolicyValue struct {
	Enabled               *bool  `cty:"enabled" hcl:"enabled"`
	ExpiryInDays          *int64 `cty:"expiry_in_days" hcl:"expiry_in_days"`
	MinLength             *int64 `cty:"min_length" hcl:"min_length"`
	RequiresSpecialChar   *bool  `cty:"requires_special_char" hcl:"requires_special_char"`
	RequiresTwoFactorAuth *bool  `cty:"requires_two_factor_auth" hcl:"requires_two_factor_auth"`
}

type PcapValue struct {
	Bucket    *string `cty:"bucket" hcl:"bucket"`
	MaxPktLen *int64  `cty:"max_pkt_len" hcl:"max_pkt_len"`
}

type SecurityValue struct {
	DisableLocalSsh     *bool   `cty:"disable_local_ssh" hcl:"disable_local_ssh"`
	FipsZeroizePassword *string `cty:"fips_zeroize_password" hcl:"fips_zeroize_password"`
	LimitSshAccess      *bool   `cty:"limit_ssh_access" hcl:"limit_ssh_access"`
}

type OrgSettingSwitchMgmtValue struct {
	ApAffinityThreshold *int64 `cty:"ap_affinity_threshold" hcl:"ap_affinity_threshold"`
}

type SyntheticTestValue struct {
	Disabled     *bool              `cty:"disabled" hcl:"disabled"`
	Vlans        []VlansValue       `cty:"vlans" hcl:"vlans"`
	WanSpeedtest *WanSpeedtestValue `cty:"wan_speedtest" hcl:"wan_speedtest"`
}

type VlansValue struct {
	CustomTestUrls []string `cty:"custom_test_urls" hcl:"custom_test_urls"`
	Disabled       *bool    `cty:"disabled" hcl:"disabled"`
	VlanIds        []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type WanSpeedtestValue struct {
	Enabled   *bool   `cty:"enabled" hcl:"enabled"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type VpnOptionsValue struct {
	AsBase   *int64  `cty:"as_base" hcl:"as_base"`
	StSubnet *string `cty:"st_subnet" hcl:"st_subnet"`
}

type WanPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type WiredPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type WirelessPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}
