package provider

type OrgSsoRoleModel struct {
	Name       string            `hcl:"name"`
	OrgId      string            `hcl:"org_id"`
	Privileges []PrivilegesValue `hcl:"privileges"`
}

// type PrivilegesValue struct {
// 	Role        string   `cty:"role"`
// 	Scope       string   `cty:"scope"`
// 	SiteId      *string  `cty:"site_id"`
// 	SitegroupId *string  `cty:"sitegroup_id"`
// 	Views       []string `cty:"views"`
// }
