package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_app_categories"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constAppCategoriesDataSource)(nil)

func NewConstAppCategoriesDataSource() datasource.DataSource {
	return &constAppCategoriesDataSource{}
}

type constAppCategoriesDataSource struct {
	client mistapi.ClientInterface
}

func (d *constAppCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constants App Catgories Datasource client")
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
func (d *constAppCategoriesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_app_categories"
}

func (d *constAppCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the of ConstAppCategories.\n" +
			"This information can be used as `app_categories` in the `mist_org_service` resource",
		Attributes: datasource_const_app_categories.ConstAppCategoriesDataSourceSchema(ctx).Attributes,
	}
}

func (d *constAppCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_app_categories.ConstAppCategoriesModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListAppCategoryDefinitions(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Unable to get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}
	constAppCategories, diags := datasource_const_app_categories.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_app_categories"), constAppCategories); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
