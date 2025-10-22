package provider

type OrgAlarmtemplateModel struct {
	Delivery OrgAlarmtemplateDeliveryValue         `hcl:"delivery"`
	Name     string                                `hcl:"name"`
	OrgId    string                                `hcl:"org_id"`
	Rules    map[string]OrgAlarmtemplateRulesValue `hcl:"rules"`
}

type OrgAlarmtemplateDeliveryValue struct {
	AdditionalEmails []string `cty:"additional_emails" hcl:"additional_emails"`
	Enabled          bool     `cty:"enabled" hcl:"enabled"`
	ToOrgAdmins      *bool    `cty:"to_org_admins" hcl:"to_org_admins"`
	ToSiteAdmins     *bool    `cty:"to_site_admins" hcl:"to_site_admins"`
}

type OrgAlarmtemplateRulesValue struct {
	Delivery *OrgAlarmtemplateDeliveryValue `cty:"delivery" hcl:"delivery"`
	Enabled  bool                           `cty:"enabled" hcl:"enabled"`
}
