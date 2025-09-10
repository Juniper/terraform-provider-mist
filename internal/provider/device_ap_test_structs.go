package provider

type DeviceApModel struct {
	Aeroscout        *AeroscoutValue                    `hcl:"aeroscout"`
	Airista          *AiristaValue                      `hcl:"airista"`
	BleConfig        *BleConfigValue                    `hcl:"ble_config"`
	Centrak          *CentrakValue                      `hcl:"centrak"`
	ClientBridge     *ClientBridgeValue                 `hcl:"client_bridge"`
	DeviceId         string                             `hcl:"device_id"`
	DisableEth1      *bool                              `hcl:"disable_eth1"`
	DisableEth2      *bool                              `hcl:"disable_eth2"`
	DisableEth3      *bool                              `hcl:"disable_eth3"`
	DisableModule    *bool                              `hcl:"disable_module"`
	EslConfig        *EslConfigValue                    `hcl:"esl_config"`
	FlowControl      *bool                              `hcl:"flow_control"`
	Height           *float64                           `hcl:"height"`
	IpConfig         *IpConfigValue                     `hcl:"ip_config"`
	LacpConfig       *LacpConfigValue                   `hcl:"lacp_config"`
	Led              *LedValue                          `hcl:"led"`
	Locked           *bool                              `hcl:"locked"`
	MapId            *string                            `hcl:"map_id"`
	Mesh             *MeshValue                         `hcl:"mesh"`
	Name             string                             `hcl:"name"`
	Notes            *string                            `hcl:"notes"`
	NtpServers       []string                           `hcl:"ntp_servers"`
	Orientation      *int64                             `hcl:"orientation"`
	PoePassthrough   *bool                              `hcl:"poe_passthrough"`
	PortConfig       map[string]DeviceAPPortConfigValue `hcl:"port_config"`
	PwrConfig        *PwrConfigValue                    `hcl:"pwr_config"`
	RadioConfig      *RadioConfigValue                  `hcl:"radio_config"`
	SiteId           string                             `hcl:"site_id"`
	UplinkPortConfig *UplinkPortConfigValue             `hcl:"uplink_port_config"`
	UsbConfig        *UsbConfigValue                    `hcl:"usb_config"`
	Vars             map[string]string                  `hcl:"vars"`
	X                *float64                           `hcl:"x"`
	Y                *float64                           `hcl:"y"`
}

type AeroscoutValue struct {
	Enabled         *bool   `cty:"enabled" hcl:"enabled"`
	Host            *string `cty:"host" hcl:"host"`
	LocateConnected *bool   `cty:"locate_connected" hcl:"locate_connected"`
	Port            *int64  `cty:"port" hcl:"port"`
}

type AiristaValue struct {
	Enabled *bool   `cty:"enabled" hcl:"enabled"`
	Host    *string `cty:"host" hcl:"host"`
	Port    *int64  `cty:"port" hcl:"port"`
}

