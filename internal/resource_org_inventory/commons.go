package resource_org_inventory

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"regexp"
)

func GenDeviceMap(devices *basetypes.MapValue) (devicesMap map[string]*InventoryValue) {
	/*
		Generate a map[string]InventoryValue from the basetypes.MapValue

		parameters:
			devices: *basetypes.SetValue
				SetNested with each device in the Inventory of the Mist Org

		returns:
			map[string]InventoryValue
				key is the device Claim Code or MAC Address, value is the DeviceValue
	*/
	devicesMap = make(map[string]*InventoryValue)
	for key, v := range devices.Elements() {
		var dsi interface{} = v
		var dev = dsi.(InventoryValue)
		devicesMap[key] = &dev
	}
	return devicesMap
}

func DetectDeviceInfoType(diags *diag.Diagnostics, deviceInfo string) (isClaimcode bool, isMac bool) {
	/*
		Function to detect the type of info (Claim Code or MAC Address)

		parameters
			diags : *diag.Diagnostics
			deviceInfo : string
				the string to test

		returns:
		bool
			true if it's a Claim Code
		bool
			true if it's a MAC Address

	*/
	reClaimcode := `^[0-9a-zA-Z]{15}$`
	reMac := `^[0-9a-fA-F]{12}$`
	if isValid, _ := regexp.MatchString(reClaimcode, deviceInfo); isValid {
		return true, false
	} else if isValid, _ = regexp.MatchString(reMac, deviceInfo); isValid {
		return false, true
	} else {
		diags.AddError(
			"Invalid device Key in \"org_inventory\" resource",
			fmt.Sprintf("Unable to identify the type of key (claim code / mac address) for the device. got: \"%s\"", deviceInfo),
		)
	}
	return false, false
}
