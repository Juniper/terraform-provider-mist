package datasource_device_gateway_stats

import (
	"context"
	"math/big"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
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
		loadAvg = misttransform.ListOfNumberSdkToTerraform(d.LoadAvg)
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
	data, e := types.ObjectValue(CpuStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
