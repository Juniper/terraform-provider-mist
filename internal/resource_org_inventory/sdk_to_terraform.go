package resource_org_inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func checkVcSiteId(device *DevicesValue, vcmac_to_site map[string]types.String) {
	if device.VcMac.ValueString() != "" && device.SiteId.ValueString() == "" {
		var vcMac string = strings.ToLower(device.VcMac.ValueString())
		if site_id, ok := vcmac_to_site[vcMac]; ok {
			device.SiteId = site_id
		}

	}
}

func SdkToTerraform(ctx context.Context, orgId string, data []models.Inventory, plan *OrgInventoryModel) (OrgInventoryModel, diag.Diagnostics) {
	var state OrgInventoryModel
	var diags diag.Diagnostics
	var devices_out []attr.Value
	devices_map_magic := make(map[string]*DevicesValue)
	devices_map_mac := make(map[string]*DevicesValue)
	vcmac_to_site := make(map[string]types.String)

	state.OrgId = types.StringValue(orgId)

	for _, d := range data {
		var claim_code basetypes.StringValue
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var org_id basetypes.StringValue
		var serial basetypes.StringValue
		var site_id basetypes.StringValue
		var device_type basetypes.StringValue
		var vc_mac basetypes.StringValue
		var hostname basetypes.StringValue
		var id basetypes.StringValue

		if d.Magic != nil {
			claim_code = types.StringValue(*d.Magic)
		}
		if d.Mac != nil {
			mac = types.StringValue(*d.Mac)
		}
		if d.Model != nil {
			model = types.StringValue(*d.Model)
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
		if d.Type != nil {
			device_type = types.StringValue(string(*d.Type))
		}
		if d.VcMac != nil {
			vc_mac = types.StringValue(*d.VcMac)
		}
		if d.Hostname != nil {
			hostname = types.StringValue(*d.Hostname)
		}
		if d.Id != nil {
			id = types.StringValue(*d.Id)
		}

		data_map_attr_type := DevicesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"claim_code": claim_code,
			"mac":        mac,
			"model":      model,
			"org_id":     org_id,
			"serial":     serial,
			"site_id":    site_id,
			"type":       device_type,
			"vc_mac":     vc_mac,
			"hostname":   hostname,
			"id":         id,
		}
		newDevice, e := NewDevicesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		var nMagic string = strings.ToUpper(newDevice.Magic.ValueString())
		var nMac string = strings.ToLower(newDevice.Mac.ValueString())
		devices_map_mac[nMac] = &newDevice
		if nMagic != "" {
			// for claimed devices
			devices_map_magic[nMagic] = &newDevice
		}

		if newDevice.VcMac.Equal(newDevice.Mac) {
			vcmac_to_site[nMac] = newDevice.SiteId
		}
	}

	// If it is for an Import (no plan.OrgId), then return all the claimed devices
	// otherwise, just return the devices from the plan to be sure to not unclaim
	// devices not managed by TF
	if plan.OrgId.ValueStringPointer() == nil {
		for _, dev := range devices_map_mac {
			checkVcSiteId(dev, vcmac_to_site)
			devices_out = append(devices_out, dev)
		}
	} else {
		for _, dev_plan_attr := range plan.Devices.Elements() {
			var dpi interface{} = dev_plan_attr
			var device = dpi.(DevicesValue)

			var magic string = strings.ReplaceAll(strings.ToUpper(device.Magic.ValueString()), "-", "")
			var mac string = strings.ToLower(device.Mac.ValueString())

			if dev_from_mist, ok := devices_map_magic[magic]; ok {
				checkVcSiteId(dev_from_mist, vcmac_to_site)
				devices_out = append(devices_out, *dev_from_mist)
			} else if dev_from_mist, ok := devices_map_mac[mac]; ok {
				checkVcSiteId(dev_from_mist, vcmac_to_site)
				devices_out = append(devices_out, *dev_from_mist)
			} else if magic != "" {
				diags.AddError("Device not found", fmt.Sprintf("Unable to find device with Claim Code \"%s\" in the Org Inventory", magic))
			} else {
				diags.AddError("Device not found", fmt.Sprintf("Unable to find device with MAC \"%s\" in the Org Inventory", mac))
			}
		}
	}

	devices_list, e := basetypes.NewListValue(DevicesValue{}.Type(ctx), devices_out)
	diags.Append(e...)
	state.Devices = devices_list

	return state, diags
}
