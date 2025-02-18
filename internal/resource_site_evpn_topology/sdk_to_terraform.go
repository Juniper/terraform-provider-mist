package resource_site_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.EvpnTopology) (SiteEvpnTopologyModel, diag.Diagnostics) {
	var state SiteEvpnTopologyModel
	var diags diag.Diagnostics

	var evpnOptions = NewEvpnOptionsValueNull()
	var id types.String
	var name types.String
	var orgId types.String
	var podNames = types.MapNull(types.StringType)
	var siteId types.String
	var switches = types.MapNull(SwitchesValue{}.Type(ctx))

	if data.EvpnOptions != nil {
		evpnOptions = evpnOptionsSdkToTerraform(ctx, &diags, data.EvpnOptions)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.PodNames != nil {
		dataMap := make(map[string]string)
		for k, v := range data.PodNames {
			dataMap[k] = v
		}
		stateResult, e := types.MapValueFrom(ctx, types.StringType, dataMap)
		diags.Append(e...)
		podNames = stateResult
	}
	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}
	if data.Switches != nil {
		switches = switchesSdkToTerraform(ctx, &diags, data.Switches)
	}

	state.EvpnOptions = evpnOptions
	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.PodNames = podNames
	state.SiteId = siteId
	state.Switches = switches

	return state, diags
}
