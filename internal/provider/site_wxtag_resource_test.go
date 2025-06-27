package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	// gwc "github.com/terraform-provider-mist/internal/resource_device_gateway_cluster"
)

func TestSiteWxtagModel(t *testing.T) {
	testSiteID := GetTestSiteId()

	type testStep struct {
		config SiteWxtagModel
	}

	type testCase struct {
		//apiVersionConstraints version.Constraints
		steps []testStep
	}

	var FixtureSiteWxtagModel SiteWxtagModel

	b, err := os.ReadFile("fixtures/site_wtag_resource/site_wtag_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fmt.Println(str)

	err = hcl.Decode(&FixtureSiteWxtagModel, str)
	if err != nil {
		fmt.Printf("error decoding hcl: %s\n", err)
	}

	match := "ip_range_subnet"
	op := "in"
	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWxtagModel{
						SiteId: testSiteID,
						Name:   "wtag_test",
						Type:   "match",
						Match:  &match,
						Op:     &op,
						Values: []string{"10.3.0.0/16"},
					},
				},
			},
		},
		// "hcl_decode": {
		// 	steps: []testStep{
		// 		{
		// 			config: FixtureSiteWxtagModel,
		// 		},
		// 	},
		// },
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_wxtag"

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

func (s *SiteWxtagModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttr", "site_id", s.SiteId)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "type", s.Type)
	checks.append(t, "TestCheckResourceAttr", "match", *s.Match)
	checks.append(t, "TestCheckResourceAttr", "op", *s.Op)
	checks.append(t, "TestCheckResourceAttr", "values.#", fmt.Sprintf("%d", len(s.Values)))
	for i, v := range s.Values {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("values.%d", i), v)
	}

	return checks
}
