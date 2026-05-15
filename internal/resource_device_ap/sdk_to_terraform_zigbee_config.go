package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func zigbeeConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApZigbee) ZigbeeConfigValue {
	var allowJoin basetypes.StringValue
	var channel basetypes.Int64Value
	var enabled basetypes.BoolValue
	var extendedPanId basetypes.StringValue
	var panId basetypes.StringValue

	if d.AllowJoin != nil {
		allowJoin = types.StringValue(string(*d.AllowJoin))
	}
	if d.Channel != nil {
		channel = types.Int64Value(int64(*d.Channel))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.ExtendedPanId.Value() != nil {
		extendedPanId = types.StringValue(*d.ExtendedPanId.Value())
	}
	if d.PanId.Value() != nil {
		panId = types.StringValue(*d.PanId.Value())
	}

	dataMapValue := map[string]attr.Value{
		"allow_join":      allowJoin,
		"channel":         channel,
		"enabled":         enabled,
		"extended_pan_id": extendedPanId,
		"pan_id":          panId,
	}
	data, e := NewZigbeeConfigValue(ZigbeeConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func zigbeeConfigTerraformToSdk(d ZigbeeConfigValue) *models.ApZigbee {
	data := models.ApZigbee{}

	if d.AllowJoin.ValueStringPointer() != nil {
		data.AllowJoin = models.ToPointer(models.ApZigbeeAllowJoinEnum(d.AllowJoin.ValueString()))
	}
	if d.Channel.ValueInt64Pointer() != nil {
		data.Channel = models.ToPointer(int(d.Channel.ValueInt64()))
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.ExtendedPanId.ValueStringPointer() != nil {
		data.ExtendedPanId = models.NewOptional(d.ExtendedPanId.ValueStringPointer())
	}
	if d.PanId.ValueStringPointer() != nil {
		data.PanId = models.NewOptional(d.PanId.ValueStringPointer())
	}

	return &data
}
