// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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

				// Focus checks on the webhook resource
				checks := step.config.testChecks(t, resourceType, tName)

				if len(config.Topics) > 0 {
					checks.append(t, "TestCheckResourceAttr", "topics.0", config.Topics[0])
				}

				steps[i] = resource.TestStep{
					Config: configStr,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration for debugging
				t.Logf("\n// ------ Config ------\n%s\n", configStr)
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})

		})
	}
}

func (s *OrgWebhookModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttr", "url", s.Url)
	checks.append(t, "TestCheckResourceAttr", "topics.#", fmt.Sprintf("%d", len(s.Topics)))

	return checks
}
