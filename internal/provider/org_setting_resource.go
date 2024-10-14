package provider

import (
	"context"
	"fmt"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_setting"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgSettingResource{}
	_ resource.ResourceWithConfigure   = &orgSettingResource{}
	_ resource.ResourceWithImportState = &orgSettingResource{}
)

func NewOrgSettingResource() resource.Resource {
	return &orgSettingResource{}
}

type orgSettingResource struct {
	client mistapi.ClientInterface
}

func (r *orgSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgSetting client")
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
func (r *orgSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_setting"
}

func (r *orgSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource manages the Org Settings.\n" +
			"The Org Settings can be used to customize the Org configuration",
		Attributes: resource_org_setting.OrgSettingResourceSchema(ctx).Attributes,
	}
}

func (r *orgSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgSetting Create")
	var plan, state resource_org_setting.OrgSettingModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgSetting, diags := resource_org_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting OrgSetting Create for Org "+plan.OrgId.ValueString())
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to create the Org Setting. %s", api_err),
		)
		return
	}

	state, diags = resource_org_setting.SdkToTerraform(ctx, &data.Data)
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

func (r *orgSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_setting.OrgSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Read: orgSetting_id "+state.OrgId.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.OrgsSetting().GetOrgSettings(ctx, orgId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_setting\" resource",
			"Unable to get the orgSetting, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_setting.SdkToTerraform(ctx, &httpr.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_setting.OrgSettingModel
	tflog.Info(ctx, "Starting OrgSetting Update")

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

	orgSetting, diags := resource_org_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Update for OrgSetting "+state.OrgId.ValueString())
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to update the Org Setting. %s", api_err),
		)
		return
	}

	state, diags = resource_org_setting.SdkToTerraform(ctx, &data.Data)
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

func (r *orgSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_setting.OrgSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgSetting Delete: orgSetting_id "+state.OrgId.ValueString())
	orgSetting, e := resource_org_setting.DeleteTerraformToSdk(ctx)
	diags.Append(e...)
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsSetting().UpdateOrgSettings(ctx, orgId, orgSetting)
	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if data.Response.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to delete the Org Setting. %s", api_err),
		)
		return
	}
}

func (r *orgSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_setting\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" must be a valid Org Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), req.ID)...)
}
