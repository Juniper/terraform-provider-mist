package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDeviceImageModel(t *testing.T) {
	testImagePath := CreateTestPNGFile(t)

	type testStep struct {
		config DeviceImageModel
	}

	type testCase struct {
		steps []testStep
	}

	// var FixtureDeviceImageModel DeviceImageModel

	// b, err := os.ReadFile("fixtures/site_setting_resource/site_setting_config.tf")
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// str := string(b) // convert content to a 'string'

	// err = hcl.Decode(&FixtureDeviceImageModel, str)
	// if err != nil {
	// 	fmt.Printf("error decoding hcl: %s\n", err)
	// }

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: DeviceImageModel{
						DeviceId:    "",
						File:        testImagePath,
						ImageNumber: 1,
					},
				},
			},
		},
	}

	for tName, tCase := range testCases {
		t.Skip("Skipping test case: " + tName) // Skip all tests for now
		t.Run(tName, func(t *testing.T) {
			resourceType := "device_image"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				combinedConfig := siteConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *DeviceImageModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "file", s.File)
	checks.append(t, "TestCheckResourceAttrSet", "device_id")
	checks.append(t, "TestCheckResourceAttr", "image_number", fmt.Sprintf("%d", s.ImageNumber))

	return checks
}
