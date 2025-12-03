package provider

type SiteSettingModel struct {
	Analytic                   *SiteSettingAnalyticValue              `hcl:"analytic"`
	ApUpdownThreshold          *int64                                 `hcl:"ap_updown_threshold"`
	AutoUpgrade                *SiteSettingAutoUpgradeValue           `hcl:"auto_upgrade"`
	AutoUpgradeEsl             *SiteSettingAutoUpgradeEslValue        `hcl:"auto_upgrade_esl"`
	BgpNeighborUpdownThreshold *int64                                 `hcl:"bgp_neighbor_updown_threshold"`
	BleConfig                  *SiteSettingBleConfigValue             `hcl:"ble_config"`
	ConfigAutoRevert           *bool                                  `hcl:"config_auto_revert"`
	ConfigPushPolicy           *SiteSettingConfigPushPolicyValue      `hcl:"config_push_policy"`
	CriticalUrlMonitoring      *SiteSettingCriticalUrlMonitoringValue `hcl:"critical_url_monitoring"`
	DefaultPortUsage           *string                                `hcl:"default_port_usage"`
	DeviceUpdownThreshold      *int64                                 `hcl:"device_updown_threshold"`
	EnableUnii4                *bool                                  `hcl:"enable_unii_4"`
	Engagement                 *SiteSettingEngagementValue            `hcl:"engagement"`
	GatewayMgmt                *SiteSettingGatewayMgmtValue           `hcl:"gateway_mgmt"`
	GatewayUpdownThreshold     *int64                                 `hcl:"gateway_updown_threshold"`
	JuniperSrx                 *SiteSettingJuniperSrxValue            `hcl:"juniper_srx"`
	Led                        *SiteSettingLedValue                   `hcl:"led"`
	Marvis                     *SiteSettingMarvisValue                `hcl:"marvis"`
	Occupancy                  *SiteSettingOccupancyValue             `hcl:"occupancy"`
	PersistConfigOnDevice      *bool                                  `hcl:"persist_config_on_device"`
	Proxy                      *SiteSettingProxyValue                 `hcl:"proxy"`
	RemoveExistingConfigs      *bool                                  `hcl:"remove_existing_configs"`
	ReportGatt                 *bool                                  `hcl:"report_gatt"`
	Rogue                      *SiteSettingRogueValue                 `hcl:"rogue"`
	Rtsa                       *SiteSettingRtsaValue                  `hcl:"rtsa"`
	SimpleAlert                *SiteSettingSimpleAlertValue           `hcl:"simple_alert"`
	SiteId                     string                                 `hcl:"site_id"`
	Skyatp                     *SiteSettingSkyatpValue                `hcl:"skyatp"`
	SleThresholds              *SiteSettingSleThresholdsValue         `hcl:"sle_thresholds"`
	SrxApp                     *SiteSettingSrxAppValue                `hcl:"srx_app"`
	SshKeys                    []string                               `hcl:"ssh_keys"`
	Ssr                        *SiteSettingSsrValue                   `hcl:"ssr"`
	SwitchUpdownThreshold      *int64                                 `hcl:"switch_updown_threshold"`
	SyntheticTest              *SiteSettingSyntheticTestValue         `hcl:"synthetic_test"`
	TrackAnonymousDevices      *bool                                  `hcl:"track_anonymous_devices"`
	UplinkPortConfig           *SiteSettingUplinkPortConfigValue      `hcl:"uplink_port_config"`
	Vars                       map[string]string                      `hcl:"vars"`
	Vna                        *SiteSettingVnaValue                   `hcl:"vna"`
	VpnPathUpdownThreshold     *int64                                 `hcl:"vpn_path_updown_threshold"`
	VpnPeerUpdownThreshold     *int64                                 `hcl:"vpn_peer_updown_threshold"`
	VsInstance                 map[string]SiteSettingVsInstanceValue  `hcl:"vs_instance"`
	WanVna                     *SiteSettingWanVnaValue                `hcl:"wan_vna"`
	Wids                       *SiteSettingWidsValue                  `hcl:"wids"`
	Wifi                       *SiteSettingWifiValue                  `hcl:"wifi"`
	WiredVna                   *SiteSettingWiredVnaValue              `hcl:"wired_vna"`
	ZoneOccupancyAlert         *SiteSettingZoneOccupancyAlertValue    `hcl:"zone_occupancy_alert"`
}

type SiteSettingAnalyticValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteSettingAutoUpgradeValue struct {
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	DayOfWeek      *string           `cty:"day_of_week" hcl:"day_of_week"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	TimeOfDay      *string           `cty:"time_of_day" hcl:"time_of_day"`
	Version        *string           `cty:"version" hcl:"version"`
}

type SiteSettingAutoUpgradeEslValue struct {
	AllowDowngrade *bool             `cty:"allow_downgrade" hcl:"allow_downgrade"`
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	DayOfWeek      *string           `cty:"day_of_week" hcl:"day_of_week"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	TimeOfDay      *string           `cty:"time_of_day" hcl:"time_of_day"`
	Version        *string           `cty:"version" hcl:"version"`
}

type SiteSettingBleConfigValue struct {
	BeaconEnabled           *bool   `cty:"beacon_enabled" hcl:"beacon_enabled"`
	BeaconRate              *int64  `cty:"beacon_rate" hcl:"beacon_rate"`
	BeaconRateMode          *string `cty:"beacon_rate_mode" hcl:"beacon_rate_mode"`
	BeamDisabled            []int64 `cty:"beam_disabled" hcl:"beam_disabled"`
	CustomBlePacketEnabled  *bool   `cty:"custom_ble_packet_enabled" hcl:"custom_ble_packet_enabled"`
	CustomBlePacketFrame    *string `cty:"custom_ble_packet_frame" hcl:"custom_ble_packet_frame"`
	CustomBlePacketFreqMsec *int64  `cty:"custom_ble_packet_freq_msec" hcl:"custom_ble_packet_freq_msec"`
	EddystoneUidAdvPower    *int64  `cty:"eddystone_uid_adv_power" hcl:"eddystone_uid_adv_power"`
	EddystoneUidBeams       *string `cty:"eddystone_uid_beams" hcl:"eddystone_uid_beams"`
	EddystoneUidEnabled     *bool   `cty:"eddystone_uid_enabled" hcl:"eddystone_uid_enabled"`
	EddystoneUidFreqMsec    *int64  `cty:"eddystone_uid_freq_msec" hcl:"eddystone_uid_freq_msec"`
	EddystoneUidInstance    *string `cty:"eddystone_uid_instance" hcl:"eddystone_uid_instance"`
	EddystoneUidNamespace   *string `cty:"eddystone_uid_namespace" hcl:"eddystone_uid_namespace"`
	EddystoneUrlAdvPower    *int64  `cty:"eddystone_url_adv_power" hcl:"eddystone_url_adv_power"`
	EddystoneUrlBeams       *string `cty:"eddystone_url_beams" hcl:"eddystone_url_beams"`
	EddystoneUrlEnabled     *bool   `cty:"eddystone_url_enabled" hcl:"eddystone_url_enabled"`
	EddystoneUrlFreqMsec    *int64  `cty:"eddystone_url_freq_msec" hcl:"eddystone_url_freq_msec"`
	EddystoneUrlUrl         *string `cty:"eddystone_url_url" hcl:"eddystone_url_url"`
	IbeaconAdvPower         *int64  `cty:"ibeacon_adv_power" hcl:"ibeacon_adv_power"`
	IbeaconBeams            *string `cty:"ibeacon_beams" hcl:"ibeacon_beams"`
	IbeaconEnabled          *bool   `cty:"ibeacon_enabled" hcl:"ibeacon_enabled"`
	IbeaconFreqMsec         *int64  `cty:"ibeacon_freq_msec" hcl:"ibeacon_freq_msec"`
	IbeaconMajor            *int64  `cty:"ibeacon_major" hcl:"ibeacon_major"`
	IbeaconMinor            *int64  `cty:"ibeacon_minor" hcl:"ibeacon_minor"`
	IbeaconUuid             *string `cty:"ibeacon_uuid" hcl:"ibeacon_uuid"`
	Power                   *int64  `cty:"power" hcl:"power"`
	PowerMode               *string `cty:"power_mode" hcl:"power_mode"`
}

