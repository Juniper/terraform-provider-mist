package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/tmunzer/mistapi-go/mistapi"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_upgrade_device"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
func (r *upgradeDeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_upgrade_device"
}

func (r *upgradeDeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages Org SSO Configuration.\n\n" +
			"Org SSO, or Single Sign-On, is a method of authentication that allows users to securely log in to multiple applications " +
			"and websites with a single set of login credentials.  \n" +
			"It involves integrating the Org portal with an Identity Provider (IdP) using the Security Assertion Markup Language (SAML) framework.  \n" +
			"This enables users to authenticate themselves through their corporate IdP, eliminating the need to remember separate " +
			"passwords or enter credentials each time they access the Org portal.",
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

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *upgradeDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *upgradeDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	return
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

	upgrade, diags := resource_upgrade_device.TerraformToSdk(ctx, &plan)
	diags.Append(diags...)
	if diags.HasError() {
		return state, diags
	}
	// Start upgrade
	data, err := r.client.UtilitiesUpgrade().UpgradeDevice(ctx, siteId, deviceId, upgrade)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		diags.AddError(
			"Error creating \"mist_upgrade_device\" resource",
			fmt.Sprintf("Unable to Upgrade the device. %s", api_err),
		)
		return state, diags
	} else if data.Data.Status == "device already running version" {
		diags.AddWarning(
			"Unable to upgrade the device",
			fmt.Sprintf("Device is already running the version %s", plan.Version.ValueString()),
		)
		state, diags = resource_upgrade_device.SdkToTerraform(ctx, plan, &data.Data)
		diags.Append(diags...)
		if diags.HasError() {
			return state, diags
		}
	} else {
		state, diags = resource_upgrade_device.SdkToTerraform(ctx, plan, &data.Data)
		diags.Append(diags...)
		if diags.HasError() {
			return state, diags
		}

		state, diags = r.refreshFwUpdate(
			ctx,
			state,
			plan.SyncUpgrade.ValueBool(),
			plan.SyncUpgradeStartTimeout.ValueInt64(),
			plan.SyncUpgradeRefreshInterval.ValueInt64(),
			plan.SyncUpgradeTimeout.ValueInt64(),
		)
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

	var done = false
	var previousStatusId int64 = -1
	var startTime = time.Now()
	var retry = 0
	var maxRetry = 3

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
		tflog.Debug(ctx, fmt.Sprintf(
			"upgrade check. "+
				"SyncUpgradeStartTimeout: %d, "+
				"sleepTime: %d, "+
				"previousStatusId: %d, "+
				"startTime: %s, "+
				"done: %t, ",
			syncUpgradeStartTimeout, syncUpgradeRefreshInterval, previousStatusId, startTime, done))

		data, err := r.client.SitesStatsDevices().GetSiteDeviceStats(ctx, siteId, deviceId, fields)

		api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
		if api_err != "" {
			diags.AddError(
				"Error creating \"mist_upgrade_device\" resource",
				fmt.Sprintf("Unable to create Upgrade the device. %s", api_err),
			)
			if retry < maxRetry {
				retry += 1
				time.Sleep(2 * time.Second)
			} else {
				return state, diags
			}
		}

		retry = 0

		state, diags = resource_upgrade_device.DeviceStatSdkToTerraform(ctx, state, &data)
		diags.Append(diags...)
		if diags.HasError() {
			diags.AddError(
				"Error when upgrading the device",
				fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
			)
			return state, diags
		}

		if !syncUpgrade {
			tflog.Debug(ctx, "upgrade check is async, do not wait")
			return state, diags
		} else {
			if previousStatusId == -1 && time.Since(startTime).Seconds() > float64(syncUpgradeStartTimeout) {
				diags.AddError(
					"Error creating \"mist_upgrade_device\" resource",
					"Upgrade start Timeout reached",
				)
				return state, diags
			} else if time.Since(startTime).Seconds() > float64(syncUpgradeTimeout) {
				diags.AddError(
					"Error creating \"mist_upgrade_device\" resource",
					"Upgrade end Timeout reached",
				)
				return state, diags
			} else if previousStatusId == -1 {
				tflog.Debug(ctx, "upgrade check is sync, wait for the upgrade to start")
				previousStatusId = state.Fwupdate.StatusId.ValueInt64()
			} else {
				tflog.Debug(ctx, fmt.Sprintf(
					"upgrade check is sync, wait for the end of the upgrade. current progress: %d, current status: %s",
					state.Fwupdate.Progress.ValueInt64(), state.Fwupdate.Status),
				)
				if state.Fwupdate.Status.ValueString() == "success" {
					return state, diags
				}
			}
			time.Sleep(time.Duration(syncUpgradeRefreshInterval) * time.Second)
		}
	}
}
