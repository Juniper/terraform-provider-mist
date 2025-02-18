package provider

import (
	"context"
	"fmt"
	"sync"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_            resource.Resource                = &siteResource{}
	_            resource.ResourceWithConfigure   = &siteResource{}
	_            resource.ResourceWithImportState = &siteResource{}
	resourceLock sync.Mutex
)

func NewSiteResource() resource.Resource {
	return &siteResource{}
}

type siteResource struct {
	client mistapi.ClientInterface
}

func (r *siteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site client")
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
func (r *siteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

func (r *siteResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This resource manages the Site basic information.\n\n" +
			"This resource can be used to assign templates to a site, or to change basic information (e.g. Site Address)",
		Attributes: resource_site.SiteResourceSchema(ctx).Attributes,
	}
}

func (r *siteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state resource_site.SiteModel

	// the resourceLock is required to avoid concurent sitegroup assign/unassign to
	// different sites. If the same sitegroup is assigned to multiple sites at the
	// same time, the `site_ids` atribute may not be updated correctly in the Mist
	// backend
	// by using this approach, the provider is creating/updating/deleting each sites
	// sequentially
	resourceLock.Lock()

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, diags := resource_site.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_site\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting Site Create for Org "+plan.OrgId.ValueString())

	data, err := r.client.OrgsSites().CreateOrgSite(ctx, orgId, site)
	// The API call has been done, we can unlock the resource
	resourceLock.Unlock()

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site\" resource",
			fmt.Sprintf("Unable to create the Mist Site in the org \"%s\". %s", orgId.String(), apiErr),
		)
		return
	}

	state, diags = resource_site.SdkToTerraform(&data.Data)
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

func (r *siteResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site.SiteModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting Site Read: site_id "+state.Id.ValueString())
	httpr, err := r.client.Sites().GetSiteInfo(ctx, siteId)
	if httpr.Response.StatusCode == 404 || httpr.Response.StatusCode == 403 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_site\" resource",
			"Unable to get the Site, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site.SdkToTerraform(&httpr.Data)
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

func (r *siteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site.SiteModel

	// the resourceLock is required to avoid concurent sitegroup assign/unassign to
	// different sites. If the same sitegroup is assigned to multiple sites at the
	// same time, the `site_ids` atribute may not be updated correctly in the Mist
	// backend
	// by using this approach, the provider is creating/updating/deleting each sites
	// sequentially
	resourceLock.Lock()

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

	site, diags := resource_site.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Site Update for Site "+state.Id.ValueString())
	data, err := r.client.Sites().UpdateSiteInfo(ctx, siteId, site)
	// The API call has been done, we can unlock the resource
	resourceLock.Unlock()

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_site\" resource",
			fmt.Sprintf("Unable to update the Mist Site. %s", apiErr),
		)
		return
	}

	state, diags = resource_site.SdkToTerraform(&data.Data)
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

func (r *siteResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site.SiteModel

	// the resourceLock is required to avoid concurent sitegroup assign/unassign to
	// different sites. If the same sitegroup is assigned to multiple sites at the
	// same time, the `site_ids` atribute may not be updated correctly in the Mist
	// backend
	// by using this approach, the provider is creating/updating/deleting each sites
	// sequentially
	resourceLock.Lock()

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Site Delete: site_id "+state.Id.ValueString())
	data, err := r.client.Sites().DeleteSite(ctx, siteId)
	// The API call has been done, we can unlock the resource
	resourceLock.Unlock()

	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_site\" resource",
			fmt.Sprintf("Unable to delete the Site. %s", apiErr),
		)
		return
	}
}

func (r *siteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" must be a valid Site Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
