package resource_device_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func routingPolicyTermActionsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GwRoutingPolicyTermAction {
	data := models.GwRoutingPolicyTermAction{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewActionsValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Accept.ValueBoolPointer() != nil {
			data.Accept = models.ToPointer(plan.Accept.ValueBool())
		}
		if !plan.AddCommunity.IsNull() && !plan.AddCommunity.IsUnknown() {
			data.AddCommunity = mistutils.ListOfStringTerraformToSdk(plan.AddCommunity)
		}
		if !plan.AddTargetVrfs.IsNull() && !plan.AddTargetVrfs.IsUnknown() {
			data.AddTargetVrfs = mistutils.ListOfStringTerraformToSdk(plan.AddTargetVrfs)
		}
		if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
			data.Community = mistutils.ListOfStringTerraformToSdk(plan.Community)
		}
		if !plan.ExcludeAsPath.IsNull() && !plan.ExcludeAsPath.IsUnknown() {
			data.ExcludeAsPath = mistutils.ListOfStringTerraformToSdk(plan.ExcludeAsPath)
		}
		if !plan.ExcludeCommunity.IsNull() && !plan.ExcludeCommunity.IsUnknown() {
			data.ExcludeCommunity = mistutils.ListOfStringTerraformToSdk(plan.ExcludeCommunity)
		}
		if !plan.ExportCommunities.IsNull() && !plan.ExportCommunities.IsUnknown() {
			data.ExportCommunities = mistutils.ListOfStringTerraformToSdk(plan.ExportCommunities)
		}
		if plan.LocalPreference.ValueStringPointer() != nil {
			data.LocalPreference = models.ToPointer(models.RoutingPolicyLocalPreferenceContainer.FromString(plan.LocalPreference.ValueString()))
		}
		if !plan.PrependAsPath.IsNull() && !plan.PrependAsPath.IsUnknown() {
			data.PrependAsPath = mistutils.ListOfStringTerraformToSdk(plan.PrependAsPath)
		}
		return &data
	}
}

func routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GwRoutingPolicyTermMatchingRouteExists {
	data := models.GwRoutingPolicyTermMatchingRouteExists{}
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

func routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GwRoutingPolicyTermMatchingVpnPathSla {
	data := models.GwRoutingPolicyTermMatchingVpnPathSla{}
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

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GwRoutingPolicyTermMatching {
	data := models.GwRoutingPolicyTermMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRoutingPolicyTermMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())
		if !plan.AsPath.IsNull() && !plan.AsPath.IsUnknown() {
			var items []models.BgpAs
			for _, item := range plan.Protocol.Elements() {
				var sInterface interface{} = item
				s := sInterface.(basetypes.StringValue)
				items = append(items, models.BgpAsContainer.FromString(s.ValueString()))
			}
			data.AsPath = items
		}
		if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
			data.Community = mistutils.ListOfStringTerraformToSdk(plan.Community)
		}
		if !plan.Network.IsNull() && !plan.Network.IsUnknown() {
			data.Network = mistutils.ListOfStringTerraformToSdk(plan.Network)
		}
		if !plan.Prefix.IsNull() && !plan.Prefix.IsUnknown() {
			data.Prefix = mistutils.ListOfStringTerraformToSdk(plan.Prefix)
		}
		if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
			var items []models.GwRoutingPolicyTermMatchingProtocolEnum
			for _, item := range plan.Protocol.Elements() {
				var sInterface interface{} = item
				s := sInterface.(basetypes.StringValue)
				items = append(items, models.GwRoutingPolicyTermMatchingProtocolEnum(s.ValueString()))
			}
			data.Protocol = items
		}
		if !plan.RouteExists.IsNull() && !plan.RouteExists.IsUnknown() {
			data.RouteExists = routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx, plan.RouteExists)
		}

		if !plan.VpnNeighborMac.IsNull() && !plan.VpnNeighborMac.IsUnknown() {
			data.VpnNeighborMac = mistutils.ListOfStringTerraformToSdk(plan.VpnNeighborMac)
		}
		if !plan.VpnPath.IsNull() && !plan.VpnPath.IsUnknown() {
			data.VpnPath = mistutils.ListOfStringTerraformToSdk(plan.VpnPath)
		}

		if !plan.VpnPathSla.IsNull() && !plan.VpnPathSla.IsUnknown() {
			data.VpnPathSla = routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx, plan.VpnPathSla)
		}
		return &data
	}
}

func routingPolicyTermTerraformToSdk(ctx context.Context, d basetypes.ListValue) []models.GwRoutingPolicyTerm {
	var dataList []models.GwRoutingPolicyTerm
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TermsValue)
		data := models.GwRoutingPolicyTerm{}

		if !plan.Actions.IsNull() && !plan.Actions.IsUnknown() {
			data.Actions = routingPolicyTermActionsTerraformToSdk(ctx, plan.Actions)
		}

		if !plan.RoutingPolicyTermMatching.IsNull() && !plan.RoutingPolicyTermMatching.IsUnknown() {
			data.Matching = routingPolicyTermMatchingTerraformToSdk(ctx, plan.RoutingPolicyTermMatching)
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func routingPoliciesTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.GwRoutingPolicy {
	dataMap := make(map[string]models.GwRoutingPolicy)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(RoutingPoliciesValue)

		data := models.GwRoutingPolicy{}
		if !plan.Terms.IsNull() && !plan.Terms.IsUnknown() {
			data.Terms = routingPolicyTermTerraformToSdk(ctx, plan.Terms)
		}

		dataMap[k] = data
	}
	return dataMap
}
