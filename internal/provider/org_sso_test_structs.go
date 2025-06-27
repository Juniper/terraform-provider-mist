package provider

type OrgSsoModel struct {
	CustomLogoutUrl      *string `hcl:"custom_logout_url"`
	DefaultRole          *string `hcl:"default_role"`
	IdpCert              string  `hcl:"idp_cert"`
	IdpSignAlgo          string  `hcl:"idp_sign_algo"`
	IdpSsoUrl            string  `hcl:"idp_sso_url"`
	IgnoreUnmatchedRoles *bool   `hcl:"ignore_unmatched_roles"`
	Issuer               string  `hcl:"issuer"`
	Name                 string  `hcl:"name"`
	NameidFormat         *string `hcl:"nameid_format"`
	OrgId                string  `hcl:"org_id"`
	RoleAttrExtraction   *string `hcl:"role_attr_extraction"`
	RoleAttrFrom         *string `hcl:"role_attr_from"`
}
