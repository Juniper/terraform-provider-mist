package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgAlarmtemplate(t *testing.T) {
	type testStep struct {
		config OrgAlarmtemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgAlarmtemplateModel{
						Name:  "test_alarm_template",
						OrgId: GetTestOrgId(),
						Delivery: OrgAlarmtemplateDeliveryValue{
							Enabled: false,
						},
						Rules: map[string]OrgAlarmtemplateRulesValue{
							"test_rule": {
								Enabled: false,
							},
						},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_alarmtemplate_resource/org_alarmtemplate_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		FixtureOrgAlarmtemplateModel := OrgAlarmtemplateModel{}

		err = hcl.Decode(&FixtureOrgAlarmtemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgAlarmtemplateModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_alarmtemplate"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config
				config.OrgId = GetTestOrgId()

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

func (s *OrgAlarmtemplateModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)
	// Required parameters
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttrSet", "delivery.%")
	checks.append(t, "TestCheckResourceAttr", "delivery.enabled", fmt.Sprintf("%t", s.Delivery.Enabled))

	// DeliveryValue optional fields
	if len(s.Delivery.AdditionalEmails) > 0 {
		checks.append(t, "TestCheckResourceAttr", "delivery.additional_emails.#", fmt.Sprintf("%d", len(s.Delivery.AdditionalEmails)))
	}
	if s.Delivery.ToOrgAdmins != nil {
		checks.append(t, "TestCheckResourceAttr", "delivery.to_org_admins", fmt.Sprintf("%t", *s.Delivery.ToOrgAdmins))
	}
	if s.Delivery.ToSiteAdmins != nil {
		checks.append(t, "TestCheckResourceAttr", "delivery.to_site_admins", fmt.Sprintf("%t", *s.Delivery.ToSiteAdmins))
	}

	// Rules map
	if len(s.Rules) > 0 {
		checks.append(t, "TestCheckResourceAttr", "rules.%", fmt.Sprintf("%d", len(s.Rules)))
		for key, rule := range s.Rules {
			attrPrefix := fmt.Sprintf("rules.%s", key)
			checks.append(t, "TestCheckResourceAttrSet", attrPrefix+".enabled")
			if rule.Delivery != nil {
				if len(rule.Delivery.AdditionalEmails) > 0 {
					checks.append(t, "TestCheckResourceAttr", attrPrefix+".delivery.additional_emails.#", fmt.Sprintf("%d", len(rule.Delivery.AdditionalEmails)))
				}
				if rule.Delivery.ToOrgAdmins != nil {
					checks.append(t, "TestCheckResourceAttr", attrPrefix+".delivery.to_org_admins", fmt.Sprintf("%t", *rule.Delivery.ToOrgAdmins))
				}
				if rule.Delivery.ToSiteAdmins != nil {
					checks.append(t, "TestCheckResourceAttr", attrPrefix+".delivery.to_site_admins", fmt.Sprintf("%t", *rule.Delivery.ToSiteAdmins))
				}
			}
		}
	}

	return checks
}
