package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pwrConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PwrConfigValue) *models.ApPwrConfig {
	data := models.ApPwrConfig{}

	if d.Base.ValueInt64Pointer() != nil {
		data.Base = models.ToPointer(int(d.Base.ValueInt64()))
	}
	if d.PreferUsbOverWifi.ValueBoolPointer() != nil {
		data.PreferUsbOverWifi = d.PreferUsbOverWifi.ValueBoolPointer()
	}

	return &data
}
