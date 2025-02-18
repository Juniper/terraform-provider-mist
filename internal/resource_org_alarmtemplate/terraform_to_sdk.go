package resource_org_alarmtemplate

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgAlarmtemplateModel) (*models.AlarmTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.AlarmTemplate{}

	data.Delivery = deliveryTerraformToSdk(plan.Delivery)
	data.Name = plan.Name.ValueStringPointer()
	data.Rules = rulesTerraformToSdk(ctx, &diags, plan.Rules)

	return &data, diags
}

func deliveryTerraformToSdk(d DeliveryValue) models.Delivery {
	var data models.Delivery
	if !d.IsNull() && !d.IsUnknown() {
		if !d.AdditionalEmails.IsNull() && !d.AdditionalEmails.IsUnknown() {
			data.AdditionalEmails = misttransform.ListOfStringTerraformToSdk(d.AdditionalEmails)
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

func rulesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.AlarmTemplateRule {
	dataMap := make(map[string]models.AlarmTemplateRule)
	for k, v := range d.Elements() {
		var vI interface{} = v
		vR := vI.(RulesValue)
		data := models.AlarmTemplateRule{}
		if !vR.Delivery.IsNull() && !vR.IsUnknown() {
			delivery, e := NewDeliveryValue(DeliveryValue{}.AttributeTypes(ctx), vR.Delivery.Attributes())
			if e != nil {
				diags.Append(e...)
			} else {
				data.Delivery = models.ToPointer(deliveryTerraformToSdk(delivery))
			}
		}
		if !vR.Enabled.IsNull() && !vR.Enabled.IsUnknown() {
			data.Enabled = vR.Enabled.ValueBoolPointer()
		}
		dataMap[k] = data
	}
	return dataMap
}
