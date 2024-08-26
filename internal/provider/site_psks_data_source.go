package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_site_psks"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*sitePsksDataSource)(nil)

func NewSitePsksDataSource() datasource.DataSource {
	return &sitePsksDataSource{}
}

type sitePsksDataSource struct {
	client mistapi.ClientInterface
}

func (d *sitePsksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site Psks Datasource client")
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
func (d *sitePsksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_psks"
}

func (d *sitePsksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance Psks." +
			"The Psks are used in the `service_policies` from the Gateway configuration and Gateway templates ",
		Attributes: datasource_site_psks.SitePsksDataSourceSchema(ctx).Attributes,
	}
}

func (d *sitePsksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_site_psks.SitePsksModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(ds.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_psks\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var name *string = ds.Name.ValueStringPointer()
	var limit *int = models.ToPointer(1000)
	var page *int
	var ssid *string = ds.Ssid.ValueStringPointer()
	var role *string = ds.Role.ValueStringPointer()

	data, err := d.client.SitesPsks().ListSitePsks(ctx, siteId, name, ssid, role, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_site_psks.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("site_psks"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
