package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_device_switch_stats"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
			"Invalid \"org_id\" value for \"device_switch_stats\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var duration string
	var end int
	var fields string = string("*")
	var mac string
	var siteId string
	var status models.DeviceStatusEnum
	var start int
	var mType models.DeviceTypeWithAllEnum = models.DeviceTypeWithAllEnum(models.DeviceTypeWithAllEnum_ENUMSWITCH)

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

	var limit int = 1000
	var page int = 0
	var total int = 9999
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
				"Error getting Switch Stats",
				"Unable to get the Switch Stats, unexpected error: "+err.Error(),
			)
			return
		}

		limit_string := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limit_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Limit value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		total_string := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(total_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Total value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		body, _ := io.ReadAll(data.Response.Body)
		mist_stats := []models.StatsSwitch{}
		json.Unmarshal(body, &mist_stats)

		diags = datasource_device_switch_stats.SdkToTerraform(ctx, &mist_stats, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_device_switch_stats.DeviceSwitchStatsValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("device_switch_stats"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
