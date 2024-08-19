package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_traffic_types"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constTrafficTypesDataSource)(nil)

func NewConstTrafficTypesDataSource() datasource.DataSource {
	return &constTrafficTypesDataSource{}
}

type constTrafficTypesDataSource struct {
	client mistapi.ClientInterface
}

func (d *constTrafficTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constant Traffic Types Datasource client")
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
func (d *constTrafficTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_traffic_types"
}

func (d *constTrafficTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the of ConstTrafficTypes." +
			"This information can be used to define the Country in the RF templates (`mist_org_rftemplate`)",
		Attributes: datasource_const_traffic_types.ConstTrafficTypesDataSourceSchema(ctx).Attributes,
	}
}

func (d *constTrafficTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_traffic_types.ConstTrafficTypesModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListTrafficTypes(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}
	constTrafficTypes, diags := datasource_const_traffic_types.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_traffic_types"), constTrafficTypes); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
