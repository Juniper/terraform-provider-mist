package datasource_device_ap_stats

import (
	"context"
	"math/big"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func lldpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStatsLldpStat) basetypes.ObjectValue {

	var chassis_id basetypes.StringValue
	var lldp_med_supported basetypes.BoolValue
	var mgmt_addr basetypes.StringValue
	var mgmt_addrs basetypes.ListValue = types.ListUnknown(types.StringType)
	var port_desc basetypes.StringValue
	var port_id basetypes.StringValue
	var power_allocated basetypes.NumberValue
	var power_draw basetypes.NumberValue
	var power_request_count basetypes.Int64Value
	var power_requested basetypes.NumberValue
	var system_desc basetypes.StringValue
	var system_name basetypes.StringValue

	if d.ChassisId.Value() != nil {
		chassis_id = types.StringValue(*d.ChassisId.Value())
	}
	if d.LldpMedSupported.Value() != nil {
		lldp_med_supported = types.BoolValue(*d.LldpMedSupported.Value())
	}
	if d.MgmtAddr.Value() != nil {
		mgmt_addr = types.StringValue(*d.MgmtAddr.Value())
	}
	if d.MgmtAddrs != nil {
		mgmt_addrs = mist_transform.ListOfStringSdkToTerraform(ctx, d.MgmtAddrs)
	}
	if d.PortDesc.Value() != nil {
		port_desc = types.StringValue(*d.PortDesc.Value())
	}
	if d.PortId.Value() != nil {
		port_id = types.StringValue(*d.PortId.Value())
	}
	if d.PowerAllocated.Value() != nil {
		power_allocated = types.NumberValue(big.NewFloat(*d.PowerAllocated.Value()))
	}
	if d.PowerDraw.Value() != nil {
		power_draw = types.NumberValue(big.NewFloat(*d.PowerDraw.Value()))
	}
	if d.PowerRequestCount.Value() != nil {
		power_request_count = types.Int64Value(int64(*d.PowerRequestCount.Value()))
	}
	if d.PowerRequested.Value() != nil {
		power_requested = types.NumberValue(big.NewFloat(*d.PowerRequested.Value()))
	}
	if d.SystemDesc.Value() != nil {
		system_desc = types.StringValue(*d.SystemDesc.Value())
	}
	if d.SystemName.Value() != nil {
		system_name = types.StringValue(*d.SystemName.Value())
	}

	data_map_attr_type := LldpStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"chassis_id":          chassis_id,
		"lldp_med_supported":  lldp_med_supported,
		"mgmt_addr":           mgmt_addr,
		"mgmt_addrs":          mgmt_addrs,
		"port_desc":           port_desc,
		"port_id":             port_id,
		"power_allocated":     power_allocated,
		"power_draw":          power_draw,
		"power_request_count": power_request_count,
		"power_requested":     power_requested,
		"system_desc":         system_desc,
		"system_name":         system_name,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
