package resource_org_deviceprofile_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ServicePolicyAppqoe {
	data := models.ServicePolicyAppqoe{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewAppqoeValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
			}
		}
	}
	return &data
}

func servicePolicyEwfRuleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicyEwfRule {
	var data_list []models.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(EwfValue)
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

		data_list = append(data_list, data)
	}
	return data_list
}

func idpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpConfig {
	data := models.IdpConfig{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewIdpValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.IdpprofileId.ValueStringPointer() != nil {
				idp_profile_id, e := uuid.Parse(plan.IdpprofileId.ValueString())
				if e != nil {
					diags.AddError("Unable to convert IdpprofileId", e.Error())
				} else {
					data.IdpprofileId = models.ToPointer(idp_profile_id)
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
		}
	}
	return &data
}

func servicePoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicy {
	var data_list []models.ServicePolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServicePoliciesValue)
		data := models.ServicePolicy{}

		if plan.Action.ValueStringPointer() != nil {
			data.Action = models.ToPointer(models.AllowDenyEnum(plan.Action.ValueString()))
		}

		if !plan.Appqoe.IsNull() && !plan.Appqoe.IsUnknown() {
			data.Appqoe = servicePolicyAppqoeTerraformToSdk(ctx, diags, plan.Appqoe)
		}

		if !plan.Ewf.IsNull() && !plan.Ewf.IsUnknown() {
			data.Ewf = servicePolicyEwfRuleTerraformToSdk(ctx, diags, plan.Ewf)
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
			service_policy_id, e := uuid.Parse(plan.ServicepolicyId.ValueString())
			if e != nil {
				diags.AddError("Unable to convert ServicepolicyId", e.Error())
			} else {
				data.ServicepolicyId = models.ToPointer(service_policy_id)
			}
		}

		if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
			data.Services = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services)
		}
		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Tenants)
		}

		data_list = append(data_list, data)
	}
	return data_list
}
