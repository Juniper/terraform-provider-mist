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
	PortRange *string  `cty:"port_range" hcl:"port_range"`
	Protocol  *string  `cty:"protocol" hcl:"protocol"`
	Subnets   []string `cty:"subnets" hcl:"subnets"`
}
