package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_site_wlans"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*siteWlansDataSource)(nil)

func NewSiteWlansDataSource() datasource.DataSource {
	return &siteWlansDataSource{}
}

type siteWlansDataSource struct {
	client mistapi.ClientInterface
}

func (d *siteWlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site Wlans Datasource client")
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
func (d *siteWlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wlans"
}

func (d *siteWlansDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This data source provides the list of Site Wlans.\n\n" +
			"The WLAN object contains all the required configuration to broadcast an SSID (Authentication, VLAN, ...)",
		Attributes: datasource_site_wlans.SiteWlansDataSourceSchema(ctx).Attributes,
	}
}

func (d *siteWlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_site_wlans.SiteWlansModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(ds.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_wlans\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var limit = 1000
	var page = 0
	var total = 9999
	var elements []attr.Value
	var diags diag.Diagnostics

	for limit*page < total {
		page += 1
		tflog.Debug(ctx, "Pagination Info", map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		})
		data, err := d.client.SitesWlans().ListSiteWlans(ctx, siteId, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Site Wlans list",
				"Unable to get the the list of Site Wlans, unexpected error: "+err.Error(),
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

		diags = datasource_site_wlans.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_site_wlans.SiteWlansValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("site_wlans"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
