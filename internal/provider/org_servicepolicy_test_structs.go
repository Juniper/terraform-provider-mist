package provider

type OrgServicepolicyModel struct {
	Aamw           *OrgServicepolicyAamwValue      `hcl:"aamw"`
	Action         *string                         `hcl:"action"`
	Antivirus      *OrgServicepolicyAntivirusValue `hcl:"antivirus"`
	Appqoe         *OrgServicepolicyAppqoeValue    `hcl:"appqoe"`
	Ewf            []OrgServicepolicyEwfValue      `hcl:"ewf"`
	Idp            *OrgServicepolicyIdpValue       `hcl:"idp"`
	LocalRouting   *bool                           `hcl:"local_routing"`
	Name           string                          `hcl:"name"`
	OrgId          string                          `hcl:"org_id"`
	PathPreference *string                         `hcl:"path_preference"`
	Services       []string                        `hcl:"services"`
	SslProxy       *OrgServicepolicySslProxyValue  `hcl:"ssl_proxy"`
	Tenants        []string                        `hcl:"tenants"`
}

type OrgServicepolicyAamwValue struct {
	AamwprofileId *string `cty:"aamwprofile_id" hcl:"aamwprofile_id"`
	Enabled       *bool   `cty:"enabled" hcl:"enabled"`
	Profile       *string `cty:"profile" hcl:"profile"`
}

type OrgServicepolicyAntivirusValue struct {
	AvprofileId *string `cty:"avprofile_id" hcl:"avprofile_id"`
	Enabled     *bool   `cty:"enabled" hcl:"enabled"`
	Profile     *string `cty:"profile" hcl:"profile"`
}

type OrgServicepolicyAppqoeValue struct {
	Enabled *bool `cty:"enabled" hcl:"enabled"`
}

type OrgServicepolicyEwfValue struct {
	AlertOnly    *bool   `cty:"alert_only" hcl:"alert_only"`
	BlockMessage *string `cty:"block_message" hcl:"block_message"`
	Enabled      *bool   `cty:"enabled" hcl:"enabled"`
	Profile      *string `cty:"profile" hcl:"profile"`
}

type OrgServicepolicyIdpValue struct {
	AlertOnly    *bool   `cty:"alert_only" hcl:"alert_only"`
	Enabled      *bool   `cty:"enabled" hcl:"enabled"`
	IdpprofileId *string `cty:"idpprofile_id" hcl:"idpprofile_id"`
	Profile      *string `cty:"profile" hcl:"profile"`
}

type OrgServicepolicySslProxyValue struct {
	CiphersCategory *string `cty:"ciphers_category" hcl:"ciphers_category"`
	Enabled         *bool   `cty:"enabled" hcl:"enabled"`
}
