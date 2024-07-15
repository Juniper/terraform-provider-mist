package datasource_org_inventory

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Inventory) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := inventorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(OrgInventoryValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func inventorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Inventory) OrgInventoryValue {
	var adopted basetypes.BoolValue
	var claim_code basetypes.StringValue
	var connected basetypes.BoolValue
	var created_time basetypes.Int64Value
	var deviceprofile_id basetypes.StringValue
	var hostname basetypes.StringValue
	var hw_rev basetypes.StringValue
	var id basetypes.StringValue
	var jsi basetypes.BoolValue
	var mac basetypes.StringValue
	var model basetypes.StringValue
	var modified_time basetypes.Int64Value
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var serial basetypes.StringValue
	var site_id basetypes.StringValue
	var sku basetypes.StringValue
	var vc_mac basetypes.StringValue

	if d.Adopted != nil {
		adopted = types.BoolValue(*d.Adopted)
	}
	if d.Connected != nil {
		connected = types.BoolValue(*d.Connected)
	}
	if d.CreatedTime != nil {
		created_time = types.Int64Value(int64(*d.CreatedTime))
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofile_id = types.StringValue(*d.DeviceprofileId.Value())
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.HwRev != nil {
		hw_rev = types.StringValue(*d.HwRev)
	}
	if d.Id != nil {
		id = types.StringValue(*d.Id)
	}
	if d.Jsi != nil {
		jsi = types.BoolValue(*d.Jsi)
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.Magic != nil {
		claim_code = types.StringValue(*d.Magic)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.ModifiedTime != nil {
		modified_time = types.Int64Value(int64(*d.ModifiedTime))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.Sku != nil {
		sku = types.StringValue(*d.Sku)
	}
	if d.VcMac != nil {
		vc_mac = types.StringValue(*d.VcMac)
	}

	data_map_attr_type := OrgInventoryValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"adopted":          adopted,
		"claim_code":       claim_code,
		"connected":        connected,
		"created_time":     created_time,
		"deviceprofile_id": deviceprofile_id,
		"hostname":         hostname,
		"hw_rev":           hw_rev,
		"id":               id,
		"jsi":              jsi,
		"mac":              mac,
		"model":            model,
		"modified_time":    modified_time,
		"name":             name,
		"org_id":           org_id,
		"serial":           serial,
		"site_id":          site_id,
		"sku":              sku,
		"vc_mac":           vc_mac,
	}
	data, e := NewOrgInventoryValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
