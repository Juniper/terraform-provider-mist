package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sso"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestOrgSsoModel(t *testing.T) {
	type testStep struct {
		config OrgSsoModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgSsoModel{
						OrgId:       GetTestOrgId(),
						Name:        "test-sso",
						IdpCert:     "-----BEGIN CERTIFICATE-----\nMIIBkTCB+wIJALJ8UUKmgH1GMA0GCSqGSIb3DQEBCwUAMBQxEjAQBgNVBAMMCXRlc3QtY2VydDAeFw0yNDEyMjcwMDAwMDBaFw0yNTEyMjcwMDAwMDBaMBQxEjAQBgNVBAMMCXRlc3QtY2VydDBcMA0GCSqGSIb3DQEBAQUAA0sAMEgCQQCxUC6+OeSgM1FhOdKqA5C1XQfFdKK0C8JxUQKHjOKE8Q1j8I+FHFOdKGY5TKZrIvOLMbOeXJGF7Wl5xD0dVhZdAgMBAAEwDQYJKoZIhvcNAQELBQADQQA3F8+8MzE5E5GHj5E5TQ==\n-----END CERTIFICATE-----",
						IdpSignAlgo: "sha256",
						IdpSsoUrl:   "https://example.com/sso",
						Issuer:      "https://example.com/issuer",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/org_sso_resource/org_sso_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		var FixtureOrgSsoModel OrgSsoModel
		err = hcl.Decode(&FixtureOrgSsoModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		FixtureOrgSsoModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_casepaskdvn_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureOrgSsoModel,
				},
			},
		}
	}

	resourceType := "org_sso"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_sso.OrgSsoResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				configStr := Render(resourceType, tName, string(f.Bytes()))

				checks := config.testChecks(t, resourceType, tName, tracker)
				chkLog := checks.string()
				stepName := fmt.Sprintf("test case %s step %d", tName, i+1)

				t.Logf("\n// ------ begin config for %s ------\n%s// -------- end config for %s ------\n\n", stepName, configStr, stepName)
				t.Logf("\n// ------ begin checks for %s ------\n%s// -------- end config for %s ------\n\n", stepName, chkLog, stepName)

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

func (o *OrgSsoModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType)+"."+tName, tracker)

	// Check required fields
	checks.append(t, "TestCheckResourceAttr", "org_id", o.OrgId)
	checks.append(t, "TestCheckResourceAttr", "name", o.Name)
	checks.append(t, "TestCheckResourceAttr", "idp_cert", o.IdpCert)
	checks.append(t, "TestCheckResourceAttr", "idp_sign_algo", o.IdpSignAlgo)
	checks.append(t, "TestCheckResourceAttr", "idp_sso_url", o.IdpSsoUrl)
	checks.append(t, "TestCheckResourceAttr", "issuer", o.Issuer)

	// Check computed-only fields (verify they exist)
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "domain")

	// Check optional string fields
	if o.CustomLogoutUrl != nil {
		checks.append(t, "TestCheckResourceAttr", "custom_logout_url", *o.CustomLogoutUrl)
	}
	if o.DefaultRole != nil {
		checks.append(t, "TestCheckResourceAttr", "default_role", *o.DefaultRole)
	}
	if o.NameidFormat != nil {
		checks.append(t, "TestCheckResourceAttr", "nameid_format", *o.NameidFormat)
	}
	if o.RoleAttrExtraction != nil {
		checks.append(t, "TestCheckResourceAttr", "role_attr_extraction", *o.RoleAttrExtraction)
	}
	if o.RoleAttrFrom != nil {
		checks.append(t, "TestCheckResourceAttr", "role_attr_from", *o.RoleAttrFrom)
	}

	// Check optional boolean fields
	if o.IgnoreUnmatchedRoles != nil {
		checks.append(t, "TestCheckResourceAttr", "ignore_unmatched_roles", fmt.Sprintf("%t", *o.IgnoreUnmatchedRoles))
	}

	return checks
}
