package provider

type OrgWebhookModel struct {
	Enabled            *bool             `hcl:"enabled"`
	Headers            map[string]string `hcl:"headers"`
	Name               string            `hcl:"name"`
	Oauth2ClientId     *string           `hcl:"oauth2_client_id"`
	Oauth2ClientSecret *string           `hcl:"oauth2_client_secret"`
	Oauth2GrantType    *string           `hcl:"oauth2_grant_type"`
	Oauth2Password     *string           `hcl:"oauth2_password"`
	Oauth2Scopes       []string          `hcl:"oauth2_scopes"`
	Oauth2TokenUrl     *string           `hcl:"oauth2_token_url"`
	Oauth2Username     *string           `hcl:"oauth2_username"`
	OrgId              string            `hcl:"org_id"`
	Secret             *string           `hcl:"secret"`
	SplunkToken        *string           `hcl:"splunk_token"`
	Topics             []string          `hcl:"topics"`
	Type               *string           `hcl:"type"`
	Url                string            `hcl:"url"`
	VerifyCert         *bool             `hcl:"verify_cert"`
}
