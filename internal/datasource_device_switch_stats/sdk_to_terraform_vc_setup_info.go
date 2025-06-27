package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vcSetupInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchVcSetupInfo) basetypes.ObjectValue {

	var configType basetypes.StringValue
	var currentStats basetypes.StringValue
	var errMissingDevIdFpc basetypes.BoolValue
	var lastUpdate basetypes.Float64Value
	var requestTime basetypes.Float64Value
	var requestType basetypes.StringValue

	if d.ConfigType != nil {
		configType = types.StringValue(*d.ConfigType)
	}
	if d.CurrentStats != nil {
		currentStats = types.StringValue(*d.CurrentStats)
	}
	if d.ErrMissingDevIdFpc != nil {
		errMissingDevIdFpc = types.BoolValue(*d.ErrMissingDevIdFpc)
	}
	if d.LastUpdate != nil {
		lastUpdate = types.Float64Value(*d.LastUpdate)
	}
	if d.RequestTime != nil {
		requestTime = types.Float64Value(*d.RequestTime)
	}
	if d.RequestType != nil {
		requestType = types.StringValue(*d.RequestType)
	}

	dataMapValue := map[string]attr.Value{
		"config_type":            configType,
		"current_stats":          currentStats,
		"err_missing_dev_id_fpc": errMissingDevIdFpc,
		"last_update":            lastUpdate,
		"request_time":           requestTime,
		"request_type":           requestType,
	}
	data, e := basetypes.NewObjectValue(VcSetupInfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
