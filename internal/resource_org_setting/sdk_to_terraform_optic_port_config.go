package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func opticPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OpticPortConfigPort) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {

		var channelized basetypes.BoolValue
		var speed basetypes.StringValue

		if d.Channelized != nil {
			channelized = types.BoolValue(*d.Channelized)
		}
		if d.Speed != nil {
			speed = types.StringValue(*d.Speed)
		}

		dataMapValue := map[string]attr.Value{
			"channelized": channelized,
			"speed":       speed,
		}
		data, e := NewOpticPortConfigValue(OpticPortConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := OpticPortConfigValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
