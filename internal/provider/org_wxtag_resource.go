package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wxtag"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgWxTagResource{}
	_ resource.ResourceWithConfigure   = &orgWxTagResource{}
	_ resource.ResourceWithImportState = &orgWxTagResource{}
)

func NewOrgWxTag() resource.Resource {
	return &orgWxTagResource{}
}

type orgWxTagResource struct {
	client mistapi.ClientInterface
}

func (r *orgWxTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WxTag client")
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
func (r *orgWxTagResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wxtag"
}

func (r *orgWxTagResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource manages the Org WxLan tags (labels).\n\n" +
			"A WxTag is a label or tag used in the mist system to classify and categorize applications, " +
			"users, and resources for the purpose of creating policies and making network management decisions. " +
			"They can be used \n" +
			"* within the Org WxRules to create filtering rules:\n" +
			"  * `mist_org_wxrule.dst_allow_wxtags`\n" +
			"  * `mist_org_wxrule.dst_deny_wxtags`\n" +
			"  * `mist_org_wxrule.dst_wxtags`\n" +
			"  * `mist_org_wxrule.src_wxtags`\n" +
			"* within the Org WxRules to assign specific VLAN:\n" +
			"  * `mist_org_wxrule.apply_tags`\n" +
			"* in the WLANs configuration to assign a WLAN to specific APs:\n" +
			"  * `mist_org_wlan.wxtag_ids`\n" +
			"* to identify unknown application used by Wi-Fi clients",
		Attributes: resource_org_wxtag.OrgWxtagResourceSchema(ctx).Attributes,
	}
}

func (r *orgWxTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxTag Create")
	var plan, state resource_org_wxtag.OrgWxtagModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wxtag, diags := resource_org_wxtag.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsWxTags().CreateOrgWxTag(ctx, orgId, wxtag)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to create the WxTag. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *orgWxTagResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_wxtag.OrgWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Read: wxtag_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	wxtagId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.OrgsWxTags().GetOrgWxTag(ctx, orgId, wxtagId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_wxtag\" resource",
			"Unable to get the WxTag, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_wxtag.SdkToTerraform(ctx, httpr.Data)
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

func (r *orgWxTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_wxtag.OrgWxtagModel
	tflog.Info(ctx, "Starting WxTag Update")

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

	wxtag, diags := resource_org_wxtag.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Update for WxTag "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	wxtagId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsWxTags().UpdateOrgWxTag(ctx, orgId, wxtagId, wxtag)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to update the WxTag. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *orgWxTagResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wxtag.OrgWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Delete: wxtag_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	wxtagId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsWxTags().DeleteOrgWxTag(ctx, orgId, wxtagId)
	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to delete the WxTag. %s", apiErr),
		)
		return
	}
}

func (r *orgWxTagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_wxtag\" resource",
			"import \"id\" format must be \"{org_id}.{wxtag_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{wxtag_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_wxtag\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{wxtag_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
