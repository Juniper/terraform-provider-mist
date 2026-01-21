package provider

import (
	"fmt"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_evpn_topology"
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

	// var FixtureOrgEvpnTopologyModel OrgEvpnTopologyModel

	// b, err := os.ReadFile("fixtures/site_setting_resource/site_setting_config.tf")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// str := string(b) // convert content to a 'string'

	// err = hcl.Decode(&FixtureOrgEvpnTopologyModel, str)
	// if err != nil {
	// 	fmt.Printf("error decoding hcl: %s\n", err)
	// }

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

	resourceType := "org_evpn_topology"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_evpn_topology.OrgEvpnTopologyResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Skip("Skipping EVPN Topology tests temporarily")
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
	tracker.FieldCoverageReport(t)
}

func (s *OrgEvpnTopologyModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	return checks
}
