package provider

type OrgDeviceprofileApModel struct {
	Aeroscout        *OrgDeviceprofileApAeroscoutValue            `hcl:"aeroscout"`
	Airista          *OrgDeviceprofileApAiristaValue              `hcl:"airista"`
	BleConfig        *OrgDeviceprofileApBleConfigValue            `hcl:"ble_config"`
	DisableEth1      *bool                                        `hcl:"disable_eth1"`
	DisableEth2      *bool                                        `hcl:"disable_eth2"`
	DisableEth3      *bool                                        `hcl:"disable_eth3"`
	DisableModule    *bool                                        `hcl:"disable_module"`
	EslConfig        *OrgDeviceprofileApEslConfigValue            `hcl:"esl_config"`
	IpConfig         *OrgDeviceprofileApIpConfigValue             `hcl:"ip_config"`
	LacpConfig       *OrgDeviceprofileApLacpConfigValue           `hcl:"lacp_config"`
	Led              *OrgDeviceprofileApLedValue                  `hcl:"led"`
	Mesh             *OrgDeviceprofileApMeshValue                 `hcl:"mesh"`
	Name             string                                       `hcl:"name"`
	NtpServers       []string                                     `hcl:"ntp_servers"`
	OrgId            string                                       `hcl:"org_id"`
	PoePassthrough   *bool                                        `hcl:"poe_passthrough"`
	PortConfig       map[string]OrgDeviceprofileApPortConfigValue `hcl:"port_config"`
	PwrConfig        *OrgDeviceprofileApPwrConfigValue            `hcl:"pwr_config"`
	RadioConfig      *OrgDeviceprofileApRadioConfigValue          `hcl:"radio_config"`
	SiteId           *string                                      `hcl:"site_id"`
	UplinkPortConfig *OrgDeviceprofileApUplinkPortConfigValue     `hcl:"uplink_port_config"`
	UsbConfig        *OrgDeviceprofileApUsbConfigValue            `hcl:"usb_config"`
	Vars             map[string]string                            `hcl:"vars"`
}

type OrgDeviceprofileApAeroscoutValue struct {
	Enabled         *bool   `cty:"enabled" hcl:"enabled"`
	Host            *string `cty:"host" hcl:"host"`
	LocateConnected *bool   `cty:"locate_connected" hcl:"locate_connected"`
	Port            *int64  `cty:"port" hcl:"port"`
}

type OrgDeviceprofileApAiristaValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Host    *string `cty:"host" hcl:"host"`
	Port    *int64  `cty:"port" hcl:"port"`
}

