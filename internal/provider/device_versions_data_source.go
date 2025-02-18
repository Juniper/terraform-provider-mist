package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_device_versions"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*deviceVersionsDataSource)(nil)

func NewDeviceVersionsDataSource() datasource.DataSource {
	return &deviceVersionsDataSource{}
}

type deviceVersionsDataSource struct {
	client mistapi.ClientInterface
}

func (d *deviceVersionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Device Versions Datasource client")
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
func (d *deviceVersionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_versions"
}

func (d *deviceVersionsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This data source provides the list of available Firmware Versions.",
		Attributes:          datasource_device_versions.DeviceVersionsDataSourceSchema(ctx).Attributes,
	}
}

func (d *deviceVersionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_device_versions.DeviceVersionsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"device_versions\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var mType models.DeviceTypeEnum
	var model string

	if !ds.Type.IsNull() && !ds.Type.IsUnknown() {
		mType = models.DeviceTypeEnum(ds.Type.ValueString())
	}
	if !ds.Model.IsNull() && !ds.Model.IsUnknown() {
		model = ds.Model.ValueString()
	}

	var diags diag.Diagnostics

	data, err := d.client.UtilitiesUpgrade().ListOrgAvailableDeviceVersions(ctx, orgId, &mType, &model)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting Device Versions",
			"Unable to get the Device Versions, unexpected error: "+err.Error(),
		)
		return
	}

	dataSet, diags := datasource_device_versions.SdkToTerraform(ctx, data.Data)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("device_versions"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
