package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func aeroscoutSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApAeroscout) AeroscoutValue {
	tflog.Debug(ctx, "aeroscoutSdkToTerraform")
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var locate_connected basetypes.BoolValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host.Value() != nil {
		host = types.StringValue(*d.Host.Value())
	}
	if d.LocateConnected != nil {
		locate_connected = types.BoolValue(*d.LocateConnected)
	}

	data_map_attr_type := AeroscoutValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":          enabled,
		"host":             host,
		"locate_connected": locate_connected,
	}
	data, e := NewAeroscoutValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
