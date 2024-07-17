package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func pwrConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApPwrConfig) PwrConfigValue {
	var base basetypes.Int64Value
	var prefer_usb_over_wifi basetypes.BoolValue

	if d.Base != nil {
		base = types.Int64Value(int64(*d.Base))
	}
	if d.PreferUsbOverWifi != nil {
		prefer_usb_over_wifi = types.BoolValue(*d.PreferUsbOverWifi)
	}

	data_map_attr_type := PwrConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"base":                 base,
		"prefer_usb_over_wifi": prefer_usb_over_wifi,
	}
	data, e := NewPwrConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
