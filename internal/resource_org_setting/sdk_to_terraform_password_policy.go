package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func passwordPolicySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingPasswordPolicy) PasswordPolicyValue {

	var enabled basetypes.BoolValue
	var expiryInDays basetypes.Int64Value
	var minLength basetypes.Int64Value
	var requiresSpecialChar basetypes.BoolValue
	var requiresTwoFactorAuth basetypes.BoolValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.ExpiryInDays != nil {
		expiryInDays = types.Int64Value(int64(*d.ExpiryInDays))
	}
	if d.MinLength != nil {
		minLength = types.Int64Value(int64(*d.MinLength))
	}
	if d.RequiresSpecialChar != nil {
		requiresSpecialChar = types.BoolValue(*d.RequiresSpecialChar)
	}
	if d.RequiresTwoFactorAuth != nil {
		requiresTwoFactorAuth = types.BoolValue(*d.RequiresTwoFactorAuth)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":                  enabled,
		"expiry_in_days":           expiryInDays,
		"min_length":               minLength,
		"requires_special_char":    requiresSpecialChar,
		"requires_two_factor_auth": requiresTwoFactorAuth,
	}
	data, e := NewPasswordPolicyValue(PasswordPolicyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
