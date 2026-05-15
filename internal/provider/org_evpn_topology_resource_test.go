package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_evpn_topology"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgEvpnTopologyModel(t *testing.T) {

	type testStep struct {
		config OrgEvpnTopologyModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgEvpnTopologyModel{
						Name:  "test_evpn_topology",
						OrgId: GetTestOrgId(),
						Switches: map[string]OrgEvpnTopologySwitchesValue{
							"000000000001": { // Placeholder MAC for testing
								Role: "none",
							},
						},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_evpn_topology_resource/org_evpn_topology_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgEvpnTopologyModel OrgEvpnTopologyModel
		err = hcl.Decode(&FixtureOrgEvpnTopologyModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgEvpnTopologyModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgEvpnTopologyModel,
				},
			},
		}
	}

	resourceType := "org_evpn_topology"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_evpn_topology.OrgEvpnTopologyResourceSchema(t.Context()).Attributes)
	t.Skipf("Skipping %s tests, as they require a real device.", resourceType)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				// log config and checks here
				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: configStr,
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

func (s *OrgEvpnTopologyModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	checks.append(t, "TestCheckResourceAttrSet", "org_id")
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
