package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_inventory"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgInventoryDataSource)(nil)

func NewOrgInventoryDataSource() datasource.DataSource {
	return &orgInventoryDataSource{}
}

type orgInventoryDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgInventoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist AP Stats")
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
func (d *orgInventoryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_inventory"
}

func (d *orgInventoryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This data source provides the list of Devices in the Org inventory.",
		Attributes:          datasource_org_inventory.OrgInventoryDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgInventoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_inventory.OrgInventoryModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_id from ds",
			"Could not get org_id, unexpected error: "+err.Error(),
		)
		return
	}

	var limit *int = models.ToPointer(1000)
	var mac *string
	var model *string
	var page *int
	var serial *string
	var siteId *string
	var unassigned *bool
	var vcMac *string

	if !ds.Mac.IsNull() && !ds.Mac.IsUnknown() {
		mac = ds.Mac.ValueStringPointer()
	}
	if !ds.Model.IsNull() && !ds.Model.IsUnknown() {
		model = ds.Model.ValueStringPointer()
	}
	if !ds.SiteId.IsNull() && !ds.SiteId.IsUnknown() {
		siteId = ds.SiteId.ValueStringPointer()
	}
	if !ds.Serial.IsNull() && !ds.Serial.IsUnknown() {
		serial = ds.Serial.ValueStringPointer()
	}
	if !ds.Unassigned.IsNull() && !ds.Unassigned.IsUnknown() {
		unassigned = ds.Unassigned.ValueBoolPointer()
	}
	if !ds.VcMac.IsNull() && !ds.VcMac.IsUnknown() {
		vcMac = ds.VcMac.ValueStringPointer()
	}

	// Read API call logic
	data, err := d.client.OrgsInventory().GetOrgInventory(ctx, orgId, serial, model, nil, mac, siteId, vcMac, nil, unassigned, limit, page)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting AP Stats",
			"Could not get AP Stats, unexpected error: "+err.Error(),
		)
		return
	}
	orgInventory, diags := datasource_org_inventory.SdkToTerraform(ctx, data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_inventory"), orgInventory); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
