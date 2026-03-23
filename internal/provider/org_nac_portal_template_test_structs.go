package provider

import ()

type OrgNacPortalTemplateModel struct {
	Alignment *string `hcl:"alignment"`
	Color *string `hcl:"color"`
	Logo *string `hcl:"logo"`
	NacportalId string `hcl:"nacportal_id"`
	OrgId string `hcl:"org_id"`
	PoweredBy *bool `hcl:"powered_by"`
}

