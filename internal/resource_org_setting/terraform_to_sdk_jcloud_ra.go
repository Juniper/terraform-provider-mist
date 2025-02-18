package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func jcloudRaTerraformToSdk(d JcloudRaValue) *models.OrgSettingJcloudRa {
	data := models.OrgSettingJcloudRa{}

	if d.OrgApitoken.ValueStringPointer() != nil {
		data.OrgApitoken = d.OrgApitoken.ValueStringPointer()
	}

	if d.OrgApitokenName.ValueStringPointer() != nil {
		data.OrgApitokenName = d.OrgApitokenName.ValueStringPointer()
	}

	if d.OrgId.ValueStringPointer() != nil {
		data.OrgId = d.OrgId.ValueStringPointer()
	}

	return &data
}
