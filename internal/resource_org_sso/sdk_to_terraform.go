package resource_org_sso

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Sso) (OrgSsoModel, diag.Diagnostics) {
	var state OrgSsoModel
	var diags diag.Diagnostics
	var custom_logout_url types.String
	var default_role types.String
	var domain types.String
	var id types.String
	var idp_cert types.String
	var idp_sign_algo types.String
	var idp_sso_url types.String
	var ignore_unmatched_roles types.Bool
	var issuer types.String
	var name types.String
	var nameid_format types.String
	var org_id types.String
	var role_attr_extraction types.String
	var role_attr_from types.String

	if data.CustomLogoutUrl != nil {
		custom_logout_url = types.StringValue(*data.CustomLogoutUrl)
	}
	if data.DefaultRole != nil {
		default_role = types.StringValue(*data.DefaultRole)
	}
	if data.Domain != nil {
		domain = types.StringValue(*data.Domain)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.IdpCert != nil {
		idp_cert = types.StringValue(*data.IdpCert)
	}
	if data.IdpSignAlgo != nil {
		idp_sign_algo = types.StringValue(string(*data.IdpSignAlgo))
	}
	if data.IdpSsoUrl != nil {
		idp_sso_url = types.StringValue(*data.IdpSsoUrl)
	}
	if data.IgnoreUnmatchedRoles != nil {
		ignore_unmatched_roles = types.BoolValue(*data.IgnoreUnmatchedRoles)
	}
	if data.Issuer != nil {
		issuer = types.StringValue(*data.Issuer)
	}

	name = types.StringValue(data.Name)

	if data.NameidFormat != nil {
		nameid_format = types.StringValue(string(*data.NameidFormat))
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.RoleAttrExtraction != nil {
		role_attr_extraction = types.StringValue(*data.RoleAttrExtraction)
	}
	if data.RoleAttrFrom != nil {
		role_attr_from = types.StringValue(*data.RoleAttrFrom)
	}

	state.CustomLogoutUrl = custom_logout_url
	state.DefaultRole = default_role
	state.Domain = domain
	state.Id = id
	state.IdpCert = idp_cert
	state.IdpSignAlgo = idp_sign_algo
	state.IdpSsoUrl = idp_sso_url
	state.IgnoreUnmatchedRoles = ignore_unmatched_roles
	state.Issuer = issuer
	state.Name = name
	state.NameidFormat = nameid_format
	state.OrgId = org_id
	state.RoleAttrExtraction = role_attr_extraction
	state.RoleAttrFrom = role_attr_from

	return state, diags
}
