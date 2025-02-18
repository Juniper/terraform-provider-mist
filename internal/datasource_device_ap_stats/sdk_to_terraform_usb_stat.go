package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func usbStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApUsbStat) basetypes.ObjectValue {
	var channel basetypes.Int64Value
	var connected basetypes.BoolValue
	var lastActivity basetypes.Int64Value
	var typeUsb basetypes.StringValue
	var up basetypes.BoolValue

	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Connected.Value() != nil {
		connected = types.BoolValue(*d.Connected.Value())
	}
	if d.LastActivity.Value() != nil {
		lastActivity = types.Int64Value(int64(*d.LastActivity.Value()))
	}
	if d.Type.Value() != nil {
		typeUsb = types.StringValue(*d.Type.Value())
	}
	if d.Up.Value() != nil {
		up = types.BoolValue(*d.Up.Value())
	}

	dataMapValue := map[string]attr.Value{
		"channel":       channel,
		"connected":     connected,
		"last_activity": lastActivity,
		"type":          typeUsb,
		"up":            up,
	}
	data, e := types.ObjectValue(UsbStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
