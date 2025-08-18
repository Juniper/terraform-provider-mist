package resource_site_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func appLimitTerraformToSdk(plan AppLimitValue) *models.WlanAppLimit {

	data := models.WlanAppLimit{}

	appLimit := make(map[string]int)
	for k, v := range plan.Apps.Elements() {
		var vInterface interface{} = v
		appLimit[k] = int(vInterface.(basetypes.Int64Value).ValueInt64())
	}

	wxtagsLimit := make(map[string]int)
	for k, v := range plan.WxtagIds.Elements() {
		var vInterface interface{} = v
		wxtagsLimit[k] = int(vInterface.(basetypes.Int64Value).ValueInt64())
	}

	data.Apps = appLimit
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.WxtagIds = wxtagsLimit

	return &data
}
