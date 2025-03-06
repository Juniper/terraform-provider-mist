package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_fingerprints"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constFingerprintsDataSource)(nil)

func NewConstFingerprintsDataSource() datasource.DataSource {
	return &constFingerprintsDataSource{}
}

type constFingerprintsDataSource struct {
	client mistapi.ClientInterface
}

func (d *constFingerprintsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *constFingerprintsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_fingerprints"
}

func (d *constFingerprintsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the of list of supported Fingerprints.\n\n" +
			"The Fingerprint information can be used within `matching` and `not_matching` attributes of the NAC Rule resource (`mist_org_nacrule`)\n\n" +
			"There is four different type of fingerprints available:" +
			"* Family\n" +
			"* Model\n" +
			"* Mfg\n" +
			"* OS Type\n",
		Attributes: datasource_const_fingerprints.ConstFingerprintsDataSourceSchema(ctx).Attributes,
	}
}

func (d *constFingerprintsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_fingerprints.ConstFingerprintsModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListFingerprintTypes(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Fingerprints",
			"Unable to get the Fingerprints, unexpected error: "+err.Error(),
		)
		return
	}
	constFingerprints, diags := datasource_const_fingerprints.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, constFingerprints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
