package provider

type OrgAlarmtemplateModel struct {
	Delivery DeliveryValue                         `hcl:"delivery"`
	Name     string                                `hcl:"name"`
	OrgId    string                                `hcl:"org_id"`
	Rules    map[string]OrgAlarmtemplateRulesValue `hcl:"rules"`
}

type DeliveryValue struct {
	AdditionalEmails []string `cty:"additional_emails"`
	Enabled          bool     `cty:"enabled"`
	ToOrgAdmins      *bool    `cty:"to_org_admins"`
	ToSiteAdmins     *bool    `cty:"to_site_admins"`
}

type OrgAlarmtemplateRulesValue struct {
	Delivery *DeliveryValue `cty:"delivery"`
	Enabled  bool           `cty:"enabled"`
}
