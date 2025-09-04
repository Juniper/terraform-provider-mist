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

func TestOrgWlanPortalTemplateModel(t *testing.T) {
	type testStep struct {
		config OrgWlanPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	// Helper values for boolean pointers
	boolTrue := true

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: OrgWlanPortalTemplateModel{
						OrgId: GetTestOrgId(),
						PortalTemplate: PortalTemplateValue{
							Company: &boolTrue,
							Email:   &boolTrue,
						},
					},
				},
			},
		},
	}

	// Load fixture data following the checklist pattern
	b, err := os.ReadFile("fixtures/org_wlan_portal_template_resource/org_wlan_portal_template_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {
		fixtureOrgWlanPortalTemplateModel := OrgWlanPortalTemplateModel{}
		err = hcl.Decode(&fixtureOrgWlanPortalTemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		fixtureOrgWlanPortalTemplateModel.OrgId = GetTestOrgId()

		// Create a test image file for the logo field (following TestOrgWlanPortalImageModel pattern)
		if fixtureOrgWlanPortalTemplateModel.PortalTemplate.Logo != nil {
			testImagePath := CreateTestPNGFile(t)
			fixtureOrgWlanPortalTemplateModel.PortalTemplate.Logo = &testImagePath
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: fixtureOrgWlanPortalTemplateModel,
				},
			},
		}
	}

	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {
			resourceType := "org_wlan_portal_template"

			// Create single-step tests with combined config (WLAN template + WLAN + portal template)
			// Since portal templates require a WLAN, and WLANs require a template,
			// we create all three in the same config but focus our checks on the portal template
			steps := make([]resource.TestStep, len(tCase.steps))

			for i, step := range tCase.steps {
				// Generate combined config: WLAN template + WLAN
				combinedConfig, wlanRef := GetOrgWlanBaseConfig(step.config.OrgId)

				// Generate the HCL configuration for the portal template
				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&step.config, f.Body())
				// Add the wlan_id attribute to the body before rendering
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig = combinedConfig + "\n\n" + Render("org_wlan_portal_template", tName, string(f.Bytes()))

				// Focus checks on the portal template resource (WLAN template and WLAN are prerequisites)
				checks := step.config.testChecks(t, resourceType, tName)

				steps[i] = resource.TestStep{
					Config: combinedConfig,
					Check:  resource.ComposeAggregateTestCheckFunc(checks.checks...),
				}

				// Log configuration and checks for debugging  
				t.Logf("\n// ------ begin config for test case %s step %d ------\n%s\n// -------- end config for test case %s step %d ------\n", tName, i+1, combinedConfig, tName, i+1)
				t.Logf("\n// ------ begin checks for test case %s step %d ------\n%s\n// -------- end checks for test case %s step %d ------\n", tName, i+1, checks.string(), tName, i+1)
			}

			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps:                    steps,
			})

		})
	}
}

