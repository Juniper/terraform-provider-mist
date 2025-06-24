// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func (s *OrgWebhookModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "name")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	return checks
}

func TestOrgWebhookModel(t *testing.T) {
	type testStep struct {
		config OrgWebhookModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWebhookModel{
						Name:  "webhook_test",
						Url:   "http://api.mistsys.com",
						OrgId: GetTestOrgId(),
						Topics: []string{
							"alarms",
						},
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_webhook"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))
				fmt.Printf("Config: %s\n", configStr)

				checks := config.testChecks(t, PrefixProviderName(resourceType), tName)
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
