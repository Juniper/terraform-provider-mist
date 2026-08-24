package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func uwbConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApUwbConfig) UwbConfigValue {
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var port basetypes.Int64Value
	var slot basetypes.Int64Value
	var uwbConfigType basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host != nil {
		host = types.StringValue(*d.Host)
	}
	if d.Port != nil {
		port = types.Int64Value(int64(*d.Port))
	}
	if d.Slot != nil {
		slot = types.Int64Value(int64(*d.Slot))
	}
	if d.Type != nil {
		uwbConfigType = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"host":    host,
		"port":    port,
		"slot":    slot,
		"type":    uwbConfigType,
	}
	data, e := NewUwbConfigValue(UwbConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
