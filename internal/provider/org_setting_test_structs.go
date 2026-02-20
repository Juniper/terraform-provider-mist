package provider

type OrgSettingModel struct {
	ApUpdownThreshold      *int64                                    `hcl:"ap_updown_threshold"`
	ApiPolicy              *OrgSettingApiPolicyValue                 `hcl:"api_policy"`
	Cacerts                []string                                  `hcl:"cacerts"`
	Celona                 *OrgSettingCelonaValue                    `hcl:"celona"`
	Cloudshark             *OrgSettingCloudsharkValue                `hcl:"cloudshark"`
	DeviceCert             *OrgSettingDeviceCertValue                `hcl:"device_cert"`
	DeviceUpdownThreshold  *int64                                    `hcl:"device_updown_threshold"`
	DisablePcap            *bool                                     `hcl:"disable_pcap"`
	DisableRemoteShell     *bool                                     `hcl:"disable_remote_shell"`
	GatewayUpdownThreshold *int64                                    `hcl:"gateway_updown_threshold"`
	Installer              *OrgSettingInstallerValue                 `hcl:"installer"`
	Jcloud                 *OrgSettingJcloudValue                    `hcl:"jcloud"`
	JcloudRa               *OrgSettingJcloudRaValue                  `hcl:"jcloud_ra"`
	JuniperSrx             *OrgSettingJuniperSrxValue                `hcl:"juniper_srx"`
	JunosShellAccess       *OrgSettingJunosShellAccessValue          `hcl:"junos_shell_access"`
	Marvis                 *OrgSettingMarvisValue                    `hcl:"marvis"`
	Mgmt                   *OrgSettingMgmtValue                      `hcl:"mgmt"`
	MistNac                *OrgSettingMistNacValue                   `hcl:"mist_nac"`
	MxedgeMgmt             *OrgSettingMxedgeMgmtValue                `hcl:"mxedge_mgmt"`
	OpticPortConfig        map[string]OrgSettingOpticPortConfigValue `hcl:"optic_port_config"`
	OrgId                  string                                    `hcl:"org_id"`
	PasswordPolicy         *OrgSettingPasswordPolicyValue            `hcl:"password_policy"`
	Security               *OrgSettingSecurityValue                  `hcl:"security"`
	Ssr                    *OrgSettingSsrValue                       `hcl:"ssr"`
	Switch                 *OrgSettingSwitchValue                    `hcl:"switch"`
	SwitchMgmt             *OrgSettingSwitchMgmtValue                `hcl:"switch_mgmt"`
	SwitchUpdownThreshold  *int64                                    `hcl:"switch_updown_threshold"`
	SyntheticTest          *OrgSettingSyntheticTestValue             `hcl:"synthetic_test"`
	UiIdleTimeout          *int64                                    `hcl:"ui_idle_timeout"`
	UiNoTracking           *bool                                     `hcl:"ui_no_tracking"`
	VpnOptions             *OrgSettingVpnOptionsValue                `hcl:"vpn_options"`
	WanPma                 *OrgSettingWanPmaValue                    `hcl:"wan_pma"`
	WiredPma               *OrgSettingWiredPmaValue                  `hcl:"wired_pma"`
	WirelessPma            *OrgSettingWirelessPmaValue               `hcl:"wireless_pma"`
}

type OrgSettingApiPolicyValue struct {
	NoReveal *bool `cty:"no_reveal" hcl:"no_reveal"`
}

type OrgSettingCelonaValue struct {
	ApiKey    string `cty:"api_key" hcl:"api_key"`
	ApiPrefix string `cty:"api_prefix" hcl:"api_prefix"`
}

type OrgSettingCloudsharkValue struct {
	Apitoken *string `cty:"apitoken" hcl:"apitoken"`
	Url      *string `cty:"url" hcl:"url"`
}

type OrgSettingCradlepointValue struct {
}

type OrgSettingDeviceCertValue struct {
	Cert string `cty:"cert" hcl:"cert"`
	Key  string `cty:"key" hcl:"key"`
}

type OrgSettingInstallerValue struct {
	AllowAllDevices *bool    `cty:"allow_all_devices" hcl:"allow_all_devices"`
	AllowAllSites   *bool    `cty:"allow_all_sites" hcl:"allow_all_sites"`
	ExtraSiteIds    []string `cty:"extra_site_ids" hcl:"extra_site_ids"`
	GracePeriod     *int64   `cty:"grace_period" hcl:"grace_period"`
}

type OrgSettingJcloudValue struct {
	OrgApitoken     string `cty:"org_apitoken" hcl:"org_apitoken"`
	OrgApitokenName string `cty:"org_apitoken_name" hcl:"org_apitoken_name"`
	OrgId           string `cty:"org_id" hcl:"org_id"`
}

type OrgSettingJcloudRaValue struct {
	OrgApitoken     *string `cty:"org_apitoken" hcl:"org_apitoken"`
	OrgApitokenName *string `cty:"org_apitoken_name" hcl:"org_apitoken_name"`
	OrgId           *string `cty:"org_id" hcl:"org_id"`
}

type OrgSettingJuniperValue struct {
}

