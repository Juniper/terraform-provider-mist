package provider

type OrgApitokenModel struct {
	Name       string                       `hcl:"name"`
	OrgId      string                       `hcl:"org_id"`
	Privileges []OrgApitokenPrivilegesValue `hcl:"privileges"`
	SrcIps     []string                     `hcl:"src_ips"`
}

type OrgApitokenPrivilegesValue struct {
	Role        string  `cty:"role" hcl:"role"`
	Scope       string  `cty:"scope" hcl:"scope"`
	SiteId      *string `cty:"site_id" hcl:"site_id"`
	SitegroupId *string `cty:"sitegroup_id" hcl:"sitegroup_id"`
}
