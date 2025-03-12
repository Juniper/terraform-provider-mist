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

func legacyCheckVcSiteId(device *DevicesValue, vcmMcToSite map[string]types.String) {
	if device.VcMac.ValueString() != "" && device.SiteId.ValueString() == "" {
		var vcMac = strings.ToUpper(device.VcMac.ValueString())
		if siteId, ok := vcmMcToSite[vcMac]; ok {
			device.SiteId = siteId
		}

	}
}
func legacyProcessMistInventory(
	ctx context.Context,
	diags *diag.Diagnostics,
	data *[]models.Inventory,
	mistDevicesByClaimCode *map[string]*DevicesValue,
	mistDevicesByMac *map[string]*DevicesValue,
	mistSiteIdByVcMac *map[string]types.String,
) {
	for _, d := range *data {
		var deviceprofileId = types.StringValue("")
		var hostname = types.StringValue("")
		var id basetypes.StringValue
		var mac basetypes.StringValue
		var claimCode = types.StringValue("")
		var model basetypes.StringValue
		var orgId basetypes.StringValue
		var serial basetypes.StringValue
		var siteId basetypes.StringValue
		var deviceType basetypes.StringValue
		var vcMac = types.StringValue("")
		var unclaimWhenDestroyed = types.BoolValue(false)

		if d.DeviceprofileId.Value() != nil {
			deviceprofileId = types.StringValue(*d.DeviceprofileId.Value())
		}
		if d.Magic != nil {
			claimCode = types.StringValue(*d.Magic)
		}
		if d.Mac != nil {
			mac = types.StringValue(*d.Mac)
		}
		if d.Model != nil {
			model = types.StringValue(*d.Model)
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
		if d.Type != nil {
			deviceType = types.StringValue(string(*d.Type))
		}
		if d.VcMac != nil {
			vcMac = types.StringValue(*d.VcMac)
		}
		if d.Hostname != nil {
			hostname = types.StringValue(*d.Hostname)
		}
		if d.Id != nil {
			id = types.StringValue(d.Id.String())
		}

		dataMapValue := map[string]attr.Value{
			"deviceprofile_id":       deviceprofileId,
			"hostname":               hostname,
			"id":                     id,
			"mac":                    mac,
			"claim_code":             claimCode,
			"model":                  model,
			"org_id":                 orgId,
			"serial":                 serial,
			"site_id":                siteId,
			"type":                   deviceType,
			"unclaim_when_destroyed": unclaimWhenDestroyed,
			"vc_mac":                 vcMac,
		}
		newDevice, e := NewDevicesValue(DevicesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		var nMagic = strings.ToUpper(newDevice.Magic.ValueString())
		var nMac = strings.ToUpper(newDevice.Mac.ValueString())
		(*mistDevicesByMac)[nMac] = &newDevice
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
	refInventory *OrgInventoryModel,
) (OrgInventoryModel, diag.Diagnostics) {
	var newState OrgInventoryModel
	var diags diag.Diagnostics
	var devicesOut []attr.Value
	mistDevicesByClaimCode := make(map[string]*DevicesValue)
	mistDevicesByMac := make(map[string]*DevicesValue)
	mistSiteIdByVcMac := make(map[string]types.String)

	legacyProcessMistInventory(ctx, &diags, data, &mistDevicesByClaimCode, &mistDevicesByMac, &mistSiteIdByVcMac)

	newState.OrgId = types.StringValue(orgId)

	for _, devRefInventoryAttr := range refInventory.Devices.Elements() {
		var dpi interface{} = devRefInventoryAttr
		var device = dpi.(DevicesValue)

		var magic = strings.ReplaceAll(strings.ToUpper(device.Magic.ValueString()), "-", "")
		var mac = strings.ToUpper(device.Mac.ValueString())

		var deviceFromMist *DevicesValue
		var ok bool
		if deviceFromMist, ok = mistDevicesByClaimCode[magic]; ok {
			legacyCheckVcSiteId(deviceFromMist, mistSiteIdByVcMac)
			deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
			devicesOut = append(devicesOut, *deviceFromMist)
		} else if deviceFromMist, ok = mistDevicesByMac[mac]; ok {
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

	devicesList, e := basetypes.NewListValue(DevicesValue{}.Type(ctx), devicesOut)
	diags.Append(e...)
	newState.Devices = devicesList

	return newState, diags
}
