package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxcluster"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgMxclusterResource{}
	_ resource.ResourceWithConfigure   = &orgMxclusterResource{}
	_ resource.ResourceWithImportState = &orgMxclusterResource{}
)

func NewOrgMxclusterResource() resource.Resource {
	return &orgMxclusterResource{}
}

type orgMxclusterResource struct {
	client mistapi.ClientInterface
}

func (r *orgMxclusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgMxcluster Resource client")
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

func (r *orgMxclusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_mxcluster"
}

func (r *orgMxclusterResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages MxCluster (cluster of MxEdge devices) in the Mist Organization.\n\n" +
			"MxCluster provides high-availability configuration for groups of MxEdge devices.",
		Attributes: resource_org_mxcluster.OrgMxclusterResourceSchema(ctx).Attributes,
	}
}

func (r *orgMxclusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgMxcluster Create")
	var plan, state resource_org_mxcluster.OrgMxclusterModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxcluster\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	mxcluster, diags := resource_org_mxcluster.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating OrgMxcluster in Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsMxClusters().CreateOrgMxEdgeCluster(ctx, orgId, mxcluster)
	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_org_mxcluster\" resource",
				fmt.Sprintf("Unable to create the MxCluster. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistMxcluster := models.Mxcluster{}
	json.Unmarshal(body, &mistMxcluster)

	state, diags = resource_org_mxcluster.SdkToTerraform(ctx, &mistMxcluster)
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

func (r *orgMxclusterResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_mxcluster.OrgMxclusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxcluster Read: mxcluster_id "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxclusterId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.OrgsMxClusters().GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_mxcluster\" resource",
			"Unable to get the MxCluster, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
	mistMxcluster := models.Mxcluster{}
	json.Unmarshal(body, &mistMxcluster)

	state, diags = resource_org_mxcluster.SdkToTerraform(ctx, &mistMxcluster)
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

func (r *orgMxclusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_mxcluster.OrgMxclusterModel
	tflog.Info(ctx, "Starting OrgMxcluster Update")

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mxcluster, diags := resource_org_mxcluster.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxcluster Update for MxCluster "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxclusterId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsMxClusters().UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, mxcluster)

	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_org_mxcluster\" resource",
				fmt.Sprintf("Unable to update the MxCluster. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistMxcluster := models.Mxcluster{}
	json.Unmarshal(body, &mistMxcluster)

	state, diags = resource_org_mxcluster.SdkToTerraform(ctx, &mistMxcluster)
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

func (r *orgMxclusterResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_mxcluster.OrgMxclusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxcluster Delete: mxcluster_id "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxclusterId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.OrgsMxClusters().DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId)
	if httpr.StatusCode != 200 && httpr.StatusCode != 404 {
		apiErr := mistapierror.ProcessApiError(httpr.StatusCode, httpr.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_mxcluster\" resource",
				fmt.Sprintf("Unable to delete the MxCluster. %s", apiErr),
			)
			return
		}
	}
}

func (r *orgMxclusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Import \"id\" format must be \"{org_id}.{mxcluster_id}\", got %s", req.ID),
		)
		return
	}

	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxcluster_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxcluster\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxcluster_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
