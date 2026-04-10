package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func routingPolicyTermMatchingRouteExistsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GwRoutingPolicyTermMatchingRouteExists) basetypes.ObjectValue {
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
func routingPolicyTermMatchingVpnSlaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GwRoutingPolicyTermMatchingVpnPathSla) basetypes.ObjectValue {

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
func routingPolicyTermMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GwRoutingPolicyTermMatching) basetypes.ObjectValue {

	var asPath = types.ListNull(types.StringType)
	var community = types.ListNull(types.StringType)
	var network = types.ListNull(types.StringType)
	var prefix = types.ListNull(types.StringType)
	var protocol = types.ListNull(types.StringType)
	var routeExists = types.ObjectNull(RouteExistsValue{}.AttributeTypes(ctx))
	var vpnNeighborMac = types.ListNull(types.StringType)
	var vpnPath = types.ListNull(types.StringType)
	var vpnPathSla = types.ObjectNull(VpnPathSlaValue{}.AttributeTypes(ctx))

	if len(d.AsPath) > 0 {
		var items []attr.Value
		for _, item := range d.AsPath {
			items = append(items, mistutils.ContainerAsString(&item))
		}
		asPath, _ = types.ListValue(basetypes.StringType{}, items)
	}
	if len(d.Community) > 0 {
		community = mistutils.ListOfStringSdkToTerraform(d.Community)
	}
	if len(d.Network) > 0 {
		network = mistutils.ListOfStringSdkToTerraform(d.Network)
	}
	if len(d.Prefix) > 0 {
		prefix = mistutils.ListOfStringSdkToTerraform(d.Prefix)
	}
	if len(d.Protocol) > 0 {
		var items []attr.Value
		for _, item := range d.Protocol {
			items = append(items, types.StringValue(string(item)))
		}
		protocol, _ = types.ListValue(basetypes.StringType{}, items)
	}
	if d.RouteExists != nil {
		routeExists = routingPolicyTermMatchingRouteExistsSdkToTerraform(ctx, diags, *d.RouteExists)
	}
	if len(d.VpnNeighborMac) > 0 {
		vpnNeighborMac = mistutils.ListOfStringSdkToTerraform(d.VpnNeighborMac)
	}
	if len(d.VpnPath) > 0 {
		vpnPath = mistutils.ListOfStringSdkToTerraform(d.VpnPath)
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
func routingPolicyTermActionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GwRoutingPolicyTermAction) basetypes.ObjectValue {

	var accept basetypes.BoolValue
	var addCommunity = types.ListNull(types.StringType)
	var addTargetVrfs = types.ListNull(types.StringType)
	var community = types.ListNull(types.StringType)
	var excludeAsPath = types.ListNull(types.StringType)
	var excludeCommunity = types.ListNull(types.StringType)
	var exportCommunities = types.ListNull(types.StringType)
	var localPreference basetypes.StringValue
	var prependAsPath = types.ListNull(types.StringType)

	if d.Accept != nil {
		accept = types.BoolValue(*d.Accept)
	}
	if len(d.AddCommunity) > 0 {
		addCommunity = mistutils.ListOfStringSdkToTerraform(d.AddCommunity)
	}
	if len(d.AddTargetVrfs) > 0 {
		addTargetVrfs = mistutils.ListOfStringSdkToTerraform(d.AddTargetVrfs)
	}
	if len(d.Community) > 0 {
		community = mistutils.ListOfStringSdkToTerraform(d.Community)
	}
	if len(d.ExcludeAsPath) > 0 {
		excludeAsPath = mistutils.ListOfStringSdkToTerraform(d.ExcludeAsPath)
	}
	if len(d.ExcludeCommunity) > 0 {
		excludeCommunity = mistutils.ListOfStringSdkToTerraform(d.ExcludeCommunity)
	}
	if len(d.ExportCommunities) > 0 {
		exportCommunities = mistutils.ListOfStringSdkToTerraform(d.ExportCommunities)
	}
	if d.LocalPreference != nil {
		localPreference = mistutils.ContainerAsString(d.LocalPreference)
	}
	if len(d.PrependAsPath) > 0 {
		prependAsPath = mistutils.ListOfStringSdkToTerraform(d.PrependAsPath)
	}

	dataMapValue := map[string]attr.Value{
		"accept":             accept,
		"add_community":      addCommunity,
		"add_target_vrfs":    addTargetVrfs,
		"community":          community,
		"exclude_as_path":    excludeAsPath,
		"exclude_community":  excludeCommunity,
		"export_communities": exportCommunities,
		"local_preference":   localPreference,
		"prepend_as_path":    prependAsPath,
	}
	data, e := basetypes.NewObjectValue(ActionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func routingPolicyTermsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.GwRoutingPolicyTerm) basetypes.SetValue {
	var dataList []TermsValue

	for _, d := range l {
		var actions = types.ObjectNull(ActionsValue{}.AttributeTypes(ctx))
		var matching = types.ObjectNull(RoutingPolicyTermMatchingValue{}.AttributeTypes(ctx))

		if d.Actions != nil {
			actions = routingPolicyTermActionsSdkToTerraform(ctx, diags, *d.Actions)
		}
		if d.Matching != nil {
			matching = routingPolicyTermMatchingSdkToTerraform(ctx, diags, *d.Matching)
		}

		dataMapValue := map[string]attr.Value{
			"actions":  actions,
			"matching": matching,
		}
		data, e := NewTermsValue(TermsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := TermsValue{}.Type(ctx)
	r, e := types.SetValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func routingPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GwRoutingPolicy) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var terms = types.SetNull(TermsValue{}.Type(ctx))

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
