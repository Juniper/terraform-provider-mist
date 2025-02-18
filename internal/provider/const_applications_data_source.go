package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_applications"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constApplicationsDataSource)(nil)

func NewConstApplicationsDataSource() datasource.DataSource {
	return &constApplicationsDataSource{}
}

type constApplicationsDataSource struct {
	client mistapi.ClientInterface
}

func (d *constApplicationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constant Applications Datasource client")
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
func (d *constApplicationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_applications"
}

func (d *constApplicationsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the of ConstApplications.\n\n" +
			"This information can be used as `apps` in:\n" +
			"* `mist_org_service` resource\n" +
			"* `mist_site_setting` resource (`mist_site_setting.gateway_mgmt.app_probing.apps`)\n" +
			"* `mist_org_setting` resource (`mist_org_setting.gateway_mgmt.app_probing.apps`)",
		Attributes: datasource_const_applications.ConstApplicationsDataSourceSchema(ctx).Attributes,
	}
}

func (d *constApplicationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_applications.ConstApplicationsModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListApplications(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Unable to get the AP Stats, unexpected error: "+err.Error(),
		)
		return
	}
	constApplications, diags := datasource_const_applications.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_applications"), constApplications); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
