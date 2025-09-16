package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wanPmaTerraformToSdk(d WanPmaValue) *models.OrgSettingWanPma {
	data := models.OrgSettingWanPma{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
