package resource_org_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.EvpnTopology) (OrgEvpnTopologyModel, diag.Diagnostics) {
	var state OrgEvpnTopologyModel
	var diags diag.Diagnostics

	var evpn_options EvpnOptionsValue = NewEvpnOptionsValueNull()
	var id types.String
	var name types.String
	var org_id types.String
	var pod_names types.Map = types.MapNull(types.StringType)
	var switches types.Map = types.MapNull(SwitchesValue{}.Type(ctx))

	if data.EvpnOptions != nil {
		evpn_options = evpnOptionsSdkToTerraform(ctx, &diags, data.EvpnOptions)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.PodNames != nil {
		data_map := make(map[string]string)
		for k, v := range data.PodNames {
			data_map[k] = v
		}
		state_result, e := types.MapValueFrom(ctx, types.StringType, data_map)
		diags.Append(e...)
		pod_names = state_result
	}
	if data.Switches != nil {
		switches = switchesSdkToTerraform(ctx, &diags, data.Switches)
	}

	state.EvpnOptions = evpn_options
	state.Id = id
	state.Name = name
	state.OrgId = org_id
	state.PodNames = pod_names
	state.Switches = switches

	return state, diags
}
