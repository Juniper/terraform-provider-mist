package provider

import ()

type OrgNacPortalImageModel struct {
	File        string `hcl:"file"`
	NacportalId string `hcl:"nacportal_id"`
	OrgId       string `hcl:"org_id"`
}
