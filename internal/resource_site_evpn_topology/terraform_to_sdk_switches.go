package resource_site_evpn_topology

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) []models.EvpnTopologySwitch {

	var data []models.EvpnTopologySwitch
	for mac, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(SwitchesValue)
		data_item := models.EvpnTopologySwitch{}
		data_item.Mac = mac

		if !plan.Pod.IsNull() && !plan.Pod.IsUnknown() {
			data_item.Pod = models.ToPointer(int(plan.Pod.ValueInt64()))
		}
		if !plan.Pods.IsNull() && !plan.Pods.IsUnknown() {
			data_item.Pods = mist_transform.ListOfIntTerraformToSdk(ctx, plan.Pods)
		}
		if !plan.Role.IsNull() && !plan.Role.IsUnknown() {
			data_item.Role = models.EvpnTopologySwitchRoleEnum(plan.Role.ValueString())
		}

		data = append(data, data_item)
	}
	return data
}
