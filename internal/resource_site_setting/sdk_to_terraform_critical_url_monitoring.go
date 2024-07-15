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

func criticalUrlMonitoringMonitorSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SiteSettingCriticalUrlMonitoringMonitor) basetypes.ListValue {
	tflog.Debug(ctx, "criticalUrlMonitoringMonitorSdkToTerraform")
	var data_list = []MonitorsValue{}
	for _, d := range l {
		var url basetypes.StringValue
		var vlan_id basetypes.Int64Value

		if d.Url != nil {
			url = types.StringValue(*d.Url)
		}
		if d.VlanId != nil {
			vlan_id = types.Int64Value(int64(*d.VlanId))
		}

		data_map_attr_type := MonitorsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"url":     url,
			"vlan_id": vlan_id,
		}
		data, e := NewMonitorsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := MonitorsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func criticalUrlMonitoringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingCriticalUrlMonitoring) CriticalUrlMonitoringValue {
	tflog.Debug(ctx, "criticalUrlMonitoringSdkToTerraform")
	var enabled basetypes.BoolValue
	var monitors basetypes.ListValue = types.ListNull(MonitorsValue{}.Type(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Monitors != nil {
		monitors = criticalUrlMonitoringMonitorSdkToTerraform(ctx, diags, d.Monitors)
	}

	data_map_attr_type := CriticalUrlMonitoringValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":  enabled,
		"monitors": monitors,
	}
	data, e := NewCriticalUrlMonitoringValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
