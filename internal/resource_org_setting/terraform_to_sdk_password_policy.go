package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func passwordPolicyTerraformToSdk(d PasswordPolicyValue) *models.OrgSettingPasswordPolicy {
	data := models.OrgSettingPasswordPolicy{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if d.ExpiryInDays.ValueInt64Pointer() != nil {
		data.ExpiryInDays = models.ToPointer(int(d.ExpiryInDays.ValueInt64()))
	}

	if d.MinLength.ValueInt64Pointer() != nil {
		data.MinLength = models.ToPointer(int(d.MinLength.ValueInt64()))
	}

	if d.RequiresSpecialChar.ValueBoolPointer() != nil {
		data.RequiresSpecialChar = d.RequiresSpecialChar.ValueBoolPointer()
	}

	if d.RequiresTwoFactorAuth.ValueBoolPointer() != nil {
		data.RequiresTwoFactorAuth = d.RequiresTwoFactorAuth.ValueBoolPointer()
	}

	return &data
}
