package provider

type UpgradeDeviceModel struct {
	DeviceId                   string `hcl:"device_id"`
	Reboot                     *bool  `hcl:"reboot"`
	RebootAt                   int64  `hcl:"reboot_at"`
	SiteId                     string `hcl:"site_id"`
	Snapshot                   *bool  `hcl:"snapshot"`
	StartTime                  int64  `hcl:"start_time"`
	SyncUpgrade                *bool  `hcl:"sync_upgrade"`
	SyncUpgradeRefreshInterval int64  `hcl:"sync_upgrade_refresh_interval"`
	SyncUpgradeStartTimeout    int64  `hcl:"sync_upgrade_start_timeout"`
	SyncUpgradeTimeout         int64  `hcl:"sync_upgrade_timeout"`
	TargetVersion              string `hcl:"target_version"`
}
