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
	var deviceVersion basetypes.StringValue
	fwupdate := NewFwupdateValueNull()
	status := types.StringValue(string(data.Status))
	timestamp := types.Float64Value(data.Timestamp)

	syncUpgrade := types.BoolValue(true)
	if !upgrade.SyncUpgrade.IsNull() && !upgrade.SyncUpgrade.IsUnknown() {
		syncUpgrade = upgrade.SyncUpgrade
	}

	syncUpgradeStartTimeout := types.Int64Value(60)
	if !upgrade.SyncUpgradeStartTimeout.IsNull() && !upgrade.SyncUpgradeStartTimeout.IsUnknown() {
		syncUpgradeStartTimeout = upgrade.SyncUpgradeStartTimeout
	}

	syncUpgradeRefreshInterval := types.Int64Value(15)
	if !upgrade.SyncUpgradeRefreshInterval.IsNull() && !upgrade.SyncUpgradeRefreshInterval.IsUnknown() {
		syncUpgradeRefreshInterval = upgrade.SyncUpgradeRefreshInterval
	}

	syncUpgradeTimeout := types.Int64Value(1800)
	if !upgrade.SyncUpgradeTimeout.IsNull() && !upgrade.SyncUpgradeTimeout.IsUnknown() {
		syncUpgradeTimeout = upgrade.SyncUpgradeTimeout
	}

	upgrade.DeviceVersion = deviceVersion
	upgrade.Fwupdate = fwupdate
	upgrade.Status = status
	upgrade.Timestamp = timestamp
	upgrade.SyncUpgrade = syncUpgrade
	upgrade.SyncUpgradeStartTimeout = syncUpgradeStartTimeout
	upgrade.SyncUpgradeRefreshInterval = syncUpgradeRefreshInterval
	upgrade.SyncUpgradeTimeout = syncUpgradeTimeout

	var diags diag.Diagnostics
	return upgrade, diags
}

func DeviceStatSdkToTerraform(ctx context.Context, upgrade UpgradeDeviceModel, data *models.ApiResponse[models.StatsDevice]) (UpgradeDeviceModel, int, diag.Diagnostics) {
	var diags diag.Diagnostics
	defaultUptime := -1
	body, err := io.ReadAll(data.Response.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading response body", map[string]interface{}{"error": err.Error()})
		return upgrade, defaultUptime, diags
	}

	var objMap map[string]interface{}
	err = json.Unmarshal(body, &objMap)
	if err != nil {
		tflog.Error(ctx, err.Error())
		return upgrade, defaultUptime, diags
	}

	deviceType, ok := objMap["type"].(string)
	if !ok {
		tflog.Error(ctx, "Device type not found / invalid")
		return upgrade, defaultUptime, diags
	}

	switch deviceType {
	case "ap":
		var stats models.StatsAp
		err := json.Unmarshal(body, &stats)
		if err != nil {
			tflog.Error(ctx, "Error unmarshaling ap stats", map[string]interface{}{"error": err.Error()})
			return upgrade, defaultUptime, diags
		}
		return processCommonFields(ctx, upgrade, &diags, stats.AutoUpgradeStat, stats.Fwupdate, stats.Version.Value(), stats.Uptime.Value())

	case "switch":
		var stats models.StatsSwitch
		err := json.Unmarshal(body, &stats)
		if err != nil {
			tflog.Error(ctx, "Error unmarshaling switch stats", map[string]interface{}{"error": err.Error()})
			return upgrade, defaultUptime, diags
		}

		if stats.ConfigTimestamp != nil {
			upgrade.ConfigTimestamp = types.Int64Value(int64(*stats.ConfigTimestamp))
		}

		if stats.ConfigVersion != nil {
			upgrade.ConfigVersion = types.Int64Value(int64(*stats.ConfigVersion))
		}

		if stats.ExtIp != nil {
			upgrade.ExtIp = types.StringValue(*stats.ExtIp)
		}

		if stats.TagId != nil {
			upgrade.TagId = types.Int64Value(int64(*stats.TagId))
		}

		if stats.TagUuid != nil {
			upgrade.TagUuid = types.StringValue(stats.TagUuid.String())
		}

		return processCommonFields(ctx, upgrade, &diags, stats.AutoUpgradeStat, stats.Fwupdate, stats.Version.Value(), stats.Uptime.Value())

	case "gateway":
		var stats models.StatsGateway
		err := json.Unmarshal(body, &stats)
		if err != nil {
			tflog.Error(ctx, "Error unmarshaling gateway stats", map[string]interface{}{"error": err.Error()})
			return upgrade, defaultUptime, diags
		}

		if stats.ConfigTimestamp != nil {
			upgrade.ConfigTimestamp = types.Int64Value(int64(*stats.ConfigTimestamp))
		}

		if stats.ConfigVersion != nil {
			upgrade.ConfigVersion = types.Int64Value(int64(*stats.ConfigVersion))
		}

		if stats.ExtIp.Value() != nil {
			upgrade.ExtIp = types.StringValue(*stats.ExtIp.Value())
		}

		if stats.TagId != nil {
			upgrade.TagId = types.Int64Value(int64(*stats.TagId))
		}

		if stats.TagUuid != nil {
			upgrade.TagUuid = types.StringValue(stats.TagUuid.String())
		}

		return processCommonFields(ctx, upgrade, &diags, stats.AutoUpgradeStat, stats.Fwupdate, stats.Version.Value(), stats.Uptime.Value())

	default:
		tflog.Error(ctx, "Unknown device type", map[string]interface{}{"type": deviceType})
		return upgrade, defaultUptime, diags
	}
}

