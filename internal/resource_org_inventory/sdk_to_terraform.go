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

func checkVcSiteId(device *InventoryValue, mistSiteIdByVcMac *map[string]types.String) {
	/*
		Function to check if the device is part of a VC / Cluster, and update the device
		siteId with the Master/Primary device SiteId

		parameters:
			device : *DeviceValue
				the device to check
			mistSiteIdByVcMac : *map[string]types.String
				a map where the key is the device MAC Address and the value the device SiteId
	*/
	if device.VcMac.ValueString() != "" && device.SiteId.ValueString() == "" {
		if siteId, ok := (*mistSiteIdByVcMac)[strings.ToUpper(device.VcMac.ValueString())]; ok {
			device.SiteId = siteId
		}
	}
}

func processMistInventory(
	ctx context.Context,
	diags *diag.Diagnostics,
	data *[]models.Inventory,
	mistDevicesByClaimCode *map[string]*InventoryValue,
	mistDevicesByMac *map[string]*InventoryValue,
	mistSiteIdByVcMac *map[string]types.String,
) {
	/*
		Function to process the Inventory list retrieved from Mist. This return the SetNested with all the devices
		in the inventory (used in OrgInventoryModel.Devices) and generate the list used in the other functions:
		- mistDevicesByClaimCode: map to find a device based on its claim code (if any)
		- mistDevicesByMac: map to find a device based on its MAC Address
		- mistSiteIdByVcMac: map to find a siteId based on the VC Master/Cluster Primary MAC Address

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			data : *[]models.Inventory
				Inventory list retrieved from Mist
			mistDevicesByClaimCode : *map[string]*InventoryValue
				map to find a device based on its claim code (if any). The Key is the device Claim Code, the value
				is the device Data
			mistDevicesByMac : *map[string]*InventoryValue
				map to find a device based on its MAC Address. The Key is the device MAC Address, the value
				is the device Data
			mistSiteIdByVcMac : *map[string]*types.String
				map to find a siteId based on the VC Master/Cluster Primary MAC Address. The Key is the device MAC
				Address, the value is the Site ID

	*/

	for _, d := range *data {
		var claimCode = types.StringValue("")
		var deviceprofileId = types.StringValue("")
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var orgId basetypes.StringValue
		var serial basetypes.StringValue
		var siteId basetypes.StringValue
		var deviceType basetypes.StringValue
		var vcMac = types.StringValue("")
		var hostname = types.StringValue("")
		var unclaimWhenDestroyed = types.BoolValue(false)
		var id basetypes.StringValue

		if d.Magic != nil {
			claimCode = types.StringValue(*d.Magic)
		}
		if d.DeviceprofileId.Value() != nil {
			deviceprofileId = types.StringValue(*d.DeviceprofileId.Value())
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
		newDevice, e := NewInventoryValue(InventoryValue{}.AttributeTypes(ctx), dataMapValue)
		if e != nil {
			diags.Append(e...)
		} else {
			var nMagic = strings.ToUpper(newDevice.Magic.ValueString())
			var nMac = strings.ToUpper(newDevice.Mac.ValueString())
			(*mistDevicesByMac)[nMac] = &newDevice
			if nMagic != "" {
				(*mistDevicesByClaimCode)[nMagic] = &newDevice
			}
			if newDevice.VcMac.Equal(newDevice.Mac) {
				(*mistSiteIdByVcMac)[nMac] = newDevice.SiteId
			}
		}
	}
}

func processImport(
	ctx context.Context,
	diags *diag.Diagnostics,
	mistDevicesByMac *map[string]*InventoryValue,
	mistSiteIdByVcMac *map[string]types.String,
) basetypes.MapValue {
	/*
		Function used when using a TF import

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			mistDevicesByMac : *map[string]*InventoryValue,
				map of the devices retrieved in the Mist Inventory
			mistSiteIdByVcMac : *map[string]*types.String
				map to find a siteId based on the VC Master/Cluster Primary MAC Address. The Key is the device MAC
				Address, the value is the Site ID

		returns:
			basetypes.MapValue
				map of DeviceValue to save. Key is the device Claim Code or the MAC Address, Value is a DeviceValue
				Nested Object with the SiteId, the UnclaimWhenDestroyed bit and the Read only attributes from Mist
	*/
	newStateDevicesMap := make(map[string]attr.Value)

	for _, device := range *mistDevicesByMac {
		checkVcSiteId(device, mistSiteIdByVcMac)
		if device.Magic.ValueStringPointer() == nil || len(device.Magic.ValueString()) == 0 {
			newStateDevicesMap[device.Mac.ValueString()] = device
		} else {
			newStateDevicesMap[device.Magic.ValueString()] = device
		}
	}
	newStateDevices, e := types.MapValueFrom(ctx, InventoryValue{}.Type(ctx), newStateDevicesMap)
	diags.Append(e...)
	return newStateDevices
}

func vcSiteIdValidation(
	diags *diag.Diagnostics,
	deviceInfo string,
	deviceSiteId string,
	vcSiteId string,
) {
	if deviceSiteId != vcSiteId && deviceSiteId != "" {
		diags.AddError(
			"Unable to claim a device in \"mist_org_inventory\"",
			fmt.Sprintf(
				"The device mist_org_inventory.inventory[%s] cannot be claimed and assigned to the site %s"+
					" because it is part of a Virtual Chassis already assigned to the site %s.\n"+
					"Please update mist_org_inventory.inventory[%s].site_id with the Virtual Chassis site_id",
				deviceInfo, deviceSiteId, vcSiteId, deviceInfo,
			),
		)
	} else if deviceSiteId != vcSiteId && deviceSiteId == "" {
		diags.AddError(
			"Unable to claim a device in \"mist_org_inventory\"",
			fmt.Sprintf(
				"The device mist_org_inventory.inventory[%s] cannot be claimed"+
					" because it is part of a Virtual Chassis already assigned to the site %s.\n"+
					"Please update mist_org_inventory.inventory[%s].site_id with the Virtual Chassis site_id",
				deviceInfo, vcSiteId, deviceInfo,
			),
		)
	}
}

func processSync(
	ctx context.Context,
	diags *diag.Diagnostics,
	refInventoryDevices *map[string]*InventoryValue,
	mistDevicesByClaimCode *map[string]*InventoryValue,
	mistDevicesByMac *map[string]*InventoryValue,
	mistSiteIdByVcMac *map[string]types.String,
) (newStateDevicesSet basetypes.MapValue) {
	/*
		Function used when using a TF import. Generated the ByClaimCode SetNested

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			ref_inventoryDevices : *basetypes.MapValue
				map of ByClaimCode to save. Key is the device Claim Code, Value is a Nested Object with the SiteId
				and the UnclaimWhenDestroyed bit
			mistDevices : *basetypes.SetValue
				SetNested of the devices retrieved in the Mist Inventory
			mistSiteIdByVcMac : *map[string]*types.String
				map to find a siteId based on the VC Master/Cluster Primary MAC Address. The Key is the device MAC
				Address, the value is the Site ID

		returns:
			basetypes.MapValue
				map of ByClaimCode to save. Key is the device Claim Code, Value is a Nested Object with the SiteId
				and the unclaimWhenDestroyed bit
	*/
	newStateDevices := make(map[string]attr.Value)

	for deviceInfo, device := range *refInventoryDevices {
		isClaimCode, isMac := DetectDeviceInfoType(diags, strings.ToUpper(deviceInfo))

		if isClaimCode {
			if deviceFromMist, ok := (*mistDevicesByClaimCode)[strings.ToUpper(deviceInfo)]; ok {
				checkVcSiteId(deviceFromMist, mistSiteIdByVcMac)
				deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
				newStateDevices[deviceInfo] = deviceFromMist
				if deviceFromMist.InventoryType.ValueString() == "switch" {
					vcSiteIdValidation(diags, deviceInfo, device.SiteId.ValueString(), deviceFromMist.SiteId.ValueString())
				}
			}
		} else if isMac {
			if deviceFromMist, ok := (*mistDevicesByMac)[strings.ToUpper(deviceInfo)]; ok {
				checkVcSiteId(deviceFromMist, mistSiteIdByVcMac)
				deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
				newStateDevices[deviceInfo] = deviceFromMist
				if deviceFromMist.InventoryType.ValueString() == "switch" {
					vcSiteIdValidation(diags, deviceInfo, device.SiteId.ValueString(), deviceFromMist.SiteId.ValueString())
				}
			}
		} else {
			diags.AddError(
				"Device not found",
				fmt.Sprintf("Invalid Claim Code / MAC Address format. Got: \"%s\"", deviceInfo),
			)
		}

	}
	newStateDevicesSet, e := types.MapValueFrom(ctx, InventoryValue{}.Type(ctx), newStateDevices)
	diags.Append(e...)
	return newStateDevicesSet
}

func mapSdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Inventory,
	refInventory *OrgInventoryModel,
) (state OrgInventoryModel, diags diag.Diagnostics) {
	mistDevicesByClaimCode := make(map[string]*InventoryValue)
	mistDevicesByMac := make(map[string]*InventoryValue)
	mistSiteIdByVcMac := make(map[string]types.String)

	processMistInventory(ctx, &diags, data, &mistDevicesByClaimCode, &mistDevicesByMac, &mistSiteIdByVcMac)

	/*
		The SetNested Devices is set/updated in both cases (done above)

		If it's for an Import (no ref_inventory.OrgId), then generate the inventory with
		- basetypes.StringValue OrgId with the import orgId
		- SetNested	ByClaimCode with the list of devices with a claim code
		- SetNested ByMac with the list of devices without a claim code

		If it's for a Sync (ref_inventory.OrgId set)
		- basetypes.StringValue OrgId with the ref_inventory.OrgId
		- setNested ByClaimCode with the list of devices in the ref_inventory and in the Mist Inventory (and take the siteId from the inventory)
		- setNested ByMac with the list of devices in the ref_inventory and in the Mist Inventory (and take the siteId from the inventory)
	*/
	if refInventory.OrgId.ValueStringPointer() == nil {
		state.OrgId = types.StringValue(orgId)
		state.Inventory = processImport(ctx, &diags, &mistDevicesByMac, &mistSiteIdByVcMac)
	} else {
		state.OrgId = refInventory.OrgId
		refInventoryDevicesMap := GenDeviceMap(&refInventory.Inventory)
		state.Inventory = processSync(ctx, &diags, &refInventoryDevicesMap, &mistDevicesByClaimCode, &mistDevicesByMac, &mistSiteIdByVcMac)
	}

	return state, diags
}

func SdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Inventory,
	refInventory *OrgInventoryModel,
) (OrgInventoryModel, diag.Diagnostics) {
	if !refInventory.Devices.IsNull() && !refInventory.Devices.IsUnknown() {
		state, diags := legacySdkToTerraform(ctx, orgId, data, refInventory)
		state.Inventory = basetypes.NewMapNull(InventoryValue{}.Type(ctx))
		return state, diags
	} else {
		state, diags := mapSdkToTerraform(ctx, orgId, data, refInventory)
		state.Devices = basetypes.NewListNull(DevicesValue{}.Type(ctx))
		return state, diags
	}
}
