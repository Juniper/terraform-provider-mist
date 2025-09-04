package provider

type OrgWlantemplateModel struct {
	Applies               *AppliesValue    `hcl:"applies"`
	DeviceprofileIds      []string         `hcl:"deviceprofile_ids"`
	Exceptions            *ExceptionsValue `hcl:"exceptions"`
	FilterByDeviceprofile *bool            `hcl:"filter_by_deviceprofile"`
	Id                    *string          `hcl:"id"`
	Name                  string           `hcl:"name"`
	OrgId                 string           `hcl:"org_id"`
}

type AppliesValue struct {
	OrgId        *string  `hcl:"org_id" cty:"org_id"`
	SiteIds      []string `hcl:"site_ids" cty:"site_ids"`
	SitegroupIds []string `hcl:"sitegroup_ids" cty:"sitegroup_ids"`
}

type ExceptionsValue struct {
	SiteIds      []string `hcl:"site_ids" cty:"site_ids"`
	SitegroupIds []string `hcl:"sitegroup_ids" cty:"sitegroup_ids"`
}
