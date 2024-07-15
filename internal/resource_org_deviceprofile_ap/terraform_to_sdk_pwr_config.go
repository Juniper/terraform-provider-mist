package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pwrConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d PwrConfigValue) *models.ApPwrConfig {
	tflog.Debug(ctx, "pwrConfigTerraformToSdk")
	data := models.ApPwrConfig{}

	if d.Base.ValueInt64Pointer() != nil {
		data.Base = models.ToPointer(int(d.Base.ValueInt64()))
	}
	if d.PreferUsbOverWifi.ValueBoolPointer() != nil {
		data.PreferUsbOverWifi = d.PreferUsbOverWifi.ValueBoolPointer()
	}

	return &data
}
