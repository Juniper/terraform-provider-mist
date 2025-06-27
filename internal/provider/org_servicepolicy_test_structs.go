package provider

type OrgServicepolicyModel struct {
	Action         *string         `hcl:"action"`
	Antivirus      *AntivirusValue `hcl:"antivirus"`
	Appqoe         *AppqoeValue    `hcl:"appqoe"`
	Ewf            []EwfValue      `hcl:"ewf"`
	Idp            *IdpValue       `hcl:"idp"`
	LocalRouting   *bool           `hcl:"local_routing"`
	Name           string          `hcl:"name"`
	OrgId          string          `hcl:"org_id"`
	PathPreference *string         `hcl:"path_preference"`
	Services       []string        `hcl:"services"`
	SslProxy       *SslProxyValue  `hcl:"ssl_proxy"`
	Tenants        []string        `hcl:"tenants"`
}
