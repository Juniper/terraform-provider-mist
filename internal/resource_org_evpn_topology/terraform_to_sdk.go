package resource_org_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgEvpnTopologyModel) (*models.EvpnTopology, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.EvpnTopology{}

	if !plan.EvpnOptions.IsNull() && !plan.EvpnOptions.IsUnknown() {
		data.EvpnOptions = evpnOptionsTerraformToSdk(ctx, &diags, plan.EvpnOptions)
	} else {
		unset["-evpn_options"] = ""
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = plan.Name.ValueStringPointer()
	}

	if !plan.PodNames.IsNull() && !plan.PodNames.IsUnknown() && len(plan.PodNames.Elements()) > 0 {
		for k, v := range plan.PodNames.Elements() {
			var s_interface interface{} = v
			s := s_interface.(basetypes.StringValue)
			data.PodNames[k] = s.ValueString()
		}
	} else {
		unset["-pod_names"] = ""
	}
	if !plan.Switches.IsNull() && !plan.Switches.IsUnknown() {
		data.Switches = switchTerraformToSdk(ctx, &diags, plan.Switches)
	} else {
		unset["-switches"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
