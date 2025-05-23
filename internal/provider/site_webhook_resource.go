package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_webhook"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &siteWebhookResource{}
	_ resource.ResourceWithConfigure   = &siteWebhookResource{}
	_ resource.ResourceWithImportState = &siteWebhookResource{}
)

func NewSiteWebhook() resource.Resource {
	return &siteWebhookResource{}
}

type siteWebhookResource struct {
	client mistapi.ClientInterface
}

func (r *siteWebhookResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Webhook client")
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
func (r *siteWebhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_webhook"
}

func (r *siteWebhookResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This resource manages Site Webhooks.\n\n" +
			"A Site Webhook is a configuration that allows real-time events and data from a specific site to be pushed to a provided url.  \n" +
			"It enables the collection of information about various topics such as device events, alarms, audits, client sessions and location updates at the site level.  \n" +
			"The Webhook can be set up and customized using the Mist API, allowing users to receive and analyze specific data from a particular site.",
		Attributes: resource_site_webhook.SiteWebhookResourceSchema(ctx).Attributes,
	}
}

func (r *siteWebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Webhook Create")
	var plan, state resource_site_webhook.SiteWebhookModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	webhook, diags := resource_site_webhook.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.SitesWebhooks().CreateSiteWebhook(ctx, siteId, &webhook)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to create the Webhook. %s", apiErr),
		)
		return
	}

	state, diags = resource_site_webhook.SdkToTerraform(ctx, &data.Data)
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

func (r *siteWebhookResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_webhook.SiteWebhookModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Webhook Read: webhook_id "+state.Id.ValueString())
	httpr, err := r.client.SitesWebhooks().GetSiteWebhook(ctx, siteId, webhookId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_site_webhook\" resource",
			"Unable to get the Webhook, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_webhook.SdkToTerraform(ctx, &httpr.Data)
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

func (r *siteWebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_webhook.SiteWebhookModel
	tflog.Info(ctx, "Starting Webhook Update")

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

	webhook, diags := resource_site_webhook.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Webhook Update for Webhook "+state.Id.ValueString())
	data, err := r.client.SitesWebhooks().UpdateSiteWebhook(ctx, siteId, webhookId, &webhook)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to update the Webhook. %s", apiErr),
		)
		return
	}

	state, diags = resource_site_webhook.SdkToTerraform(ctx, &data.Data)
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

func (r *siteWebhookResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_webhook.SiteWebhookModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesWebhooks().DeleteSiteWebhook(ctx, siteId, webhookId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_site_webhook\" resource",
				fmt.Sprintf("Unable to delete the Webhook. %s", apiErr),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_site_webhook\" resource",
			"Unable to delete the Webhook, unexpected error: "+err.Error(),
		)
		return
	}
	resp.State.RemoveResource(ctx)
}

func (r *siteWebhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site_webhook\" resource",
			"import \"id\" format must be \"{site_id}.{webhook_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{webhook_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_site_webhook\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{webhook_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
