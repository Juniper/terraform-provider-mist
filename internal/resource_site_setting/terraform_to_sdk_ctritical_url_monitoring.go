package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func criticalUrlMonitoringMonitorsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SiteSettingCriticalUrlMonitoringMonitor {
	var data_list []models.SiteSettingCriticalUrlMonitoringMonitor
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MonitorsValue)
		data := models.SiteSettingCriticalUrlMonitoringMonitor{}
		data.Url = plan.Url.ValueStringPointer()
		if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
			data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func criticalUrlMonitoringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CriticalUrlMonitoringValue) *models.SiteSettingCriticalUrlMonitoring {
	data := models.SiteSettingCriticalUrlMonitoring{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	monitors := criticalUrlMonitoringMonitorsTerraformToSdk(ctx, diags, d.Monitors)
	data.Monitors = monitors

	return &data
}
