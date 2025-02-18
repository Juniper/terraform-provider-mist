package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func eslStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApEslStat) basetypes.ObjectValue {
	var channel basetypes.Int64Value
	var connected basetypes.BoolValue
	var typeEsl basetypes.StringValue
	var up basetypes.BoolValue

	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Connected.Value() != nil {
		connected = types.BoolValue(*d.Connected.Value())
	}
	if d.Type.Value() != nil {
		typeEsl = types.StringValue(*d.Type.Value())
	}
	if d.Up.Value() != nil {
		up = types.BoolValue(*d.Up.Value())
	}

	dataMapValue := map[string]attr.Value{
		"channel":   channel,
		"connected": connected,
		"type":      typeEsl,
		"up":        up,
	}
	data, e := types.ObjectValue(EslStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
