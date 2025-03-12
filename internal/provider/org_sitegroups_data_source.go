package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_sitegroups"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgSitegroupsDataSource)(nil)

func NewOrgSitegroupsDataSource() datasource.DataSource {
	return &orgSitegroupsDataSource{}
}

type orgSitegroupsDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgSitegroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Site Groups Datasource client")
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
func (d *orgSitegroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sitegroups"
}

func (d *orgSitegroupsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This data source provides the list of Org Site Groups (sitegroups).\n\n" +
			"A site group is a feature that allows users to group multiple sites together based on regions, functions, or other parameters for efficient management of devices.  \n" +
			"Sites can exist in multiple groups simultaneously, and site groups can be used to ensure consistent settings, manage administrator access, and apply specific templates to groups of sites.",
		Attributes: datasource_org_sitegroups.OrgSitegroupsDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgSitegroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_sitegroups.OrgSitegroupsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sitegroups\" data_source",
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
		data, err := d.client.OrgsSitegroups().ListOrgSiteGroups(ctx, orgId, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org Site Groups list",
				"Unable to get the the list of Org Site Groups, unexpected error: "+err.Error(),
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

		diags = datasource_org_sitegroups.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_sitegroups.OrgSitegroupsValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_sitegroups"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