type OrgSettingAccountsValue struct {
}

type OrgSettingJuniperSrxValue struct {
	SrxAutoUpgrade *OrgSettingSrxAutoUpgradeValue `cty:"auto_upgrade" hcl:"auto_upgrade"`
}

type OrgSettingSrxAutoUpgradeValue struct {
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	Snapshot       *bool             `cty:"snapshot" hcl:"snapshot"`
	Version        *string           `cty:"version" hcl:"version"`
}

type OrgSettingJunosShellAccessValue struct {
	Admin    *string `cty:"admin" hcl:"admin"`
	Helpdesk *string `cty:"helpdesk" hcl:"helpdesk"`
	Read     *string `cty:"read" hcl:"read"`
	Write    *string `cty:"write" hcl:"write"`
}

type OrgSettingMarvisValue struct {
	AutoOperations *OrgSettingAutoOperationsValue `cty:"auto_operations" hcl:"auto_operations"`
}

type OrgSettingAutoOperationsValue struct {
	ApInsufficientCapacity                 *bool `cty:"ap_insufficient_capacity" hcl:"ap_insufficient_capacity"`
	ApLoop                                 *bool `cty:"ap_loop" hcl:"ap_loop"`
	ApNonCompliant                         *bool `cty:"ap_non_compliant" hcl:"ap_non_compliant"`
	BouncePortForAbnormalPoeClient         *bool `cty:"bounce_port_for_abnormal_poe_client" hcl:"bounce_port_for_abnormal_poe_client"`
	DisablePortWhenDdosProtocolViolation   *bool `cty:"disable_port_when_ddos_protocol_violation" hcl:"disable_port_when_ddos_protocol_violation"`
	DisablePortWhenRogueDhcpServerDetected *bool `cty:"disable_port_when_rogue_dhcp_server_detected" hcl:"disable_port_when_rogue_dhcp_server_detected"`
	GatewayNonCompliant                    *bool `cty:"gateway_non_compliant" hcl:"gateway_non_compliant"`
	SwitchMisconfiguredPort                *bool `cty:"switch_misconfigured_port" hcl:"switch_misconfigured_port"`
	SwitchPortStuck                        *bool `cty:"switch_port_stuck" hcl:"switch_port_stuck"`
}

type OrgSettingMgmtValue struct {
	MxtunnelIds []string `cty:"mxtunnel_ids" hcl:"mxtunnel_ids"`
	UseMxtunnel *bool    `cty:"use_mxtunnel" hcl:"use_mxtunnel"`
	UseWxtunnel *bool    `cty:"use_wxtunnel" hcl:"use_wxtunnel"`
}

type OrgSettingMistNacValue struct {
	Cacerts                   []string                       `cty:"cacerts" hcl:"cacerts"`
	DefaultIdpId              *string                        `cty:"default_idp_id" hcl:"default_idp_id"`
	DisableRsaeAlgorithms     *bool                          `cty:"disable_rsae_algorithms" hcl:"disable_rsae_algorithms"`
	EapSslSecurityLevel       *int64                         `cty:"eap_ssl_security_level" hcl:"eap_ssl_security_level"`
	EuOnly                    *bool                          `cty:"eu_only" hcl:"eu_only"`
	Fingerprinting            *OrgSettingFingerprintingValue `cty:"fingerprinting" hcl:"fingerprinting"`
	IdpMachineCertLookupField *string                        `cty:"idp_machine_cert_lookup_field" hcl:"idp_machine_cert_lookup_field"`
	IdpUserCertLookupField    *string                        `cty:"idp_user_cert_lookup_field" hcl:"idp_user_cert_lookup_field"`
	Idps                      []OrgSettingIdpsValue          `cty:"idps" hcl:"idps"`
	ServerCert                *OrgSettingServerCertValue     `cty:"server_cert" hcl:"server_cert"`
	UseIpVersion              *string                        `cty:"use_ip_version" hcl:"use_ip_version"`
	UseSslPort                *bool                          `cty:"use_ssl_port" hcl:"use_ssl_port"`
	UsermacExpiry             *int64                         `cty:"usermac_expiry" hcl:"usermac_expiry"`
}

type OrgSettingFingerprintingValue struct {
	Enabled             *bool   `cty:"enabled" hcl:"enabled"`
	GenerateCoa         *bool   `cty:"generate_coa" hcl:"generate_coa"`
	GenerateWirelessCoa *bool   `cty:"generate_wireless_coa" hcl:"generate_wireless_coa"`
	WirelessCoaType     *string `cty:"wireless_coa_type" hcl:"wireless_coa_type"`
}

type OrgSettingIdpsValue struct {
	ExcludeRealms []string `cty:"exclude_realms" hcl:"exclude_realms"`
	Id            string   `cty:"id" hcl:"id"`
	UserRealms    []string `cty:"user_realms" hcl:"user_realms"`
}

type OrgSettingServerCertValue struct {
	Cert     *string `cty:"cert" hcl:"cert"`
	Key      *string `cty:"key" hcl:"key"`
	Password *string `cty:"password" hcl:"password"`
}

