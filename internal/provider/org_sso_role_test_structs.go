package provider

type OrgSsoRoleModel struct {
	Name       string                      `hcl:"name"`
	OrgId      string                      `hcl:"org_id"`
	Privileges []OrgSsoRolePrivilegesValue `hcl:"privileges"`
}

type OrgSsoRolePrivilegesValue struct {
	Role        string   `cty:"role" hcl:"role"`
	Scope       string   `cty:"scope" hcl:"scope"`
	SiteId      *string  `cty:"site_id" hcl:"site_id"`
	SitegroupId *string  `cty:"sitegroup_id" hcl:"sitegroup_id"`
	View        *string  `cty:"view" hcl:"view"`
	Views       []string `cty:"views" hcl:"views"`
}
