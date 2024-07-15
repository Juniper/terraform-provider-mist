package provider

import (
	"context"
	"fmt"
	"terraform-provider-mist/internal/resource_org"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgResource{}
	_ resource.ResourceWithConfigure = &orgResource{}
)

func NewOrgResource() resource.Resource {
	return &orgResource{}
}

type orgResource struct {
	client mistapi.ClientInterface
}

func (r *orgResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *models.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *orgResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org"
}

func (r *orgResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org.OrgResourceSchema(ctx)
}

func (r *orgResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org Create")
	var plan, state resource_org.OrgModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	org, diags := resource_org.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := r.client.Orgs().CreateOrg(ctx, org)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating org",
			"Could not create org, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org.SdkToTerraform(ctx, data.Data)
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

func (r *orgResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org.OrgModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting Org Read: org_id "+state.Id.ValueString())
	data, err := r.client.Orgs().GetOrg(ctx, orgId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org",
			"Could not get org, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org.SdkToTerraform(ctx, data.Data)
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

func (r *orgResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org.OrgModel
	tflog.Info(ctx, "Starting Org Update")

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

	orgId := uuid.MustParse(state.Id.ValueString())
	org, diags := resource_org.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Org Update for Org "+orgId.String())
	data, err := r.client.Orgs().UpdateOrg(ctx, orgId, org)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating site",
			"Could not update site, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org.SdkToTerraform(ctx, data.Data)
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

func (r *orgResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org.OrgModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.Id.ValueString())

	tflog.Info(ctx, "Starting Org Delete: org_id "+state.Id.ValueString())
	_, err := r.client.Orgs().DeleteOrg(ctx, orgId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting org",
			"Could not delete org, unexpected error: "+err.Error(),
		)
		return
	}
}
