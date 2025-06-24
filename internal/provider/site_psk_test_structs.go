package provider

import ()

type SitePskModel struct {
	Email *string `hcl:"email"`
	ExpireTime *int64 `hcl:"expire_time"`
	ExpiryNotificationTime *int64 `hcl:"expiry_notification_time"`
	Mac *string `hcl:"mac"`
	Name string `hcl:"name"`
	Note *string `hcl:"note"`
	NotifyExpiry *bool `hcl:"notify_expiry"`
	NotifyOnCreateOrEdit *bool `hcl:"notify_on_create_or_edit"`
	OldPassphrase *string `hcl:"old_passphrase"`
	Passphrase string `hcl:"passphrase"`
	Role *string `hcl:"role"`
	SiteId string `hcl:"site_id"`
	Ssid string `hcl:"ssid"`
	Usage *string `hcl:"usage"`
	VlanId *string `hcl:"vlan_id"`
}

