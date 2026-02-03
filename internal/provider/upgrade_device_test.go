package provider

import (
	"fmt"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_upgrade_device"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestUpgradeDeviceModel(t *testing.T) {
	resourceType := "upgrade_device"
	t.Skipf("Skipping %s tests, as they require a real device.", resourceType)

	type testStep struct {
		config UpgradeDeviceModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: UpgradeDeviceModel{
						DeviceId:      "test-device-id",
						TargetVersion: "0.14.29543",
					},
				},
			},
		},
	}

	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_upgrade_device.UpgradeDeviceResourceSchema(t.Context()).Attributes)
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

func (s *UpgradeDeviceModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Required attributes
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "device_id", s.DeviceId)
	checks.append(t, "TestCheckResourceAttr", "target_version", s.TargetVersion)

	// Computed-only attributes (check presence only)
	checks.append(t, "TestCheckResourceAttrSet", "auto_upgrade_stat.%")
	checks.append(t, "TestCheckResourceAttrSet", "config_timestamp")
	checks.append(t, "TestCheckResourceAttrSet", "config_version")
	checks.append(t, "TestCheckResourceAttrSet", "device_version")
	checks.append(t, "TestCheckResourceAttrSet", "ext_ip")
	checks.append(t, "TestCheckResourceAttrSet", "status")
	checks.append(t, "TestCheckResourceAttrSet", "tag_id")
	checks.append(t, "TestCheckResourceAttrSet", "tag_uuid")
	checks.append(t, "TestCheckResourceAttrSet", "timestamp")

	return checks
}
