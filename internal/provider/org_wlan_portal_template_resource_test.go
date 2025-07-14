// WIP
package provider

import (
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWlanPortalTemplateModel(t *testing.T) {
	testOrgID := GetTestOrgId()

	type testStep struct {
		config OrgWlanPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanPortalTemplateModel{
						OrgId: GetTestOrgId(),
						PortalTemplate: PortalTemplateValue{
							// Basic portal template configuration
							Company: &[]bool{true}[0],
							Email:   &[]bool{true}[0],
						},
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wlan_portal_template"

			// Create single-step tests with combined config (WLAN template + WLAN + portal template)
			// Since portal templates require a WLAN, and WLANs require a template,
			// we create all three in the same config but focus our checks on the portal template
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN
				combinedConfig, wlanRef := GetOrgWlanBaseConfig(testOrgID)

				// Generate the HCL configuration for the portal template
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				// Add the wlan_id attribute to the body before rendering
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig = combinedConfig + "\n\n" + Render("org_wlan_portal_template", tName, string(f.Bytes()))

				// Focus checks on the portal template resource (WLAN template and WLAN are prerequisites)
				// Note: This resource doesn't have an "id" field as it's more like a configuration overlay
				checks := newTestChecks(PrefixProviderName(resourceType) + "." + tName)
				checks.append(t, "TestCheckResourceAttr", "org_id", step.config.OrgId)
				checks.append(t, "TestCheckResourceAttrSet", "wlan_id")

				// Basic checks for portal template configuration
				// Check the element count and specific nested attributes
				checks.append(t, "TestCheckResourceAttrSet", "portal_template.%")
				if step.config.PortalTemplate.Company != nil && *step.config.PortalTemplate.Company {
					checks.append(t, "TestCheckResourceAttr", "portal_template.company", "true")
				}
				if step.config.PortalTemplate.Email != nil && *step.config.PortalTemplate.Email {
					checks.append(t, "TestCheckResourceAttr", "portal_template.email", "true")
				}

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
