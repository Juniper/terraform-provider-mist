package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ledSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApLed) LedValue {

	var brightness basetypes.Int64Value
	var enabled basetypes.BoolValue

	if d != nil && d.Brightness != nil {
		brightness = types.Int64Value(int64(*d.Brightness))
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := LedValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"brightness": brightness,
		"enabled":    enabled,
	}
	data, e := NewLedValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
