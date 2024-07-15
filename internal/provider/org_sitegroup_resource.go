package provider

import (
	"context"
	"fmt"
	"terraform-provider-mist/internal/resource_org_sitegroup"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgSiteGroupResource{}
	_ resource.ResourceWithConfigure = &orgSiteGroupResource{}
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
	resp.Schema = resource_org_sitegroup.OrgSitegroupResourceSchema(ctx)
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
	orgId := uuid.MustParse(plan.OrgId.ValueString())
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

	orgId := uuid.MustParse(state.OrgId.ValueString())
	sitegroupId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting SiteGroup Read: sitegroup_id "+state.Id.ValueString())
	data, err := r.client.OrgsSitegroups().GetOrgSiteGroup(ctx, orgId, sitegroupId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting SiteGroup",
			"Could not get SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_sitegroup.SdkToTerraform(&data.Data)
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

	orgId := uuid.MustParse(state.OrgId.ValueString())
	sitegroupId := uuid.MustParse(state.Id.ValueString())
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

	orgId := uuid.MustParse(state.OrgId.ValueString())
	sitegroupId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting SiteGroup Delete: sitegroup_id "+state.Id.ValueString())
	_, err := r.client.OrgsSitegroups().DeleteOrgSiteGroup(ctx, orgId, sitegroupId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting SiteGroup",
			"Could not delete SiteGroup, unexpected error: "+err.Error(),
		)
		return
	}
}
