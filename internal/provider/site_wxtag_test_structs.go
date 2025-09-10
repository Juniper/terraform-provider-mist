package provider

type SiteWxtagModel struct {
	Mac    *string              `hcl:"mac"`
	Match  *string              `hcl:"match"`
	Name   string               `hcl:"name"`
	Op     *string              `hcl:"op"`
	SiteId string               `hcl:"site_id"`
	Specs  []OrgWxtagSpecsValue `hcl:"specs"`
	Type   string               `hcl:"type"`
	Values []string             `hcl:"values"`
	VlanId *string              `hcl:"vlan_id"`
}
