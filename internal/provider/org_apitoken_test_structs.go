package provider

type OrgApitokenModel struct {
	Name       string            `hcl:"name"`
	OrgId      string            `hcl:"org_id"`
	Privileges []PrivilegesValue `hcl:"privileges"`
	SrcIps     []string          `hcl:"src_ips"`
}

type PrivilegesValue struct {
	Role        string  `cty:"role"`
	Scope       string  `cty:"scope"`
	SiteId      *string `cty:"site_id"`
	SitegroupId *string `cty:"sitegroup_id"`
}
