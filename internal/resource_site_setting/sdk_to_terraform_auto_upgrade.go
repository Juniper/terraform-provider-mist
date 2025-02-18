package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func autoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteSettingAutoUpgrade) AutoUpgradeValue {

	var customVersions = types.MapNull(types.StringType)
	var dayOfWeek basetypes.StringValue
	var enabled basetypes.BoolValue
	var timeOfDay basetypes.StringValue
	var version basetypes.StringValue

	if d.CustomVersions != nil {
		customVersionsMapValue := make(map[string]attr.Value)
		for k, v := range d.CustomVersions {
			customVersionsMapValue[k] = types.StringValue(v)
		}
		customVersions = types.MapValueMust(types.StringType, customVersionsMapValue)
	}

	if d.DayOfWeek != nil {
		dayOfWeek = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}
	if d.Version != nil {
		version = types.StringValue(string(*d.Version))
	}

	dataMapValue := map[string]attr.Value{
		"custom_versions": customVersions,
		"day_of_week":     dayOfWeek,
		"enabled":         enabled,
		"time_of_day":     timeOfDay,
		"version":         version,
	}
	data, e := NewAutoUpgradeValue(AutoUpgradeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
