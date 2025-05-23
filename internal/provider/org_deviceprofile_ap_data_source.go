package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_deviceprofiles_ap"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgDeviceprofilesApDataSource)(nil)

func NewOrgDeviceprofilesApDataSource() datasource.DataSource {
	return &orgDeviceprofilesApDataSource{}
}

type orgDeviceprofilesApDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgDeviceprofilesApDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Device Profiles AP Datasource client")
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
func (d *orgDeviceprofilesApDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofiles_ap"
}

func (d *orgDeviceprofilesApDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This data source provides the list of AP Device Profiles.\n\n" +
			"AP Device profiles are used to specify a configuration that can be applied to a select set of aps from any site in the organization. " +
			"They are providing efficient application of configurations based on ap groups, wlan groups, RF settings, and sites. " +
			"Device profiles enable various use cases such as activating ethernet passthrough, applying different rf settings, applying mesh configuration, " +
			"activating specific features like esl or vble, and more.",
		Attributes: datasource_org_deviceprofiles_ap.OrgDeviceprofilesApDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgDeviceprofilesApDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_deviceprofiles_ap.OrgDeviceprofilesApModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofiles_ap\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var mType = models.DeviceTypeDefaultApEnum_AP

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
		data, err := d.client.OrgsDeviceProfiles().ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
		if data.Response.StatusCode != 200 && err != nil {
			resp.Diagnostics.AddError(
				"Error getting AP Device Profiles",
				"Unable to get the the list of AP Device Profiles, unexpected error: "+err.Error(),
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
		var mistDeviceprofiles []models.DeviceprofileAp
		json.Unmarshal(body, &mistDeviceprofiles)

		diags = datasource_org_deviceprofiles_ap.SdkToTerraform(ctx, &mistDeviceprofiles, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_deviceprofiles_ap.OrgDeviceprofilesApValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("deviceprofiles"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
