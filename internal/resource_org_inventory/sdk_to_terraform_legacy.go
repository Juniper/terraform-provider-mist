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

func legacyCheckVcSiteId(device *DevicesValue, vcmac_to_site map[string]types.String) {
	if device.VcMac.ValueString() != "" && device.SiteId.ValueString() == "" {
		var vcMac string = strings.ToUpper(device.VcMac.ValueString())
		if site_id, ok := vcmac_to_site[vcMac]; ok {
			device.SiteId = site_id
		}

	}
}
func legacyProcessMistInventory(
	ctx context.Context,
	diags *diag.Diagnostics,
	data *[]models.Inventory,
	mistDevicesByClaimCode *map[string]*DevicesValue,
	mistDevicesbyMac *map[string]*DevicesValue,
	mistSiteIdByVcMac *map[string]types.String,
) {
	for _, d := range *data {
		var deviceprofile_id basetypes.StringValue = types.StringValue("")
		var hostname basetypes.StringValue = types.StringValue("")
		var id basetypes.StringValue
		var mac basetypes.StringValue
		var claim_code basetypes.StringValue = types.StringValue("")
		var model basetypes.StringValue
		var org_id basetypes.StringValue
		var serial basetypes.StringValue
		var site_id basetypes.StringValue
		var device_type basetypes.StringValue
		var vc_mac basetypes.StringValue = types.StringValue("")
		var unclaim_when_destroyed basetypes.BoolValue = types.BoolValue(false)

		if d.DeviceprofileId.Value() != nil {
			deviceprofile_id = types.StringValue(*d.DeviceprofileId.Value())
		}
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
			id = types.StringValue(d.Id.String())
		}

		data_map_attr_type := DevicesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"deviceprofile_id":       deviceprofile_id,
			"hostname":               hostname,
			"id":                     id,
			"mac":                    mac,
			"claim_code":             claim_code,
			"model":                  model,
			"org_id":                 org_id,
			"serial":                 serial,
			"site_id":                site_id,
			"type":                   device_type,
			"unclaim_when_destroyed": unclaim_when_destroyed,
			"vc_mac":                 vc_mac,
		}
		newDevice, e := NewDevicesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		var nMagic string = strings.ToUpper(newDevice.Magic.ValueString())
		var nMac string = strings.ToUpper(newDevice.Mac.ValueString())
		(*mistDevicesbyMac)[nMac] = &newDevice
		if nMagic != "" {
			// for claimed devices
			(*mistDevicesByClaimCode)[nMagic] = &newDevice
		}
		if newDevice.VcMac.Equal(newDevice.Mac) {
			(*mistSiteIdByVcMac)[nMac] = newDevice.SiteId
		}
	}
}
func legacySdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Inventory,
	ref_inventory *OrgInventoryModel,
) (OrgInventoryModel, diag.Diagnostics) {
	var newState OrgInventoryModel
	var diags diag.Diagnostics
	var devicesOut []attr.Value
	mistDevicesByClaimCode := make(map[string]*DevicesValue)
	mistDevicesbyMac := make(map[string]*DevicesValue)
	mistSiteIdByVcMac := make(map[string]types.String)

	legacyProcessMistInventory(ctx, &diags, data, &mistDevicesByClaimCode, &mistDevicesbyMac, &mistSiteIdByVcMac)

	newState.OrgId = types.StringValue(orgId)

	for _, devref_inventoryAttr := range ref_inventory.Devices.Elements() {
		var dpi interface{} = devref_inventoryAttr
		var device = dpi.(DevicesValue)

		var magic string = strings.ReplaceAll(strings.ToUpper(device.Magic.ValueString()), "-", "")
		var mac string = strings.ToUpper(device.Mac.ValueString())

		if deviceFromMist, ok := mistDevicesByClaimCode[magic]; ok {
			legacyCheckVcSiteId(deviceFromMist, mistSiteIdByVcMac)
			deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
			devicesOut = append(devicesOut, *deviceFromMist)
		} else if deviceFromMist, ok := mistDevicesbyMac[mac]; ok {
			legacyCheckVcSiteId(deviceFromMist, mistSiteIdByVcMac)
			deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
			devicesOut = append(devicesOut, *deviceFromMist)
		} else if magic != "" {
			diags.AddWarning("Device not found", fmt.Sprintf("Unable to find device with Claim Code \"%s\" in the Org Inventory", magic))
			devicesOut = append(devicesOut, NewDevicesValueNull())
		} else if mac != "" {
			diags.AddWarning("Device not found", fmt.Sprintf("Unable to find device with MAC \"%s\" in the Org Inventory", mac))
			devicesOut = append(devicesOut, NewDevicesValueNull())
		}
	}

	devices_list, e := basetypes.NewListValue(DevicesValue{}.Type(ctx), devicesOut)
	diags.Append(e...)
	newState.Devices = devices_list

	return newState, diags
}
