package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeviceGatewayModel(t *testing.T) {
	resourceType := "device_gateway"
	t.Skipf("Skipping %s tests, as they require a real device.", resourceType)

	type testStep struct {
		config DeviceGatewayModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: DeviceGatewayModel{
						DeviceId: "00000000-0000-0000-1000-f4bfa8426080",
						Name:     "test_device_gateway",
						SiteId:   "2c107c8e-2e06-404a-ba61-e25b5757ecea",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/device_gateway_resource/device_gateway_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureDeviceGatewayModel DeviceGatewayModel
		err = hcl.Decode(&FixtureDeviceGatewayModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureDeviceGatewayModel,
				},
			},
		}
	}

	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_device_gateway.DeviceGatewayResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				// f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *DeviceGatewayModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, s)

	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttrSet", "device_id")
	checks.append(t, "TestCheckResourceAttrSet", "image1_url")
	checks.append(t, "TestCheckResourceAttrSet", "image2_url")
	checks.append(t, "TestCheckResourceAttrSet", "image3_url")
	checks.append(t, "TestCheckResourceAttrSet", "mac")
	checks.append(t, "TestCheckResourceAttrSet", "model")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "serial")

	return checks
}
