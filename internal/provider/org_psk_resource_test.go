package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_psk"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgPskModel(t *testing.T) {
	type testStep struct {
		config OrgPskModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgPskModel{
						OrgId:      GetTestOrgId(),
						Name:       "test-psk",
						Passphrase: "testpassphrase",
						Ssid:       "test-ssid",
					},
				},
			},
		},
	}

	// Load fixture data
	b, err := os.ReadFile("fixtures/org_psk_resource/org_psk_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgPskModel := OrgPskModel{}
		err = hcl.Decode(&fixtureOrgPskModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgPskModel.OrgId = GetTestOrgId()

		// Set expire_time to now + 1 year (in epoch seconds) if it was set in fixture
		if fixtureOrgPskModel.ExpireTime != nil && *fixtureOrgPskModel.ExpireTime == 1 {
			futureTime := time.Now().AddDate(1, 0, 0).Unix()
			fixtureOrgPskModel.ExpireTime = &futureTime
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgPskModel,
				},
			},
		}
	}

	resourceType := "org_psk"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_psk.OrgPskResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate Terraform configuration using automated HCL generation
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				configStr := Render("org_psk", tName, string(f.Bytes()))

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

func (o *OrgPskModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Check fields in struct order
	// 1. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 2. Email (optional)
	if o.Email != nil {
		checks.append(t, "TestCheckResourceAttr", "email", *o.Email)
	}

	// 3. ExpireTime (optional)
	if o.ExpireTime != nil {
		checks.append(t, "TestCheckResourceAttr", "expire_time", fmt.Sprintf("%d", *o.ExpireTime))
	}

	// 4. ExpiryNotificationTime (optional)
	if o.ExpiryNotificationTime != nil {
		checks.append(t, "TestCheckResourceAttr", "expiry_notification_time", fmt.Sprintf("%d", *o.ExpiryNotificationTime))
	}

	// 5. Mac (optional)
	if o.Mac != nil {
		checks.append(t, "TestCheckResourceAttr", "mac", *o.Mac)
	}

	// 6. Macs (optional array)
	if len(o.Macs) > 0 {
		checks.append(t, "TestCheckResourceAttr", "macs.#", fmt.Sprintf("%d", len(o.Macs)))
		for i, mac := range o.Macs {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("macs.%d", i), mac)
		}
	}

	// 7. MaxUsage (optional)
	if o.MaxUsage != nil {
		checks.append(t, "TestCheckResourceAttr", "max_usage", fmt.Sprintf("%d", *o.MaxUsage))
	}

	// 8. Name (required)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)

	// 9. Note (optional)
	if o.Note != nil {
		checks.append(t, "TestCheckResourceAttr", "note", *o.Note)
	}

	// 10. NotifyExpiry (optional)
	if o.NotifyExpiry != nil {
		checks.append(t, "TestCheckResourceAttr", "notify_expiry", fmt.Sprintf("%t", *o.NotifyExpiry))
	}

	// 11. NotifyOnCreateOrEdit (optional)
	if o.NotifyOnCreateOrEdit != nil {
		checks.append(t, "TestCheckResourceAttr", "notify_on_create_or_edit", fmt.Sprintf("%t", *o.NotifyOnCreateOrEdit))
	}

	// 12. OldPassphrase (optional, sensitive - use TestCheckResourceAttrSet)
	if o.OldPassphrase != nil {
		checks.append(t, "TestCheckResourceAttrSet", "old_passphrase")
	}

	// 13. OrgId (required)
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)

	// 14. Passphrase (required)
	checks.append(t, "TestCheckResourceAttr", "passphrase", o.Passphrase)

	// 15. Role (optional)
	if o.Role != nil {
		checks.append(t, "TestCheckResourceAttr", "role", *o.Role)
	}

	// 16. Ssid (required)
	checks.append(t, "TestCheckResourceAttr", "ssid", o.Ssid)

	// 17. Usage (optional)
	if o.Usage != nil {
		checks.append(t, "TestCheckResourceAttr", "usage", *o.Usage)
	}

	// 18. VlanId (optional)
	if o.VlanId != nil {
		checks.append(t, "TestCheckResourceAttr", "vlan_id", *o.VlanId)
	}

	return checks
}
