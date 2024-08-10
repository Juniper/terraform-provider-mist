package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func clientsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsSwitchClientItem) basetypes.ListValue {

	var data_list = []ClientsValue{}

	for _, d := range l {
		var device_mac basetypes.StringValue
		var hostname basetypes.StringValue
		var mac basetypes.StringValue
		var port_id basetypes.StringValue

		if d.DeviceMac != nil {
			device_mac = types.StringValue(*d.DeviceMac)
		}
		if d.Hostname != nil {
			hostname = types.StringValue(*d.Hostname)
		}
		if d.Mac != nil {
			mac = types.StringValue(*d.Mac)
		}
		if d.PortId != nil {
			port_id = types.StringValue(*d.PortId)
		}

		data_map_attr_type := ClientsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"device_mac": device_mac,
			"hostname":   hostname,
			"mac":        mac,
			"port_id":    port_id,
		}
		data, e := NewClientsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ClientsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
