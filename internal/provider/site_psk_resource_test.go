package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_psk"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSitePskModel(t *testing.T) {
	type testStep struct {
		config SitePskModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SitePskModel{
						Name:       "test-site-psk",
						Passphrase: "test-passphrase",
						Ssid:       "test-ssid",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_psk_resource/site_psk_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureSitePskModel SitePskModel
		err = hcl.Decode(&FixtureSitePskModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSitePskModel,
				},
			},
		}
	}

	resourceType := "site_psk"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site_psk.SitePskResourceSchema(t.Context()).Attributes)
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

func (s *SitePskModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "passphrase", s.Passphrase)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "ssid", s.Ssid)

	if s.Email != nil {
		checks.append(t, "TestCheckResourceAttr", "email", *s.Email)
	}

	if s.ExpireTime != nil {
		checks.append(t, "TestCheckResourceAttr", "expire_time", fmt.Sprintf("%d", *s.ExpireTime))
	}

	if s.ExpiryNotificationTime != nil {
		checks.append(t, "TestCheckResourceAttr", "expiry_notification_time", fmt.Sprintf("%d", *s.ExpiryNotificationTime))
	}

	if s.Mac != nil {
		checks.append(t, "TestCheckResourceAttr", "mac", *s.Mac)
	}

	if s.Note != nil {
		checks.append(t, "TestCheckResourceAttr", "note", *s.Note)
	}

	if s.NotifyExpiry != nil {
		checks.append(t, "TestCheckResourceAttr", "notify_expiry", fmt.Sprintf("%t", *s.NotifyExpiry))
	}

	if s.NotifyOnCreateOrEdit != nil {
		checks.append(t, "TestCheckResourceAttr", "notify_on_create_or_edit", fmt.Sprintf("%t", *s.NotifyOnCreateOrEdit))
	}

	if s.OldPassphrase != nil {
		checks.append(t, "TestCheckResourceAttr", "old_passphrase", *s.OldPassphrase)
	}

	if s.Role != nil {
		checks.append(t, "TestCheckResourceAttr", "role", *s.Role)
	}

	if s.Usage != nil {
		checks.append(t, "TestCheckResourceAttr", "usage", *s.Usage)
	}

	if s.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *s.VlanId)
	}

	return checks
}
