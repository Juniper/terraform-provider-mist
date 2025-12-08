package provider

import ()

type OrgMxedgeInventoryModel struct {
	Mxedges map[string]OrgMxedgeInventoryMxedgesValue `hcl:"mxedges"`
	OrgId string `hcl:"org_id"`
}

type OrgMxedgeInventoryMxedgesValue struct {
	Magic *string `cty:"magic" hcl:"magic"`
	SiteId *string `cty:"site_id" hcl:"site_id"`
}

