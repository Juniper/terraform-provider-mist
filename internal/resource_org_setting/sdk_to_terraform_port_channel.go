package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func PortChannelSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.PortChannelization) PortChannelizationValue {

	var config basetypes.MapValue = types.MapNull(types.StringType)
	var enabled basetypes.BoolValue

	if len(d.AdditionalProperties) > 0 {
		config_items := make(map[string]string)
		for k, v := range d.AdditionalProperties {
			if k != "enabled" {
				var i interface{} = v
				s := i.(string)
				config_items[k] = s
			}
		}
		config_map, e := types.MapValueFrom(ctx, types.StringType, config_items)
		if e != nil {
			diags.Append(e...)
		} else {
			config = config_map
		}
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := PortChannelizationValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config":  config,
		"enabled": enabled,
	}
	data, e := NewPortChannelizationValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
