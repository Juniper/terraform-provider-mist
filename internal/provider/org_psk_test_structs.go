package provider

import ()

type OrgPskModel struct {
	Email *string `hcl:"email"`
	ExpireTime int64 `hcl:"expire_time"`
	ExpiryNotificationTime int64 `hcl:"expiry_notification_time"`
	Mac *string `hcl:"mac"`
	Macs []string `hcl:"macs"`
	MaxUsage int64 `hcl:"max_usage"`
	Name string `hcl:"name"`
	Note *string `hcl:"note"`
	NotifyExpiry *bool `hcl:"notify_expiry"`
	NotifyOnCreateOrEdit *bool `hcl:"notify_on_create_or_edit"`
	OldPassphrase *string `hcl:"old_passphrase"`
	OrgId string `hcl:"org_id"`
	Passphrase string `hcl:"passphrase"`
	Role *string `hcl:"role"`
	Ssid string `hcl:"ssid"`
	Usage *string `hcl:"usage"`
	VlanId *string `hcl:"vlan_id"`
}

