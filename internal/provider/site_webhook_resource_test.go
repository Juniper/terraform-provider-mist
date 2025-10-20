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

func TestSiteWebhookModel(t *testing.T) {
	type testStep struct {
		config SiteWebhookModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWebhookModel{
						Name: "test-site-webhook",
						Topics: []string{
							"device-events",
							"alarms",
						},
						Url: "https://example.com/webhook",
					},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_webhook_resource/site_webhook_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {

		var FixtureSiteWebhookModel SiteWebhookModel
		err = hcl.Decode(&FixtureSiteWebhookModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteWebhookModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "site_webhook"

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				siteConfig, siteRef := GetSiteBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				configStr := siteConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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

func (s *SiteWebhookModel) testChecks(t testing.TB, rType, rName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + rName)

	// Required attributes
	checks.append(t, "TestCheckResourceAttr", "name", s.Name)
	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttr", "url", s.Url)

	// Check topics list
	if len(s.Topics) > 0 {
		checks.append(t, "TestCheckResourceAttr", "topics.#", fmt.Sprintf("%d", len(s.Topics)))
		for i, topic := range s.Topics {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("topics.%d", i), topic)
		}
	}

	// Optional attributes with conditional checks
	if len(s.AssetfilterIds) > 0 {
		checks.append(t, "TestCheckResourceAttr", "assetfilter_ids.#", fmt.Sprintf("%d", len(s.AssetfilterIds)))
		for i, id := range s.AssetfilterIds {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("assetfilter_ids.%d", i), id)
		}
	}

	if s.Enabled != nil {
		checks.append(t, "TestCheckResourceAttr", "enabled", fmt.Sprintf("%t", *s.Enabled))
	}

	if len(s.Headers) > 0 {
		for key, value := range s.Headers {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("headers.%s", key), value)
		}
	}

	if s.Oauth2ClientId != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_id", *s.Oauth2ClientId)
	}

	if s.Oauth2ClientSecret != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_client_secret", *s.Oauth2ClientSecret)
	}

	if s.Oauth2GrantType != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_grant_type", *s.Oauth2GrantType)
	}

	if s.Oauth2Password != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_password", *s.Oauth2Password)
	}

	if len(s.Oauth2Scopes) > 0 {
		checks.append(t, "TestCheckResourceAttr", "oauth2_scopes.#", fmt.Sprintf("%d", len(s.Oauth2Scopes)))
		for i, scope := range s.Oauth2Scopes {
			checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("oauth2_scopes.%d", i), scope)
		}
	}

	if s.Oauth2TokenUrl != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_token_url", *s.Oauth2TokenUrl)
	}

	if s.Oauth2Username != nil {
		checks.append(t, "TestCheckResourceAttr", "oauth2_username", *s.Oauth2Username)
	}

	if s.Secret != nil {
		checks.append(t, "TestCheckResourceAttr", "secret", *s.Secret)
	}

	if s.SingleEventPerMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "single_event_per_message", fmt.Sprintf("%t", *s.SingleEventPerMessage))
	}

	if s.SplunkToken != nil {
		checks.append(t, "TestCheckResourceAttr", "splunk_token", *s.SplunkToken)
	}

	if s.Type != nil {
		checks.append(t, "TestCheckResourceAttr", "type", *s.Type)
	}

	if s.VerifyCert != nil {
		checks.append(t, "TestCheckResourceAttr", "verify_cert", fmt.Sprintf("%t", *s.VerifyCert))
	}

	// Computed attributes that should be set
	checks.append(t, "TestCheckResourceAttrSet", "id")
	checks.append(t, "TestCheckResourceAttrSet", "org_id")

	return checks
}
