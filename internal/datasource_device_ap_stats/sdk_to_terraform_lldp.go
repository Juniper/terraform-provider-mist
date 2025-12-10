package datasource_device_ap_stats

import (
	"context"
	"math/big"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func lldpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApLldpStat) basetypes.ObjectValue {

	var chassisId basetypes.StringValue
	var lldpMedSupported basetypes.BoolValue
	var mgmtAddr basetypes.StringValue
	var mgmtAddrs = types.ListUnknown(types.StringType)
	var portDesc basetypes.StringValue
	var portId basetypes.StringValue
	var powerAllocated basetypes.NumberValue
	var powerAvail basetypes.Int64Value
	var powerBudget basetypes.Int64Value
	var powerConstrained basetypes.BoolValue
	var powerDraw basetypes.NumberValue
	var powerNeeded basetypes.Int64Value
	var powerOpmode basetypes.StringValue
	var powerSrc basetypes.StringValue
	var powerSrcs = types.ListNull(types.StringType)
	var powerRequestCount basetypes.Int64Value
	var powerRequested basetypes.NumberValue
	var systemDesc basetypes.StringValue
	var systemName basetypes.StringValue

	if d.ChassisId.Value() != nil {
		chassisId = types.StringValue(*d.ChassisId.Value())
	}
	if d.LldpMedSupported.Value() != nil {
		lldpMedSupported = types.BoolValue(*d.LldpMedSupported.Value())
	}
	if d.MgmtAddr.Value() != nil {
		mgmtAddr = types.StringValue(*d.MgmtAddr.Value())
	}
	if d.MgmtAddrs != nil {
		mgmtAddrs = mistutils.ListOfStringSdkToTerraform(d.MgmtAddrs)
	}
	if d.PortDesc.Value() != nil {
		portDesc = types.StringValue(*d.PortDesc.Value())
	}
	if d.PortId.Value() != nil {
		portId = types.StringValue(*d.PortId.Value())
	}
	if d.PowerAllocated.Value() != nil {
		powerAllocated = types.NumberValue(big.NewFloat(*d.PowerAllocated.Value()))
	}
	if d.PowerAvail != nil {
		powerAvail = types.Int64Value(int64(*d.PowerAvail))
	}
	if d.PowerBudget != nil {
		powerBudget = types.Int64Value(int64(*d.PowerBudget))
	}
	if d.PowerConstrained != nil {
		powerConstrained = types.BoolValue(*d.PowerConstrained)
	}
	if d.PowerDraw.Value() != nil {
		powerDraw = types.NumberValue(big.NewFloat(*d.PowerDraw.Value()))
	}
	if d.PowerNeeded != nil {
		powerNeeded = types.Int64Value(int64(*d.PowerNeeded))
	}
	if d.PowerOpmode != nil {
		powerOpmode = types.StringValue(*d.PowerOpmode)
	}
	if d.PowerSrc != nil {
		powerSrc = types.StringValue(*d.PowerSrc)
	}
	if d.PowerSrcs != nil {
		powerSrcs = mistutils.ListOfStringSdkToTerraform(d.PowerSrcs)
	}
	if d.PowerRequestCount.Value() != nil {
		powerRequestCount = types.Int64Value(int64(*d.PowerRequestCount.Value()))
	}
	if d.PowerRequested.Value() != nil {
		powerRequested = types.NumberValue(big.NewFloat(*d.PowerRequested.Value()))
	}
	if d.SystemDesc.Value() != nil {
		systemDesc = types.StringValue(*d.SystemDesc.Value())
	}
	if d.SystemName.Value() != nil {
		systemName = types.StringValue(*d.SystemName.Value())
	}

	dataMapValue := map[string]attr.Value{
		"chassis_id":          chassisId,
		"lldp_med_supported":  lldpMedSupported,
		"mgmt_addr":           mgmtAddr,
		"mgmt_addrs":          mgmtAddrs,
		"port_desc":           portDesc,
		"port_id":             portId,
		"power_allocated":     powerAllocated,
		"power_avail":         powerAvail,
		"power_budget":        powerBudget,
		"power_constrained":   powerConstrained,
		"power_draw":          powerDraw,
		"power_needed":        powerNeeded,
		"power_opmode":        powerOpmode,
		"power_src":           powerSrc,
		"power_srcs":          powerSrcs,
		"power_request_count": powerRequestCount,
		"power_requested":     powerRequested,
		"system_desc":         systemDesc,
		"system_name":         systemName,
	}
	data, e := basetypes.NewObjectValue(LldpStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
