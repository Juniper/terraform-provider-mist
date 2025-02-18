package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func celonaTerraformToSdk(d CelonaValue) *models.OrgSettingCelona {
	data := models.OrgSettingCelona{}

	if d.ApiKey.ValueStringPointer() != nil {
		data.ApiKey = d.ApiKey.ValueStringPointer()
	}

	if d.ApiPrefix.ValueStringPointer() != nil {
		data.ApiPrefix = d.ApiPrefix.ValueStringPointer()
	}
	return &data
}
