package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_rftemplate"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgRftemplateModel(t *testing.T) {
	type testStep struct {
		config OrgRftemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgRftemplateModel{
						OrgId: GetTestOrgId(),
						Name:  "TestRfTemplate",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_rftemplate_resource/org_rftemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgRftemplateModel := OrgRftemplateModel{}
		err = hcl.Decode(&fixtureOrgRftemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgRftemplateModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgRftemplateModel,
				},
			},
		}
	}

	resourceType := "org_rftemplate"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_rftemplate.OrgRftemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			// Create single-step tests with generated config
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				// Generate checks for the rftemplate resource
				checks := step.config.testChecks(t, resourceType, tName, tracker)
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

func (o *OrgRftemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, o)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	return checks
}
