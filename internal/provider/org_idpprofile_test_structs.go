package provider

type OrgIdpprofileModel struct {
	BaseProfile string                         `hcl:"base_profile"`
	Name        string                         `hcl:"name"`
	OrgId       string                         `hcl:"org_id"`
	Overwrites  []OrgIdpprofileOverwritesValue `hcl:"overwrites"`
}

type OrgIdpprofileOverwritesValue struct {
	Action   *string                     `cty:"action" hcl:"action"`
	Matching *OrgIdpprofileMatchingValue `cty:"matching" hcl:"matching"`
	Name     string                      `cty:"name" hcl:"name"`
}

type OrgIdpprofileMatchingValue struct {
	AttackName []string `cty:"attack_name" hcl:"attack_name"`
	DstSubnet  []string `cty:"dst_subnet" hcl:"dst_subnet"`
	Severity   []string `cty:"severity" hcl:"severity"`
}
