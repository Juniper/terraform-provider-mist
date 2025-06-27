// WIP
package provider

import (
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanPortalImageModel(t *testing.T) {
	testOrgID := GetTestOrgId()
	// Create a test PNG file
	testImagePath := CreateTestPNGFile(t)

	type testStep struct {
		config OrgWlanPortalImageModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanPortalImageModel{
						OrgId: testOrgID,
						File:  testImagePath,
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wlan_portal_image"

			// Create single-step tests with combined config (WLAN template + WLAN + portal image)
			// Since portal images require a WLAN, and WLANs require a template,
			// we create all three in the same config but focus our checks on the portal image
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN
				combinedConfig, wlanRef := GetOrgWlanBaseConfig(testOrgID)

				// Generate the HCL configuration for the portal template
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				// Add the wlan_id attribute to the body before rendering
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig = combinedConfig + "\n\n" + Render("org_wlan_portal_image", tName, string(f.Bytes()))

				// Focus checks on the portal image resource (WLAN template and WLAN are prerequisites)
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

func (s *OrgWlanPortalImageModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")
	checks.append(t, "TestCheckResourceAttr", "file", s.File)

	return checks
}
