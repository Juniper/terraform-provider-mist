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

var _ function.Function = &SearchVcByMemberClaimcodeFunction{}

type SearchVcByMemberClaimcodeFunction struct{}

func NewSearchVcByMemberClaimcodeFunction() function.Function {
	return &SearchVcByMemberClaimcodeFunction{}
}

func (f *SearchVcByMemberClaimcodeFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "search_vc_by_member_claimcode"
}

func (f *SearchVcByMemberClaimcodeFunction) Definition(ctx context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: docCategoryDevices + "Retrieve a Virtual Chassis in the `mist_org_inventory` resource based on one of its member Claim Code",
		MarkdownDescription: "Given `mist_org_inventory` resource and a Claim Code string, will return the Device object of the Virtual Chassis having " +
			"one of its member with the provided Claim Code.  \n" +
			"If the provided Claim Code belongs to a device that is not part of a Virtual Chassis, the function will return the device itself.\n\n" +
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
					"devices":   types.ListType{ElemType: resource_org_inventory.DevicesValue{}.Type(ctx)},
					"inventory": types.MapType{ElemType: resource_org_inventory.InventoryValue{}.Type(ctx)},
				},
			},
			function.StringParameter{
				Name:                "claim_code",
				Description:         "Device Claim Code",
				MarkdownDescription: "Device Claim Code",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: resource_org_inventory.DevicesValue{}.AttributeTypes(ctx),
		},
	}
}

func (f *SearchVcByMemberClaimcodeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inventory resource_org_inventory.OrgInventoryModel
	var claimcode string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inventory, &claimcode))
	if resp.Error != nil {
		return
	}

	if !inventory.Devices.IsNull() && !inventory.Devices.IsUnknown() && len(inventory.Devices.Elements()) > 0 {
		for _, v := range inventory.Devices.Elements() {
			var vi interface{} = v
			vcMember := vi.(resource_org_inventory.DevicesValue)
			if !vcMember.Magic.IsNull() && !vcMember.Magic.IsUnknown() && strings.EqualFold(vcMember.Magic.ValueString(), claimcode) {
				if vcMember.DevicesType.ValueString() == "switch" {
					vc, err := f.genVirtualChassFromDevices(ctx, &vcMember)
					if err != nil {
						for _, e := range err.Errors() {
							resp.Error = function.NewFuncError(e.Detail())
						}
					}
					resp.Error = resp.Result.Set(ctx, &vc)
					return
				} else {
					resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("The provided Claim Code \"%s\" does not belong to a switch (%s)", claimcode, vcMember.DevicesType.ValueString()))
				}
			}
		}
	} else if !inventory.Inventory.IsNull() && !inventory.Inventory.IsUnknown() && len(inventory.Inventory.Elements()) > 0 {
		for _, v := range inventory.Inventory.Elements() {
			var vi interface{} = v
			vcMember := vi.(resource_org_inventory.InventoryValue)
			if !vcMember.Magic.IsNull() && !vcMember.Magic.IsUnknown() && strings.EqualFold(vcMember.Magic.ValueString(), claimcode) {
				if vcMember.InventoryType.ValueString() == "switch" {
					vc, err := f.genVirtualChassFromInventory(ctx, &vcMember)
					if err != nil {
						for _, e := range err.Errors() {
							resp.Error = function.NewFuncError(e.Detail())
						}
					}
					resp.Error = resp.Result.Set(ctx, &vc)
					return
				} else {
					resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("The provided Claim Code \"%s\" does not belong to a switch (%s)", claimcode, vcMember.InventoryType.ValueString()))
				}
			}
		}
	} else {
		resp.Error = function.NewArgumentFuncError(0, "The provided inventory is emtpy")
	}

	resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("Unable to find a device with Claim Code \"%s\" in the provided inventory", claimcode))
}

func (f *SearchVcByMemberClaimcodeFunction) genVirtualChassFromDevices(
	ctx context.Context,
	vcMember *resource_org_inventory.DevicesValue,
) (resource_org_inventory.DevicesValue, diag.Diagnostics) {
	if !vcMember.VcMac.IsNull() && !vcMember.VcMac.IsUnknown() && vcMember.VcMac.ValueString() != "" {
		var claimCode basetypes.StringValue
		var deviceprofileId = vcMember.DeviceprofileId
		var mac = vcMember.VcMac
		var model = vcMember.Model
		var orgId = vcMember.OrgId
		var serial basetypes.StringValue
		var siteId = vcMember.SiteId
		var deviceType = vcMember.DevicesType
		var vcMac = vcMember.VcMac
		var hostname = vcMember.Hostname
		var unclaimWhenDestroyed = vcMember.UnclaimWhenDestroyed
		var id = types.StringValue(fmt.Sprintf("00000000-0000-0000-1000-%s", vcMember.VcMac.ValueString()))

		dataMapValue := map[string]attr.Value{
			"deviceprofile_id":       deviceprofileId,
			"hostname":               hostname,
			"id":                     id,
			"mac":                    mac,
			"claimCode":              claimCode,
			"model":                  model,
			"org_id":                 orgId,
			"serial":                 serial,
			"site_id":                siteId,
			"type":                   deviceType,
			"unclaim_when_destroyed": unclaimWhenDestroyed,
			"vc_mac":                 vcMac,
		}
		vc, err := resource_org_inventory.NewDevicesValue(resource_org_inventory.InventoryValue{}.AttributeTypes(ctx), dataMapValue)
		return vc, err.Errors()
	}
	return *vcMember, nil
}

func (f *SearchVcByMemberClaimcodeFunction) genVirtualChassFromInventory(
	ctx context.Context,
	vcMember *resource_org_inventory.InventoryValue,
) (resource_org_inventory.InventoryValue, diag.Diagnostics) {
	if !vcMember.VcMac.IsNull() && !vcMember.VcMac.IsUnknown() && vcMember.VcMac.ValueString() != "" {
		var claimCode basetypes.StringValue
		var deviceprofileId = vcMember.DeviceprofileId
		var mac = vcMember.VcMac
		var model = vcMember.Model
		var orgId = vcMember.OrgId
		var serial basetypes.StringValue
		var siteId = vcMember.SiteId
		var deviceType = vcMember.InventoryType
		var vcMac = vcMember.VcMac
		var hostname = vcMember.Hostname
		var unclaimWhenDestroyed = vcMember.UnclaimWhenDestroyed
		var id = types.StringValue(fmt.Sprintf("00000000-0000-0000-1000-%s", vcMember.VcMac.ValueString()))

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
		vc, err := resource_org_inventory.NewInventoryValue(resource_org_inventory.InventoryValue{}.AttributeTypes(ctx), dataMapValue)
		return vc, err.Errors()
	}
	return *vcMember, nil
}
