package resource_org_vpn

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgVpnModel) (*models.Vpn, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Vpn{}

	data.Name = plan.Name.ValueString()
	data.Paths = vpnPathsTerraformToSdk(ctx, &diags, plan.Paths)
	data.PathSelection = vpnPathSelectionTerraformToSdk(plan.PathSelection)
	data.Type = (*models.VpnModeEnum)(plan.Type.ValueStringPointer())

	return &data, diags
}

func vpnPathSelectionTerraformToSdk(d PathSelectionValue) (data *models.VpnPathSelection) {
	if d.Strategy.ValueStringPointer() != nil {
		data.Strategy = (*models.VpnPathSelectionStrategyEnum)(d.Strategy.ValueStringPointer())
	}
	return data
}

func vpnPathsPeerPathsTerraformToSdk(d basetypes.MapValue) (dataMap map[string]models.VpnPathPeerPathsPeer) {
	dataMap = make(map[string]models.VpnPathPeerPathsPeer)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PeerPathsValue)
		data := models.VpnPathPeerPathsPeer{}

		if plan.Preference.ValueInt64Pointer() != nil {
			data.Preference = models.ToPointer(int(plan.Preference.ValueInt64()))
		}

		dataMap[k] = data
	}
	return dataMap
}

func vpnPathsTrafficShapingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) (data *models.VpnPathTrafficShaping) {

	d, e := NewTrafficShapingValue(o.AttributeTypes(ctx), o.Attributes())
	if e != nil {
		diags.Append(e...)
	} else {
		if !d.ClassPercentage.IsNull() && !d.ClassPercentage.IsUnknown() {
			data.ClassPercentage = mistutils.ListOfIntTerraformToSdk(d.ClassPercentage)
		}
		if d.Enabled.ValueBoolPointer() != nil {
			data.Enabled = d.Enabled.ValueBoolPointer()
		}
		if d.MaxTxKbps.ValueInt64Pointer() != nil {
			data.MaxTxKbps = models.NewOptional(models.ToPointer(int(d.MaxTxKbps.ValueInt64())))
		}
	}
	return data
}

func vpnPathsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) (dataMap map[string]models.VpnPath) {
	dataMap = make(map[string]models.VpnPath)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PathsValue)
		data := models.VpnPath{}

		if plan.BfdProfile.ValueStringPointer() != nil {
			data.BfdProfile = (*models.VpnPathBfdProfileEnum)(plan.BfdProfile.ValueStringPointer())
		}
		if plan.BfdUseTunnelMode.ValueBoolPointer() != nil {
			data.BfdUseTunnelMode = plan.BfdUseTunnelMode.ValueBoolPointer()
		}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = plan.Ip.ValueStringPointer()
		}
		if !plan.PeerPaths.IsNull() && !plan.PeerPaths.IsUnknown() {
			data.PeerPaths = vpnPathsPeerPathsTerraformToSdk(plan.PeerPaths)
		}
		if plan.Pod.ValueInt64Pointer() != nil {
			data.Pod = models.ToPointer(int(plan.Pod.ValueInt64()))
		}
		if !plan.TrafficShaping.IsNull() && !plan.TrafficShaping.IsUnknown() {
			data.TrafficShaping = vpnPathsTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		}

		dataMap[k] = data
	}
	return dataMap
}
