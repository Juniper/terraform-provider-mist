package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_avprofiles"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgAvprofilesDataSource)(nil)

func NewOrgAvprofilesDataSource() datasource.DataSource {
	return &orgAvprofilesDataSource{}
}

type orgAvprofilesDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgAvprofilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org AV Profiles Datasource client")
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
func (d *orgAvprofilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_avprofiles"
}

func (d *orgAvprofilesDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance Antivirus Profiles.\n\n" +
			"An Antivirus Profile is used to configure the Antivirus feature on SRX devices. " +
			"It specifies which content the Antivirus should analyse and which content should be ignored.\n\n" +
			"The Antivirus profiles can be used within the following resources: \n" +
			" * `mist_org_servicepolicy.antivirus` \n" +
			" * `mist_org_gatewaytemplate.service_policies.antivirus` \n" +
			" * `mist_org_deviceprofile_gateway.service_policies.antivirus` \n" +
			" * `mist_devicee_gateway.service_policies.antivirus` \n",
		Attributes: datasource_org_avprofiles.OrgAvprofilesDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgAvprofilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_avprofiles.OrgAvprofilesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_avprofiles\" data_source",
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
		data, err := d.client.OrgsAntivirusProfiles().ListOrgAntivirusProfiles(ctx, orgId, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org AV Profiles list",
				"Unable to get the the list of Org AV Profiles, unexpected error: "+err.Error(),
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

		diags = datasource_org_avprofiles.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_avprofiles.OrgAvprofilesValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_avprofiles"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
