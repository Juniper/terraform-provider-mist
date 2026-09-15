package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxcluster"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgMxclusterModel(t *testing.T) {
	type testStep struct {
		config OrgMxclusterModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgMxclusterModel{
						Name:  "Test Org Mxcluster",
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_mxcluster_resource/org_mxcluster_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgMxclusterModel := OrgMxclusterModel{}
		err = hcl.Decode(&fixtureOrgMxclusterModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		// Need to set org_id to required field as it is used in the url for the resource
		fixtureOrgMxclusterModel.OrgId = GetTestOrgId()
		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgMxclusterModel,
				},
			},
		}
	}

	resourceType := "org_mxcluster"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_mxcluster.OrgMxclusterResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				siteConfig, siteRef := "", ""

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())

				combinedConfig := Render(resourceType, tName, string(f.Bytes()))
				if config.SiteId != nil && *config.SiteId != "" {
					siteConfig, siteRef = GetSiteBaseConfig(GetTestOrgId())
				}

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

func (o *OrgMxclusterModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	// site_id is injected as a terraform reference, not a literal value
	appendReflectChecks(t, &checks, o, "site_id")
	if o.SiteId != nil {
		checks.append(t, "TestCheckResourceAttrSet", "site_id")
	}
	return checks
}
