// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nac_portal_image"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacPortalImageModel(t *testing.T) {
	testOrgID := GetTestOrgId()
	// Create a test PNG file
	testImagePath := CreateTestPNGFile(t)

	type testStep struct {
		config OrgNacPortalImageModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacPortalImageModel{
						OrgId: testOrgID,
						File:  testImagePath,
					},
				},
			},
		},
	}

	resourceType := "org_nac_portal_image"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nac_portal_image.OrgNacPortalImageResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			testSteps := make([]resource.TestStep, 0)
			for i, step := range tCase.steps {
				// Generate combined config: NAC portal + portal image
				combinedConfig, nacPortalRef := GetOrgNacPortalBaseConfig(testOrgID)

				// Generate the HCL configuration for the portal image
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				f.Body().SetAttributeRaw("nacportal_id", hclwrite.TokensForIdentifier(nacPortalRef))
				combinedConfig = combinedConfig + "\n\n" + Render("org_nac_portal_image", tName, string(f.Bytes()))

				// Create the portal image and verify it exists
				checks := step.config.testChecks(t, resourceType, tName, tracker)
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// Log config and checks
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, checks.string(), stepName)

				testSteps = append(testSteps, resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				})
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    testSteps,
			})

		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *OrgNacPortalImageModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttrSet", "nacportal_id")
	checks.append(t, "TestCheckResourceAttr", "file", s.File)

	return checks
}
