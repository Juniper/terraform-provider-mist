package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_ap"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgDeviceprofileApResource{}
	_ resource.ResourceWithConfigure   = &orgDeviceprofileApResource{}
	_ resource.ResourceWithImportState = &orgDeviceprofileApResource{}
)

func NewOrgDeviceprofileAp() resource.Resource {
	return &orgDeviceprofileApResource{}
}

type orgDeviceprofileApResource struct {
	client mistapi.ClientInterface
}

func (r *orgDeviceprofileApResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceprofileAp client")
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
func (r *orgDeviceprofileApResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofile_ap"
}

func (r *orgDeviceprofileApResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource manages the AP Device Profiles.\n" +
			"AP Device profiles for aps are used to specify a configuration that can be applied to a select set of aps from any site in the organization. " +
			"They allow for efficient application of configurations based on ap groups, wlan groups, RF settings, and sites. " +
			"Device profiles enable various use cases such as activating ethernet passthrough, applying different rf settings, applying mesh configuration, " +
			"activating specific features like esl or vble, and more.\n\n" +
			"The AP Devide Profile can be assigned to a gateway with the `mist_org_deviceprofile_assign` resource.",
		Attributes: resource_org_deviceprofile_ap.OrgDeviceprofileApResourceSchema(ctx).Attributes,
	}
}

func (r *orgDeviceprofileApResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceprofileAp Create")
	var plan, state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceprofileAp, diags := resource_org_deviceprofile_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsDeviceProfiles().CreateOrgDeviceProfiles(ctx, orgId, &deviceprofileAp)
	if data.Response.StatusCode != 200 {

		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_org_deviceprofile_ap\" resource",
				fmt.Sprintf("Unable to create the AP Device Profile. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistDeviceprofileAp := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mistDeviceprofileAp)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mistDeviceprofileAp)
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

func (r *orgDeviceprofileApResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileApid, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Read: deviceprofile_ap_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsDeviceProfiles().GetOrgDeviceProfile(ctx, orgId, deviceprofileApid)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_deviceprofile_ap\" resource",
			"Unable to get the AP Device Profile, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
	mistDeviceprofileAp := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mistDeviceprofileAp)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mistDeviceprofileAp)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgDeviceprofileApResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_deviceprofile_ap.OrgDeviceprofileApModel
	tflog.Info(ctx, "Starting DeviceprofileAp Update")

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

	deviceprofileAp, diags := resource_org_deviceprofile_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileApid, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Update for DeviceprofileAp "+state.Id.ValueString())
	data, err := r.client.OrgsDeviceProfiles().UpdateOrgDeviceProfile(ctx, orgId, deviceprofileApid, &deviceprofileAp)

	if data.Response.StatusCode != 200 {

		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_org_deviceprofile_ap\" resource",
				fmt.Sprintf("Unable to update the AP Device Profile. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistDeviceprofileAp := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mistDeviceprofileAp)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mistDeviceprofileAp)
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

func (r *orgDeviceprofileApResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileApid, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Delete: deviceprofile_ap_id "+state.Id.ValueString())
	data, err := r.client.OrgsDeviceProfiles().DeleteOrgDeviceProfile(ctx, orgId, deviceprofileApid)
	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to delete the AP Device Profile. %s", apiErr),
		)
		return
	}
}

func (r *orgDeviceprofileApResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_deviceprofile_ap\" resource",
			"import \"id\" format must be \"{org_id}.{deviceprofile_ap_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{deviceprofile_ap_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_deviceprofile_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{deviceprofile_ap_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
