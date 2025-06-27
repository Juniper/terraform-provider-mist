package provider

type OrgSsoRoleModel struct {
	Name       string                      `hcl:"name"`
	OrgId      string                      `hcl:"org_id"`
	Privileges []OrgSsoRolePrivilegesValue `hcl:"privileges"`
}

type OrgSsoRolePrivilegesValue struct {
	Role        string   `hcl:"role" cty:"role"`
	Scope       string   `hcl:"scope" cty:"scope"`
	SiteId      *string  `hcl:"site_id" cty:"site_id"`
	SitegroupId *string  `hcl:"sitegroup_id" cty:"sitegroup_id"`
	Views       []string `hcl:"views" cty:"views"`
}
