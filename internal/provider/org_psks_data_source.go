package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_psks"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgPsksDataSource)(nil)

func NewOrgPsksDataSource() datasource.DataSource {
	return &orgPsksDataSource{}
}

type orgPsksDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgPsksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Psks Datasource client")
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
func (d *orgPsksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_psks"
}

func (d *orgPsksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance Psks." +
			"The Psks are used in the `service_policies` from the Gateway configuration and Gateway templates ",
		Attributes: datasource_org_psks.OrgPsksDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgPsksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_psks.OrgPsksModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psks\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var name *string = ds.Name.ValueStringPointer()
	var limit *int = models.ToPointer(1000)
	var page *int
	var ssid *string = ds.Ssid.ValueStringPointer()
	var role *string = ds.Role.ValueStringPointer()

	data, err := d.client.OrgsPsks().ListOrgPsks(ctx, orgId, name, ssid, role, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_psks.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_psks"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
