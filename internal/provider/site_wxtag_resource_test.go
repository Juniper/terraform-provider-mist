package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wxtag"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteWxtagModel(t *testing.T) {
	type testStep struct {
		config SiteWxtagModel
	}

	type testCase struct {
		steps []testStep
	}

	match := "ip_range_subnet"
	op := "in"
	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWxtagModel{
						Name:   "wtag_test",
						Type:   "match",
						Match:  &match,
						Op:     &op,
						Values: []string{"10.3.0.0/16"},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_wxtag_resource/site_wxtag_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureSiteWxtagModel SiteWxtagModel
		err = hcl.Decode(&FixtureSiteWxtagModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteWxtagModel,
				},
			},
		}
	}

	resourceType := "site_wxtag"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site_wxtag.SiteWxtagResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *SiteWxtagModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Always present attributes
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "type", s.Type)

	// Optional attributes with conditional checks
	if s.Mac != nil {
		checks.append(t, "TestCheckResourceAttr", "mac", *s.Mac)
	}
	if s.Match != nil {
		checks.append(t, "TestCheckResourceAttr", "match", *s.Match)
	}
	if s.Op != nil {
		checks.append(t, "TestCheckResourceAttr", "op", *s.Op)
	}
	if s.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *s.VlanId)
	}

	// Values array checks
	if len(s.Values) > 0 {
		checks.append(t, "TestCheckResourceAttr", "values.#", fmt.Sprintf("%d", len(s.Values)))
		for i, v := range s.Values {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("values.%d", i), v)
		}
	} else {
		checks.append(t, "TestCheckResourceAttr", "values.#", "0")
	}

	// Specs array checks
	if len(s.Specs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "specs.#", fmt.Sprintf("%d", len(s.Specs)))
		for i, spec := range s.Specs {
			if spec.PortRange != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.port_range", i), *spec.PortRange)
			}
			if spec.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.protocol", i), *spec.Protocol)
			}
			if len(spec.Subnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.subnets.#", i), fmt.Sprintf("%d", len(spec.Subnets)))
				for j, subnet := range spec.Subnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.subnets.%d", i, j), subnet)
				}
			}
		}
	} else {
		checks.append(t, "TestCheckResourceAttr", "specs.#", "0")
	}

	return checks
}
