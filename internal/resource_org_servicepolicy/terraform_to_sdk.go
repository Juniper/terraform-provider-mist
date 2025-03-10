package resource_org_servicepolicy

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgServicepolicyModel) (models.OrgServicePolicy, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.OrgServicePolicy{}
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueStringPointer()

	if !plan.Aamw.IsNull() && !plan.Aamw.IsUnknown() {
		data.Aamw = aamwTerraformToSdk(&diags, plan.Aamw)
	} else {
		unset["-aamw"] = ""
	}

	if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
		data.Action = (*models.AllowDenyEnum)(plan.Action.ValueStringPointer())
	} else {
		unset["-action"] = ""
	}

	if !plan.Antivirus.IsNull() && !plan.Antivirus.IsUnknown() {
		data.Antivirus = avTerraformToSdk(&diags, plan.Antivirus)
	} else {
		unset["-antivirus"] = ""
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
		data.Services = mistutils.ListOfStringTerraformToSdk(plan.Services)
	} else {
		unset["-services"] = ""
	}

	if !plan.SslProxy.IsNull() && !plan.SslProxy.IsUnknown() {
		data.SslProxy = sslProxyTerraformToSdk(&diags, plan.SslProxy)
	} else {
		unset["-ssl_proxy"] = ""
	}

	if !plan.Tenants.IsNull() && !plan.Services.IsUnknown() {
		data.Tenants = mistutils.ListOfStringTerraformToSdk(plan.Tenants)
	} else {
		unset["-tenants"] = ""
	}

	data.AdditionalProperties = unset

	return data, diags
}
