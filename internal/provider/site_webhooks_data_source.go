package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_site_webhooks"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*siteWebhooksDataSource)(nil)

func NewSiteWebhooksDataSource() datasource.DataSource {
	return &siteWebhooksDataSource{}
}

type siteWebhooksDataSource struct {
	client mistapi.ClientInterface
}

func (d *siteWebhooksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site Webhooks Datasource client")
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

	d.client = client
}
func (d *siteWebhooksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_webhooks"
}

func (d *siteWebhooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This data source provides the list of Site Webhooks.",
		Attributes:          datasource_site_webhooks.SiteWebhooksDataSourceSchema(ctx).Attributes,
	}
}

func (d *siteWebhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_site_webhooks.SiteWebhooksModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(ds.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_webhooks\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var page *int

	data, err := d.client.SitesWebhooks().ListSiteWebhooks(ctx, siteId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_site_webhooks.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("site_webhooks"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
