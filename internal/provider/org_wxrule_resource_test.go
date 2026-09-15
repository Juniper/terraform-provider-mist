package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wxrule"
	"github.com/hashicorp/hcl"
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

	// Load fixture data following the org_wlan pattern
	b, err := os.ReadFile("fixtures/org_wxrule_resource/org_wxrule_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgWxruleModel := OrgWxruleModel{}
		err = hcl.Decode(&fixtureOrgWxruleModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWxruleModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWxruleModel,
				},
			},
		}
	}

	resourceType := "org_wxrule"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_wxrule.OrgWxruleResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			templateName := "test_template"

			// Create single-step tests with combined config (template + WX Rule)
			// Since WX Rules require a template, we create both in the same config
			// but focus our checks on the WX Rule resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: template + WX Rule
				combinedConfig := generateOrgWxruleConfig(templateName, tName, step.config)

				// Focus checks on the WX Rule resource (template is just a prerequisite)
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

func (s *OrgWxruleModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, s)
	checks.append(t, "TestCheckResourceAttrSet", "template_id")

	return checks
}
