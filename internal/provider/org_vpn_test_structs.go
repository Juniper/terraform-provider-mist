package provider

type OrgVpnModel struct {
	Name  string                      `hcl:"name"`
	OrgId string                      `hcl:"org_id"`
	Paths map[string]OrgVpnPathsValue `hcl:"paths"`
}

type OrgVpnPathsValue struct {
	BfdProfile *string `hcl:"bfd_profile" cty:"bfd_profile"`
	Ip         *string `hcl:"ip" cty:"ip"`
	Pod        *int64  `hcl:"pod" cty:"pod"`
}
