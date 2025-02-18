package resource_org_vpn

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(plan *OrgVpnModel) (*models.Vpn, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Vpn{}

	data.Name = plan.Name.ValueString()
	data.Paths = vpnPathsTerraformToSdk(plan.Paths)
	return &data, diags

}

func vpnPathsTerraformToSdk(d basetypes.MapValue) map[string]models.VpnPath {
	dataMap := make(map[string]models.VpnPath)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PathsValue)
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

		dataMap[k] = data
	}
	return dataMap
}
