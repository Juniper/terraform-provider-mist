package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func uplinkPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApUplinkPortConfig) UplinkPortConfigValue {
	tflog.Debug(ctx, "uplinkPortConfigSdkToTerraform")
	var dot1x basetypes.BoolValue
	var keep_wlans_up_if_down basetypes.BoolValue

	if d.Dot1x != nil {
		dot1x = types.BoolValue(*d.Dot1x)
	}
	if d.KeepWlansUpIfDown != nil {
		keep_wlans_up_if_down = types.BoolValue(*d.KeepWlansUpIfDown)
	}

	data_map_attr_type := UplinkPortConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"dot1x":                 dot1x,
		"keep_wlans_up_if_down": keep_wlans_up_if_down,
	}
	data, e := NewUplinkPortConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
