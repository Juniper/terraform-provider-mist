// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanModel(t *testing.T) {
	type testStep struct {
		config OrgWlanModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanModel{
						OrgId: GetTestOrgId(),
						Ssid:  "TestSSID",
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wlan"
			templateName := "test_template"

			// Create single-step tests with combined config (template + WLAN)
			// Since WLANs require a template, we create both in the same config
			// but focus our checks on the WLAN resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: template + WLAN
				combinedConfig := generateOrgWlanTestConfig(templateName, tName, step.config)

				// Focus checks on the WLAN resource (template is just a prerequisite)
				checks := step.config.testChecks(t, resourceType, tName)

				// Custom checks for WLAN-specific attributes

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration for debugging
				t.Logf("\n// ------ Combined Config ------\n%s\n", combinedConfig)
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})

		})
	}
}

// generateOrgWlanTestConfig creates a combined configuration with both a WLAN template and a WLAN
// This handles the prerequisite that WLANs require a template to exist
func generateOrgWlanTestConfig(templateName, wlanName string, wlanConfig OrgWlanModel) string {
	// Create the prerequisite WLAN template
	templateConfig := OrgWlantemplateModel{
		Name:  "Test_WLAN_Template",
		OrgId: wlanConfig.OrgId,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&templateConfig, f.Body())
	templateConfigStr := Render("org_wlantemplate", templateName, string(f.Bytes()))

	// Create the WLAN that references the template
	templateRef := fmt.Sprintf("mist_org_wlantemplate.%s.id", templateName)

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wlanConfig, f.Body())

	// Add the template_id attribute to the body before rendering
	f.Body().SetAttributeRaw("template_id", hclwrite.TokensForIdentifier(templateRef))
	wlanConfigStr := Render("org_wlan", wlanName, string(f.Bytes()))

	// Combine both configs
	return templateConfigStr + "\n\n" + wlanConfigStr
}

func (s *OrgWlanModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttr", "ssid", s.Ssid)
	checks.append(t, "TestCheckResourceAttrSet", "template_id")

	return checks
}
