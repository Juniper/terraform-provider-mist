package resource_org_networktemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func routingPolicyTermActionsTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.SwRoutingPolicyTermAction {
	data := models.SwRoutingPolicyTermAction{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRoutingPolicyTermActionsValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Accept.ValueBoolPointer() != nil {
			data.Accept = models.ToPointer(plan.Accept.ValueBool())
		}
		if !plan.Community.IsNull() && !plan.Community.IsUnknown() {
			data.Community = mistutils.ListOfStringTerraformToSdk(plan.Community)
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

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.SwRoutingPolicyTermMatching {
	data := models.SwRoutingPolicyTermMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())
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
		if !plan.Prefix.IsNull() && !plan.Prefix.IsUnknown() {
			data.Prefix = mistutils.ListOfStringTerraformToSdk(plan.Prefix)
		}
		if !plan.Protocol.IsNull() && !plan.Protocol.IsUnknown() {
			var items []models.SwRoutingPolicyTermMatchingProtocolEnum
			for _, item := range plan.Protocol.Elements() {
				var sInterface interface{} = item
				s := sInterface.(basetypes.StringValue)
				items = append(items, models.SwRoutingPolicyTermMatchingProtocolEnum(s.ValueString()))
			}
			data.Protocol = items
		}
		return &data
	}
}

func routingPolicyTermTerraformToSdk(ctx context.Context, d basetypes.SetValue) []models.SwRoutingPolicyTerm {
	var dataList []models.SwRoutingPolicyTerm
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(TermsValue)
		data := models.SwRoutingPolicyTerm{}

		if !plan.RoutingPolicyTermActions.IsNull() && !plan.RoutingPolicyTermActions.IsUnknown() {
			data.Actions = routingPolicyTermActionsTerraformToSdk(ctx, plan.RoutingPolicyTermActions)
		}

		if !plan.Matching.IsNull() && !plan.Matching.IsUnknown() {
			data.Matching = routingPolicyTermMatchingTerraformToSdk(ctx, plan.Matching)
		}

		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueString()
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func routingPoliciesTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.SwRoutingPolicy {
	dataMap := make(map[string]models.SwRoutingPolicy)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(RoutingPoliciesValue)

		data := models.SwRoutingPolicy{}
		if !plan.Terms.IsNull() && !plan.Terms.IsUnknown() {
			data.Terms = routingPolicyTermTerraformToSdk(ctx, plan.Terms)
		}

		dataMap[k] = data
	}
	return dataMap
}
