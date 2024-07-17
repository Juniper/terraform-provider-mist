package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_idpprofiles"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgIdpprofilesDataSource)(nil)

func NewOrgIdpprofilesDataSource() datasource.DataSource {
	return &orgIdpprofilesDataSource{}
}

type orgIdpprofilesDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgIdpprofilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist AP Stats")
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
func (d *orgIdpprofilesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_idpprofiles"
}

func (d *orgIdpprofilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance IDP Profiles." +
			"An IDP Profile is a configuration setting that defines the behavior and actions of an intrusion detection and prevention (IDP) system." +
			"It specifies how the idp system should detect and respond to potential security threats or attacks on a network." +
			"The profile includes rules and policies that determine which types of traffic or attacks should be monitored," +
			"what actions should be taken when a threat is detected, and any exceptions or exclusions for specific destinations or attack types.",
		Attributes: datasource_org_idpprofiles.OrgIdpprofilesDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgIdpprofilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_idpprofiles.OrgIdpprofilesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id from ds",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var page *int

	data, err := d.client.OrgsIDPProfiles().ListOrgIdpProfiles(ctx, orgId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_idpprofiles.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_idpprofiles"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
