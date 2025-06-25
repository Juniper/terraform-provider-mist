// WIP
package provider

import (
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWxtagModel(t *testing.T) {
	type testStep struct {
		config OrgWxtagModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWxtagModel{
						OrgId: GetTestOrgId(),
						Name:  "Test_WxTag",
						Type:  "",
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wxtag"

			// Create single-step tests with WX Tag config
			// WX Tags are standalone resources and don't require prerequisites
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate WX Tag config directly
				config := generateOrgWxtagTestConfig(tName, step.config)

				// Focus checks on the WX Tag resource
				checks := newTestChecks(PrefixProviderName(resourceType) + "." + tName)
				checks.append(t, "TestCheckResourceAttrSet", "id")
				checks.append(t, "TestCheckResourceAttrSet", "org_id")
				checks.append(t, "TestCheckResourceAttr", "org_id", step.config.OrgId)
				checks.append(t, "TestCheckResourceAttr", "name", step.config.Name)
				checks.append(t, "TestCheckResourceAttr", "type", step.config.Type)

				steps[i] = resource.TestStep{
					Config: config,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration for debugging
				t.Logf("\n// ------ WX Tag Config ------\n%s\n", config)
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})

		})
	}
}

// generateOrgWxtagTestConfig creates a configuration for a standalone WX Tag
// WX Tags don't require prerequisites, so this is a simple direct generation
func generateOrgWxtagTestConfig(wxtagName string, wxtagConfig OrgWxtagModel) string {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&wxtagConfig, f.Body())
	return Render("org_wxtag", wxtagName, string(f.Bytes()))
}
