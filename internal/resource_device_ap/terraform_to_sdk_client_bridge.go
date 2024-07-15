package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clientBridgeAuthTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApClientBridgeAuth {
	tflog.Debug(ctx, "clientBridgeAuthTerraformToSdk")
	data := models.ApClientBridgeAuth{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		var di interface{} = d
		dv := di.(AuthValue)
		if dv.Psk.ValueStringPointer() != nil {
			data.Psk = dv.Psk.ValueStringPointer()
		}
		if dv.AuthType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.ApClientBridgeAuthTypeEnum(dv.AuthType.ValueString()))
		}

		return &data
	}
}
func clientBridgeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ClientBridgeValue) *models.ApClientBridge {
	tflog.Debug(ctx, "clientBridgeTerraformToSdk")
	data := models.ApClientBridge{}

	auth := clientBridgeAuthTerraformToSdk(ctx, diags, d.Auth)
	if !d.Auth.IsNull() && !d.Auth.IsUnknown() {
		data.Auth = auth
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Ssid.ValueStringPointer() != nil {
		data.Ssid = d.Ssid.ValueStringPointer()
	}

	return &data
}
