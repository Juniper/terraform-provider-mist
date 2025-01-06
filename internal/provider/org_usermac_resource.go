package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_usermac"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgNacEndpointResource{}
	_ resource.ResourceWithConfigure   = &orgNacEndpointResource{}
	_ resource.ResourceWithImportState = &orgNacEndpointResource{}
)

func NewOrgNacEndpointResource() resource.Resource {
	return &orgNacEndpointResource{}
}

type orgNacEndpointResource struct {
	client mistapi.ClientInterface
}

func (r *orgNacEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgNacEndpoint client")
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

func (r *orgNacEndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nac_endpoint"
}
func (r *orgNacEndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This resource manages the NAC Endpoints (User MACs).\n\n" +
			"NAC Endpoints (User MACs) provide a database of endpoints identified by their MAC addresses. " +
			"They can be used assign each endpoint with various attributes, such as name, VLAN, role and client label.  " +
			"Once an endpoint is labeled, the label name can be used to create `mist_org_nactag` resource as match criteria.\n\n" +
			"The `mist_org_nactag` resource can be used to create Tags regrouping one or multiple endpoint MAC Addresses, " +
			"but the use of the User MACs provides additional features:\n" +
			"* possitility to assign specific attributes, like a Name, a Radius Group, a VLAN ID, ...\n" +
			"* possibility to assign one or multiple Tags (Labels) to a User MAC\n" +
			"* improved management for large list of MAC Addresses",
		Attributes: resource_org_usermac.OrgUsermacResourceSchema(ctx).Attributes,
	}
}

func (r *orgNacEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgNacEndpoint Create")
	var plan, state resource_org_usermac.OrgUsermacModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nacEndpoint, diags := resource_org_usermac.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting OrgNacEndpoint Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsUserMACs().CreateOrgUserMacs(ctx, orgId, &nacEndpoint)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to create the NacEndpoint. %s", api_err),
		)
		return
	}

	state, diags = resource_org_usermac.SdkToTerraform(ctx, &data.Data, &orgId)
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

func (r *orgNacEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_usermac.OrgUsermacModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgNacEndpoint Read: nac_endpoint_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	nacEndpointId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.OrgsUserMACs().GetOrgUserMac(ctx, orgId, nacEndpointId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_nac_endpoint\" resource",
			"Unable to get the NacEndpoint, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_usermac.SdkToTerraform(ctx, &httpr.Data, &orgId)
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

func (r *orgNacEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_usermac.OrgUsermacModel
	tflog.Info(ctx, "Starting OrgNacEndpoint Update")

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

	nacEndpoint, diags := resource_org_usermac.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgNacEndpoint Update for OrgNacEndpoint "+plan.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	serviceId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsUserMACs().UpdateOrgUserMac(ctx, orgId, serviceId, &nacEndpoint)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to update the NacEndpoint. %s", api_err),
		)
		return
	}

	state, diags = resource_org_usermac.SdkToTerraform(ctx, &data.Data, &orgId)
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

func (r *orgNacEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_usermac.OrgUsermacModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgNacEndpoint Delete: nac_endpoint_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	nacEndpointId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsUserMACs().DeleteOrgUserMac(ctx, orgId, nacEndpointId)
	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to delete the NacEndpoint. %s", api_err),
		)
		return
	}
}

func (r *orgNacEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_nac_endpoint\" resource",
			"import \"id\" format must be \"{org_id}.{nac_endpoint_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{nac_endpoint_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_nac_endpoint\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{nac_endpoint_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
