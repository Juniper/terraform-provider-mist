package resource_org_nacidp

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Sso) (OrgNacidpModel, diag.Diagnostics) {
	var state OrgNacidpModel
	var diags diag.Diagnostics
	var id types.String
	var idp_type types.String
	var ldap_base_dn types.String
	var ldap_bind_dn types.String
	var ldap_bind_password types.String
	var ldap_ca_certs types.List = types.ListNull(types.StringType)
	var ldap_client_cert types.String
	var ldap_client_key types.String
	var ldap_group_attr types.String
	var ldap_group_dn types.String
	var ldap_group_filter types.String
	var ldap_member_filter types.String
	var ldap_resolve_groups types.Bool
	var ldap_server_hosts types.List = types.ListNull(types.StringType)
	var ldap_type types.String
	var ldap_user_filter types.String
	var name types.String
	var oauth_cc_client_id types.String
	var oauth_cc_client_secret types.String
	var oauth_discovery_url types.String
	var oauth_ropc_client_id types.String
	var oauth_ropc_client_secret types.String
	var oauth_tenant_id types.String
	var oauth_type types.String
	var org_id types.String

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.IdpType != nil {
		idp_type = types.StringValue(string(*data.IdpType))
	}
	if data.LdapBaseDn != nil {
		ldap_base_dn = types.StringValue(*data.LdapBaseDn)
	}
	if data.LdapBindDn != nil {
		ldap_bind_dn = types.StringValue(*data.LdapBindDn)
	}
	if data.LdapBindPassword != nil {
		ldap_bind_password = types.StringValue(*data.LdapBindPassword)
	}
	if data.LdapCacerts != nil {
		ldap_ca_certs = mist_transform.ListOfStringSdkToTerraform(ctx, data.LdapCacerts)
	}
	if data.LdapClientCert != nil {
		ldap_client_cert = types.StringValue(*data.LdapClientCert)
	}
	if data.LdapClientKey != nil {
		ldap_client_key = types.StringValue(*data.LdapClientKey)
	}
	if data.LdapGroupAttr != nil {
		ldap_group_attr = types.StringValue(*data.LdapGroupAttr)
	}
	if data.LdapGroupDn != nil {
		ldap_group_dn = types.StringValue(*data.LdapGroupDn)
	}
	if data.GroupFilter != nil {
		ldap_group_filter = types.StringValue(*data.GroupFilter)
	}
	if data.MemberFilter != nil {
		ldap_member_filter = types.StringValue(*data.MemberFilter)
	}
	if data.LdapResolveGroups != nil {
		ldap_resolve_groups = types.BoolValue(*data.LdapResolveGroups)
	}
	if data.LdapServerHosts != nil {
		ldap_server_hosts = mist_transform.ListOfStringSdkToTerraform(ctx, data.LdapServerHosts)
	}
	if data.LdapType != nil {
		ldap_type = types.StringValue(string(*data.LdapType))
	}
	if data.LdapUserFilter != nil {
		ldap_user_filter = types.StringValue(*data.LdapUserFilter)
	}

	name = types.StringValue(data.Name)

	if data.OauthCcClientId != nil {
		oauth_cc_client_id = types.StringValue(*data.OauthCcClientId)
	}
	if data.OauthCcClientSecret != nil {
		oauth_cc_client_secret = types.StringValue(*data.OauthCcClientSecret)
	}
	if data.OauthDiscoveryUrl != nil {
		oauth_discovery_url = types.StringValue(*data.OauthDiscoveryUrl)
	}
	if data.OauthRopcClientId != nil {
		oauth_ropc_client_id = types.StringValue(*data.OauthRopcClientId)
	}
	if data.OauthRopcClientSecret != nil {
		oauth_ropc_client_secret = types.StringValue(*data.OauthRopcClientSecret)
	}
	if data.OauthTenantId != nil {
		oauth_tenant_id = types.StringValue(*data.OauthTenantId)
	}
	if data.OauthType != nil {
		/// TEMP WORKAROUND
		if string(*data.OauthType) == "standards" {
			oauth_type = types.StringValue("azure")
		} else {
			oauth_type = types.StringValue(string(*data.OauthType))
		}
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}

	state.Id = id
	state.IdpType = idp_type
	state.LdapBaseDn = ldap_base_dn
	state.LdapBindDn = ldap_bind_dn
	state.LdapBindPassword = ldap_bind_password
	state.LdapCacerts = ldap_ca_certs
	state.LdapClientCert = ldap_client_cert
	state.LdapClientKey = ldap_client_key
	state.LdapGroupAttr = ldap_group_attr
	state.LdapGroupDn = ldap_group_dn
	state.GroupFilter = ldap_group_filter
	state.MemberFilter = ldap_member_filter
	state.LdapResolveGroups = ldap_resolve_groups
	state.LdapServerHosts = ldap_server_hosts
	state.LdapType = ldap_type
	state.LdapUserFilter = ldap_user_filter
	state.Name = name
	state.OauthCcClientId = oauth_cc_client_id
	state.OauthCcClientSecret = oauth_cc_client_secret
	state.OauthDiscoveryUrl = oauth_discovery_url
	state.OauthRopcClientId = oauth_ropc_client_id
	state.OauthRopcClientSecret = oauth_ropc_client_secret
	state.OauthTenantId = oauth_tenant_id
	state.OauthType = oauth_type
	state.OrgId = org_id

	return state, diags
}
