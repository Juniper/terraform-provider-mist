package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicVlanValue) *models.WlanDynamicVlan {

	var local_vlan_ids []models.WlanDynamicVlanLocalVlanIds
	for _, item := range plan.LocalVlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.StringValue)
		j := models.WlanDynamicVlanLocalVlanIdsContainer.FromString(i.ValueString())
		local_vlan_ids = append(local_vlan_ids, j)
	}

	vlans := make(map[string]string)
	for k, v := range plan.Vlans.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		vlans[k] = v_plan.ValueString()
	}

	var default_vlan_ids []models.WlanDynamicVlanDefaultVlanIds
	for _, item := range plan.DefaultVlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.StringValue)
		j := models.WlanDynamicVlanDefaultVlanIdsContainer.FromString(i.ValueString())
		default_vlan_ids = append(default_vlan_ids, j)
	}

	data := models.WlanDynamicVlan{}
	data.DefaultVlanIds = default_vlan_ids
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.LocalVlanIds = local_vlan_ids
	data.Type = models.ToPointer(models.WlanDynamicVlanTypeEnum(string(plan.DynamicVlanType.ValueString())))
	data.Vlans = vlans

	return &data
}
