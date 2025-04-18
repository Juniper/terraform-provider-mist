package resource_site_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *SiteEvpnTopologyModel) (*models.EvpnTopology, diag.Diagnostics) {
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
		dataMap := make(map[string]string)
		for k, v := range plan.PodNames.Elements() {
			var sInterface interface{} = v
			s := sInterface.(basetypes.StringValue)
			dataMap[k] = s.ValueString()
		}
		data.PodNames = dataMap
	} else {
		unset["-pod_names"] = ""
	}
	if !plan.Switches.IsNull() && !plan.Switches.IsUnknown() {
		data.Switches = switchTerraformToSdk(plan.Switches)
	} else {
		unset["-switches"] = ""
	}

	data.Overwrite = models.ToPointer(true)

	data.AdditionalProperties = unset
	return &data, diags
}
