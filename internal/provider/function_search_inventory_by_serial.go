package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ function.Function = &SearchInventoryBySerialFunction{}

type SearchInventoryBySerialFunction struct{}

func NewSearchInventoryBySerialFunction() function.Function {
	return &SearchInventoryBySerialFunction{}
}

func (f *SearchInventoryBySerialFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "search_inventory_by_serial"
}

func (f *SearchInventoryBySerialFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: docCategoryDevices + "Parse an RFC3339 timestamp string into an object",
		MarkdownDescription: "Given `mist_org_inventory.devices` Map and a Serial Number string, will return the Device object having the provided Serial Number." +
			"The response object will contain all the information from the Mist Inventory:\n" +
			"* `claim_code`: Claim Code of the device \n" +
			"* `deviceprofile_id`: deviceprofile id if assigned\n" +
			"* `hostname`: hostname reported by the device\n" +
			"* `id`: ID of the device\n" +
			"* `mac`: MAC Address of the device\n" +
			"* `model`: Model of the device\n" +
			"* `org_id`: Org ID of the device\n" +
			"* `serial`: Serial of the device\n" +
			"* `site_id`: Site ID of the device\n" +
			"* `type`: Type of device\n" +
			"* `unclaim_when_destroyed`: If the device will be unclaimed when removed from the `mist_org_inventory` resource\n" +
			"* `vc_mac`: only if `type`==`switch` of `type`==`gateway`, MAC Address of the Virtual Chassis Primary switch or the Gateway Cluster Master\n\n" +
			"> NOTE: \n> The search function is case-insensitive\n",

		Parameters: []function.Parameter{
			function.MapParameter{
				Name:                "devices",
				Description:         "Device Map from the `mist_org_inventory` resource (full path `mist_org_inventory.devices`)",
				MarkdownDescription: "Device Map from the `mist_org_inventory` resource (full path `mist_org_inventory.devices`)",
				ElementType:         resource_org_inventory.DevicesValue{}.Type(ctx),
			},
			function.StringParameter{
				Name:                "serial",
				Description:         "Device Serial",
				MarkdownDescription: "Device Serial",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: resource_org_inventory.DevicesValue{}.AttributeTypes(ctx),
		},
	}
}

func (f *SearchInventoryBySerialFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var devices basetypes.MapValue
	var serial string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &devices, &serial))
	if resp.Error != nil {
		return
	}

	if devices.IsNull() || devices.IsUnknown() || len(devices.Elements()) == 0 {
		resp.Error = function.NewArgumentFuncError(0, "The inventory provided is emtpy")
	}

	for _, v := range devices.Elements() {
		var vi interface{} = v
		device := vi.(resource_org_inventory.DevicesValue)
		if !device.Serial.IsNull() && !device.Serial.IsUnknown() && strings.EqualFold(device.Serial.ValueString(), serial) {
			resp.Error = resp.Result.Set(ctx, &device)
			return
		}
	}

	resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("Unable to find a device with Serial \"%s\" in the provided inventory", serial))

}
