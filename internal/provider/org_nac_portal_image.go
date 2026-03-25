package provider

import (
	"context"
	"fmt"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nac_portal_image"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNacPortalImageResource{}
	_ resource.ResourceWithConfigure = &orgNacPortalImageResource{}
)

func NewOrgNacPortalImage() resource.Resource {
	return &orgNacPortalImageResource{}
}

type orgNacPortalImageResource struct {
	client mistapi.ClientInterface
}

func (r *orgNacPortalImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org NAC Portal Image client")
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

func (r *orgNacPortalImageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nac_portal_image"
}

func (r *orgNacPortalImageResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This resource is used to upload a NAC Portal background image.",
		Attributes:          resource_org_nac_portal_image.OrgNacPortalImageResourceSchema(ctx).Attributes,
	}
}

func (r *orgNacPortalImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org NAC Portal Image Create")
	var plan, state resource_org_nac_portal_image.OrgNacPortalImageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(plan.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"file\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json = ""

	data, err := r.client.OrgsNACPortals().UploadOrgNacPortalImage(ctx, orgId, nacPortalId, &file, &json)

	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Unable to create the Portal Image. %s", apiErr),
		)
		return
	}

	state.File = plan.File
	state.OrgId = plan.OrgId
	state.NacportalId = plan.NacportalId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgNacPortalImageResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {

}

func (r *orgNacPortalImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_nac_portal_image.OrgNacPortalImageModel
	tflog.Info(ctx, "Starting Org NAC Portal Image Update")

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(plan.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"file\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json = ""

	data, err := r.client.OrgsNACPortals().UploadOrgNacPortalImage(ctx, orgId, nacPortalId, &file, &json)

	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Unable to update the Portal Image. %s", apiErr),
		)
		return
	}

	state.File = plan.File
	state.OrgId = plan.OrgId
	state.NacportalId = plan.NacportalId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgNacPortalImageResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_nac_portal_image.OrgNacPortalImageModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(state.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsNACPortals().DeleteOrgNacPortalImage(ctx, orgId, nacPortalId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_nac_portal_image\" resource",
				"Could not delete Portal Image, unexpected error: "+err.Error(),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_nac_portal_image\" resource",
			"Unable to delete the Portal Image, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}
