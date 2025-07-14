package provider

type DeviceApModel struct {
	Aeroscout        *AeroscoutValue        `hcl:"aeroscout"`
	BleConfig        *BleConfigValue        `hcl:"ble_config"`
	Centrak          *CentrakValue          `hcl:"centrak"`
	ClientBridge     *ClientBridgeValue     `hcl:"client_bridge"`
	DeviceId         string                 `hcl:"device_id"`
	DisableEth1      *bool                  `hcl:"disable_eth1"`
	DisableEth2      *bool                  `hcl:"disable_eth2"`
	DisableEth3      *bool                  `hcl:"disable_eth3"`
	DisableModule    *bool                  `hcl:"disable_module"`
	EslConfig        *EslConfigValue        `hcl:"esl_config"`
	FlowControl      *bool                  `hcl:"flow_control"`
	Height           *float64               `hcl:"height"`
	IpConfig         *IpConfigValue         `hcl:"ip_config"`
	Led              *LedValue              `hcl:"led"`
	Locked           *bool                  `hcl:"locked"`
	MapId            *string                `hcl:"map_id"`
	Mesh             *MeshValue             `hcl:"mesh"`
	Name             string                 `hcl:"name"`
	Notes            *string                `hcl:"notes"`
	NtpServers       []string               `hcl:"ntp_servers"`
	Orientation      *int64                 `hcl:"orientation"`
	PoePassthrough   *bool                  `hcl:"poe_passthrough"`
	PwrConfig        *PwrConfigValue        `hcl:"pwr_config"`
	RadioConfig      *RadioConfigValue      `hcl:"radio_config"`
	SiteId           string                 `hcl:"site_id"`
	UplinkPortConfig *UplinkPortConfigValue `hcl:"uplink_port_config"`
	UsbConfig        *UsbConfigValue        `hcl:"usb_config"`
	Vars             map[string]string      `hcl:"vars"`
	X                *float64               `hcl:"x"`
	Y                *float64               `hcl:"y"`
}

type AeroscoutValue struct {
	Enabled         *bool   `cty:"enabled"`
	Host            *string `cty:"host"`
	LocateConnected *bool   `cty:"locate_connected"`
}

type BleConfigValue struct {
	BeaconEnabled           *bool   `cty:"beacon_enabled"`
	BeaconRate              *int64  `cty:"beacon_rate"`
	BeaconRateMode          *string `cty:"beacon_rate_mode"`
	BeamDisabled            []int64 `cty:"beam_disabled"`
	CustomBlePacketEnabled  *bool   `cty:"custom_ble_packet_enabled"`
	CustomBlePacketFrame    *string `cty:"custom_ble_packet_frame"`
	CustomBlePacketFreqMsec *int64  `cty:"custom_ble_packet_freq_msec"`
	EddystoneUidAdvPower    *int64  `cty:"eddystone_uid_adv_power"`
	EddystoneUidBeams       *string `cty:"eddystone_uid_beams"`
	EddystoneUidEnabled     *bool   `cty:"eddystone_uid_enabled"`
	EddystoneUidFreqMsec    *int64  `cty:"eddystone_uid_freq_msec"`
	EddystoneUidInstance    *string `cty:"eddystone_uid_instance"`
	EddystoneUidNamespace   *string `cty:"eddystone_uid_namespace"`
	EddystoneUrlAdvPower    *int64  `cty:"eddystone_url_adv_power"`
	EddystoneUrlBeams       *string `cty:"eddystone_url_beams"`
	EddystoneUrlEnabled     *bool   `cty:"eddystone_url_enabled"`
	EddystoneUrlFreqMsec    *int64  `cty:"eddystone_url_freq_msec"`
	EddystoneUrlUrl         *string `cty:"eddystone_url_url"`
	IbeaconAdvPower         *int64  `cty:"ibeacon_adv_power"`
	IbeaconBeams            *string `cty:"ibeacon_beams"`
	IbeaconEnabled          *bool   `cty:"ibeacon_enabled"`
	IbeaconFreqMsec         *int64  `cty:"ibeacon_freq_msec"`
	IbeaconMajor            *int64  `cty:"ibeacon_major"`
	IbeaconMinor            *int64  `cty:"ibeacon_minor"`
	IbeaconUuid             *string `cty:"ibeacon_uuid"`
	Power                   *int64  `cty:"power"`
	PowerMode               *string `cty:"power_mode"`
}

