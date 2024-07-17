package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func passwordPolicyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PasswordPolicyValue) *models.OrgSettingPasswordPolicy {
	data := models.OrgSettingPasswordPolicy{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if d.Freshness.ValueInt64Pointer() != nil {
		data.Freshness = models.ToPointer(int(d.Freshness.ValueInt64()))
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
