package resource_site_setting

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func zoneOccupancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteZoneOccupancyAlert) ZoneOccupancyAlertValue {
	tflog.Debug(ctx, "zoneOccupancySdkToTerraform")
	var email_notifiers basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var enabled basetypes.BoolValue
	var threshold basetypes.Int64Value

	if d.EmailNotifiers != nil {
		email_notifiers = mist_transform.ListOfStringSdkToTerraform(ctx, d.EmailNotifiers)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Threshold != nil {
		threshold = types.Int64Value(int64(*d.Threshold))
	}

	data_map_attr_type := ZoneOccupancyAlertValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"email_notifiers": email_notifiers,
		"enabled":         enabled,
		"threshold":       threshold,
	}
	data, e := NewZoneOccupancyAlertValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
