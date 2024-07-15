package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func routingPolocyTermMatchingRouteExistsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingRouteExists) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingRouteExistsSdkToTerraform")
	var route basetypes.StringValue
	var vrf_name basetypes.StringValue = types.StringValue("default")

	if d.Route != nil {
		route = types.StringValue(*d.Route)
	}
	if d.VrfName != nil {
		vrf_name = types.StringValue(*d.VrfName)
	}

	data_map_attr_type := RouteExistsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"route":    route,
		"vrf_name": vrf_name,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func routingPolocyTermMatchingVpnSlaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingVpnPathSla) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingVpnSlaSdkToTerraform")

	var max_jitter basetypes.Int64Value
	var max_latency basetypes.Int64Value
	var max_loss basetypes.Int64Value

	if d.MaxJitter.Value() != nil {
		max_jitter = types.Int64Value(int64(*d.MaxJitter.Value()))
	}
	if d.MaxLatency.Value() != nil {
		max_latency = types.Int64Value(int64(*d.MaxLatency.Value()))
	}
	if d.MaxLoss.Value() != nil {
		max_loss = types.Int64Value(int64(*d.MaxLoss.Value()))
	}

	data_map_attr_type := VpnPathSlaValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"max_jitter":  max_jitter,
		"max_latency": max_latency,
		"max_loss":    max_loss,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
func routingPolocyTermMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatching) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermMatchingSdkToTerraform")

	var as_path basetypes.ListValue = types.ListNull(types.StringType)
	var community basetypes.ListValue = types.ListNull(types.StringType)
	var network basetypes.ListValue = types.ListNull(types.StringType)
	var prefix basetypes.ListValue = types.ListNull(types.StringType)
	var protocol basetypes.ListValue = types.ListNull(types.StringType)
	var route_exists basetypes.ObjectValue = types.ObjectNull(RouteExistsValue{}.AttributeTypes(ctx))
	var vpn_neighbor_mac basetypes.ListValue = types.ListNull(types.StringType)
	var vpn_path basetypes.ListValue = types.ListNull(types.StringType)
	var vpn_path_sla basetypes.ObjectValue = types.ObjectNull(VpnPathSlaValue{}.AttributeTypes(ctx))

	if d.AsPath != nil {
		as_path = mist_transform.ListOfStringSdkToTerraform(ctx, d.AsPath)
	}
	if d.Community != nil {
		community = mist_transform.ListOfStringSdkToTerraform(ctx, d.Community)
	}
	if d.Network != nil {
		network = mist_transform.ListOfStringSdkToTerraform(ctx, d.Network)
	}
	if d.Prefix != nil {
		prefix = mist_transform.ListOfStringSdkToTerraform(ctx, d.Prefix)
	}
	if d.Protocol != nil {
		protocol = mist_transform.ListOfStringSdkToTerraform(ctx, d.Protocol)
	}
	if d.RouteExists != nil {
		route_exists = routingPolocyTermMatchingRouteExistsSdkToTerraform(ctx, diags, *d.RouteExists)
	}
	if d.VpnNeighborMac != nil {
		vpn_neighbor_mac = mist_transform.ListOfStringSdkToTerraform(ctx, d.VpnNeighborMac)
	}
	if d.VpnPath != nil {
		vpn_path = mist_transform.ListOfStringSdkToTerraform(ctx, d.VpnPath)
	}
	if d.VpnPathSla != nil {
		vpn_path_sla = routingPolocyTermMatchingVpnSlaSdkToTerraform(ctx, diags, *d.VpnPathSla)
	}

	data_map_attr_type := RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"as_path":          as_path,
		"community":        community,
		"network":          network,
		"prefix":           prefix,
		"protocol":         protocol,
		"route_exists":     route_exists,
		"vpn_neighbor_mac": vpn_neighbor_mac,
		"vpn_path":         vpn_path,
		"vpn_path_sla":     vpn_path_sla,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func routingPolocyTermActionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermAction) basetypes.ObjectValue {
	tflog.Debug(ctx, "routingPolocyTermActionSdkToTerraform")

	var accept basetypes.BoolValue
	var add_community basetypes.ListValue = types.ListNull(types.StringType)
	var add_target_vrfs basetypes.ListValue = types.ListNull(types.StringType)
	var community basetypes.ListValue = types.ListNull(types.StringType)
	var exclude_as_path basetypes.ListValue = types.ListNull(types.StringType)
	var exclude_community basetypes.ListValue = types.ListNull(types.StringType)
	var export_communitites basetypes.ListValue = types.ListNull(types.StringType)
	var local_preference basetypes.StringValue
	var prepend_as_path basetypes.ListValue = types.ListNull(types.StringType)

	if d.Accept != nil {
		accept = types.BoolValue(*d.Accept)
	}
	if d.AddCommunity != nil {
		add_community = mist_transform.ListOfStringSdkToTerraform(ctx, d.AddCommunity)
	}
	if d.AddTargetVrfs != nil {
		add_target_vrfs = mist_transform.ListOfStringSdkToTerraform(ctx, d.AddTargetVrfs)
	}
	if d.Community != nil {
		community = mist_transform.ListOfStringSdkToTerraform(ctx, d.Community)
	}
	if d.ExcludeAsPath != nil {
		exclude_as_path = mist_transform.ListOfStringSdkToTerraform(ctx, d.ExcludeAsPath)
	}
	if d.ExcludeCommunity != nil {
		exclude_community = mist_transform.ListOfStringSdkToTerraform(ctx, d.ExcludeCommunity)
	}
	if d.ExportCommunitites != nil {
		export_communitites = mist_transform.ListOfStringSdkToTerraform(ctx, d.ExportCommunitites)
	}
	if d.LocalPreference != nil {
		local_preference = types.StringValue(*d.LocalPreference)
	}
	if d.PrependAsPath != nil {
		prepend_as_path = mist_transform.ListOfStringSdkToTerraform(ctx, d.PrependAsPath)
	}

	data_map_attr_type := ActionValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"accept":              accept,
		"add_community":       add_community,
		"add_target_vrfs":     add_target_vrfs,
		"community":           community,
		"exclude_as_path":     exclude_as_path,
		"exclude_community":   exclude_community,
		"export_communitites": export_communitites,
		"local_preference":    local_preference,
		"prepend_as_path":     prepend_as_path,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func routingPolocyTermsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RoutingPolicyTerm) basetypes.ListValue {
	tflog.Debug(ctx, "routingPolocyTermsSdkToTerraform")
	var data_list = []TermsValue{}

	for _, d := range l {
		var action basetypes.ObjectValue = types.ObjectNull(ActionValue{}.AttributeTypes(ctx))
		var matching basetypes.ObjectValue //= types.ObjectNull(RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx))

		if d.Action != nil {
			action = routingPolocyTermActionSdkToTerraform(ctx, diags, *d.Action)
		}
		if d.Matching != nil {
			matching = routingPolocyTermMatchingSdkToTerraform(ctx, diags, *d.Matching)
		}

		data_map_attr_type := TermsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":   action,
			"matching": matching,
		}
		data, e := NewTermsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := TermsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func routingPolociesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.RoutingPolicy) basetypes.MapValue {
	tflog.Debug(ctx, "routingPolociesSdkToTerraform")
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {

		var terms basetypes.ListValue = types.ListNull(TermsValue{}.Type(ctx))

		if d.Terms != nil {
			terms = routingPolocyTermsSdkToTerraform(ctx, diags, d.Terms)
		}

		data_map_attr_type := RoutingPoliciesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"terms": terms,
		}
		data, e := NewRoutingPoliciesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := RoutingPoliciesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
