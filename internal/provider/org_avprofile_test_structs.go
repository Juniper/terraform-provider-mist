package provider

type OrgAvprofileModel struct {
	FallbackAction *string  `hcl:"fallback_action"`
	MaxFilesize    *int64   `hcl:"max_filesize"`
	MimeWhitelist  []string `hcl:"mime_whitelist"`
	Name           string   `hcl:"name"`
	OrgId          string   `hcl:"org_id"`
	Protocols      []string `hcl:"protocols"`
	UrlWhitelist   []string `hcl:"url_whitelist"`
}
