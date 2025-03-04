package resource_org_vpn

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, d *models.Vpn) (OrgVpnModel, diag.Diagnostics) {
	var state OrgVpnModel
	var diags diag.Diagnostics

	var id types.String
	var name types.String
	var orgId types.String
	var paths = types.MapNull(PathsValue{}.Type(ctx))
	pathSelection := NewPathSelectionValueNull()
	var vpnType types.String

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, &diags, d.Paths)
	}
	if d.PathSelection != nil {
		pathSelection = vpnPathSelectionSdkToTerraform(ctx, &diags, d.PathSelection)
	}
	if d.Type != nil {
		vpnType = types.StringValue(string(*d.Type))
	}

	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.Paths = paths
	state.PathSelection = pathSelection
	state.Type = vpnType

	return state, diags
}

func vpnPathSelectionSdkToTerraform(_ context.Context, _ *diag.Diagnostics, d *models.VpnPathSelection) (data PathSelectionValue) {

	if d.Strategy != nil {
		data.Strategy = types.StringValue(string(*d.Strategy))
	}
	return data
}

func vpnPathsPeerPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPathPeerPathsPeer) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var preference basetypes.Int64Value

		if d.Preference != nil {
			preference = types.Int64Value(int64(*d.Preference))
		}

		dataMapValue := map[string]attr.Value{
			"preference": preference,
		}

		data, e := NewPathsValue(PeerPathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PeerPathsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}

func vpnPathsTrafficShapingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VpnPathTrafficShaping) basetypes.ObjectValue {
	classPercentage := types.ListNull(types.Int64Type)
	var enabled basetypes.BoolValue
	var maxTxKbps basetypes.Int64Value

	if d.ClassPercentage != nil {
		classPercentage = misttransform.ListOfIntSdkToTerraform(d.ClassPercentage)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(bool(*d.Enabled))
	}
	if d.MaxTxKbps.Value() != nil {
		maxTxKbps = types.Int64Value(int64(*d.MaxTxKbps.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"class_percentage": classPercentage,
		"enabled":          enabled,
		"max_tx_kbps":      maxTxKbps,
	}
	data, e := basetypes.NewObjectValue(TrafficShapingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func vpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPath) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var bfdProfile basetypes.StringValue
		var ip basetypes.StringValue
		peerPaths := types.MapNull(PeerPathsValue{}.Type(ctx))
		var pod basetypes.Int64Value
		traffigShapping := types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))

		if d.BfdProfile != nil {
			bfdProfile = types.StringValue(string(*d.BfdProfile))
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.PeerPaths != nil {
			peerPaths = vpnPathsPeerPathsSdkToTerraform(ctx, diags, d.PeerPaths)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}
		if d.TrafficShaping != nil {
			traffigShapping = vpnPathsTrafficShapingSdkToTerraform(ctx, diags, d.TrafficShaping)
		}

		dataMapValue := map[string]attr.Value{
			"bfd_profile":      bfdProfile,
			"ip":               ip,
			"peer_paths":       peerPaths,
			"pod":              pod,
			"traffic_shapping": traffigShapping,
		}
		data, e := NewPathsValue(PathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PathsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
