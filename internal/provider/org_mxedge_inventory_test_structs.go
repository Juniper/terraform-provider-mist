package provider

type OrgMxedgeInventoryModel struct {
	Mxedges map[string]OrgMxedgeInventoryMxedgesValue `hcl:"mxedges"`
	OrgId   string                                    `hcl:"org_id"`
}

type OrgMxedgeInventoryMxedgesValue struct {
	SiteId *string `cty:"site_id" hcl:"site_id"`
}
