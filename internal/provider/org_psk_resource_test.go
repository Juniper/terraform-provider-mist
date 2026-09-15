package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_psk"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgPskModel(t *testing.T) {
	type testStep struct {
		config OrgPskModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgPskModel{
						OrgId:      GetTestOrgId(),
						Name:       "test-psk",
						Passphrase: "testpassphrase",
						Ssid:       "test-ssid",
					},
				},
			},
		},
	}

	// Load fixture data
	b, err := os.ReadFile("fixtures/org_psk_resource/org_psk_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgPskModel := OrgPskModel{}
		err = hcl.Decode(&fixtureOrgPskModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgPskModel.OrgId = GetTestOrgId()

		// Set expire_time to now + 1 year (in epoch seconds) if it was set in fixture
		if fixtureOrgPskModel.ExpireTime != nil && *fixtureOrgPskModel.ExpireTime == 1 {
			futureTime := time.Now().AddDate(1, 0, 0).Unix()
			fixtureOrgPskModel.ExpireTime = &futureTime
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgPskModel,
				},
			},
		}
	}

	resourceType := "org_psk"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_psk.OrgPskResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render("org_psk", tName, string(f.Bytes()))

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

func (o *OrgPskModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, o)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	return checks
}
