package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_sso_roles"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgSsoRolesDataSource)(nil)

func NewOrgSsoRolesDataSource() datasource.DataSource {
	return &orgSsoRolesDataSource{}
}

type orgSsoRolesDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgSsoRolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org SSO Roles Datasource client")
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
func (d *orgSsoRolesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sso_roles"
}

func (d *orgSsoRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource provides the list of SSO Roles.\n\n" +
			"SSO roles refer to the different functions assigned to users within a Single Sign-On (SSO) system.  \n" +
			"These roles determine the tasks and actions that users can perform within the SSO system. " +
			"There are typically predefined roles and custom roles in an SSO system.  \n" +
			"Roles in SSO provide a well-defined separation of responsibility and visibility, allowing for granular-level access control on SSO objects.",
		Attributes: datasource_org_sso_roles.OrgSsoRolesDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgSsoRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_sso_roles.OrgSsoRolesModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_sso_roles\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
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
		data, err := d.client.OrgsSSORoles().ListOrgSsoRoles(ctx, orgId, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org SSO Roles list",
				"Unable to get the the list of SSO Roles, unexpected error: "+err.Error(),
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

		diags = datasource_org_sso_roles.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_sso_roles.OrgSsoRolesValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_sso_roles"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
