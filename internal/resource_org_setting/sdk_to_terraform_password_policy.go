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
	var freshness basetypes.Int64Value
	var min_length basetypes.Int64Value
	var requires_special_char basetypes.BoolValue
	var requires_two_factor_auth basetypes.BoolValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Freshness != nil {
		freshness = types.Int64Value(int64(*d.Freshness))
	}
	if d.MinLength != nil {
		min_length = types.Int64Value(int64(*d.MinLength))
	}
	if d.RequiresSpecialChar != nil {
		requires_special_char = types.BoolValue(*d.RequiresSpecialChar)
	}
	if d.RequiresTwoFactorAuth != nil {
		requires_two_factor_auth = types.BoolValue(*d.RequiresTwoFactorAuth)
	}

	data_map_attr_type := PasswordPolicyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":                  enabled,
		"freshness":                freshness,
		"min_length":               min_length,
		"requires_special_char":    requires_special_char,
		"requires_two_factor_auth": requires_two_factor_auth,
	}
	data, e := NewPasswordPolicyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
