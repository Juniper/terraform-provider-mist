package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgMxedgeInventoryModel(t *testing.T) {
	type testStep struct {
		config OrgMxedgeInventoryModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxedgeInventoryModel{
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_inventory_resource/org_inventory_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxedgeInventoryModel := OrgMxedgeInventoryModel{}
		err = hcl.Decode(&fixtureOrgMxedgeInventoryModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgMxedgeInventoryModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxedgeInventoryModel,
				},
			},
		}
	}

	resourceType := "org_mxedge_inventory"
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			// Skip fixture cases that require real devices with valid MAC addresses
			if strings.HasPrefix(tName, "fixture_case") {
				t.Skip("Skipping fixture case as it requires real devices with valid MAC addresses.")
			}

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				siteConfig, siteRef := "", ""

				// Check if any mxedge items need site_id and set up site config
				if config.Mxedges != nil {
					for key, mxedgeItem := range config.Mxedges {
						if mxedgeItem.SiteId != nil {
							// Set placeholder for site_id in mxedge item
							mxedgeItem.SiteId = stringPtr("{site_id}")
							config.Mxedges[key] = mxedgeItem
						}
					}
				}

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))
				siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())

				configStr := ""
				if siteConfig != "" {
					combinedConfig = strings.ReplaceAll(combinedConfig, "\"{site_id}\"", siteRef)
					configStr = siteConfig + "\n\n"
				}
				combinedConfig = configStr + combinedConfig

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

func (o *OrgMxedgeInventoryModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

	// Check inventory map and all nested fields if inventory is configured
	if len(o.Mxedges) > 0 {
		// Validate the inventory map length
		checks.append(t, "TestCheckResourceAttr", "mxedges.%", fmt.Sprintf("%d", len(o.Mxedges)))

		// Validate each inventory device
		for key, device := range o.Mxedges {
			// Test all computed fields with TestCheckResourceAttrSet (since they're populated by the API)
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.id", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.model", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.org_id", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.serial", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.type", key))

			// Test configurable fields with expected values
			if device.SiteId != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("mxedges.%s.site_id", key))
			}

		}
	}

	return checks
}
