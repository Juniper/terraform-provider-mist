package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clientAuthBridgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApClientBridgeAuth) basetypes.ObjectValue {

	var psk basetypes.StringValue
	var typeAuth basetypes.StringValue

	if d.Psk != nil {
		psk = types.StringValue(*d.Psk)
	}
	if d.Type != nil {
		typeAuth = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"psk":  psk,
		"type": typeAuth,
	}
	data, e := basetypes.NewObjectValue(AuthValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func clientBridgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApClientBridge) ClientBridgeValue {
	var auth = types.ObjectNull(AuthValue{}.AttributeTypes(ctx))
	var enabled basetypes.BoolValue
	var ssid basetypes.StringValue

	if d != nil && d.Auth != nil {
		auth = clientAuthBridgeSdkToTerraform(ctx, diags, d.Auth)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Ssid != nil {
		ssid = types.StringValue(*d.Ssid)
	}

	dataMapValue := map[string]attr.Value{
		"auth":    auth,
		"enabled": enabled,
		"ssid":    ssid,
	}
	data, e := NewClientBridgeValue(ClientBridgeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