type SiteSettingConfigPushPolicyValue struct {
	NoPush     *bool                       `cty:"no_push" hcl:"no_push"`
	PushWindow *SiteSettingPushWindowValue `cty:"push_window" hcl:"push_window"`
}

type SiteSettingPushWindowValue struct {
	Enabled *bool                  `cty:"enabled" hcl:"enabled"`
	Hours   *SiteSettingHoursValue `cty:"hours" hcl:"hours"`
}

type SiteSettingHoursValue struct {
	Fri *string `cty:"fri" hcl:"fri"`
	Mon *string `cty:"mon" hcl:"mon"`
	Sat *string `cty:"sat" hcl:"sat"`
	Sun *string `cty:"sun" hcl:"sun"`
	Thu *string `cty:"thu" hcl:"thu"`
	Tue *string `cty:"tue" hcl:"tue"`
	Wed *string `cty:"wed" hcl:"wed"`
}

type SiteSettingCriticalUrlMonitoringValue struct {
	Enabled  *bool                      `cty:"enabled" hcl:"enabled"`
	Monitors []SiteSettingMonitorsValue `cty:"monitors" hcl:"monitors"`
}

type SiteSettingMonitorsValue struct {
	Url    *string `cty:"url" hcl:"url"`
	VlanId *string `cty:"vlan_id" hcl:"vlan_id"`
}

type SiteSettingEngagementValue struct {
	DwellTagNames *SiteSettingDwellTagNamesValue `cty:"dwell_tag_names" hcl:"dwell_tag_names"`
	DwellTags     *SiteSettingDwellTagsValue     `cty:"dwell_tags" hcl:"dwell_tags"`
	Hours         *SiteSettingHoursValue         `cty:"hours" hcl:"hours"`
	MaxDwell      *int64                         `cty:"max_dwell" hcl:"max_dwell"`
	MinDwell      *int64                         `cty:"min_dwell" hcl:"min_dwell"`
}

type SiteSettingDwellTagNamesValue struct {
	Bounce    *string `cty:"bounce" hcl:"bounce"`
	Engaged   *string `cty:"engaged" hcl:"engaged"`
	Passerby  *string `cty:"passerby" hcl:"passerby"`
	Stationed *string `cty:"stationed" hcl:"stationed"`
}

type SiteSettingDwellTagsValue struct {
	Bounce    *string `cty:"bounce" hcl:"bounce"`
	Engaged   *string `cty:"engaged" hcl:"engaged"`
	Passerby  *string `cty:"passerby" hcl:"passerby"`
	Stationed *string `cty:"stationed" hcl:"stationed"`
}

