package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_device_gateway_stats"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*deviceGatewayStatsDataSource)(nil)

func NewDeviceGatewayStatsDataSource() datasource.DataSource {
	return &deviceGatewayStatsDataSource{}
}

type deviceGatewayStatsDataSource struct {
	client mistapi.ClientInterface
}

func (d *deviceGatewayStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Gateway Stats Datasource client")
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
func (d *deviceGatewayStatsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway_stats"
}

func (d *deviceGatewayStatsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This data source provides the list of Gateways with their statistics.",
		Attributes:          datasource_device_gateway_stats.DeviceGatewayStatsDataSourceSchema(ctx).Attributes,
	}
}

func (d *deviceGatewayStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_device_gateway_stats.DeviceGatewayStatsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"device_gateway_stats\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var duration string
	var end int
	var fields = "*"
	var mac string
	var siteId string
	var status models.DeviceStatusEnum
	var start int
	var mType = models.DeviceTypeWithAllEnum_GATEWAY

	if !ds.Duration.IsNull() && !ds.Duration.IsUnknown() {
		duration = ds.Duration.ValueString()
	}
	if !ds.End.IsNull() && !ds.End.IsUnknown() {
		end = int(ds.End.ValueInt64())
	}
	if !ds.Mac.IsNull() && !ds.Mac.IsUnknown() {
		mac = ds.Mac.ValueString()
	}
	if !ds.SiteId.IsNull() && !ds.SiteId.IsUnknown() {
		siteId = ds.SiteId.ValueString()
	}
	if !ds.Status.IsNull() && !ds.Status.IsUnknown() {
		status = (models.DeviceStatusEnum)(ds.Status.ValueString())
	}
	if !ds.Start.IsNull() && !ds.Start.IsUnknown() {
		start = int(ds.Start.ValueInt64())
	}

	var limit = 1000
	var page = 0
	var total = 9999
	var elements []attr.Value
	var diags diag.Diagnostics

	for limit*page < total {
		tflog.Debug(ctx, "Pagination Info", map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		})
		page += 1
		data, err := d.client.OrgsStatsDevices().ListOrgDevicesStats(
			ctx,
			orgId,
			&mType,
			&status,
			&siteId,
			&mac,
			nil,
			nil,
			&fields,
			&start,
			&end,
			&duration,
			&limit,
			&page,
		)

		if data.Response.StatusCode != 200 && err != nil {
			resp.Diagnostics.AddError(
				"Error getting Gateway Stats",
				"Unable to get the Gateway Stats, unexpected error: "+err.Error(),
			)
			return
		}

		limitString := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limitString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Limit value into int, unexpected error: "+err.Error(),
			)
			return
		}

		totalString := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(totalString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Total value into int, unexpected error: "+err.Error(),
			)
			return
		}

		body, _ := io.ReadAll(data.Response.Body)
		var mistStats []models.StatsGateway
		json.Unmarshal(body, &mistStats)

		diags = datasource_device_gateway_stats.SdkToTerraform(ctx, &mistStats, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_device_gateway_stats.DeviceGatewayStatsValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("device_gateway_stats"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