type OrgDeviceprofileApBleConfigValue struct {
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

type OrgDeviceprofileApEslConfigValue struct {
	Cacert        *string `cty:"cacert" hcl:"cacert"`
	Channel       *int64  `cty:"channel" hcl:"channel"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	Host          *string `cty:"host" hcl:"host"`
	Port          *int64  `cty:"port" hcl:"port"`
	EslConfigType *string `cty:"type" hcl:"type"`
	VerifyCert    *bool   `cty:"verify_cert" hcl:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id" hcl:"vlan_id"`
}

type OrgDeviceprofileApIpConfigValue struct {
	Dns          []string `cty:"dns" hcl:"dns"`
	DnsSuffix    []string `cty:"dns_suffix" hcl:"dns_suffix"`
	Gateway      *string  `cty:"gateway" hcl:"gateway"`
	Gateway6     *string  `cty:"gateway6" hcl:"gateway6"`
	Ip           *string  `cty:"ip" hcl:"ip"`
	Ip6          *string  `cty:"ip6" hcl:"ip6"`
	Mtu          *int64   `cty:"mtu" hcl:"mtu"`
	Netmask      *string  `cty:"netmask" hcl:"netmask"`
	Netmask6     *string  `cty:"netmask6" hcl:"netmask6"`
	IpConfigType *string  `cty:"type" hcl:"type"`
	Type6        *string  `cty:"type6" hcl:"type6"`
	VlanId       *int64   `cty:"vlan_id" hcl:"vlan_id"`
}

type OrgDeviceprofileApLacpConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileApLedValue struct {
	Brightness *int64 `cty:"brightness" hcl:"brightness"`
	Enabled    *bool  `cty:"enabled" hcl:"enabled"`
}

type OrgDeviceprofileApMeshValue struct {
	Bands   []string `cty:"bands" hcl:"bands"`
	Enabled *bool    `cty:"enabled" hcl:"enabled"`
	Group   *int64   `cty:"group" hcl:"group"`
	Role    *string  `cty:"role" hcl:"role"`
}

type OrgDeviceprofileApPortConfigValue struct {
	Disabled         *bool                                `cty:"disabled" hcl:"disabled"`
	DynamicVlan      *OrgDeviceprofileApDynamicVlanValue  `cty:"dynamic_vlan" hcl:"dynamic_vlan"`
	EnableMacAuth    *bool                                `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	Forwarding       *string                              `cty:"forwarding" hcl:"forwarding"`
	MacAuthPreferred *bool                                `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol  *string                              `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MistNac          *OrgDeviceprofileApMistNacValue      `cty:"mist_nac" hcl:"mist_nac"`
	MxTunnelId       *string                              `cty:"mx_tunnel_id" hcl:"mx_tunnel_id"`
	MxtunnelName     *string                              `cty:"mxtunnel_name" hcl:"mxtunnel_name"`
	PortAuth         *string                              `cty:"port_auth" hcl:"port_auth"`
	PortVlanId       *int64                               `cty:"port_vlan_id" hcl:"port_vlan_id"`
	RadiusConfig     *OrgDeviceprofileApRadiusConfigValue `cty:"radius_config" hcl:"radius_config"`
	Radsec           *OrgDeviceprofileApRadsecValue       `cty:"radsec" hcl:"radsec"`
	VlanId           *int64                               `cty:"vlan_id" hcl:"vlan_id"`
	VlanIds          []int64                              `cty:"vlan_ids" hcl:"vlan_ids"`
	WxtunnelId       *string                              `cty:"wxtunnel_id" hcl:"wxtunnel_id"`
	WxtunnelRemoteId *string                              `cty:"wxtunnel_remote_id" hcl:"wxtunnel_remote_id"`
}

type OrgDeviceprofileApDynamicVlanValue struct {
	DefaultVlanId   *int64            `cty:"default_vlan_id" hcl:"default_vlan_id"`
	Enabled         *bool             `cty:"enabled" hcl:"enabled"`
	DynamicVlanType *string           `cty:"type" hcl:"type"`
	Vlans           map[string]string `cty:"vlans" hcl:"vlans"`
}

type OrgDeviceprofileApMistNacValue struct {
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

type OrgDeviceprofileApRadiusConfigValue struct {
	AcctInterimInterval *int64                               `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AcctServers         []OrgDeviceprofileApAcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServers         []OrgDeviceprofileApAuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	AuthServersRetries  *int64                               `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64                               `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool                                `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *int64                               `cty:"coa_port" hcl:"coa_port"`
	Network             *string                              `cty:"network" hcl:"network"`
	SourceIp            *string                              `cty:"source_ip" hcl:"source_ip"`
}

type OrgDeviceprofileApAcctServersValue struct {
	Host           string  `cty:"host" hcl:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port           *string `cty:"port" hcl:"port"`
	Secret         string  `cty:"secret" hcl:"secret"`
}

type OrgDeviceprofileApAuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type OrgDeviceprofileApRadsecValue struct {
	CoaEnabled    *bool                            `cty:"coa_enabled" hcl:"coa_enabled"`
	Enabled       *bool                            `cty:"enabled" hcl:"enabled"`
	IdleTimeout   *string                          `cty:"idle_timeout" hcl:"idle_timeout"`
	MxclusterIds  []string                         `cty:"mxcluster_ids" hcl:"mxcluster_ids"`
	ProxyHosts    []string                         `cty:"proxy_hosts" hcl:"proxy_hosts"`
	ServerName    *string                          `cty:"server_name" hcl:"server_name"`
	Servers       []OrgDeviceprofileApServersValue `cty:"servers" hcl:"servers"`
	UseMxedge     *bool                            `cty:"use_mxedge" hcl:"use_mxedge"`
	UseSiteMxedge *bool                            `cty:"use_site_mxedge" hcl:"use_site_mxedge"`
}

type OrgDeviceprofileApServersValue struct {
	Host *string `cty:"host" hcl:"host"`
	Port *int64  `cty:"port" hcl:"port"`
}

type OrgDeviceprofileApPwrConfigValue struct {
	Base              *int64 `cty:"base" hcl:"base"`
	PreferUsbOverWifi *bool  `cty:"prefer_usb_over_wifi" hcl:"prefer_usb_over_wifi"`
}

type OrgDeviceprofileApRadioConfigValue struct {
	AllowRrmDisable  *bool                                  `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain24        *int64                                 `cty:"ant_gain_24" hcl:"ant_gain_24"`
	AntGain5         *int64                                 `cty:"ant_gain_5" hcl:"ant_gain_5"`
	AntGain6         *int64                                 `cty:"ant_gain_6" hcl:"ant_gain_6"`
	AntennaMode      *string                                `cty:"antenna_mode" hcl:"antenna_mode"`
	AntennaSelect    *string                                `cty:"antenna_select" hcl:"antenna_select"`
	Band24           *OrgDeviceprofileApBand24Value         `cty:"band_24" hcl:"band_24"`
	Band24Usage      *string                                `cty:"band_24_usage" hcl:"band_24_usage"`
	Band5            *OrgDeviceprofileApBand5Value          `cty:"band_5" hcl:"band_5"`
	Band5On24Radio   *OrgDeviceprofileApBand5On24RadioValue `cty:"band_5_on_24_radio" hcl:"band_5_on_24_radio"`
	Band6            *OrgDeviceprofileApBand6Value          `cty:"band_6" hcl:"band_6"`
	FullAutomaticRrm *bool                                  `cty:"full_automatic_rrm" hcl:"full_automatic_rrm"`
	IndoorUse        *bool                                  `cty:"indoor_use" hcl:"indoor_use"`
	RrmManaged       *bool                                  `cty:"rrm_managed" hcl:"rrm_managed"`
	ScanningEnabled  *bool                                  `cty:"scanning_enabled" hcl:"scanning_enabled"`
}

type OrgDeviceprofileApBand24Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain" hcl:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode" hcl:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth" hcl:"bandwidth"`
	Channel         *int64  `cty:"channel" hcl:"channel"`
	Channels        []int64 `cty:"channels" hcl:"channels"`
	Disabled        *bool   `cty:"disabled" hcl:"disabled"`
	Power           *int64  `cty:"power" hcl:"power"`
	PowerMax        *int64  `cty:"power_max" hcl:"power_max"`
	PowerMin        *int64  `cty:"power_min" hcl:"power_min"`
	Preamble        *string `cty:"preamble" hcl:"preamble"`
}

type OrgDeviceprofileApBand5Value struct {
	AllowRrmDisable    *bool   `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain            *int64  `cty:"ant_gain" hcl:"ant_gain"`
	AntennaBeamPattern *string `cty:"antenna_beam_pattern" hcl:"antenna_beam_pattern"`
	AntennaMode        *string `cty:"antenna_mode" hcl:"antenna_mode"`
	Bandwidth          *int64  `cty:"bandwidth" hcl:"bandwidth"`
	Channel            *int64  `cty:"channel" hcl:"channel"`
	Channels           []int64 `cty:"channels" hcl:"channels"`
	Disabled           *bool   `cty:"disabled" hcl:"disabled"`
	Power              *int64  `cty:"power" hcl:"power"`
	PowerMax           *int64  `cty:"power_max" hcl:"power_max"`
	PowerMin           *int64  `cty:"power_min" hcl:"power_min"`
	Preamble           *string `cty:"preamble" hcl:"preamble"`
}

type OrgDeviceprofileApBand5On24RadioValue struct {
	AllowRrmDisable    *bool   `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain            *int64  `cty:"ant_gain" hcl:"ant_gain"`
	AntennaBeamPattern *string `cty:"antenna_beam_pattern" hcl:"antenna_beam_pattern"`
	AntennaMode        *string `cty:"antenna_mode" hcl:"antenna_mode"`
	Bandwidth          *int64  `cty:"bandwidth" hcl:"bandwidth"`
	Channel            *int64  `cty:"channel" hcl:"channel"`
	Channels           []int64 `cty:"channels" hcl:"channels"`
	Disabled           *bool   `cty:"disabled" hcl:"disabled"`
	Power              *int64  `cty:"power" hcl:"power"`
	PowerMax           *int64  `cty:"power_max" hcl:"power_max"`
	PowerMin           *int64  `cty:"power_min" hcl:"power_min"`
	Preamble           *string `cty:"preamble" hcl:"preamble"`
}

type OrgDeviceprofileApBand6Value struct {
	AllowRrmDisable    *bool   `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain            *int64  `cty:"ant_gain" hcl:"ant_gain"`
	AntennaBeamPattern *string `cty:"antenna_beam_pattern" hcl:"antenna_beam_pattern"`
	AntennaMode        *string `cty:"antenna_mode" hcl:"antenna_mode"`
	Bandwidth          *int64  `cty:"bandwidth" hcl:"bandwidth"`
	Channel            *int64  `cty:"channel" hcl:"channel"`
	Channels           []int64 `cty:"channels" hcl:"channels"`
	Disabled           *bool   `cty:"disabled" hcl:"disabled"`
	Power              *int64  `cty:"power" hcl:"power"`
	PowerMax           *int64  `cty:"power_max" hcl:"power_max"`
	PowerMin           *int64  `cty:"power_min" hcl:"power_min"`
	Preamble           *string `cty:"preamble" hcl:"preamble"`
	StandardPower      *bool   `cty:"standard_power" hcl:"standard_power"`
}

type OrgDeviceprofileApUplinkPortConfigValue struct {
	Dot1x             *bool `cty:"dot1x" hcl:"dot1x"`
	KeepWlansUpIfDown *bool `cty:"keep_wlans_up_if_down" hcl:"keep_wlans_up_if_down"`
}

type OrgDeviceprofileApUsbConfigValue struct {
	Cacert        *string `cty:"cacert" hcl:"cacert"`
	Channel       *int64  `cty:"channel" hcl:"channel"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	Host          *string `cty:"host" hcl:"host"`
	Port          *int64  `cty:"port" hcl:"port"`
	UsbConfigType *string `cty:"type" hcl:"type"`
	VerifyCert    *bool   `cty:"verify_cert" hcl:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id" hcl:"vlan_id"`
}
