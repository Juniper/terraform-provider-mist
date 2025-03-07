package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_usermacs"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgNacEndpointsDataSource)(nil)

func NewOrgNacEndpointsDataSource() datasource.DataSource {
	return &orgNacEndpointsDataSource{}
}

type orgNacEndpointsDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgNacEndpointsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org NacEndpoints Datasource client")
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
func (d *orgNacEndpointsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nac_endpoints"
}

func (d *orgNacEndpointsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This data source provides the list of NAC Endpoints (User MACs).\n\n" +
			"NAC Endpoints (User MACs) provide a database of endpoints identified by their MAC addresses. " +
			"They can be used assign each endpoint with various attributes, such as name, VLAN, role and client label.  " +
			"Once an endpoint is labeled, the label name can be used to create `mist_org_nactag` resource as match criteria.",
		Attributes: datasource_org_usermacs.OrgUsermacsDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgNacEndpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_usermacs.OrgUsermacsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_NacEndpoints\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var mac = ds.Mac.ValueStringPointer()
	var labels []string
	if !ds.Labels.IsNull() && !ds.Labels.IsUnknown() {
		labels = mistutils.ListOfStringTerraformToSdk(ds.Labels)
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
		data, err := d.client.OrgsUserMACs().SearchOrgUserMacs(ctx, orgId, mac, labels, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org NacEndpoints list",
				"Unable to get the the list of Org NacEndpoints, unexpected error: "+err.Error(),
			)
			return
		}

		limitString := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limitString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Limit value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		totalString := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(totalString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Total value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		diags = datasource_org_usermacs.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_usermacs.OrgUsermacsValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_usermacs"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
