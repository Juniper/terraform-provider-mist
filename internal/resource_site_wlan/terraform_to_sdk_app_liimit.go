package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appLimitTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppLimitValue) *models.WlanAppLimit {

	data := models.WlanAppLimit{}

	app_limit := make(map[string]int)
	for k, v := range plan.Apps.Elements() {
		var v_interface interface{} = v
		app_limit[k] = int(v_interface.(int64))
	}

	wxtags_limit := make(map[string]int)
	for k, v := range plan.WxtagIds.Elements() {
		var v_interface interface{} = v
		wxtags_limit[k] = int(v_interface.(int64))
	}

	data.Apps = app_limit
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.WxtagIds = wxtags_limit

	return &data
}
