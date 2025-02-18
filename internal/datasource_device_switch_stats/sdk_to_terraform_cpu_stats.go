package datasource_device_switch_stats

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func cpuStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.CpuStat) basetypes.ObjectValue {

	var idle basetypes.NumberValue
	var interrupt basetypes.NumberValue
	var loadAvg = types.ListNull(types.NumberType)
	var system basetypes.NumberValue
	var user basetypes.NumberValue

	if d.Idle.Value() != nil {
		idle = types.NumberValue(big.NewFloat(*d.Idle.Value()))
	}
	if d.Interrupt.Value() != nil {
		interrupt = types.NumberValue(big.NewFloat(*d.Interrupt.Value()))
	}
	if d.LoadAvg != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.NumberType{}
		for _, item := range d.LoadAvg {
			items = append(items, types.NumberValue(big.NewFloat(item)))
		}
		list, e := types.ListValue(itemsType, items)
		if e != nil {
			diags.Append(e...)
		}

		loadAvg = list
	}
	if d.System.Value() != nil {
		system = types.NumberValue(big.NewFloat(*d.System.Value()))
	}
	if d.User.Value() != nil {
		user = types.NumberValue(big.NewFloat(*d.User.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"idle":      idle,
		"interrupt": interrupt,
		"load_avg":  loadAvg,
		"system":    system,
		"user":      user,
	}
	data, e := basetypes.NewObjectValue(CpuStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
