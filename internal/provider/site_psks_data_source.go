package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_site_psks"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*sitePsksDataSource)(nil)

func NewSitePsksDataSource() datasource.DataSource {
	return &sitePsksDataSource{}
}

type sitePsksDataSource struct {
	client mistapi.ClientInterface
}

func (d *sitePsksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site Psks Datasource client")
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
func (d *sitePsksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_psks"
}

func (d *sitePsksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance Psks." +
			"The Psks are used in the `service_policies` from the Gateway configuration and Gateway templates ",
		Attributes: datasource_site_psks.SitePsksDataSourceSchema(ctx).Attributes,
	}
}

func (d *sitePsksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_site_psks.SitePsksModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(ds.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_psks\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var name string
	var ssid string
	var role string

	if ds.Name.ValueStringPointer() != nil {
		name = ds.Name.ValueString()
	}
	if ds.Ssid.ValueStringPointer() != nil {
		ssid = ds.Ssid.ValueString()
	}
	if ds.Role.ValueStringPointer() != nil {
		role = ds.Role.ValueString()
	}

	var limit int = 1000
	var page int = 0
	var total int = 9999
	var elements []attr.Value
	var diags diag.Diagnostics

	for limit*page < total {
		page += 1
		tflog.Debug(ctx, "Pagination Info", map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		})
		data, err := d.client.SitesPsks().ListSitePsks(ctx, siteId, &name, &ssid, &role, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Site PSKs list",
				"Could not get the list of Site PSKs, unexpected error: "+err.Error(),
			)
			return
		}

		limit_string := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limit_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Could not convert X-Page-Limit value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		total_string := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(total_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Could not convert X-Page-Total value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		diags = datasource_site_psks.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_site_psks.SitePsksValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("site_psks"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
