package provider

type OrgDeviceprofileApModel struct {
	Aeroscout        *AeroscoutValue        `hcl:"aeroscout"`
	BleConfig        *BleConfigValue        `hcl:"ble_config"`
	DisableEth1      *bool                  `hcl:"disable_eth1"`
	DisableEth2      *bool                  `hcl:"disable_eth2"`
	DisableEth3      *bool                  `hcl:"disable_eth3"`
	DisableModule    *bool                  `hcl:"disable_module"`
	EslConfig        *EslConfigValue        `hcl:"esl_config"`
	IpConfig         *IpConfigValue         `hcl:"ip_config"`
	Led              *LedValue              `hcl:"led"`
	Mesh             *MeshValue             `hcl:"mesh"`
	Name             string                 `hcl:"name"`
	NtpServers       []string               `hcl:"ntp_servers"`
	OrgId            string                 `hcl:"org_id"`
	PoePassthrough   *bool                  `hcl:"poe_passthrough"`
	PwrConfig        *PwrConfigValue        `hcl:"pwr_config"`
	RadioConfig      *RadioConfigValue      `hcl:"radio_config"`
	SiteId           *string                `hcl:"site_id"`
	UplinkPortConfig *UplinkPortConfigValue `hcl:"uplink_port_config"`
	UsbConfig        *UsbConfigValue        `hcl:"usb_config"`
	Vars             map[string]string      `hcl:"vars"`
}
