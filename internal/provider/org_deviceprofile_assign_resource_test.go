package provider

import (
	"fmt"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_assign"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgDeviceprofileAssignModel(t *testing.T) {
	type testStep struct {
		config OrgDeviceprofileAssignModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgDeviceprofileAssignModel{
						Macs:  []string{"00:11:22:33:44:55"},
						OrgId: GetTestOrgId(),
					},
				},
			},
		},
	}

	resourceType := "org_deviceprofile_assign"
	var checks testChecks
	for tName, tCase := range testCases {
		t.Skip("Skipping Device Profile Assign tests temporarily")
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				deviceProfileConfig, deviceprofileRef := GetOrgDeviceprofileApBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("deviceprofile_id", hclwrite.TokensForIdentifier(deviceprofileRef))
				combinedConfig := deviceProfileConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

				checks = config.testChecks(t, resourceType, tName)
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
	FieldCoverageReport(t, &checks)
}

func GetOrgDeviceprofileApBaseConfig(org_ID string) (config string, deviceprofileRef string) {
	deviceProfileApConfig := OrgDeviceprofileApModel{
		OrgId: org_ID,
		Name:  "test_ap_profile",
	}

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&deviceProfileApConfig, f.Body())
	deviceProfileApConfigStr := Render("org_deviceprofile_ap", deviceProfileApConfig.Name, string(f.Bytes()))

	return deviceProfileApConfigStr, fmt.Sprintf("mist_org_deviceprofile_ap.%s.id", deviceProfileApConfig.Name)
}

func (s *OrgDeviceprofileAssignModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	TrackFieldCoverage(t, &checks, "org_deviceprofile_assign", resource_org_deviceprofile_assign.OrgDeviceprofileAssignResourceSchema)

	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "deviceprofile_id")
	checks.append(t, "TestCheckResourceAttr", "macs.#", fmt.Sprintf("%d", len(s.Macs)))
	for i, mac := range s.Macs {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("macs.%d", i), mac)
	}

	return checks
}
