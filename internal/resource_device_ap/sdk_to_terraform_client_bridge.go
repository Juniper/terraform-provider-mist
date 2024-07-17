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
	var type_auth basetypes.StringValue

	if d.Psk != nil {
		psk = types.StringValue(*d.Psk)
	}
	if d.Type != nil {
		type_auth = types.StringValue(string(*d.Type))
	}

	data_map_attr_type := AuthValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"psk":  psk,
		"type": type_auth,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func clientBridgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApClientBridge) ClientBridgeValue {
	var auth basetypes.ObjectValue = types.ObjectNull(AuthValue{}.AttributeTypes(ctx))
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

	data_map_attr_type := ClientBridgeValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auth":    auth,
		"enabled": enabled,
		"ssid":    ssid,
	}
	data, e := NewClientBridgeValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
