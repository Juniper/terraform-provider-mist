package resource_device_gateway

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func routingPolicyTermActionTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.RoutingPolicyTermAction {
	data := models.RoutingPolicyTermAction{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewActionValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Accept.ValueBoolPointer() != nil {
			data.Accept = models.ToPointer(plan.Accept.ValueBool())
		}
		if !plan.AddCommunity.IsNull() && !plan.AddCommunity.IsUnknown() {
			data.AddCommunity = misttransform.ListOfStringTerraformToSdk(plan.AddCommunity)
		}
		if !plan.AddTargetVrfs.IsNull() && !plan.AddTargetVrfs.IsUnknown() {
			data.AddTargetVrfs = misttransform.ListOfStringTerraformToSdk(plan.AddTargetVrfs)
		}
		if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
			data.Community = misttransform.ListOfStringTerraformToSdk(plan.Community)
		}
		if !plan.ExcludeAsPath.IsNull() && !plan.ExcludeAsPath.IsUnknown() {
			data.ExcludeAsPath = misttransform.ListOfStringTerraformToSdk(plan.ExcludeAsPath)
		}
		if !plan.ExcludeCommunity.IsNull() && !plan.ExcludeCommunity.IsUnknown() {
			data.ExcludeCommunity = misttransform.ListOfStringTerraformToSdk(plan.ExcludeCommunity)
		}
		if !plan.ExportCommunities.IsNull() && !plan.ExportCommunities.IsUnknown() {
			data.ExportCommunities = misttransform.ListOfStringTerraformToSdk(plan.ExportCommunities)
		}
		if plan.LocalPreference.ValueStringPointer() != nil {
			data.LocalPreference = models.ToPointer(plan.LocalPreference.ValueString())
		}
		if !plan.PrependAsPath.IsNull() && !plan.PrependAsPath.IsUnknown() {
			data.PrependAsPath = misttransform.ListOfStringTerraformToSdk(plan.PrependAsPath)
		}
		return &data
	}
}

func routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingRouteExists {
	data := models.RoutingPolicyTermMatchingRouteExists{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRouteExistsValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Route.ValueStringPointer() != nil {
			data.Route = models.ToPointer(plan.Route.ValueString())
		}
		if plan.VrfName.ValueStringPointer() != nil {
			data.VrfName = models.ToPointer(plan.VrfName.ValueString())
		}
		return &data
	}
}

func routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingVpnPathSla {
	data := models.RoutingPolicyTermMatchingVpnPathSla{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewVpnPathSlaValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.MaxJitter.ValueInt64Pointer() != nil {
			data.MaxJitter = models.NewOptional(models.ToPointer(int(plan.MaxJitter.ValueInt64())))
		}
		if plan.MaxLatency.ValueInt64Pointer() != nil {
			data.MaxLatency = models.NewOptional(models.ToPointer(int(plan.MaxLatency.ValueInt64())))
		}
		if plan.MaxLoss.ValueInt64Pointer() != nil {
			data.MaxLoss = models.NewOptional(models.ToPointer(int(plan.MaxLoss.ValueInt64())))
		}
		return &data
	}
}

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.RoutingPolicyTermMatching {
	data := models.RoutingPolicyTermMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRoutingPolicyTermMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())
		if !plan.AsPath.IsNull() && !plan.AsPath.IsUnknown() {
			data.AsPath = misttransform.ListOfStringTerraformToSdk(plan.AsPath)
		}
		if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
			data.Community = misttransform.ListOfStringTerraformToSdk(plan.Community)
		}
		if !plan.Network.IsNull() && !plan.Network.IsUnknown() {
			data.Network = misttransform.ListOfStringTerraformToSdk(plan.Network)
		}
		if !plan.Prefix.IsNull() && !plan.Prefix.IsUnknown() {
			data.Prefix = misttransform.ListOfStringTerraformToSdk(plan.Prefix)
		}
		if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
			data.Protocol = misttransform.ListOfStringTerraformToSdk(plan.Protocol)
		}

		if !plan.RouteExists.IsNull() && !plan.RouteExists.IsUnknown() {
			data.RouteExists = routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx, plan.RouteExists)
		}

		if !plan.VpnNeighborMac.IsNull() && !plan.VpnNeighborMac.IsUnknown() {
			data.VpnNeighborMac = misttransform.ListOfStringTerraformToSdk(plan.VpnNeighborMac)
		}
		if !plan.VpnPath.IsNull() && !plan.VpnPath.IsUnknown() {
			data.VpnPath = misttransform.ListOfStringTerraformToSdk(plan.VpnPath)
		}

		if !plan.VpnPathSla.IsNull() && !plan.VpnPathSla.IsUnknown() {
			data.VpnPathSla = routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx, plan.VpnPathSla)
		}
		return &data
	}
}

func routingPolicyTermerraformToSdk(ctx context.Context, d basetypes.ListValue) []models.RoutingPolicyTerm {
	var dataList []models.RoutingPolicyTerm
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TermsValue)
		data := models.RoutingPolicyTerm{}

		if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
			data.Action = routingPolicyTermActionTerraformToSdk(ctx, plan.Action)
		}

		if !plan.RoutingPolicyTermMatching.IsNull() && !plan.RoutingPolicyTermMatching.IsUnknown() {
			data.Matching = routingPolicyTermMatchingTerraformToSdk(ctx, plan.RoutingPolicyTermMatching)
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func routingPoliciesTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.RoutingPolicy {
	dataMap := make(map[string]models.RoutingPolicy)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(RoutingPoliciesValue)

		data := models.RoutingPolicy{}
		if !plan.Terms.IsNull() && !plan.Terms.IsUnknown() {
			data.Terms = routingPolicyTermerraformToSdk(ctx, plan.Terms)
		}

		dataMap[k] = data
	}
	return dataMap
}
