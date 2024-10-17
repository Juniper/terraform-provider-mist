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

func checkVcSiteId(device *DevicesValue, mistSiteIdByVcMac *map[string]types.String) {
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
	mistDevicesByClaimCode *map[string]*DevicesValue,
	mistDevicesbyMac *map[string]*DevicesValue,
	mistSiteIdByVcMac *map[string]types.String,
) {
	/*
		Function to process the Inventory list retrieved from Mist. This return the SetNested with all the devices
		in the inventory (used in OrgInventoryModel.Devices) and generate the list used in the other functions:
		- mistDevicesByClaimCode: map to find a device based on its claim code (if any)
		- mistDevicesByMac: map to find a device based on it's MAC Address
		- mistSiteIdByVcMac: map to find a siteId based on the VC Master/Cluster Primary MAC Address

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			data : *[]models.Inventory
				Inventory list retrived from Mist
			mistDevicesByClaimCode : *map[string]*DevicesValue
				map to find a device based on its claim code (if any). The Key is the device Claim Code, the value
				is the device Data
			mistDevicesbyMac : *map[string]*DevicesValue
				map to find a device based on its MAC Address. The Key is the device MAC Address, the value
				is the device Data
			mistSiteIdByVcMac : *map[string]*types.String
				map to find a siteId based on the VC Master/Cluster Primary MAC Address. The Key is the device MAC
				Address, the value is the Site ID

	*/

	for _, d := range *data {
		var claim_code basetypes.StringValue = types.StringValue("")
		var deviceprofile_id basetypes.StringValue = types.StringValue("")
		var mac basetypes.StringValue
		var model basetypes.StringValue
		var org_id basetypes.StringValue
		var serial basetypes.StringValue
		var site_id basetypes.StringValue = types.StringValue("")
		var device_type basetypes.StringValue
		var vc_mac basetypes.StringValue = types.StringValue("")
		var hostname basetypes.StringValue = types.StringValue("")
		var unclaim_when_destroyed basetypes.BoolValue = types.BoolValue(false)
		var id basetypes.StringValue

		if d.Magic != nil {
			claim_code = types.StringValue(*d.Magic)
		}
		if d.DeviceprofileId.Value() != nil {
			deviceprofile_id = types.StringValue(*d.DeviceprofileId.Value())
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
		if e != nil {
			diags.Append(e...)
		} else {
			var nMagic string = strings.ToUpper(newDevice.Magic.ValueString())
			var nMac string = strings.ToUpper(newDevice.Mac.ValueString())
			(*mistDevicesbyMac)[nMac] = &newDevice
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
	mistDevices *basetypes.MapValue,
	mistSiteIdByVcMac *map[string]types.String,
) basetypes.MapValue {
	/*
		Function used when using a TF import

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			mistDevices : *basetypes.SetValue
				SetNested of the devices retrieved in the Mist Inventory
			mistSiteIdByVcMac : *map[string]*types.String
				map to find a siteId based on the VC Master/Cluster Primary MAC Address. The Key is the device MAC
				Address, the value is the Site ID

		returns:
			basetypes.MapValue
				map of DeviceValue to save. Key is the device Claim Code or the MAC Address, Value is a DeviceValue
				Nested Object with the SiteId, the UnclaimWhenDestroyed bit and the Read only attributes from Mist
	*/
	newStateDevicesMap := make(map[string]attr.Value)

	for _, d := range mistDevices.Elements() {
		var di interface{} = d
		var device = di.(DevicesValue)
		checkVcSiteId(&device, mistSiteIdByVcMac)
		if !device.Magic.IsNull() && !device.Magic.IsUnknown() {
			newStateDevicesMap[device.Magic.ValueString()] = device
		} else {
			newStateDevicesMap[device.Mac.ValueString()] = device
		}
	}
	newStateDevices, e := types.MapValueFrom(ctx, DevicesValue{}.Type(ctx), newStateDevicesMap)
	diags.Append(e...)
	return newStateDevices
}

func processSync(
	ctx context.Context,
	diags *diag.Diagnostics,
	planDevices *map[string]DevicesValue,
	mistDevicesByClaimCode *map[string]*DevicesValue,
	mistDevicesByMac *map[string]*DevicesValue,
	mistSiteIdByVcMac *map[string]types.String,
) basetypes.MapValue {
	/*
		Function used when using a TF import. Generated the ByClaimCode SetNested

		parameters:
			ctx : context.Context
			diags :  *diag.Diagnostics
			planDevices : *basetypes.MapValue
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
				and the UnclaimWhenDestroyed bit
	*/
	newStateDevices := make(map[string]attr.Value)

	for deviceInfo, d := range *planDevices {
		isClaimCode, isMac := DetectDeviceInfoType(diags, deviceInfo)

		var di interface{} = d
		var device = di.(DevicesValue)

		if isClaimCode {
			if deviceFromMist, ok := (*mistDevicesByClaimCode)[strings.ToUpper(deviceInfo)]; ok {
				checkVcSiteId(deviceFromMist, mistSiteIdByVcMac)
				deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
				newStateDevices[deviceInfo] = deviceFromMist
			} else {
				diags.AddError("Device not found", fmt.Sprintf("Unable to find device with the Claim Code \"%s\" in the Org Inventory", deviceInfo))
			}
		} else if isMac {
			if deviceFromMist, ok := (*mistDevicesByMac)[strings.ToUpper(deviceInfo)]; ok {
				checkVcSiteId(deviceFromMist, mistSiteIdByVcMac)
				deviceFromMist.UnclaimWhenDestroyed = device.UnclaimWhenDestroyed
				newStateDevices[deviceInfo] = deviceFromMist
			} else {
				diags.AddError("Device not found", fmt.Sprintf("Unable to find device with the MAC Address \"%s\" in the Org Inventory", deviceInfo))
			}
		} else {
			diags.AddError(
				"Device not found",
				fmt.Sprintf("Invalid Claim Code / MAC Address format. Got: \"%s\"", deviceInfo),
			)
		}

	}
	newStateDevicesSet, e := types.MapValueFrom(ctx, DevicesValue{}.Type(ctx), newStateDevices)
	diags.Append(e...)
	return newStateDevicesSet
}

func SdkToTerraform(
	ctx context.Context,
	orgId string,
	data *[]models.Inventory,
	plan *OrgInventoryModel,
) (OrgInventoryModel, diag.Diagnostics) {
	var state OrgInventoryModel
	var diags diag.Diagnostics
	mistDevicesByClaimCode := make(map[string]*DevicesValue)
	mistDevicesbyMac := make(map[string]*DevicesValue)
	mistSiteIdByVcMac := make(map[string]types.String)

	processMistInventory(ctx, &diags, data, &mistDevicesByClaimCode, &mistDevicesbyMac, &mistSiteIdByVcMac)

	/*
		The SetNested Devices is set/updated in both cases (done above)

		If it's for an Import (no plan.OrgId), then generate the inventory with
		- basetypes.StringValue OrgId with the import orgId
		- SetNested	ByClaimCode with the list of devices with a claim code
		- SetNested ByMac with the list of devices without a claim code

		If it's for a Sync (plan.OrgId set)
		- baetypes.StringValue OrgId with the plan.OrgId
		- setNested ByClaimCode with the list of devices in the plan and in the Mist Inventory (and take the siteId from the inventory)
		- setNested ByMac with the list of devices in the plan and in the Mist Inventory (and take the siteId from the inventory)
	*/
	if plan.OrgId.ValueStringPointer() == nil {
		state.OrgId = types.StringValue(orgId)
		state.Devices = processImport(ctx, &diags, &state.Devices, &mistSiteIdByVcMac)
	} else {
		state.OrgId = plan.OrgId
		planDevicesMap := GenDeviceMap(&plan.Devices)
		state.Devices = processSync(ctx, &diags, &planDevicesMap, &mistDevicesByClaimCode, &mistDevicesbyMac, &mistSiteIdByVcMac)
	}

	return state, diags
}
