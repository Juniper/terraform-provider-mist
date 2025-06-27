package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_upgrade_device"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &upgradeDeviceResource{}
	_ resource.ResourceWithConfigure = &upgradeDeviceResource{}
)

func NewUpgradeDevice() resource.Resource {
	return &upgradeDeviceResource{}
}

type upgradeDeviceResource struct {
	client mistapi.ClientInterface
}

func (r *upgradeDeviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SSO client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *upgradeDeviceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_upgrade_device"
}

func (r *upgradeDeviceResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource can be used to upgrade the firmware of a single device (Wi-Fi Access Points, Switches and SRX/SSR Gateways).\n\n" +
			"The resource will send the upgrade command to Mist, which will take care of deploying the new firmware version to the " +
			"device, and reboot it if required.\n\n" +
			"The time required to upgrade a device depends on the type of device and its hardware. " +
			"By default, the resource will track the upgrade process and only return the result once the device is upgraded and rebooted " +
			"(unless `reboot`==`false` or `reboot_at` is set).  \n" +
			"If required it is possible to run the upgrade in async mode (attribute `sync`=`false`). " +
			"In this case, the resource will only trigger the upgrade and return the Mist response, but will not track the upgrade progress.\n\n" +
			"The list of available firmware versions can be retrieved with the `mist_device_versions` data source.",
		Attributes: resource_upgrade_device.UpgradeDeviceResourceSchema(ctx).Attributes,
	}
}

func (r *upgradeDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state resource_upgrade_device.UpgradeDeviceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, diags = r.startFwUpdate(ctx, state, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *upgradeDeviceResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_upgrade_device.UpgradeDeviceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, diags = r.refreshFwUpdate(ctx, state, false, 0, 0, 0)
	diags.Append(diags...)
	if diags.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *upgradeDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var state, plan resource_upgrade_device.UpgradeDeviceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, diags = r.startFwUpdate(ctx, state, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *upgradeDeviceResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}

func (r *upgradeDeviceResource) startFwUpdate(
	ctx context.Context,
	state resource_upgrade_device.UpgradeDeviceModel,
	plan resource_upgrade_device.UpgradeDeviceModel,
) (resource_upgrade_device.UpgradeDeviceModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"site_id\" value for \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return state, diags
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"device_id\" value for \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return state, diags
	}

	upgrade, diags := resource_upgrade_device.TerraformToSdk(&plan)
	diags.Append(diags...)
	if diags.HasError() {
		return state, diags
	}
	// Start upgrade
	data, err := r.client.UtilitiesUpgrade().UpgradeDevice(ctx, siteId, deviceId, upgrade)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		diags.AddError(
			"Error creating \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to Upgrade the device. %s", apiErr),
		)
		return state, diags
	} else if data.Data.Status == "device already running version" {
		diags.AddAttributeWarning(
			path.Root("version"),
			"Unable to upgrade the device",
			fmt.Sprintf("Device is already running the version %s", plan.TargetVersion.ValueString()),
		)
		state, diags = resource_upgrade_device.SdkToTerraform(plan, &data.Data)
		diags.Append(diags...)
		if diags.HasError() {
			return state, diags
		}
	} else {
		state, diags = resource_upgrade_device.SdkToTerraform(plan, &data.Data)
		diags.Append(diags...)
		if diags.HasError() {
			return state, diags
		}
		if !plan.SyncUpgrade.ValueBool() {
			tflog.Debug(ctx, "upgrade check is async, do not wait")
			state.Fwupdate.Progress = types.Int64Value(0)
			state.Fwupdate.Status = types.StringValue("scheduled")
			return state, diags
		} else {
			state, diags = r.refreshFwUpdate(
				ctx,
				state,
				true,
				plan.SyncUpgradeStartTimeout.ValueInt64(),
				plan.SyncUpgradeRefreshInterval.ValueInt64(),
				plan.SyncUpgradeTimeout.ValueInt64(),
			)
		}
		diags.Append(diags...)
		if diags.HasError() {
			return state, diags
		}
	}
	return state, diags
}

