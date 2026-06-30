package provider

type OrgSsoModel struct {
	CustomLogoutUrl          *string  `hcl:"custom_logout_url"`
	DefaultRole              *string  `hcl:"default_role"`
	IdpCert                  string   `hcl:"idp_cert"`
	IdpSignAlgo              string   `hcl:"idp_sign_algo"`
	IdpSsoUrl                string   `hcl:"idp_sso_url"`
	IgnoreUnmatchedRoles     *bool    `hcl:"ignore_unmatched_roles"`
	Issuer                   string   `hcl:"issuer"`
	Name                     string   `hcl:"name"`
	NameidFormat             *string  `hcl:"nameid_format"`
	OauthProviderDomain      *string  `hcl:"oauth_provider_domain"`
	OpenroamingSsids         []string `hcl:"openroaming_ssids"`
	OpenroamingWbaClientCert *string  `hcl:"openroaming_wba_client_cert"`
	OpenroamingWbaClientKey  *string  `hcl:"openroaming_wba_client_key"`
	OrgId                    string   `hcl:"org_id"`
	RoleAttrExtraction       *string  `hcl:"role_attr_extraction"`
	RoleAttrFrom             *string  `hcl:"role_attr_from"`
}
