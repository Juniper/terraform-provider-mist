package provider

type OrgSsoRoleModel struct {
	Name       string            `hcl:"name"`
	OrgId      string            `hcl:"org_id"`
	Privileges []PrivilegesValue `hcl:"privileges"`
}
