package provider

type OrgNacruleModel struct {
	Action         string            `hcl:"action"`
	ApplyTags      []string          `hcl:"apply_tags"`
	Enabled        *bool             `hcl:"enabled"`
	GuestAuthState *string           `hcl:"guest_auth_state"`
	Matching       *MatchingValue    `hcl:"matching"`
	Name           string            `hcl:"name"`
	NotMatching    *NotMatchingValue `hcl:"not_matching"`
	Order          int64             `hcl:"order"`
	OrgId          string            `hcl:"org_id"`
}

type MatchingValue struct {
	AuthType     *string  `cty:"auth_type" hcl:"auth_type"`
	Family       []string `cty:"family" hcl:"family"`
	Mfg          []string `cty:"mfg" hcl:"mfg"`
	Model        []string `cty:"model" hcl:"model"`
	Nactags      []string `cty:"nactags" hcl:"nactags"`
	OsType       []string `cty:"os_type" hcl:"os_type"`
	PortTypes    []string `cty:"port_types" hcl:"port_types"`
	SiteIds      []string `cty:"site_ids" hcl:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids" hcl:"sitegroup_ids"`
	Vendor       []string `cty:"vendor" hcl:"vendor"`
}

type NotMatchingValue struct {
	AuthType     *string  `cty:"auth_type" hcl:"auth_type"`
	Family       []string `cty:"family" hcl:"family"`
	Mfg          []string `cty:"mfg" hcl:"mfg"`
	Model        []string `cty:"model" hcl:"model"`
	Nactags      []string `cty:"nactags" hcl:"nactags"`
	OsType       []string `cty:"os_type" hcl:"os_type"`
	PortTypes    []string `cty:"port_types" hcl:"port_types"`
	SiteIds      []string `cty:"site_ids" hcl:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids" hcl:"sitegroup_ids"`
	Vendor       []string `cty:"vendor" hcl:"vendor"`
}
