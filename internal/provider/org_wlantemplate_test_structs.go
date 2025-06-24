package provider

type OrgWlantemplateModel struct {
	Applies               *AppliesValue    `hcl:"applies"`
	DeviceprofileIds      []string         `hcl:"deviceprofile_ids"`
	Exceptions            *ExceptionsValue `hcl:"exceptions"`
	FilterByDeviceprofile *bool            `hcl:"filter_by_deviceprofile"`
	Name                  string           `hcl:"name"`
	OrgId                 string           `hcl:"org_id"`
}

type AppliesValue struct {
	OrgId        *string  `cty:"org_id"`
	SiteIds      []string `cty:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids"`
}

type ExceptionsValue struct {
	SiteIds      []string `cty:"site_ids"`
	SitegroupIds []string `cty:"sitegroup_ids"`
}
