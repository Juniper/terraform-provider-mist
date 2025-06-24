// WIP
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteEvpnTopologyModel(t *testing.T) {
	t.Skip("This test is currently skipped because it requires a real device with a valid MAC address.")
	type testStep struct {
		config SiteEvpnTopologyModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteEvpnTopologyModel{
						Name: "Test_EVPN_Topology",
						// Use a placeholder MAC address for acceptance tests
						// This won't correspond to a real device but satisfies the schema requirement
						Switches: map[string]SwitchesValue{
							"000000000001": { // Placeholder MAC for testing
								Role: "access",
							},
						},
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_evpn_topology"
			siteName := "test_site"

			// Create single-step tests with combined config (site + EVPN topology)
			// Since site EVPN topologies require a site, we create both in the same config
			// but focus our checks on the EVPN topology resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: site + EVPN topology
				combinedConfig := generateSiteEvpnTopologyTestConfig(siteName, tName, step.config)

				// Focus checks on the EVPN topology resource (site is just a prerequisite)
				checks := newTestChecks(PrefixProviderName(resourceType) + "." + tName)
				checks.append(t, "TestCheckResourceAttrSet", "id")
				checks.append(t, "TestCheckResourceAttrSet", "site_id")
				checks.append(t, "TestCheckResourceAttr", "name", step.config.Name)
				checks.append(t, "TestCheckResourceAttrSet", "org_id")

				// Basic checks for switches configuration (using placeholder MAC addresses)
				if len(step.config.Switches) > 0 {
					checks.append(t, "TestCheckResourceAttrSet", "switches.%")
					// Note: We use placeholder MAC addresses for testing since actual devices
					// are not available in the test environment
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

// generateSiteEvpnTopologyTestConfig creates a combined configuration with both a site and a site EVPN topology
// This handles the prerequisite that site EVPN topologies require a site to exist
func generateSiteEvpnTopologyTestConfig(siteName, evpnTopologyName string, evpnTopologyConfig SiteEvpnTopologyModel) string {
	// Create the prerequisite site
	siteConfig := SiteModel{
		Name:    "Test_Site",
		Address: "Test Address",
		OrgId:   GetTestOrgId(),
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&siteConfig, f.Body())
	siteConfigStr := Render("site", siteName, string(f.Bytes()))

	// Create the EVPN topology that references the site
	siteRef := fmt.Sprintf("mist_site.%s.id", siteName)

	f = hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&evpnTopologyConfig, f.Body())

	// Add the site_id attribute to the body before rendering
	f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
	evpnTopologyConfigStr := Render("site_evpn_topology", evpnTopologyName, string(f.Bytes()))

	// Combine both configs
	return siteConfigStr + "\n\n" + evpnTopologyConfigStr
}
