package resource_site_wlan

import (
	"context"

	misthours "github.com/Juniper/terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func scheduleSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanSchedule) ScheduleValue {
	var enabled basetypes.BoolValue
	var hours = types.ObjectNull(HoursValue{}.AttributeTypes(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Hours != nil && len(d.Hours.String()) > 0 {
		hours = misthours.HoursSdkToTerraform(diags, d.Hours)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"hours":   hours,
	}
	data, e := NewScheduleValue(ScheduleValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
