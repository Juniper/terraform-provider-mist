package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_webhooks"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constWebhooksDataSource)(nil)

func NewConstWebhooksDataSource() datasource.DataSource {
	return &constWebhooksDataSource{}
}

type constWebhooksDataSource struct {
	client mistapi.ClientInterface
}

func (d *constWebhooksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constant Webhooks Datasource client")
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
func (d *constWebhooksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_webhooks"
}

func (d *constWebhooksDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the list of Webhook Topics.\n\n" +
			"This information can be used to configure webhooks at the Org level (`mist_org_webhook` resource) " +
			"or at the Site level (`mist_site_webhook` resource).\n\n" +
			"-> Only the Webhook topics with `for_org`==` true` are supported at the Org level.",
		Attributes: datasource_const_webhooks.ConstWebhooksDataSourceSchema(ctx).Attributes,
	}
}

func (d *constWebhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_webhooks.ConstWebhooksModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListWebhookTopics(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Unable to get the AP Stats, unexpected error: "+err.Error(),
		)
		return
	}
	constWebhooks, diags := datasource_const_webhooks.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_webhooks"), constWebhooks); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
