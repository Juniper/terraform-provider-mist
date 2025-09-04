package provider

type OrgWxtagModel struct {
	Id     *string              `hcl:"id"`
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
	PortRange *string  `hcl:"port_range" cty:"port_range"`
	Protocol  *string  `hcl:"protocol" cty:"protocol"`
	Subnets   []string `hcl:"subnets" cty:"subnets"`
}
