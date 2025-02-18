package resource_site_setting

import (
	"context"

	misthours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func configPushPolicyWindowSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.PushPolicyPushWindow) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	var hours = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Hours != nil {
		hours = misthours.HoursSdkToTerraform(diags, d.Hours)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"hours":   hours,
	}
	data, e := basetypes.NewObjectValue(PushWindowValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func configPushPolicySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingConfigPushPolicy) ConfigPushPolicyValue {
	var noPush basetypes.BoolValue
	var pushWindow = types.ObjectNull(PushWindowValue{}.AttributeTypes(ctx))

	if d != nil && d.NoPush != nil {
		noPush = types.BoolValue(*d.NoPush)
	}
	if d != nil && d.PushWindow != nil {
		pushWindow = configPushPolicyWindowSdkToTerraform(ctx, diags, d.PushWindow)
	}

	dataMapValue := map[string]attr.Value{
		"no_push":     noPush,
		"push_window": pushWindow,
	}
	data, e := NewConfigPushPolicyValue(ConfigPushPolicyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
