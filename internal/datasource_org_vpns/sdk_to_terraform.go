package datasource_org_vpns

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.Vpn, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := vpnSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func vpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Vpn) OrgVpnsValue {
	var createdTime basetypes.Float64Value
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name types.String
	var orgId types.String
	var paths = types.MapNull(PathsValue{}.Type(ctx))
	var pathSelection = types.ObjectNull(PathSelectionValue{}.AttributeTypes(ctx))
	var vpnType types.String

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, diags, d.Paths)
	}
	if d.PathSelection != nil {
		pathSelection = vpnPathSelectionSdkToTerraform(ctx, diags, d.PathSelection)
	}
	if d.Type != nil {
		vpnType = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"created_time":   createdTime,
		"id":             id,
		"modified_time":  modifiedTime,
		"name":           name,
		"org_id":         orgId,
		"paths":          paths,
		"path_selection": pathSelection,
		"type":           vpnType,
	}
	data, e := NewOrgVpnsValue(OrgVpnsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func vpnPathSelectionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VpnPathSelection) basetypes.ObjectValue {
	var strategy types.String
	if d.Strategy != nil {
		strategy = types.StringValue(string(*d.Strategy))
	}
	dataMapValue := map[string]attr.Value{
		"strategy": strategy,
	}
	data, e := basetypes.NewObjectValue(PathSelectionValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return data
}

func vpnPathsTrafficShapingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VpnPathTrafficShaping) basetypes.ObjectValue {
	classPercentage := types.ListNull(types.Int64Type)
	var enabled basetypes.BoolValue
	var maxTxKbps basetypes.Int64Value

	if d.ClassPercentage != nil {
		classPercentage = mistutils.ListOfIntSdkToTerraform(d.ClassPercentage)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
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

func vpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPath) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var bfdProfile basetypes.StringValue
		var bfdUseTunnelMode basetypes.BoolValue
		var ip basetypes.StringValue
		peerPaths := types.MapNull(PeerPathsValue{}.Type(ctx))
		var pod basetypes.Int64Value
		trafficShaping := types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))

		if d.BfdProfile != nil {
			bfdProfile = types.StringValue(string(*d.BfdProfile))
		}
		if d.BfdUseTunnelMode != nil {
			bfdUseTunnelMode = types.BoolValue(*d.BfdUseTunnelMode)
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
			trafficShaping = vpnPathsTrafficShapingSdkToTerraform(ctx, diags, d.TrafficShaping)
		}

		dataMapValue := map[string]attr.Value{
			"bfd_profile":         bfdProfile,
			"bfd_use_tunnel_mode": bfdUseTunnelMode,
			"ip":                  ip,
			"peer_paths":          peerPaths,
			"pod":                 pod,
			"traffic_shaping":     trafficShaping,
		}
		data, e := NewPathsValue(PathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PathsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
