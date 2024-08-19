package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_device_switch_stats"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*deviceSwitchStatsDataSource)(nil)

func NewDeviceSwitchStatsDataSource() datasource.DataSource {
	return &deviceSwitchStatsDataSource{}
}

type deviceSwitchStatsDataSource struct {
	client mistapi.ClientInterface
}

func (d *deviceSwitchStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Switch Stats Datasource client")
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
func (d *deviceSwitchStatsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_switch_stats"
}

func (d *deviceSwitchStatsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWired + "This data source provides the list of Switches with their statistics.",
		Attributes:          datasource_device_switch_stats.DeviceSwitchStatsDataSourceSchema(ctx).Attributes,
	}
}

func (d *deviceSwitchStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_device_switch_stats.DeviceSwitchStatsModel
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

	var duration *string
	var end *int
	var evpnUnused *string
	var evpnTopoId *string
	var fields *string = models.ToPointer(string("*"))
	var limit *int = models.ToPointer(1000)
	var mac *string
	var page *int
	var siteId *string
	var status *models.DeviceStatusEnum
	var start *int
	var mType *models.DeviceTypeWithAllEnum = models.ToPointer(models.DeviceTypeWithAllEnum(models.DeviceTypeWithAllEnum_ENUMSWITCH))

	if !ds.Duration.IsNull() && !ds.Duration.IsUnknown() {
		duration = ds.Duration.ValueStringPointer()
	}
	if !ds.End.IsNull() && !ds.End.IsUnknown() {
		end_int := int(ds.End.ValueInt64())
		end = &end_int
	}
	if !ds.EvpnUnused.IsNull() && !ds.EvpnUnused.IsUnknown() {
		evpnUnused = ds.EvpnUnused.ValueStringPointer()
	}
	if !ds.EvpntopoId.IsNull() && !ds.EvpntopoId.IsUnknown() {
		evpnTopoId = ds.EvpntopoId.ValueStringPointer()
	}
	if !ds.Mac.IsNull() && !ds.Mac.IsUnknown() {
		mac = ds.Mac.ValueStringPointer()
	}
	if !ds.SiteId.IsNull() && !ds.SiteId.IsUnknown() {
		siteId = ds.SiteId.ValueStringPointer()
	}
	if !ds.Status.IsNull() && !ds.Status.IsUnknown() {
		status = (*models.DeviceStatusEnum)(ds.Status.ValueStringPointer())
	}
	if !ds.Start.IsNull() && !ds.Start.IsUnknown() {
		start_int := int(ds.Start.ValueInt64())
		start = &start_int
	}

	data, err := d.client.OrgsStatsDevices().ListOrgDevicesStats(ctx, orgId, mType, status, siteId, mac, evpnTopoId, evpnUnused, fields, page, limit, start, end, duration)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Switch Stats",
			"Could not get Switch Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceSwitchStat, diags := datasource_device_switch_stats.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("device_switch_stats"), deviceSwitchStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
