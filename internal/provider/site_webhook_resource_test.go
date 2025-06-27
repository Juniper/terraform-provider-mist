package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	// gwc "github.com/terraform-provider-mist/internal/resource_device_gateway_cluster"
)

func TestSiteWebhookModel(t *testing.T) {
	type testStep struct {
		config SiteWebhookModel
	}

	type testCase struct {
		//apiVersionConstraints version.Constraints
		steps []testStep
	}

	var FixtureSiteWebhookModel SiteWebhookModel

	b, err := os.ReadFile("fixtures/site_webhook_resource/site_webhook_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fmt.Println(str)

	err = hcl.Decode(&FixtureSiteWebhookModel, str)
	if err != nil {
		fmt.Printf("error decoding hcl: %s\n", err)
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWebhookModel{
						Name:   "test-site-webhook",
						SiteId: "2c107c8e-2e06-404a-ba61-e25b5757ecea",
						Topics: []string{
							"device-events",
							"alarms",
						},
						Url: "https://example.com/webhook",
					},
				},
			},
		},
		// "hcl_decode": {
		// 	steps: []testStep{
		// 		{
		// 			config: FixtureSiteWebhookModel,
		// 		},
		// 	},
		// },
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_webhook"

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

func (s *SiteWebhookModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "site_id", s.SiteId)
	checks.append(t, "TestCheckResourceAttr", "topics.#", fmt.Sprintf("%d", len(s.Topics)))
	checks.append(t, "TestCheckResourceAttr", "topics.0", s.Topics[0])
	checks.append(t, "TestCheckResourceAttr", "topics.1", s.Topics[1])
	checks.append(t, "TestCheckResourceAttr", "url", s.Url)

	return checks
}
