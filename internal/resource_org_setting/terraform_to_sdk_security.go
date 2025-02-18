package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func securityTerraformToSdk(d SecurityValue) *models.OrgSettingSecurity {
	data := models.OrgSettingSecurity{}

	if d.DisableLocalSsh.ValueBoolPointer() != nil {
		data.DisableLocalSsh = d.DisableLocalSsh.ValueBoolPointer()
	}

	if d.FipsZeroizePassword.ValueStringPointer() != nil {
		data.FipsZeroizePassword = d.FipsZeroizePassword.ValueStringPointer()
	}

	if d.LimitSshAccess.ValueBoolPointer() != nil {
		data.LimitSshAccess = d.LimitSshAccess.ValueBoolPointer()
	}

	return &data
}
