package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicPskTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicPskValue) *models.WlanDynamicPsk {

	var vlan_ids []int
	for _, item := range plan.VlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.Int64Value)
		j := int(i.ValueInt64())
		vlan_ids = append(vlan_ids, j)
	}

	data := models.WlanDynamicPsk{}
	data.DefaultPsk = plan.DefaultPsk.ValueStringPointer()
	data.DefaultVlanId = models.NewOptional(models.ToPointer(int(plan.DefaultVlanId.ValueInt64())))
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.ForceLookup = plan.ForceLookup.ValueBoolPointer()
	data.Source = models.ToPointer(models.DynamicPskSourceEnum(string(plan.Source.ValueString())))
	data.VlanIds = vlan_ids

	return &data
}
