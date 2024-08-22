package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sitegroup"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgSiteGroupResource{}
	_ resource.ResourceWithConfigure   = &orgSiteGroupResource{}
	_ resource.ResourceWithImportState = &orgSiteGroupResource{}
)

func NewOrgSiteGroupResource() resource.Resource {
	return &orgSiteGroupResource{}
}

type orgSiteGroupResource struct {
	client mistapi.ClientInterface
}

func (r *orgSiteGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SiteGroup client")
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
func (r *orgSiteGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sitegroup"
}

func (r *orgSiteGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource managed the Org Site Groups (sitegroups)." +
			"A site group is a feature that allows users to group multiple sites together based on regions, functions, or other parameters for efficient management of devices. " +
			"Sites can exist in multiple groups simultaneously, and site groups can be used to ensure consistent settings, manage administrator access, and apply specific templates to groups of sites.",
		Attributes: resource_org_sitegroup.OrgSitegroupResourceSchema(ctx).Attributes,
	}
}

func (r *orgSiteGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting SiteGroup Create")
	var plan, state resource_org_sitegroup.OrgSitegroupModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sitegroup, diags := resource_org_sitegroup.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsSitegroups().CreateOrgSiteGroup(ctx, orgId, sitegroup)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating SiteGroup",
			"Could not create SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sitegroup.SdkToTerraform(&data.Data)
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

func (r *orgSiteGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_sitegroup.OrgSitegroupModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	sitegroupId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting SiteGroup Read: sitegroup_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsSitegroups().GetOrgSiteGroup(ctx, orgId, sitegroupId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting SiteGroup",
			"Could not get SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_sitegroup.SdkToTerraform(&httpr.Data)
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

func (r *orgSiteGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_sitegroup.OrgSitegroupModel
	tflog.Info(ctx, "Starting SiteGroup Update")

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

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	sitegroupId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	sitegroup_name := models.NameString{}
	sitegroup_name.Name = plan.Name.ValueStringPointer()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting SiteGroup Update for Site "+plan.Id.ValueString())
	data, err := r.client.OrgsSitegroups().UpdateOrgSiteGroup(ctx, orgId, sitegroupId, &sitegroup_name)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating SiteGroup",
			"Could not update SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sitegroup.SdkToTerraform(&data.Data)
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

func (r *orgSiteGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_sitegroup.OrgSitegroupModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	sitegroupId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_sitegroup\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting SiteGroup Delete: sitegroup_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsSitegroups().DeleteOrgSiteGroup(ctx, orgId, sitegroupId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting SiteGroup",
			"Could not delete SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgSiteGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