type BleConfigValue struct {
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

type CentrakValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type ClientBridgeValue struct {
	Auth    *AuthValue `cty:"auth" hcl:"auth"`
	Enabled *bool      `cty:"enabled" hcl:"enabled"`
	Ssid    *string    `cty:"ssid" hcl:"ssid"`
}

type AuthValue struct {
	Psk      *string `cty:"psk" hcl:"psk"`
	AuthType *string `cty:"type" hcl:"type"`
}

type EslConfigValue struct {
	Cacert        *string `cty:"cacert" hcl:"cacert"`
	Channel       *int64  `cty:"channel" hcl:"channel"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	Host          *string `cty:"host" hcl:"host"`
	Port          *int64  `cty:"port" hcl:"port"`
	EslConfigType *string `cty:"type" hcl:"type"`
	VerifyCert    *bool   `cty:"verify_cert" hcl:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id" hcl:"vlan_id"`
}

type IpConfigValue struct {
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

type LacpConfigValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type LedValue struct {
	Brightness *int64 `cty:"brightness" hcl:"brightness"`
	Enabled    *bool  `cty:"enabled" hcl:"enabled"`
}

type MeshValue struct {
	Bands   []string `cty:"bands" hcl:"bands"`
	Enabled *bool    `cty:"enabled" hcl:"enabled"`
	Group   *int64   `cty:"group" hcl:"group"`
	Role    *string  `cty:"role" hcl:"role"`
}

type DeviceAPPortConfigValue struct {
	Disabled         *bool                      `cty:"disabled" hcl:"disabled"`
	DynamicVlan      *DeviceAPDynamicVlanValue  `cty:"dynamic_vlan" hcl:"dynamic_vlan"`
	EnableMacAuth    *bool                      `cty:"enable_mac_auth" hcl:"enable_mac_auth"`
	Forwarding       *string                    `cty:"forwarding" hcl:"forwarding"`
	MacAuthPreferred *bool                      `cty:"mac_auth_preferred" hcl:"mac_auth_preferred"`
	MacAuthProtocol  *string                    `cty:"mac_auth_protocol" hcl:"mac_auth_protocol"`
	MistNac          *DeviceAPMistNacValue      `cty:"mist_nac" hcl:"mist_nac"`
	MxTunnelId       *string                    `cty:"mx_tunnel_id" hcl:"mx_tunnel_id"`
	MxtunnelName     *string                    `cty:"mxtunnel_name" hcl:"mxtunnel_name"`
	PortAuth         *string                    `cty:"port_auth" hcl:"port_auth"`
	PortVlanId       *int64                     `cty:"port_vlan_id" hcl:"port_vlan_id"`
	RadiusConfig     *DeviceAPRadiusConfigValue `cty:"radius_config" hcl:"radius_config"`
	Radsec           *DeviceAPRadsecValue       `cty:"radsec" hcl:"radsec"`
	VlanId           *int64                     `cty:"vlan_id" hcl:"vlan_id"`
	VlanIds          []int64                    `cty:"vlan_ids" hcl:"vlan_ids"`
	WxtunnelId       *string                    `cty:"wxtunnel_id" hcl:"wxtunnel_id"`
	WxtunnelRemoteId *string                    `cty:"wxtunnel_remote_id" hcl:"wxtunnel_remote_id"`
}

type DeviceAPDynamicVlanValue struct {
	DefaultVlanId   *int64            `cty:"default_vlan_id" hcl:"default_vlan_id"`
	Enabled         *bool             `cty:"enabled" hcl:"enabled"`
	DynamicVlanType *string           `cty:"type" hcl:"type"`
	Vlans           map[string]string `cty:"vlans" hcl:"vlans"`
}

type DeviceAPMistNacValue struct {
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

type DeviceAPRadiusConfigValue struct {
	AcctInterimInterval *int64                     `cty:"acct_interim_interval" hcl:"acct_interim_interval"`
	AcctServers         []DeviceAPAcctServersValue `cty:"acct_servers" hcl:"acct_servers"`
	AuthServers         []DeviceAPAuthServersValue `cty:"auth_servers" hcl:"auth_servers"`
	AuthServersRetries  *int64                     `cty:"auth_servers_retries" hcl:"auth_servers_retries"`
	AuthServersTimeout  *int64                     `cty:"auth_servers_timeout" hcl:"auth_servers_timeout"`
	CoaEnabled          *bool                      `cty:"coa_enabled" hcl:"coa_enabled"`
	CoaPort             *int64                     `cty:"coa_port" hcl:"coa_port"`
	Network             *string                    `cty:"network" hcl:"network"`
	SourceIp            *string                    `cty:"source_ip" hcl:"source_ip"`
}

type DeviceAPAcctServersValue struct {
	Host           string  `cty:"host" hcl:"host"`
	KeywrapEnabled *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat  *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek     *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack    *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port           *string `cty:"port" hcl:"port"`
	Secret         string  `cty:"secret" hcl:"secret"`
}

type DeviceAPAuthServersValue struct {
	Host                        string  `cty:"host" hcl:"host"`
	KeywrapEnabled              *bool   `cty:"keywrap_enabled" hcl:"keywrap_enabled"`
	KeywrapFormat               *string `cty:"keywrap_format" hcl:"keywrap_format"`
	KeywrapKek                  *string `cty:"keywrap_kek" hcl:"keywrap_kek"`
	KeywrapMack                 *string `cty:"keywrap_mack" hcl:"keywrap_mack"`
	Port                        *string `cty:"port" hcl:"port"`
	RequireMessageAuthenticator *bool   `cty:"require_message_authenticator" hcl:"require_message_authenticator"`
	Secret                      string  `cty:"secret" hcl:"secret"`
}

type DeviceAPRadsecValue struct {
	CoaEnabled    *bool                  `cty:"coa_enabled" hcl:"coa_enabled"`
	Enabled       *bool                  `cty:"enabled" hcl:"enabled"`
	IdleTimeout   *string                `cty:"idle_timeout" hcl:"idle_timeout"`
	MxclusterIds  []string               `cty:"mxcluster_ids" hcl:"mxcluster_ids"`
	ProxyHosts    []string               `cty:"proxy_hosts" hcl:"proxy_hosts"`
	ServerName    *string                `cty:"server_name" hcl:"server_name"`
	Servers       []DeviceAPServersValue `cty:"servers" hcl:"servers"`
	UseMxedge     *bool                  `cty:"use_mxedge" hcl:"use_mxedge"`
	UseSiteMxedge *bool                  `cty:"use_site_mxedge" hcl:"use_site_mxedge"`
}

type DeviceAPServersValue struct {
	Host *string `cty:"host" hcl:"host"`
	Port *int64  `cty:"port" hcl:"port"`
}

type PwrConfigValue struct {
	Base              *int64 `cty:"base" hcl:"base"`
	PreferUsbOverWifi *bool  `cty:"prefer_usb_over_wifi" hcl:"prefer_usb_over_wifi"`
}

type RadioConfigValue struct {
	AllowRrmDisable  *bool                `cty:"allow_rrm_disable" hcl:"allow_rrm_disable"`
	AntGain24        *int64               `cty:"ant_gain_24" hcl:"ant_gain_24"`
	AntGain5         *int64               `cty:"ant_gain_5" hcl:"ant_gain_5"`
	AntGain6         *int64               `cty:"ant_gain_6" hcl:"ant_gain_6"`
	AntennaMode      *string              `cty:"antenna_mode" hcl:"antenna_mode"`
	Band24           *Band24Value         `cty:"band_24" hcl:"band_24"`
	Band24Usage      *string              `cty:"band_24_usage" hcl:"band_24_usage"`
	Band5            *Band5Value          `cty:"band_5" hcl:"band_5"`
	Band5On24Radio   *Band5On24RadioValue `cty:"band_5_on_24_radio" hcl:"band_5_on_24_radio"`
	Band6            *Band6Value          `cty:"band_6" hcl:"band_6"`
	FullAutomaticRrm *bool                `cty:"full_automatic_rrm" hcl:"full_automatic_rrm"`
	IndoorUse        *bool                `cty:"indoor_use" hcl:"indoor_use"`
	ScanningEnabled  *bool                `cty:"scanning_enabled" hcl:"scanning_enabled"`
}

type Band24Value struct {
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

type Band5Value struct {
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

type Band5On24RadioValue struct {
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

type Band6Value struct {
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
	StandardPower   *bool   `cty:"standard_power" hcl:"standard_power"`
}

type UplinkPortConfigValue struct {
	Dot1x             *bool `cty:"dot1x" hcl:"dot1x"`
	KeepWlansUpIfDown *bool `cty:"keep_wlans_up_if_down" hcl:"keep_wlans_up_if_down"`
}

type UsbConfigValue struct {
	Cacert        *string `cty:"cacert" hcl:"cacert"`
	Channel       *int64  `cty:"channel" hcl:"channel"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	Host          *string `cty:"host" hcl:"host"`
	Port          *int64  `cty:"port" hcl:"port"`
	UsbConfigType *string `cty:"type" hcl:"type"`
	VerifyCert    *bool   `cty:"verify_cert" hcl:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id" hcl:"vlan_id"`
}
