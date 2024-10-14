package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_sso_metadata"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgSsoMetadataDataSource)(nil)

func NewOrgSsoMetadataDataSource() datasource.DataSource {
	return &orgSsoMetadataDataSource{}
}

type orgSsoMetadataDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgSsoMetadataDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org SSO Metadata Datasource client")
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
func (d *orgSsoMetadataDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sso_metadata"
}

func (d *orgSsoMetadataDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This data source provides the SSO Metadata information.\n" +
			"The provided information (`entity_id`, `acs_url`, `logout_url` and `metadata`) are the information" +
			"required to configure the IDP",
		Attributes: datasource_org_sso_metadata.OrgSsoMetadataDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgSsoMetadataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_sso_metadata.OrgSsoMetadataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sso_metadata\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	ssoId, err := uuid.Parse(ds.SsoId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"sso_id\" value for \"org_sso_metadata\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	data, err := d.client.OrgsSSO().GetOrgSsoSamlMetadata(ctx, orgId, ssoId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting SSO Metadata",
			"Unable to get the the SSO Metadata, unexpected error: "+err.Error(),
		)
		return
	}

	metadata := datasource_org_sso_metadata.SdkToTerraform(ctx, &data.Data)
	metadata.SsoId = ds.SsoId
	metadata.OrgId = ds.OrgId

	diags := resp.State.Set(ctx, metadata)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
