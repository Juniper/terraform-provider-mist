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

func TestOrgVpnModel(t *testing.T) {
	type testStep struct {
		config OrgVpnModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgVpnModel{
						OrgId: stringPtr(GetTestOrgId()),
						Name:  "test-vpn",
						Paths: map[string]OrgVpnPathsValue{
							"path1": {},
						},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_vpn_resource/org_vpn_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgVpnModel OrgVpnModel
		err = hcl.Decode(&FixtureOrgVpnModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgVpnModel.OrgId = stringPtr(GetTestOrgId())

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgVpnModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_vpn"
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, PrefixProviderName(resourceType), tName)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

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
}

func (o *OrgVpnModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// Check computed-only fields (verify they exist)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// Check optional string fields
	if o.OrgId != nil {
		checks.append(t, "TestCheckResourceAttr", "org_id", *o.OrgId)
	}
	if o.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *o.Type)
	}

	// Check PathSelection nested object
	if o.PathSelection != nil {
		if o.PathSelection.Strategy != nil {
			checks.append(t, "TestCheckResourceAttr", "path_selection.strategy", *o.PathSelection.Strategy)
		}
	}

	// Check Paths map
	if len(o.Paths) > 0 {
		checks.append(t, "TestCheckResourceAttr", "paths.%", fmt.Sprintf("%d", len(o.Paths)))
		for pathKey, path := range o.Paths {
			// Check optional string fields in each path
			if path.BfdProfile != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.bfd_profile", pathKey), *path.BfdProfile)
			}
			if path.Ip != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.ip", pathKey), *path.Ip)
			}

			// Check optional boolean fields in each path
			if path.BfdUseTunnelMode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.bfd_use_tunnel_mode", pathKey), fmt.Sprintf("%t", *path.BfdUseTunnelMode))
			}

			// Check optional integer fields in each path
			if path.Pod != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.pod", pathKey), fmt.Sprintf("%d", *path.Pod))
			}

			// Check PeerPaths nested map
			if len(path.PeerPaths) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.peer_paths.%%", pathKey), fmt.Sprintf("%d", len(path.PeerPaths)))
				for peerKey, peer := range path.PeerPaths {
					if peer.Preference != nil {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.peer_paths.%s.preference", pathKey, peerKey), fmt.Sprintf("%d", *peer.Preference))
					}
				}
			}

			// Check TrafficShaping nested object
			if path.TrafficShaping != nil {
				if path.TrafficShaping.Enabled != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.traffic_shaping.enabled", pathKey), fmt.Sprintf("%t", *path.TrafficShaping.Enabled))
				}
				if path.TrafficShaping.MaxTxKbps != nil {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.traffic_shaping.max_tx_kbps", pathKey), fmt.Sprintf("%d", *path.TrafficShaping.MaxTxKbps))
				}
				if len(path.TrafficShaping.ClassPercentage) > 0 {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.traffic_shaping.class_percentage.#", pathKey), fmt.Sprintf("%d", len(path.TrafficShaping.ClassPercentage)))
					for i, percentage := range path.TrafficShaping.ClassPercentage {
						checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("paths.%s.traffic_shaping.class_percentage.%d", pathKey, i), fmt.Sprintf("%d", percentage))
					}
				}
			}
		}
	}

	return checks
}
