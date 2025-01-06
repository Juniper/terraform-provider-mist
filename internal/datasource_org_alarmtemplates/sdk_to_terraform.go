package datasource_org_alarmtemplates

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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

	var created_time basetypes.Float64Value
	var delivery basetypes.ObjectValue = types.ObjectNull(DeliveryValue{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var modified_time basetypes.Float64Value
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var rules basetypes.MapValue = basetypes.NewMapNull(RulesValue{}.Type(ctx))

	if d.CreatedTime != nil {
		created_time = types.Float64Value(float64(*d.CreatedTime))
	}
	delivery = deliverySdkToTerraform(ctx, diags, &d.Delivery)
	id = types.StringValue(d.Id.String())
	if d.ModifiedTime != nil {
		modified_time = types.Float64Value(float64(*d.ModifiedTime))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	org_id = types.StringValue(d.OrgId.String())
	rules = rulesSdkToTerraform(ctx, diags, d.Rules)

	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"delivery":      delivery,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
		"rules":         rules,
	}
	data, e := NewOrgAlarmtemplatesValue(OrgAlarmtemplatesValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) basetypes.ObjectValue {
	var additional_emails types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var enabled types.Bool
	var to_org_admins types.Bool
	var to_site_admins types.Bool

	if data != nil {
		if data.AdditionalEmails != nil && len(data.AdditionalEmails) > 0 {
			additional_emails = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalEmails)
		}
		enabled = types.BoolValue(data.Enabled)
		if data.ToOrgAdmins != nil {
			to_org_admins = types.BoolValue(*data.ToOrgAdmins)
		}
		if data.ToSiteAdmins != nil {
			to_site_admins = types.BoolValue(*data.ToSiteAdmins)
		}
	}

	data_map_attr_type := DeliveryValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"additional_emails": additional_emails,
		"enabled":           enabled,
		"to_org_admins":     to_org_admins,
		"to_site_admins":    to_site_admins,
	}
	r, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)
	return r
}

func rulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]models.AlarmTemplateRule) basetypes.MapValue {
	rules_map := make(map[string]attr.Value)
	for k, v := range data {
		var delivery types.Object = basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
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

		data_map_attr_type := RulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"delivery": delivery,
			"enabled":  enabled,
		}
		data, e := NewRulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		rules_map[k] = data
	}
	r, e := basetypes.NewMapValueFrom(ctx, RulesValue{}.Type(ctx), rules_map)
	if e != nil {
		diags.Append(e...)
	}
	return r
}
