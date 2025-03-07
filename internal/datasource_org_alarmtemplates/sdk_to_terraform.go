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

func SdkToTerraform(ctx context.Context, l *[]models.AlarmTemplate, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := alarmTempalteSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func alarmTempalteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AlarmTemplate) OrgAlarmtemplatesValue {

	var createdTime basetypes.Float64Value
	var delivery = types.ObjectNull(DeliveryValue{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var rules = basetypes.NewMapNull(RulesValue{}.Type(ctx))

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	delivery = deliverySdkToTerraform(ctx, diags, &d.Delivery)
	id = types.StringValue(d.Id.String())
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	orgId = types.StringValue(d.OrgId.String())
	rules = rulesSdkToTerraform(ctx, diags, d.Rules)

	dataMapValue := map[string]attr.Value{
		"created_time":  createdTime,
		"delivery":      delivery,
		"id":            id,
		"modified_time": modifiedTime,
		"name":          name,
		"org_id":        orgId,
		"rules":         rules,
	}
	data, e := NewOrgAlarmtemplatesValue(OrgAlarmtemplatesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) basetypes.ObjectValue {
	var additionalEmails = mistutils.ListOfStringSdkToTerraformEmpty()
	var enabled types.Bool
	var toOrgAdmins types.Bool
	var toSiteAdmins types.Bool

	if data != nil {
		if data.AdditionalEmails != nil && len(data.AdditionalEmails) > 0 {
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
	r, e := basetypes.NewObjectValue(DeliveryValue{}.AttributeTypes(ctx), dataMapValue)
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
