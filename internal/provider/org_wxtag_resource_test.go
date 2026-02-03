package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wxtag"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgWxtagModel(t *testing.T) {
	type testStep struct {
		config OrgWxtagModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWxtagModel{
						OrgId: GetTestOrgId(),
						Name:  "Test_WxTag",
						Type:  "",
					},
				},
			},
		},
	}

	// Load fixture data
	b, err := os.ReadFile("fixtures/org_wxtag_resource/org_wxtag_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgWxtagModel := OrgWxtagModel{}
		err = hcl.Decode(&fixtureOrgWxtagModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWxtagModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWxtagModel,
				},
			},
		}
	}

	resourceType := "org_wxtag"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_wxtag.OrgWxtagResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render("org_wxtag", tName, string(f.Bytes()))

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

func (s *OrgWxtagModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Check fields in struct order
	// 1. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 2. Mac (optional)
	if s.Mac != nil {
		checks.append(t, "TestCheckResourceAttr", "mac", *s.Mac)
	}

	// 3. Match (optional)
	if s.Match != nil {
		checks.append(t, "TestCheckResourceAttr", "match", *s.Match)
	}

	// 4. Name (required)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// 5. Op (optional)
	if s.Op != nil {
		checks.append(t, "TestCheckResourceAttr", "op", *s.Op)
	}

	// 6. OrgId (required)
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)

	// 7. Specs (optional nested object array) - test child attributes only
	if len(s.Specs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "specs.#", fmt.Sprintf("%d", len(s.Specs)))
		for i, spec := range s.Specs {
			if spec.Protocol != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.protocol", i), *spec.Protocol)
			}
			if spec.PortRange != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.port_range", i), *spec.PortRange)
			}
			if len(spec.Subnets) > 0 {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.subnets.#", i), fmt.Sprintf("%d", len(spec.Subnets)))
				for j, subnet := range spec.Subnets {
					checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("specs.%d.subnets.%d", i, j), subnet)
				}
			}
		}
	}

	// 8. Type (required)
	checks.append(t, "TestCheckResourceAttr", "type", s.Type)

	// 9. Values (optional array)
	if len(s.Values) > 0 {
		checks.append(t, "TestCheckResourceAttr", "values.#", fmt.Sprintf("%d", len(s.Values)))
		for i, value := range s.Values {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("values.%d", i), value)
		}
	}

	// 10. VlanId (optional)
	if s.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *s.VlanId)
	}

	return checks
}
