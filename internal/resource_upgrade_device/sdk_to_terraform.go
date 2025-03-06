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

func SdkToTerraform(upgrade UpgradeDeviceModel, data *models.ResponseDeviceUpgrade) (UpgradeDeviceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	var deviceVersion basetypes.StringValue
	var fwupdate = NewFwupdateValueNull()
	var status = types.StringValue(string(data.Status))
	var timestamp = types.Float64Value(data.Timestamp)
	var syncUpgrade = types.BoolValue(true)
	var syncUpgradeStartTimeout = types.Int64Value(60)
	var syncUpgradeRefreshInterval = types.Int64Value(15)
	var syncUpgradeTimeout = types.Int64Value(1800)

	if !upgrade.SyncUpgrade.IsNull() && !upgrade.SyncUpgrade.IsUnknown() {
		syncUpgrade = upgrade.SyncUpgrade
	}
	if !upgrade.SyncUpgradeStartTimeout.IsNull() && !upgrade.SyncUpgradeStartTimeout.IsUnknown() {
		syncUpgradeStartTimeout = upgrade.SyncUpgradeStartTimeout
	}
	if !upgrade.SyncUpgradeRefreshInterval.IsNull() && !upgrade.SyncUpgradeRefreshInterval.IsUnknown() {
		syncUpgradeRefreshInterval = upgrade.SyncUpgradeRefreshInterval
	}
	if !upgrade.SyncUpgradeTimeout.IsNull() && !upgrade.SyncUpgradeTimeout.IsUnknown() {
		syncUpgradeTimeout = upgrade.SyncUpgradeTimeout
	}

	upgrade.DeviceVersion = deviceVersion
	upgrade.Fwupdate = fwupdate
	upgrade.Status = status
	upgrade.Timestamp = types.Number(timestamp)
	upgrade.SyncUpgrade = syncUpgrade
	upgrade.SyncUpgradeStartTimeout = syncUpgradeStartTimeout
	upgrade.SyncUpgradeRefreshInterval = syncUpgradeRefreshInterval
	upgrade.SyncUpgradeTimeout = syncUpgradeTimeout

	return upgrade, diags
}

func DeviceStatSdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.ApiResponse[models.StatsDevice]) (UpgradeDeviceModel, int, diag.Diagnostics) {

	var diags diag.Diagnostics

	var fwupdate = NewFwupdateValueNull()
	var deviceVersion basetypes.StringValue
	var uptime = -1

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
	var statusId basetypes.Int64Value
	var timestamp basetypes.Float64Value
	var willRetry basetypes.BoolValue

	if deviceFwUpdate != nil {
		if deviceFwUpdate.Progress.Value() != nil {
			progress = types.Int64Value(int64(*deviceFwUpdate.Progress.Value()))
		}
		if deviceFwUpdate.Status.Value() != nil {
			status = types.StringValue(string(*deviceFwUpdate.Status.Value()))
		}
		if deviceFwUpdate.StatusId.Value() != nil {
			statusId = types.Int64Value(int64(*deviceFwUpdate.StatusId.Value()))
		}
		if deviceFwUpdate.Timestamp != nil {
			timestamp = types.Float64Value(*deviceFwUpdate.Timestamp)
		}
		if deviceFwUpdate.WillRetry.Value() != nil {
			willRetry = types.BoolValue(*deviceFwUpdate.WillRetry.Value())
		}
	}

	dataMapValue := map[string]attr.Value{
		"progress":   progress,
		"status":     status,
		"status_id":  statusId,
		"timestamp":  timestamp,
		"will_retry": willRetry,
	}
	fwupdate, e := NewFwupdateValue(FwupdateValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return fwupdate
}
