package resource_device_gateway

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.ServicePolicyAppqoe {
	data := models.ServicePolicyAppqoe{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewAppqoeValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
		return &data
	}
}

func servicePolicyEwfRuleTerraformToSdk(d basetypes.ListValue) []models.ServicePolicyEwfRule {
	var dataList []models.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(EwfValue)
		data := models.ServicePolicyEwfRule{}
		if plan.AlertOnly.ValueBoolPointer() != nil {
			data.AlertOnly = models.ToPointer(plan.AlertOnly.ValueBool())
		}
		if plan.BlockMessage.ValueStringPointer() != nil {
			data.BlockMessage = models.ToPointer(plan.BlockMessage.ValueString())
		}
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
		if plan.Profile.ValueStringPointer() != nil {
			data.Profile = models.ToPointer(models.ServicePolicyEwfRuleProfileEnum(plan.Profile.ValueString()))
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func idpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpConfig {
	data := models.IdpConfig{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewIdpValueMust(d.AttributeTypes(ctx), d.Attributes())

		if plan.IdpprofileId.ValueStringPointer() != nil {
			idpProfileId, e := uuid.Parse(plan.IdpprofileId.ValueString())
			if e == nil {
				data.IdpprofileId = models.ToPointer(idpProfileId)
			} else {
				diags.AddError("Bad value for idpprofile_id", e.Error())
			}
		}

		if plan.AlertOnly.ValueBoolPointer() != nil {
			data.AlertOnly = models.ToPointer(plan.AlertOnly.ValueBool())
		}
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
		if plan.Profile.ValueStringPointer() != nil {
			data.Profile = models.ToPointer(plan.Profile.ValueString())
		}
		return &data
	}
}

func servicePoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicy {
	var dataList []models.ServicePolicy
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ServicePoliciesValue)
		data := models.ServicePolicy{}

		if plan.Action.ValueStringPointer() != nil {
			data.Action = models.ToPointer(models.AllowDenyEnum(plan.Action.ValueString()))
		}

		if !plan.Appqoe.IsNull() && !plan.Appqoe.IsUnknown() {
			data.Appqoe = servicePolicyAppqoeTerraformToSdk(ctx, plan.Appqoe)
		}

		if !plan.Ewf.IsNull() && !plan.Ewf.IsUnknown() {
			data.Ewf = servicePolicyEwfRuleTerraformToSdk(plan.Ewf)
		}

		if !plan.Idp.IsNull() && !plan.Idp.IsUnknown() {
			data.Idp = idpConfigTerraformToSdk(ctx, diags, plan.Idp)
		}

		if plan.LocalRouting.ValueBoolPointer() != nil {
			data.LocalRouting = models.ToPointer(plan.LocalRouting.ValueBool())
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}
		if plan.PathPreference.ValueStringPointer() != nil {
			data.PathPreference = models.ToPointer(plan.PathPreference.ValueString())
		}
		if plan.ServicepolicyId.ValueStringPointer() != nil {
			servicePolicyId, e := uuid.Parse(plan.ServicepolicyId.ValueString())
			if e == nil {
				data.ServicepolicyId = models.ToPointer(servicePolicyId)
			} else {
				diags.AddError("Bad value for servicepolicy_id", e.Error())
			}
		}

		if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
			data.Services = misttransform.ListOfStringTerraformToSdk(plan.Services)
		}
		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = misttransform.ListOfStringTerraformToSdk(plan.Tenants)
		}

		dataList = append(dataList, data)
	}
	return dataList
}
