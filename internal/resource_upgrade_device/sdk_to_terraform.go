package resource_upgrade_device

import (
	"context"
	"encoding/json"
	"io"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func SdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.ResponseDeviceUpgrade) (UpgradeDeviceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	var device_version basetypes.StringValue
	var fwupdate FwupdateValue = NewFwupdateValueNull()
	var status basetypes.StringValue = types.StringValue(string(data.Status))
	var timestamp basetypes.Float64Value = types.Float64Value(float64(data.Timestamp))
	var sync_upgrade basetypes.BoolValue = types.BoolValue(true)
	var sync_upgrade_start_timeout basetypes.Int64Value = types.Int64Value(60)
	var sync_upgrade_refresh_interval basetypes.Int64Value = types.Int64Value(15)
	var sync_upgrade_timeout basetypes.Int64Value = types.Int64Value(1800)

	if !upgrade.SyncUpgrade.IsNull() && !upgrade.SyncUpgrade.IsUnknown() {
		sync_upgrade = upgrade.SyncUpgrade
	}
	if !upgrade.SyncUpgradeStartTimeout.IsNull() && !upgrade.SyncUpgradeStartTimeout.IsUnknown() {
		sync_upgrade_start_timeout = upgrade.SyncUpgradeStartTimeout
	}
	if !upgrade.SyncUpgradeRefreshInterval.IsNull() && !upgrade.SyncUpgradeRefreshInterval.IsUnknown() {
		sync_upgrade_refresh_interval = upgrade.SyncUpgradeRefreshInterval
	}
	if !upgrade.SyncUpgradeTimeout.IsNull() && !upgrade.SyncUpgradeTimeout.IsUnknown() {
		sync_upgrade_timeout = upgrade.SyncUpgradeTimeout
	}

	upgrade.DeviceVersion = device_version
	upgrade.Fwupdate = fwupdate
	upgrade.Status = status
	upgrade.Timestamp = types.Number(timestamp)
	upgrade.SyncUpgrade = sync_upgrade
	upgrade.SyncUpgradeStartTimeout = sync_upgrade_start_timeout
	upgrade.SyncUpgradeRefreshInterval = sync_upgrade_refresh_interval
	upgrade.SyncUpgradeTimeout = sync_upgrade_timeout

	return upgrade, diags
}

func DeviceStatSdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.ApiResponse[models.StatsDevice]) (UpgradeDeviceModel, int, diag.Diagnostics) {

	var diags diag.Diagnostics

	var fwupdate FwupdateValue = NewFwupdateValueNull()
	var deviceVersion basetypes.StringValue
	var uptime int = -1

	body, _ := io.ReadAll(data.Response.Body)
	var objmap map[string]interface{}
	if err := json.Unmarshal(body, &objmap); err != nil {
		tflog.Error(ctx, err.Error())
	} else {
		if objmap["type"] == "ap" {
			stats := models.StatsAp{}
			json.Unmarshal(body, &stats)
			fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
			if stats.Version.Value() != nil {
				deviceVersion = types.StringValue(*stats.Version.Value())
			}
			if stats.Uptime.Value() != nil {
				uptime = int(*stats.Uptime.Value())
			}
		} else if objmap["type"] == "switch" {
			stats := models.StatsSwitch{}
			json.Unmarshal(body, &stats)
			fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
			if stats.Version.Value() != nil {
				deviceVersion = types.StringValue(*stats.Version.Value())
			}
			if stats.Uptime.Value() != nil {
				uptime = int(*stats.Uptime.Value())
			}
		} else if objmap["type"] == "gateway" {
			stats := models.StatsGateway{}
			json.Unmarshal(body, &stats)
			fwupdate = fwUpdateSdtToTerraform(ctx, &diags, stats.Fwupdate)
			if stats.Version.Value() != nil {
				deviceVersion = types.StringValue(*stats.Version.Value())
			}
			if stats.Uptime.Value() != nil {
				uptime = int(*stats.Uptime.Value())
			}
		}
	}

	upgrade.Fwupdate = fwupdate
	upgrade.Status = fwupdate.Status
	upgrade.DeviceVersion = deviceVersion
	return upgrade, uptime, diags
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
