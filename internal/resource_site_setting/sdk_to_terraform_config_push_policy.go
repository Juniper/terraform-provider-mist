package resource_site_setting

import (
	"context"

	mist_hours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func configPushPolicyWindowSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.PushPolicyPushWindow) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	var hours basetypes.ObjectValue = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Hours != nil {
		hours = mist_hours.HoursSdkToTerraform(ctx, diags, d.Hours)
	}

	data_map_attr_type := PushWindowValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
		"hours":   hours,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func configPushPolicySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingConfigPushPolicy) ConfigPushPolicyValue {
	var no_push basetypes.BoolValue
	var push_window basetypes.ObjectValue = types.ObjectNull(PushWindowValue{}.AttributeTypes(ctx))

	if d != nil && d.NoPush != nil {
		no_push = types.BoolValue(*d.NoPush)
	}
	if d != nil && d.PushWindow != nil {
		push_window = configPushPolicyWindowSdkToTerraform(ctx, diags, d.PushWindow)
	}

	data_map_attr_type := ConfigPushPolicyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"no_push":     no_push,
		"push_window": push_window,
	}
	data, e := NewConfigPushPolicyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
