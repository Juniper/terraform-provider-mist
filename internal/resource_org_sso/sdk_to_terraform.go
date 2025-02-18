package resource_org_sso

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(data *models.Sso) (OrgSsoModel, diag.Diagnostics) {
	var state OrgSsoModel
	var diags diag.Diagnostics
	var customLogoutUrl types.String
	var defaultRole types.String
	var domain types.String
	var id types.String
	var idpCert types.String
	var idpSignAlgo types.String
	var idpSsoUrl types.String
	var ignoreUnmatchedRoles types.Bool
	var issuer types.String
	var name types.String
	var nameidFormat types.String
	var orgId types.String
	var roleAttrExtraction types.String
	var roleAttrFrom types.String

	if data.CustomLogoutUrl != nil {
		customLogoutUrl = types.StringValue(*data.CustomLogoutUrl)
	}
	if data.DefaultRole != nil {
		defaultRole = types.StringValue(*data.DefaultRole)
	}
	if data.Domain != nil {
		domain = types.StringValue(*data.Domain)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.IdpCert != nil {
		idpCert = types.StringValue(*data.IdpCert)
	}
	if data.IdpSignAlgo != nil {
		idpSignAlgo = types.StringValue(string(*data.IdpSignAlgo))
	}
	if data.IdpSsoUrl != nil {
		idpSsoUrl = types.StringValue(*data.IdpSsoUrl)
	}
	if data.IgnoreUnmatchedRoles != nil {
		ignoreUnmatchedRoles = types.BoolValue(*data.IgnoreUnmatchedRoles)
	}
	if data.Issuer != nil {
		issuer = types.StringValue(*data.Issuer)
	}

	name = types.StringValue(data.Name)

	if data.NameidFormat != nil {
		nameidFormat = types.StringValue(string(*data.NameidFormat))
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.RoleAttrExtraction != nil {
		roleAttrExtraction = types.StringValue(*data.RoleAttrExtraction)
	}
	if data.RoleAttrFrom != nil {
		roleAttrFrom = types.StringValue(*data.RoleAttrFrom)
	}

	state.CustomLogoutUrl = customLogoutUrl
	state.DefaultRole = defaultRole
	state.Domain = domain
	state.Id = id
	state.IdpCert = idpCert
	state.IdpSignAlgo = idpSignAlgo
	state.IdpSsoUrl = idpSsoUrl
	state.IgnoreUnmatchedRoles = ignoreUnmatchedRoles
	state.Issuer = issuer
	state.Name = name
	state.NameidFormat = nameidFormat
	state.OrgId = orgId
	state.RoleAttrExtraction = roleAttrExtraction
	state.RoleAttrFrom = roleAttrFrom

	return state, diags
}
