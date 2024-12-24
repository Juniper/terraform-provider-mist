package resource_org_deviceprofile_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func routingPolicyTermActionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermAction {
	data := models.RoutingPolicyTermAction{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewActionValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Accept.ValueBoolPointer() != nil {
				data.Accept = models.ToPointer(plan.Accept.ValueBool())
			}
			if !plan.AddCommunity.IsNull() && !plan.AddCommunity.IsUnknown() {
				data.AddCommunity = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddCommunity)
			}
			if !plan.AddTargetVrfs.IsNull() && !plan.AddTargetVrfs.IsUnknown() {
				data.AddTargetVrfs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddTargetVrfs)
			}
			if !plan.Aggregate.IsNull() && !plan.Aggregate.IsUnknown() {
				data.Aggregate = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Aggregate)
			}
			if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
				data.Community = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community)
			}
			if !plan.ExcludeAsPath.IsNull() && !plan.ExcludeAsPath.IsUnknown() {
				data.ExcludeAsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeAsPath)
			}
			if !plan.ExcludeCommunity.IsNull() && !plan.ExcludeCommunity.IsUnknown() {
				data.ExcludeCommunity = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeCommunity)
			}
			if plan.LocalPreference.ValueStringPointer() != nil {
				data.LocalPreference = models.ToPointer(plan.LocalPreference.ValueString())
			}
			if !plan.PrependAsPath.IsNull() && !plan.PrependAsPath.IsUnknown() {
				data.PrependAsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.PrependAsPath)
			}
		}
	}
	return &data
}

func routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingRouteExists {
	data := models.RoutingPolicyTermMatchingRouteExists{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewRouteExistsValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Route.ValueStringPointer() != nil {
				data.Route = models.ToPointer(plan.Route.ValueString())
			}
			if plan.VrfName.ValueStringPointer() != nil {
				data.VrfName = models.ToPointer(plan.VrfName.ValueString())
			}
		}
	}
	return &data
}

func routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingVpnPathSla {
	data := models.RoutingPolicyTermMatchingVpnPathSla{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewVpnPathSlaValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.MaxJitter.ValueInt64Pointer() != nil {
				data.MaxJitter = models.NewOptional(models.ToPointer(int(plan.MaxJitter.ValueInt64())))
			}
			if plan.MaxLatency.ValueInt64Pointer() != nil {
				data.MaxLatency = models.NewOptional(models.ToPointer(int(plan.MaxLatency.ValueInt64())))
			}
			if plan.MaxLoss.ValueInt64Pointer() != nil {
				data.MaxLoss = models.NewOptional(models.ToPointer(int(plan.MaxLoss.ValueInt64())))
			}
		}
	}
	return &data
}

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatching {
	data := models.RoutingPolicyTermMatching{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewRoutingPolicyTermMatchingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.AsPath.IsNull() && !plan.AsPath.IsUnknown() {
				data.AsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AsPath)
			}
			if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
				data.Community = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community)
			}
			if !plan.Network.IsNull() && !plan.Network.IsUnknown() {
				data.Network = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Network)
			}
			if !plan.Prefix.IsNull() && !plan.Prefix.IsUnknown() {
				data.Prefix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Prefix)
			}
			if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
				data.Protocol = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Protocol)
			}

			if !plan.RouteExists.IsNull() && !plan.RouteExists.IsUnknown() {
				data.RouteExists = routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx, diags, plan.RouteExists)
			}

			if !plan.VpnNeighborMac.IsNull() && !plan.VpnNeighborMac.IsUnknown() {
				data.VpnNeighborMac = mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnNeighborMac)
			}
			if !plan.VpnPath.IsNull() && !plan.VpnPath.IsUnknown() {
				data.VpnPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnPath)
			}

			if !plan.VpnPathSla.IsNull() && !plan.VpnPathSla.IsUnknown() {
				data.VpnPathSla = routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx, diags, plan.VpnPathSla)
			}
		}
	}
	return &data
}

func routingPolicyTermerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RoutingPolicyTerm {
	var data_list []models.RoutingPolicyTerm
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TermsValue)
		data := models.RoutingPolicyTerm{}

		if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
			data.Action = routingPolicyTermActionTerraformToSdk(ctx, diags, plan.Action)
		}

		if !plan.RoutingPolicyTermMatching.IsNull() && !plan.RoutingPolicyTermMatching.IsUnknown() {
			data.Matching = routingPolicyTermMatchingTerraformToSdk(ctx, diags, plan.RoutingPolicyTermMatching)
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func routingPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.RoutingPolicy {
	data_map := make(map[string]models.RoutingPolicy)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(RoutingPoliciesValue)

		data := models.RoutingPolicy{}
		if !plan.Terms.IsNull() && !plan.Terms.IsUnknown() {
			data.Terms = routingPolicyTermerraformToSdk(ctx, diags, plan.Terms)
		}

		data_map[k] = data
	}
	return data_map
}
