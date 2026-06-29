package resource_org_deviceprofile_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func iotConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchIotPort) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)

	for k, d := range m {
		var alarmClass basetypes.StringValue
		var enabled basetypes.BoolValue
		var inputSrc basetypes.StringValue
		var name basetypes.StringValue

		if d.AlarmClass != nil {
			alarmClass = types.StringValue(string(*d.AlarmClass))
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		if d.InputSrc != nil {
			inputSrc = types.StringValue(string(*d.InputSrc))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		dataMapValue := map[string]attr.Value{
			"alarm_class": alarmClass,
			"enabled":     enabled,
			"input_src":   inputSrc,
			"name":        name,
		}
		data, e := NewIotConfigValue(IotConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}

	r, e := types.MapValueFrom(ctx, IotConfigValue{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return r
}
