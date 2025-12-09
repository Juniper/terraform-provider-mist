package datasource_org_inventory

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Inventory, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := inventorySdkToTerraform(ctx, &diags, d)
		*elements = append(*elements, elem)
	}

	return diags
}

func inventorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Inventory) OrgInventoryValue {
	var adopted basetypes.BoolValue
	var chassisMac basetypes.StringValue
	var chassisSerial basetypes.StringValue
	var claimCode basetypes.StringValue
	var connected basetypes.BoolValue
	var deviceprofileId basetypes.StringValue
	var hostname basetypes.StringValue
	var hwRev basetypes.StringValue
	var id basetypes.StringValue
	var jsi basetypes.BoolValue
	var mac basetypes.StringValue
	var model basetypes.StringValue
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var serial basetypes.StringValue
	var siteId basetypes.StringValue
	var sku basetypes.StringValue
	var deviceType basetypes.StringValue
	var vcMac basetypes.StringValue

	if d.Adopted != nil {
		adopted = types.BoolValue(*d.Adopted)
	}
	if d.ChassisMac != nil {
		chassisMac = types.StringValue(*d.ChassisMac)
	}
	if d.ChassisSerial != nil {
		chassisSerial = types.StringValue(*d.ChassisSerial)
	}
	if d.Connected != nil {
		connected = types.BoolValue(*d.Connected)
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofileId = types.StringValue(*d.DeviceprofileId.Value())
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.HwRev != nil {
		hwRev = types.StringValue(*d.HwRev)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Jsi != nil {
		jsi = types.BoolValue(*d.Jsi)
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.Magic != nil {
		claimCode = types.StringValue(*d.Magic)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.Sku != nil {
		sku = types.StringValue(*d.Sku)
	}
	if d.Type != nil {
		deviceType = types.StringValue(string(*d.Type))
	}
	if d.VcMac != nil {
		vcMac = types.StringValue(*d.VcMac)
	}

	dataMapValue := map[string]attr.Value{
		"adopted":          adopted,
		"chassis_mac":      chassisMac,
		"chassis_serial":   chassisSerial,
		"claim_code":       claimCode,
		"connected":        connected,
		"deviceprofile_id": deviceprofileId,
		"hostname":         hostname,
		"hw_rev":           hwRev,
		"id":               id,
		"jsi":              jsi,
		"mac":              mac,
		"model":            model,
		"name":             name,
		"org_id":           orgId,
		"serial":           serial,
		"site_id":          siteId,
		"sku":              sku,
		"type":             deviceType,
		"vc_mac":           vcMac,
	}
	data, e := NewOrgInventoryValue(OrgInventoryValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
