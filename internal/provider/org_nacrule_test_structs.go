package provider

import ()

type OrgNacruleModel struct {
	Action string `hcl:"action"`
	ApplyTags []string `hcl:"apply_tags"`
	Enabled *bool `hcl:"enabled"`
	Matching *MatchingValue `hcl:"matching"`
	Name string `hcl:"name"`
	NotMatching *NotMatchingValue `hcl:"not_matching"`
	Order int64 `hcl:"order"`
	OrgId string `hcl:"org_id"`
}

type MatchingValue struct {
	AuthType *string `cty:"auth_type"`
	Nactags []string `cty:"nactags"`
	PortTypes []string `cty:"port_types"`
	SiteIds []string `cty:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids"`
	Vendor []string `cty:"vendor"`
}

type NotMatchingValue struct {
	AuthType *string `cty:"auth_type"`
	Nactags []string `cty:"nactags"`
	PortTypes []string `cty:"port_types"`
	SiteIds []string `cty:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids"`
	Vendor []string `cty:"vendor"`
}

