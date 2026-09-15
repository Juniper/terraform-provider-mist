package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nacrule"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacruleModel(t *testing.T) {
	type testStep struct {
		config OrgNacruleModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacruleModel{
						OrgId:  GetTestOrgId(),
						Name:   "test-nacrule",
						Action: "allow",
						Order:  1,
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nacrule_resource/org_nacrule_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNacruleModel OrgNacruleModel
		err = hcl.Decode(&FixtureOrgNacruleModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNacruleModel,
				},
			},
		}
	}

	resourceType := "org_nacrule"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nacrule.OrgNacruleResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				config.OrgId = GetTestOrgId()
				siteConfig, siteRef := "", ""

				// Check if config needs site_id and set up site config
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := string(f.Bytes())

				if strings.Contains(configStr, "{site_id}") || strings.Contains(configStr, "{sitegroup_id}") {
					siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
					// For now, use the same site ID for sitegroup_id as a placeholder
					configStr = strings.ReplaceAll(configStr, "\"{site_id}\"", siteRef)
					configStr = strings.ReplaceAll(configStr, "\"{sitegroup_id}\"", siteRef)
				}

				combinedConfig := Render(resourceType, tName, configStr)

				if siteConfig != "" {
					combinedConfig = siteConfig + "\n\n" + combinedConfig
				}

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

func (o *OrgNacruleModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	// site_ids and sitegroup_ids hold terraform reference placeholders, not literal UUIDs
	appendReflectChecks(t, &checks, o,
		"matching.site_ids", "matching.sitegroup_ids",
		"not_matching.site_ids", "not_matching.sitegroup_ids")
	checks.append(t, "TestCheckResourceAttrSet", "id")

	return checks
}
