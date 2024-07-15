package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func autoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteSettingAutoUpgrade) AutoUpgradeValue {
	tflog.Debug(ctx, "autoUpgradeSdkToTerraform")

	var custom_versions basetypes.MapValue = types.MapNull(types.StringType)
	var day_of_week basetypes.StringValue
	var enabled basetypes.BoolValue
	var time_of_day basetypes.StringValue
	var version basetypes.StringValue

	custom_versions_map_value := make(map[string]string)
	for k, v := range d.CustomVersions {
		custom_versions_map_value[k] = v
	}
	custom_versions, e := types.MapValueFrom(ctx, types.StringType, custom_versions_map_value)
	diags.Append(e...)

	if d.DayOfWeek != nil {
		day_of_week = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		time_of_day = types.StringValue(*d.TimeOfDay)
	}
	if d.Version != nil {
		version = types.StringValue(string(*d.Version))
	}

	data_map_attr_type := AutoUpgradeValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"custom_versions": custom_versions,
		"day_of_week":     day_of_week,
		"enabled":         enabled,
		"time_of_day":     time_of_day,
		"version":         version,
	}
	data, e := NewAutoUpgradeValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
