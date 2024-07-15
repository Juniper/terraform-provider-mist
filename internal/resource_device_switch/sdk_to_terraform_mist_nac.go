package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// ////////////////// MIST NAC ///////////////////////
func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMistNac) MistNacValue {

	var enabled basetypes.BoolValue
	var network basetypes.StringValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Network != nil {
		network = types.StringValue(*d.Network)
	}

	data_map_attr_type := MistNacValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
		"network": network,
	}
	data, e := NewMistNacValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
