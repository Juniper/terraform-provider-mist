package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgApitoken(t *testing.T) {
	type testStep struct {
		config OrgApitokenModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgApitokenModel{
						Name:  "apitoken1",
						OrgId: GetTestOrgId(),
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

	b, err := os.ReadFile("fixtures/org_apitoken_resource/org_apitoken_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "$")

	for i, fixture := range fixtures {
		var FixtureOrgApitokenModel OrgApitokenModel
		err = hcl.Decode(&FixtureOrgApitokenModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgApitokenModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_apitoken"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName)
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

func (s *OrgApitokenModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "privileges.#", fmt.Sprintf("%d", len(s.Privileges)))
	for i, priv := range s.Privileges {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.role", i), priv.Role)
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.scope", i), priv.Scope)
		if priv.SiteId != nil {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.site_id", i), *priv.SiteId)
		}
		if priv.SitegroupId != nil {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("privileges.%d.sitegroup_id", i), *priv.SitegroupId)
		}
	}

	return checks
}
