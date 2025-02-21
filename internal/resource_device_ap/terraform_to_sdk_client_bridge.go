package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clientBridgeAuthTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApClientBridgeAuth {
	data := models.ApClientBridgeAuth{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewAuthValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Psk.ValueStringPointer() != nil {
				data.Psk = plan.Psk.ValueStringPointer()
			}
			if plan.AuthType.ValueStringPointer() != nil {
				data.Type = models.ToPointer(models.ApClientBridgeAuthTypeEnum(plan.AuthType.ValueString()))
			}
		}
	}
	return &data
}
func clientBridgeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ClientBridgeValue) *models.ApClientBridge {
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
