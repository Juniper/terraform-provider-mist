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
	var name types.String
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}

	var diags diag.Diagnostics
	result := OrgAlarmtemplateModel{
		Delivery: deliverySdkToTerraform(ctx, &diags, &data.Delivery),
		Id:       types.StringValue(data.Id.String()),
		Name:     name,
		OrgId:    types.StringValue(data.OrgId.String()),
		Rules:    rulesSdkToTerraform(ctx, &diags, data.Rules),
	}

	return result, diags
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) DeliveryValue {
	if data == nil {
		return NewDeliveryValueNull()
	}

	var additionalEmails = types.ListValueMust(types.StringType, []attr.Value{})
	if len(data.AdditionalEmails) > 0 {
		additionalEmails = mistutils.ListOfStringSdkToTerraform(data.AdditionalEmails)
	}

	enabled := types.BoolValue(data.Enabled)

	var toOrgAdmins types.Bool
	if data.ToOrgAdmins != nil {
		toOrgAdmins = types.BoolValue(*data.ToOrgAdmins)
	}

	var toSiteAdmins types.Bool
	if data.ToSiteAdmins != nil {
		toSiteAdmins = types.BoolValue(*data.ToSiteAdmins)
	}

	dataMapValue := map[string]attr.Value{
		"additional_emails": additionalEmails,
		"enabled":           enabled,
		"to_org_admins":     toOrgAdmins,
		"to_site_admins":    toSiteAdmins,
	}
	result, err := NewDeliveryValue(DeliveryValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}

func rulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]models.AlarmTemplateRule) basetypes.MapValue {
	rules := make(map[string]attr.Value)
	for key, item := range data {
		var delivery = basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
		if item.Delivery != nil {
			var err diag.Diagnostics
			delivery, err = deliverySdkToTerraform(ctx, diags, item.Delivery).ToObjectValue(ctx)
			if err != nil {
				diags.Append(err...)
			}
		}

		var enabled types.Bool
		if item.Enabled != nil {
			enabled = types.BoolValue(*item.Enabled)
		}

		dataMap := map[string]attr.Value{
			"delivery": delivery,
			"enabled":  enabled,
		}
		item, err := NewRulesValue(RulesValue{}.AttributeTypes(ctx), dataMap)
		diags.Append(err...)

		rules[key] = item
	}

	result, err := basetypes.NewMapValueFrom(ctx, RulesValue{}.Type(ctx), rules)
	if err != nil {
		diags.Append(err...)
	}

	return result
}
