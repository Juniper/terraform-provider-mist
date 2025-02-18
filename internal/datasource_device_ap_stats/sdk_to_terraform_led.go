package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ledSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApLed) basetypes.ObjectValue {

	var brightness basetypes.Int64Value
	var enabled basetypes.BoolValue

	if d.Brightness != nil {
		brightness = types.Int64Value(int64(*d.Brightness))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"brightness": brightness,
		"enabled":    enabled,
	}
	data, e := basetypes.NewObjectValue(LedValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
