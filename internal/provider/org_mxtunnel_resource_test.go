package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxtunnel"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgMxtunnelModel(t *testing.T) {
	type testStep struct {
		config OrgMxtunnelModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxtunnelModel{
						Name:  "Test Org Mxtunnel",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_mxtunnel_resource/org_mxtunnel_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxtunnelModel := OrgMxtunnelModel{}
		err = hcl.Decode(&fixtureOrgMxtunnelModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgMxtunnelModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxtunnelModel,
				},
			},
		}
	}

	resourceType := "org_mxtunnel"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_mxtunnel.OrgMxtunnelResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())

				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, combinedConfig, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end checks for %s ------\n\n", stepName, chkLog, stepName)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}
			}

			resource.Test(t, resource.TestCase{
				IsUnitTest:               true,
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})
		})
	}
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (o *OrgMxtunnelModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, o)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")

	return checks
}
