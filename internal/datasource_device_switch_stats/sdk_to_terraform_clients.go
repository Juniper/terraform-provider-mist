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

	var dataList []ClientsValue

	for _, d := range l {
		var deviceMac basetypes.StringValue
		var hostname basetypes.StringValue
		var mac basetypes.StringValue
		var portId basetypes.StringValue

		if d.DeviceMac != nil {
			deviceMac = types.StringValue(*d.DeviceMac)
		}
		if d.Hostname != nil {
			hostname = types.StringValue(*d.Hostname)
		}
		if d.Mac != nil {
			mac = types.StringValue(*d.Mac)
		}
		if d.PortId != nil {
			portId = types.StringValue(*d.PortId)
		}

		dataMapValue := map[string]attr.Value{
			"device_mac": deviceMac,
			"hostname":   hostname,
			"mac":        mac,
			"port_id":    portId,
		}
		data, e := NewClientsValue(ClientsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ClientsValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
