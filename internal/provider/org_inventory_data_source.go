package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_inventory"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	tflog.Info(ctx, "Configuring Mist Org Inventory Datasource client")
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
func (d *orgInventoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_inventory"
}

func (d *orgInventoryDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"Invalid \"org_id\" value for \"org_inventory\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var mac string
	var model string
	var serial string
	var siteId *uuid.UUID
	var unassigned bool
	var modifiedAfter int
	var vc bool
	var vcMac string
	var deviceType models.DeviceTypeEnum

	if !ds.Mac.IsNull() && !ds.Mac.IsUnknown() {
		mac = ds.Mac.ValueString()
	}
	if !ds.Model.IsNull() && !ds.Model.IsUnknown() {
		model = ds.Model.ValueString()
	}
	if !ds.SiteId.IsNull() && !ds.SiteId.IsUnknown() {
		tmp, err := uuid.Parse(ds.SiteId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid \"siteId\" value for \"org_inventory\" data_source",
				"Could parse the UUID: "+err.Error(),
			)
			return
		}
		siteId = &tmp
	}
	if !ds.Serial.IsNull() && !ds.Serial.IsUnknown() {
		serial = ds.Serial.ValueString()
	}
	if !ds.Unassigned.IsNull() && !ds.Unassigned.IsUnknown() {
		unassigned = ds.Unassigned.ValueBool()
	}
	if !ds.VcMac.IsNull() && !ds.VcMac.IsUnknown() {
		vcMac = ds.VcMac.ValueString()
	}
	if !ds.Vc.IsNull() && !ds.Vc.IsUnknown() {
		vc = ds.Vc.ValueBool()
	}
	if !ds.Type.IsNull() && !ds.Type.IsUnknown() {
		deviceType = models.DeviceTypeEnum(ds.Type.ValueString())
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
		// Read API call logic
		data, err := d.client.OrgsInventory().GetOrgInventory(ctx, orgId, &serial, &model, &deviceType, &mac, siteId, &vcMac, &vc, &unassigned, &modifiedAfter, &limit, &page)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org Inventory",
				"Unable to get the Org Inventory, unexpected error: "+err.Error(),
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

		diags = datasource_org_inventory.SdkToTerraform(ctx, &data.Data, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	dataSet, diags := types.SetValue(datasource_org_inventory.OrgInventoryValue{}.Type(ctx), elements)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := resp.State.SetAttribute(ctx, path.Root("org_inventory"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
