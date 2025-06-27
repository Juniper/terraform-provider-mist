// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWxruleModel(t *testing.T) {
	type testStep struct {
		config OrgWxruleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWxruleModel{
						OrgId:  GetTestOrgId(),
						Order:  1,
						Action: "allow",
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wxrule"
			templateName := "test_template"

			// Create single-step tests with combined config (template + WX Rule)
			// Since WX Rules require a template, we create both in the same config
			// but focus our checks on the WX Rule resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: template + WX Rule
				combinedConfig := generateOrgWxruleConfig(templateName, tName, step.config)

				// Focus checks on the WX Rule resource (template is just a prerequisite)
				checks := step.config.testChecks(t, resourceType, tName)

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

// generateOrgWxruleConfig creates a combined configuration with both a WLAN template and a WX Rule
// This handles the prerequisite that WX Rules require a template to exist
func generateOrgWxruleConfig(templateName, wxRuleName string, wxRuleConfig OrgWxruleModel) string {
	// Create the prerequisite WLAN template
	templateConfig := OrgWlantemplateModel{
		Name:  "Test_WLAN_Template",
		OrgId: wxRuleConfig.OrgId,
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&templateConfig, f.Body())
	templateConfigStr := Render("org_wlantemplate", templateName, string(f.Bytes()))

	// Create the WX Rule that references the template
	templateRef := fmt.Sprintf("mist_org_wlantemplate.%s.id", templateName)

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wxRuleConfig, f.Body())

	// Add the template_id attribute to the body before rendering
	f.Body().SetAttributeRaw("template_id", hclwrite.TokensForIdentifier(templateRef))
	wxRuleConfigStr := Render("org_wxrule", wxRuleName, string(f.Bytes()))

	// Combine both configs
	return templateConfigStr + "\n\n" + wxRuleConfigStr
}

func (s *OrgWxruleModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttr", "order", fmt.Sprintf("%d", s.Order))
	checks.append(t, "TestCheckResourceAttr", "action", s.Action)
	checks.append(t, "TestCheckResourceAttrSet", "template_id")

	return checks
}
