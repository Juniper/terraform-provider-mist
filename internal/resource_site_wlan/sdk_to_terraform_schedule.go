package resource_site_wlan

import (
	"context"

	mist_hours "terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func scheduleSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanSchedule) ScheduleValue {
	var enabled basetypes.BoolValue
	var hours basetypes.ObjectValue = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Hours != nil {
		hours = mist_hours.HoursSdkToTerraform(ctx, diags, d.Hours)
	}

	data_map_attr_type := ScheduleValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
		"hours":   hours,
	}
	data, e := NewScheduleValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
