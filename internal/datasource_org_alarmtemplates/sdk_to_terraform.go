package datasource_org_alarmtemplates

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *[]models.AlarmTemplate, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics
	for _, d := range *data {
		elem := alarmTemplateSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func alarmTemplateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.AlarmTemplate) OrgAlarmtemplatesValue {
	if data == nil {
		return OrgAlarmtemplatesValue{}
	}

	var createdTime basetypes.Float64Value
	if data.CreatedTime != nil {
		createdTime = types.Float64Value(*data.CreatedTime)
	}

	delivery := deliverySdkToTerraform(ctx, diags, &data.Delivery)

	id := types.StringValue(data.Id.String())

	var modifiedTime basetypes.Float64Value
	if data.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*data.ModifiedTime)
	}

	var name basetypes.StringValue
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}

	orgId := types.StringValue(data.OrgId.String())

	rules := rulesSdkToTerraform(ctx, diags, data.Rules)

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"delivery":      delivery,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"rules":         rules,
	}
	result, e := NewOrgAlarmtemplatesValue(OrgAlarmtemplatesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return result
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) basetypes.ObjectValue {
	if data == nil {
		return basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
	}

	var additionalEmails = mistutils.ListOfStringSdkToTerraformEmpty()
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
	result, err := basetypes.NewObjectValue(DeliveryValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}

func rulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]models.AlarmTemplateRule) basetypes.MapValue {
	rulesMap := make(map[string]attr.Value)
	for k, v := range data {
		var delivery = basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
		if v.Delivery != nil {
			var err diag.Diagnostics
			delivery, err = deliverySdkToTerraform(ctx, diags, v.Delivery).ToObjectValue(ctx)
			if err != nil {
				diags.Append(err...)
			}
		}

		var enabled types.Bool
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}

		dataMapValue := map[string]attr.Value{
			"delivery": delivery,
			"enabled":  enabled,
		}
		ruleValue, e := NewRulesValue(RulesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		rulesMap[k] = ruleValue
	}

	r, e := basetypes.NewMapValueFrom(ctx, RulesValue{}.Type(ctx), rulesMap)
	if e != nil {
		diags.Append(e...)
	}

	return r
}
