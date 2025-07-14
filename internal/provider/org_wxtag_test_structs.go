package provider

type OrgWxtagModel struct {
	Mac    *string              `hcl:"mac"`
	Match  *string              `hcl:"match"`
	Name   string               `hcl:"name"`
	Op     *string              `hcl:"op"`
	OrgId  string               `hcl:"org_id"`
	Specs  []OrgWxtagSpecsValue `hcl:"specs"`
	Type   string               `hcl:"type"`
	Values []string             `hcl:"values"`
	VlanId *string              `hcl:"vlan_id"`
}

type OrgWxtagSpecsValue struct {
	PortRange *string  `cty:"port_range"`
	Protocol  *string  `cty:"protocol"`
	Subnets   []string `cty:"subnets"`
}
