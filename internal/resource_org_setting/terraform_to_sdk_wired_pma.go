package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wiredPmaTerraformToSdk(d WiredPmaValue) *models.OrgSettingWiredPma {
	data := models.OrgSettingWiredPma{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
