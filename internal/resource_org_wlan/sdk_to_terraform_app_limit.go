package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appLimitSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanAppLimit) AppLimitValue {
	var apps = types.MapNull(types.Int64Type)
	var enabled basetypes.BoolValue
	var wxtagIds = types.MapNull(types.Int64Type)

	appLimitAttr := make(map[string]attr.Value)
	for k, v := range d.Apps {
		appLimitAttr[k] = types.Int64Value(int64(v))
	}
	apps = types.MapValueMust(types.Int64Type, appLimitAttr)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	wxtagLimitAttr := make(map[string]attr.Value)
	for k, v := range d.WxtagIds {
		wxtagLimitAttr[k] = types.Int64Value(int64(v))
	}
	wxtagIds = types.MapValueMust(types.Int64Type, wxtagLimitAttr)

	dataMapValue := map[string]attr.Value{
		"apps":      apps,
		"enabled":   enabled,
		"wxtag_ids": wxtagIds,
	}
	data, e := NewAppLimitValue(AppLimitValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