func processCommonFields(
	ctx context.Context,
	upgrade UpgradeDeviceModel,
	diags *diag.Diagnostics,
	autoUpgradeStat *models.StatsApAutoUpgrade,
	fwupdate *models.FwupdateStat,
	version *string,
	uptime *float64,
) (UpgradeDeviceModel, int, diag.Diagnostics) {
	upgrade.AutoUpgradeStat = autoUpgradeStatSdkToTerraform(ctx, diags, autoUpgradeStat)
	upgrade.Fwupdate = fwUpdateSdtToTerraform(ctx, diags, fwupdate)
	upgrade.Status = upgrade.Fwupdate.Status

	if version != nil {
		upgrade.DeviceVersion = types.StringValue(*version)
	}

	uptimeInt := -1
	if uptime != nil {
		uptimeInt = int(*uptime)
	}

	return upgrade, uptimeInt, *diags
}

func autoUpgradeStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.StatsApAutoUpgrade) AutoUpgradeStatValue {
	if data == nil {
		return NewAutoUpgradeStatValueNull()
	}

	var lastcheck basetypes.Int64Value
	if data.Lastcheck.Value() != nil {
		lastcheck = types.Int64Value(int64(*data.Lastcheck.Value()))
	}

	dataMap := map[string]attr.Value{
		"lastcheck": lastcheck,
	}
	result, err := NewAutoUpgradeStatValue(AutoUpgradeStatValue{}.AttributeTypes(ctx), dataMap)
	diags.Append(err...)

	return result
}

func fwUpdateSdtToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.FwupdateStat) FwupdateValue {
	if data == nil {
		return NewFwupdateValueNull()
	}

	var progress basetypes.Int64Value
	if data.Progress.Value() != nil {
		progress = types.Int64Value(int64(*data.Progress.Value()))
	}

	var status basetypes.StringValue
	if data.Status.Value() != nil {
		status = types.StringValue(string(*data.Status.Value()))
	}

	var statusId basetypes.Int64Value
	if data.StatusId.Value() != nil {
		statusId = types.Int64Value(int64(*data.StatusId.Value()))
	}

	var timestamp basetypes.Float64Value
	if data.Timestamp != nil {
		timestamp = types.Float64Value(*data.Timestamp)
	}

	var willRetry basetypes.BoolValue
	if data.WillRetry.Value() != nil {
		willRetry = types.BoolValue(*data.WillRetry.Value())
	}

	dataMap := map[string]attr.Value{
		"progress":   progress,
		"status":     status,
		"status_id":  statusId,
		"timestamp":  timestamp,
		"will_retry": willRetry,
	}
	result, err := NewFwupdateValue(FwupdateValue{}.AttributeTypes(ctx), dataMap)
	diags.Append(err...)

	return result
}
