package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	// gwc "github.com/terraform-provider-mist/internal/resource_device_gateway_cluster"
)

func TestSiteWlanPortalImageModel(t *testing.T) {
	testSiteID := GetTestSiteId()
	// Create a test PNG file
	testImagePath := CreateTestPNGFile(t)

	type testStep struct {
		config SiteWlanPortalImageModel
	}

	type testCase struct {
		//apiVersionConstraints version.Constraints
		steps []testStep
	}

	var FixtureSiteWlanPortalImageModel SiteWlanPortalImageModel

	b, err := os.ReadFile("fixtures/site_wlan_portal_image_resource/site_wlan_portal_image_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fmt.Println(str)

	err = hcl.Decode(&FixtureSiteWlanPortalImageModel, str)
	if err != nil {
		fmt.Printf("error decoding hcl: %s\n", err)
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWlanPortalImageModel{
						File:   testImagePath,
						SiteId: testSiteID,
					},
				},
			},
		},
		// "hcl_decode": {
		// 	steps: []testStep{
		// 		{
		// 			config: FixtureSiteWlanPortalImageModel,
		// 		},
		// 	},
		// },
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_wlan_portal_image"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN
				wlanConfig, wlanRef := GetSiteWlanBaseConfig(testSiteID)
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				// Add the wlan_id attribute to the body before rendering
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig := wlanConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName)
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
}

func (s *SiteWlanPortalImageModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttr", "file", s.File)
	checks.append(t, "TestCheckResourceAttr", "site_id", s.SiteId)
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")

	return checks
}
