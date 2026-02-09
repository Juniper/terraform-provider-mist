package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appLimitSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.WlanAppLimit) basetypes.ObjectValue {
	appLimitAttr := make(map[string]attr.Value)
	for attr, val := range data.Apps {
		appLimitAttr[attr] = types.Int64Value(int64(val))
	}
	apps := types.MapValueMust(types.Int64Type, appLimitAttr)

	var enabled basetypes.BoolValue
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	wxtagLimitAttr := make(map[string]attr.Value)
	for k, v := range data.WxtagIds {
		wxtagLimitAttr[k] = types.Int64Value(int64(v))
	}
	wxtagIds := types.MapValueMust(types.Int64Type, wxtagLimitAttr)

	dataMapValue := map[string]attr.Value{
		"apps":      apps,
		"enabled":   enabled,
		"wxtag_ids": wxtagIds,
	}
	result, err := basetypes.NewObjectValue(AppLimitValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}
