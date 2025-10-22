package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wirelessPmaTerraformToSdk(d WirelessPmaValue) *models.OrgSettingWirelessPma {
	data := models.OrgSettingWirelessPma{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
