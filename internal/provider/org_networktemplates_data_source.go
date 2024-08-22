package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_networktemplates"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgNetworktemplatesDataSource)(nil)

func NewOrgNetworktemplatesDataSource() datasource.DataSource {
	return &orgNetworktemplatesDataSource{}
}

type orgNetworktemplatesDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgNetworktemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Network Templates Datasource client")
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
func (d *orgNetworktemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_networktemplates"
}

func (d *orgNetworktemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWired + "This data source provides the list of Org Network Templates (Switch templates)." +
			"A network template is a predefined configuration that provides a consistent and reusable set of network settings for devices within an organization. " +
			"It includes various parameters such as ip addressing, vlan configurations, routing protocols, security policies, and other network-specific settings. " +
			"Network templates simplify the deployment and management of switches by ensuring consistent configurations across multiple devices and sites. " +
			"They help enforce standardization, reduce human error, and streamline troubleshooting and maintenance tasks.",
		Attributes: datasource_org_networktemplates.OrgNetworktemplatesDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgNetworktemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_networktemplates.OrgNetworktemplatesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_networktemplates\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var page *int

	data, err := d.client.OrgsNetworkTemplates().ListOrgNetworkTemplates(ctx, orgId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_networktemplates.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_networktemplates"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
