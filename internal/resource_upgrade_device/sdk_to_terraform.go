package resource_upgrade_device

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.ResponseUpgradeDevice) (UpgradeDeviceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	var fwupdate FwupdateValue = NewFwupdateValueNull()
	var status basetypes.StringValue = types.StringValue(string(data.Status))
	var timestamp basetypes.Float64Value = types.Float64Value(float64(data.Timestamp))

	upgrade.Fwupdate = fwupdate
	upgrade.Status = status
	upgrade.Timestamp = types.Number(timestamp)

	return upgrade, diags
}

func DeviceStatSdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.StatsDevice) (UpgradeDeviceModel, diag.Diagnostics) {

	var diags diag.Diagnostics

	var fwupdate FwupdateValue = NewFwupdateValueNull()

	if stats, ok := data.AsStatsAp(); ok {
		fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
	} else if stats, ok := data.AsStatsGateway(); ok {
		fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
	} else if stats, ok := data.AsStatsSwitch(); ok {
		fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
	}

	upgrade.Fwupdate = fwupdate
	upgrade.Status = fwupdate.Status
	return upgrade, diags
}

func fwUpdateSdtToTerraform(ctx context.Context, diags *diag.Diagnostics, deviceFwUpdate *models.FwupdateStat) FwupdateValue {

	var progress basetypes.Int64Value
	var status basetypes.StringValue
	var status_id basetypes.Int64Value
	var timestamp basetypes.Float64Value
	var will_retry basetypes.BoolValue

	if deviceFwUpdate != nil {
		if deviceFwUpdate.Progress.Value() != nil {
			progress = types.Int64Value(int64(*deviceFwUpdate.Progress.Value()))
		}
		if deviceFwUpdate.Status.Value() != nil {
			status = types.StringValue(string(*deviceFwUpdate.Status.Value()))
		}
		if deviceFwUpdate.StatusId.Value() != nil {
			status_id = types.Int64Value(int64(*deviceFwUpdate.StatusId.Value()))
		}
		if deviceFwUpdate.Timestamp.Value() != nil {
			timestamp = types.Float64Value(float64(*deviceFwUpdate.Timestamp.Value()))
		}
		if deviceFwUpdate.WillRetry.Value() != nil {
			will_retry = types.BoolValue(*deviceFwUpdate.WillRetry.Value())
		}
	}

	data_map_value := map[string]attr.Value{
		"progress":   progress,
		"status":     status,
		"status_id":  status_id,
		"timestamp":  timestamp,
		"will_retry": will_retry,
	}
	fwupdate, e := NewFwupdateValue(FwupdateValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)
	return fwupdate
}
