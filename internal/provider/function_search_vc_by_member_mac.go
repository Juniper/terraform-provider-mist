package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ function.Function = &SearchVcByMemberMacFunction{}

type SearchVcByMemberMacFunction struct{}

func NewSearchVcByMemberMacFunction() function.Function {
	return &SearchVcByMemberMacFunction{}
}

func (f *SearchVcByMemberMacFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "search_vc_by_member_mac"
}

func (f *SearchVcByMemberMacFunction) Definition(ctx context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: docCategoryDevices + "Retrieve a Virtual Chassis in the `mist_org_inventory` resource based on one of its member MAC Address",
		MarkdownDescription: "Given `mist_org_inventory` resource and a MAC Address string, will return the Device object of the Virtual Chassis having " +
			"one of its member with the provided MAC Address.  \n" +
			"If the provided MAC Address belongs to a device that is not part of a Virtual Chassis, the function will return the device itself.\n\n" +
			"The response object will contain all the information from the Mist Inventory:\n" +
			"* `claim_code`: `nil`\n" +
			"* `deviceprofile_id`: deviceprofile id if assigned\n" +
			"* `hostname`: hostname reported by the Virtual Chassis\n" +
			"* `id`: ID of the Virtual Chassis\n" +
			"* `mac`: MAC Address of the Virtual Chassis\n" +
			"* `model`: Model of the Virtual Chassis\n" +
			"* `org_id`: Org ID of the Virtual Chassis\n" +
			"* `serial`: `nil`\n" +
			"* `site_id`: Site ID of the Virtual Chassis\n" +
			"* `type`: `switch`\n" +
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
				Name:                "mac",
				Description:         "Device MAC Address",
				MarkdownDescription: "Device MAC Address",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: resource_org_inventory.InventoryValue{}.AttributeTypes(ctx),
		},
	}
}

func (f *SearchVcByMemberMacFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inventory resource_org_inventory.OrgInventoryModel
	var mac string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inventory, &mac))
	if resp.Error != nil {
		return
	}

	if inventory.Inventory.IsNull() || inventory.Inventory.IsUnknown() || len(inventory.Inventory.Elements()) == 0 {
		resp.Error = function.NewArgumentFuncError(0, "The provided inventory is empty")
		return
	}

	for _, v := range inventory.Inventory.Elements() {
		vcMember := v.(resource_org_inventory.InventoryValue)
		if !strings.EqualFold(vcMember.Mac.ValueString(), mac) {
			continue // Skip non-matching devices
		}

		if vcMember.InventoryType.ValueString() != "switch" {
			resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("The provided MAC Address \"%s\" does not belong to a switch (%s)", mac, vcMember.InventoryType.ValueString()))
			return
		}

		vc, err := f.genVirtualChassisFromInventory(ctx, vcMember)
		if err != nil {
			for _, e := range err.Errors() {
				resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(e.Detail()))
			}
			return
		}

		resp.Error = resp.Result.Set(ctx, &vc)
		return
	}

	resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("Unable to find a device with MAC Address \"%s\" in the provided inventory", mac))
}

func (f *SearchVcByMemberMacFunction) genVirtualChassisFromInventory(
	ctx context.Context,
	vcMember resource_org_inventory.InventoryValue,
) (
	resource_org_inventory.InventoryValue,
	diag.Diagnostics,
) {
	if vcMember.VcMac.IsNull() || vcMember.VcMac.IsUnknown() || vcMember.VcMac.ValueString() == "" {
		return vcMember, nil
	}

	dataMapValue := map[string]attr.Value{
		"id":                     types.StringValue(fmt.Sprintf("00000000-0000-0000-1000-%s", vcMember.VcMac.ValueString())),
		"org_id":                 vcMember.OrgId,
		"site_id":                vcMember.SiteId,
		"deviceprofile_id":       vcMember.DeviceprofileId,
		"serial":                 basetypes.StringValue{},
		"claim_code":             basetypes.StringValue{},
		"hostname":               vcMember.Hostname,
		"mac":                    vcMember.Mac,
		"vc_mac":                 vcMember.VcMac,
		"model":                  vcMember.Model,
		"type":                   vcMember.InventoryType,
		"unclaim_when_destroyed": vcMember.UnclaimWhenDestroyed,
	}
	vc, err := resource_org_inventory.NewInventoryValue(resource_org_inventory.InventoryValue{}.AttributeTypes(ctx), dataMapValue)

	return vc, err.Errors()
}
