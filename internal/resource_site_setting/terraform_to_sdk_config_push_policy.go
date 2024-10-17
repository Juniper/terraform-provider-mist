package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_hours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pushPolicyPushWindowConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.PushPolicyPushWindow {
	data := models.PushPolicyPushWindow{}

	if !d.IsNull() && !d.IsUnknown() {
		vd, e := NewPushWindowValue(PushWindowValue{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Enabled = vd.Enabled.ValueBoolPointer()

			hours := mist_hours.HoursTerraformToSdk(ctx, diags, vd.Hours)
			data.Hours = hours
		}
	}
	return &data
}

func pushPolicyConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ConfigPushPolicyValue) *models.SiteSettingConfigPushPolicy {
	data := models.SiteSettingConfigPushPolicy{}

	if d.NoPush.ValueBoolPointer() != nil {
		data.NoPush = d.NoPush.ValueBoolPointer()
	}
	if !d.PushWindow.IsNull() && !d.PushWindow.IsUnknown() {
		data.PushWindow = pushPolicyPushWindowConfigTerraformToSdk(ctx, diags, d.PushWindow)
	}

	return &data
}