func (r *upgradeDeviceResource) refreshFwUpdate(
	ctx context.Context,
	state resource_upgrade_device.UpgradeDeviceModel,
	syncUpgrade bool,
	syncUpgradeStartTimeout int64,
	syncUpgradeRefreshInterval int64,
	syncUpgradeTimeout int64,
) (resource_upgrade_device.UpgradeDeviceModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var fields *string

	var previousStatusId int64 = -1
	var startTime = time.Now()
	var retry = 0
	var maxRetry = 3
	var deviceUptime = -1
	var upgradeStarted = false
	var uploadDone = false
	var upgradeDone = false
	var rebootDone = false

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"site_id\" value for \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return state, diags
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"device_id\" value for \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return state, diags
	}

	// Read status
	for {
		data, err := r.client.SitesStatsDevices().GetSiteDeviceStats(ctx, siteId, deviceId, fields)

		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			if retry < maxRetry {
				retry += 1
				time.Sleep(5 * time.Second)
			} else {
				diags.AddError(
					"Error reading device status for the \"mist_upgrade_device\" resource",
					fmt.Sprintf("Unable to retrieve the device upgrade status. %s", apiErr),
				)
				return state, diags
			}
		}

		retry = 0

		state, deviceUptime, diags = resource_upgrade_device.DeviceStatSdkToTerraform(ctx, state, &data)
		diags.Append(diags...)
		if diags.HasError() {
			diags.AddError(
				"Error reading device status for the \"mist_upgrade_device\" resource",
				fmt.Sprintf("Unable to retrieve the device upgrade status. %s", err.Error()),
			)
			return state, diags
		}

		if !syncUpgrade {
			tflog.Debug(ctx, "upgrade check is async, do not wait")
			return state, diags
		} else {
			if !upgradeStarted {
				// save the initial upgrade status id retrieved from Mist (it will increase once the upgrade starts)
				// this will be used to detect if/when the upgrade starts
				if previousStatusId == -1 {
					tflog.Debug(ctx, "upgrade check is sync, wait for the upgrade to start")
					previousStatusId = state.Fwupdate.StatusId.ValueInt64()
				}
				// If the upgrade didn't start but the start timeout is reached, raise an error
				if previousStatusId == state.Fwupdate.StatusId.ValueInt64() && time.Since(startTime).Seconds() > float64(syncUpgradeStartTimeout) {
					diags.AddError(
						"Error during device upgrade",
						"Upgrade start Timeout reached",
					)
					return state, diags
				}
				if previousStatusId != state.Fwupdate.StatusId.ValueInt64() {
					upgradeStarted = true
				}
			}

			// this can be executed during the same round. If we detect the upgrade started
			// we want to process the device stats
			if upgradeStarted {

				// if the upgrade timeout is reached, raise an error
				if time.Since(startTime).Seconds() > float64(syncUpgradeTimeout) {
					diags.AddError(
						"Error during device upgrade",
						fmt.Sprintf(
							"Upgrade end Timeout reached. The upgrade process will continue but will not be monitored "+
								"(current upgrade progress: %d%%)",
							state.Fwupdate.Progress.ValueInt64(),
						),
					)
					return state, diags
				}

				uploadDone, upgradeDone, rebootDone = checkUpgradeProgress(state, startTime, deviceUptime)

				// detect when the upload start / end (this is mostly between 0 and 70% of the progress)
				// this is only for logging purpose
				if !uploadDone {
					tflog.Info(ctx, fmt.Sprintf(
						"upgrade check is sync, wait for the end of the upload. "+
							"current progress: %d, current status: %s, current version: %s, requested version: %s",
						state.Fwupdate.Progress.ValueInt64(), state.Fwupdate.Status, state.DeviceVersion, state.TargetVersion),
					)

					// try to detect when the upgrade start / end (this is mostly between 70 and 100% of the progress)
					// this is only for logging purpose
				} else if !upgradeDone {
					tflog.Info(ctx, fmt.Sprintf(
						"upgrade check is sync, wait for the end of the upgrade. "+
							"current progress: %d, current status: %s, current version: %s, requested version: %s",
						state.Fwupdate.Progress.ValueInt64(), state.Fwupdate.Status, state.DeviceVersion, state.TargetVersion),
					)

					// if the upgrade process is at 100% and reboot is not requested or postponed,
					// do not wait for the reboot and return a success
				} else if !state.Reboot.ValueBool() || (!state.RebootAt.IsNull() && !state.RebootAt.IsUnknown()) {
					tflog.Info(ctx, fmt.Sprintf(
						"upgrade check is sync but do not wait for the reboot (not requested or postponed reboot). "+
							"current progress: %d, current status: %s, current version: %s, requested version: %s",
						state.Fwupdate.Progress.ValueInt64(), state.Fwupdate.Status, state.DeviceVersion, state.TargetVersion),
					)
					return state, diags

					// if the upgrade process is at 100% and reboot is  requested, wait for the device reboot and return a success
					// this is detected based on the device uptime
					// this is only for logging purpose
				} else if !rebootDone {
					tflog.Info(ctx, fmt.Sprintf(
						"upgrade check is sync, wait for the end of the reboot. "+
							"current progress: %d, current status: %s, current version: %s, requested version: %s",
						state.Fwupdate.Progress.ValueInt64(), state.Fwupdate.Status, state.DeviceVersion, state.TargetVersion),
					)

					// last part, if we detect the running version is the same as the requested version for any reason
					// return a success
				} else if rebootDone {
					if state.DeviceVersion != state.TargetVersion {
						diags.AddError(
							"Error during device upgrade",
							fmt.Sprintf(
								"Upgrade process finished but the running firmware version reported by the device (%s) "+
									"is different from the requested target firmware version (%s)",
								state.DeviceVersion.ValueString(), state.TargetVersion.ValueString(),
							),
						)
					}
					return state, diags
				}
			}
			time.Sleep(time.Duration(syncUpgradeRefreshInterval) * time.Second)
		}
	}
}

func checkUpgradeProgress(
	state resource_upgrade_device.UpgradeDeviceModel,
	startTime time.Time,
	deviceUptime int,
) (bool, bool, bool) {
	var uploadDone = false
	var upgradeDone = false
	var rebootDone = false

	if state.Fwupdate.Progress.ValueInt64() >= 70 {
		uploadDone = true
		if state.Fwupdate.Progress.ValueInt64() == 100 {
			upgradeDone = true
			// we are not checking the connection status (connected/disconnected) because we have to wait
			// for the update of the device info to get the "new" running version
			// the easiest way to do it is to check the uptime value, which is refreshed at the same time
			// as the running firmware version
			// once the uptime value is updated, we consider the upgrade process finished and will compare
			// the target_version with the device_version
			if float64(deviceUptime) > 0 && time.Since(startTime).Seconds() > float64(deviceUptime) {
				rebootDone = true
			}
		}
	}

	return uploadDone, upgradeDone, rebootDone
}
