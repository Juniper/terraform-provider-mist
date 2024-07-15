package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_device_ap_stats"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*deviceApStatsDataSource)(nil)

func NewDeviceApStatsDataSource() datasource.DataSource {
	return &deviceApStatsDataSource{}
}

type deviceApStatsDataSource struct {
	client mistapi.ClientInterface
}

func (d *deviceApStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *deviceApStatsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ap_stats"
}

func (d *deviceApStatsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasource_device_ap_stats.DeviceApStatsDataSourceSchema(ctx)
}

func (d *deviceApStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_device_ap_stats.DeviceApStatsModel
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
	var fields *string = models.ToPointer(string("*"))
	var limit *int = models.ToPointer(1000)
	var mac *string
	var page *int
	var siteId *string
	var status *models.DeviceStatusEnum
	var start *int
	var mType *models.DeviceTypeWithAllEnum = models.ToPointer(models.DeviceTypeWithAllEnum(models.DeviceTypeWithAllEnum_AP))

	if !ds.Duration.IsNull() && !ds.Duration.IsUnknown() {
		duration = ds.Duration.ValueStringPointer()
	}
	if !ds.End.IsNull() && !ds.End.IsUnknown() {
		end_int := int(ds.End.ValueInt64())
		end = &end_int
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

	data, err := d.client.OrgsDevicesStats().ListOrgDevicesStats(ctx, orgId, mType, status, siteId, mac, nil, nil, fields, page, limit, start, end, duration)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}

	deviceApStat, diags := datasource_device_ap_stats.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("device_ap_stats"), deviceApStat); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
