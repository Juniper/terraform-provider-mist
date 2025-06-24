package provider

import ()

type OrgNacidpModel struct {
	GroupFilter *string `hcl:"group_filter"`
	IdpType string `hcl:"idp_type"`
	LdapBaseDn *string `hcl:"ldap_base_dn"`
	LdapBindDn *string `hcl:"ldap_bind_dn"`
	LdapBindPassword *string `hcl:"ldap_bind_password"`
	LdapCacerts []string `hcl:"ldap_cacerts"`
	LdapClientCert *string `hcl:"ldap_client_cert"`
	LdapClientKey *string `hcl:"ldap_client_key"`
	LdapGroupAttr *string `hcl:"ldap_group_attr"`
	LdapGroupDn *string `hcl:"ldap_group_dn"`
	LdapResolveGroups *bool `hcl:"ldap_resolve_groups"`
	LdapServerHosts []string `hcl:"ldap_server_hosts"`
	LdapType *string `hcl:"ldap_type"`
	LdapUserFilter *string `hcl:"ldap_user_filter"`
	MemberFilter *string `hcl:"member_filter"`
	Name string `hcl:"name"`
	OauthCcClientId *string `hcl:"oauth_cc_client_id"`
	OauthCcClientSecret *string `hcl:"oauth_cc_client_secret"`
	OauthDiscoveryUrl *string `hcl:"oauth_discovery_url"`
	OauthPingIdentityRegion *string `hcl:"oauth_ping_identity_region"`
	OauthRopcClientId *string `hcl:"oauth_ropc_client_id"`
	OauthRopcClientSecret *string `hcl:"oauth_ropc_client_secret"`
	OauthTenantId *string `hcl:"oauth_tenant_id"`
	OauthType *string `hcl:"oauth_type"`
	OrgId string `hcl:"org_id"`
	ScimEnabled *bool `hcl:"scim_enabled"`
	ScimSecretToken *string `hcl:"scim_secret_token"`
}

