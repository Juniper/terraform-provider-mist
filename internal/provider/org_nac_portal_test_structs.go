package provider

type OrgNacPortalModel struct {
	AccessType              *string                  `hcl:"access_type"`
	AdditionalCacerts       []string                 `hcl:"additional_cacerts"`
	AdditionalNacServerName []string                 `hcl:"additional_nac_server_name"`
	CertExpireTime          *int64                   `hcl:"cert_expire_time"`
	EapType                 *string                  `hcl:"eap_type"`
	EnableTelemetry         *bool                    `hcl:"enable_telemetry"`
	ExpiryNotificationTime  *int64                   `hcl:"expiry_notification_time"`
	Name                    string                   `hcl:"name"`
	NotifyExpiry            *bool                    `hcl:"notify_expiry"`
	OrgId                   string                   `hcl:"org_id"`
	Portal                  *OrgNacPortalPortalValue `hcl:"portal"`
	Ssid                    *string                  `hcl:"ssid"`
	Sso                     *OrgNacPortalSsoValue    `hcl:"sso"`
	Tos                     *string                  `hcl:"tos"`
	Type                    *string                  `hcl:"type"`
}

type OrgNacPortalPortalValue struct {
	Auth              *string `cty:"auth" hcl:"auth"`
	Expire            *int64  `cty:"expire" hcl:"expire"`
	ExternalPortalUrl *string `cty:"external_portal_url" hcl:"external_portal_url"`
	ForceReconnect    *bool   `cty:"force_reconnect" hcl:"force_reconnect"`
	Forward           *bool   `cty:"forward" hcl:"forward"`
	ForwardUrl        *string `cty:"forward_url" hcl:"forward_url"`
	MaxNumDevices     *int64  `cty:"max_num_devices" hcl:"max_num_devices"`
	Privacy           *bool   `cty:"privacy" hcl:"privacy"`
}

type OrgNacPortalSsoValue struct {
	IdpCert           *string                            `cty:"idp_cert" hcl:"idp_cert"`
	IdpSignAlgo       *string                            `cty:"idp_sign_algo" hcl:"idp_sign_algo"`
	IdpSsoUrl         *string                            `cty:"idp_sso_url" hcl:"idp_sso_url"`
	Issuer            *string                            `cty:"issuer" hcl:"issuer"`
	NameidFormat      *string                            `cty:"nameid_format" hcl:"nameid_format"`
	SsoRoleMatching   []OrgNacPortalSsoRoleMatchingValue `cty:"sso_role_matching" hcl:"sso_role_matching"`
	UseSsoRoleForCert *bool                              `cty:"use_sso_role_for_cert" hcl:"use_sso_role_for_cert"`
}

type OrgNacPortalSsoRoleMatchingValue struct {
	Assigned *string `cty:"assigned" hcl:"assigned"`
	Match    *string `cty:"match" hcl:"match"`
}
