package provider

type SiteWlanPortalImageModel struct {
	File   string `hcl:"file"`
	SiteId string `hcl:"site_id"`
	WlanId string `hcl:"wlan_id"`
}
