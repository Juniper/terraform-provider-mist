package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlan_portal_template"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanPortalTemplateModel(t *testing.T) {
	type testStep struct {
		config OrgWlanPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	// Helper values for boolean pointers
	boolTrue := true

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanPortalTemplateModel{
						OrgId: GetTestOrgId(),
						PortalTemplate: PortalTemplateValue{
							Company: &boolTrue,
							Email:   &boolTrue,
						},
					},
				},
			},
		},
	}

	// Load fixture data following the checklist pattern
	b, err := os.ReadFile("fixtures/org_wlan_portal_template_resource/org_wlan_portal_template_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")
	testImagePath := CreateTestPNGFile(t)

	for i, fixture := range fixtures {
		fixtureOrgWlanPortalTemplateModel := OrgWlanPortalTemplateModel{}
		err = hcl.Decode(&fixtureOrgWlanPortalTemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWlanPortalTemplateModel.OrgId = GetTestOrgId()

		// Create a test image file for the logo field (following TestOrgWlanPortalImageModel pattern)
		if fixtureOrgWlanPortalTemplateModel.PortalTemplate.Logo != nil {
			fixtureOrgWlanPortalTemplateModel.PortalTemplate.Logo = &testImagePath
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWlanPortalTemplateModel,
				},
			},
		}
	}

	resourceType := "org_wlan_portal_template"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_wlan_portal_template.OrgWlanPortalTemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			// Create single-step tests with combined config (WLAN template + WLAN + portal template)
			// Since portal templates require a WLAN, and WLANs require a template,
			// we create all three in the same config but focus our checks on the portal template
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN
				combinedConfig, wlanRef := GetOrgWlanBaseConfig(step.config.OrgId)

				// Generate the HCL configuration for the portal template
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				// Add the wlan_id attribute to the body before rendering
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig = combinedConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

				// Focus checks on the portal template resource (WLAN template and WLAN are prerequisites)
				checks := step.config.testChecks(t, resourceType, tName, tracker)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration and checks for debugging
				t.Logf("\n// ------ begin config for test case %s step %d ------\n%s\n// -------- end config for test case %s step %d ------\n", tName, i+1, combinedConfig, tName, i+1)
				t.Logf("\n// ------ begin checks for test case %s step %d ------\n%s\n// -------- end checks for test case %s step %d ------\n", tName, i+1, checks.string(), tName, i+1)
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

func (s *OrgWlanPortalTemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, s)

	return checks
}
