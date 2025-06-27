package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgSsoRole(t *testing.T) {
	type testStep struct {
		config OrgSsoRoleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgSsoRoleModel{
						OrgId: GetTestOrgId(),
						Name:  "test-sso-role",
						Privileges: []PrivilegesValue{
							{
								Role:  "read",
								Scope: "org",
							},
						},
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
				configStr := Render("org_sso_role", tName, string(f.Bytes()))

				checks := config.testChecks(t, PrefixProviderName("org_sso_role"), tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

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

func (o *OrgSsoRoleModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check privileges
	if len(o.Privileges) > 0 {
		// Check the first privilege entry
		checks.append(t, "TestCheckResourceAttr", "privileges.0.role", o.Privileges[0].Role)
		checks.append(t, "TestCheckResourceAttr", "privileges.0.scope", o.Privileges[0].Scope)

		// Check optional fields in privileges if they are set
		if o.Privileges[0].SiteId != nil {
			checks.append(t, "TestCheckResourceAttr", "privileges.0.site_id", *o.Privileges[0].SiteId)
		}
		if o.Privileges[0].SitegroupId != nil {
			checks.append(t, "TestCheckResourceAttr", "privileges.0.sitegroup_id", *o.Privileges[0].SitegroupId)
		}
	}

	return checks
}
