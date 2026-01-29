package provider

import (
	"fmt"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wlan_portal_image"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteWlanPortalImageModel(t *testing.T) {
	// Create a test PNG file
	testImagePath := CreateTestPNGFile(t)

	type testStep struct {
		config SiteWlanPortalImageModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWlanPortalImageModel{
						File: testImagePath,
					},
				},
			},
		},
	}

	resourceType := "site_wlan_portal_image"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site_wlan_portal_image.SiteWlanPortalImageResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate combined config: Site + WLAN
				wlanConfig, siteRef, wlanRef := GetSiteWlanBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				// Add the site_id and wlan_id attributes to the body before rendering
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig := wlanConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *SiteWlanPortalImageModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	checks.append(t, "TestCheckResourceAttr", "file", s.File)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")

	return checks
}
