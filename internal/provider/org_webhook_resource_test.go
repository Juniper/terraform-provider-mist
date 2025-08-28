// WIP
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

func TestOrgWebhookModel(t *testing.T) {
	type testStep struct {
		config OrgWebhookModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWebhookModel{
						Name:  "webhook_test",
						Url:   "http://api.mistsys.com",
						OrgId: GetTestOrgId(),
						Topics: []string{
							"alarms",
						},
					},
				},
			},
		},
	}

	fixtures, err := os.ReadFile("fixtures/org_webhook_resource/org_webhook_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	for i, fixture := range strings.Split(string(fixtures), "␞") {
		fixtureOrgWebhookModel := OrgWebhookModel{}
		err = hcl.Decode(&fixtureOrgWebhookModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWebhookModel.OrgId = GetTestOrgId()

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWebhookModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_webhook"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				config := step.config

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

func (s *OrgWebhookModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)

	// Always check required fields
	checks.append(t, "TestCheckResourceAttrSet", "org_id")
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttr", "url", s.Url)
	checks.append(t, "TestCheckResourceAttr", "topics.#", fmt.Sprintf("%d", len(s.Topics)))

	// Check individual topics
	for i, topic := range s.Topics {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("topics.%d", i), topic)
	}

	// Check type if set
	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
	}

	// Check enabled if explicitly set
	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}

	// Check verify_cert if explicitly set
	if s.VerifyCert != nil {
		checks.append(t, "TestCheckResourceAttr", "verify_cert", fmt.Sprintf("%t", *s.VerifyCert))
	}

	// Check single_event_per_message if explicitly set
	if s.SingleEventPerMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "single_event_per_message", fmt.Sprintf("%t", *s.SingleEventPerMessage))
	}

	// Check secret if set
	if s.Secret != nil {
		checks.append(t, "TestCheckResourceAttr", "secret", *s.Secret)
	}

	// Check OAuth2 fields if set
	if s.Oauth2GrantType != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_grant_type", *s.Oauth2GrantType)
	}
	if s.Oauth2ClientId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_id", *s.Oauth2ClientId)
	}
	if s.Oauth2ClientSecret != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_secret", *s.Oauth2ClientSecret)
	}
	if s.Oauth2Username != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_username", *s.Oauth2Username)
	}
	if s.Oauth2Password != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_password", *s.Oauth2Password)
	}
	if s.Oauth2TokenUrl != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_token_url", *s.Oauth2TokenUrl)
	}
	if len(s.Oauth2Scopes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "oauth2_scopes.#", fmt.Sprintf("%d", len(s.Oauth2Scopes)))
		for i, scope := range s.Oauth2Scopes {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("oauth2_scopes.%d", i), scope)
		}
	}

	// Check Splunk token if set
	if s.SplunkToken != nil {
		checks.append(t, "TestCheckResourceAttr", "splunk_token", *s.SplunkToken)
	}

	// Check headers if set
	if len(s.Headers) > 0 {
		for key, value := range s.Headers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("headers.%s", key), value)
		}
	}

	return checks
}
