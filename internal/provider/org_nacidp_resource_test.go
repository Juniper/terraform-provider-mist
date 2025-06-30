package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacidp(t *testing.T) {
	type testStep struct {
		config OrgNacidpModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacidpModel{
						OrgId:            GetTestOrgId(),
						Name:             "test-nacidp",
						IdpType:          "ldap",
						LdapBaseDn:       stringPtr("dc=example,dc=com"),
						LdapBindDn:       stringPtr("cn=admin,dc=example,dc=com"),
						LdapBindPassword: stringPtr("password123"),
						LdapType:         stringPtr("azure"),
						LdapServerHosts:  []string{"ldap.example.com"},
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render("org_nacidp", tName, string(f.Bytes()))

				checks := config.testChecks(t, PrefixProviderName("org_nacidp"), tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: configStr,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
}

func (o *OrgNacidpModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "idp_type", o.IdpType)

	// Check optional fields that are set in the test
	if o.LdapBaseDn != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_base_dn", *o.LdapBaseDn)
	}
	if o.LdapBindDn != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_bind_dn", *o.LdapBindDn)
	}
	if o.LdapBindPassword != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_bind_password", *o.LdapBindPassword)
	}
	if o.LdapType != nil {
		checks.append(t, "TestCheckResourceAttr", "ldap_type", *o.LdapType)
	}
	if len(o.LdapServerHosts) > 0 {
		checks.append(t, "TestCheckResourceAttr", "ldap_server_hosts.0", o.LdapServerHosts[0])
	}

	return checks
}
