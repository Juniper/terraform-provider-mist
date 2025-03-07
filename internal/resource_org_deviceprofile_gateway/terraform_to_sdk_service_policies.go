package resource_org_deviceprofile_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ServicePolicyAppqoe {
	data := models.ServicePolicyAppqoe{}
	if !d.IsNull() && !d.IsUnknown() {
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
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewIdpValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
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
		}
	}
	return &data
}

func avTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ServicePolicyAntivirus {
	data := models.ServicePolicyAntivirus{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewAntivirusValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.AvprofileId.ValueStringPointer() != nil {
				avprofileId, e := uuid.Parse(plan.AvprofileId.ValueString())
				if e != nil {
					diags.AddError("Unable to parse IDP Profile ID", e.Error())
				} else {
					data.AvprofileId = models.ToPointer(avprofileId)
				}
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}
			if plan.Profile.ValueStringPointer() != nil {
				data.Profile = plan.Profile.ValueStringPointer()
			}
		}
	}
	return &data
}

func sslProxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ServicePolicySslProxy {
	data := models.ServicePolicySslProxy{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewSslProxyValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.CiphersCategory.ValueStringPointer() != nil {
				data.CiphersCategory = (*models.SslProxyCiphersCategoryEnum)(plan.CiphersCategory.ValueStringPointer())
			}
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}
		}
	}
	return &data
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

		if !plan.Antivirus.IsNull() && !plan.Antivirus.IsUnknown() {
			data.Antivirus = avTerraformToSdk(ctx, diags, plan.Antivirus)
		}

		if !plan.Appqoe.IsNull() && !plan.Appqoe.IsUnknown() {
			data.Appqoe = servicePolicyAppqoeTerraformToSdk(ctx, diags, plan.Appqoe)
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
			data.Services = mistutils.ListOfStringTerraformToSdk(plan.Services)
		}
		if !plan.SslProxy.IsNull() && !plan.SslProxy.IsUnknown() {
			data.SslProxy = sslProxyTerraformToSdk(ctx, diags, plan.SslProxy)
		}
		if !plan.Tenants.IsNull() && !plan.Tenants.IsUnknown() {
			data.Tenants = mistutils.ListOfStringTerraformToSdk(plan.Tenants)
		}

		dataList = append(dataList, data)
	}
	return dataList
}
