package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nac_portal_template"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacPortalTemplateModel(t *testing.T) {
	type testStep struct {
		config OrgNacPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	// Helper values for pointers
	alignmentCenter := "center"
	colorBlue := "#1074bc"
	poweredByFalse := false

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacPortalTemplateModel{
						OrgId:     GetTestOrgId(),
						Alignment: &alignmentCenter,
						Color:     &colorBlue,
						PoweredBy: &poweredByFalse,
					},
				},
			},
		},
	}

	// Load fixture data following the checklist pattern
	b, err := os.ReadFile("fixtures/org_nac_portal_template_resource/org_nac_portal_template_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")
	testImagePath := CreateTestPNGFile(t)

	for i, fixture := range fixtures {
		fixtureOrgNacPortalTemplateModel := OrgNacPortalTemplateModel{}
		err = hcl.Decode(&fixtureOrgNacPortalTemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgNacPortalTemplateModel.OrgId = GetTestOrgId()

		// Create a test image file for the logo field
		if fixtureOrgNacPortalTemplateModel.Logo != nil {
			fixtureOrgNacPortalTemplateModel.Logo = &testImagePath
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgNacPortalTemplateModel,
				},
			},
		}
	}

	resourceType := "org_nac_portal_template"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nac_portal_template.OrgNacPortalTemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			// Create single-step tests with combined config (NAC portal + NAC portal template)
			// Since portal templates require a NAC portal, we create both in the same config
			// but focus our checks on the portal template
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: NAC portal
				combinedConfig, nacPortalRef := GetOrgNacPortalBaseConfig(step.config.OrgId)

				// Generate the HCL configuration for the portal template
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				// Add the nacportal_id attribute to the body before rendering
				f.Body().SetAttributeRaw("nacportal_id", hclwrite.TokensForIdentifier(nacPortalRef))
				combinedConfig = combinedConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

				// Focus checks on the portal template resource (NAC portal is a prerequisite)
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

func (s *OrgNacPortalTemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Check fields in struct order
	// 1. OrgId (required)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	// 2. NacportalId (required but dynamic reference) - Use AttrSet since it's a dynamic reference
	checks.append(t, "TestCheckResourceAttrSet", "nacportal_id")

	// Check optional fields
	if s.Alignment != nil {
		checks.append(t, "TestCheckResourceAttr", "alignment", *s.Alignment)
	}
	if s.Color != nil {
		checks.append(t, "TestCheckResourceAttr", "color", *s.Color)
	}
	if s.Logo != nil {
		// Logo is a file path, but after processing it becomes base64, so just check it's set
		checks.append(t, "TestCheckResourceAttrSet", "logo")
	}
	if s.PoweredBy != nil {
		checks.append(t, "TestCheckResourceAttr", "powered_by", fmt.Sprintf("%t", *s.PoweredBy))
	}

	return checks
}