func (s *OrgWlanPortalTemplateModel) testChecks(t testing.TB, rType, tName string) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	
	// Check basic resource attributes
	checks.append(t, "TestCheckResourceAttr", "org_id", s.OrgId)
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")
	
	// Check portal template presence
	checks.append(t, "TestCheckResourceAttrSet", "portal_template.%")
	
	// Check boolean fields
	if s.PortalTemplate.Company != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company", fmt.Sprintf("%t", *s.PortalTemplate.Company))
	}
	if s.PortalTemplate.Email != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email", fmt.Sprintf("%t", *s.PortalTemplate.Email))
	}
	if s.PortalTemplate.Field1 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1", fmt.Sprintf("%t", *s.PortalTemplate.Field1))
	}
	if s.PortalTemplate.Field1required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1required", fmt.Sprintf("%t", *s.PortalTemplate.Field1required))
	}
	if s.PortalTemplate.Field2 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2", fmt.Sprintf("%t", *s.PortalTemplate.Field2))
	}
	if s.PortalTemplate.Field2required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2required", fmt.Sprintf("%t", *s.PortalTemplate.Field2required))
	}
	if s.PortalTemplate.Field3 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field3", fmt.Sprintf("%t", *s.PortalTemplate.Field3))
	}
	if s.PortalTemplate.Field4 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field4", fmt.Sprintf("%t", *s.PortalTemplate.Field4))
	}
	if s.PortalTemplate.MultiAuth != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.multi_auth", fmt.Sprintf("%t", *s.PortalTemplate.MultiAuth))
	}
	if s.PortalTemplate.Name != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.name", fmt.Sprintf("%t", *s.PortalTemplate.Name))
	}
	if s.PortalTemplate.OptOutDefault != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.opt_out_default", fmt.Sprintf("%t", *s.PortalTemplate.OptOutDefault))
	}
	if s.PortalTemplate.Optout != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.optout", fmt.Sprintf("%t", *s.PortalTemplate.Optout))
	}
	if s.PortalTemplate.PoweredBy != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.powered_by", fmt.Sprintf("%t", *s.PortalTemplate.PoweredBy))
	}
	if s.PortalTemplate.Privacy != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.privacy", fmt.Sprintf("%t", *s.PortalTemplate.Privacy))
	}
	if s.PortalTemplate.ResponsiveLayout != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.responsive_layout", fmt.Sprintf("%t", *s.PortalTemplate.ResponsiveLayout))
	}
	if s.PortalTemplate.SmsIsTwilio != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_is_twilio", fmt.Sprintf("%t", *s.PortalTemplate.SmsIsTwilio))
	}
	if s.PortalTemplate.Tos != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.tos", fmt.Sprintf("%t", *s.PortalTemplate.Tos))
	}
	
	// Check string fields
	if s.PortalTemplate.Alignment != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.alignment", *s.PortalTemplate.Alignment)
	}
	if s.PortalTemplate.AuthButtonAmazon != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_amazon", *s.PortalTemplate.AuthButtonAmazon)
	}
	if s.PortalTemplate.AuthButtonAzure != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_azure", *s.PortalTemplate.AuthButtonAzure)
	}
	if s.PortalTemplate.AuthButtonEmail != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_email", *s.PortalTemplate.AuthButtonEmail)
	}
	if s.PortalTemplate.AuthButtonFacebook != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_facebook", *s.PortalTemplate.AuthButtonFacebook)
	}
	if s.PortalTemplate.AuthButtonGoogle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_google", *s.PortalTemplate.AuthButtonGoogle)
	}
	if s.PortalTemplate.AuthButtonMicrosoft != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_microsoft", *s.PortalTemplate.AuthButtonMicrosoft)
	}
	if s.PortalTemplate.AuthButtonPassphrase != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_passphrase", *s.PortalTemplate.AuthButtonPassphrase)
	}
	if s.PortalTemplate.AuthButtonSms != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_sms", *s.PortalTemplate.AuthButtonSms)
	}
	if s.PortalTemplate.AuthButtonSponsor != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_button_sponsor", *s.PortalTemplate.AuthButtonSponsor)
	}
	if s.PortalTemplate.AuthLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.auth_label", *s.PortalTemplate.AuthLabel)
	}
	if s.PortalTemplate.BackLink != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.back_link", *s.PortalTemplate.BackLink)
	}
	if s.PortalTemplate.Color != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.color", *s.PortalTemplate.Color)
	}
	if s.PortalTemplate.ColorDark != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.color_dark", *s.PortalTemplate.ColorDark)
	}
	if s.PortalTemplate.ColorLight != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.color_light", *s.PortalTemplate.ColorLight)
	}
	if s.PortalTemplate.CompanyError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company_error", *s.PortalTemplate.CompanyError)
	}
	if s.PortalTemplate.CompanyLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company_label", *s.PortalTemplate.CompanyLabel)
	}
	if s.PortalTemplate.EmailAccessDomainError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_access_domain_error", *s.PortalTemplate.EmailAccessDomainError)
	}
	if s.PortalTemplate.EmailCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_cancel", *s.PortalTemplate.EmailCancel)
	}
	if s.PortalTemplate.EmailCodeCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_cancel", *s.PortalTemplate.EmailCodeCancel)
	}
	if s.PortalTemplate.EmailCodeError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_error", *s.PortalTemplate.EmailCodeError)
	}
	if s.PortalTemplate.EmailCodeFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_field_label", *s.PortalTemplate.EmailCodeFieldLabel)
	}
	if s.PortalTemplate.EmailCodeMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_message", *s.PortalTemplate.EmailCodeMessage)
	}
	if s.PortalTemplate.EmailCodeSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_submit", *s.PortalTemplate.EmailCodeSubmit)
	}
	if s.PortalTemplate.EmailCodeTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_code_title", *s.PortalTemplate.EmailCodeTitle)
	}
	if s.PortalTemplate.EmailError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_error", *s.PortalTemplate.EmailError)
	}
	if s.PortalTemplate.EmailFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_field_label", *s.PortalTemplate.EmailFieldLabel)
	}
	if s.PortalTemplate.EmailLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_label", *s.PortalTemplate.EmailLabel)
	}
	if s.PortalTemplate.EmailMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_message", *s.PortalTemplate.EmailMessage)
	}
	if s.PortalTemplate.EmailSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_submit", *s.PortalTemplate.EmailSubmit)
	}
	if s.PortalTemplate.EmailTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email_title", *s.PortalTemplate.EmailTitle)
	}
	if s.PortalTemplate.Field1error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1error", *s.PortalTemplate.Field1error)
	}
	if s.PortalTemplate.Field1label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1label", *s.PortalTemplate.Field1label)
	}
	if s.PortalTemplate.Field2error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2error", *s.PortalTemplate.Field2error)
	}
	if s.PortalTemplate.Field2label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2label", *s.PortalTemplate.Field2label)
	}
	if s.PortalTemplate.Logo != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.logo", *s.PortalTemplate.Logo)
	}
	if s.PortalTemplate.Message != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.message", *s.PortalTemplate.Message)
	}
	if s.PortalTemplate.NameError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.name_error", *s.PortalTemplate.NameError)
	}
	if s.PortalTemplate.NameLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.name_label", *s.PortalTemplate.NameLabel)
	}
	if s.PortalTemplate.OptoutLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.optout_label", *s.PortalTemplate.OptoutLabel)
	}
	if s.PortalTemplate.PageTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.page_title", *s.PortalTemplate.PageTitle)
	}
	if s.PortalTemplate.PassphraseCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_cancel", *s.PortalTemplate.PassphraseCancel)
	}
	if s.PortalTemplate.PassphraseError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_error", *s.PortalTemplate.PassphraseError)
	}
	if s.PortalTemplate.PassphraseLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_label", *s.PortalTemplate.PassphraseLabel)
	}
	if s.PortalTemplate.PassphraseMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_message", *s.PortalTemplate.PassphraseMessage)
	}
	if s.PortalTemplate.PassphraseSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_submit", *s.PortalTemplate.PassphraseSubmit)
	}
	if s.PortalTemplate.PassphraseTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.passphrase_title", *s.PortalTemplate.PassphraseTitle)
	}
	if s.PortalTemplate.PrivacyPolicyAcceptLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.privacy_policy_accept_label", *s.PortalTemplate.PrivacyPolicyAcceptLabel)
	}
	if s.PortalTemplate.PrivacyPolicyError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.privacy_policy_error", *s.PortalTemplate.PrivacyPolicyError)
	}
	if s.PortalTemplate.PrivacyPolicyLink != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.privacy_policy_link", *s.PortalTemplate.PrivacyPolicyLink)
	}
	if s.PortalTemplate.PrivacyPolicyText != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.privacy_policy_text", *s.PortalTemplate.PrivacyPolicyText)
	}
	if s.PortalTemplate.RequiredFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.required_field_label", *s.PortalTemplate.RequiredFieldLabel)
	}
	if s.PortalTemplate.SignInLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sign_in_label", *s.PortalTemplate.SignInLabel)
	}
	if s.PortalTemplate.TosAcceptLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.tos_accept_label", *s.PortalTemplate.TosAcceptLabel)
	}
	if s.PortalTemplate.TosError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.tos_error", *s.PortalTemplate.TosError)
	}
	if s.PortalTemplate.TosLink != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.tos_link", *s.PortalTemplate.TosLink)
	}
	if s.PortalTemplate.TosText != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.tos_text", *s.PortalTemplate.TosText)
	}
	
	// Check int64 fields
	if s.PortalTemplate.SmsValidityDuration != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_validity_duration", fmt.Sprintf("%d", *s.PortalTemplate.SmsValidityDuration))
	}

	return checks
}
