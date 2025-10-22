package provider

type SiteWxruleModel struct {
	Action         string   `hcl:"action"`
	ApplyTags      []string `hcl:"apply_tags"`
	BlockedApps    []string `hcl:"blocked_apps"`
	DstAllowWxtags []string `hcl:"dst_allow_wxtags"`
	DstDenyWxtags  []string `hcl:"dst_deny_wxtags"`
	DstWxtags      []string `hcl:"dst_wxtags"`
	Enabled        *bool    `hcl:"enabled"`
	Order          int64    `hcl:"order"`
	SiteId         string   `hcl:"site_id"`
	SrcWxtags      []string `hcl:"src_wxtags"`
}
