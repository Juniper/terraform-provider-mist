package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func criticalUrlMonitoringMonitorsTerraformToSdk(d basetypes.ListValue) []models.SiteSettingCriticalUrlMonitoringMonitor {
	var dataList []models.SiteSettingCriticalUrlMonitoringMonitor
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(MonitorsValue)
		data := models.SiteSettingCriticalUrlMonitoringMonitor{}
		data.Url = plan.Url.ValueStringPointer()
		if !plan.VlanId.IsNull() && !plan.VlanId.IsUnknown() {
			data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func criticalUrlMonitoringTerraformToSdk(d CriticalUrlMonitoringValue) *models.SiteSettingCriticalUrlMonitoring {
	data := models.SiteSettingCriticalUrlMonitoring{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	monitors := criticalUrlMonitoringMonitorsTerraformToSdk(d.Monitors)
	data.Monitors = monitors

	return &data
}
