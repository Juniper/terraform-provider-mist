package provider

type SiteModel struct {
	Address           string       `hcl:"address"`
	AlarmtemplateId   *string      `hcl:"alarmtemplate_id"`
	AptemplateId      *string      `hcl:"aptemplate_id"`
	CountryCode       *string      `hcl:"country_code"`
	GatewaytemplateId *string      `hcl:"gatewaytemplate_id"`
	Latlng            *LatlngValue `hcl:"latlng"`
	Name              string       `hcl:"name"`
	NetworktemplateId *string      `hcl:"networktemplate_id"`
	Notes             *string      `hcl:"notes"`
	OrgId             string       `hcl:"org_id"`
	RftemplateId      *string      `hcl:"rftemplate_id"`
	SecpolicyId       *string      `hcl:"secpolicy_id"`
	SitegroupIds      []string     `hcl:"sitegroup_ids"`
	SitetemplateId    *string      `hcl:"sitetemplate_id"`
	Timezone          *string      `hcl:"timezone"`
}
