package provider

type OrgVpnModel struct {
	Name  string                      `hcl:"name"`
	OrgId string                      `hcl:"org_id"`
	Paths map[string]OrgVpnPathsValue `hcl:"paths"`
}

type OrgVpnPathsValue struct {
	BfdProfile *string `cty:"bfd_profile"`
	Ip         *string `cty:"ip"`
	Pod        int64   `cty:"pod"`
}
