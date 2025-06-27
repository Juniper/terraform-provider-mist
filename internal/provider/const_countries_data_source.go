package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_countries"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constCountriesDataSource)(nil)

func NewConstCountriesDataSource() datasource.DataSource {
	return &constCountriesDataSource{}
}

type constCountriesDataSource struct {
	client mistapi.ClientInterface
}

func (d *constCountriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constant Countries Datasource client")
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
func (d *constCountriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_countries"
}

func (d *constCountriesDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the of ConstCountries.\n\n" +
			"This information can be used to define the Country in the RF templates (`mist_org_rftemplate`)",
		Attributes: datasource_const_countries.ConstCountriesDataSourceSchema(ctx).Attributes,
	}
}

func (d *constCountriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_countries.ConstCountriesModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListCountryCodes(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Country Definitions",
			"Unable to get the Country Definitions, unexpected error: "+err.Error(),
		)
		return
	}
	constCountries, diags := datasource_const_countries.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_countries"), constCountries); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
