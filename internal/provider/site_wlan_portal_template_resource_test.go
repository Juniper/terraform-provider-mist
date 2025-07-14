package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteWlanPortalTemplateModel(t *testing.T) {
	type testStep struct {
		config SiteWlanPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	// var FixtureSiteWlanPortalTemplateModel SiteWlanPortalTemplateModel

	// b, err := os.ReadFile("fixtures/site_wlan_portal_template_resource/site_wlan_portal_template_config.tf")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// str := string(b) // convert content to a 'string'

	// err = hcl.Decode(&FixtureSiteWlanPortalTemplateModel, str)
	// if err != nil {
	// 	fmt.Printf("error decoding hcl: %s\n", err)
	// }

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWlanPortalTemplateModel{},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_wlan_portal_template"

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

func (s *SiteWlanPortalTemplateModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")

	return checks
}
