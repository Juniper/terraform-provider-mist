// WIP
package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlan_portal_image"
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

	resourceType := "org_wlan_portal_image"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_wlan_portal_image.OrgWlanPortalImageResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			testSteps := make([]resource.TestStep, 0)
			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN + portal image
				combinedConfig, wlanRef := GetOrgWlanBaseConfig(testOrgID)

				// Generate the HCL configuration for the portal image
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig = combinedConfig + "\n\n" + Render("org_wlan_portal_image", tName, string(f.Bytes()))

				// Step 1: Create the portal image and verify it exists
				checks := step.config.testChecks(t, resourceType, tName, tracker)
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// Log config and checks for step 1
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, checks.string(), stepName)

				testSteps = append(testSteps, resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				})

				// Step 2: Same config but now also check the WLAN's portal_image field
				// This gives the API time to propagate the changes
				wlanPortalImageChecks := newTestChecks("mist_org_wlan.wlanName", tracker)
				wlanPortalImageChecks.append(t, "TestCheckResourceAttr", "portal_image", "present")

				// Combine both portal image and WLAN checks for the second step
				allChecks := make([]resource.TestCheckFunc, 0)
				allChecks = append(allChecks, checks.checks...)
				allChecks = append(allChecks, wlanPortalImageChecks.checks...)

				// Combined check log
				combinedCheckLog := checks.string() + wlanPortalImageChecks.string()
				stepName2 := fmt.Sprintf("test case %s step %d", tName, i+2)

				// Log checks for step 2
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName2, combinedCheckLog, stepName2)

				testSteps = append(testSteps, resource.TestStep{
					PreConfig: func() {
						// Small delay to allow API propagation
						time.Sleep(2 * time.Second)
					},
					// NOTE: Using the same combinedConfig as Step 1 - this doesn't recreate resources!
					// Terraform sees identical config and runs "no changes" plan, but executes new checks.
					// The delay above allows Mist API time to propagate the portal image URL to the WLAN's portal_image field.
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(allChecks...),
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

func (s *OrgWlanPortalImageModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")
	checks.append(t, "TestCheckResourceAttr", "file", s.File)

	return checks
}
