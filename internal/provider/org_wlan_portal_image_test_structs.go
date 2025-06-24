package provider

type OrgWlanPortalImageModel struct {
	File   string `hcl:"file"`
	OrgId  string `hcl:"org_id"`
	WlanId string `hcl:"wlan_id"`
}
