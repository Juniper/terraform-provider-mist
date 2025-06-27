package provider

type OrgVpnModel struct {
	Name  string                `hcl:"name"`
	OrgId string                `hcl:"org_id"`
	Paths map[string]PathsValue `hcl:"paths"`
}
