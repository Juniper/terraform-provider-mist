package resource_org_vpn

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgVpnModel) (*models.Vpn, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Vpn{}

	data.Name = plan.Name.ValueString()
	data.Paths = vpnPathsTerraformToSdk(ctx, &diags, plan.Paths)
	return &data, diags

}

func vpnPathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.VpnPath {
	data_map := make(map[string]models.VpnPath)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PathsValue)
		data := models.VpnPath{}

		if plan.BfdProfile.ValueStringPointer() != nil {
			data.BfdProfile = (*models.VpnPathBfdProfileEnum)(plan.BfdProfile.ValueStringPointer())
		}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = plan.Ip.ValueStringPointer()
		}
		if plan.Pod.ValueInt64Pointer() != nil {
			data.Pod = models.ToPointer(int(plan.Pod.ValueInt64()))
		}

		data_map[k] = data
	}
	return data_map
}
