package resource_org_sso

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgSsoModel) (*models.Sso, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.Sso{}

	if !plan.CustomLogoutUrl.IsNull() && !plan.CustomLogoutUrl.IsUnknown() {
		data.CustomLogoutUrl = plan.CustomLogoutUrl.ValueStringPointer()
	} else {
		unset["-custom_logout_url"] = ""
	}

	if !plan.DefaultRole.IsNull() && !plan.DefaultRole.IsUnknown() {
		data.DefaultRole = plan.DefaultRole.ValueStringPointer()
	} else {
		unset["-default_role"] = ""
	}

	if !plan.IdpCert.IsNull() && !plan.IdpCert.IsUnknown() {
		data.IdpCert = plan.IdpCert.ValueStringPointer()
	} else {
		unset["-idp_cert"] = ""
	}

	if !plan.IdpSignAlgo.IsNull() && !plan.IdpSignAlgo.IsUnknown() {
		data.IdpSignAlgo = (*models.SsoIdpSignAlgoEnum)(plan.IdpSignAlgo.ValueStringPointer())
	} else {
		unset["-idp_sign_algo"] = ""
	}

	if !plan.IdpSsoUrl.IsNull() && !plan.IdpSsoUrl.IsUnknown() {
		data.IdpSsoUrl = plan.IdpSsoUrl.ValueStringPointer()
	} else {
		unset["-idp_sso_url"] = ""
	}

	if !plan.IgnoreUnmatchedRoles.IsNull() && !plan.IgnoreUnmatchedRoles.IsUnknown() {
		data.IgnoreUnmatchedRoles = plan.IgnoreUnmatchedRoles.ValueBoolPointer()
	} else {
		unset["-ignore_unmatched_roles"] = ""
	}

	if !plan.Issuer.IsNull() && !plan.Issuer.IsUnknown() {
		data.Issuer = plan.Issuer.ValueStringPointer()
	} else {
		unset["-issuer"] = ""
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = plan.Name.ValueString()
	} else {
		unset["-name"] = ""
	}

	if !plan.NameidFormat.IsNull() && !plan.NameidFormat.IsUnknown() {
		data.NameidFormat = (*models.SsoNameidFormatEnum)(plan.NameidFormat.ValueStringPointer())
	} else {
		unset["-nameid_format"] = ""
	}

	if !plan.RoleAttrExtraction.IsNull() && !plan.RoleAttrExtraction.IsUnknown() {
		data.RoleAttrExtraction = plan.RoleAttrExtraction.ValueStringPointer()
	} else {
		unset["-role_attr_extraction"] = ""
	}

	if !plan.RoleAttrFrom.IsNull() && !plan.RoleAttrFrom.IsUnknown() {
		data.RoleAttrFrom = plan.RoleAttrFrom.ValueStringPointer()
	} else {
		unset["-role_attr_from"] = ""
	}

	if !plan.ScimEnabled.IsNull() && !plan.ScimEnabled.IsUnknown() {
		data.ScimEnabled = plan.ScimEnabled.ValueBoolPointer()
	} else {
		unset["-scim_enabled"] = ""
	}

	if !plan.ScimSecretToken.IsNull() && !plan.ScimSecretToken.IsUnknown() {
		data.ScimSecretToken = plan.ScimSecretToken.ValueStringPointer()
	} else {
		unset["-scim_secret_token"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