type SiteSettingGatewayMgmtValue struct {
	AdminSshkeys               []string                             `cty:"admin_sshkeys" hcl:"admin_sshkeys"`
	AppProbing                 *SiteSettingAppProbingValue          `cty:"app_probing" hcl:"app_probing"`
	AppUsage                   *bool                                `cty:"app_usage" hcl:"app_usage"`
	AutoSignatureUpdate        *SiteSettingAutoSignatureUpdateValue `cty:"auto_signature_update" hcl:"auto_signature_update"`
	ConfigRevertTimer          *int64                               `cty:"config_revert_timer" hcl:"config_revert_timer"`
	DisableConsole             *bool                                `cty:"disable_console" hcl:"disable_console"`
	DisableOob                 *bool                                `cty:"disable_oob" hcl:"disable_oob"`
	DisableUsb                 *bool                                `cty:"disable_usb" hcl:"disable_usb"`
	FipsEnabled                *bool                                `cty:"fips_enabled" hcl:"fips_enabled"`
	ProbeHosts                 []string                             `cty:"probe_hosts" hcl:"probe_hosts"`
	ProbeHostsv6               []string                             `cty:"probe_hostsv6" hcl:"probe_hostsv6"`
	ProtectRe                  *SiteSettingProtectReValue           `cty:"protect_re" hcl:"protect_re"`
	RootPassword               *string                              `cty:"root_password" hcl:"root_password"`
	SecurityLogSourceAddress   *string                              `cty:"security_log_source_address" hcl:"security_log_source_address"`
	SecurityLogSourceInterface *string                              `cty:"security_log_source_interface" hcl:"security_log_source_interface"`
}

type SiteSettingAppProbingValue struct {
	Apps       []string                     `cty:"apps" hcl:"apps"`
	CustomApps []SiteSettingCustomAppsValue `cty:"custom_apps" hcl:"custom_apps"`
	Enabled    *bool                        `cty:"enabled" hcl:"enabled"`
}

type SiteSettingCustomAppsValue struct {
	AppType    *string  `cty:"app_type" hcl:"app_type"`
	Hostnames  []string `cty:"hostnames" hcl:"hostnames"`
	Name       string   `cty:"name" hcl:"name"`
	Network    *string  `cty:"network" hcl:"network"`
	PacketSize *int64   `cty:"packet_size" hcl:"packet_size"`
	Protocol   string   `cty:"protocol" hcl:"protocol"`
	Vrf        *string  `cty:"vrf" hcl:"vrf"`
}

