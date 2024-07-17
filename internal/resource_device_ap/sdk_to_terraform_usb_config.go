package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func usbConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApUsb) UsbConfigValue {
	var cacert basetypes.StringValue
	var channel basetypes.Int64Value
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var port basetypes.Int64Value
	var type_usb basetypes.StringValue
	var verify_cert basetypes.BoolValue
	var vlan_id basetypes.Int64Value

	if d.Cacert.Value() != nil {
		cacert = types.StringValue(*d.Cacert.Value())
	}
	if d.Channel != nil {
		channel = types.Int64Value(int64(*d.Channel))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host != nil {
		host = types.StringValue(*d.Host)
	}
	if d.Port != nil {
		port = types.Int64Value(int64(*d.Port))
	}
	if d.Type != nil {
		type_usb = types.StringValue(string(*d.Type))
	}
	if d.VerifyCert != nil {
		verify_cert = types.BoolValue(*d.VerifyCert)
	}
	if d.VlanId != nil {
		vlan_id = types.Int64Value(int64(*d.VlanId))
	}

	data_map_attr_type := UsbConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cacert":      cacert,
		"channel":     channel,
		"enabled":     enabled,
		"host":        host,
		"port":        port,
		"type":        type_usb,
		"verify_cert": verify_cert,
		"vlan_id":     vlan_id,
	}
	data, e := NewUsbConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
