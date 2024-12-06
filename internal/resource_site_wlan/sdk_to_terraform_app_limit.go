package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appLimitSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanAppLimit) AppLimitValue {
	var apps basetypes.MapValue = types.MapNull(types.Int64Type)
	var enabled basetypes.BoolValue
	var wxtag_ids basetypes.MapValue = types.MapNull(types.Int64Type)

	//if d != nil && d.Apps != nil && len(d.Apps) > 0 {
	app_limit_attr := make(map[string]attr.Value)
	for k, v := range d.Apps {
		app_limit_attr[k] = types.Int64Value(int64(v))
	}
	apps = types.MapValueMust(types.Int64Type, app_limit_attr)
	//}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	//if d != nil && d.WxtagIds != nil && len(d.WxtagIds) > 0 {
	wxtag_limit_attr := make(map[string]attr.Value)
	for k, v := range d.WxtagIds {
		wxtag_limit_attr[k] = types.Int64Value(int64(v))
	}
	wxtag_ids = types.MapValueMust(types.Int64Type, wxtag_limit_attr)
	//}

	data_map_attr_type := AppLimitValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"apps":      apps,
		"enabled":   enabled,
		"wxtag_ids": wxtag_ids,
	}
	data, e := NewAppLimitValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
