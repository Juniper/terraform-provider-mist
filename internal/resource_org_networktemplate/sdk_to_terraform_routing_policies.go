package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func routingPolicyTermMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwRoutingPolicyTermMatching) basetypes.ObjectValue {

	var asPath = types.ListNull(types.StringType)
	var community = types.ListNull(types.StringType)
	var prefix = types.ListNull(types.StringType)
	var protocol = types.ListNull(types.StringType)

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

	dataMapValue := map[string]attr.Value{
		"as_path":   asPath,
		"community": community,
		"prefix":    prefix,
		"protocol":  protocol,
	}
	data, e := basetypes.NewObjectValue(MatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func routingPolicyTermActionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwRoutingPolicyTermAction) basetypes.ObjectValue {

	var accept basetypes.BoolValue
	var community = types.ListNull(types.StringType)
	var localPreference basetypes.StringValue
	var prependAsPath = types.ListNull(types.StringType)

	if d.Accept != nil {
		accept = types.BoolValue(*d.Accept)
	}
	if len(d.Community) > 0 {
		community = mistutils.ListOfStringSdkToTerraform(d.Community)
	}
	if d.LocalPreference != nil {
		localPreference = mistutils.ContainerAsString(d.LocalPreference)
	}
	if len(d.PrependAsPath) > 0 {
		prependAsPath = mistutils.ListOfStringSdkToTerraform(d.PrependAsPath)
	}

	dataMapValue := map[string]attr.Value{
		"accept":           accept,
		"community":        community,
		"local_preference": localPreference,
		"prepend_as_path":  prependAsPath,
	}
	data, e := basetypes.NewObjectValue(RoutingPolicyTermActionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func routingPolicyTermsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwRoutingPolicyTerm) basetypes.SetValue {
	var dataList []TermsValue

	for _, d := range l {
		var actions = types.ObjectNull(RoutingPolicyTermActionsValue{}.AttributeTypes(ctx))
		var matching = types.ObjectNull(MatchingValue{}.AttributeTypes(ctx))
		var name basetypes.StringValue

		if d.Actions != nil {
			actions = routingPolicyTermActionsSdkToTerraform(ctx, diags, *d.Actions)
		}
		if d.Matching != nil {
			matching = routingPolicyTermMatchingSdkToTerraform(ctx, diags, *d.Matching)
		}
		name = types.StringValue(d.Name)

		dataMapValue := map[string]attr.Value{
			"actions":  actions,
			"matching": matching,
			"name":     name,
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

func routingPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwRoutingPolicy) basetypes.MapValue {
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
