package resource_org_alarmtemplate

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.AlarmTemplate) (OrgAlarmtemplateModel, diag.Diagnostics) {
	var state OrgAlarmtemplateModel
	var diags diag.Diagnostics

	var delivery DeliveryValue
	var id types.String
	var name types.String
	var orgId types.String
	var rules = basetypes.NewMapNull(RulesValue{}.Type(ctx))

	delivery = deliverySdkToTerraform(ctx, &diags, &data.Delivery)
	id = types.StringValue(data.Id.String())
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	orgId = types.StringValue(data.OrgId.String())
	rules = rulesSdkToTerraform(ctx, &diags, data.Rules)

	state.Delivery = delivery
	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.Rules = rules

	return state, diags
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) DeliveryValue {
	var additionalEmails = types.ListValueMust(types.StringType, []attr.Value{})
	var enabled types.Bool
	var toOrgAdmins types.Bool
	var toSiteAdmins types.Bool

	if data != nil {
		if len(data.AdditionalEmails) > 0 {
			additionalEmails = mistutils.ListOfStringSdkToTerraform(data.AdditionalEmails)
		}
		enabled = types.BoolValue(data.Enabled)
		if data.ToOrgAdmins != nil {
			toOrgAdmins = types.BoolValue(*data.ToOrgAdmins)
		}
		if data.ToSiteAdmins != nil {
			toSiteAdmins = types.BoolValue(*data.ToSiteAdmins)
		}
	}

	dataMapValue := map[string]attr.Value{
		"additional_emails": additionalEmails,
		"enabled":           enabled,
		"to_org_admins":     toOrgAdmins,
		"to_site_admins":    toSiteAdmins,
	}
	r, e := NewDeliveryValue(DeliveryValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return r
}

func rulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]models.AlarmTemplateRule) basetypes.MapValue {
	rulesMap := make(map[string]attr.Value)
	for k, v := range data {
		var delivery = basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
		var enabled types.Bool

		if v.Delivery != nil {
			tmp, e := deliverySdkToTerraform(ctx, diags, v.Delivery).ToObjectValue(ctx)
			if e != nil {
				diags.Append(e...)
			} else {
				delivery = tmp
			}
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}

		dataMapValue := map[string]attr.Value{
			"delivery": delivery,
			"enabled":  enabled,
		}
		data, e := NewRulesValue(RulesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		rulesMap[k] = data
	}
	r, e := basetypes.NewMapValueFrom(ctx, RulesValue{}.Type(ctx), rulesMap)
	if e != nil {
		diags.Append(e...)
	}
	return r
}
