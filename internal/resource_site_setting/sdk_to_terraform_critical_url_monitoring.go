package resource_site_setting

import (
	"context"

	mistapi "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func criticalUrlMonitoringMonitorSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SiteSettingCriticalUrlMonitoringMonitor) basetypes.ListValue {
	var dataList []MonitorsValue
	for _, d := range l {
		var url basetypes.StringValue
		var vlanId basetypes.StringValue

		if d.Url != nil {
			url = types.StringValue(*d.Url)
		}
		if d.VlanId != nil {
			vlanId = mistapi.VlanAsString(*d.VlanId)
		}

		dataMapValue := map[string]attr.Value{
			"url":     url,
			"vlan_id": vlanId,
		}
		data, e := NewMonitorsValue(MonitorsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := MonitorsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func criticalUrlMonitoringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingCriticalUrlMonitoring) CriticalUrlMonitoringValue {
	var enabled basetypes.BoolValue
	var monitors = types.ListNull(MonitorsValue{}.Type(ctx))

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Monitors != nil {
		monitors = criticalUrlMonitoringMonitorSdkToTerraform(ctx, diags, d.Monitors)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":  enabled,
		"monitors": monitors,
	}
	data, e := NewCriticalUrlMonitoringValue(CriticalUrlMonitoringValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
