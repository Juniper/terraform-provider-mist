package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	// gwc "github.com/terraform-provider-mist/internal/resource_device_gateway_cluster"
)

func (s *SiteModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(rType + "." + rName)
	checks.append(t, "TestCheckResourceAttr", "address", s.Address)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	return checks
}

func TestSiteModel(t *testing.T) {
	type testStep struct {
		config SiteModel
	}

	type testCase struct {
		//apiVersionConstraints version.Constraints
		steps []testStep
	}

	var FixtureSiteModel SiteModel

	b, err := os.ReadFile("fixtures/site_resource/site_resource_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fmt.Println(str)

	err = hcl.Decode(&FixtureSiteModel, str)
	if err != nil {
		fmt.Printf("error decoding hcl: %s\n", err)
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteModel{
						Address: "48 Balfour St, Woodstock, Cape Town, 7915, South Africa",
						Name:    "test-site",
						OrgId:   "901c5705-ca11-4bf1-9158-31f7195618ef",
					},
				},
			},
		},
		// "hcl_decode": {
		// 	steps: []testStep{
		// 		{
		// 			config: FixtureSiteModel,
		// 		},
		// 	},
		// },
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "mist_site"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName)
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
}
