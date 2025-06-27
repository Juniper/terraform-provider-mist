package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mxEdgeMgmtTerraformToSdk(d MxedgeMgmtValue) *models.MxedgeMgmt {
	data := models.MxedgeMgmt{}

	if d.ConfigAutoRevert.ValueBoolPointer() != nil {
		data.ConfigAutoRevert = d.ConfigAutoRevert.ValueBoolPointer()
	}

	if d.FipsEnabled.ValueBoolPointer() != nil {
		data.FipsEnabled = d.FipsEnabled.ValueBoolPointer()
	}

	if d.MistPassword.ValueStringPointer() != nil {
		data.MistPassword = d.MistPassword.ValueStringPointer()
	}

	if d.OobIpType.ValueStringPointer() != nil {
		data.OobIpType = (*models.MxedgeMgmtOobIpTypeEnum)(d.OobIpType.ValueStringPointer())
	}

	if d.OobIpType6.ValueStringPointer() != nil {
		data.OobIpType6 = (*models.MxedgeMgmtOobIpType6Enum)(d.OobIpType6.ValueStringPointer())
	}

	if d.RootPassword.ValueStringPointer() != nil {
		data.RootPassword = d.RootPassword.ValueStringPointer()
	}

	return &data
}
