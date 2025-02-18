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

	dataMapValue := map[string]attr.Value{
		"brightness": brightness,
		"enabled":    enabled,
	}
	data, e := NewLedValue(LedValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
