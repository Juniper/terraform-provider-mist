package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func evpnConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnConfig) EvpnConfigValue {
	var enabled basetypes.BoolValue
	var role basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Role != nil {
		role = types.StringValue(string(*d.Role))
	}

	data_map_attr_type := EvpnConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
		"role":    role,
	}
	data, e := NewEvpnConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
