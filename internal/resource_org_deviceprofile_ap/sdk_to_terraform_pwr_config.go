package resource_org_deviceprofile_ap

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
	var preferUsbOverWifi basetypes.BoolValue

	if d.Base != nil {
		base = types.Int64Value(int64(*d.Base))
	}
	if d.PreferUsbOverWifi != nil {
		preferUsbOverWifi = types.BoolValue(*d.PreferUsbOverWifi)
	}

	dataMapValue := map[string]attr.Value{
		"base":                 base,
		"prefer_usb_over_wifi": preferUsbOverWifi,
	}
	data, e := NewPwrConfigValue(PwrConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