type CentrakValue struct {
	Enabled *bool `cty:"enabled"`
}

type ClientBridgeValue struct {
	Auth    *AuthValue `cty:"auth"`
	Enabled *bool      `cty:"enabled"`
	Ssid    *string    `cty:"ssid"`
}

type AuthValue struct {
	Psk      *string `cty:"psk"`
	AuthType *string `cty:"type"`
}

type EslConfigValue struct {
	Cacert        *string `cty:"cacert"`
	Channel       *int64  `cty:"channel"`
	Enabled       *bool   `cty:"enabled"`
	Host          *string `cty:"host"`
	Port          *int64  `cty:"port"`
	EslConfigType *string `cty:"type"`
	VerifyCert    *bool   `cty:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id"`
}

type IpConfigValue struct {
	Dns          []string `cty:"dns"`
	DnsSuffix    []string `cty:"dns_suffix"`
	Gateway      *string  `cty:"gateway"`
	Gateway6     *string  `cty:"gateway6"`
	Ip           *string  `cty:"ip"`
	Ip6          *string  `cty:"ip6"`
	Mtu          *int64   `cty:"mtu"`
	Netmask      *string  `cty:"netmask"`
	Netmask6     *string  `cty:"netmask6"`
	IpConfigType *string  `cty:"type"`
	Type6        *string  `cty:"type6"`
	VlanId       *int64   `cty:"vlan_id"`
}

type LedValue struct {
	Brightness *int64 `cty:"brightness"`
	Enabled    *bool  `cty:"enabled"`
}

type MeshValue struct {
	Enabled *bool   `cty:"enabled"`
	Group   *int64  `cty:"group"`
	Role    *string `cty:"role"`
}

type PwrConfigValue struct {
	Base              *int64 `cty:"base"`
	PreferUsbOverWifi *bool  `cty:"prefer_usb_over_wifi"`
}

type RadioConfigValue struct {
	AllowRrmDisable *bool                `cty:"allow_rrm_disable"`
	AntGain24       *int64               `cty:"ant_gain_24"`
	AntGain5        *int64               `cty:"ant_gain_5"`
	AntGain6        *int64               `cty:"ant_gain_6"`
	AntennaMode     *string              `cty:"antenna_mode"`
	Band24          *Band24Value         `cty:"band_24"`
	Band24Usage     *string              `cty:"band_24_usage"`
	Band5           *Band5Value          `cty:"band_5"`
	Band5On24Radio  *Band5On24RadioValue `cty:"band_5_on_24_radio"`
	Band6           *Band6Value          `cty:"band_6"`
	IndoorUse       *bool                `cty:"indoor_use"`
	ScanningEnabled *bool                `cty:"scanning_enabled"`
}

type Band24Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channel         *int64  `cty:"channel"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type Band5Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channel         *int64  `cty:"channel"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type Band5On24RadioValue struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channel         *int64  `cty:"channel"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
}

type Band6Value struct {
	AllowRrmDisable *bool   `cty:"allow_rrm_disable"`
	AntGain         *int64  `cty:"ant_gain"`
	AntennaMode     *string `cty:"antenna_mode"`
	Bandwidth       *int64  `cty:"bandwidth"`
	Channel         *int64  `cty:"channel"`
	Channels        []int64 `cty:"channels"`
	Disabled        *bool   `cty:"disabled"`
	Power           *int64  `cty:"power"`
	PowerMax        *int64  `cty:"power_max"`
	PowerMin        *int64  `cty:"power_min"`
	Preamble        *string `cty:"preamble"`
	StandardPower   *bool   `cty:"standard_power"`
}

type UplinkPortConfigValue struct {
	Dot1x             *bool `cty:"dot1x"`
	KeepWlansUpIfDown *bool `cty:"keep_wlans_up_if_down"`
}

type UsbConfigValue struct {
	Cacert        *string `cty:"cacert"`
	Channel       *int64  `cty:"channel"`
	Enabled       *bool   `cty:"enabled"`
	Host          *string `cty:"host"`
	Port          *int64  `cty:"port"`
	UsbConfigType *string `cty:"type"`
	VerifyCert    *bool   `cty:"verify_cert"`
	VlanId        *int64  `cty:"vlan_id"`
}