type SiteSettingAutoSignatureUpdateValue struct {
	DayOfWeek *string `cty:"day_of_week" hcl:"day_of_week"`
	Enable    *bool   `cty:"enable" hcl:"enable"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type SiteSettingProtectReValue struct {
	AllowedServices []string                 `cty:"allowed_services" hcl:"allowed_services"`
	Custom          []SiteSettingCustomValue `cty:"custom" hcl:"custom"`
	Enabled         *bool                    `cty:"enabled" hcl:"enabled"`
	HitCount        *bool                    `cty:"hit_count" hcl:"hit_count"`
	TrustedHosts    []string                 `cty:"trusted_hosts" hcl:"trusted_hosts"`
}

type SiteSettingCustomValue struct {
	PortRange *string  `cty:"port_range" hcl:"port_range"`
	Protocol  *string  `cty:"protocol" hcl:"protocol"`
	Subnets   []string `cty:"subnets" hcl:"subnets"`
}

type SiteSettingJuniperSrxValue struct {
	Gateways            []SiteSettingGatewaysValue      `cty:"gateways" hcl:"gateways"`
	SendMistNacUserInfo *bool                           `cty:"send_mist_nac_user_info" hcl:"send_mist_nac_user_info"`
	SrxAutoUpgrade      *SiteSettingSrxAutoUpgradeValue `cty:"auto_upgrade" hcl:"auto_upgrade"`
}

type SiteSettingGatewaysValue struct {
	ApiKey      *string `cty:"api_key" hcl:"api_key"`
	ApiPassword *string `cty:"api_password" hcl:"api_password"`
	ApiUrl      *string `cty:"api_url" hcl:"api_url"`
}

type SiteSettingSrxAutoUpgradeValue struct {
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
	Snapshot       *bool             `cty:"snapshot" hcl:"snapshot"`
}

type SiteSettingLedValue struct {
	Brightness *int64 `cty:"brightness" hcl:"brightness"`
	Enabled    *bool  `cty:"enabled" hcl:"enabled"`
}

type SiteSettingMarvisValue struct {
	AutoOperations *SiteSettingAutoOperationsValue `cty:"auto_operations" hcl:"auto_operations"`
}

type SiteSettingAutoOperationsValue struct {
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

type SiteSettingOccupancyValue struct {
	AssetsEnabled             *bool  `cty:"assets_enabled" hcl:"assets_enabled"`
	ClientsEnabled            *bool  `cty:"clients_enabled" hcl:"clients_enabled"`
	MinDuration               *int64 `cty:"min_duration" hcl:"min_duration"`
	SdkclientsEnabled         *bool  `cty:"sdkclients_enabled" hcl:"sdkclients_enabled"`
	UnconnectedClientsEnabled *bool  `cty:"unconnected_clients_enabled" hcl:"unconnected_clients_enabled"`
}

type SiteSettingProxyValue struct {
	Disabled *bool   `cty:"disabled" hcl:"disabled"`
	Url      *string `cty:"url" hcl:"url"`
}

type SiteSettingRogueValue struct {
	AllowedVlanIds    []int64  `cty:"allowed_vlan_ids" hcl:"allowed_vlan_ids"`
	Enabled           *bool    `cty:"enabled" hcl:"enabled"`
	HoneypotEnabled   *bool    `cty:"honeypot_enabled" hcl:"honeypot_enabled"`
	MinDuration       *int64   `cty:"min_duration" hcl:"min_duration"`
	MinRogueDuration  *int64   `cty:"min_rogue_duration" hcl:"min_rogue_duration"`
	MinRogueRssi      *int64   `cty:"min_rogue_rssi" hcl:"min_rogue_rssi"`
	MinRssi           *int64   `cty:"min_rssi" hcl:"min_rssi"`
	WhitelistedBssids []string `cty:"whitelisted_bssids" hcl:"whitelisted_bssids"`
	WhitelistedSsids  []string `cty:"whitelisted_ssids" hcl:"whitelisted_ssids"`
}

type SiteSettingRtsaValue struct {
	AppWaking             *bool `cty:"app_waking" hcl:"app_waking"`
	DisableDeadReckoning  *bool `cty:"disable_dead_reckoning" hcl:"disable_dead_reckoning"`
	DisablePressureSensor *bool `cty:"disable_pressure_sensor" hcl:"disable_pressure_sensor"`
	Enabled               *bool `cty:"enabled" hcl:"enabled"`
	TrackAsset            *bool `cty:"track_asset" hcl:"track_asset"`
}

type SiteSettingSimpleAlertValue struct {
	ArpFailure  *SiteSettingArpFailureValue  `cty:"arp_failure" hcl:"arp_failure"`
	DhcpFailure *SiteSettingDhcpFailureValue `cty:"dhcp_failure" hcl:"dhcp_failure"`
	DnsFailure  *SiteSettingDnsFailureValue  `cty:"dns_failure" hcl:"dns_failure"`
}

type SiteSettingArpFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type SiteSettingDhcpFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type SiteSettingDnsFailureValue struct {
	ClientCount   *int64 `cty:"client_count" hcl:"client_count"`
	Duration      *int64 `cty:"duration" hcl:"duration"`
	IncidentCount *int64 `cty:"incident_count" hcl:"incident_count"`
}

type SiteSettingSkyatpValue struct {
	Enabled          *bool `cty:"enabled" hcl:"enabled"`
	SendIpMacMapping *bool `cty:"send_ip_mac_mapping" hcl:"send_ip_mac_mapping"`
}

type SiteSettingSleThresholdsValue struct {
	Capacity      *int64 `cty:"capacity" hcl:"capacity"`
	Coverage      *int64 `cty:"coverage" hcl:"coverage"`
	Throughput    *int64 `cty:"throughput" hcl:"throughput"`
	Timetoconnect *int64 `cty:"timetoconnect" hcl:"timetoconnect"`
}

type SiteSettingSrxAppValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteSettingSsrValue struct {
	ConductorHosts []string                        `cty:"conductor_hosts" hcl:"conductor_hosts"`
	ConductorToken *string                         `cty:"conductor_token" hcl:"conductor_token"`
	DisableStats   *bool                           `cty:"disable_stats" hcl:"disable_stats"`
	Proxy          *SiteSettingProxyValue          `cty:"proxy" hcl:"proxy"`
	SsrAutoUpgrade *SiteSettingSsrAutoUpgradeValue `cty:"auto_upgrade" hcl:"auto_upgrade"`
}

type SiteSettingSsrAutoUpgradeValue struct {
	Channel        *string           `cty:"channel" hcl:"channel"`
	CustomVersions map[string]string `cty:"custom_versions" hcl:"custom_versions"`
	Enabled        *bool             `cty:"enabled" hcl:"enabled"`
}

type SiteSettingSyntheticTestValue struct {
	Aggressiveness *string                                 `cty:"aggressiveness" hcl:"aggressiveness"`
	CustomProbes   map[string]SiteSettingCustomProbesValue `cty:"custom_probes" hcl:"custom_probes"`
	Disabled       *bool                                   `cty:"disabled" hcl:"disabled"`
	LanNetworks    []SiteSettingLanNetworksValue           `cty:"lan_networks" hcl:"lan_networks"`
	Vlans          []SiteSettingVlansValue                 `cty:"vlans" hcl:"vlans"`
	WanSpeedtest   *SiteSettingWanSpeedtestValue           `cty:"wan_speedtest" hcl:"wan_speedtest"`
}

type SiteSettingCustomProbesValue struct {
	Aggressiveness   *string `cty:"aggressiveness" hcl:"aggressiveness"`
	Host             *string `cty:"host" hcl:"host"`
	Port             *int64  `cty:"port" hcl:"port"`
	Threshold        *int64  `cty:"threshold" hcl:"threshold"`
	CustomProbesType *string `cty:"type" hcl:"type"`
	Url              *string `cty:"url" hcl:"url"`
}

type SiteSettingLanNetworksValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
	Probes   []string `cty:"probes" hcl:"probes"`
}

type SiteSettingVlansValue struct {
	CustomTestUrls []string `cty:"custom_test_urls" hcl:"custom_test_urls"`
	Disabled       *bool    `cty:"disabled" hcl:"disabled"`
	Probes         []string `cty:"probes" hcl:"probes"`
	VlanIds        []string `cty:"vlan_ids" hcl:"vlan_ids"`
}

type SiteSettingWanSpeedtestValue struct {
	Enabled   *bool   `cty:"enabled" hcl:"enabled"`
	TimeOfDay *string `cty:"time_of_day" hcl:"time_of_day"`
}

type SiteSettingUplinkPortConfigValue struct {
	Dot1x             *bool `cty:"dot1x" hcl:"dot1x"`
	KeepWlansUpIfDown *bool `cty:"keep_wlans_up_if_down" hcl:"keep_wlans_up_if_down"`
}

type SiteSettingVnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteSettingVsInstanceValue struct {
	Networks []string `cty:"networks" hcl:"networks"`
}

type SiteSettingWanVnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteSettingWidsValue struct {
	RepeatedAuthFailures *SiteSettingRepeatedAuthFailuresValue `cty:"repeated_auth_failures" hcl:"repeated_auth_failures"`
}

type SiteSettingRepeatedAuthFailuresValue struct {
	Duration  *int64 `cty:"duration" hcl:"duration"`
	Threshold *int64 `cty:"threshold" hcl:"threshold"`
}

type SiteSettingWifiValue struct {
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

type SiteSettingWiredVnaValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type SiteSettingZoneOccupancyAlertValue struct {
	EmailNotifiers []string `cty:"email_notifiers" hcl:"email_notifiers"`
	Enabled        *bool    `cty:"enabled" hcl:"enabled"`
	Threshold      *int64   `cty:"threshold" hcl:"threshold"`
}
