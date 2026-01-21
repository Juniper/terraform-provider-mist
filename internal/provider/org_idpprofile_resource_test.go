package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_idpprofile"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgIdpprofileModel(t *testing.T) {
	type testStep struct {
		config OrgIdpprofileModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgIdpprofileModel{
						BaseProfile: "standard",
						Name:        "test_org_idp_profile",
						OrgId:       GetTestOrgId(),
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_idpprofile_resource/org_idpprofile_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgIdpprofileModel OrgIdpprofileModel

		err = hcl.Decode(&FixtureOrgIdpprofileModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgIdpprofileModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgIdpprofileModel,
				},
			},
		}
	}

	resourceType := "org_idpprofile"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_idpprofile.OrgIdpprofileResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				config.OrgId = GetTestOrgId()

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
	if tracker != nil {
		tracker.FieldCoverageReport(t)
	}
}

func (s *OrgIdpprofileModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "base_profile", s.BaseProfile)
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// Check overwrites list if present
	if len(s.Overwrites) > 0 {
		checks.append(t, "TestCheckResourceAttr", "overwrites.#", fmt.Sprintf("%d", len(s.Overwrites)))

		for i, overwrite := range s.Overwrites {
			prefix := fmt.Sprintf("overwrites.%d", i)

			// Required name field
			checks.append(t, "TestCheckResourceAttr", prefix+".name", overwrite.Name)

			// Optional action field
			if overwrite.Action != nil {
				checks.append(t, "TestCheckResourceAttr", prefix+".action", *overwrite.Action)
			}

			// Optional matching field
			if overwrite.Matching != nil {
				matching := overwrite.Matching

				// Check attack_name list if present
				if len(matching.AttackName) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".matching.attack_name.#", fmt.Sprintf("%d", len(matching.AttackName)))
					for j, attackName := range matching.AttackName {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".matching.attack_name.%d", j), attackName)
					}
				}

				// Check dst_subnet list if present
				if len(matching.DstSubnet) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".matching.dst_subnet.#", fmt.Sprintf("%d", len(matching.DstSubnet)))
					for j, dstSubnet := range matching.DstSubnet {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".matching.dst_subnet.%d", j), dstSubnet)
					}
				}

				// Check severity list if present
				if len(matching.Severity) > 0 {
					checks.append(t, "TestCheckResourceAttr", prefix+".matching.severity.#", fmt.Sprintf("%d", len(matching.Severity)))
					for j, severity := range matching.Severity {
						checks.append(t, "TestCheckResourceAttr", prefix+fmt.Sprintf(".matching.severity.%d", j), severity)
					}
				}
			}
		}
	}

	return checks
}
