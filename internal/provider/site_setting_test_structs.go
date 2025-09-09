package provider

type SiteSettingModel struct {
	Analytic               *AnalyticValue                 `hcl:"analytic"`
	ApUpdownThreshold      *int64                         `hcl:"ap_updown_threshold"`
	AutoUpgrade            *AutoUpgradeValue              `hcl:"auto_upgrade"`
	BleConfig              *BleConfigValue                `hcl:"ble_config"`
	ConfigAutoRevert       *bool                          `hcl:"config_auto_revert"`
	ConfigPushPolicy       *ConfigPushPolicyValue         `hcl:"config_push_policy"`
	CriticalUrlMonitoring  *CriticalUrlMonitoringValue    `hcl:"critical_url_monitoring"`
	DeviceUpdownThreshold  *int64                         `hcl:"device_updown_threshold"`
	Engagement             *EngagementValue               `hcl:"engagement"`
	GatewayMgmt            *GatewayMgmtValue              `hcl:"gateway_mgmt"`
	GatewayUpdownThreshold *int64                         `hcl:"gateway_updown_threshold"`
	JuniperSrx             *JuniperSrxValue               `hcl:"juniper_srx"`
	Led                    *LedValue                      `hcl:"led"`
	Occupancy              *OccupancyValue                `hcl:"occupancy"`
	PersistConfigOnDevice  *bool                          `hcl:"persist_config_on_device"`
	Proxy                  *ProxyValue                    `hcl:"proxy"`
	RemoveExistingConfigs  *bool                          `hcl:"remove_existing_configs"`
	ReportGatt             *bool                          `hcl:"report_gatt"`
	Rogue                  *RogueValue                    `hcl:"rogue"`
	Rtsa                   *RtsaValue                     `hcl:"rtsa"`
	SimpleAlert            *SimpleAlertValue              `hcl:"simple_alert"`
	SiteId                 string                         `hcl:"site_id"`
	Skyatp                 *SkyatpValue                   `hcl:"skyatp"`
	SrxApp                 *SrxAppValue                   `hcl:"srx_app"`
	SshKeys                []string                       `hcl:"ssh_keys"`
	Ssr                    *SsrValue                      `hcl:"ssr"`
	SwitchUpdownThreshold  *int64                         `hcl:"switch_updown_threshold"`
	SyntheticTest          *SiteSettingSyntheticTestValue `hcl:"synthetic_test"`
	TrackAnonymousDevices  *bool                          `hcl:"track_anonymous_devices"`
	UplinkPortConfig       *UplinkPortConfigValue         `hcl:"uplink_port_config"`
	Vars                   map[string]string              `hcl:"vars"`
	Vna                    *VnaValue                      `hcl:"vna"`
	VsInstance             map[string]VsInstanceValue     `hcl:"vs_instance"`
	WanVna                 *WanVnaValue                   `hcl:"wan_vna"`
	Wids                   *WidsValue                     `hcl:"wids"`
	Wifi                   *WifiValue                     `hcl:"wifi"`
	WiredVna               *WiredVnaValue                 `hcl:"wired_vna"`
	ZoneOccupancyAlert     *ZoneOccupancyAlertValue       `hcl:"zone_occupancy_alert"`
}

type AnalyticValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type AutoUpgradeValue struct {
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	DayOfWeek      *string           `cty:"day_of_week" hcl:"day_of_week"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	TimeOfDay      *string           `cty:"time_of_day" hcl:"time_of_day"`
	Version        *string           `cty:"version" hcl:"version"`
}

type ConfigPushPolicyValue struct {
	NoPush     *bool            `cty:"no_push" hcl:"no_push"`
	PushWindow *PushWindowValue `cty:"push_window" hcl:"push_window"`
}

type PushWindowValue struct {
	Enabled *bool       `cty:"enabled" hcl:"enabled"`
	Hours   *HoursValue `cty:"hours" hcl:"hours"`
}

type CriticalUrlMonitoringValue struct {
	Enabled  *bool           `cty:"enabled" hcl:"enabled"`
	Monitors []MonitorsValue `cty:"monitors" hcl:"monitors"`
}

type MonitorsValue struct {
	Url    *string `cty:"url" hcl:"url"`
	VlanId *string `cty:"vlan_id" hcl:"vlan_id"`
}

type EngagementValue struct {
	DwellTagNames *DwellTagNamesValue `cty:"dwell_tag_names" hcl:"dwell_tag_names"`
	DwellTags     *DwellTagsValue     `cty:"dwell_tags" hcl:"dwell_tags"`
	Hours         *HoursValue         `cty:"hours" hcl:"hours"`
	MaxDwell      *int64              `cty:"max_dwell" hcl:"max_dwell"`
	MinDwell      *int64              `cty:"min_dwell" hcl:"min_dwell"`
}

type DwellTagNamesValue struct {
	Bounce    *string `cty:"bounce" hcl:"bounce"`
	Engaged   *string `cty:"engaged" hcl:"engaged"`
	Passerby  *string `cty:"passerby" hcl:"passerby"`
	Stationed *string `cty:"stationed" hcl:"stationed"`
}

type DwellTagsValue struct {
	Bounce    *string `cty:"bounce" hcl:"bounce"`
	Engaged   *string `cty:"engaged" hcl:"engaged"`
	Passerby  *string `cty:"passerby" hcl:"passerby"`
	Stationed *string `cty:"stationed" hcl:"stationed"`
}

type GatewayMgmtValue struct {
	AdminSshkeys               []string                  `cty:"admin_sshkeys" hcl:"admin_sshkeys"`
	AppProbing                 *AppProbingValue          `cty:"app_probing" hcl:"app_probing"`
	AppUsage                   *bool                     `cty:"app_usage" hcl:"app_usage"`
	AutoSignatureUpdate        *AutoSignatureUpdateValue `cty:"auto_signature_update" hcl:"auto_signature_update"`
	ConfigRevertTimer          *int64                    `cty:"config_revert_timer" hcl:"config_revert_timer"`
	DisableConsole             *bool                     `cty:"disable_console" hcl:"disable_console"`
	DisableOob                 *bool                     `cty:"disable_oob" hcl:"disable_oob"`
	ProbeHosts                 []string                  `cty:"probe_hosts" hcl:"probe_hosts"`
	ProtectRe                  *ProtectReValue           `cty:"protect_re" hcl:"protect_re"`
	RootPassword               *string                   `cty:"root_password" hcl:"root_password"`
	SecurityLogSourceAddress   *string                   `cty:"security_log_source_address" hcl:"security_log_source_address"`
	SecurityLogSourceInterface *string                   `cty:"security_log_source_interface" hcl:"security_log_source_interface"`
}

type AppProbingValue struct {
	Apps       []string          `cty:"apps" hcl:"apps"`
	CustomApps []CustomAppsValue `cty:"custom_apps" hcl:"custom_apps"`
	Enabled    *bool             `cty:"enabled" hcl:"enabled"`
}

type CustomAppsValue struct {
	AppType    *string  `cty:"app_type" hcl:"app_type"`
	Hostnames  []string `cty:"hostnames" hcl:"hostnames"`
	Name       string   `cty:"name" hcl:"name"`
	Network    *string  `cty:"network" hcl:"network"`
	PacketSize *int64   `cty:"packet_size" hcl:"packet_size"`
	Protocol   string   `cty:"protocol" hcl:"protocol"`
	Vrf        *string  `cty:"vrf" hcl:"vrf"`
}

type AutoSignatureUpdateValue struct {
	DayOfWeek *string `cty:"day_of_week" hcl:"day_of_week"`
	Enable    *bool   `cty:"enable" hcl:"enable"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type JuniperSrxValue struct {
	Gateways            []GatewaysValue `cty:"gateways" hcl:"gateways"`
	SendMistNacUserInfo *bool           `cty:"send_mist_nac_user_info" hcl:"send_mist_nac_user_info"`
}

type GatewaysValue struct {
	ApiKey *string `cty:"api_key" hcl:"api_key"`
	ApiUrl *string `cty:"api_url" hcl:"api_url"`
}

type OccupancyValue struct {
	AssetsEnabled             *bool  `cty:"assets_enabled" hcl:"assets_enabled"`
	ClientsEnabled            *bool  `cty:"clients_enabled" hcl:"clients_enabled"`
	MinDuration               *int64 `cty:"min_duration" hcl:"min_duration"`
	SdkclientsEnabled         *bool  `cty:"sdkclients_enabled" hcl:"sdkclients_enabled"`
	UnconnectedClientsEnabled *bool  `cty:"unconnected_clients_enabled" hcl:"unconnected_clients_enabled"`
}

type ProxyValue struct {
	Url *string `cty:"url" hcl:"url"`
}

type RogueValue struct {
	Enabled           *bool    `cty:"enabled" hcl:"enabled"`
	HoneypotEnabled   *bool    `cty:"honeypot_enabled" hcl:"honeypot_enabled"`
	MinDuration       *int64   `cty:"min_duration" hcl:"min_duration"`
	MinRssi           *int64   `cty:"min_rssi" hcl:"min_rssi"`
	WhitelistedBssids []string `cty:"whitelisted_bssids" hcl:"whitelisted_bssids"`
	WhitelistedSsids  []string `cty:"whitelisted_ssids" hcl:"whitelisted_ssids"`
}

type RtsaValue struct {
	AppWaking             *bool `cty:"app_waking" hcl:"app_waking"`
	DisableDeadReckoning  *bool `cty:"disable_dead_reckoning" hcl:"disable_dead_reckoning"`
	DisablePressureSensor *bool `cty:"disable_pressure_sensor" hcl:"disable_pressure_sensor"`
	Enabled               *bool `cty:"enabled" hcl:"enabled"`
	TrackAsset            *bool `cty:"track_asset" hcl:"track_asset"`
}

type SimpleAlertValue struct {
	ArpFailure  *ArpFailureValue  `cty:"arp_failure" hcl:"arp_failure"`
	DhcpFailure *DhcpFailureValue `cty:"dhcp_failure" hcl:"dhcp_failure"`
	DnsFailure  *DnsFailureValue  `cty:"dns_failure" hcl:"dns_failure"`
}

type ArpFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type DhcpFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type DnsFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type SkyatpValue struct {
	Enabled          *bool `cty:"enabled" hcl:"enabled"`
	SendIpMacMapping *bool `cty:"send_ip_mac_mapping" hcl:"send_ip_mac_mapping"`
}

type SrxAppValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SsrValue struct {
	ConductorHosts []string `cty:"conductor_hosts" hcl:"conductor_hosts"`
	ConductorToken *string  `cty:"conductor_token" hcl:"conductor_token"`
	DisableStats   *bool    `cty:"disable_stats" hcl:"disable_stats"`
}

type VnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type VsInstanceValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
}

type WanVnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type WidsValue struct {
	RepeatedAuthFailures *RepeatedAuthFailuresValue `cty:"repeated_auth_failures" hcl:"repeated_auth_failures"`
}

type RepeatedAuthFailuresValue struct {
	Duration  *int64 `cty:"duration" hcl:"duration"`
	Threshold *int64 `cty:"threshold" hcl:"threshold"`
}

type WifiValue struct {
	CiscoEnabled                      *bool   `cty:"cisco_enabled" hcl:"cisco_enabled"`
	Disable11k                        *bool   `cty:"disable_11k" hcl:"disable_11k"`
	DisableRadiosWhenPowerConstrained *bool   `cty:"disable_radios_when_power_constrained" hcl:"disable_radios_when_power_constrained"`
	EnableArpSpoofCheck               *bool   `cty:"enable_arp_spoof_check" hcl:"enable_arp_spoof_check"`
	EnableSharedRadioScanning         *bool   `cty:"enable_shared_radio_scanning" hcl:"enable_shared_radio_scanning"`
	Enabled                           *bool   `cty:"enabled" hcl:"enabled"`
	LocateConnected                   *bool   `cty:"locate_connected" hcl:"locate_connected"`
	LocateUnconnected                 *bool   `cty:"locate_unconnected" hcl:"locate_unconnected"`
	MeshAllowDfs                      *bool   `cty:"mesh_allow_dfs" hcl:"mesh_allow_dfs"`
	MeshEnableCrm                     *bool   `cty:"mesh_enable_crm" hcl:"mesh_enable_crm"`
	MeshEnabled                       *bool   `cty:"mesh_enabled" hcl:"mesh_enabled"`
	MeshPsk                           *string `cty:"mesh_psk" hcl:"mesh_psk"`
	MeshSsid                          *string `cty:"mesh_ssid" hcl:"mesh_ssid"`
	ProxyArp                          *string `cty:"proxy_arp" hcl:"proxy_arp"`
}

type WiredVnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type ZoneOccupancyAlertValue struct {
	EmailNotifiers []string `cty:"email_notifiers" hcl:"email_notifiers"`
	Enabled        *bool    `cty:"enabled" hcl:"enabled"`
	Threshold      *int64   `cty:"threshold" hcl:"threshold"`
}

type SiteSettingSyntheticTestValue struct {
	Aggressiveness *string                      `cty:"aggressiveness" hcl:"aggressiveness"`
	CustomProbes   map[string]CustomProbesValue `cty:"custom_probes" hcl:"custom_probes"`
	Disabled       *bool                        `cty:"disabled" hcl:"disabled"`
	LanNetworks    []LanNetworksValue           `cty:"lan_networks" hcl:"lan_networks"`
	Vlans          []SiteSettingVlansValue      `cty:"vlans" hcl:"vlans"`
	WanSpeedtest   *WanSpeedtestValue           `cty:"wan_speedtest" hcl:"wan_speedtest"`
}

type CustomProbesValue struct {
	Aggressiveness *string `cty:"aggressiveness" hcl:"aggressiveness"`
	Host           *string `cty:"host" hcl:"host"`
	Port           *int64  `cty:"port" hcl:"port"`
	Threshold      *int64  `cty:"threshold" hcl:"threshold"`
	Type           *string `cty:"type" hcl:"type"`
	Url            *string `cty:"url" hcl:"url"`
}

type LanNetworksValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
	Probes   []string `cty:"probes" hcl:"probes"`
}

type SiteSettingVlansValue struct {
	CustomTestUrls []string `cty:"custom_test_urls" hcl:"custom_test_urls"`
	Disabled       *bool    `cty:"disabled" hcl:"disabled"`
	Probes         []string `cty:"probes" hcl:"probes"`
	VlanIds        []string `cty:"vlan_ids" hcl:"vlan_ids"`
}
