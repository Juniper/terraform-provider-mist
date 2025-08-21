package provider

import (
	"fmt"
	"testing"

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

	// var FixtureOrgDeviceprofileAssignModel OrgDeviceprofileAssignModel

	// b, err := os.ReadFile("fixtures/site_setting_resource/site_setting_config.tf")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// str := string(b) // convert content to a 'string'

	// err = hcl.Decode(&FixtureOrgDeviceprofileAssignModel, str)
	// if err != nil {
	// 	fmt.Printf("error decoding hcl: %s\n", err)
	// }

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

	for tName, tCase := range testCases {
		t.Skip("Skipping org_deviceprofile_assign tests, as they require a device to assign.")
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_deviceprofile_assign"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				deviceProfileConfig, deviceprofileRef := GetOrgDeviceprofileApBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("deviceprofile_id", hclwrite.TokensForIdentifier(deviceprofileRef))
				combinedConfig := deviceProfileConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName)
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
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttrSet", "deviceprofile_id")
	checks.append(t, "TestCheckResourceAttr", "macs.#", fmt.Sprintf("%d", len(s.Macs)))
	for i, mac := range s.Macs {
		checks.append(t, "TestCheckTypeSetElemAttr", fmt.Sprintf("macs.%d", i), mac)
	}

	return checks
}
