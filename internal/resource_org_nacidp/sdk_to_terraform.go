package resource_org_nacidp

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(data *models.Sso) (OrgNacidpModel, diag.Diagnostics) {
	var state OrgNacidpModel
	var diags diag.Diagnostics
	var id types.String
	var idpType types.String
	var ldapBaseDn types.String
	var ldapBindDn types.String
	var ldapBindPassword types.String
	var ldapCaCerts = types.ListNull(types.StringType)
	var ldapClientCert types.String
	var ldapClientKey types.String
	var ldapGroupAttr types.String
	var ldapGroupDn types.String
	var ldapGroupFilter types.String
	var ldapMemberFilter types.String
	var ldapResolveGroups types.Bool
	var ldapServerHosts = types.ListNull(types.StringType)
	var ldapType types.String
	var ldapUserFilter types.String
	var name types.String
	var oauthCcClientId types.String
	var oauthCcClientSecret types.String
	var oauthDiscoveryUrl types.String
	var oauthPingIdentityRegion types.String
	var oauthRopcClientId types.String
	var oauthRopcClientSecret types.String
	var oauthTenantId types.String
	var oauthType types.String
	var orgId types.String
	var scimEnabled = types.BoolValue(false)
	var scimSecretToken = types.StringValue("")

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.IdpType != nil {
		idpType = types.StringValue(string(*data.IdpType))
	}
	if data.LdapBaseDn != nil {
		ldapBaseDn = types.StringValue(*data.LdapBaseDn)
	}
	if data.LdapBindDn != nil {
		ldapBindDn = types.StringValue(*data.LdapBindDn)
	}
	if data.LdapBindPassword != nil {
		ldapBindPassword = types.StringValue(*data.LdapBindPassword)
	}
	if data.LdapCacerts != nil {
		ldapCaCerts = misttransform.ListOfStringSdkToTerraform(data.LdapCacerts)
	}
	if data.LdapClientCert != nil {
		ldapClientCert = types.StringValue(*data.LdapClientCert)
	}
	if data.LdapClientKey != nil {
		ldapClientKey = types.StringValue(*data.LdapClientKey)
	}
	if data.LdapGroupAttr != nil {
		ldapGroupAttr = types.StringValue(*data.LdapGroupAttr)
	}
	if data.LdapGroupDn != nil {
		ldapGroupDn = types.StringValue(*data.LdapGroupDn)
	}
	if data.GroupFilter != nil {
		ldapGroupFilter = types.StringValue(*data.GroupFilter)
	}
	if data.MemberFilter != nil {
		ldapMemberFilter = types.StringValue(*data.MemberFilter)
	}
	if data.LdapResolveGroups != nil {
		ldapResolveGroups = types.BoolValue(*data.LdapResolveGroups)
	}
	if data.LdapServerHosts != nil {
		ldapServerHosts = misttransform.ListOfStringSdkToTerraform(data.LdapServerHosts)
	}
	if data.LdapType != nil {
		ldapType = types.StringValue(string(*data.LdapType))
	}
	if data.LdapUserFilter != nil {
		ldapUserFilter = types.StringValue(*data.LdapUserFilter)
	}

	name = types.StringValue(data.Name)

	if data.OauthCcClientId != nil {
		oauthCcClientId = types.StringValue(*data.OauthCcClientId)
	}
	if data.OauthCcClientSecret != nil {
		oauthCcClientSecret = types.StringValue(*data.OauthCcClientSecret)
	}
	if data.OauthDiscoveryUrl != nil {
		oauthDiscoveryUrl = types.StringValue(*data.OauthDiscoveryUrl)
	}
	if data.OauthPingIdentityRegion != nil {
		oauthPingIdentityRegion = types.StringValue(string(*data.OauthPingIdentityRegion))
	}
	if data.OauthRopcClientId != nil {
		oauthRopcClientId = types.StringValue(*data.OauthRopcClientId)
	}
	if data.OauthRopcClientSecret != nil {
		oauthRopcClientSecret = types.StringValue(*data.OauthRopcClientSecret)
	}
	if data.OauthTenantId != nil {
		oauthTenantId = types.StringValue(*data.OauthTenantId)
	}
	if data.OauthType != nil {
		/// TEMP WORKAROUND
		if string(*data.OauthType) == "standards" {
			oauthType = types.StringValue("azure")
		} else {
			oauthType = types.StringValue(string(*data.OauthType))
		}
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.ScimEnabled != nil {
		scimEnabled = types.BoolValue(*data.ScimEnabled)
	}
	if data.ScimSecretToken != nil {
		scimSecretToken = types.StringValue(*data.ScimSecretToken)
	}

	state.Id = id
	state.IdpType = idpType
	state.LdapBaseDn = ldapBaseDn
	state.LdapBindDn = ldapBindDn
	state.LdapBindPassword = ldapBindPassword
	state.LdapCacerts = ldapCaCerts
	state.LdapClientCert = ldapClientCert
	state.LdapClientKey = ldapClientKey
	state.LdapGroupAttr = ldapGroupAttr
	state.LdapGroupDn = ldapGroupDn
	state.GroupFilter = ldapGroupFilter
	state.MemberFilter = ldapMemberFilter
	state.LdapResolveGroups = ldapResolveGroups
	state.LdapServerHosts = ldapServerHosts
	state.LdapType = ldapType
	state.LdapUserFilter = ldapUserFilter
	state.Name = name
	state.OauthCcClientId = oauthCcClientId
	state.OauthCcClientSecret = oauthCcClientSecret
	state.OauthDiscoveryUrl = oauthDiscoveryUrl
	state.OauthPingIdentityRegion = oauthPingIdentityRegion
	state.OauthRopcClientId = oauthRopcClientId
	state.OauthRopcClientSecret = oauthRopcClientSecret
	state.OauthTenantId = oauthTenantId
	state.OauthType = oauthType
	state.OrgId = orgId
	state.ScimEnabled = scimEnabled
	state.ScimSecretToken = scimSecretToken

	return state, diags
}
