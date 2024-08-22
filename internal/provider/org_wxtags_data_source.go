package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_wxtags"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgWxtagsDataSource)(nil)

func NewOrgWxtagsDataSource() datasource.DataSource {
	return &orgWxtagsDataSource{}
}

type orgWxtagsDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgWxtagsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org WxTags Datasource client")
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
func (d *orgWxtagsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wxtags"
}

func (d *orgWxtagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource provides the list of Org WxLan tags (labels)." +
			"A WxTag is a label or tag used in the mist system to classify and categorize applications, " +
			"users, and resources for the purpose of creating policies and making network management decisions." +
			"They can be used " +
			"  * within the WxRules to create filtering rules, or assign specific VLAN" +
			"  * in the WLANs configuration to assign a WLAN to specific APs" +
			"  * to identify unknown application used by Wi-Fi clients",
		Attributes: datasource_org_wxtags.OrgWxtagsDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgWxtagsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_wxtags.OrgWxtagsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_wxtags\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var page *int

	data, err := d.client.OrgsWxTags().ListOrgWxTags(ctx, orgId, page, limit)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_org_wxtags.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_wxtags"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
