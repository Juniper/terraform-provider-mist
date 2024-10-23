package resource_org_inventory

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func GenDeviceMap(devices *basetypes.MapValue) map[string]InventoryValue {
	/*
		Generate a map[string]InventoryValue from the basetypes.MapValue

		parameters:
			deivces: *basetypes.SetValue
				SetNested with each devices in the Inventory of the Mist Org

		returns:
			map[string]InventoryValue
				key is the device Claim Code or MAC Address, value is the DeviceValue
	*/
	devices_map := make(map[string]InventoryValue)
	for key, v := range devices.Elements() {
		var dsi interface{} = v
		var dev = dsi.(InventoryValue)
		devices_map[key] = dev
	}
	return devices_map
}

func DetectDeviceInfoType(diags *diag.Diagnostics, deviceInfo string) (bool, bool) {
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
	re_claimcode := `^[0-9a-zA-Z]{15}$`
	re_mac := `^[0-9a-fA-F]{12}$`
	if is_valid, _ := regexp.MatchString(re_claimcode, deviceInfo); is_valid {
		return true, false
	} else if is_valid, _ := regexp.MatchString(re_mac, deviceInfo); is_valid {
		return false, true
	} else {
		diags.AddError(
			"Invalid device Key in \"org_inventory\" resource",
			fmt.Sprintf("Unable to identidy the type of key (claim code / mac address) for the device. got: \"%s\"", deviceInfo),
		)
	}
	return false, false
}
