package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = &SearchInventoryBySerialFunction{}

type SearchInventoryBySerialFunction struct{}

func NewSearchInventoryBySerialFunction() function.Function {
	return &SearchInventoryBySerialFunction{}
}

func (f *SearchInventoryBySerialFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "search_inventory_by_serial"
}

func (f *SearchInventoryBySerialFunction) Definition(ctx context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: docCategoryDevices + "Retrieve a device in the `mist_org_inventory` resource based on its Serial Number",
		MarkdownDescription: "Given `mist_org_inventory` resource and a Serial Number string, will return the Device object having the provided Serial Number. " +
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
			"-> The search function is case-insensitive\n",
		Parameters: []function.Parameter{
			function.ObjectParameter{
				Name:                "inventory",
				Description:         "`mist_org_inventory` resource",
				MarkdownDescription: "`mist_org_inventory` resource",
				AttributeTypes: map[string]attr.Type{
					"org_id":    types.StringType,
					"inventory": types.MapType{ElemType: resource_org_inventory.InventoryValue{}.Type(ctx)},
				},
			},
			function.StringParameter{
				Name:                "serial",
				Description:         "Device Serial",
				MarkdownDescription: "Device Serial",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: resource_org_inventory.InventoryValue{}.AttributeTypes(ctx),
		},
	}
}

func (f *SearchInventoryBySerialFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inventory resource_org_inventory.OrgInventoryModel
	var serial string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inventory, &serial))
	if resp.Error != nil {
		return
	}

	if !inventory.Inventory.IsNull() && !inventory.Inventory.IsUnknown() && len(inventory.Inventory.Elements()) > 0 {
		for _, v := range inventory.Inventory.Elements() {
			var vi interface{} = v
			device := vi.(resource_org_inventory.InventoryValue)
			if !device.Serial.IsNull() && !device.Serial.IsUnknown() && strings.EqualFold(device.Serial.ValueString(), serial) {
				resp.Error = resp.Result.Set(ctx, &device)
				return
			}
		}
	} else {
		resp.Error = function.NewArgumentFuncError(0, "The provided inventory is empty")
	}

	resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("Unable to find a device with Serial \"%s\" in the provided inventory", serial))

}