type OrgSettingMxedgeMgmtValue struct {
	ConfigAutoRevert *bool   `cty:"config_auto_revert" hcl:"config_auto_revert"`
	FipsEnabled      *bool   `cty:"fips_enabled" hcl:"fips_enabled"`
	MistPassword     *string `cty:"mist_password" hcl:"mist_password"`
	OobIpType        *string `cty:"oob_ip_type" hcl:"oob_ip_type"`
	OobIpType6       *string `cty:"oob_ip_type6" hcl:"oob_ip_type6"`
	RootPassword     *string `cty:"root_password" hcl:"root_password"`
}

type OrgSettingOpticPortConfigValue struct {
	Channelized *bool   `cty:"channelized" hcl:"channelized"`
	Speed       *string `cty:"speed" hcl:"speed"`
}

type OrgSettingPasswordPolicyValue struct {
	Enabled               *bool  `cty:"enabled" hcl:"enabled"`
	ExpiryInDays          *int64 `cty:"expiry_in_days" hcl:"expiry_in_days"`
	MinLength             *int64 `cty:"min_length" hcl:"min_length"`
	RequiresSpecialChar   *bool  `cty:"requires_special_char" hcl:"requires_special_char"`
	RequiresTwoFactorAuth *bool  `cty:"requires_two_factor_auth" hcl:"requires_two_factor_auth"`
}

type OrgSettingPcapValue struct {
}

type OrgSettingSecurityValue struct {
	DisableLocalSsh     *bool   `cty:"disable_local_ssh" hcl:"disable_local_ssh"`
	FipsZeroizePassword *string `cty:"fips_zeroize_password" hcl:"fips_zeroize_password"`
	LimitSshAccess      *bool   `cty:"limit_ssh_access" hcl:"limit_ssh_access"`
}

type OrgSettingSsrValue struct {
	ConductorHosts []string                       `cty:"conductor_hosts" hcl:"conductor_hosts"`
	ConductorToken *string                        `cty:"conductor_token" hcl:"conductor_token"`
	DisableStats   *bool                          `cty:"disable_stats" hcl:"disable_stats"`
	Proxy          *OrgSettingProxyValue          `cty:"proxy" hcl:"proxy"`
	SsrAutoUpgrade *OrgSettingSsrAutoUpgradeValue `cty:"auto_upgrade" hcl:"auto_upgrade"`
}

type OrgSettingProxyValue struct {
	Disabled *bool   `cty:"disabled" hcl:"disabled"`
	Url      *string `cty:"url" hcl:"url"`
}

type OrgSettingSsrAutoUpgradeValue struct {
	Channel        *string           `cty:"channel" hcl:"channel"`
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	Version        *string           `cty:"version" hcl:"version"`
}

type OrgSettingSwitchValue struct {
	AutoUpgrade *OrgSettingAutoUpgradeValue `cty:"auto_upgrade" hcl:"auto_upgrade"`
}

type OrgSettingAutoUpgradeValue struct {
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	Snapshot       *bool             `cty:"snapshot" hcl:"snapshot"`
}

type OrgSettingSwitchMgmtValue struct {
	ApAffinityThreshold *int64 `cty:"ap_affinity_threshold" hcl:"ap_affinity_threshold"`
}

type OrgSettingSyntheticTestValue struct {
	Aggressiveness *string                                `cty:"aggressiveness" hcl:"aggressiveness"`
	CustomProbes   map[string]OrgSettingCustomProbesValue `cty:"custom_probes" hcl:"custom_probes"`
	Disabled       *bool                                  `cty:"disabled" hcl:"disabled"`
	LanNetworks    []OrgSettingLanNetworksValue           `cty:"lan_networks" hcl:"lan_networks"`
	Vlans          []OrgSettingVlansValue                 `cty:"vlans" hcl:"vlans"`
	WanSpeedtest   *OrgSettingWanSpeedtestValue           `cty:"wan_speedtest" hcl:"wan_speedtest"`
}

type OrgSettingCustomProbesValue struct {
	Aggressiveness   *string `cty:"aggressiveness" hcl:"aggressiveness"`
	Target           *string `cty:"target" hcl:"target"`
	Threshold        *int64  `cty:"threshold" hcl:"threshold"`
	CustomProbesType *string `cty:"type" hcl:"type"`
}

type OrgSettingLanNetworksValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
	Probes   []string `cty:"probes" hcl:"probes"`
}

type OrgSettingVlansValue struct {
	CustomTestUrls []string `cty:"custom_test_urls" hcl:"custom_test_urls"`
	Disabled       *bool    `cty:"disabled" hcl:"disabled"`
	Probes         []string `cty:"probes" hcl:"probes"`
	VlanIds        []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type OrgSettingWanSpeedtestValue struct {
	Enabled   *bool   `cty:"enabled" hcl:"enabled"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type OrgSettingVpnOptionsValue struct {
	AsBase     *int64  `cty:"as_base" hcl:"as_base"`
	EnableIpv6 *bool   `cty:"enable_ipv6" hcl:"enable_ipv6"`
	StSubnet   *string `cty:"st_subnet" hcl:"st_subnet"`
}

type OrgSettingWanPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgSettingWiredPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgSettingWirelessPmaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}
