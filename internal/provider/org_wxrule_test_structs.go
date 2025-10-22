package provider

type OrgWxruleModel struct {
	Action         string   `hcl:"action"`
	ApplyTags      []string `hcl:"apply_tags"`
	BlockedApps    []string `hcl:"blocked_apps"`
	DstAllowWxtags []string `hcl:"dst_allow_wxtags"`
	DstDenyWxtags  []string `hcl:"dst_deny_wxtags"`
	DstWxtags      []string `hcl:"dst_wxtags"`
	Enabled        *bool    `hcl:"enabled"`
	Id             *string  `hcl:"id"`
	Order          int64    `hcl:"order"`
	OrgId          string   `hcl:"org_id"`
	SrcWxtags      []string `hcl:"src_wxtags"`
	TemplateId     string   `hcl:"template_id"`
}
