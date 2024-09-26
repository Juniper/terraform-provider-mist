package resource_org_alarmtemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgAlarmtemplateModel) (*models.AlarmTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.AlarmTemplate{}

	data.Delivery = deliveryTerraformToSdk(ctx, &diags, plan.Delivery)
	data.Name = plan.Name.ValueStringPointer()
	data.Rules = rulesTerraforToSdk(ctx, &diags, plan.Rules)

	return &data, diags
}

func deliveryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DeliveryValue) models.Delivery {
	var data models.Delivery
	if !d.IsNull() && !d.IsUnknown() {
		if !d.AdditionalEmails.IsNull() && !d.AdditionalEmails.IsUnknown() {
			data.AdditionalEmails = mist_transform.ListOfStringTerraformToSdk(ctx, d.AdditionalEmails)
		}
		data.Enabled = *d.Enabled.ValueBoolPointer()
		if !d.ToOrgAdmins.IsNull() && !d.ToOrgAdmins.IsUnknown() {
			data.ToOrgAdmins = d.ToOrgAdmins.ValueBoolPointer()
		}
		if !d.ToSiteAdmins.IsNull() && !d.ToSiteAdmins.IsUnknown() {
			data.ToSiteAdmins = d.ToSiteAdmins.ValueBoolPointer()
		}
	}
	return data
}

func rulesTerraforToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.AlarmTemplateRule {
	data_map := make(map[string]models.AlarmTemplateRule)
	for k, v := range d.Elements() {
		var v_i interface{} = v
		v_r := v_i.(RulesValue)
		data := models.AlarmTemplateRule{}
		if !v_r.Delivery.IsNull() && !v_r.IsUnknown() {
			delivery, e := NewDeliveryValue(DeliveryValue{}.AttributeTypes(ctx), v_r.Delivery.Attributes())
			if e != nil {
				diags.Append(e...)
			} else {
				data.Delivery = models.ToPointer(deliveryTerraformToSdk(ctx, diags, delivery))
			}
		}
		if !v_r.Enabled.IsNull() && !v_r.Enabled.IsUnknown() {
			data.Enabled = v_r.Enabled.ValueBoolPointer()
		}
		data_map[k] = data
	}
	return data_map
}
