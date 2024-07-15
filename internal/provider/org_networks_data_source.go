package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"terraform-provider-mist/internal/datasource_org_networks"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgNetworksDataSource)(nil)

func NewOrgNetworksDataSource() datasource.DataSource {
	return &orgNetworksDataSource{}
}

type orgNetworksDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgNetworksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *orgNetworksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_networks"
}

func (d *orgNetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_org_networks.OrgNetworksDataSourceSchema(ctx)
}

func (d *orgNetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_networks.OrgNetworksModel
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

	data, err := d.client.OrgsNetworks().ListOrgNetworks(ctx, orgId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_networks.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_networks"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
