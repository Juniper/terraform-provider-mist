package provider

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Juniper/terraform-provider-mist/internal/provider/validators"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wlan_portal_template"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSiteWlanPortalTemplateModel(t *testing.T) {
	type testStep struct {
		config SiteWlanPortalTemplateModel
	}

	type testCase struct {
		steps []testStep
	}

	testCases := map[string]testCase{
		"simple_case": {
			steps: []testStep{
				{
					config: SiteWlanPortalTemplateModel{},
				},
			},
		},
	}

	b, err := os.ReadFile("fixtures/site_wlan_portal_template_resource/site_wlan_portal_template_config.tf")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	fixtures := strings.Split(str, "␞")

	for i, fixture := range fixtures {

		var FixtureSiteWlanPortalTemplateModel SiteWlanPortalTemplateModel
		err = hcl.Decode(&FixtureSiteWlanPortalTemplateModel, fixture)
		if err != nil {
			fmt.Printf("error decoding hcl: %s\n", err)
		}

		testCases[fmt.Sprintf("fixture_case_%d", i)] = testCase{
			steps: []testStep{
				{
					config: FixtureSiteWlanPortalTemplateModel,
				},
			},
		}
	}

	resourceType := "site_wlan_portal_template"
	tracker := validators.FieldCoverageTrackerWithSchema(resourceType, resource_site_wlan_portal_template.SiteWlanPortalTemplateResourceSchema(t.Context()).Attributes)
	for tName, tCase := range testCases {
		t.Run(tName, func(t *testing.T) {

			steps := make([]resource.TestStep, len(tCase.steps))
			for i, step := range tCase.steps {
				// Generate combined config: Site + WLAN
				wlanConfig, siteRef, wlanRef := GetSiteWlanBaseConfig(GetTestOrgId())
				config := step.config

				f := hclwrite.NewEmptyFile()
				gohcl.EncodeIntoBody(&config, f.Body())
				// Add the site_id and wlan_id attributes to the body before rendering
				f.Body().SetAttributeRaw("site_id", hclwrite.TokensForIdentifier(siteRef))
				f.Body().SetAttributeRaw("wlan_id", hclwrite.TokensForIdentifier(wlanRef))
				combinedConfig := wlanConfig + "\n\n" + Render(resourceType, tName, string(f.Bytes()))

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
	tracker.FieldCoverageReport(t)
}

func (s *SiteWlanPortalTemplateModel) testChecks(t testing.TB, rType, tName string, tracker *validators.FieldCoverageTracker) testChecks {
	checks := newTestChecks(PrefixProviderName(rType) + "." + tName)
	checks.SetTracker(tracker)

	checks.append(t, "TestCheckResourceAttrSet", "site_id")
	checks.append(t, "TestCheckResourceAttrSet", "wlan_id")

	// Portal template attributes
	if s.PortalTemplate.AccessCodeAlternateEmail != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.access_code_alternate_email", *s.PortalTemplate.AccessCodeAlternateEmail)
	}
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
	if s.PortalTemplate.Company != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company", fmt.Sprintf("%t", *s.PortalTemplate.Company))
	}
	if s.PortalTemplate.CompanyError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company_error", *s.PortalTemplate.CompanyError)
	}
	if s.PortalTemplate.CompanyLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.company_label", *s.PortalTemplate.CompanyLabel)
	}
	if s.PortalTemplate.Email != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.email", fmt.Sprintf("%t", *s.PortalTemplate.Email))
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
	if s.PortalTemplate.Field1 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1", fmt.Sprintf("%t", *s.PortalTemplate.Field1))
	}
	if s.PortalTemplate.Field1error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1error", *s.PortalTemplate.Field1error)
	}
	if s.PortalTemplate.Field1label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1label", *s.PortalTemplate.Field1label)
	}
	if s.PortalTemplate.Field1required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field1required", fmt.Sprintf("%t", *s.PortalTemplate.Field1required))
	}
	if s.PortalTemplate.Field2 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2", fmt.Sprintf("%t", *s.PortalTemplate.Field2))
	}
	if s.PortalTemplate.Field2error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2error", *s.PortalTemplate.Field2error)
	}
	if s.PortalTemplate.Field2label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2label", *s.PortalTemplate.Field2label)
	}
	if s.PortalTemplate.Field2required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field2required", fmt.Sprintf("%t", *s.PortalTemplate.Field2required))
	}
	if s.PortalTemplate.Field3 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field3", fmt.Sprintf("%t", *s.PortalTemplate.Field3))
	}
	if s.PortalTemplate.Field3error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field3error", *s.PortalTemplate.Field3error)
	}
	if s.PortalTemplate.Field3label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field3label", *s.PortalTemplate.Field3label)
	}
	if s.PortalTemplate.Field3required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field3required", fmt.Sprintf("%t", *s.PortalTemplate.Field3required))
	}
	if s.PortalTemplate.Field4 != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field4", fmt.Sprintf("%t", *s.PortalTemplate.Field4))
	}
	if s.PortalTemplate.Field4error != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field4error", *s.PortalTemplate.Field4error)
	}
	if s.PortalTemplate.Field4label != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field4label", *s.PortalTemplate.Field4label)
	}
	if s.PortalTemplate.Field4required != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.field4required", fmt.Sprintf("%t", *s.PortalTemplate.Field4required))
	}
	if s.PortalTemplate.Logo != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.logo", *s.PortalTemplate.Logo)
	}
	if s.PortalTemplate.Message != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.message", *s.PortalTemplate.Message)
	}
	if s.PortalTemplate.Name != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.name", fmt.Sprintf("%t", *s.PortalTemplate.Name))
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
	if s.PortalTemplate.SmsCarrierDefault != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_carrier_default", *s.PortalTemplate.SmsCarrierDefault)
	}
	if s.PortalTemplate.SmsCarrierError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_carrier_error", *s.PortalTemplate.SmsCarrierError)
	}
	if s.PortalTemplate.SmsCarrierFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_carrier_field_label", *s.PortalTemplate.SmsCarrierFieldLabel)
	}
	if s.PortalTemplate.SmsCodeCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_cancel", *s.PortalTemplate.SmsCodeCancel)
	}
	if s.PortalTemplate.SmsCodeError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_error", *s.PortalTemplate.SmsCodeError)
	}
	if s.PortalTemplate.SmsCodeFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_field_label", *s.PortalTemplate.SmsCodeFieldLabel)
	}
	if s.PortalTemplate.SmsCodeMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_message", *s.PortalTemplate.SmsCodeMessage)
	}
	if s.PortalTemplate.SmsCodeSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_submit", *s.PortalTemplate.SmsCodeSubmit)
	}
	if s.PortalTemplate.SmsCodeTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_code_title", *s.PortalTemplate.SmsCodeTitle)
	}
	if s.PortalTemplate.SmsCountryFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_country_field_label", *s.PortalTemplate.SmsCountryFieldLabel)
	}
	if s.PortalTemplate.SmsCountryFormat != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_country_format", *s.PortalTemplate.SmsCountryFormat)
	}
	if s.PortalTemplate.SmsHaveAccessCode != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_have_access_code", *s.PortalTemplate.SmsHaveAccessCode)
	}
	if s.PortalTemplate.SmsIsTwilio != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_is_twilio", fmt.Sprintf("%t", *s.PortalTemplate.SmsIsTwilio))
	}
	if s.PortalTemplate.SmsMessageFormat != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_message_format", *s.PortalTemplate.SmsMessageFormat)
	}
	if s.PortalTemplate.SmsNumberCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_cancel", *s.PortalTemplate.SmsNumberCancel)
	}
	if s.PortalTemplate.SmsNumberError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_error", *s.PortalTemplate.SmsNumberError)
	}
	if s.PortalTemplate.SmsNumberFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_field_label", *s.PortalTemplate.SmsNumberFieldLabel)
	}
	if s.PortalTemplate.SmsNumberFormat != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_format", *s.PortalTemplate.SmsNumberFormat)
	}
	if s.PortalTemplate.SmsNumberMessage != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_message", *s.PortalTemplate.SmsNumberMessage)
	}
	if s.PortalTemplate.SmsNumberSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_submit", *s.PortalTemplate.SmsNumberSubmit)
	}
	if s.PortalTemplate.SmsNumberTitle != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_number_title", *s.PortalTemplate.SmsNumberTitle)
	}
	if s.PortalTemplate.SmsUsernameFormat != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sms_username_format", *s.PortalTemplate.SmsUsernameFormat)
	}
	if s.PortalTemplate.SponsorBackLink != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_back_link", *s.PortalTemplate.SponsorBackLink)
	}
	if s.PortalTemplate.SponsorCancel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_cancel", *s.PortalTemplate.SponsorCancel)
	}
	if s.PortalTemplate.SponsorEmail != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_email", *s.PortalTemplate.SponsorEmail)
	}
	if s.PortalTemplate.SponsorEmailError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_email_error", *s.PortalTemplate.SponsorEmailError)
	}
	if s.PortalTemplate.SponsorInfoApproved != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_info_approved", *s.PortalTemplate.SponsorInfoApproved)
	}
	if s.PortalTemplate.SponsorInfoDenied != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_info_denied", *s.PortalTemplate.SponsorInfoDenied)
	}
	if s.PortalTemplate.SponsorInfoPending != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_info_pending", *s.PortalTemplate.SponsorInfoPending)
	}
	if s.PortalTemplate.SponsorName != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_name", *s.PortalTemplate.SponsorName)
	}
	if s.PortalTemplate.SponsorNameError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_name_error", *s.PortalTemplate.SponsorNameError)
	}
	if s.PortalTemplate.SponsorNotePending != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_note_pending", *s.PortalTemplate.SponsorNotePending)
	}
	if s.PortalTemplate.SponsorRequestAccess != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_request_access", *s.PortalTemplate.SponsorRequestAccess)
	}
	if s.PortalTemplate.SponsorStatusApproved != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_status_approved", *s.PortalTemplate.SponsorStatusApproved)
	}
	if s.PortalTemplate.SponsorStatusDenied != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_status_denied", *s.PortalTemplate.SponsorStatusDenied)
	}
	if s.PortalTemplate.SponsorStatusPending != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_status_pending", *s.PortalTemplate.SponsorStatusPending)
	}
	if s.PortalTemplate.SponsorSubmit != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsor_submit", *s.PortalTemplate.SponsorSubmit)
	}
	if s.PortalTemplate.SponsorsError != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsors_error", *s.PortalTemplate.SponsorsError)
	}
	if s.PortalTemplate.SponsorsFieldLabel != nil {
		checks.append(t, "TestCheckResourceAttr", "portal_template.sponsors_field_label", *s.PortalTemplate.SponsorsFieldLabel)
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

	// Handle locales map
	if s.PortalTemplate.Locales != nil {
		for locale, localeValue := range s.PortalTemplate.Locales {
			if localeValue.AuthButtonAmazon != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_amazon", locale), *localeValue.AuthButtonAmazon)
			}
			if localeValue.AuthButtonAzure != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_azure", locale), *localeValue.AuthButtonAzure)
			}
			if localeValue.AuthButtonEmail != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_email", locale), *localeValue.AuthButtonEmail)
			}
			if localeValue.AuthButtonFacebook != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_facebook", locale), *localeValue.AuthButtonFacebook)
			}
			if localeValue.AuthButtonGoogle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_google", locale), *localeValue.AuthButtonGoogle)
			}
			if localeValue.AuthButtonMicrosoft != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_microsoft", locale), *localeValue.AuthButtonMicrosoft)
			}
			if localeValue.AuthButtonPassphrase != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_passphrase", locale), *localeValue.AuthButtonPassphrase)
			}
			if localeValue.AuthButtonSms != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_sms", locale), *localeValue.AuthButtonSms)
			}
			if localeValue.AuthButtonSponsor != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_button_sponsor", locale), *localeValue.AuthButtonSponsor)
			}
			if localeValue.AuthLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.auth_label", locale), *localeValue.AuthLabel)
			}
			if localeValue.BackLink != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.back_link", locale), *localeValue.BackLink)
			}
			if localeValue.CompanyError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.company_error", locale), *localeValue.CompanyError)
			}
			if localeValue.CompanyLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.company_label", locale), *localeValue.CompanyLabel)
			}
			if localeValue.EmailAccessDomainError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_access_domain_error", locale), *localeValue.EmailAccessDomainError)
			}
			if localeValue.EmailCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_cancel", locale), *localeValue.EmailCancel)
			}
			if localeValue.EmailCodeCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_cancel", locale), *localeValue.EmailCodeCancel)
			}
			if localeValue.EmailCodeError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_error", locale), *localeValue.EmailCodeError)
			}
			if localeValue.EmailCodeFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_field_label", locale), *localeValue.EmailCodeFieldLabel)
			}
			if localeValue.EmailCodeMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_message", locale), *localeValue.EmailCodeMessage)
			}
			if localeValue.EmailCodeSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_submit", locale), *localeValue.EmailCodeSubmit)
			}
			if localeValue.EmailCodeTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_code_title", locale), *localeValue.EmailCodeTitle)
			}
			if localeValue.EmailError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_error", locale), *localeValue.EmailError)
			}
			if localeValue.EmailFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_field_label", locale), *localeValue.EmailFieldLabel)
			}
			if localeValue.EmailLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_label", locale), *localeValue.EmailLabel)
			}
			if localeValue.EmailMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_message", locale), *localeValue.EmailMessage)
			}
			if localeValue.EmailSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_submit", locale), *localeValue.EmailSubmit)
			}
			if localeValue.EmailTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.email_title", locale), *localeValue.EmailTitle)
			}
			if localeValue.Field1error != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field1error", locale), *localeValue.Field1error)
			}
			if localeValue.Field1label != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field1label", locale), *localeValue.Field1label)
			}
			if localeValue.Field2error != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field2error", locale), *localeValue.Field2error)
			}
			if localeValue.Field2label != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field2label", locale), *localeValue.Field2label)
			}
			if localeValue.Field3error != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field3error", locale), *localeValue.Field3error)
			}
			if localeValue.Field3label != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field3label", locale), *localeValue.Field3label)
			}
			if localeValue.Field4error != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field4error", locale), *localeValue.Field4error)
			}
			if localeValue.Field4label != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.field4label", locale), *localeValue.Field4label)
			}
			if localeValue.Message != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.message", locale), *localeValue.Message)
			}
			if localeValue.NameError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.name_error", locale), *localeValue.NameError)
			}
			if localeValue.NameLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.name_label", locale), *localeValue.NameLabel)
			}
			if localeValue.OptoutLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.optout_label", locale), *localeValue.OptoutLabel)
			}
			if localeValue.PageTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.page_title", locale), *localeValue.PageTitle)
			}
			if localeValue.PassphraseCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_cancel", locale), *localeValue.PassphraseCancel)
			}
			if localeValue.PassphraseError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_error", locale), *localeValue.PassphraseError)
			}
			if localeValue.PassphraseLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_label", locale), *localeValue.PassphraseLabel)
			}
			if localeValue.PassphraseMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_message", locale), *localeValue.PassphraseMessage)
			}
			if localeValue.PassphraseSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_submit", locale), *localeValue.PassphraseSubmit)
			}
			if localeValue.PassphraseTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.passphrase_title", locale), *localeValue.PassphraseTitle)
			}
			if localeValue.PrivacyPolicyAcceptLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.privacy_policy_accept_label", locale), *localeValue.PrivacyPolicyAcceptLabel)
			}
			if localeValue.PrivacyPolicyError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.privacy_policy_error", locale), *localeValue.PrivacyPolicyError)
			}
			if localeValue.PrivacyPolicyLink != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.privacy_policy_link", locale), *localeValue.PrivacyPolicyLink)
			}
			if localeValue.PrivacyPolicyText != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.privacy_policy_text", locale), *localeValue.PrivacyPolicyText)
			}
			if localeValue.RequiredFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.required_field_label", locale), *localeValue.RequiredFieldLabel)
			}
			if localeValue.SignInLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sign_in_label", locale), *localeValue.SignInLabel)
			}
			if localeValue.SmsCarrierDefault != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_carrier_default", locale), *localeValue.SmsCarrierDefault)
			}
			if localeValue.SmsCarrierError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_carrier_error", locale), *localeValue.SmsCarrierError)
			}
			if localeValue.SmsCarrierFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_carrier_field_label", locale), *localeValue.SmsCarrierFieldLabel)
			}
			if localeValue.SmsCodeCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_cancel", locale), *localeValue.SmsCodeCancel)
			}
			if localeValue.SmsCodeError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_error", locale), *localeValue.SmsCodeError)
			}
			if localeValue.SmsCodeFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_field_label", locale), *localeValue.SmsCodeFieldLabel)
			}
			if localeValue.SmsCodeMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_message", locale), *localeValue.SmsCodeMessage)
			}
			if localeValue.SmsCodeSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_submit", locale), *localeValue.SmsCodeSubmit)
			}
			if localeValue.SmsCodeTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_code_title", locale), *localeValue.SmsCodeTitle)
			}
			if localeValue.SmsCountryFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_country_field_label", locale), *localeValue.SmsCountryFieldLabel)
			}
			if localeValue.SmsCountryFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_country_format", locale), *localeValue.SmsCountryFormat)
			}
			if localeValue.SmsHaveAccessCode != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_have_access_code", locale), *localeValue.SmsHaveAccessCode)
			}
			if localeValue.SmsMessageFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_message_format", locale), *localeValue.SmsMessageFormat)
			}
			if localeValue.SmsNumberCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_cancel", locale), *localeValue.SmsNumberCancel)
			}
			if localeValue.SmsNumberError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_error", locale), *localeValue.SmsNumberError)
			}
			if localeValue.SmsNumberFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_field_label", locale), *localeValue.SmsNumberFieldLabel)
			}
			if localeValue.SmsNumberFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_format", locale), *localeValue.SmsNumberFormat)
			}
			if localeValue.SmsNumberMessage != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_message", locale), *localeValue.SmsNumberMessage)
			}
			if localeValue.SmsNumberSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_submit", locale), *localeValue.SmsNumberSubmit)
			}
			if localeValue.SmsNumberTitle != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_number_title", locale), *localeValue.SmsNumberTitle)
			}
			if localeValue.SmsUsernameFormat != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sms_username_format", locale), *localeValue.SmsUsernameFormat)
			}
			if localeValue.SponsorBackLink != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_back_link", locale), *localeValue.SponsorBackLink)
			}
			if localeValue.SponsorCancel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_cancel", locale), *localeValue.SponsorCancel)
			}
			if localeValue.SponsorEmail != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_email", locale), *localeValue.SponsorEmail)
			}
			if localeValue.SponsorEmailError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_email_error", locale), *localeValue.SponsorEmailError)
			}
			if localeValue.SponsorInfoApproved != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_info_approved", locale), *localeValue.SponsorInfoApproved)
			}
			if localeValue.SponsorInfoDenied != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_info_denied", locale), *localeValue.SponsorInfoDenied)
			}
			if localeValue.SponsorInfoPending != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_info_pending", locale), *localeValue.SponsorInfoPending)
			}
			if localeValue.SponsorName != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_name", locale), *localeValue.SponsorName)
			}
			if localeValue.SponsorNameError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_name_error", locale), *localeValue.SponsorNameError)
			}
			if localeValue.SponsorNotePending != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_note_pending", locale), *localeValue.SponsorNotePending)
			}
			if localeValue.SponsorRequestAccess != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_request_access", locale), *localeValue.SponsorRequestAccess)
			}
			if localeValue.SponsorStatusApproved != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_status_approved", locale), *localeValue.SponsorStatusApproved)
			}
			if localeValue.SponsorStatusDenied != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_status_denied", locale), *localeValue.SponsorStatusDenied)
			}
			if localeValue.SponsorStatusPending != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_status_pending", locale), *localeValue.SponsorStatusPending)
			}
			if localeValue.SponsorSubmit != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsor_submit", locale), *localeValue.SponsorSubmit)
			}
			if localeValue.SponsorsError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsors_error", locale), *localeValue.SponsorsError)
			}
			if localeValue.SponsorsFieldLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.sponsors_field_label", locale), *localeValue.SponsorsFieldLabel)
			}
			if localeValue.TosAcceptLabel != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.tos_accept_label", locale), *localeValue.TosAcceptLabel)
			}
			if localeValue.TosError != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.tos_error", locale), *localeValue.TosError)
			}
			if localeValue.TosLink != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.tos_link", locale), *localeValue.TosLink)
			}
			if localeValue.TosText != nil {
				checks.append(t, "TestCheckResourceAttr", fmt.Sprintf("portal_template.locales.%s.tos_text", locale), *localeValue.TosText)
			}

		}
	}

	return checks
}
