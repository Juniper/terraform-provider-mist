// WIP
package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_webhook"
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

	resourceType := "org_webhook"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_org_webhook.OrgWebhookResourceSchema(t.Context()).Attributes)
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

func (s *OrgWebhookModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	// Check fields in struct order
	// 1. Enabled
	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}

	// 2. Headers
	if len(s.Headers) > 0 {
		for key, value := range s.Headers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("headers.%s", key), value)
		}
	}

	// 3. Id (computed-only)
	checks.append(t, "TestCheckResourceAttrSet", "id")

	// 4. Name
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)

	// 5. Oauth2ClientId
	if s.Oauth2ClientId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_id", *s.Oauth2ClientId)
	}

	// 6. Oauth2ClientSecret
	if s.Oauth2ClientSecret != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_secret", *s.Oauth2ClientSecret)
	}

	// 7. Oauth2GrantType
	if s.Oauth2GrantType != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_grant_type", *s.Oauth2GrantType)
	}

	// 8. Oauth2Password
	if s.Oauth2Password != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_password", *s.Oauth2Password)
	}

	// 9. Oauth2Scopes
	if len(s.Oauth2Scopes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "oauth2_scopes.#", fmt.Sprintf("%d", len(s.Oauth2Scopes)))
		for i, scope := range s.Oauth2Scopes {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("oauth2_scopes.%d", i), scope)
		}
	}

	// 10. Oauth2TokenUrl
	if s.Oauth2TokenUrl != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_token_url", *s.Oauth2TokenUrl)
	}

	// 11. Oauth2Username
	if s.Oauth2Username != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_username", *s.Oauth2Username)
	}

	// 12. OrgId (computed)
	checks.append(t, "TestCheckResourceAttrSet", "org_id")

	// 13. Secret
	if s.Secret != nil {
		checks.append(t, "TestCheckResourceAttr", "secret", *s.Secret)
	}

	// 14. SingleEventPerMessage
	if s.SingleEventPerMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "single_event_per_message", fmt.Sprintf("%t", *s.SingleEventPerMessage))
	}

	// 15. SplunkToken
	if s.SplunkToken != nil {
		checks.append(t, "TestCheckResourceAttr", "splunk_token", *s.SplunkToken)
	}

	// 16. Topics
	checks.append(t, "TestCheckResourceAttr", "topics.#", fmt.Sprintf("%d", len(s.Topics)))
	for i, topic := range s.Topics {
		checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("topics.%d", i), topic)
	}

	// 17. Type
	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
	}

	// 18. Url
	checks.append(t, "TestCheckResourceAttr", "url", s.Url)

	// 19. VerifyCert
	if s.VerifyCert != nil {
		checks.append(t, "TestCheckResourceAttr", "verify_cert", fmt.Sprintf("%t", *s.VerifyCert))
	}

	return checks
}
