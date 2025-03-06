package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_const_alarms"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*constAlarmsDataSource)(nil)

func NewConstAlarmsDataSource() datasource.DataSource {
	return &constAlarmsDataSource{}
}

type constAlarmsDataSource struct {
	client mistapi.ClientInterface
}

func (d *constAlarmsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Constant Alarms Datasource client")
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
func (d *constAlarmsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_const_alarms"
}

func (d *constAlarmsDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryConst + "This data source provides the list of  available Alarms.\n\n" +
			"The alarm `key` can be used to configure the `mist_org_alarmtemplate.rules`.",
		Attributes: datasource_const_alarms.ConstAlarmsDataSourceSchema(ctx).Attributes,
	}
}

func (d *constAlarmsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_const_alarms.ConstAlarmsModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	data, err := d.client.ConstantsDefinitions().ListAlarmDefinitions(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Alarm Definitions",
			"Unable to get the Alarm Definitions, unexpected error: "+err.Error(),
		)
		return
	}
	constAlarms, diags := datasource_const_alarms.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("const_alarms"), constAlarms); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
