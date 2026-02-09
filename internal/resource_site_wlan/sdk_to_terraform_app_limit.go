package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appLimitSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.WlanAppLimit) AppLimitValue {
	if data == nil {
		return AppLimitValue{}
	}

	appLimitAttr := make(map[string]attr.Value)
	for key, value := range data.Apps {
		appLimitAttr[key] = types.Int64Value(int64(value))
	}
	apps := types.MapValueMust(types.Int64Type, appLimitAttr)

	wxtagLimitAttr := make(map[string]attr.Value)
	for key, value := range data.WxtagIds {
		wxtagLimitAttr[key] = types.Int64Value(int64(value))
	}
	wxtagIds := types.MapValueMust(types.Int64Type, wxtagLimitAttr)

	var enabled basetypes.BoolValue
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"apps":      apps,
		"enabled":   enabled,
		"wxtag_ids": wxtagIds,
	}
	result, err := NewAppLimitValue(AppLimitValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}
