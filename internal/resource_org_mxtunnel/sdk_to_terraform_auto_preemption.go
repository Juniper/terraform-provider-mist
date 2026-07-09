package resource_org_mxtunnel

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func autoPreemptionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AutoPreemption) AutoPreemptionValue {
	var dayOfWeek = types.StringNull()
	var enabled = types.BoolNull()
	var timeOfDay = types.StringNull()

	if d.DayOfWeek != nil {
		dayOfWeek = types.StringValue(string(*d.DayOfWeek))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}

	dataMapAttrType := AutoPreemptionValue{}.AttributeTypes(ctx)
	dataMapValue := map[string]attr.Value{
		"day_of_week": dayOfWeek,
		"enabled":     enabled,
		"time_of_day": timeOfDay,
	}
	data, e := NewAutoPreemptionValue(dataMapAttrType, dataMapValue)
	diags.Append(e...)

	return data
}
