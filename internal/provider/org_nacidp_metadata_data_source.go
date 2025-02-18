package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_nacidp_metadata"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgNacidpMetadataDataSource)(nil)

func NewOrgNacidpMetadataDataSource() datasource.DataSource {
	return &orgNacidpMetadataDataSource{}
}

type orgNacidpMetadataDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgNacidpMetadataDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org NAC IDP Metadata Datasource client")
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
func (d *orgNacidpMetadataDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nacidp_metadata"
}

func (d *orgNacidpMetadataDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This data source provides the NAC IDP Metadata information.\n\n" +
			"The provided information (`entity_id`, `acs_url`, `logout_url` and `metadata`) are the information" +
			"required to configure the IDP",
		Attributes: datasource_org_nacidp_metadata.OrgNacidpMetadataDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgNacidpMetadataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_nacidp_metadata.OrgNacidpMetadataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_nacidp_metadata\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	nacidpId, err := uuid.Parse(ds.NacidpId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacidp_id\" value for \"org_nacidp_metadata\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	data, err := d.client.OrgsSSO().GetOrgSamlMetadata(ctx, orgId, nacidpId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NAC IDP Metadata",
			"Unable to get the the NAC IDP Metadata, unexpected error: "+err.Error(),
		)
		return
	}

	metadata := datasource_org_nacidp_metadata.SdkToTerraform(&data.Data)
	metadata.NacidpId = ds.NacidpId
	metadata.OrgId = ds.OrgId

	diags := resp.State.Set(ctx, metadata)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
