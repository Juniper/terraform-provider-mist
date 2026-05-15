// WIP
package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_evpn_topology"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteEvpnTopologyModel(t *testing.T) {
	resourceType := "site_evpn_topology"
	t.Skipf("Skipping %s tests, as they require a real device.", resourceType)

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
						Switches: map[string]SiteEvpnTopologySwitchesValue{
							"000000000001": { // Placeholder MAC for testing
								Role: "access",
							},
						},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_evpn_topology_resource/site_evpn_topology_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureSiteEvpnTopologyModel SiteEvpnTopologyModel
		err = hcl.Decode(&FixtureSiteEvpnTopologyModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteEvpnTopologyModel,
				},
			},
		}
	}

	siteName := "test_site"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site_evpn_topology.SiteEvpnTopologyResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			// Create single-step tests with combined config (site + EVPN topology)
			// Since site EVPN topologies require a site, we create both in the same config
			// but focus our checks on the EVPN topology resource being tested
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: site + EVPN topology
				combinedConfig := generateSiteEvpnTopologyTestConfig(siteName, tName, step.config)

				// Focus checks on the EVPN topology resource (site is just a prerequisite)
				checks := step.config.testChecks(t, resourceType, tName, tracker)

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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
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

func (s *SiteEvpnTopologyModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	if s.EvpnOptions != nil {
		if s.EvpnOptions.AutoLoopbackSubnet != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.auto_loopback_subnet", *s.EvpnOptions.AutoLoopbackSubnet)
		}
		if s.EvpnOptions.AutoLoopbackSubnet6 != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.auto_loopback_subnet6", *s.EvpnOptions.AutoLoopbackSubnet6)
		}
		if s.EvpnOptions.AutoRouterIdSubnet != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.auto_router_id_subnet", *s.EvpnOptions.AutoRouterIdSubnet)
		}
		if s.EvpnOptions.AutoRouterIdSubnet6 != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.auto_router_id_subnet6", *s.EvpnOptions.AutoRouterIdSubnet6)
		}
		if s.EvpnOptions.CoreAsBorder != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.core_as_border", fmt.Sprintf("%t", *s.EvpnOptions.CoreAsBorder))
		}
		if s.EvpnOptions.EnableInbandMgmt != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.enable_inband_mgmt", fmt.Sprintf("%t", *s.EvpnOptions.EnableInbandMgmt))
		}
		if s.EvpnOptions.EnableInbandZtp != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.enable_inband_ztp", fmt.Sprintf("%t", *s.EvpnOptions.EnableInbandZtp))
		}
		if s.EvpnOptions.Overlay != nil {
			if s.EvpnOptions.Overlay.As != nil {
				checks.append(t, "TestCheckResourceAttr", "evpn_options.overlay.as", fmt.Sprintf("%d", *s.EvpnOptions.Overlay.As))
			}
		}
		if s.EvpnOptions.PerVlanVgaV4Mac != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.per_vlan_vga_v4_mac", fmt.Sprintf("%t", *s.EvpnOptions.PerVlanVgaV4Mac))
		}
		if s.EvpnOptions.PerVlanVgaV6Mac != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.per_vlan_vga_v6_mac", fmt.Sprintf("%t", *s.EvpnOptions.PerVlanVgaV6Mac))
		}
		if s.EvpnOptions.RoutedAt != nil {
			checks.append(t, "TestCheckResourceAttr", "evpn_options.routed_at", *s.EvpnOptions.RoutedAt)
		}
		if s.EvpnOptions.Underlay != nil {
			if s.EvpnOptions.Underlay.AsBase != nil {
				checks.append(t, "TestCheckResourceAttr", "evpn_options.underlay.as_base", fmt.Sprintf("%d", *s.EvpnOptions.Underlay.AsBase))
			}
			if s.EvpnOptions.Underlay.RoutedIdPrefix != nil {
				checks.append(t, "TestCheckResourceAttr", "evpn_options.underlay.routed_id_prefix", *s.EvpnOptions.Underlay.RoutedIdPrefix)
			}
			if s.EvpnOptions.Underlay.Subnet != nil {
				checks.append(t, "TestCheckResourceAttr", "evpn_options.underlay.subnet", *s.EvpnOptions.Underlay.Subnet)
			}
			if s.EvpnOptions.Underlay.UseIpv6 != nil {
				checks.append(t, "TestCheckResourceAttr", "evpn_options.underlay.use_ipv6", fmt.Sprintf("%t", *s.EvpnOptions.Underlay.UseIpv6))
			}
		}
		for key, vsInstance := range s.EvpnOptions.VsInstances {
			if len(vsInstance.Networks) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("evpn_options.vs_instances.%s.networks.#", key), fmt.Sprintf("%d", len(vsInstance.Networks)))
				for i, network := range vsInstance.Networks {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("evpn_options.vs_instances.%s.networks.%d", key, i), network)
				}
			}
		}
	}
	if len(s.PodNames) > 0 {
		for key, value := range s.PodNames {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("pod_names.%s", key), value)
		}
	}
	for key, sw := range s.Switches {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switches.%s.role", key), sw.Role)
		if sw.Pod != nil {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switches.%s.pod", key), fmt.Sprintf("%d", *sw.Pod))
		}
		if len(sw.Pods) > 0 {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switches.%s.pods.#", key), fmt.Sprintf("%d", len(sw.Pods)))
			for i, pod := range sw.Pods {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("switches.%s.pods.%d", key, i), fmt.Sprintf("%d", pod))
			}
		}
	}

	return checks
}
