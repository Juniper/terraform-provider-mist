package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func routingPolicyTermMatchingRouteExistsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingRouteExists) basetypes.ObjectValue {
	var route basetypes.StringValue
	var vrfName = types.StringValue("default")

	if d.Route != nil {
		route = types.StringValue(*d.Route)
	}
	if d.VrfName != nil {
		vrfName = types.StringValue(*d.VrfName)
	}

	dataMapValue := map[string]attr.Value{
		"route":    route,
		"vrf_name": vrfName,
	}
	data, e := basetypes.NewObjectValue(RouteExistsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func routingPolicyTermMatchingVpnSlaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatchingVpnPathSla) basetypes.ObjectValue {

	var maxJitter basetypes.Int64Value
	var maxLatency basetypes.Int64Value
	var maxLoss basetypes.Int64Value

	if d.MaxJitter.Value() != nil {
		maxJitter = types.Int64Value(int64(*d.MaxJitter.Value()))
	}
	if d.MaxLatency.Value() != nil {
		maxLatency = types.Int64Value(int64(*d.MaxLatency.Value()))
	}
	if d.MaxLoss.Value() != nil {
		maxLoss = types.Int64Value(int64(*d.MaxLoss.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"max_jitter":  maxJitter,
		"max_latency": maxLatency,
		"max_loss":    maxLoss,
	}
	data, e := basetypes.NewObjectValue(VpnPathSlaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
func routingPolicyTermMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermMatching) basetypes.ObjectValue {

	var asPath = types.ListNull(types.StringType)
	var community = types.ListNull(types.StringType)
	var network = types.ListNull(types.StringType)
	var prefix = types.ListNull(types.StringType)
	var protocol = types.ListNull(types.StringType)
	var routeExists = types.ObjectNull(RouteExistsValue{}.AttributeTypes(ctx))
	var vpnNeighborMac = types.ListNull(types.StringType)
	var vpnPath = types.ListNull(types.StringType)
	var vpnPathSla = types.ObjectNull(VpnPathSlaValue{}.AttributeTypes(ctx))

	if d.AsPath != nil {
		asPath = misttransform.ListOfStringSdkToTerraform(d.AsPath)
	}
	if d.Community != nil {
		community = misttransform.ListOfStringSdkToTerraform(d.Community)
	}
	if d.Network != nil {
		network = misttransform.ListOfStringSdkToTerraform(d.Network)
	}
	if d.Prefix != nil {
		prefix = misttransform.ListOfStringSdkToTerraform(d.Prefix)
	}
	if d.Protocol != nil {
		protocol = misttransform.ListOfStringSdkToTerraform(d.Protocol)
	}
	if d.RouteExists != nil {
		routeExists = routingPolicyTermMatchingRouteExistsSdkToTerraform(ctx, diags, *d.RouteExists)
	}
	if d.VpnNeighborMac != nil {
		vpnNeighborMac = misttransform.ListOfStringSdkToTerraform(d.VpnNeighborMac)
	}
	if d.VpnPath != nil {
		vpnPath = misttransform.ListOfStringSdkToTerraform(d.VpnPath)
	}
	if d.VpnPathSla != nil {
		vpnPathSla = routingPolicyTermMatchingVpnSlaSdkToTerraform(ctx, diags, *d.VpnPathSla)
	}

	dataMapValue := map[string]attr.Value{
		"as_path":          asPath,
		"community":        community,
		"network":          network,
		"prefix":           prefix,
		"protocol":         protocol,
		"route_exists":     routeExists,
		"vpn_neighbor_mac": vpnNeighborMac,
		"vpn_path":         vpnPath,
		"vpn_path_sla":     vpnPathSla,
	}
	data, e := basetypes.NewObjectValue(RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func routingPolicyTermActionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.RoutingPolicyTermAction) basetypes.ObjectValue {

	var accept basetypes.BoolValue
	var addCommunity = types.ListNull(types.StringType)
	var addTargetVrfs = types.ListNull(types.StringType)
	var aggregate = types.ListNull(types.StringType)
	var community = types.ListNull(types.StringType)
	var excludeAsPath = types.ListNull(types.StringType)
	var excludeCommunity = types.ListNull(types.StringType)
	var exportCommunitites = types.ListNull(types.StringType)
	var localPreference basetypes.StringValue
	var prependAsPath = types.ListNull(types.StringType)

	if d.Accept != nil {
		accept = types.BoolValue(*d.Accept)
	}
	if d.AddCommunity != nil {
		addCommunity = misttransform.ListOfStringSdkToTerraform(d.AddCommunity)
	}
	if d.AddTargetVrfs != nil {
		addTargetVrfs = misttransform.ListOfStringSdkToTerraform(d.AddTargetVrfs)
	}
	if d.Aggregate != nil {
		aggregate = misttransform.ListOfStringSdkToTerraform(d.Aggregate)
	}
	if d.Community != nil {
		community = misttransform.ListOfStringSdkToTerraform(d.Community)
	}
	if d.ExcludeAsPath != nil {
		excludeAsPath = misttransform.ListOfStringSdkToTerraform(d.ExcludeAsPath)
	}
	if d.ExcludeCommunity != nil {
		excludeCommunity = misttransform.ListOfStringSdkToTerraform(d.ExcludeCommunity)
	}
	if d.ExportCommunitites != nil {
		exportCommunitites = misttransform.ListOfStringSdkToTerraform(d.ExportCommunitites)
	}
	if d.LocalPreference != nil {
		localPreference = types.StringValue(*d.LocalPreference)
	}
	if d.PrependAsPath != nil {
		prependAsPath = misttransform.ListOfStringSdkToTerraform(d.PrependAsPath)
	}

	dataMapValue := map[string]attr.Value{
		"accept":              accept,
		"add_community":       addCommunity,
		"add_target_vrfs":     addTargetVrfs,
		"aggregate":           aggregate,
		"community":           community,
		"exclude_as_path":     excludeAsPath,
		"exclude_community":   excludeCommunity,
		"export_communitites": exportCommunitites,
		"local_preference":    localPreference,
		"prepend_as_path":     prependAsPath,
	}
	data, e := basetypes.NewObjectValue(ActionValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func routingPolicyTermsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RoutingPolicyTerm) basetypes.ListValue {
	var dataList []TermsValue

	for _, d := range l {
		var action = types.ObjectNull(ActionValue{}.AttributeTypes(ctx))
		var matching = types.ObjectNull(RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx))

		if d.Action != nil {
			action = routingPolicyTermActionSdkToTerraform(ctx, diags, *d.Action)
		}
		if d.Matching != nil {
			matching = routingPolicyTermMatchingSdkToTerraform(ctx, diags, *d.Matching)
		}

		dataMapValue := map[string]attr.Value{
			"action":   action,
			"matching": matching,
		}
		data, e := NewTermsValue(TermsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := TermsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func routingPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.RoutingPolicy) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var terms = types.ListNull(TermsValue{}.Type(ctx))

		if d.Terms != nil {
			terms = routingPolicyTermsSdkToTerraform(ctx, diags, d.Terms)
		}

		dataMapValue := map[string]attr.Value{
			"terms": terms,
		}
		data, e := NewRoutingPoliciesValue(RoutingPoliciesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := RoutingPoliciesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
