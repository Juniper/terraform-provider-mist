package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxtunnel"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgMxtunnelResource{}
	_ resource.ResourceWithConfigure   = &orgMxtunnelResource{}
	_ resource.ResourceWithImportState = &orgMxtunnelResource{}
)

func NewOrgMxtunnelResource() resource.Resource {
	return &orgMxtunnelResource{}
}

type orgMxtunnelResource struct {
	client mistapi.ClientInterface
}

func (r *orgMxtunnelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgMxtunnel Resource client")
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

func (r *orgMxtunnelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_mxtunnel"
}

func (r *orgMxtunnelResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages Mist Tunnels in the Mist Organization.\n\n" +
			"A Mist Tunnel (MxTunnel) defines tunneling configuration used to carry AP user VLANs to Mist Edge clusters.",
		Attributes: resource_org_mxtunnel.OrgMxtunnelResourceSchema(ctx).Attributes,
	}
}

func (r *orgMxtunnelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgMxtunnel Create")
	var plan, state resource_org_mxtunnel.OrgMxtunnelModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxtunnel\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	mxtunnel, diags := resource_org_mxtunnel.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating OrgMxtunnel in Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsMxTunnels().CreateOrgMxTunnel(ctx, orgId, mxtunnel)
	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_org_mxtunnel\" resource",
				fmt.Sprintf("Unable to create the MxTunnel. %s", apiErr),
			)
			return
		}
	}

	body, err := io.ReadAll(data.Response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read API response body", err.Error())
		return
	}
	mistMxtunnel := models.Mxtunnel{}
	if err = json.Unmarshal(body, &mistMxtunnel); err != nil {
		resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
		return
	}

	state, diags = resource_org_mxtunnel.SdkToTerraform(ctx, &mistMxtunnel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *orgMxtunnelResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_mxtunnel.OrgMxtunnelModel

	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxtunnel Read: mxtunnel_id "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxtunnelId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.OrgsMxTunnels().GetOrgMxTunnel(ctx, orgId, mxtunnelId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_mxtunnel\" resource",
			"Unable to get the MxTunnel, unexpected error: "+err.Error(),
		)
		return
	}

	body, err := io.ReadAll(httpr.Response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read API response body", err.Error())
		return
	}
	mistMxtunnel := models.Mxtunnel{}
	if err = json.Unmarshal(body, &mistMxtunnel); err != nil {
		resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
		return
	}

	state, diags := resource_org_mxtunnel.SdkToTerraform(ctx, &mistMxtunnel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *orgMxtunnelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_mxtunnel.OrgMxtunnelModel
	tflog.Info(ctx, "Starting OrgMxtunnel Update")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	mxtunnel, diags := resource_org_mxtunnel.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxtunnel Update for MxTunnel "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxtunnelId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsMxTunnels().UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, mxtunnel)
	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_org_mxtunnel\" resource",
				fmt.Sprintf("Unable to update the MxTunnel. %s", apiErr),
			)
			return
		}
	}

	body, err := io.ReadAll(data.Response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read API response body", err.Error())
		return
	}
	mistMxtunnel := models.Mxtunnel{}
	if err = json.Unmarshal(body, &mistMxtunnel); err != nil {
		resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
		return
	}

	state, diags = resource_org_mxtunnel.SdkToTerraform(ctx, &mistMxtunnel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *orgMxtunnelResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_mxtunnel.OrgMxtunnelModel

	resp.Diagnostics.Append(resp.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxtunnel Delete: mxtunnel_id "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxtunnelId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.OrgsMxTunnels().DeleteOrgMxTunnel(ctx, orgId, mxtunnelId)
	if httpr.StatusCode != 200 && httpr.StatusCode != 404 {
		apiErr := mistapierror.ProcessApiError(httpr.StatusCode, httpr.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_mxtunnel\" resource",
				fmt.Sprintf("Unable to delete the MxTunnel. %s", apiErr),
			)
			return
		}
	}
}

func (r *orgMxtunnelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Import \"id\" format must be \"{org_id}.{mxtunnel_id}\", got %s", req.ID),
		)
		return
	}

	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxtunnel_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxtunnel\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxtunnel_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
