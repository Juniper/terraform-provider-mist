package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgInventoryModel(t *testing.T) {
	type testStep struct {
		config OrgInventoryModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgInventoryModel{
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
		fixtureOrgInventoryModel := OrgInventoryModel{}
		err = hcl.Decode(&fixtureOrgInventoryModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgInventoryModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgInventoryModel,
				},
			},
		}
	}

	resourceType := "org_inventory"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_inventory.OrgInventoryResourceSchema(t.Context()).Attributes)
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

				// Check if any inventory items need site_id and set up site config
				if config.Inventory != nil {
					for key, inventoryItem := range config.Inventory {
						if inventoryItem.SiteId != nil {
							// Set placeholder for site_id in inventory item
							inventoryItem.SiteId = stringPtr("{site_id}")
							config.Inventory[key] = inventoryItem
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

func (o *OrgInventoryModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

	// Check inventory map and all nested fields if inventory is configured
	if len(o.Inventory) > 0 {
		// Validate the inventory map length
		checks.append(t, "TestCheckResourceAttr", "inventory.%", fmt.Sprintf("%d", len(o.Inventory)))

		// Validate each inventory device
		for key, device := range o.Inventory {
			// Test all computed fields with TestCheckResourceAttrSet (since they're populated by the API)
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.id", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.mac", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.model", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.org_id", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.serial", key))
			checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.type", key))

			// Test configurable fields with expected values
			if device.SiteId != nil {
				checks.append(t, "TestCheckResourceAttrSet", fmt.Sprintf("inventory.%s.site_id", key))
			}
			if device.UnclaimWhenDestroyed != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("inventory.%s.unclaim_when_destroyed", key), fmt.Sprintf("%t", *device.UnclaimWhenDestroyed))
			}
		}
	}

	return checks
}
