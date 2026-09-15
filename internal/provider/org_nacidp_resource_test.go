package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nacidp"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgNacidpModel(t *testing.T) {
	type testStep struct {
		config OrgNacidpModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgNacidpModel{
						OrgId:            GetTestOrgId(),
						Name:             "test-nacidp",
						IdpType:          "ldap",
						LdapBaseDn:       stringPtr("dc=example,dc=com"),
						LdapBindDn:       stringPtr("cn=admin,dc=example,dc=com"),
						LdapBindPassword: stringPtr("password123"),
						LdapType:         stringPtr("azure"),
						LdapServerHosts:  []string{"ldap.example.com"},
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_nacidp_resource/org_nacidp_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	fixtures := strings.Split(string(b), "␞")

	for i, fixture := range fixtures {
		var FixtureOrgNacidpModel OrgNacidpModel
		err = hcl.Decode(&FixtureOrgNacidpModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgNacidpModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgNacidpModel,
				},
			},
		}
	}

	resourceType := "org_nacidp"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_nacidp.OrgNacidpResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				combinedConfig := Render(resourceType, tName, string(f.Bytes()))

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

func (o *OrgNacidpModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)
	appendReflectChecks(t, &checks, o)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	return checks
}
