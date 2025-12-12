package provider

import ()

type OrgMxedgeInventoryModel struct {
	Mxedges map[string]OrgMxedgeInventoryMxedgesValue `hcl:"mxedges"`
	OrgId string `hcl:"org_id"`
}

type OrgMxedgeInventoryMxedgesValue struct {
	Magic *string `cty:"claim_code" hcl:"claim_code"`
	SiteId *string `cty:"site_id" hcl:"site_id"`
}

