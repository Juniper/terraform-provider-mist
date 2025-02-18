package resource_org_servicepolicy

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgServicepolicyModel) (models.OrgServicePolicy, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.OrgServicePolicy{}
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueStringPointer()

	if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
		data.Action = (*models.AllowDenyEnum)(plan.Action.ValueStringPointer())
	} else {
		unset["-action"] = ""
	}

	if !plan.Appqoe.IsNull() && !plan.Appqoe.IsUnknown() {
		data.Appqoe = appqoeTerraformToSdk(plan.Appqoe)
	} else {
		unset["-appqoe"] = ""
	}

	if !plan.Ewf.IsNull() && !plan.Ewf.IsUnknown() {
		data.Ewf = ewfRuleTerraformToSdk(plan.Ewf)
	} else {
		unset["-ewf"] = ""
	}

	if !plan.Idp.IsNull() && !plan.Idp.IsUnknown() {
		data.Idp = idpTerraformToSdk(&diags, plan.Idp)
	} else {
		unset["-idp"] = ""
	}

	if plan.LocalRouting.ValueBoolPointer() != nil {
		data.LocalRouting = plan.LocalRouting.ValueBoolPointer()
	} else {
		unset["-local_routing"] = ""
	}

	if plan.Name.ValueStringPointer() != nil {
		data.Name = plan.Name.ValueStringPointer()
	}

	if plan.PathPreference.ValueStringPointer() != nil {
		data.PathPreference = plan.PathPreference.ValueStringPointer()
	} else {
		unset["-path_preferences"] = ""
	}

	if !plan.Services.IsNull() && !plan.Services.IsUnknown() {
		data.Services = misttransform.ListOfStringTerraformToSdk(plan.Services)
	} else {
		unset["-services"] = ""
	}

	if !plan.Tenants.IsNull() && !plan.Services.IsUnknown() {
		data.Tenants = misttransform.ListOfStringTerraformToSdk(plan.Tenants)
	} else {
		unset["-tenants"] = ""
	}

	data.AdditionalProperties = unset

	return data, diags
}
