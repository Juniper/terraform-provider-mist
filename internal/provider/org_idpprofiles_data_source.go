package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_idpprofiles"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgIdpprofilesDataSource)(nil)

func NewOrgIdpprofilesDataSource() datasource.DataSource {
	return &orgIdpprofilesDataSource{}
}

type orgIdpprofilesDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgIdpprofilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org IDP Profiles Datasource client")
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
func (d *orgIdpprofilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_idpprofiles"
}

func (d *orgIdpprofilesDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This data source provides the list of WAN Assurance IDP Profiles.\n\n" +
			"An IDP Profile is a configuration setting that defines the behavior and actions of an intrusion detection and prevention (IDP) system. " +
			"It specifies how the idp system should detect and respond to potential security threats or attacks on a network.\n\n" +
			"The profile includes rules and policies that determine which types of traffic or attacks should be monitored, " +
			"what actions should be taken when a threat is detected, and any exceptions or exclusions for specific destinations or attack types.",
		Attributes: datasource_org_idpprofiles.OrgIdpprofilesDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgIdpprofilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_idpprofiles.OrgIdpprofilesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_idpprofiles\" data_source",
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
		data, err := d.client.OrgsIDPProfiles().ListOrgIdpProfiles(ctx, orgId, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org IDP Profiles list",
				"Unable to get the the list of Org IDP Profiles, unexpected error: "+err.Error(),
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

		diags = datasource_org_idpprofiles.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_idpprofiles.OrgIdpprofilesValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_idpprofiles"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
